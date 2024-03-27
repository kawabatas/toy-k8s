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
)

// Implementation of RESTStorage for the api server.
type ControllerRegistryStorage struct {
	registry ControllerRegistry
}

func MakeControllerRegistryStorage(registry ControllerRegistry) apiserver.RESTStorage {
	return &ControllerRegistryStorage{
		registry: registry,
	}
}

func (storage *ControllerRegistryStorage) List(*url.URL) (interface{}, error) {
	// TODO
	return nil, nil
}

func (storage *ControllerRegistryStorage) Get(id string) (interface{}, error) {
	// TODO
	return nil, nil
}

func (storage *ControllerRegistryStorage) Delete(id string) error {
	// TODO
	return nil
}

func (storage *ControllerRegistryStorage) Extract(body string) (interface{}, error) {
	// TODO
	return nil, nil
}

func (storage *ControllerRegistryStorage) Create(controller interface{}) error {
	// TODO
	return nil
}

func (storage *ControllerRegistryStorage) Update(controller interface{}) error {
	// TODO
	return nil
}
