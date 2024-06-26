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
// The kubelet binary is responsible for maintaining a set of containers on a particular host VM.
// It sync's data from both configuration file as well as from a quorum of etcd servers.
// It then queries Docker to see what is currently running.  It synchronizes the configuration data,
// with the running set of containers by starting or stopping Docker containers.
package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"time"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/kawabatas/toy-k8s/pkg/kubelet"
	"github.com/kawabatas/toy-k8s/third_party/github.com/coreos/go-etcd/etcd"
)

var (
	file               = flag.String("config", "", "Path to the config file")
	etcdServers        = flag.String("etcd_servers", "", "Url of etcd servers in the cluster")
	syncFrequency      = flag.Duration("sync_frequency", 10*time.Second, "Max seconds between synchronizing running containers and config")
	fileCheckFrequency = flag.Duration("file_check_frequency", 20*time.Second, "Seconds between checking file for new data")
	httpCheckFrequency = flag.Duration("http_check_frequency", 20*time.Second, "Seconds between checking http for new data")
	manifestURL        = flag.String("manifest_url", "", "URL for accessing the container manifest")
	address            = flag.String("address", "127.0.0.1", "The address for the info server to serve on")
	port               = flag.Uint("port", 10250, "The port for the info server to serve on")
	hostnameOverride   = flag.String("hostname_override", "", "If non-empty, will use this string as identification instead of the actual hostname.")
)

const dockerBinary = "/usr/bin/docker"

func main() {
	flag.Parse()

	// Set up logger for etcd client
	etcd.SetLogger(log.New(os.Stderr, "etcd ", log.LstdFlags))

	endpoint := "unix:///var/run/docker.sock"
	dockerClient, err := docker.NewClient(endpoint)
	if err != nil {
		log.Fatal("Couldn't connnect to docker.")
	}

	hostname := []byte(*hostnameOverride)
	if string(hostname) == "" {
		hostname, err = exec.Command("hostname", "-f").Output()
		if err != nil {
			log.Fatalf("Couldn't determine hostname: %v", err)
		}
	}

	myKubelet := kubelet.Kubelet{
		DockerClient:       dockerClient,
		FileCheckFrequency: *fileCheckFrequency,
		SyncFrequency:      *syncFrequency,
		HTTPCheckFrequency: *httpCheckFrequency,
		Hostname:           string(hostname),
	}
	myKubelet.RunKubelet(*file, *manifestURL, *etcdServers, *address, *port)
}
