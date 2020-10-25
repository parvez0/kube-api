package kube

import (
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (c *KubeClientSet) GetCM(ns string, name string, opts v1.GetOptions) (*v12.ConfigMap, error) {
	client := c.GetCMClient(ns)

	return client.Get(c.Ctx, name, opts)
}

func (c *KubeClientSet) ListCM(ns string, opts v1.ListOptions) (*v12.ConfigMapList, error) {
	client := c.GetCMClient(ns)
	return client.List(c.Ctx, opts)
}

func (c *KubeClientSet) CreateCM(ns string, name string, cm *v12.ConfigMap, opts v1.CreateOptions) (*v12.ConfigMap, error) {
	client := c.GetCMClient(ns)
	return client.Create(c.Ctx, cm, opts)
}

func (c *KubeClientSet) UpdateCM(namespace string, cm *v12.ConfigMap, options v1.UpdateOptions) (*v12.ConfigMap, error) {
	client := c.GetCMClient(namespace)
	return client.Update(c.Ctx, cm, options)
}

func (c *KubeClientSet) PatchCM(namespace string, name string, patchType types.PatchType, data []byte, opts v1.PatchOptions, subresources string) (*v12.ConfigMap, error) {
	client := c.GetCMClient(namespace)
	if patchType == "" {
		patchType = types.StrategicMergePatchType
	}
	return client.Patch(c.Ctx, name, patchType, data, opts, subresources)
}
