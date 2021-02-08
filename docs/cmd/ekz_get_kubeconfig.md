## ekz get kubeconfig

Get kubeconfig

### Synopsis

This command obtains the KubeConfig of the EKS-D cluster and writes to the target file.

```
ekz get kubeconfig [flags]
```

### Examples

```
  # Get the KubeConfig from the cluster and write to $HOME/.kube/config
  ekz get kubeconfig

  # Get the KubeConfig of the 'dev' cluster
  ekz get kubeconfig --name=dev

  # Get the KubeConfig of the 'dev' cluster (alternative syntax) 
  ekz get kubeconfig dev

  # Get the KubeConfig and save to $PWD/kubeconfig
  ekz get kubeconfig -o $PWD/kubeconfig

  # Get the KubeConfig from the default KIND-based cluster
  ekz get kubeconfig --provider=kind

  # Get the KubeConfig from the 'dev' KIND-based cluster
  ekz get kubeconfig --provider=kind --name=dev

```

### Options

```
  -h, --help            help for kubeconfig
      --name string     cluster name (default "ekz")
  -o, --output string   specify output file to write kubeconfig to (default "~/.kube/config")
```

### Options inherited from parent commands

```
      --provider string   cluster provider possible values: "ekz", "kind". env: EKZ_PROVIDER (default "ekz")
      --verbose           run verbosely
```

### SEE ALSO

* [ekz get](ekz_get.md)	 - Get properties of an EKS-D cluster

