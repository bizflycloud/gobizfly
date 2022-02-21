package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Slack is represents slack payload - which will be use create a receiver
type Slack struct {
	SlackChannelName string `json:"channel_name"`
	WebhookURL       string `json:"webhook_url"`
}

// ReceiverCreateRequest contains receiver information.
type ReceiverCreateRequest struct {
	AutoScale      *AutoScalingWebhook `json:"autoscale,omitempty"`
	EmailAddress   string              `json:"email_address,omitempty"`
	Name           string              `json:"name"`
	Slack          *Slack              `json:"slack,omitempty"`
	SMSNumber      string              `json:"sms_number,omitempty"`
	TelegramChatID string              `json:"telegram_chat_id,omitempty"`
	WebhookURL     string              `json:"webhook_url,omitempty"`
}

// Receivers contains receiver information.
type Receivers struct {
	AutoScale              *AutoScalingWebhook `json:"autoscale,omitempty"`
	Created                string              `json:"_created"`
	Creator                string              `json:"creator"`
	EmailAddress           string              `json:"email_address,omitempty"`
	Name                   string              `json:"name"`
	ProjectID              string              `json:"project_id,omitempty"`
	ReceiverID             string              `json:"_id"`
	Slack                  *Slack              `json:"slack,omitempty"`
	SMSNumber              string              `json:"sms_number,omitempty"`
	TelegramChatID         string              `json:"telegram_chat_id,omitempty"`
	UserID                 string              `json:"user_id,omitempty"`
	VerifiedEmailDddress   bool                `json:"verified_email_address,omitempty"`
	VerifiedSMSNumber      bool                `json:"verified_sms_number,omitempty"`
	VerifiedTelegramChatID bool                `json:"verified_telegram_chat_id,omitempty"`
	VerifiedWebhookURL     bool                `json:"verified_webhook_url,omitempty"`
	WebhookURL             string              `json:"webhook_url,omitempty"`
}

func (r *receivers) resourcePath() string {
	return strings.Join([]string{receiversResourcePath}, "/")
}

func (r *receivers) itemPath(id string) string {
	return strings.Join([]string{receiversResourcePath, id}, "/")
}

func (r *receivers) verificationPath() string {
	return strings.Join([]string{getVerificationPath}, "/")
}

// List receivers
func (r *receivers) List(ctx context.Context, filters *string) ([]*Receivers, error) {
	req, err := r.client.NewRequest(ctx, http.MethodGet, cloudwatcherServiceName, r.resourcePath(), nil)
	if err != nil {
		return nil, err
	}

	if filters != nil {
		q := req.URL.Query()
		q.Add("where", *filters)
		req.URL.RawQuery = q.Encode()
	}

	resp, err := r.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Receivers []*Receivers `json:"_items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Receivers, nil
}

// Create a receiver
func (r *receivers) Create(ctx context.Context, rcr *ReceiverCreateRequest) (*ResponseRequest, error) {
	req, err := r.client.NewRequest(ctx, http.MethodPost, cloudwatcherServiceName, r.resourcePath(), &rcr)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData = &ResponseRequest{}
	if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
		return nil, err
	}
	return respData, nil
}

// Get a receiver
func (r *receivers) Get(ctx context.Context, id string) (*Receivers, error) {
	req, err := r.client.NewRequest(ctx, http.MethodGet, cloudwatcherServiceName, r.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	receiver := &Receivers{}
	if err := json.NewDecoder(resp.Body).Decode(receiver); err != nil {
		return nil, err
	}
	return receiver, nil
}

// Update receiver
func (r *receivers) Update(ctx context.Context, id string, rur *ReceiverCreateRequest) (*ResponseRequest, error) {
	req, err := r.client.NewRequest(ctx, http.MethodPut, cloudwatcherServiceName, r.itemPath(id), &rur)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData := &ResponseRequest{}
	if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
		return nil, err
	}

	return respData, nil
}

// Delete receiver
func (r *receivers) Delete(ctx context.Context, id string) error {
	req, err := r.client.NewRequest(ctx, http.MethodDelete, cloudwatcherServiceName, r.itemPath(id), nil)
	if err != nil {
		return err
	}
	resp, err := r.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}

// ResendVerificationLink is use get a link verification
func (r *receivers) ResendVerificationLink(ctx context.Context, id string, rType string) error {
	req, err := r.client.NewRequest(ctx, http.MethodGet, cloudwatcherServiceName, r.verificationPath(), nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("rec_id", id)
	q.Add("rec_type", rType)
	req.URL.RawQuery = q.Encode()

	resp, err := r.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}
