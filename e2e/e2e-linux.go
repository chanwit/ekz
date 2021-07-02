package main

import (
	"github.com/chanwit/script"
)

const (
	V1_18_Image = "v1.18.16-eks-1-18-6"
	V1_19_Image = "v1.19.12-eks-1-19-5"
	V1_20_Image = "v1.20.7-eks-1-20-2"

	LastestImage = V1_20_Image
)

func main() {
	script.Debug = true
	for _, provider := range []string{"ekz", "kind"} {
		for eksdVersion, expected := range map[string]string{
			"v1.20": V1_20_Image,
			"v1.19": V1_19_Image,
			"v1.18": V1_18_Image,
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
