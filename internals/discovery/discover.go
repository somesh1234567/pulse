package discovery

import "k8s.io/client-go/kubernetes"

func BuildState(client *kubernetes.Clientset) (*ClusterState, error) {
	//get the nodes
	nodes, err := GetNodes(client)
	if err != nil {
		return nil, err
	}
	// get the namespaces
	namespaces, err := GetNamespaces(client)
	if err != nil {
		return nil, err
	}
	// get the pods list
	pods, err := GetPods(client)
	if err != nil {
		return nil, err
	}

	return &ClusterState{
		Nodes:      nodes,
		Namespaces: namespaces,
		Pods:       pods,
	}, nil
}
