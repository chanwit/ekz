## ekz

Command line utility for creating EKS-D clusters on desktop

### Synopsis

This program is a command line utility for creating and managing EKS-D clusters on desktop.
It currently supports clusters provided by EKZ (k0s-based) and KIND.
All EKS-D cluster is single-node and run inside Docker.

### Examples

```
  # Create an EKS-D cluster with the default provider
  ekz create cluster

  # Create a 1.18 EKS-D cluster with the default provider
  ekz create cluster --eksd-version=v1.18.9-eks-1-18-1

  # Create a 1.19 EKS-D cluster with the default provider
  ekz create cluster --eksd-version=v1.19.6-eks-1-19-1

  # Delete the default cluster
  ekz delete cluster

  # List all clusters
  ekz list clusters

  # List all clusters (shorter syntax)
  ekz ls

  # Obtain KubeConfig of the cluster and write to $HOME/.kube/config
  ekz get kubeconfig

  # Create, delete, get the kube config for the default KIND-based cluster
  ekz create cluster --provider=kind
  ekz get kubeconfig --provider=kind
  ekz delete cluster --provider=kind

```

### Options

```
  -h, --help              help for ekz
      --provider string   cluster provider possible values: "ekz", "kind". env: EKZ_PROVIDER (default "ekz")
      --verbose           run verbosely
```

### SEE ALSO

* [ekz create](ekz_create.md)	 - Create clusters
* [ekz delete](ekz_delete.md)	 - Delete clusters
* [ekz get](ekz_get.md)	 - Get properties of an EKS-D cluster
* [ekz list](ekz_list.md)	 - List clusters
* [ekz load](ekz_load.md)	 - Load artifacts into the cluster
* [ekz ui](ekz_ui.md)	 - Start the UI

