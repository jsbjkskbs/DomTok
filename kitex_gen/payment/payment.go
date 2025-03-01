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

// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package payment

import (
	"context"
	"fmt"
	"strings"

	"github.com/west2-online/DomTok/kitex_gen/model"
)

type PaymentTokenRequest struct {
	OrderID int64 `thrift:"orderID,1,required" frugal:"1,required,i64" json:"orderID"`
}

func NewPaymentTokenRequest() *PaymentTokenRequest {
	return &PaymentTokenRequest{}
}

func (p *PaymentTokenRequest) InitDefault() {
}

func (p *PaymentTokenRequest) GetOrderID() (v int64) {
	return p.OrderID
}
func (p *PaymentTokenRequest) SetOrderID(val int64) {
	p.OrderID = val
}

func (p *PaymentTokenRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentTokenRequest(%+v)", *p)
}

func (p *PaymentTokenRequest) DeepEqual(ano *PaymentTokenRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.OrderID) {
		return false
	}
	return true
}

func (p *PaymentTokenRequest) Field1DeepEqual(src int64) bool {

	if p.OrderID != src {
		return false
	}
	return true
}

var fieldIDToName_PaymentTokenRequest = map[int16]string{
	1: "orderID",
}

type PaymentTokenResponse struct {
	Base      *model.BaseResp         `thrift:"base,1" frugal:"1,default,model.BaseResp" json:"base"`
	TokenInfo *model.PaymentTokenInfo `thrift:"tokenInfo,2" frugal:"2,default,model.PaymentTokenInfo" json:"tokenInfo"`
}

func NewPaymentTokenResponse() *PaymentTokenResponse {
	return &PaymentTokenResponse{}
}

func (p *PaymentTokenResponse) InitDefault() {
}

var PaymentTokenResponse_Base_DEFAULT *model.BaseResp

func (p *PaymentTokenResponse) GetBase() (v *model.BaseResp) {
	if !p.IsSetBase() {
		return PaymentTokenResponse_Base_DEFAULT
	}
	return p.Base
}

var PaymentTokenResponse_TokenInfo_DEFAULT *model.PaymentTokenInfo

func (p *PaymentTokenResponse) GetTokenInfo() (v *model.PaymentTokenInfo) {
	if !p.IsSetTokenInfo() {
		return PaymentTokenResponse_TokenInfo_DEFAULT
	}
	return p.TokenInfo
}
func (p *PaymentTokenResponse) SetBase(val *model.BaseResp) {
	p.Base = val
}
func (p *PaymentTokenResponse) SetTokenInfo(val *model.PaymentTokenInfo) {
	p.TokenInfo = val
}

func (p *PaymentTokenResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *PaymentTokenResponse) IsSetTokenInfo() bool {
	return p.TokenInfo != nil
}

func (p *PaymentTokenResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentTokenResponse(%+v)", *p)
}

func (p *PaymentTokenResponse) DeepEqual(ano *PaymentTokenResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Base) {
		return false
	}
	if !p.Field2DeepEqual(ano.TokenInfo) {
		return false
	}
	return true
}

