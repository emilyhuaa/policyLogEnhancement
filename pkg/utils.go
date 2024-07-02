package k8s

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
)

// ListPods lists Kubernetes Pods by namespace(s)
func ListPods(namespace string, client kubernetes.Interface) (*v1.PodList, error) {
	fmt.Println("Get Kubernetes Pods")
	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pods: %v\n", err)
		return nil, err
	}
	return pods, nil
}
