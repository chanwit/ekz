## ekz create cluster

Create a cluster

### Synopsis

The create sub-commands create EKS-D clusters.

```
ekz create cluster [flags]
```

### Examples

```
  # Create an EKS-D cluster with the default provider
  # The KubeConfig will be merged to $HOME/.kube/config. The default cluster name is 'ekz'.
  ekz create cluster

  # Create cluster and name it 'dev'
  ekz create cluster --name=dev

  # Create the 'dev' cluster (alternative syntax)
  ekz create cluster dev

  # Create an EKS-D cluster with the EKZ provider
  # This command creates an EKS-D-compatible K0s-based cluster.
  ekz --provider=ekz create cluster

  # Create an EKS-D cluster with the KIND provider
  # This command creates an EKS-D-compatible KIND cluster.
  ekz --provider=kind create cluster

  # Create an EKS-D cluster and write KubeConfig to $PWD/kubeconfig
  # If the file already exists, the new KubeConfig will be merged into it.
  ekz create cluster -o kubeconfig

  # Create EKS-D cluster with a specific version of EKS-D
  ekz create --eksd-version=v1.18.9-eks-1-18-1 cluster 

```

### Options

```
      --eksd-version string   specify a version of EKS-D (default "v1.18.9-eks-1-18-1")
  -h, --help                  help for cluster
      --name string           cluster name (default "ekz")
  -o, --output string         specify output file to write kubeconfig to (default "~/.kube/config")
```

### Options inherited from parent commands

```
      --provider string   cluster provider possible values: "ekz", "kind". env: EKZ_PROVIDER (default "ekz")
      --verbose           run verbosely
```

### SEE ALSO

* [ekz create](ekz_create.md)	 - Create clusters

