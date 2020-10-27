package kube

import (
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func (c *KubeClientSet) GetPvc(ns string, svc string) (*v12.PersistentVolumeClaim, error) {
	client := c.GetPVCClient(ns)
	return client.Get(c.Ctx, svc, v1.GetOptions{})
}

func (c *KubeClientSet) CreatePvc(ns string, svc *v12.PersistentVolumeClaim, opts v1.CreateOptions) (*v12.PersistentVolumeClaim, error) {
	client := c.GetPVCClient(ns)
	return client.Create(c.Ctx, svc, opts)
}

func (c *KubeClientSet) DeletePvc(ns string, svc string, delOptions v1.DeleteOptions) error {
	client := c.GetPVCClient(ns)
	return client.Delete(c.Ctx, svc, delOptions)
}

func (c *KubeClientSet) WatchPvc(ns string, listOpts v1.ListOptions) (watch.Interface, error) {
	client := c.GetPVCClient(ns)
	return client.Watch(c.Ctx, listOpts)
}

