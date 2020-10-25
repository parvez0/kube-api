package test
//
//import (
//	"fmt"
//	"github.com/parvez0/kube-api/kube"
//	"k8s.io/api/autoscaling/v2beta2"
//	"k8s.io/apimachinery/pkg/api/resource"
//	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
//)
//
//func main() {
//	cli := kube.CreateClient("production-backup")
//	ns := "production-bots"
//	bot := "x1532195756954"
//	var minRep int32 = 1
//	var cpuThreshold int32 = 1200
//	var cpuTarget int64 = 200
//	var memoryThreshold int32 = 1000
//	var memoryTarget int64 = 250
//	dsd := v2beta2.HorizontalPodAutoscaler{
//		ObjectMeta: v13.ObjectMeta{
//			Name:      bot,
//			Namespace: ns,
//		},
//		Spec: v2beta2.HorizontalPodAutoscalerSpec{
//			ScaleTargetRef: v2beta2.CrossVersionObjectReference{
//				Kind:       "Deployment",
//				Name:       bot,
//				APIVersion: "apps/v1",
//			},
//			MinReplicas: &minRep,
//			MaxReplicas: 3,
//			Metrics: []v2beta2.MetricSpec{
//				{
//					Type: "Resource",
//					Resource: &v2beta2.ResourceMetricSource{
//						Name: "cpu",
//						Target: v2beta2.MetricTarget{
//							Type:               v2beta2.UtilizationMetricType,
//							AverageValue:       resource.NewQuantity(cpuTarget, "m"),
//							AverageUtilization: &cpuThreshold,
//						},
//					},
//				},
//				{
//					Type: "Resource",
//					Resource: &v2beta2.ResourceMetricSource{
//						Name: "memory",
//						Target: v2beta2.MetricTarget{
//							AverageValue:       resource.NewQuantity(memoryTarget, "Mi"),
//							AverageUtilization: &memoryThreshold,
//						},
//					},
//				},
//			},
//		},
//	}
//	hpa, err := cli.CreateHpa(ns, &dsd, v13.CreateOptions{})
//	if err != nil {
//		fmt.Errorf("failed to get HPA - %+v", err)
//	}
//	fmt.Printf("list hpa - %+v", *hpa)
//}
