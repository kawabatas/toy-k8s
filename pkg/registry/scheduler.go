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
	"math/rand"

	"github.com/kawabatas/toy-k8s/pkg/api"
)

// Scheduler is an interface implemented by things that know how to schedule tasks onto machines.
type Scheduler interface {
	Schedule(api.Task) (string, error)
}

// RandomScheduler choses machines uniformly at random.
type RandomScheduler struct {
	machines []string
	random   rand.Rand
}

func MakeRandomScheduler(machines []string, random rand.Rand) Scheduler {
	return &RandomScheduler{
		machines: machines,
		random:   random,
	}
}

func (s *RandomScheduler) Schedule(task api.Task) (string, error) {
	return s.machines[s.random.Int()%len(s.machines)], nil
}

// RoundRobinScheduler chooses machines in order.
type RoundRobinScheduler struct {
	machines     []string
	currentIndex int
}

func MakeRoundRobinScheduler(machines []string) Scheduler {
	return &RoundRobinScheduler{
		machines:     machines,
		currentIndex: 0,
	}
}

func (s *RoundRobinScheduler) Schedule(task api.Task) (string, error) {
	// TODO
	return "", nil
}

type FirstFitScheduler struct {
	machines []string
	registry TaskRegistry
}

func MakeFirstFitScheduler(machines []string, registry TaskRegistry) Scheduler {
	return &FirstFitScheduler{
		machines: machines,
		registry: registry,
	}
}

func (s *FirstFitScheduler) Schedule(task api.Task) (string, error) {
	// TODO
	return "", nil
}