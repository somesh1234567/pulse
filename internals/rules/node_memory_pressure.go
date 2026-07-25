package rules

import (
	"fmt"
	"pulse/internals/discovery"
	"pulse/internals/shared"

	corev1 "k8s.io/api/core/v1"
)

type NodeMemoryRule struct{}

func (r NodeMemoryRule) Name() string {
	return "NodeUnderMemoryPressure"
}

func (r NodeMemoryRule) Run(state *discovery.ClusterState) []shared.Finding {
	var findings []shared.Finding
	for _, node := range state.Nodes.Items {
		fmt.Println(node.Name)
		for _, condition := range node.Status.Conditions {
			if condition.Type == corev1.NodeReady &&
				condition.Status != corev1.ConditionTrue {

				findings = append(findings, shared.Finding{
					Rule:      r.Name(),
					Severity:  "Critical",
					Category:  "NodeHealth",
					Resource:  node.Name,
					Message:   "Node is NotReady",
					Diagnosis: "Inspect kubelet and node health.",
				})
			}

		}
	}
	return findings
}
