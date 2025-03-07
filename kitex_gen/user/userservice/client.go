/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kitex v0.12.1. DO NOT EDIT.

package userservice

import (
	"context"

	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"

	user "github.com/west2-online/DomTok/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error)
	Login(ctx context.Context, req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error)
	GetAddress(ctx context.Context, req *user.GetAddressRequest, callOptions ...callopt.Option) (r *user.GetAddressResponse, err error)
	AddAddress(ctx context.Context, req *user.AddAddressRequest, callOptions ...callopt.Option) (r *user.AddAddressResponse, err error)
	BanUser(ctx context.Context, req *user.BanUserReq, callOptions ...callopt.Option) (r *user.BanUserResp, err error)
	LiftBandUser(ctx context.Context, req *user.LiftBanUserReq, callOptions ...callopt.Option) (r *user.LiftBanUserResp, err error)
	Logout(ctx context.Context, req *user.LogoutReq, callOptions ...callopt.Option) (r *user.LogoutResp, err error)
	SetAdministrator(ctx context.Context, req *user.SetAdministratorReq, callOptions ...callopt.Option) (r *user.SetAdministratorResp, err error)
	GetUserInfo(ctx context.Context, req *user.GetUserInfoReq, callOptions ...callopt.Option) (r *user.GetUserInfoResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Register(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, req)
}

func (p *kUserServiceClient) Login(ctx context.Context, req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kUserServiceClient) GetAddress(ctx context.Context, req *user.GetAddressRequest, callOptions ...callopt.Option) (r *user.GetAddressResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAddress(ctx, req)
}

func (p *kUserServiceClient) AddAddress(ctx context.Context, req *user.AddAddressRequest, callOptions ...callopt.Option) (r *user.AddAddressResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddAddress(ctx, req)
}

func (p *kUserServiceClient) BanUser(ctx context.Context, req *user.BanUserReq, callOptions ...callopt.Option) (r *user.BanUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BanUser(ctx, req)
}

func (p *kUserServiceClient) LiftBandUser(ctx context.Context, req *user.LiftBanUserReq, callOptions ...callopt.Option) (r *user.LiftBanUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LiftBandUser(ctx, req)
}

func (p *kUserServiceClient) Logout(ctx context.Context, req *user.LogoutReq, callOptions ...callopt.Option) (r *user.LogoutResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Logout(ctx, req)
}

func (p *kUserServiceClient) SetAdministrator(ctx context.Context, req *user.SetAdministratorReq, callOptions ...callopt.Option) (r *user.SetAdministratorResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetAdministrator(ctx, req)
}

func (p *kUserServiceClient) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq, callOptions ...callopt.Option) (r *user.GetUserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, req)
}
