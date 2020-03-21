// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// InferenceClassificationOverrideRequestBuilder is request builder for InferenceClassificationOverride
type InferenceClassificationOverrideRequestBuilder struct{ BaseRequestBuilder }

// Request returns InferenceClassificationOverrideRequest
func (b *InferenceClassificationOverrideRequestBuilder) Request() *InferenceClassificationOverrideRequest {
	return &InferenceClassificationOverrideRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InferenceClassificationOverrideRequest is request for InferenceClassificationOverride
type InferenceClassificationOverrideRequest struct{ BaseRequest }

// Get performs GET request for InferenceClassificationOverride
func (r *InferenceClassificationOverrideRequest) Get(ctx context.Context) (resObj *InferenceClassificationOverride, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InferenceClassificationOverride
func (r *InferenceClassificationOverrideRequest) Update(ctx context.Context, reqObj *InferenceClassificationOverride) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InferenceClassificationOverride
func (r *InferenceClassificationOverrideRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}