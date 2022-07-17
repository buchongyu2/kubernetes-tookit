package main

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := "~/.kube/config"
	if strings.HasPrefix(path, "~/") {
		dirname, _ := os.UserHomeDir()
		path = filepath.Join(dirname, path[2:])
	}
	fmt.Println("config_path: ", path)

	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		panic(err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pods, err := client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range pods.Items {
		fmt.Printf("命名空间: %v \n pod名字 %v\n pod_id: %v \n IP: %v \n", v.Namespace, v.Name, v.Status.PodIP, v.Status.HostIP)
	}
}
