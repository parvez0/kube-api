package kube

import (
	"encoding/json"
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"time"
)

// GetStatefulSet returns a stateful from the given namespace or returns an error
func (c *KubeClientSet) GetStatefulSet(ns string, statefulset string, getOptions v1.GetOptions) (*v12.StatefulSet, error) {
	client := c.GetStatefulSetClient(ns)
	return client.Get(c.Ctx, statefulset, getOptions)
}

// ListStatefulSet returns a statefulset from the given namespace or returns an error
func (c *KubeClientSet) ListStatefulSet(ns string, opts v1.ListOptions) (*v12.StatefulSetList, error) {
	client := c.GetStatefulSetClient(ns)
	return client.List(c.Ctx, opts)
}

// CreateStatefulSet creates a namespace and returns response or error if not created
func (c *KubeClientSet) CreateStatefulSet(ns string, body *v12.StatefulSet, createOptions v1.CreateOptions) (*v12.StatefulSet, error) {
	client := c.GetStatefulSetClient(ns)
	return client.Create(c.Ctx, body, createOptions)
}

// PatchStatefulSets update the resource partially and returns and error if not successful
func (c *KubeClientSet) PatchStatefulSets(ns string, statefulsets string, patchType types.PatchType, data []byte, patchOptions v1.PatchOptions, subresource string) (*v12.StatefulSet, error) {
	client := c.GetStatefulSetClient(ns)
	if patchType == "" {
		patchType = types.StrategicMergePatchType
	}
	return client.Patch(c.Ctx, statefulsets, patchType, data, patchOptions, subresource)
}

// RestartStatefulSets restarts the resource by setting annotations to latest date
func (c *KubeClientSet) RestartStatefulSets(ns string, ss string) error {
	patchSS, err := c.GetStatefulSet(ns, ss, v1.GetOptions{})
	if err != nil {
		return err
	}
	patchSS.Spec.Template.Annotations = map[string]string{"kubectl.kubernetes.io/restartedAt": time.Now().String()}
	data, err := json.Marshal(patchSS)
	if err != nil {
		return err
	}
	_, err = c.PatchStatefulSets(ns, ss, "", data, v1.PatchOptions{}, "")
	return err
}

// DeleteStatefulSets deletes a resource or returns an error if not found
func (c *KubeClientSet) DeleteStatefulSets(ns string, statefulsets string, deleteOptions v1.DeleteOptions) error {
	client := c.GetStatefulSetClient(ns)
	return client.Delete(c.Ctx, statefulsets, deleteOptions)
}

// WatchStatefulSets provides the watch for statefulset event changes
func (c *KubeClientSet) WatchStatefulSets(ns string) (watch.Interface, error) {
	client := c.GetStatefulSetClient(ns)
	return client.Watch(c.Ctx, v1.ListOptions{})
}
