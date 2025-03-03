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

package http

import (
	"context"
	"net/url"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/west2-online/DomTok/app/assistant/service"

	"github.com/west2-online/DomTok/app/assistant/cli/server/adapter"
)

type Client struct {
	adapter.ServerCaller

	cli     *client.Client
	baseUrl string
}

type ClientConfig struct {
	BaseUrl string
}

func NewClient(opt *ClientConfig) *Client {
	cli, _ := client.NewClient(
		client.WithDialTimeout(time.Second),
		client.WithClientReadTimeout(time.Second),
		client.WithWriteTimeout(time.Second),
	)
	return &Client{cli: cli, baseUrl: opt.BaseUrl}
}

func (c *Client) buildUrl(path string) string {
	u, _ := url.JoinPath(c.baseUrl, path)
	return u
}

func (c *Client) do(ctx context.Context, req *protocol.Request, resp *protocol.Response) error {
	auth, ok := ctx.Value(service.CtxKeyAuthHeader).(string)
	if ok {
		req.SetHeader("Authorization", auth)
	}
	return c.cli.Do(ctx, req, resp)
}
