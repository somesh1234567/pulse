package discovery

import "k8s.io/client-go/kubernetes"

func BuildState(client *kubernetes.Clientset) (*ClusterState, error) {
	namespaces, err := GetNamespaces(client)
	if err != nil {
		return nil, err
	}

	return &ClusterState{
		Namespaces: namespaces,
	}, nil
}