func (p *PaymentTokenResponse) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}
func (p *PaymentTokenResponse) Field2DeepEqual(src *model.PaymentTokenInfo) bool {

	if !p.TokenInfo.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentTokenResponse = map[int16]string{
	1: "base",
	2: "tokenInfo",
}

type PaymentRequest struct {
	OrderID     int64                 `thrift:"orderID,1,required" frugal:"1,required,i64" json:"orderID"`
	UserID      int64                 `thrift:"userID,2,required" frugal:"2,required,i64" json:"userID"`
	CreditCard  *model.CreditCardInfo `thrift:"creditCard,4,required" frugal:"4,required,model.CreditCardInfo" json:"creditCard"`
	Description *string               `thrift:"description,5,optional" frugal:"5,optional,string" json:"description,omitempty"`
}

func NewPaymentRequest() *PaymentRequest {
	return &PaymentRequest{}
}

func (p *PaymentRequest) InitDefault() {
}

func (p *PaymentRequest) GetOrderID() (v int64) {
	return p.OrderID
}

func (p *PaymentRequest) GetUserID() (v int64) {
	return p.UserID
}

var PaymentRequest_CreditCard_DEFAULT *model.CreditCardInfo

func (p *PaymentRequest) GetCreditCard() (v *model.CreditCardInfo) {
	if !p.IsSetCreditCard() {
		return PaymentRequest_CreditCard_DEFAULT
	}
	return p.CreditCard
}

var PaymentRequest_Description_DEFAULT string

func (p *PaymentRequest) GetDescription() (v string) {
	if !p.IsSetDescription() {
		return PaymentRequest_Description_DEFAULT
	}
	return *p.Description
}
func (p *PaymentRequest) SetOrderID(val int64) {
	p.OrderID = val
}
func (p *PaymentRequest) SetUserID(val int64) {
	p.UserID = val
}
func (p *PaymentRequest) SetCreditCard(val *model.CreditCardInfo) {
	p.CreditCard = val
}
func (p *PaymentRequest) SetDescription(val *string) {
	p.Description = val
}

func (p *PaymentRequest) IsSetCreditCard() bool {
	return p.CreditCard != nil
}

func (p *PaymentRequest) IsSetDescription() bool {
	return p.Description != nil
}

func (p *PaymentRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentRequest(%+v)", *p)
}

func (p *PaymentRequest) DeepEqual(ano *PaymentRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.OrderID) {
		return false
	}
	if !p.Field2DeepEqual(ano.UserID) {
		return false
	}
	if !p.Field4DeepEqual(ano.CreditCard) {
		return false
	}
	if !p.Field5DeepEqual(ano.Description) {
		return false
	}
	return true
}

func (p *PaymentRequest) Field1DeepEqual(src int64) bool {

	if p.OrderID != src {
		return false
	}
	return true
}
func (p *PaymentRequest) Field2DeepEqual(src int64) bool {

	if p.UserID != src {
		return false
	}
	return true
}
func (p *PaymentRequest) Field4DeepEqual(src *model.CreditCardInfo) bool {

	if !p.CreditCard.DeepEqual(src) {
		return false
	}
	return true
}
func (p *PaymentRequest) Field5DeepEqual(src *string) bool {

	if p.Description == src {
		return true
	} else if p.Description == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Description, *src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_PaymentRequest = map[int16]string{
	1: "orderID",
	2: "userID",
	4: "creditCard",
	5: "description",
}

type PaymentResponse struct {
	Base      *model.BaseResp `thrift:"base,1" frugal:"1,default,model.BaseResp" json:"base"`
	PaymentID int64           `thrift:"paymentID,2,required" frugal:"2,required,i64" json:"paymentID"`
	Status    int64           `thrift:"status,3,required" frugal:"3,required,i64" json:"status"`
}

func NewPaymentResponse() *PaymentResponse {
	return &PaymentResponse{}
}

func (p *PaymentResponse) InitDefault() {
}

var PaymentResponse_Base_DEFAULT *model.BaseResp

func (p *PaymentResponse) GetBase() (v *model.BaseResp) {
	if !p.IsSetBase() {
		return PaymentResponse_Base_DEFAULT
	}
	return p.Base
}

func (p *PaymentResponse) GetPaymentID() (v int64) {
	return p.PaymentID
}

func (p *PaymentResponse) GetStatus() (v int64) {
	return p.Status
}
func (p *PaymentResponse) SetBase(val *model.BaseResp) {
	p.Base = val
}
func (p *PaymentResponse) SetPaymentID(val int64) {
	p.PaymentID = val
}
func (p *PaymentResponse) SetStatus(val int64) {
	p.Status = val
}

func (p *PaymentResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *PaymentResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentResponse(%+v)", *p)
}

func (p *PaymentResponse) DeepEqual(ano *PaymentResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Base) {
		return false
	}
	if !p.Field2DeepEqual(ano.PaymentID) {
		return false
	}
	if !p.Field3DeepEqual(ano.Status) {
		return false
	}
	return true
}

func (p *PaymentResponse) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}
func (p *PaymentResponse) Field2DeepEqual(src int64) bool {

	if p.PaymentID != src {
		return false
	}
	return true
}
func (p *PaymentResponse) Field3DeepEqual(src int64) bool {

	if p.Status != src {
		return false
	}
	return true
}

var fieldIDToName_PaymentResponse = map[int16]string{
	1: "base",
	2: "paymentID",
	3: "status",
}

