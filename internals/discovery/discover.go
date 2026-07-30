package discovery

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

func BuildState(client kubernetes.Interface) (*ClusterState, error) {
	//get the nodes
	nodes, err := GetNodes(client)
	if err != nil {
		return nil, err
	}
	// get the node event
	events, _ := GetNodeEvent(client)
	nodeEvents := make(map[string][]corev1.Event)

	for _, event := range events.Items {
		nodeName := event.InvolvedObject.Name
		nodeEvents[nodeName] = append(nodeEvents[nodeName], event)
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
		NodeEvent:  nodeEvents,
		Namespaces: namespaces,
		Pods:       pods,
	}, nil
}
