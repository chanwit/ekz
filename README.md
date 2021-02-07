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

#### Homebrew (macOS & Linux)
```sh
brew install ekz-io/tap/ekz
```

#### CURL One-liner (macOS & Linux)
```sh
curl -sSL https://bit.ly/install-ekz | bash
```

#### Wget One-liner (macOS & Linux)
```sh
wget -qO- https://bit.ly/install-ekz | bash
```

#### Scoop (Windows)
```
scoop bucket add ekz-io https://github.com/ekz-io/scoop-ekz.git
scoop install ekz-io/ekz
```

#### Chocolatey (Windows)
```
choco install -y ekz
```

Then you can start your first EKS-D cluster using the following command:
```
ekz create cluster
```

You can also use the KIND provider, so that your EKS-D clusters will be KIND-compatible.
To use the KIND provider, you can use either flag `--provider=kind` or `export EKZ_PROVIDER=kind`. 

Here' the example of using the `--provider=kind` flag:
```
ekz create cluster --provider=kind
```

In case you'd like to use KIND as the default provider, it's better to set the EKZ_PROVIDER environment variable:
```
export EKZ_PROVIDER=kind
ekz create cluster
```

### Without CLI

If you don't want to install the CLI, you could also start a cluster using one of `ekz` containers.

#### macOS & Linux

```sh
$ docker run -d --name ekz-controller \
   --hostname controller \
   --privileged -v /var/lib/ekz \
   -p 6443:6443 quay.io/ekz-io/ekz:v1.18.9-eks-1-18-1.6
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
  1. Packaged with k0s v0.10
  1. Amazon Linux 2 base image
  1. Enable network policy by default via the Calico CNI

The KIND provider

  1. EKS-D binaries from v1.18.9-eks-1-18-1
  1. Using KIND v0.10
  1. Packaged using KIND v1.18.15 node image
  1. Enable network policy by default via the Calico CNI