type RefundTokenRequest struct {
	OrderID int64 `thrift:"orderID,1,required" frugal:"1,required,i64" json:"orderID"`
}

func NewRefundTokenRequest() *RefundTokenRequest {
	return &RefundTokenRequest{}
}

func (p *RefundTokenRequest) InitDefault() {
}

func (p *RefundTokenRequest) GetOrderID() (v int64) {
	return p.OrderID
}
func (p *RefundTokenRequest) SetOrderID(val int64) {
	p.OrderID = val
}

func (p *RefundTokenRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RefundTokenRequest(%+v)", *p)
}

func (p *RefundTokenRequest) DeepEqual(ano *RefundTokenRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.OrderID) {
		return false
	}
	return true
}

func (p *RefundTokenRequest) Field1DeepEqual(src int64) bool {

	if p.OrderID != src {
		return false
	}
	return true
}

var fieldIDToName_RefundTokenRequest = map[int16]string{
	1: "orderID",
}

type RefundTokenResponse struct {
	Base     *model.BaseResp `thrift:"base,1" frugal:"1,default,model.BaseResp" json:"base"`
	RefundID int64           `thrift:"refundID,2,required" frugal:"2,required,i64" json:"refundID"`
}

func NewRefundTokenResponse() *RefundTokenResponse {
	return &RefundTokenResponse{}
}

func (p *RefundTokenResponse) InitDefault() {
}

var RefundTokenResponse_Base_DEFAULT *model.BaseResp

func (p *RefundTokenResponse) GetBase() (v *model.BaseResp) {
	if !p.IsSetBase() {
		return RefundTokenResponse_Base_DEFAULT
	}
	return p.Base
}

func (p *RefundTokenResponse) GetRefundID() (v int64) {
	return p.RefundID
}
func (p *RefundTokenResponse) SetBase(val *model.BaseResp) {
	p.Base = val
}
func (p *RefundTokenResponse) SetRefundID(val int64) {
	p.RefundID = val
}

func (p *RefundTokenResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *RefundTokenResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RefundTokenResponse(%+v)", *p)
}

func (p *RefundTokenResponse) DeepEqual(ano *RefundTokenResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Base) {
		return false
	}
	if !p.Field2DeepEqual(ano.RefundID) {
		return false
	}
	return true
}

func (p *RefundTokenResponse) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}
func (p *RefundTokenResponse) Field2DeepEqual(src int64) bool {

	if p.RefundID != src {
		return false
	}
	return true
}

var fieldIDToName_RefundTokenResponse = map[int16]string{
	1: "base",
	2: "refundID",
}

type RefundRequest struct {
	OrderID      int64   `thrift:"orderID,1,required" frugal:"1,required,i64" json:"orderID"`
	RefundAmount float64 `thrift:"refundAmount,2,required" frugal:"2,required,double" json:"refundAmount"`
	RefundReason string  `thrift:"refundReason,3,required" frugal:"3,required,string" json:"refundReason"`
}

func NewRefundRequest() *RefundRequest {
	return &RefundRequest{}
}

func (p *RefundRequest) InitDefault() {
}

func (p *RefundRequest) GetOrderID() (v int64) {
	return p.OrderID
}

func (p *RefundRequest) GetRefundAmount() (v float64) {
	return p.RefundAmount
}

func (p *RefundRequest) GetRefundReason() (v string) {
	return p.RefundReason
}
func (p *RefundRequest) SetOrderID(val int64) {
	p.OrderID = val
}
func (p *RefundRequest) SetRefundAmount(val float64) {
	p.RefundAmount = val
}
func (p *RefundRequest) SetRefundReason(val string) {
	p.RefundReason = val
}

func (p *RefundRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RefundRequest(%+v)", *p)
}

func (p *RefundRequest) DeepEqual(ano *RefundRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.OrderID) {
		return false
	}
	if !p.Field2DeepEqual(ano.RefundAmount) {
		return false
	}
	if !p.Field3DeepEqual(ano.RefundReason) {
		return false
	}
	return true
}

