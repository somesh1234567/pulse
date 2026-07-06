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
	for _, ns := range state.Namespaces.Items {
		fmt.Printf(". %s\n", ns.Name)
	}
	return nil
}
