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
	"net/url"

	"github.com/kawabatas/toy-k8s/pkg/apiserver"
	"github.com/kawabatas/toy-k8s/pkg/client"
)

// TaskRegistryStorage implements the RESTStorage interface in terms of a TaskRegistry
type TaskRegistryStorage struct {
	registry      TaskRegistry
	containerInfo client.ContainerInfo
	scheduler     Scheduler
}

func MakeTaskRegistryStorage(registry TaskRegistry, containerInfo client.ContainerInfo, scheduler Scheduler) apiserver.RESTStorage {
	return &TaskRegistryStorage{
		registry:      registry,
		containerInfo: containerInfo,
		scheduler:     scheduler,
	}
}

func (storage *TaskRegistryStorage) List(url *url.URL) (interface{}, error) {
	// TODO
	return nil, nil
}

func (storage *TaskRegistryStorage) Get(id string) (interface{}, error) {
	// TODO
	return nil, nil
}

func (storage *TaskRegistryStorage) Delete(id string) error {
	// TODO
	return nil
}

func (storage *TaskRegistryStorage) Extract(body string) (interface{}, error) {
	// TODO
	return nil, nil
}

func (storage *TaskRegistryStorage) Create(task interface{}) error {
	// TODO
	return nil
}

func (storage *TaskRegistryStorage) Update(task interface{}) error {
	// TODO
	return nil
}
