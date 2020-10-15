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

func (c *KubeClientSet) PatchStatefulSets(ns string, statefulsets string, patchType types.PatchType, data []byte, patchOptions v1.PatchOptions, subresource string) (*v12.StatefulSet, error) {
	client := c.GetStatefulSetClient(ns)
	if patchType == "" {
		patchType = types.StrategicMergePatchType
	}
	return client.Patch(context.TODO(), statefulsets, patchType, data, patchOptions, subresource)
}

func (c *KubeClientSet) RestartStatefulSets(ns string, ss string) error {
	patchSS := v12.StatefulSet{
		Spec: v12.StatefulSetSpec{
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
	data, err := json.Marshal(patchSS)
	if err != nil {
		return err
	}
	_, err = c.PatchStatefulSets(ns, ss, "", data, v1.PatchOptions{}, "")
	return err
}