func (p *RefundRequest) Field1DeepEqual(src int64) bool {

	if p.OrderID != src {
		return false
	}
	return true
}
func (p *RefundRequest) Field2DeepEqual(src float64) bool {

	if p.RefundAmount != src {
		return false
	}
	return true
}
func (p *RefundRequest) Field3DeepEqual(src string) bool {

	if strings.Compare(p.RefundReason, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_RefundRequest = map[int16]string{
	1: "orderID",
	2: "refundAmount",
	3: "refundReason",
}

type RefundResponse struct {
	Base       *model.BaseResp           `thrift:"base,1" frugal:"1,default,model.BaseResp" json:"base"`
	RefundInfo *model.RefundResponseInfo `thrift:"refundInfo,2" frugal:"2,default,model.RefundResponseInfo" json:"refundInfo"`
}

func NewRefundResponse() *RefundResponse {
	return &RefundResponse{}
}

func (p *RefundResponse) InitDefault() {
}

var RefundResponse_Base_DEFAULT *model.BaseResp

func (p *RefundResponse) GetBase() (v *model.BaseResp) {
	if !p.IsSetBase() {
		return RefundResponse_Base_DEFAULT
	}
	return p.Base
}

var RefundResponse_RefundInfo_DEFAULT *model.RefundResponseInfo

func (p *RefundResponse) GetRefundInfo() (v *model.RefundResponseInfo) {
	if !p.IsSetRefundInfo() {
		return RefundResponse_RefundInfo_DEFAULT
	}
	return p.RefundInfo
}
func (p *RefundResponse) SetBase(val *model.BaseResp) {
	p.Base = val
}
func (p *RefundResponse) SetRefundInfo(val *model.RefundResponseInfo) {
	p.RefundInfo = val
}

func (p *RefundResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *RefundResponse) IsSetRefundInfo() bool {
	return p.RefundInfo != nil
}

func (p *RefundResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RefundResponse(%+v)", *p)
}

func (p *RefundResponse) DeepEqual(ano *RefundResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Base) {
		return false
	}
	if !p.Field2DeepEqual(ano.RefundInfo) {
		return false
	}
	return true
}

func (p *RefundResponse) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}
func (p *RefundResponse) Field2DeepEqual(src *model.RefundResponseInfo) bool {

	if !p.RefundInfo.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_RefundResponse = map[int16]string{
	1: "base",
	2: "refundInfo",
}

type RefundReviewRequest struct {
	OrderID int64 `thrift:"orderID,1,required" frugal:"1,required,i64" json:"orderID"`
	Passed  bool  `thrift:"passed,2,required" frugal:"2,required,bool" json:"passed"`
}

func NewRefundReviewRequest() *RefundReviewRequest {
	return &RefundReviewRequest{}
}

func (p *RefundReviewRequest) InitDefault() {
}

func (p *RefundReviewRequest) GetOrderID() (v int64) {
	return p.OrderID
}

func (p *RefundReviewRequest) GetPassed() (v bool) {
	return p.Passed
}
func (p *RefundReviewRequest) SetOrderID(val int64) {
	p.OrderID = val
}
func (p *RefundReviewRequest) SetPassed(val bool) {
	p.Passed = val
}

func (p *RefundReviewRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RefundReviewRequest(%+v)", *p)
}

func (p *RefundReviewRequest) DeepEqual(ano *RefundReviewRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.OrderID) {
		return false
	}
	if !p.Field2DeepEqual(ano.Passed) {
		return false
	}
	return true
}

func (p *RefundReviewRequest) Field1DeepEqual(src int64) bool {

	if p.OrderID != src {
		return false
	}
	return true
}
func (p *RefundReviewRequest) Field2DeepEqual(src bool) bool {

	if p.Passed != src {
		return false
	}
	return true
}

var fieldIDToName_RefundReviewRequest = map[int16]string{
	1: "orderID",
	2: "passed",
}

type RefundReviewResponse struct {
	Base *model.BaseResp `thrift:"base,1" frugal:"1,default,model.BaseResp" json:"base"`
}

func NewRefundReviewResponse() *RefundReviewResponse {
	return &RefundReviewResponse{}
}

func (p *RefundReviewResponse) InitDefault() {
}

var RefundReviewResponse_Base_DEFAULT *model.BaseResp

func (p *RefundReviewResponse) GetBase() (v *model.BaseResp) {
	if !p.IsSetBase() {
		return RefundReviewResponse_Base_DEFAULT
	}
	return p.Base
}
func (p *RefundReviewResponse) SetBase(val *model.BaseResp) {
	p.Base = val
}

