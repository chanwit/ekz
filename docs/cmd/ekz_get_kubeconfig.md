## ekz get kubeconfig

Get kubeconfig

### Synopsis

This command obtains the KubeConfig of the EKS-D cluster and writes to the target file.

```
ekz get kubeconfig [flags]
```

### Examples

```
  # Get the KubeConfig from the cluster and write to $PWD/kubeconfig
  ekz get kubeconfig

  # Get the KubeConfig of the 'dev' cluster
  ekz get kubeconfig --name=dev

  # Get the KubeConfig of the 'dev' cluster (alternative syntax) 
  ekz get kubeconfig dev

  # Get the KubeConfig and writes to $HOME/.kube/config
  # Please note that this example overwrites the content of $HOME/.kube/config file.
  ekz get kubeconfig -o $HOME/.kube/config

```

### Options

```
  -h, --help            help for kubeconfig
      --name string     cluster name (default "ekz")
  -o, --output string   specify output file to write kubeconfig to (default "/home/chanwit/.kube/config")
```

### Options inherited from parent commands

```
      --provider string   cluster provider possible values: "ekz", "kind". env: EKZ_PROVIDER (default "ekz")
      --verbose           run verbosely
```

### SEE ALSO

* [ekz get](ekz_get.md)	 - Get properties of an EKS-D cluster

