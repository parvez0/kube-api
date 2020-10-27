package kube

import (
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *KubeClientSet) GetNamespace(ns string) (*v12.Namespace, error) {
	client := c.GetNamespaceClient()
	return client.Get(c.Ctx, ns, v1.GetOptions{})
}

func (c *KubeClientSet) CreateNamespace(ns *v12.Namespace, opts v1.CreateOptions) (*v12.Namespace, error) {
	client := c.GetNamespaceClient()
	return client.Create(c.Ctx, ns, opts)
}

func (c *KubeClientSet) NamespaceList(opts v1.ListOptions) (*v12.NamespaceList, error) {
	nsClient := c.GetNamespaceClient()
	list, err := nsClient.List(c.Ctx, opts)
	return list, err
}

func (c *KubeClientSet) DeleteNamespace(ns string, opts v1.DeleteOptions) error {
	nsClient := c.GetNamespaceClient()
	return nsClient.Delete(c.Ctx, ns, opts)
}
