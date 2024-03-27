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
	"sync"

	"github.com/kawabatas/toy-k8s/pkg/api"
	"github.com/kawabatas/toy-k8s/pkg/client"
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
	// TODO
}

func (r RealTaskControl) deleteTask(taskID string) error {
	// TODO
	return nil
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

func (rm *ReplicationManager) WatchControllers() {
	// TODO
}

func (rm *ReplicationManager) Synchronize() {
	// TODO
}
