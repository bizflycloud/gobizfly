// This file is part of gobizfly
//
// Copyright (C) 2020  BizFly Cloud
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	userPath = "/user"
)

var _ UserService = (*userService)(nil)

type userService struct {
	client *Client
}

type UserService interface {
	Get(ctx context.Context) (*User, error)
}

type User struct {
	Service           string             `json:"service"`
	URLType           string             `json:"url_type"`
	Origin            string             `json:"origin"`
	ClientType        string             `json:"client_type"`
	BillingBalance    int                `json:"billing_balance"`
	Balances          map[string]float32 `json:"balances"`
	PaymentMethod     string             `json:"payment_method"`
	BillingAccID      string             `json:"billing_acc_id"`
	Debit             bool               `json:"debit"`
	Email             string             `json:"email"`
	Phone             string             `json:"phone"`
	FullName          string             `json:"full_name"`
	VerifiedEmail     bool               `json:"verified_email"`
	VerifiedPhone     bool               `json:"verified_phone"`
	LoginAlert        bool               `json:"login_alert"`
	VerifiedPayment   bool               `json:"verified_payment"`
	LastRegion        string             `json:"last_region"`
	LastProject       string             `json:"last_project"`
	Type              string             `json:"type"`
	OTP               bool               `json:"otp"`
	Services          []Service          `json:"services"`
	Whitelist         []string           `json:"whitelist"`
	Expires           string             `json:"expires"`
	TenantID          string             `json:"tenant_id"`
	TenantName        string             `json:"tenant_name"`
	KsUserID          string             `json:"ks_user_id"`
	IAM               IAM                `json:"iam"`
	Domains           []string           `json:"domains"`
	PaymentType       string             `json:"payment_type"`
	DOB               string             `json:"dob"`
	Gender            string             `json:"_gender"`
	Trial             Trial              `json:"trial"`
	HasExpiredInvoice bool               `json:"has_expired_invoice"`
	NegativeBalance   bool               `json:"negative_balance"`
	Promotion         []string           `json:"promotion"`
}

type IAM struct {
	Expire          string `json:"expire"`
	TenantID        string `json:"tenant_id"`
	TenantName      string `json:"tenant_name"`
	VerifiedPhone   bool   `json:"verified_phone"`
	VerifiedEmail   bool   `json:"verified_email"`
	VerifiedPayment bool   `json:"verified_payment"`
}

type Trial struct {
	StartedAt    string `json:"started_at"`
	ExpiredAt    string `json:"expired_at"`
	Active       bool   `json:"active"`
	Enable       bool   `json:"enable"`
	ServiceLevel int    `json:"service_level"`
}

func (u userService) resourcePath() string {
	return userPath
}

func (u userService) Get(ctx context.Context) (*User, error) {
	req, err := u.client.NewRequest(ctx, http.MethodGet, serverServiceName, u.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := u.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var user *User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}
