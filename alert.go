// This file is part of gobizfly

package gobizfly

var _ CloudWatcherService = (*cloudwatcherService)(nil)

type cloudwatcherService struct {
	client *Client
}

// CloudWatcherService is the interface wrap other resource's interfaces
type CloudWatcherService interface {
	Agents() *agents
	Alarms() *alarms
	Histories() *histories
	Receivers() *receivers
	Secrets() *secrets
}

// Agents is the interface wrap cloudwatcher agents interface
func (cws *cloudwatcherService) Agents() *agents {
	return &agents{client: cws.client}
}

// Alarms is the interface wrap cloudwatcher alarms interface
func (cws *cloudwatcherService) Alarms() *alarms {
	return &alarms{client: cws.client}
}

// Receivers is the interface wrap cloudwatcher receivers interface
func (cws *cloudwatcherService) Receivers() *receivers {
	return &receivers{client: cws.client}
}

// Histories is the interface wrap cloudwatcher histories interface
func (cws *cloudwatcherService) Histories() *histories {
	return &histories{client: cws.client}
}

// Secrets is the interface wrap cloudwatcher secrets interface
func (cws *cloudwatcherService) Secrets() *secrets {
	return &secrets{client: cws.client}
}

const (
	agentsResourcePath    = "/agents"
	alarmsResourcePath    = "/alarms"
	getVerificationPath   = "/resend"
	historiesResourcePath = "/histories"
	receiversResourcePath = "/receivers"
	secretsResourcePath   = "/secrets"
)

type agents struct {
	client *Client
}

type alarms struct {
	client *Client
}

type receivers struct {
	client *Client
}

type histories struct {
	client *Client
}

type secrets struct {
	client *Client
}
