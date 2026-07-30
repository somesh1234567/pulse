package scanner

import (
	"fmt"
	"pulse/internals/discovery"
	"pulse/internals/rules"
	"pulse/internals/shared"
)

func Run() error {
	// client, err := kubernetes.NewClient()
	client := discovery.FakeClient()
	state, err := discovery.BuildState(client)
	if err != nil {
		return err
	}
	var findings []shared.Finding
	nodeRule := rules.NodeHealthRule{}
	memoryRule := rules.NodeMemoryRule{}
	findings = append(findings, nodeRule.Run(state)...)
	findings = append(findings, memoryRule.Run(state)...)
	fmt.Println("Scan Summary..")
	fmt.Println("===============")
	fmt.Println()
	fmt.Printf("Found %d nodes:\n", len(state.Nodes.Items))
	for _, node := range state.Nodes.Items {
		fmt.Printf(". %s\n", node.Name)
	}
	fmt.Println()
	fmt.Printf("Found %d namespaces:\n", len(state.Namespaces.Items))
	for _, ns := range state.Namespaces.Items {
		fmt.Printf(". %s\n", ns.Name)
	}
	fmt.Println()
	fmt.Printf("Found %d pods:\n", len(state.Pods.Items))
	for _, pod := range state.Pods.Items {
		fmt.Printf(". %s\n", pod.Name)
	}

	// FINDINGS //
	fmt.Println()
	fmt.Println("Findings")
	fmt.Println("========")
	if len(findings) == 0 {
		fmt.Println("No issues found")
		return nil
	}
	//fmt.Println(findings)
	for _, finding := range findings {
		fmt.Printf("[%s] %s\n", finding.Severity, finding.Rule)
		fmt.Printf("Resource : %s\n", finding.Resource)
		fmt.Printf("Message  : %s\n", finding.Message)
		fmt.Printf("Diagnosis: %s\n", finding.Diagnosis)
		fmt.Printf("Related Events to help: %s\n", finding.RelatedEvents)
		fmt.Println()
	}
	return nil
}
