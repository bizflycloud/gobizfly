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

// Comparison is represents comparison payload in alarms
type Comparison struct {
	CompareType string  `json:"compare_type"`
	Measurement string  `json:"measurement"`
	RangeTime   int     `json:"range_time"`
	Value       float64 `json:"value"`
}

// AlarmInstancesMonitors is represents instances payload - which servers will be monitored
type AlarmInstancesMonitors struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AlarmVolumesMonitor is represents volumes payload - which volumes will be monitored
type AlarmVolumesMonitor struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Device string `json:"device,omitempty"`
}

// HTTPHeaders is is represents http headers - which using call to http_url
type HTTPHeaders struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// AlarmLoadBalancersMonitor is represents load balancer payload - which load balancer will be monitored
type AlarmLoadBalancersMonitor struct {
	LoadBalancerID   string `json:"load_balancer_id"`
	LoadBalancerName string `json:"load_balancer_name"`
	TargetID         string `json:"target_id"`
	TargetName       string `json:"target_name,omitempty"`
	TargetType       string `json:"target_type"`
}

// AlarmReceiversUse is represents receiver's payload will be use create alarms
type AlarmReceiversUse struct {
	AutoscaleClusterName string `json:"autoscale_cluster_name,omitempty"`
	EmailAddress         string `json:"email_address,omitempty"`
	Name                 string `json:"name"`
	ReceiverID           string `json:"receiver_id"`
	SlackChannelName     string `json:"slack_channel_name,omitempty"`
	SMSInterval          int    `json:"sms_interval,omitempty"`
	SMSNumber            string `json:"sms_number,omitempty"`
	TelegramChatID       string `json:"telegram_chat_id,omitempty"`
	WebhookURL           string `json:"webhook_url,omitempty"`
}

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
