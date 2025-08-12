package main

import (
	"fmt"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

func main() {
	helm()
}

type HelmValues struct {
	OnboardingClusterKubeconfigSecretName string `yaml:"onboardingClusterKubeconfigSecretName"`
	PlatformClusterKubeconfigSecretName   string `yaml:"platformClusterKubeconfigSecretName"`
}

func helm() {
	// Example values, replace with your own logic
	vals := HelmValues{
		// OnboardingClusterKubeconfigSecretName: "onboarding-cluster-kubeconfig",
		PlatformClusterKubeconfigSecretName: "platform-cluster-kubeconfig",
	}

	// Serialize values to a temporary YAML file
	tmpFile, err := os.CreateTemp("", "helm-values-*.yaml")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	enc := yaml.NewEncoder(tmpFile)
	if err := enc.Encode(vals); err != nil {
		panic(err)
	}
	tmpFile.Close()

	// Get chart path and output dir from args
	if len(os.Args) < 3 {
		panic("usage: main <chart-path> <output-dir>")
	}
	chartPath := os.Args[1]
	outputDir := os.Args[2]

	// Run helm template with output-dir and values file
	cmd := exec.Command(
		"helm", "template", "--output-dir", outputDir, chartPath,
		"--values", tmpFile.Name(),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running: helm template --output-dir %s %s --values %s\n", outputDir, chartPath, tmpFile.Name())
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
