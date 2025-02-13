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

// Code generated by hertz generator.

package cart

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	api "github.com/west2-online/DomTok/app/gateway/model/api/cart"
)

// AddGoodsIntoCart .
// @router /api/v1/api/add [POST]
func AddGoodsIntoCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.AddGoodsIntoCartRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.AddGoodsIntoCartResponse)

	c.JSON(consts.StatusOK, resp)
}

// ShowCartGoodsList .
// @router /api/v1/api/show [GET]
func ShowCartGoodsList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ShowCartGoodsListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ShowCartGoodsListResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateCartGoods .
// @router /api/v1/api/update [PUT]
func UpdateCartGoods(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateCartGoodsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UpdateCartGoodsResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteCartGoods .
// @router /api/v1/api/delete [DELETE]
func DeleteCartGoods(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DeleteCartGoodsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DeleteCartGoodsResponse)

	c.JSON(consts.StatusOK, resp)
}
