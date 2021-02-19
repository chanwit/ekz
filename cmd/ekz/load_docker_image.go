package main

import (
	"fmt"
	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var loadDockerImageCmd = &cobra.Command{
	Use: "docker-image <IMAGE>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("name of image is required")
		}
		return nil
	},
	Short: "Load a Docker image",
	Long:  "Load docker-image command loads a Docker image into the EKS-D cluster.",
	Example: `  # Load the busybox:latest into the default EKZ cluster
  ekz load docker-image busybox:latest

  # Load the busybox:latest into the default EKZ cluster
  ekz load docker-image busybox:latest --name=ekz

  # Load the busybox:latest into the default KIND cluster
  ekz --provider=kind load docker-image busybox:latest

  # Load the busybox:latest into the staging KIND cluster
  ekz --provider=kind load docker-image busybox:latest --name=staging
`,
	RunE: loadDockerImageCmdRun,
}

func init() {
	loadDockerImageCmd.Flags().StringVar(&clusterName, "name", "ekz", "cluster name")

	loadCmd.AddCommand(loadDockerImageCmd)
}

func loadDockerImageCmdRun(cmd *cobra.Command, args []string) error {
	logger.Successf("the default provider is: %s", provider)

	// How to load image, for example:
	// docker save chanwit/spring-boot-on-kubernetes-with-jib-example:2.0.0-alpha-1 |
	// docker exec -i ekz-controller-0 /var/lib/ekz/bin/ctr --address=/var/lib/ekz/run/containerd.sock -n k8s.io image import -
	var (
		ctr               string
		containerName     string
		containerdAddress string
	)

	if provider == "ekz" {
		ctr = "/var/lib/ekz/bin/ctr"
		containerName = fmt.Sprintf("%s-controller-0", clusterName)
		containerdAddress = "/var/lib/k0s/run/containerd.sock"
	} else if provider == "kind" {
		ctr = "ctr"
		containerName = fmt.Sprintf("%s-control-plane", clusterName)
		containerdAddress = "/run/containerd/containerd.sock"
	}

	imageName := args[0]

	out, err := script.
		Exec("docker", "save", imageName).
		Exec("docker", "exec", "-i", containerName,
			ctr,
			"--address", containerdAddress,
			"--namespace", "k8s.io",
			"image", "import", "-").
		CombinedOutput()
	if err != nil {
		logger.Failuref(string(out))
		return errors.Wrap(err, "error importing image")
	}

	logger.Successf("imaged %s loaded", imageName)
	return nil
}
