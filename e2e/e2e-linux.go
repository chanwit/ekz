package main

import (
	"github.com/chanwit/script"
)

func main() {
	script.Debug = true
	for _, provider := range []string{"ekz", "kind"} {
		for eksdVersion, expected := range map[string]string{
			"v1.20": "v1.20.4-eks-1-20-1",
			"v1.19": "v1.19.8-eks-1-19-4",
			"v1.18": "v1.18.16-eks-1-18-5",
		} {
			if err := script.Run("ekz", "create", "cluster", "--provider="+provider, "--eksd-version="+eksdVersion); err != nil {
				panic(err)
			}

			kubeletVersion := script.Var()
			if err := script.Exec("kubectl", "get", "nodes", "-oyaml").
				Exec("yq", "e", ".items[0].status.nodeInfo.kubeletVersion", "-").To(kubeletVersion); err != nil {
				panic(err)
			}

			if err := script.Run("ekz", "delete", "cluster", "--provider="+provider); err != nil {
				panic(err)
			}

			if kubeletVersion.String() != expected {
				panic("expected " + expected)
			}
		}
	}
}
