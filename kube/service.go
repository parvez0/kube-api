package kube

import (
	"context"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *KubeClientSet) GetService(ns string, svc string) (*v12.Service, error) {
	client := c.GetSVCClient(ns)
	return client.Get(context.TODO(), svc, v1.GetOptions{})
}
