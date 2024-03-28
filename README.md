# toy-k8s
(学習用)自作k8s

[Kubernetes](https://kubernetes.io) is a system for automating deployment, scaling, and management of containerized applications.
This is a repository to learn how the [components](https://kubernetes.io/docs/concepts/overview/components/) of a Kubernetes cluster works.

# Local up example
#### Install etcd
Older etcd is needed. I use etcd [v0.4.2](https://github.com/etcd-io/etcd/releases/tag/v0.4.2)

#### Build binaries
```
./src/scripts/build-go.sh
```

#### Run a local kubernetes cluster
```
(sudo) ./src/scripts/local-up-cluster.sh
```

#### Create a nginx pod
```json
{
  "id": "nginxController",
  "desiredState": {
    "replicas": 1,
    "replicasInSet": {"name": "nginx"},
    "taskTemplate": {
      "desiredState": {
        "manifest": {
          "containers": [{
            "name": "nginx",
            "image": "nginx",
            "ports": [{"containerPort": 80, "hostPort": 8888}]
          }]
        }
      },
      "labels": {"name": "nginx"}
    }
  },
  "labels": {"name": "nginx"}
}
```

Note: one replica is only created on local.

```
(sudo) ./bin/cloudcfg -h http://127.0.0.1:8080 -c examples/nginx-controller.json create /replicationControllers
```

#### Check your running containers
```
docker ps
```

#### Delete the nginx container
```
docker kill [CONTAINER]
```

#### Wait a few seconds, check containers again
```
docker ps
```

# References
- [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes) - [2c4b3a5](https://github.com/kubernetes/kubernetes/commit/2c4b3a5)
