package kubeconfig

import (
	"fmt"
	"net/url"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func RewriteKubeConfigForEKZ(clusterName string, content string, port string) (string, error) {
	obj, err := yaml.Parse(content)
	if err != nil {
		return content, err
	}

	server, err := obj.Pipe(yaml.Lookup("clusters", "[name=local]", "cluster", "server"))
	if err != nil {
		return content, err
	}

	serverUrl := yaml.GetValue(server)
	u, err := url.Parse(serverUrl)
	if err != nil {
		return content, err
	}

	rewroteUrlScalar := yaml.NewScalarRNode(fmt.Sprintf("%s://%s:%s", u.Scheme, u.Hostname(), port))
	clusterNameScalar := yaml.NewScalarRNode(clusterName)
	contextScalar := yaml.NewScalarRNode("ekz-" + clusterName)
	userNameScalar := yaml.NewScalarRNode(clusterName + "-user")

	err = obj.PipeE(
		yaml.Tee(
			yaml.Lookup("clusters", "[name=local]"),
			yaml.Tee(
				yaml.Lookup("cluster", "server"),
				yaml.Set(rewroteUrlScalar),
			),
			yaml.SetField("name", clusterNameScalar),
		),
		yaml.Tee(
			yaml.Lookup("contexts", "[name=Default]"),
			yaml.Tee(yaml.Lookup("context"), yaml.SetField("cluster", clusterNameScalar)),
			yaml.Tee(yaml.Lookup("context"), yaml.SetField("user", userNameScalar)),
			yaml.SetField("name", contextScalar),
		),
		yaml.Tee(
			yaml.Lookup("users", "[name=user]"),
			yaml.SetField("name", userNameScalar),
		),
		yaml.SetField("current-context", contextScalar),
	)
	if err != nil {
		return content, err
	}

	return obj.String()
}
