# Using EKZ Without CLI

If you don't want to install the CLI, you could also start a cluster using one of `ekz` containers.

## macOS & Linux

```bash
$ docker run -d --name ekz-controller \
   --hostname controller \
   --privileged -v /var/lib/ekz \
   -p 6443:6443 quay.io/ekz-io/ekz:v1.18.9-eks-1-18-1.6
```

in case you'd like to try the dev version \(from the main branch\):

```bash
$ docker run -d --name ekz-controller \
   --hostname controller \
   --privileged -v /var/lib/ekz \
   -p 6443:6443 quay.io/chanwit/ekz:v1.18.9-eks-1-18-1.dev
```

## Windows

`ekz` also runs on Windows if you've got Docker Desktop installed.

## PowerShell

```bash
$ docker run -d --name ekz-controller `
   --hostname controller `
   --privileged -v /var/lib/ekz `
   -p 6443:6443 quay.io/chanwit/ekz:v1.18.9-eks-1-18-1.dev
```

## Command Prompt

```bash
$ docker run -d --name ekz-controller ^
   --hostname controller ^
   --privileged -v /var/lib/ekz ^
   -p 6443:6443 quay.io/chanwit/ekz:v1.18.9-eks-1-18-1.dev
```

Then we can obtain KUBECONFIG by running:

```bash
$ docker exec ekz-controller cat /var/lib/ekz/pki/admin.conf > ~/.kube/config
```