func (p *RefundReviewResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *RefundReviewResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RefundReviewResponse(%+v)", *p)
}

func (p *RefundReviewResponse) DeepEqual(ano *RefundReviewResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *RefundReviewResponse) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_RefundReviewResponse = map[int16]string{
	1: "base",
}

type PaymentCheckoutRequest struct {
	OrderID int64  `thrift:"orderID,1,required" frugal:"1,required,i64" json:"orderID"`
	Token   string `thrift:"token,2,required" frugal:"2,required,string" json:"token"`
}

func NewPaymentCheckoutRequest() *PaymentCheckoutRequest {
	return &PaymentCheckoutRequest{}
}

func (p *PaymentCheckoutRequest) InitDefault() {
}

func (p *PaymentCheckoutRequest) GetOrderID() (v int64) {
	return p.OrderID
}

func (p *PaymentCheckoutRequest) GetToken() (v string) {
	return p.Token
}
func (p *PaymentCheckoutRequest) SetOrderID(val int64) {
	p.OrderID = val
}
func (p *PaymentCheckoutRequest) SetToken(val string) {
	p.Token = val
}

func (p *PaymentCheckoutRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentCheckoutRequest(%+v)", *p)
}

func (p *PaymentCheckoutRequest) DeepEqual(ano *PaymentCheckoutRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.OrderID) {
		return false
	}
	if !p.Field2DeepEqual(ano.Token) {
		return false
	}
	return true
}

func (p *PaymentCheckoutRequest) Field1DeepEqual(src int64) bool {

	if p.OrderID != src {
		return false
	}
	return true
}
func (p *PaymentCheckoutRequest) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Token, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_PaymentCheckoutRequest = map[int16]string{
	1: "orderID",
	2: "token",
}

type PaymentCheckoutResponse struct {
	Base *model.BaseResp `thrift:"base,1" frugal:"1,default,model.BaseResp" json:"base"`
}

func NewPaymentCheckoutResponse() *PaymentCheckoutResponse {
	return &PaymentCheckoutResponse{}
}

func (p *PaymentCheckoutResponse) InitDefault() {
}

var PaymentCheckoutResponse_Base_DEFAULT *model.BaseResp

func (p *PaymentCheckoutResponse) GetBase() (v *model.BaseResp) {
	if !p.IsSetBase() {
		return PaymentCheckoutResponse_Base_DEFAULT
	}
	return p.Base
}
func (p *PaymentCheckoutResponse) SetBase(val *model.BaseResp) {
	p.Base = val
}

func (p *PaymentCheckoutResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *PaymentCheckoutResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentCheckoutResponse(%+v)", *p)
}

func (p *PaymentCheckoutResponse) DeepEqual(ano *PaymentCheckoutResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *PaymentCheckoutResponse) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentCheckoutResponse = map[int16]string{
	1: "base",
}

type PaymentService interface {
	ProcessPayment(ctx context.Context, request *PaymentRequest) (r *PaymentResponse, err error)

	RequestPaymentToken(ctx context.Context, request *PaymentTokenRequest) (r *PaymentTokenResponse, err error)

	RequestPaymentCheckout(ctx context.Context, request *PaymentCheckoutRequest) (r *PaymentCheckoutResponse, err error)

	RefundReview(ctx context.Context, request *RefundReviewRequest) (r *RefundReviewResponse, err error)

	RequestRefund(ctx context.Context, request *RefundRequest) (r *RefundResponse, err error)
}

type PaymentServiceProcessPaymentArgs struct {
	Request *PaymentRequest `thrift:"request,1" frugal:"1,default,PaymentRequest" json:"request"`
}

func NewPaymentServiceProcessPaymentArgs() *PaymentServiceProcessPaymentArgs {
	return &PaymentServiceProcessPaymentArgs{}
}

func (p *PaymentServiceProcessPaymentArgs) InitDefault() {
}

var PaymentServiceProcessPaymentArgs_Request_DEFAULT *PaymentRequest

