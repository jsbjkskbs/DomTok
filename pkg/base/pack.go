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

package base

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/west2-online/DomTok/kitex_gen/model"
	"github.com/west2-online/DomTok/pkg/constants"
	"github.com/west2-online/DomTok/pkg/errno"
	"github.com/west2-online/DomTok/pkg/logger"
)

func BuildBaseResp(err error) *model.BaseResp {
	if err == nil {
		return &model.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.Success.ErrorMsg,
		}
	}
	Errno := errno.ConvertErr(err)
	return &model.BaseResp{
		Code: Errno.ErrorCode,
		Msg:  Errno.ErrorMsg,
	}
}

func BuildSuccessResp() *model.BaseResp {
	return BuildBaseResp(nil) // 直接调用原始函数，传入 nil 表示无错误
}

func LogError(err error) {
	if err == nil {
		return
	}

	e := errno.ConvertErr(err)
	if e.StackTrace() != nil {
		logger.Error(err.Error(), zap.String(constants.StackTraceKey, fmt.Sprintf("%+v", e.StackTrace())))
		return
	}
	logger.Error(err.Error())
}

func BuildTypeList[T any, U any](items []U, buildFunc func(U) T) []T {
	if len(items) == 0 {
		return nil
	}

	list := make([]T, len(items))
	for i, item := range items {
		list[i] = buildFunc(item)
	}
	return list
}
