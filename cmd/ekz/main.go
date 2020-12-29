package main

import (
	"github.com/chanwit/script"
	"log"
	"os"

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
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		script.Debug = verbose
		return nil
	},
	Short: "Command line utility for creating EKS-D clusters on desktop",
	Long:  "Command line utility for creating EKS-D clusters on desktop",
	Example: `  # Create cluster
  ekz create cluster
`,
}

var (
	verbose  bool
	provider string
	logger   fluxlog.Logger = stderrLogger{stderr: os.Stderr}
)

func init() {
	defaultProvider := os.Getenv("EKZ_PROVIDER")
	if defaultProvider == "" {
		defaultProvider = "ekz"
	}
	if defaultProvider != "ekz" &&
		defaultProvider != "kind" {
		logger.Failuref("EKZ_PROVIDER=%s is not supported. Possible values: 'ekz', 'kind'.", defaultProvider)
		os.Exit(1)
	}

	rootCmd.PersistentFlags().StringVar(&provider, "provider", defaultProvider, "cluster provider possible values: \"ekz\", \"kind\". env: EKZ_PROVIDER")
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