package main

import (
	"fmt"
	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var loadDockerImageCmd = &cobra.Command{
	Use:   "docker-image",
	Args:  cobra.MaximumNArgs(1),
	Short: "Load a Docker image",
	Long:  "Load docker-image command loads a Docker image into the EKS cluster.",
	Example: `  # Load the busybox:latest into the cluster
  ekz load docker-image busybox:latest
`,
	RunE: loadDockerImageCmdRun,
}

var (
	imageName string
)

func init() {
	loadDockerImageCmd.Flags().StringVar(&clusterName, "name", "ekz", "cluster name")
	loadDockerImageCmd.Flags().StringVar(&imageName, "image", "", "a Docker image name to load")

	loadCmd.AddCommand(loadDockerImageCmd)
}

func loadDockerImageCmdRun(cmd *cobra.Command, args []string) error {
	if provider != "ekz" {
		return errors.Errorf("not implemented yet for provider: %s", provider)
	}

	// How to load image, for example:
	// docker save chanwit/spring-boot-on-kubernetes-with-jib-example:2.0.0-alpha-1 |
	// docker exec -i ekz-controller-0 /var/lib/ekz/bin/ctr --address=/var/lib/ekz/run/containerd.sock -n k8s.io image import -
	if imageName == "" && len(args) == 0 {
		return errors.New("error no Docker image name is specified")
	}
	if imageName == "" && len(args) == 1 {
		imageName = args[0]
	}
	containerName := fmt.Sprintf("%s-controller-0", clusterName)

	out, err := script.Exec("docker", "save", imageName).
		Exec("docker", "exec", "-i", containerName,
			"/var/lib/ekz/bin/ctr",
			"--address=/var/lib/ekz/run/containerd.sock",
			"-n", "k8s.io",
			"image", "import", "-").
		CombinedOutput()
	if err != nil {
		fmt.Print(string(out))
		return errors.Wrap(err, "error importing image")
	}

	return nil
}
