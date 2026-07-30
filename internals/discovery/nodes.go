package discovery

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNodes(client kubernetes.Interface) (*corev1.NodeList, error) {
	return client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
}

func GetNodeEvent(client kubernetes.Interface) (*corev1.EventList, error) {
	return client.CoreV1().Events("").List(
		context.Background(),
		metav1.ListOptions{
			FieldSelector: "involvedObject.kind=Node",
		},
	)
}
