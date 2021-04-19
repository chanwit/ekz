---
description: An easy way to run EKS clusters on your desktop
---

# The EKZ Project

`ekz` is an opinionated Kubernetes distribution built using binaries from [the AWS EKS Distro](https://distro.eks.amazonaws.com/) \(EKS-D\). It aims to be the easiest way to run EKS clusters on desktop.

What is EKS-D?

> _EKS-D provides the same software that has enabled tens of thousands of Kubernetes clusters on Amazon EKS._

`ekz` aims at solving the EKS compatibility problem for developers as much as possible by implanting EKS-D binaries to KIND and the k0s project, so that we can easily spin EKS-compatible clusters up to test our Kubernetes applications.

## Architecture

The architecture of EKZ has been designed to support EKS-D in multiple implementations, called providers. Currently, we ship the [k0s](https://github.com/k0sproject/k0s)-based \(EKZ provider\), and [KIND](https://github.com/kubernetes-sigs/kind/)-based \(KIND provider\) implementations. A provider can be specified when creating a cluster, or via the `EKZ_PROVIDER` variable.

Here's EKS-D versions supported by EKZ.

| EKS-D version       | EKZ provider  | KIND provider|
| ------------------- | :-----------: | :----------: |
| v1.18.9-eks-1-18-1  | ✓             | ✓            |
| v1.18.9-eks-1-18-3  | ✓             | ✓            |
| v1.19.6-eks-1-19-1  | ✓             | ✓            |
| v1.19.6-eks-1-19-3  | ✓             | ✓            |

## Getting Started

`ekz` creates a cluster for you inside a Docker container on your laptop. You can start a cluster with or without using the CLI.

### CLI Installation

You could install the CLI with one the following options.

#### Homebrew \(macOS & Linux\)

| Install | Upgrade |
|---------|---------|
|`brew install ekz-io/tap/ekz` | `brew upgrade ekz-io/tap/ekz`

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
controller   Ready    <none>   111s   v1.19.6-eks-1-19-3
```

### Without CLI

To use EKZ without using the CLI, please refer to this [document](without_cli.md).

## Features

### The EKZ provider

1. EKS-D binaries from v1.18.9-eks-1-18-{1,3}, and v1.19.6-eks-1-19-{1,3}
2. Packaged with k0s v0.13.0
3. Amazon Linux 2 base image
4. Enable network policy by default via the Calico CNI
5. Bundled with a local storage class

### The KIND provider

1. EKS-D binaries from v1.18.9-eks-1-18-{1,3}, and v1.19.6-eks-1-19-{1,3}
2. Using KIND v0.10
3. Packaged using KIND v1.18.x and v1.19.x node images
4. Enable network policy by default via the Calico CNI

