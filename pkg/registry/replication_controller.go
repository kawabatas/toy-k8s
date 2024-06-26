/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package registry

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/kawabatas/toy-k8s/pkg/api"
	"github.com/kawabatas/toy-k8s/pkg/client"
	"github.com/kawabatas/toy-k8s/pkg/util"
	"github.com/kawabatas/toy-k8s/third_party/github.com/coreos/go-etcd/etcd"
)

// ReplicationManager is responsible for synchronizing ReplicationController objects stored in etcd
// with actual running tasks.
// TODO: Remove the etcd dependency and re-factor in terms of a generic watch interface
type ReplicationManager struct {
	etcdClient  *etcd.Client
	kubeClient  client.ClientInterface
	taskControl TaskControlInterface
	updateLock  sync.Mutex
}

// An interface that knows how to add or delete tasks
// created as an interface to allow testing.
type TaskControlInterface interface {
	createReplica(controllerSpec api.ReplicationController)
	deleteTask(taskID string) error
}

type RealTaskControl struct {
	kubeClient client.ClientInterface
}

func (r RealTaskControl) createReplica(controllerSpec api.ReplicationController) {
	labels := controllerSpec.DesiredState.TaskTemplate.Labels
	if labels != nil {
		labels["replicationController"] = controllerSpec.ID
	}
	task := api.Task{
		JSONBase: api.JSONBase{
			ID: fmt.Sprintf("%x", rand.Int()),
		},
		DesiredState: controllerSpec.DesiredState.TaskTemplate.DesiredState,
		Labels:       controllerSpec.DesiredState.TaskTemplate.Labels,
	}
	_, err := r.kubeClient.CreateTask(task)
	if err != nil {
		log.Printf("%#v\n", err)
	}
}

func (r RealTaskControl) deleteTask(taskID string) error {
	return r.kubeClient.DeleteTask(taskID)
}

func MakeReplicationManager(etcdClient *etcd.Client, kubeClient client.ClientInterface) *ReplicationManager {
	return &ReplicationManager{
		kubeClient: kubeClient,
		etcdClient: etcdClient,
		taskControl: RealTaskControl{
			kubeClient: kubeClient,
		},
	}
}

func (rm *ReplicationManager) Synchronize() {
	for {
		response, err := rm.etcdClient.Get("/registry/controllers", false, false)
		if err != nil {
			log.Printf("Synchronization error %#v", err)
		}
		// TODO(bburns): There is a race here, if we get a version of the controllers, and then it is
		// updated, its possible that the watch will pick up the change first, and then we will execute
		// using the old version of the controller.
		// Probably the correct thing to do is to use the version number in etcd to detect when
		// we are stale.
		// Punting on this for now, but this could lead to some nasty bugs, so we should really fix it
		// sooner rather than later.
		if response != nil && response.Node != nil && response.Node.Nodes != nil {
			for _, value := range response.Node.Nodes {
				var controllerSpec api.ReplicationController
				err := json.Unmarshal([]byte(value.Value), &controllerSpec)
				if err != nil {
					log.Printf("Unexpected error: %#v", err)
					continue
				}
				err = rm.syncReplicationController(controllerSpec)
				if err != nil {
					log.Printf("Error synchronizing: %#v", err)
				}
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func (rm *ReplicationManager) syncReplicationController(controllerSpec api.ReplicationController) error {
	rm.updateLock.Lock()
	taskList, err := rm.kubeClient.ListTasks(controllerSpec.DesiredState.ReplicasInSet)
	if err != nil {
		return err
	}
	filteredList := rm.filterActiveTasks(taskList.Items)
	diff := len(filteredList) - controllerSpec.DesiredState.Replicas
	if diff < 0 {
		diff *= -1
		log.Printf("Too few replicas, creating %d\n", diff)
		for i := 0; i < diff; i++ {
			rm.taskControl.createReplica(controllerSpec)
			log.Printf("creating...")
			time.Sleep(time.Second) // wait 1 second, to check multi replicas starting easily.
		}
	} else if diff > 0 {
		log.Print("Too many replicas, deleting")
		for i := 0; i < diff; i++ {
			rm.taskControl.deleteTask(filteredList[i].ID)
		}
	}
	rm.updateLock.Unlock()
	return nil
}

func (rm *ReplicationManager) filterActiveTasks(tasks []api.Task) []api.Task {
	var result []api.Task
	for _, value := range tasks {
		if !strings.Contains(value.CurrentState.Status, "Exit") {
			result = append(result, value)
		}
	}
	return result
}

func (rm *ReplicationManager) WatchControllers() {
	watchChannel := make(chan *etcd.Response)
	go util.Forever(func() { rm.etcdClient.Watch("/registry/controllers", 0, true, watchChannel, nil) }, 0)
	for {
		watchResponse := <-watchChannel
		if watchResponse == nil {
			time.Sleep(time.Second * 10)
			continue
		}
		log.Printf("Got watch: %#v", watchResponse)
		controller, err := rm.handleWatchResponse(watchResponse)
		if err != nil {
			log.Printf("Error handling data: %#v, %#v", err, watchResponse)
			continue
		}
		rm.syncReplicationController(*controller)
	}
}

func (rm *ReplicationManager) handleWatchResponse(response *etcd.Response) (*api.ReplicationController, error) {
	if response.Action == "set" {
		if response.Node != nil {
			var controllerSpec api.ReplicationController
			err := json.Unmarshal([]byte(response.Node.Value), &controllerSpec)
			if err != nil {
				return nil, err
			}
			return &controllerSpec, nil
		} else {
			return nil, fmt.Errorf("response node is null %#v", response)
		}
	}
	return nil, nil
}
