package kube

import (
	"context"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *KubeClientSet) NamespaceList() (*v12.NamespaceList, error) {
	nsClient := c.GetNamespaceClient()
	list, err := nsClient.List(context.TODO(), v1.ListOptions{})
	return list, err
}
