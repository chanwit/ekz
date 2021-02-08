## ekz delete cluster

Delete a cluster

### Synopsis

The delete sub-commands delete EKS-D clusters.

```
ekz delete cluster [flags]
```

### Examples

```
  # Delete the cluster, the default name is 'ekz'
  ekz delete cluster

  # Delete the 'dev' cluster
  ekz delete cluster dev

  # Delete the 'dev' cluster (alternative syntax)
  ekz delete cluster --name=dev

  # Delete the cluster created by the EKZ provider
  ekz --provider=ekz delete cluster

  # Delete the cluster created by the KIND provider
  ekz --provider=kind delete cluster

```

### Options

```
  -h, --help            help for cluster
      --name string     cluster name (default "ekz")
  -o, --output string   specify output file to write kubeconfig to (default "kubeconfig")
```

### Options inherited from parent commands

```
      --provider string   cluster provider possible values: "ekz", "kind". env: EKZ_PROVIDER (default "ekz")
      --verbose           run verbosely
```

### SEE ALSO

* [ekz delete](ekz_delete.md)	 - Delete clusters

