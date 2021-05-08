package main

import (
	"log"
	"os"

	"github.com/chanwit/script"
	fluxlog "github.com/fluxcd/flux2/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var VERSION = "0.0.0-dev.0"

var rootCmd = &cobra.Command{
	Use:           "ekz",
	Version:       VERSION,
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		script.Debug = verbose
	},
	Short: "Command line utility for creating EKS-D clusters on desktop",
	Long: `This program is a command line utility for creating and managing EKS-D clusters on desktop.
It currently supports clusters provided by EKZ (k0s-based) and KIND.
All EKS-D cluster is single-node and run inside Docker.`,
	Example: `  # Create an EKS-D cluster with the default provider
  ekz create cluster

  # Create a 1.18 EKS-D cluster with the default provider
  ekz create cluster --eksd-version=v1.18.9-eks-1-18-1

  # Create a 1.19 EKS-D cluster with the default provider
  ekz create cluster --eksd-version=v1.19.6-eks-1-19-1

  # Delete the default cluster
  ekz delete cluster

  # List all clusters
  ekz list clusters

  # List all clusters (shorter syntax)
  ekz ls

  # Obtain KubeConfig of the cluster and write to $HOME/.kube/config
  ekz get kubeconfig

  # Create, delete, get the kube config for the default KIND-based cluster
  ekz create cluster --provider=kind
  ekz get kubeconfig --provider=kind
  ekz delete cluster --provider=kind
`,
}

var (
	verbose  bool
	provider string
	logger   fluxlog.Logger = stderrLogger{stderr: os.Stderr}
)

var (
	enableExperimental string = os.Getenv("EKZ_EXPERIMENTAL")
)

func init() {
	defaultProvider := os.Getenv("EKZ_PROVIDER")
	if defaultProvider == "" {
		defaultProvider = "ekz"
	}
	if defaultProvider != "ekz" && defaultProvider != "kind" {
		logger.Failuref("EKZ_PROVIDER=%s is not supported. Possible values: 'ekz', 'kind'.", defaultProvider)
		os.Exit(1)
	}

	rootCmd.PersistentFlags().StringVarP(&provider, "provider", "p", defaultProvider, "cluster provider possible values: \"ekz\", \"kind\". env: EKZ_PROVIDER")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "run verbosely")
}

func main() {
	generateDocs()
	if err := rootCmd.Execute(); err != nil {
		logger.Failuref("%v", err)
		os.Exit(1)
	}
}

func generateDocs() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "docgen" {
		rootCmd.DisableAutoGenTag = true
		err := doc.GenMarkdownTree(rootCmd, "./docs/cmd")
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}
