package main

import (
	"context"
	"fmt"
	"os"

	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Set up the Kubernetes client
	//kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", "/tmp/kubeconfig.yaml")
	if err != nil {
		fmt.Printf("Failed to load kubeconfig: %v\n", err)
		os.Exit(1)
	}

	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	fmt.Printf("Failed to create clientset: %v\n", err)
	// 	os.Exit(1)
	// }

	apiextClient, err := apiextensionsclientset.NewForConfig(config)
	if err != nil {
		fmt.Printf("Failed to create API Extensions client: %v\n", err)
		os.Exit(1)
	}

	// Specify the name of the CRD you want to load
	crdName := "cumulocityiotedges.edge.cumulocity.com"

	// Retrieve the CRD object
	crd, err := apiextClient.ApiextensionsV1().CustomResourceDefinitions().Get(context.TODO(), crdName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Failed to get CRD: %v\n", err)
		os.Exit(1)
	}

	// Use the loaded CRD object as needed
	fmt.Printf("Loaded CRD: %s\n\n", crd.Name)
	fmt.Printf("API Version: %s\n\n", crd.Spec.Group+"/"+crd.Spec.Versions[0].Name)
	// Add any other processing or handling of the CRD object

	//	deployedValues := crd.Spec.Versions[0].Schema.OpenAPIV3Schema.Properties

	// Print the deployed values
	//	fmt.Printf("Deployed values of CRD: %+v\n\n", deployedValues["mongodb"])

}
