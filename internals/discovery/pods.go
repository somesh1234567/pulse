package discovery

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetPods(client *kubernetes.Clientset) (*corev1.PodList, error) {
	return client.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
}
