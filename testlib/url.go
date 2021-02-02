package testlib

func CloudServerURL(path string) string {
	return "/iaas-cloud/api" + path
}

func LoadBalancerURL(path string) string {
	return "/api/loadbalancers" + path
}

func CloudWatcherURL(path string) string {
	return "/api/alert" + path
}

func AutoScalingURL(path string) string {
	return "/api/auto-scaling" + path
}

func AuthURL(path string) string {
	return "/api" + path
}

func AccountURL(path string) string {
	return "/api/account" + path
}

func K8sURL(path string) string {
	return "/api/kubernetes-engine" + path
}

func RegistryURL(path string) string {
	return "/api/container-registry" + path
}

func CDNURL(path string) string {
	return "/api/cdn" + path
}

func DNSURL(path string) string {
	return "/api/dns" + path
}
