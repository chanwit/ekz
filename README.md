# ekz
Kubernetes built on EKS Distro and K0s

`ekz` is an opinionated Kubernetes distribution built on top of the [AWS EKS Distro](https://distro.eks.amazonaws.com/) (EKS-D) and the [K0s](https://k0sproject.io/) projects.
`ekz` replaces Kubernetes components of K0s with binaries from EKS-D, resulting in an easy-to-use single-binary EKS-compatible Kubernetes for development and testing purpose.

`ekz` is intended to run inside its Docker container, on your laptop, with the following command:

```
$ docker run --name ekz-controller \
   --hostname controller \
   --privileged -v /var/lib/ekz \
   -p 6443:6443 chanwit/ekz:v1.18.9-eks-1-18-1
```

Then we can obtain KUBECONFIG by running:

```sh
$ docker exec ekz-controller cat /var/lib/ekz/pki/admin.conf > ~/.kube/config
```

Please wait for a couple of minutes and an EKS-D cluster will be ready on your laptop.

```sh
$ kubectl get nodes
NAME         STATUS   ROLES    AGE   VERSION
controller   Ready    <none>   42s   v1.18.9-eks-1-18-1
```
