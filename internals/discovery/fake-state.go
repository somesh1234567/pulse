package discovery

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
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
			Capacity: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("2"),
				corev1.ResourceMemory: resource.MustParse("4Gi"),
			},
			Allocatable: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("1800m"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
			Conditions: []corev1.NodeCondition{
				{
					Status: corev1.ConditionTrue,
					Type:   corev1.NodeReady,
				},
				{
					Type:   corev1.NodeMemoryPressure,
					Status: corev1.ConditionTrue,
				},
			},
		},
	}

	event1 := &corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "worker-1-registered",
			Namespace: "default",
		},
		InvolvedObject: corev1.ObjectReference{
			Kind: "Node",
			Name: "worker-1",
		},
		Reason:  "RegisteredNode",
		Type:    corev1.EventTypeNormal,
		Message: "registering node",
		Count:   1,
	}

	event2 := &corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "worker-1-memory",
			Namespace: "default",
		},
		InvolvedObject: corev1.ObjectReference{
			Kind: "Node",
			Name: "worker-1",
		},
		Reason:  "NodeHasSufficientMemory",
		Type:    corev1.EventTypeNormal,
		Message: "normal event for memory",
		Count:   1,
	}

	event3 := &corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "worker-1-systemoom",
			Namespace: "default",
		},
		InvolvedObject: corev1.ObjectReference{
			Kind: "Node",
			Name: "worker-1",
		},
		Reason:  "SystemOOM",
		Type:    corev1.EventTypeWarning,
		Message: "System OOM encountered...",
		Count:   2,
	}

	pod1 := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "memory-hog",
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "sample-container",
					Image: "nginx-latest",
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceMemory: resource.MustParse("900Mi"),
							corev1.ResourceCPU:    resource.MustParse("250m"),
						},
					},
				},
			},
			NodeName: "worker-1",
		},
	}

	pod2 := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "well-behaved",
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			NodeName: "worker-1",
			Containers: []corev1.Container{
				{
					Name:  "sample-container",
					Image: "nginx:latest",
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceMemory: resource.MustParse("256Mi"),
							corev1.ResourceCPU:    resource.MustParse("250m"),
						},
						Limits: corev1.ResourceList{
							corev1.ResourceMemory: resource.MustParse("256Mi"), // limit set!
							corev1.ResourceCPU:    resource.MustParse("250m"),
						},
					},
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
	return fake.NewSimpleClientset(node1, node2, event1, event2, event3, pod1, pod2)
}
