## ekz load docker-image

Load a Docker image

### Synopsis

Load docker-image command loads a Docker image into the EKS cluster.

```
ekz load docker-image [flags]
```

### Examples

```
  # Load the busybox:latest into the cluster
  ekz load docker-image busybox:latest

```

### Options

```
  -h, --help           help for docker-image
      --image string   a Docker image name to load
      --name string    cluster name (default "ekz")
```

### Options inherited from parent commands

```
      --provider string   cluster provider possible values: "ekz", "kind". env: EKZ_PROVIDER (default "ekz")
      --verbose           run verbosely
```

### SEE ALSO

* [ekz load](ekz_load.md)	 - Load artifacts into the cluster

