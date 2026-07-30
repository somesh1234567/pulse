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
		// fmt.Println(node.Name)
		events := state.NodeEvent[node.Name]
		for _, condition := range node.Status.Conditions {
			if condition.Type == corev1.NodeMemoryPressure &&
				condition.Status == corev1.ConditionTrue {

				var relatedEvents []string

				for _, event := range events {
					relatedEvents = append(relatedEvents,
						fmt.Sprintf("[%s] %s: %s",
							event.Type,
							event.Reason,
							event.Message,
						),
					)
				}

				findings = append(findings, shared.Finding{
					Rule:          r.Name(),
					Severity:      "Critical",
					Category:      "NodeMemory",
					Resource:      node.Name,
					Message:       "Node is Under High Memory Pressure",
					RelatedEvents: relatedEvents,
					Diagnosis:     "probably need to check if any pod has status of OOM Killed which can be taking up memory",
				})
			}

		}
	}
	return findings
}
