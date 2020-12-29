package main

import (
	"fmt"
	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/sets"
	"text/tabwriter"
)

var listClusterCmd = &cobra.Command{
	Use:     "cluster",
	Aliases: []string{"clusters"},
	Short:   "List clusters",
	Long:    "List all EKS-D clusters.",
	RunE:    listClusterCmdRun,
}

func init() {
	listCmd.AddCommand(listClusterCmd)
}

// How to heuristically detect a kind / ekz cluster
// "io.x-k8s.kind.cluster": "ekz"

// How to heuristically detect a ekz cluster
// io.x-k8s.ekz.cluster=ekz

func listClusters(clusterLabelKey string) ([]string, error) {
	output := script.Var()
	err := script.Exec("docker",
		"ps",
		"-a", // show stopped nodes
		// filter for nodes with the cluster label
		"--filter", "label="+clusterLabelKey,
		// format to include the cluster name
		"--format", fmt.Sprintf(`{{.Label "%s"}}`, clusterLabelKey),
	).To(output)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list clusters")
	}
	return sets.NewString(output.Lines()...).List(), nil
}

func listClusterCmdRun(cmd *cobra.Command, args []string) error {
	eksClusters, err := listClusters("io.x-k8s.ekz.cluster")
	if err != nil {
		return err
	}
	kindClusters, err := listClusters("io.x-k8s.kind.cluster")
	if err != nil {
		return err
	}
	w := tabwriter.NewWriter(script.Stdout(), 0, 2, 3, ' ', 0)
	fmt.Fprintf(w, "NAME\tPROVIDER\n")
	for _, c := range eksClusters {
		fmt.Fprintf(w, "%s\tekz\n", c)
	}
	for _, c := range kindClusters {
		fmt.Fprintf(w, "%s\tkind\n", c)
	}
	w.Flush()

	return nil
}
