package main

import (
	"context"
	"flag"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
)

var clientset *kubernetes.Clientset

func main()  {
	k8sconfig := flag.String("k8sconfig","C:\\Users\\fangyulong\\go\\go_basic_grammar\\1-BeginLearn\\10 go-k8s\\k8s.config","kubernetes config file path")
	flag.Parse()
	config , err := clientcmd.BuildConfigFromFlags("",*k8sconfig)
	if err != nil {
		log.Println(err)
	}
	clientset , err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("connect k8s success")
	}

	//获取POD
	pods,err := clientset.CoreV1().Pods("kube-system").List(context.Background(),metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}

	for _,v := range pods.Items{
		fmt.Println(v.Name)
		fmt.Println(v.CreationTimestamp)
		fmt.Println(v.Labels)
		fmt.Println(v.Namespace)
		fmt.Println(v.Status.HostIP)
		fmt.Println(v.Status.PodIP)
		fmt.Println(v.Status.StartTime)
		fmt.Println(v.Status.Phase)
		fmt.Println(v.Status.ContainerStatuses[0].RestartCount)   //重启次数
		fmt.Println(v.Status.ContainerStatuses[0].Image) //获取重启时间
	}


	//获取NODE
	fmt.Println("##################")
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(),metav1.ListOptions{})


	for _,v :=range nodes.Items{
		fmt.Println(v.Name)
		fmt.Println(v.CreationTimestamp)    //加入集群时间
		fmt.Println(v.Status.NodeInfo)
		fmt.Println(v.Status.Conditions[len(v.Status.Conditions)-1].Type)
		fmt.Println(v.Status.Allocatable.Memory().String())
	}

}