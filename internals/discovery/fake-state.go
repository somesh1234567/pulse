package discovery

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func FakeClient() kubernetes.Interface {
	node1 := &corev1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "worker-1",
		},
		Status: corev1.NodeStatus{
			Conditions: []corev1.NodeCondition{
				{
					Status: corev1.ConditionTrue,
					Type:   corev1.NodeReady,
				},
			},
		},
	}

	node2 := &corev1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "worker-2",
		},
		Status: corev1.NodeStatus{
			Conditions: []corev1.NodeCondition{
				{
					Status: corev1.ConditionFalse,
					Type:   corev1.NodeReady,
				},
			},
		},
	}
	return fake.NewSimpleClientset(node1, node2)
}
