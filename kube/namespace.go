package kube

import (
	"context"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *KubeClientSet) NamespaceList(opts v1.ListOptions) (*v12.NamespaceList, error) {
	nsClient := c.GetNamespaceClient()
	list, err := nsClient.List(c.Ctx, v1.ListOptions{})
	return list, err
}

func (c *KubeClientSet) DeleteNamespace(ns string, opts v1.DeleteOptions) error {
	nsClient := c.GetNamespaceClient()
	return nsClient.Delete(c.Ctx, ns, opts)
}
