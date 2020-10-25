package kube

import (
	"k8s.io/api/autoscaling/v2beta2"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *KubeClientSet) GetHpa(ns string, hpa string, opts v12.GetOptions) (*v2beta2.HorizontalPodAutoscaler, error) {
	cli := c.GetHPAClient(ns)
	return cli.Get(c.Ctx, hpa, opts)
}

func (c *KubeClientSet) ListHpa(ns string, opts v12.ListOptions) (*v2beta2.HorizontalPodAutoscalerList, error) {
	cli := c.GetHPAClient(ns)
	return cli.List(c.Ctx, opts)
}

func (c *KubeClientSet) CreateHpa(ns string, autoscaler *v2beta2.HorizontalPodAutoscaler, opts v12.CreateOptions) (*v2beta2.HorizontalPodAutoscaler, error) {
	cli := c.GetHPAClient(ns)
	return cli.Create(c.Ctx, autoscaler, opts)
}
