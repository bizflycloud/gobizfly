package testlib

func CloudServerURL(path string) string {
	return "/iaas-cloud/api" + path
}

func LoadBalancerURL(path string) string {
	return "/api/loadbalancers" + path
}

func AlertURL(path string) string {
	return "/api/alert" + path
}

func AutoScalingURL(path string) string {
	return "/api/auto-scaling" + path
}

func AuthURL(path string) string {
	return "/api" + path
}
