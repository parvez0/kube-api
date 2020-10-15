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

func (c *KubeClientSet) PatchDeployment(ns string, deployment string, patchType types.PatchType, data []byte, patchOptions v1.PatchOptions, subresource string) (*v12.Deployment, error) {
	client := c.GetDeploymentClient(ns)
	if patchType == "" {
		patchType = types.StrategicMergePatchType
	}
	return client.Patch(context.TODO(), deployment, patchType, data, patchOptions, subresource)
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
