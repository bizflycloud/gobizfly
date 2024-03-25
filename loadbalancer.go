// This file is part of gobizfly

package gobizfly

const (
	loadBalancersPath        = "/loadbalancers"
	loadBalancerResourcePath = "/loadbalancer"
	listenerPath             = "/listener"
	poolPath                 = "/pool"
	healthMonitorPath        = "/healthmonitor"
	l7PolicyPath             = "/l7policy"
)

type resourceID struct {
	ID string
}
