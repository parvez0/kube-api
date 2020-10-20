package kubetests

import (
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

// TestCreateService will create a test verify the cluster ip if generated
func TestCreateService(t *testing.T)  {
	svc := &v12.Service{
		TypeMeta:   v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:         ResourceName,
			Namespace:    Namespace,
		},
		Spec: v12.ServiceSpec{
			Ports: []v12.ServicePort{
				{
					Port: int32(ServicePorts),
				},
			},
			Selector: map[string]string{
				"app": ResourceName,
			},
			Type: "ClusterIP",
		},
	}
	_, err := kube_cli.CreateService(Namespace, svc, v1.CreateOptions{})
	if err != nil{
		t.Errorf("failed to create service - %+v", err)
	}
	// waiting till the service gets created
	wa, err := kube_cli.WatchService(Namespace, v1.ListOptions{})
	if err != nil{
		t.Errorf("failed to setup watcher for service %+v", err)
	}
	for i := range wa.ResultChan(){
		svc := i.Object.(*v12.Service)
		got := svc.Spec.ClusterIP
		if got != ""{
			wa.Stop()
		}
	}
}

func TestDeleteService(t *testing.T)  {
	err := kube_cli.DeleteService(Namespace, ResourceName, v1.DeleteOptions{})
	if err != nil{
		t.Errorf("failed to delete service %+v", err)
	}
}

