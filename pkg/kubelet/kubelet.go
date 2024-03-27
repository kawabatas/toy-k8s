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

// Package kubelet is ...
package kubelet

import (
	"sync"
	"time"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/kawabatas/toy-k8s/pkg/registry"
)

// State, sub object of the Docker JSON data
type State struct {
	Running bool
}

// Interface for testability
type DockerInterface interface {
	ListContainers(options docker.ListContainersOptions) ([]docker.APIContainers, error)
	InspectContainer(id string) (*docker.Container, error)
	CreateContainer(docker.CreateContainerOptions) (*docker.Container, error)
	StartContainer(id string, hostConfig *docker.HostConfig) error
	StopContainer(id string, timeout uint) error
}

// The main kubelet implementation
type Kubelet struct {
	Client             registry.EtcdClient
	DockerClient       DockerInterface
	FileCheckFrequency time.Duration
	SyncFrequency      time.Duration
	HTTPCheckFrequency time.Duration
	pullLock           sync.Mutex
	Hostname           string
}

// Starts background goroutines. If file, manifest_url, or address are empty,
// they are not watched. Never returns.
func (sl *Kubelet) RunKubelet(file, manifest_url, etcd_servers, address string, port uint) {
	// TODO
}
