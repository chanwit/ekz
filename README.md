# ekz
A Kubernetes distribution built on EKS-D and k0s

`ekz` is an opinionated Kubernetes distribution built on top of the [AWS EKS Distro](https://distro.eks.amazonaws.com/) (EKS-D) and the [k0s](https://k0sproject.io/) project. 

What is EKS-D?

  > *EKS-D provides the same software that has enabled tens of thousands of Kubernetes clusters on Amazon EKS.*

`ekz` replaces Kubernetes components of `k0s` with binaries from EKS-D, resulting in an easy-to-use single-binary EKS-compatible Kubernetes for development and testing purpose.

## Getting Started
`ekz` is intended to run inside its Docker container, on your laptop, with the following command:

### Linux & macOS

```sh
$ docker run -d --name ekz-controller \
   --hostname controller \
   --privileged -v /var/lib/ekz \
   -p 6443:6443 quay.io/chanwit/ekz:v1.18.9-eks-1-18-1.3
```

in case you'd like to try the dev version (from the main branch):

```sh
$ docker run -d --name ekz-controller \
   --hostname controller \
   --privileged -v /var/lib/ekz \
   -p 6443:6443 quay.io/chanwit/ekz:v1.18.9-eks-1-18-1.dev
```
### Windows

`ekz` also runs on Windows if you've got Docker Desktop installed.

#### PowerShell

```sh
$ docker run -d --name ekz-controller `
   --hostname controller `
   --privileged -v /var/lib/ekz `
   -p 6443:6443 quay.io/chanwit/ekz:v1.18.9-eks-1-18-1.dev
```
#### Command Prompt

```sh
$ docker run -d --name ekz-controller ^
   --hostname controller ^
   --privileged -v /var/lib/ekz ^
   -p 6443:6443 quay.io/chanwit/ekz:v1.18.9-eks-1-18-1.dev
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

## Features

  1. EKS-D binaries from v1.18.9-eks-1-18-1
  2. Packaged with k0s v0.8.1
  3. Amazon Linux 2 base image
