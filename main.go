package test
//
//import (
//	"fmt"
//	"github.com/parvez0/kube-api/kube"
//	v1 "k8s.io/api/apps/v1"
//)
//
//
//func main() {
//	cli := kube.CreateClient("production")
//	wat, err := cli.WatchStatefulSets("esstack")
//	if err != nil{
//		fmt.Printf("failed - %+v", err)
//	}
//	for i := range wat.ResultChan(){
//		ss := i.Object.(*v1.StatefulSet)
//		fmt.Printf("got watch result - %+v\n", ss.Status.ReadyReplicas)
//	}
//}
