package kube

import (
	"context"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func (c *KubeClientSet) GetService(ns string, svc string) (*v12.Service, error) {
	client := c.GetSVCClient(ns)
	return client.Get(c.Ctx, svc, v1.GetOptions{})
}

func (c *KubeClientSet) CreateService(ns string, svc *v12.Service, opts v1.CreateOptions) (*v12.Service, error) {
	client := c.GetSVCClient(ns)
	return client.Create(c.Ctx, svc, opts)
}

func (c *KubeClientSet) DeleteService(ns string, svc string, delOptions v1.DeleteOptions) error {
	client := c.GetSVCClient(ns)
	return client.Delete(c.Ctx, svc, delOptions)
}

func (c *KubeClientSet) WatchService(ns string, listOpts v1.ListOptions) (watch.Interface, error) {
	client := c.GetSVCClient(ns)
	return client.Watch(c.Ctx, listOpts)
}
