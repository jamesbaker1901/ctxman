package cmd

import (
	"fmt"
	"os/exec"
)

func (e *Environment) SetKubernetes(ns string) error {
	if e.Kubernetes.Namespace != "" || ns != "" {
		fmt.Println("Changing to Kubernetes cluster:", e.Kubernetes.Cluster)
		fmt.Println("Setting Kubernetes namespace:", e.Kubernetes.Namespace)

		var namespace string

		if ns == "" {
			namespace = e.Kubernetes.Namespace
		} else {
			namespace = ns
		}
		cmd := exec.Command("kubectl", "config", "set-context", e.Kubernetes.Cluster, "--namespace", namespace)
		err := cmd.Run()
		if err != nil {
			return err
		}
		cmd = exec.Command("kubectl", "config", "use-context", e.Kubernetes.Cluster)
		err = cmd.Run()
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Changing to Kubernetes cluster:", e.Kubernetes.Cluster)
		cmd := exec.Command("kubectl", "config", "use-context", e.Kubernetes.Cluster)
		err := cmd.Run()
		if err != nil {
			return err
		}

	}

	return nil
}
