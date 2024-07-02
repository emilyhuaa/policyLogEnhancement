package pkg

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ListPods lists Kubernetes Pods by namespace(s)
func ListPods(namespace string, client kubernetes.Interface) (*v1.PodList, error) {
	fmt.Println("Get Kubernetes Pods")
	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pods: %v", err)
		return nil, err
	}
	return pods, nil
}

// ListNamespaces lists all Namespaces
func ListNamespaces(client kubernetes.Interface) (*v1.NamespaceList, error) {
	fmt.Println("Get Kubernetes Namespaces")
	namespaces, err := client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting namespaces: %v", err)
		return nil, err
	}
	return namespaces, nil
}

func ListPodIPAddresses(namespace string, client kubernetes.Interface) ([]string, error) {
	fmt.Println("Get Kubernetes Pod IP addresses")
	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pods: %v", err)
		return nil, err
	}

	var podIPAddresses []string
	for _, pod := range pods.Items {
		podIPAddresses = append(podIPAddresses, pod.Status.PodIP)
	}

	return podIPAddresses, nil
}
