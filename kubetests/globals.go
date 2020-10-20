package kubetests

import "github.com/parvez0/kube-api/kube"

const (
	Namespace = "test"
	ResourceName = "testkube"
	Image = "nginx"
	ServicePorts = 8080
)

var kube_cli = kube.CreateClient("staging")
