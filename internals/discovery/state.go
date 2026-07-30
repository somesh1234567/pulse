package discovery

import corev1 "k8s.io/api/core/v1"

type ClusterState struct {
	Nodes      *corev1.NodeList
	NodeEvent  map[string][]corev1.Event
	Namespaces *corev1.NamespaceList
	Pods       *corev1.PodList
}
