package scanner

import (
	"fmt"
	"pulse/internals/discovery"
	"pulse/internals/kubernetes"
)

func Run() error {
	client, err := kubernetes.NewClient()
	if err != nil {
		return err
	}
	state, err := discovery.BuildState(client)
	if err != nil {
		return err
	}
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
	return nil
}
