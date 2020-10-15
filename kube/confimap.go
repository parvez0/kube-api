package kube

import (
	"context"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (c *KubeClientSet)UpdateCM(namespace string, cm *v12.ConfigMap, options v1.UpdateOptions) (*v12.ConfigMap, error) {
	client := c.GetCMClient(namespace)
	return client.Update(context.TODO(), cm, options)
}

func (c *KubeClientSet)PatchCM(namespace string, name string, patchType types.PatchType, data []byte, opts v1.PatchOptions, subresources string) (*v12.ConfigMap, error) {
	client := c.GetCMClient(namespace)
	if patchType == "" {
		patchType = types.StrategicMergePatchType
	}
	return client.Patch(context.TODO(), name, patchType, data, opts, subresources)
}