func (p *PaymentServiceProcessPaymentArgs) GetRequest() (v *PaymentRequest) {
	if !p.IsSetRequest() {
		return PaymentServiceProcessPaymentArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PaymentServiceProcessPaymentArgs) SetRequest(val *PaymentRequest) {
	p.Request = val
}

func (p *PaymentServiceProcessPaymentArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PaymentServiceProcessPaymentArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceProcessPaymentArgs(%+v)", *p)
}

func (p *PaymentServiceProcessPaymentArgs) DeepEqual(ano *PaymentServiceProcessPaymentArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *PaymentServiceProcessPaymentArgs) Field1DeepEqual(src *PaymentRequest) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentServiceProcessPaymentArgs = map[int16]string{
	1: "request",
}

type PaymentServiceProcessPaymentResult struct {
	Success *PaymentResponse `thrift:"success,0,optional" frugal:"0,optional,PaymentResponse" json:"success,omitempty"`
}

func NewPaymentServiceProcessPaymentResult() *PaymentServiceProcessPaymentResult {
	return &PaymentServiceProcessPaymentResult{}
}

func (p *PaymentServiceProcessPaymentResult) InitDefault() {
}

var PaymentServiceProcessPaymentResult_Success_DEFAULT *PaymentResponse

func (p *PaymentServiceProcessPaymentResult) GetSuccess() (v *PaymentResponse) {
	if !p.IsSetSuccess() {
		return PaymentServiceProcessPaymentResult_Success_DEFAULT
	}
	return p.Success
}
func (p *PaymentServiceProcessPaymentResult) SetSuccess(x interface{}) {
	p.Success = x.(*PaymentResponse)
}

func (p *PaymentServiceProcessPaymentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentServiceProcessPaymentResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceProcessPaymentResult(%+v)", *p)
}

func (p *PaymentServiceProcessPaymentResult) DeepEqual(ano *PaymentServiceProcessPaymentResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *PaymentServiceProcessPaymentResult) Field0DeepEqual(src *PaymentResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentServiceProcessPaymentResult = map[int16]string{
	0: "success",
}

type PaymentServiceRequestPaymentTokenArgs struct {
	Request *PaymentTokenRequest `thrift:"request,1" frugal:"1,default,PaymentTokenRequest" json:"request"`
}

func NewPaymentServiceRequestPaymentTokenArgs() *PaymentServiceRequestPaymentTokenArgs {
	return &PaymentServiceRequestPaymentTokenArgs{}
}

func (p *PaymentServiceRequestPaymentTokenArgs) InitDefault() {
}

var PaymentServiceRequestPaymentTokenArgs_Request_DEFAULT *PaymentTokenRequest

func (p *PaymentServiceRequestPaymentTokenArgs) GetRequest() (v *PaymentTokenRequest) {
	if !p.IsSetRequest() {
		return PaymentServiceRequestPaymentTokenArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PaymentServiceRequestPaymentTokenArgs) SetRequest(val *PaymentTokenRequest) {
	p.Request = val
}

func (p *PaymentServiceRequestPaymentTokenArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PaymentServiceRequestPaymentTokenArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceRequestPaymentTokenArgs(%+v)", *p)
}

func (p *PaymentServiceRequestPaymentTokenArgs) DeepEqual(ano *PaymentServiceRequestPaymentTokenArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *PaymentServiceRequestPaymentTokenArgs) Field1DeepEqual(src *PaymentTokenRequest) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentServiceRequestPaymentTokenArgs = map[int16]string{
	1: "request",
}

type PaymentServiceRequestPaymentTokenResult struct {
	Success *PaymentTokenResponse `thrift:"success,0,optional" frugal:"0,optional,PaymentTokenResponse" json:"success,omitempty"`
}

func NewPaymentServiceRequestPaymentTokenResult() *PaymentServiceRequestPaymentTokenResult {
	return &PaymentServiceRequestPaymentTokenResult{}
}

func (p *PaymentServiceRequestPaymentTokenResult) InitDefault() {
}

var PaymentServiceRequestPaymentTokenResult_Success_DEFAULT *PaymentTokenResponse

func (p *PaymentServiceRequestPaymentTokenResult) GetSuccess() (v *PaymentTokenResponse) {
	if !p.IsSetSuccess() {
		return PaymentServiceRequestPaymentTokenResult_Success_DEFAULT
	}
	return p.Success
}
func (p *PaymentServiceRequestPaymentTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*PaymentTokenResponse)
}

func (p *PaymentServiceRequestPaymentTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentServiceRequestPaymentTokenResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceRequestPaymentTokenResult(%+v)", *p)
}

func (p *PaymentServiceRequestPaymentTokenResult) DeepEqual(ano *PaymentServiceRequestPaymentTokenResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *PaymentServiceRequestPaymentTokenResult) Field0DeepEqual(src *PaymentTokenResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentServiceRequestPaymentTokenResult = map[int16]string{
	0: "success",
}

type PaymentServiceRequestPaymentCheckoutArgs struct {
	Request *PaymentCheckoutRequest `thrift:"request,1" frugal:"1,default,PaymentCheckoutRequest" json:"request"`
}

func NewPaymentServiceRequestPaymentCheckoutArgs() *PaymentServiceRequestPaymentCheckoutArgs {
	return &PaymentServiceRequestPaymentCheckoutArgs{}
}

func (p *PaymentServiceRequestPaymentCheckoutArgs) InitDefault() {
}

var PaymentServiceRequestPaymentCheckoutArgs_Request_DEFAULT *PaymentCheckoutRequest

func (p *PaymentServiceRequestPaymentCheckoutArgs) GetRequest() (v *PaymentCheckoutRequest) {
	if !p.IsSetRequest() {
		return PaymentServiceRequestPaymentCheckoutArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PaymentServiceRequestPaymentCheckoutArgs) SetRequest(val *PaymentCheckoutRequest) {
	p.Request = val
}

func (p *PaymentServiceRequestPaymentCheckoutArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PaymentServiceRequestPaymentCheckoutArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceRequestPaymentCheckoutArgs(%+v)", *p)
}

func (p *PaymentServiceRequestPaymentCheckoutArgs) DeepEqual(ano *PaymentServiceRequestPaymentCheckoutArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *PaymentServiceRequestPaymentCheckoutArgs) Field1DeepEqual(src *PaymentCheckoutRequest) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentServiceRequestPaymentCheckoutArgs = map[int16]string{
	1: "request",
}

type PaymentServiceRequestPaymentCheckoutResult struct {
	Success *PaymentCheckoutResponse `thrift:"success,0,optional" frugal:"0,optional,PaymentCheckoutResponse" json:"success,omitempty"`
}

func NewPaymentServiceRequestPaymentCheckoutResult() *PaymentServiceRequestPaymentCheckoutResult {
	return &PaymentServiceRequestPaymentCheckoutResult{}
}

func (p *PaymentServiceRequestPaymentCheckoutResult) InitDefault() {
}

var PaymentServiceRequestPaymentCheckoutResult_Success_DEFAULT *PaymentCheckoutResponse

func (p *PaymentServiceRequestPaymentCheckoutResult) GetSuccess() (v *PaymentCheckoutResponse) {
	if !p.IsSetSuccess() {
		return PaymentServiceRequestPaymentCheckoutResult_Success_DEFAULT
	}
	return p.Success
}
func (p *PaymentServiceRequestPaymentCheckoutResult) SetSuccess(x interface{}) {
	p.Success = x.(*PaymentCheckoutResponse)
}

func (p *PaymentServiceRequestPaymentCheckoutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentServiceRequestPaymentCheckoutResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceRequestPaymentCheckoutResult(%+v)", *p)
}

func (p *PaymentServiceRequestPaymentCheckoutResult) DeepEqual(ano *PaymentServiceRequestPaymentCheckoutResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *PaymentServiceRequestPaymentCheckoutResult) Field0DeepEqual(src *PaymentCheckoutResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentServiceRequestPaymentCheckoutResult = map[int16]string{
	0: "success",
}

type PaymentServiceRefundReviewArgs struct {
	Request *RefundReviewRequest `thrift:"request,1" frugal:"1,default,RefundReviewRequest" json:"request"`
}

func NewPaymentServiceRefundReviewArgs() *PaymentServiceRefundReviewArgs {
	return &PaymentServiceRefundReviewArgs{}
}

func (p *PaymentServiceRefundReviewArgs) InitDefault() {
}

var PaymentServiceRefundReviewArgs_Request_DEFAULT *RefundReviewRequest

func (p *PaymentServiceRefundReviewArgs) GetRequest() (v *RefundReviewRequest) {
	if !p.IsSetRequest() {
		return PaymentServiceRefundReviewArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PaymentServiceRefundReviewArgs) SetRequest(val *RefundReviewRequest) {
	p.Request = val
}

func (p *PaymentServiceRefundReviewArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PaymentServiceRefundReviewArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceRefundReviewArgs(%+v)", *p)
}

func (p *PaymentServiceRefundReviewArgs) DeepEqual(ano *PaymentServiceRefundReviewArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *PaymentServiceRefundReviewArgs) Field1DeepEqual(src *RefundReviewRequest) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentServiceRefundReviewArgs = map[int16]string{
	1: "request",
}

type PaymentServiceRefundReviewResult struct {
	Success *RefundReviewResponse `thrift:"success,0,optional" frugal:"0,optional,RefundReviewResponse" json:"success,omitempty"`
}

func NewPaymentServiceRefundReviewResult() *PaymentServiceRefundReviewResult {
	return &PaymentServiceRefundReviewResult{}
}

func (p *PaymentServiceRefundReviewResult) InitDefault() {
}

var PaymentServiceRefundReviewResult_Success_DEFAULT *RefundReviewResponse

func (p *PaymentServiceRefundReviewResult) GetSuccess() (v *RefundReviewResponse) {
	if !p.IsSetSuccess() {
		return PaymentServiceRefundReviewResult_Success_DEFAULT
	}
	return p.Success
}
func (p *PaymentServiceRefundReviewResult) SetSuccess(x interface{}) {
	p.Success = x.(*RefundReviewResponse)
}

func (p *PaymentServiceRefundReviewResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentServiceRefundReviewResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceRefundReviewResult(%+v)", *p)
}

func (p *PaymentServiceRefundReviewResult) DeepEqual(ano *PaymentServiceRefundReviewResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *PaymentServiceRefundReviewResult) Field0DeepEqual(src *RefundReviewResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentServiceRefundReviewResult = map[int16]string{
	0: "success",
}

type PaymentServiceRequestRefundArgs struct {
	Request *RefundRequest `thrift:"request,1" frugal:"1,default,RefundRequest" json:"request"`
}

func NewPaymentServiceRequestRefundArgs() *PaymentServiceRequestRefundArgs {
	return &PaymentServiceRequestRefundArgs{}
}

func (p *PaymentServiceRequestRefundArgs) InitDefault() {
}

var PaymentServiceRequestRefundArgs_Request_DEFAULT *RefundRequest

func (p *PaymentServiceRequestRefundArgs) GetRequest() (v *RefundRequest) {
	if !p.IsSetRequest() {
		return PaymentServiceRequestRefundArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PaymentServiceRequestRefundArgs) SetRequest(val *RefundRequest) {
	p.Request = val
}

func (p *PaymentServiceRequestRefundArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PaymentServiceRequestRefundArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceRequestRefundArgs(%+v)", *p)
}

func (p *PaymentServiceRequestRefundArgs) DeepEqual(ano *PaymentServiceRequestRefundArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *PaymentServiceRequestRefundArgs) Field1DeepEqual(src *RefundRequest) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentServiceRequestRefundArgs = map[int16]string{
	1: "request",
}

type PaymentServiceRequestRefundResult struct {
	Success *RefundResponse `thrift:"success,0,optional" frugal:"0,optional,RefundResponse" json:"success,omitempty"`
}

func NewPaymentServiceRequestRefundResult() *PaymentServiceRequestRefundResult {
	return &PaymentServiceRequestRefundResult{}
}

func (p *PaymentServiceRequestRefundResult) InitDefault() {
}

var PaymentServiceRequestRefundResult_Success_DEFAULT *RefundResponse

func (p *PaymentServiceRequestRefundResult) GetSuccess() (v *RefundResponse) {
	if !p.IsSetSuccess() {
		return PaymentServiceRequestRefundResult_Success_DEFAULT
	}
	return p.Success
}
func (p *PaymentServiceRequestRefundResult) SetSuccess(x interface{}) {
	p.Success = x.(*RefundResponse)
}

func (p *PaymentServiceRequestRefundResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentServiceRequestRefundResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceRequestRefundResult(%+v)", *p)
}

func (p *PaymentServiceRequestRefundResult) DeepEqual(ano *PaymentServiceRequestRefundResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *PaymentServiceRequestRefundResult) Field0DeepEqual(src *RefundResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_PaymentServiceRequestRefundResult = map[int16]string{
	0: "success",
}
