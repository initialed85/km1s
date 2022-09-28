# km1s

`km1s` or `k-1s` literally "K minus S" is named in the theme of `k8s`, `k3s`, `k0s` etc but is a negative number to highlight that it's not
in fact Kubernetes.

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
