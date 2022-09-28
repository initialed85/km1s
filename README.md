# km1s

`km1s` or `k-1s` literally "kay-minus-ones" is named in the theme of [k0s](https://github.com/k0sproject/k0s)
, [k3s](https://github.com/k3s-io/k3s), [k8s](https://github.com/kubernetes/kubernetes) etc but is a negative number to highlight that it's
not in fact Kubernetes.

## Goal

Provide a lightweight framework that takes some complexity out of deploying software services across many nodes (specifically targeting
small services and numerous nodes with constrained resources).

## Concept

After failing to get `k3s` to do anything useful on my old Raspberry Pi 3's because 1 GB of RAM wasn't enough and SD cards are too slow I
decided I need something that would allow me to do Kubernetes-like stuff but in a much lighter way.

Here's my thinking on the approach:

- Provide a Kubernetes pod-like abstraction called a container
    - This is where the containerization will happen
- Provide a Kubernetes container-like abstraction called a process
    - This will permit running one or more executables inside a container
- Provide a simple overlay network between all containers
    - This will permit any process to contact any other process regardless of container
- Provide a Kubernetes service-like abstraction called a service
    - This will employ DNS to hide the complexity of IP addresses and permit DNS-level load balancing
    - This may also need some sort of proxy / load balancer at the network level for per packet load balancing, sticky sessions etc
- Provide an artifact store for easy upload of executables to be run as processes inside a container
    - This will allow easy distribution of executables
- Provide a distributed multi-node network abstraction
    - This will allow containers on any node to contact containers on any other node
- Provide a distributed multi-node scheduler system
    - This will employ the artifact store, orchestrate containers and processes and expose everything via services hiding the complexity of
      deploying software on multiple nodes

It'd be great if this whole thing never needs to be more than a single Go executable (per node); but I'm a bit concerned about the network
load balancer side of it to be honest- surely I can't hope to reimplement all the learning behind the likes
of [nginx](https://github.com/nginx/nginx) and [haproxy]()https://github.com/haproxy/haproxy).

## Status / TODOs

- (In progress) Container abstraction
    - (In progress) Network isolation
        - Change usage of [iproute2](https://github.com/shemminger/iproute2) to a suitable library
        - Fix intermittent leaky namespaces / resources etc
    - (TODO) UTS / PID isolation
    - (TODO) UID / mount isolation
    - (TODO) Cgroups for resource limits
- (In progress) Process abstraction
    - Write the Manager and expose it as an excutable command
    - Fix intermittent leaky child processes
    - Improve STDIN / STDOUT / STDERR interface for logging and shells
- (TODO) Services
    - Build a dynamic DNS server on top of [miekg/dns](https://github.com/miekg/dns)
- (TODO) Artifact store
- (TODO) Distributed networking

## Usage

This thing is so WIP right now you can't really use it; currently the best demo is a test:

```shell
make test-shell
go test -v ./pkg/container/ -run TestContainer
```

Here's what that test did under the hood:

- Spin up a Docker container to run everything in
- In the default `netns` (let's call it the host)
    - Create a `bridge` and give it an IP
- In a newly created `netns` (let's cal it the container)
    - Create a `veth` with a leg on the host and leg on the container
    - Add it to the `bridge` on the host
    - Give it an IP in the container
    - Add a default route for the container
    - Spin up an instance of `netcat` listening for some UDP
- Back on the host
    - Send some UDP to the contaner
- Back in the container
    - Check the container logs to ensure `netcat` saw the UDP

When you're all done:

```shell
exit
make test-down
```

## Development

Also very WIP; if you want to run the tests:

```shell
make test
```

They're super flaky though, no idea why- I get a little bit more success with:

```shell
make test-shell

# as many times you like while developing
go test -v ./...

# when you're done
exit
make test-down
```

Keep an eye out, my deferred cleanups aren't consistent for some reason and often resources leak (not beyond the test container though in my
experience).

## References

- https://baturin.org/docs/iproute2/#ip-link-add-bridge
- https://etherarp.net/connecting-network-namespaces-with-veth/index.html
- https://github.com/vishvananda/netns
- https://blog.scottlowe.org/2013/09/04/introducing-linux-network-namespaces/
- https://stackoverflow.com/questions/32492349/how-to-list-all-network-links-in-all-namespaces-in-linux
- https://songrgg.github.io/programming/linux-namespace-part01-uts-pid/
- https://serverfault.com/questions/662699/how-to-configure-a-linux-network-namespace-that-allows-udp-broadcast
- https://github.com/coredns/coredns
- https://github.com/miekg/dns
