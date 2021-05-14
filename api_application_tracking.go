/*
 * Neucore API
 *
 * Client library of Neucore API
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ApplicationTrackingApiService ApplicationTrackingApi service
type ApplicationTrackingApiService service

type ApiMemberTrackingV1Request struct {
	ctx _context.Context
	ApiService *ApplicationTrackingApiService
	id int32
	inactive *int32
	active *int32
	account *string
}

func (r ApiMemberTrackingV1Request) Inactive(inactive int32) ApiMemberTrackingV1Request {
	r.inactive = &inactive
	return r
}
func (r ApiMemberTrackingV1Request) Active(active int32) ApiMemberTrackingV1Request {
	r.active = &active
	return r
}
func (r ApiMemberTrackingV1Request) Account(account string) ApiMemberTrackingV1Request {
	r.account = &account
	return r
}

func (r ApiMemberTrackingV1Request) Execute() ([]CorporationMember, *_nethttp.Response, error) {
	return r.ApiService.MemberTrackingV1Execute(r)
}

/*
 * MemberTrackingV1 Return corporation member tracking data.
 * Needs role: app-tracking
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id EVE corporation ID.
 * @return ApiMemberTrackingV1Request
 */
func (a *ApplicationTrackingApiService) MemberTrackingV1(ctx _context.Context, id int32) ApiMemberTrackingV1Request {
	return ApiMemberTrackingV1Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return []CorporationMember
 */
func (a *ApplicationTrackingApiService) MemberTrackingV1Execute(r ApiMemberTrackingV1Request) ([]CorporationMember, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []CorporationMember
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationTrackingApiService.MemberTrackingV1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/app/v1/corporation/{id}/member-tracking"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.inactive != nil {
		localVarQueryParams.Add("inactive", parameterToString(*r.inactive, ""))
	}
	if r.active != nil {
		localVarQueryParams.Add("active", parameterToString(*r.active, ""))
	}
	if r.account != nil {
		localVarQueryParams.Add("account", parameterToString(*r.account, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
