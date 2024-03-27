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
	"errors"

	"github.com/kawabatas/toy-k8s/pkg/api"
	"github.com/kawabatas/toy-k8s/third_party/github.com/coreos/go-etcd/etcd"
)

// TODO: Need to add a reconciler loop that makes sure that things in tasks are reflected into
//       kubelet (and vice versa)

// EtcdClient is an injectable interface for testing.
type EtcdClient interface {
	AddChild(key, data string, ttl uint64) (*etcd.Response, error)
	Get(key string, sort, recursive bool) (*etcd.Response, error)
	Set(key, value string, ttl uint64) (*etcd.Response, error)
	Create(key, value string, ttl uint64) (*etcd.Response, error)
	Delete(key string, recursive bool) (*etcd.Response, error)
	// I'd like to use directional channels here (e.g. <-chan) but this interface mimics
	// the etcd client interface which doesn't, and it doesn't seem worth it to wrap the api.
	Watch(prefix string, waitIndex uint64, recursive bool, receiver chan *etcd.Response, stop chan bool) (*etcd.Response, error)
}

// EtcdRegistry is an implementation of both ControllerRegistry and TaskRegistry which is backed with etcd.
type EtcdRegistry struct {
	etcdClient      EtcdClient
	machines        []string
	manifestFactory ManifestFactory
}

// MakeEtcdRegistry creates an etcd registry.
// 'client' is the connection to etcd
// 'machines' is the list of machines
// 'scheduler' is the scheduling algorithm to use.
func MakeEtcdRegistry(client EtcdClient, machines []string) *EtcdRegistry {
	registry := &EtcdRegistry{
		etcdClient: client,
		machines:   machines,
	}
	registry.manifestFactory = &BasicManifestFactory{}
	return registry
}

func makeTaskKey(machine, taskID string) string {
	return "/registry/hosts/" + machine + "/tasks/" + taskID
}

func makeContainerKey(machine string) string {
	return "/registry/hosts/" + machine + "/kubelet"
}

func makeControllerKey(id string) string {
	return "/registry/controllers/" + id
}

func (registry *EtcdRegistry) ListTasks(query *map[string]string) ([]api.Task, error) {
	// TODO
	return nil, nil
}

func (registry *EtcdRegistry) GetTask(taskID string) (*api.Task, error) {
	// TODO
	return nil, nil
}

func (registry *EtcdRegistry) CreateTask(machineIn string, task api.Task) error {
	// TODO
	return nil
}

func (registry *EtcdRegistry) UpdateTask(task api.Task) error {
	return errors.New("not implemented")
}

func (registry *EtcdRegistry) DeleteTask(taskID string) error {
	// TODO
	return nil
}

func (registry *EtcdRegistry) ListControllers() ([]api.ReplicationController, error) {
	// TODO
	return nil, nil
}

func (registry *EtcdRegistry) GetController(controllerID string) (*api.ReplicationController, error) {
	// TODO
	return nil, nil
}

func (registry *EtcdRegistry) CreateController(controller api.ReplicationController) error {
	// TODO : check for existence here and error.
	return registry.UpdateController(controller)
}

func (registry *EtcdRegistry) UpdateController(controller api.ReplicationController) error {
	// TODO
	return nil
}

func (registry *EtcdRegistry) DeleteController(controllerID string) error {
	// TODO
	return nil
}
