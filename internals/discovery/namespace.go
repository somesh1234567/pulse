package discovery

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNamespaces(client kubernetes.Interface) (*corev1.NamespaceList, error) {
	return client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
}
