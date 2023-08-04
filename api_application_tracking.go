/*
Neucore API

Client library of Neucore API

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ApplicationTrackingApiService ApplicationTrackingApi service
type ApplicationTrackingApiService service

type ApiMemberTrackingV1Request struct {
	ctx context.Context
	ApiService *ApplicationTrackingApiService
	id int32
	inactive *int32
	active *int32
	account *string
}

// Limit to members who have been inactive for x days or longer.
func (r ApiMemberTrackingV1Request) Inactive(inactive int32) ApiMemberTrackingV1Request {
	r.inactive = &inactive
	return r
}

// Limit to members who were active in the last x days.
func (r ApiMemberTrackingV1Request) Active(active int32) ApiMemberTrackingV1Request {
	r.active = &active
	return r
}

// Limit to members with (true) or without (false) an account.
func (r ApiMemberTrackingV1Request) Account(account string) ApiMemberTrackingV1Request {
	r.account = &account
	return r
}

func (r ApiMemberTrackingV1Request) Execute() ([]CorporationMember, *http.Response, error) {
	return r.ApiService.MemberTrackingV1Execute(r)
}

/*
MemberTrackingV1 Return corporation member tracking data.

Needs role: app-tracking

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id EVE corporation ID.
 @return ApiMemberTrackingV1Request
*/
func (a *ApplicationTrackingApiService) MemberTrackingV1(ctx context.Context, id int32) ApiMemberTrackingV1Request {
	return ApiMemberTrackingV1Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return []CorporationMember
func (a *ApplicationTrackingApiService) MemberTrackingV1Execute(r ApiMemberTrackingV1Request) ([]CorporationMember, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CorporationMember
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationTrackingApiService.MemberTrackingV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/app/v1/corporation/{id}/member-tracking"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.inactive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inactive", r.inactive, "")
	}
	if r.active != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "active", r.active, "")
	}
	if r.account != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account", r.account, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
