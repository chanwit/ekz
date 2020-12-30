package main

import (
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"text/tabwriter"

	"github.com/chanwit/script"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/sets"
)

var listClusterCmd = &cobra.Command{
	Use:     "cluster",
	Aliases: []string{"clusters"},
	Short:   "List clusters",
	Long:    "List cluster command list all EKS-D clusters.",
	Example: `  # List all EKS-D clusters
  ekz list clusters

  # List all EKS-D clusters (shorter syntax)
  ekz ls
`,
	RunE: listClusterCmdRun,
}

func init() {
	listCmd.AddCommand(listClusterCmd)
}

// How to heuristically detect a kind / ekz cluster
// "io.x-k8s.kind.cluster": "ekz"

// How to heuristically detect a ekz cluster
// io.x-k8s.ekz.cluster=ekz
const (
	EKZClusterLabel  = "io.x-k8s.ekz.cluster"
	KINDClusterLabel = "io.x-k8s.kind.cluster"
)

func listClusters(clusterLabelKey string) ([]string, error) {
	output := script.Var()
	err := script.Exec("docker", "ps",
		"-a", // show stopped nodes
		// filter for nodes with the cluster label
		"--filter", "label="+clusterLabelKey,
		// format to include the cluster name
		"--format", fmt.Sprintf(`{{.Label "%s"}}`, clusterLabelKey)).
		To(output)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list clusters")
	}

	return sets.NewString(output.Lines()...).List(), nil
}

func listClusterCmdRun(cmd *cobra.Command, args []string) error {
	config, err := clientcmd.LoadFromFile(clientcmd.RecommendedHomeFile)
	if err != nil {
		return err
	}

	eksClusters, err := listClusters(EKZClusterLabel)
	if err != nil {
		return err
	}
	kindClusters, err := listClusters(KINDClusterLabel)
	if err != nil {
		return err
	}
	w := tabwriter.NewWriter(script.Stdout(), 0, 2, 3, ' ', 0)
	fmt.Fprintf(w, "CLUSTER-NAME\tPROVIDER\tACTIVE\n")
	for _, c := range eksClusters {
		fmt.Fprintf(w, "%s\tekz\t", c)
		active := " "
		if config.CurrentContext == "ekz-"+c {
			active = "  *"
		}
		fmt.Fprintf(w, "%s\n", active)
	}
	for _, c := range kindClusters {
		fmt.Fprintf(w, "%s\tkind\t", c)
		active := " "
		if config.CurrentContext == "kind-"+c {
			active = "  *"
		}
		fmt.Fprintf(w, "%s\n", active)
	}
	w.Flush()

	return nil
}
