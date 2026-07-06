package discovery

import corev1 "k8s.io/api/core/v1"

type ClusterState struct {
	Namespaces *corev1.NamespaceList
}
