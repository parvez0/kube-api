package kube

import (
	"context"
	"fmt"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

// custom definition for client to add customize methods for fetching resource clients
type KubeClientSet struct {
	Client *kubernetes.Clientset
	Config *rest.Config
	Ctx    context.Context
}

// creating a kubernetes client object with the provided cluster name
// it will return a wrapper for client and config which can be used to get other resource clients like deployments, namespace etc
func CreateClient(ctx context.Context, cluster string) *KubeClientSet {
	var kubeconfig string
	var config *rest.Config
	var err error
	// setting config based on the environment if incluster is used
	// a cluster admin access or appropriate read write service account shoule be attached to the pod
	if os.Getenv("IN_CLUSTER") != "true" {
		// taking the home directory for config path
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		}
		// taking the config path from environment variable
		if path := os.Getenv("CONFIG_PATH"); path != "" {
			kubeconfig = path
		}
		// creating non interactive kube context from given cluster name
		config, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{
				ExplicitPath: kubeconfig,
			},
			&clientcmd.ConfigOverrides{
				CurrentContext: cluster,
			}).ClientConfig()
		if err != nil {
			panic(fmt.Sprintf("failed to create config from given config path - %+v", err))
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(fmt.Sprintf("failed to create inClusterConfig - %+v", err))
		}
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("failed to create client - %+v", err))
	}
	return &KubeClientSet{Client: client, Config: config, Ctx: ctx}
}

// GetNamespaceClient returns a namespace client interface generated from KubeClientSet config
func (c *KubeClientSet) GetNamespaceClient() v12.NamespaceInterface {
	return c.Client.CoreV1().Namespaces()
}

// GetDeploymentClient returns a deployment client interface generated from KubeClientSet config
func (c *KubeClientSet) GetDeploymentClient(namespace string) v1.DeploymentInterface {
	return c.Client.AppsV1().Deployments(namespace)
}

// GetStatefulSetClient returns a statefulset interface generated from KubeClientSet config
func (c *KubeClientSet) GetStatefulSetClient(namespace string) v1.StatefulSetInterface {
	return c.Client.AppsV1().StatefulSets(namespace)
}

// GetPVCClient returns a pvc interface generated from KubeClientSet config
func (c *KubeClientSet) GetPVCClient(namespace string) v12.PersistentVolumeClaimInterface {
	return c.Client.CoreV1().PersistentVolumeClaims(namespace)
}

// GetSVCClient returns a service interface generated from KubeClientSet config
func (c *KubeClientSet) GetSVCClient(namespace string) v12.ServiceInterface {
	return c.Client.CoreV1().Services(namespace)
}

// GetCMClient returns a config map client
func (c *KubeClientSet) GetCMClient(namespace string) v12.ConfigMapInterface {
	return c.Client.CoreV1().ConfigMaps(namespace)
}

// GetHPAClient returns a config map client
func (c *KubeClientSet) GetHPAClient(namespace string) v2beta2.HorizontalPodAutoscalerInterface {
	return c.Client.AutoscalingV2beta2().HorizontalPodAutoscalers(namespace)
}
