package kubetests

import (
	v12 "k8s.io/api/apps/v1"
	v13 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

// TestCreate will create a dummy statefulset
func TestCreateStatefulSet(t *testing.T) {
	var replica int32 = 1
	// statefulset spec for running a single container of nginx
	ss := v12.StatefulSet{
		ObjectMeta: v1.ObjectMeta{
			Name: ResourceName,
			Namespace: Namespace,
		},
		Spec: v12.StatefulSetSpec{
			Replicas: &replica,
			Selector: &v1.LabelSelector{
				MatchLabels: map[string]string{ "app": ResourceName },
			},
			Template: v13.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: map[string]string{ "app": ResourceName },
				},
				Spec:       v13.PodSpec{
					Containers: []v13.Container{
						{
							Name: ResourceName,
							Image: "nginx",
						},
					},
				},
			},
		},
	}
	_, err := kube_cli.CreateStatefulSet(Namespace, &ss, v1.CreateOptions{})
	if err != nil{
		t.Errorf("failed to create statefulset - %+v", err)
		return
	}
}

// It will fetch the statefulset if any error occurs fails the test
func TestGetStatefulSet(t *testing.T)  {
	_, err := kube_cli.GetStatefulSet(Namespace, ResourceName, v1.GetOptions{})
	if err != nil{
		t.Errorf("failed to get statefulset - %+v", err)
		return
	}
}

// It will fetch the statefulset list if any error occurs fails the test
func TestListStatefulSet(t *testing.T)  {
	_, err := kube_cli.ListStatefulSet(Namespace, v1.ListOptions{})
	if err != nil{
		t.Errorf("failed to get statefulset - %+v", err)
		return
	}
}

// Restart will cover patch statefulset functionality as well
func TestRestartStatefulSets(t *testing.T)  {
	err := kube_cli.RestartStatefulSets(Namespace, ResourceName)
	if err != nil{
		t.Errorf("failed to restart - %+v", err)
		return
	}
}

// TestDelete will delete statefulsets
func TestStatefulSets(t *testing.T) {
	err := kube_cli.DeleteStatefulSets(Namespace, ResourceName, v1.DeleteOptions{})
	if err != nil{
		t.Errorf("failed to delete statefulset - %+v", err)
		return
	}
}