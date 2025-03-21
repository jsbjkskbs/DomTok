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

package cache

import (
	"context"

	"github.com/west2-online/DomTok/pkg/constants"
	"github.com/west2-online/DomTok/pkg/errno"
)

func (c *userCache) SetUserBaned(ctx context.Context, key string) error {
	err := c.client.Set(ctx, key, constants.UserBanned, constants.NeverExpire).Err()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, err.Error())
	}
	return nil
}

func (c *userCache) DeleteUserBaned(ctx context.Context, key string) error {
	err := c.client.Del(ctx, key).Err()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, err.Error())
	}
	return nil
}
