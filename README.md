# ekz
A Kubernetes distribution built on EKS-D

`ekz` is an opinionated Kubernetes distribution built using binaries from the [AWS EKS Distro](https://distro.eks.amazonaws.com/) (EKS-D).

What is EKS-D?

  > *EKS-D provides the same software that has enabled tens of thousands of Kubernetes clusters on Amazon EKS.*

## Architecture

The EKZ architecture supports EKS-D in multiple implementations, called providers. Currently, we ship the k0s-based, and KIND-based implementations. The provider can be specified when creating a cluster.

## Getting Started

`ekz` creates a cluster for you inside a Docker container on your laptop. You can start a cluster with or without using the CLI.

### With CLI

You could install the CLI with one the following options.

#### Homebrew
```sh
brew install ekz-io/tap/ekz
```

#### One-line Shell Script
```sh
curl -sSL https://bit.ly/install-ekz | sudo bash
```

Then you can start your first EKS-D cluster using the following command:
```
ekz create cluster
```

### Without CLI

If you don't want to install the CLI, you could also start a cluster using one of `ekz` containers.

#### Linux & macOS

```sh
$ docker run -d --name ekz-controller \
   --hostname controller \
   --privileged -v /var/lib/ekz \
   -p 6443:6443 quay.io/ekz-io/ekz:v1.18.9-eks-1-18-1.5
```

in case you'd like to try the dev version (from the main branch):

```sh
$ docker run -d --name ekz-controller \
   --hostname controller \
   --privileged -v /var/lib/ekz \
   -p 6443:6443 quay.io/chanwit/ekz:v1.18.9-eks-1-18-1.dev
```
#### Windows

`ekz` also runs on Windows if you've got Docker Desktop installed.

##### PowerShell

```sh
$ docker run -d --name ekz-controller `
   --hostname controller `
   --privileged -v /var/lib/ekz `
   -p 6443:6443 quay.io/chanwit/ekz:v1.18.9-eks-1-18-1.dev
```
##### Command Prompt

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

The EKZ provider

  1. EKS-D binaries from v1.18.9-eks-1-18-1
  1. Packaged with k0s v0.9.1
  1. Amazon Linux 2 base image

The KIND provider

  1. Using KIND 0.9
  1. Packaged using KIND 1.18.9 image
