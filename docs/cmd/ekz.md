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

  # Delete the default cluster
  ekz delete cluster

  # List all clusters
  ekz list clusters

  # List all clusters (shorter syntax)
  ekz ls

  # Obtain KubeConfig of the cluster and write to $PWD/kubeconfig
  ekz get kubeconfig

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

