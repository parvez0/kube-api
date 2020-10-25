package kube

import (
	"context"
	"encoding/json"
	v12 "k8s.io/api/apps/v1"
	v13 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

// GetDeployment returns a deployment from the given namespace or returns an error
func (c *KubeClientSet) GetDeployment(ns string, deploy string, opts v1.GetOptions) (*v12.Deployment, error) {
	client := c.GetDeploymentClient(ns)
	return client.Get(c.Ctx, deploy, opts)
}

// GetDeployment returns a deployment from the given namespace or returns an error
func (c *KubeClientSet) ListDeployment(ns string, deploy string, opts v1.ListOptions) (*v12.DeploymentList, error) {
	client := c.GetDeploymentClient(ns)
	return client.List(c.Ctx, opts)
}

// CreateDeployment creates a deployment and returns response or error if not created
func (c *KubeClientSet) CreateDeployment(ns string, body *v12.Deployment, opts v1.CreateOptions) (*v12.Deployment, error) {
	client := c.GetDeploymentClient(ns)
	return client.Create(c.Ctx, body, opts)
}

func (c *KubeClientSet) PatchDeployment(ns string, deployment string, patchType types.PatchType, data []byte, patchOptions v1.PatchOptions, subresource string) (*v12.Deployment, error) {
	client := c.GetDeploymentClient(ns)
	if patchType == "" {
		patchType = types.StrategicMergePatchType
	}
	return client.Patch(c.Ctx, deployment, patchType, data, patchOptions, subresource)
}

func (c *KubeClientSet) RestartDeployment(ns string, deployment string) error {
	patchDeploy := v12.Deployment{
		Spec: v12.DeploymentSpec{
			Selector: &v1.LabelSelector{},
			Template: v13.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Annotations: map[string]string{
						"kubectl.kubernetes.io/restartedAt": time.Now().String(),
					},
				},
				Spec: v13.PodSpec{
					Containers: []v13.Container{},
				},
			},
		},
	}
	data, err := json.Marshal(patchDeploy)
	if err != nil {
		return err
	}
	_, err = c.PatchDeployment(ns, deployment, "", data, v1.PatchOptions{}, "")
	return err
}
