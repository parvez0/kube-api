package kubetests

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestListNamespaces(t *testing.T)  {
	_, err := kube_cli.NamespaceList(v1.ListOptions{})
	if err != nil{
		t.Errorf("failed to list namespaces %+v", err)
	}
}