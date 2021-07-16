---
description: An easy way to run EKS clusters on your desktop
---

# The EKZ Project

`ekz` is an opinionated Kubernetes distribution built using binaries from [the AWS EKS Distro](https://distro.eks.amazonaws.com/) \(EKS-D\). It aims to be the easiest way to run EKS clusters on desktops (Linux, macOS and Windows).

What is EKS-D?

> _EKS-D provides the same software that has enabled tens of thousands of Kubernetes clusters on Amazon EKS._

`ekz` aims at solving the EKS compatibility problems for developers as much as possible by implanting EKS-D binaries to KinD and the k0s projects, so that we can easily spin EKS-compatible clusters up to test our Kubernetes applications.

## Architecture

The architecture of EKZ has been designed to support EKS-D in multiple implementations, called providers. Currently, we ship the [k0s](https://github.com/k0sproject/k0s)-based \(EKZ provider\), and [KinD](https://github.com/kubernetes-sigs/kind/)-based \(KinD provider\) implementations. A provider can be specified when creating a cluster, or via the `EKZ_PROVIDER` variable.

Here's EKS-D versions supported by EKZ.

| Kubernetes | EKS-D Release | EKZ provider  | KIND provider|
|------------|:-------------:| :-----------: | :----------: |
| 1-18       | 7             | ✓             | ✓            |
| 1-19       | 6             | ✓             | ✓            |
| 1-20       | 3             | ✓             | ✓            |

## Getting Started

`ekz` creates a cluster for you inside a Docker container on your laptop. You can start a cluster with or without using the CLI.

### CLI Installation

You could install the CLI with one the following options.

#### Homebrew \(macOS & Linux\)

##### Install:
```bash
brew install ekz-io/tap/ekz
```
##### Upgrade:
```bash
brew upgrade ekz-io/tap/ekz
```

#### CURL One-liner \(macOS & Linux\)

```bash
curl -sSL https://bit.ly/install-ekz | bash
```

#### Wget One-liner \(macOS & Linux\)

```bash
wget -qO- https://bit.ly/install-ekz | bash
```

#### Scoop \(Windows\)

```bash
scoop bucket add ekz-io https://github.com/ekz-io/scoop-ekz.git
scoop install ekz-io/ekz
```

#### Chocolatey \(Windows\)

```bash
choco install -y ekz
```

Then you can start your first EKS-D cluster using the following command:

```bash
ekz create cluster
```

You can also use the KIND provider, so that your EKS-D clusters will be KIND-compatible. To use the KIND provider, you can use either flag `--provider=kind` or `export EKZ_PROVIDER=kind`.

Here' the example of using the `--provider=kind` flag:

```bash
ekz create cluster --provider=kind
```

In case you'd like to use KIND as the default provider, it's better to set the EKZ\_PROVIDER environment variable:

```bash
export EKZ_PROVIDER=kind
ekz create cluster
```

Please wait for a couple of minutes and an EKS-D cluster will be ready on your laptop.

```bash
❯ kubectl get nodes
NAME         STATUS   ROLES    AGE    VERSION
controller   Ready    <none>   87s   v1.20.7-eks-1-20-3
```

### Without CLI

To use EKZ without using the CLI, please refer to this [document](without_cli.md).

## Features

### EKZ provider

1. EKS-D binaries
2. Packaged with the k0s skeleton 
3. Base image: Amazon Linux 2 
4. Enable network policy by default via Kube-Router (Calico is optional)
5. Bundled with a local storage class, and a load balancer

### KinD provider

1. EKS-D binaries
2. Using KinD v0.11 as the skeleton
3. Packaged using KinD v1.18, v1.19, v1.20 node images
4. Enable network policy by default via the Calico CNI
5. Bundled with a local storage class (from KinD), and a load balancer
