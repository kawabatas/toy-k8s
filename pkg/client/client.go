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

// A client for the Kubernetes cluster management API
// There are three fundamental objects
//
//	Task - A single running container
//	TaskForce - A set of co-scheduled Task(s)
//	ReplicationController - A manager for replicating TaskForces
package client

import (
	"net/http"

	"github.com/kawabatas/toy-k8s/pkg/api"
)

// ClientInterface holds the methods for clients of Kubenetes, an interface to allow mock testing
type ClientInterface interface {
	ListTasks(labelQuery map[string]string) (api.TaskList, error)
	GetTask(name string) (api.Task, error)
	DeleteTask(name string) error
	CreateTask(api.Task) (api.Task, error)
	UpdateTask(api.Task) (api.Task, error)

	GetReplicationController(name string) (api.ReplicationController, error)
	CreateReplicationController(api.ReplicationController) (api.ReplicationController, error)
	UpdateReplicationController(api.ReplicationController) (api.ReplicationController, error)
	DeleteReplicationController(string) error

	// Service has not implemented yet...
}

// AuthInfo is used to store authorization information
type AuthInfo struct {
	User     string
	Password string
}

// Client is the actual implementation of a Kubernetes client.
// Host is the http://... base for the URL
type Client struct {
	Host       string
	Auth       *AuthInfo
	httpClient *http.Client
}

// ListTasks takes a label query, and returns the list of tasks that match that query
func (client Client) ListTasks(labelQuery map[string]string) (api.TaskList, error) {
	// TODO
	return api.TaskList{}, nil
}

// GetTask takes the name of the task, and returns the corresponding Task object, and an error if it occurs
func (client Client) GetTask(name string) (api.Task, error) {
	// TODO
	return api.Task{}, nil
}

// DeleteTask takes the name of the task, and returns an error if one occurs
func (client Client) DeleteTask(name string) error {
	// TODO
	return nil
}

// CreateTask takes the representation of a task.  Returns the server's representation of the task, and an error, if it occurs
func (client Client) CreateTask(task api.Task) (api.Task, error) {
	// TODO
	return api.Task{}, nil
}

// UpdateTask takes the representation of a task to update.  Returns the server's representation of the task, and an error, if it occurs
func (client Client) UpdateTask(task api.Task) (api.Task, error) {
	// TODO
	return api.Task{}, nil
}

// GetReplicationController returns information about a particular replication controller
func (client Client) GetReplicationController(name string) (api.ReplicationController, error) {
	// TODO
	return api.ReplicationController{}, nil
}

// CreateReplicationController creates a new replication controller
func (client Client) CreateReplicationController(controller api.ReplicationController) (api.ReplicationController, error) {
	// TODO
	return api.ReplicationController{}, nil
}

// UpdateReplicationController updates an existing replication controller
func (client Client) UpdateReplicationController(controller api.ReplicationController) (api.ReplicationController, error) {
	// TODO
	return api.ReplicationController{}, nil
}

func (client Client) DeleteReplicationController(name string) error {
	// TODO
	return nil
}
