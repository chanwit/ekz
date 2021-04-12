package main

import (
	"github.com/chanwit/ekz/pkg/constants"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func expandKubeConfigFile() string {
	if kubeConfigFile == constants.BackTickHomeFile {
		kubeConfigFile = clientcmd.RecommendedHomeFile
	}
	return os.ExpandEnv(kubeConfigFile)
}
