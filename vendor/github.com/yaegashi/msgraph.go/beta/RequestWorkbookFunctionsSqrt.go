// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsSqrtRequestBuilder struct{ BaseRequestBuilder }

// Sqrt action undocumented
func (b *WorkbookFunctionsRequestBuilder) Sqrt(reqObj *WorkbookFunctionsSqrtRequestParameter) *WorkbookFunctionsSqrtRequestBuilder {
	bb := &WorkbookFunctionsSqrtRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sqrt"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsSqrtRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsSqrtRequestBuilder) Request() *WorkbookFunctionsSqrtRequest {
	return &WorkbookFunctionsSqrtRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsSqrtRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsSqrtPiRequestBuilder struct{ BaseRequestBuilder }

// SqrtPi action undocumented
func (b *WorkbookFunctionsRequestBuilder) SqrtPi(reqObj *WorkbookFunctionsSqrtPiRequestParameter) *WorkbookFunctionsSqrtPiRequestBuilder {
	bb := &WorkbookFunctionsSqrtPiRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sqrtPi"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsSqrtPiRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsSqrtPiRequestBuilder) Request() *WorkbookFunctionsSqrtPiRequest {
	return &WorkbookFunctionsSqrtPiRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsSqrtPiRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}