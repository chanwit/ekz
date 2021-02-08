## ekz load docker-image

Load a Docker image

### Synopsis

Load docker-image command loads a Docker image into the EKS-D cluster.

```
ekz load docker-image <IMAGE> [flags]
```

### Examples

```
  # Load the busybox:latest into the default EKZ cluster
  ekz load docker-image busybox:latest

  # Load the busybox:latest into the default KIND cluster
  ekz --provider=kind load docker-image busybox:latest

  # Load the busybox:latest into the staging KIND cluster
  ekz --provider=kind load docker-image busybox:latest --cluster=staging

```

### Options

```
      --cluster string   cluster name (default "ekz")
  -h, --help             help for docker-image
```

### Options inherited from parent commands

```
      --provider string   cluster provider possible values: "ekz", "kind". env: EKZ_PROVIDER (default "ekz")
      --verbose           run verbosely
```

### SEE ALSO

* [ekz load](ekz_load.md)	 - Load artifacts into the cluster

