package main

import (
	"context"
	"fmt"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8stypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const namespace = "dev"
const deployment_name = "nginx"

func main() {
	ctx := context.Background()

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	data := fmt.Sprintf(`{"spec": {"template": {"metadata": {"annotations": {"kubectl.kubernetes.io/restartedAt": "%s"}}}}}`, time.Now().Format("20060102150405"))
	deployment, err := deploymentsClient.Patch(ctx, deployment_name, k8stypes.StrategicMergePatchType, []byte(data), v1.PatchOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Deployment successfully patched:", deployment.Name)
}
