/*
 * NSX-V
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nsxv

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type SpoofguardApiService service

/*
SpoofguardApiService spoofGuardPoliciesList
Retrieve information about all SpoofGuard policies.  Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

*/
func (a *SpoofguardApiService) ServicesSpoofguardPoliciesGet(ctx context.Context) (*http.Response, error) {
	var (
		lHttpMethod = strings.ToUpper("Get")
		lPostBody   interface{}
		lFileName   string
		lFileBytes  []byte
	)

	// create path and map variables
	lPath := a.client.cfg.BasePath + "/api/4.0/services/spoofguard/policies"

	lHeaderParams := make(map[string]string)
	lQueryParams := url.Values{}
	lFormParams := url.Values{}

	// to determine the Content-Type header
	lHttpContentTypes := []string{}

	// set Content-Type header
	lHttpContentType := selectHeaderContentType(lHttpContentTypes)
	if lHttpContentType != "" {
		lHeaderParams["Content-Type"] = lHttpContentType
	}

	// to determine the Accept header
	lHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	lHttpHeaderAccept := selectHeaderAccept(lHttpHeaderAccepts)
	if lHttpHeaderAccept != "" {
		lHeaderParams["Accept"] = lHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, lPath, lHttpMethod, lPostBody, lHeaderParams, lQueryParams, lFormParams, lFileName, lFileBytes)
	if err != nil {
		return nil, err
	}

	lHttpResponse, err := a.client.callAPI(r)
	if err != nil || lHttpResponse == nil {
		return lHttpResponse, err
	}

	lBody, err := ioutil.ReadAll(lHttpResponse.Body)
	lHttpResponse.Body.Close()
	if err != nil {
		return lHttpResponse, err
	}

	if lHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  lBody,
			error: lHttpResponse.Status,
		}
		return lHttpResponse, newErr
	}

	return lHttpResponse, nil
}

/*
SpoofguardApiService spoofGuardPolicyDelete
Delete the specified SpoofGuard policy.  Parameters:  policyID: SpoofGuard policy ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param policyID

*/
func (a *SpoofguardApiService) ServicesSpoofguardPoliciesPolicyIDDelete(ctx context.Context, policyID string) (*http.Response, error) {
	var (
		lHttpMethod = strings.ToUpper("Delete")
		lPostBody   interface{}
		lFileName   string
		lFileBytes  []byte
	)

	// create path and map variables
	lPath := a.client.cfg.BasePath + "/api/4.0/services/spoofguard/policies/{policyID}"
	lPath = strings.Replace(lPath, "{"+"policyID"+"}", fmt.Sprintf("%v", policyID), -1)

	lHeaderParams := make(map[string]string)
	lQueryParams := url.Values{}
	lFormParams := url.Values{}

	// to determine the Content-Type header
	lHttpContentTypes := []string{}

	// set Content-Type header
	lHttpContentType := selectHeaderContentType(lHttpContentTypes)
	if lHttpContentType != "" {
		lHeaderParams["Content-Type"] = lHttpContentType
	}

	// to determine the Accept header
	lHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	lHttpHeaderAccept := selectHeaderAccept(lHttpHeaderAccepts)
	if lHttpHeaderAccept != "" {
		lHeaderParams["Accept"] = lHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, lPath, lHttpMethod, lPostBody, lHeaderParams, lQueryParams, lFormParams, lFileName, lFileBytes)
	if err != nil {
		return nil, err
	}

	lHttpResponse, err := a.client.callAPI(r)
	if err != nil || lHttpResponse == nil {
		return lHttpResponse, err
	}

	lBody, err := ioutil.ReadAll(lHttpResponse.Body)
	lHttpResponse.Body.Close()
	if err != nil {
		return lHttpResponse, err
	}

	if lHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  lBody,
			error: lHttpResponse.Status,
		}
		return lHttpResponse, newErr
	}

	return lHttpResponse, nil
}

/*
SpoofguardApiService spoofGuardPolicyRead
Retrieve information about the specified SpoofGuard policy.   Parameters:  policyID: SpoofGuard policy ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param policyID

*/
func (a *SpoofguardApiService) ServicesSpoofguardPoliciesPolicyIDGet(ctx context.Context, policyID string) (*http.Response, error) {
	var (
		lHttpMethod = strings.ToUpper("Get")
		lPostBody   interface{}
		lFileName   string
		lFileBytes  []byte
	)

	// create path and map variables
	lPath := a.client.cfg.BasePath + "/api/4.0/services/spoofguard/policies/{policyID}"
	lPath = strings.Replace(lPath, "{"+"policyID"+"}", fmt.Sprintf("%v", policyID), -1)

	lHeaderParams := make(map[string]string)
	lQueryParams := url.Values{}
	lFormParams := url.Values{}

	// to determine the Content-Type header
	lHttpContentTypes := []string{}

	// set Content-Type header
	lHttpContentType := selectHeaderContentType(lHttpContentTypes)
	if lHttpContentType != "" {
		lHeaderParams["Content-Type"] = lHttpContentType
	}

	// to determine the Accept header
	lHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	lHttpHeaderAccept := selectHeaderAccept(lHttpHeaderAccepts)
	if lHttpHeaderAccept != "" {
		lHeaderParams["Accept"] = lHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, lPath, lHttpMethod, lPostBody, lHeaderParams, lQueryParams, lFormParams, lFileName, lFileBytes)
	if err != nil {
		return nil, err
	}

	lHttpResponse, err := a.client.callAPI(r)
	if err != nil || lHttpResponse == nil {
		return lHttpResponse, err
	}

	lBody, err := ioutil.ReadAll(lHttpResponse.Body)
	lHttpResponse.Body.Close()
	if err != nil {
		return lHttpResponse, err
	}

	if lHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  lBody,
			error: lHttpResponse.Status,
		}
		return lHttpResponse, newErr
	}

	return lHttpResponse, nil
}

/*
SpoofguardApiService spoofGuardPolicyUpdate
Modify the specified SpoofGuard policy.  Parameters:  policyID: SpoofGuard policy ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param policyID
 * @param optional nil or *SpoofguardApiServicesSpoofguardPoliciesPolicyIDPutOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -

*/

type SpoofguardApiServicesSpoofguardPoliciesPolicyIDPutOpts struct {
	Body        optional.Interface
	ContentType optional.String
}

func (a *SpoofguardApiService) ServicesSpoofguardPoliciesPolicyIDPut(ctx context.Context, policyID string, lOptionals *SpoofguardApiServicesSpoofguardPoliciesPolicyIDPutOpts) (*http.Response, error) {
	var (
		lHttpMethod = strings.ToUpper("Put")
		lPostBody   interface{}
		lFileName   string
		lFileBytes  []byte
	)

	// create path and map variables
	lPath := a.client.cfg.BasePath + "/api/4.0/services/spoofguard/policies/{policyID}"
	lPath = strings.Replace(lPath, "{"+"policyID"+"}", fmt.Sprintf("%v", policyID), -1)

	lHeaderParams := make(map[string]string)
	lQueryParams := url.Values{}
	lFormParams := url.Values{}

	// to determine the Content-Type header
	lHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	lHttpContentType := selectHeaderContentType(lHttpContentTypes)
	if lHttpContentType != "" {
		lHeaderParams["Content-Type"] = lHttpContentType
	}

	// to determine the Accept header
	lHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	lHttpHeaderAccept := selectHeaderAccept(lHttpHeaderAccepts)
	if lHttpHeaderAccept != "" {
		lHeaderParams["Accept"] = lHttpHeaderAccept
	}
	if lOptionals != nil && lOptionals.ContentType.IsSet() {
		lHeaderParams["Content-Type"] = parameterToString(lOptionals.ContentType.Value(), "")
	}
	// body params
	if lOptionals != nil && lOptionals.Body.IsSet() {

		lOptionalBody := lOptionals.Body.Value()
		lPostBody = &lOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, lPath, lHttpMethod, lPostBody, lHeaderParams, lQueryParams, lFormParams, lFileName, lFileBytes)
	if err != nil {
		return nil, err
	}

	lHttpResponse, err := a.client.callAPI(r)
	if err != nil || lHttpResponse == nil {
		return lHttpResponse, err
	}

	lBody, err := ioutil.ReadAll(lHttpResponse.Body)
	lHttpResponse.Body.Close()
	if err != nil {
		return lHttpResponse, err
	}

	if lHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  lBody,
			error: lHttpResponse.Status,
		}
		return lHttpResponse, newErr
	}

	return lHttpResponse, nil
}

/*
SpoofguardApiService spoofGuardPoliciesCreate
Create a SpoofGuard policy to specify the operation mode for networks.   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *SpoofguardApiServicesSpoofguardPoliciesPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -

*/

type SpoofguardApiServicesSpoofguardPoliciesPostOpts struct {
	Body        optional.Interface
	ContentType optional.String
}

func (a *SpoofguardApiService) ServicesSpoofguardPoliciesPost(ctx context.Context, lOptionals *SpoofguardApiServicesSpoofguardPoliciesPostOpts) (*http.Response, error) {
	var (
		lHttpMethod = strings.ToUpper("Post")
		lPostBody   interface{}
		lFileName   string
		lFileBytes  []byte
	)

	// create path and map variables
	lPath := a.client.cfg.BasePath + "/api/4.0/services/spoofguard/policies"

	lHeaderParams := make(map[string]string)
	lQueryParams := url.Values{}
	lFormParams := url.Values{}

	// to determine the Content-Type header
	lHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	lHttpContentType := selectHeaderContentType(lHttpContentTypes)
	if lHttpContentType != "" {
		lHeaderParams["Content-Type"] = lHttpContentType
	}

	// to determine the Accept header
	lHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	lHttpHeaderAccept := selectHeaderAccept(lHttpHeaderAccepts)
	if lHttpHeaderAccept != "" {
		lHeaderParams["Accept"] = lHttpHeaderAccept
	}
	if lOptionals != nil && lOptionals.ContentType.IsSet() {
		lHeaderParams["Content-Type"] = parameterToString(lOptionals.ContentType.Value(), "")
	}
	// body params
	if lOptionals != nil && lOptionals.Body.IsSet() {

		lOptionalBody := lOptionals.Body.Value()
		lPostBody = &lOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, lPath, lHttpMethod, lPostBody, lHeaderParams, lQueryParams, lFormParams, lFileName, lFileBytes)
	if err != nil {
		return nil, err
	}

	lHttpResponse, err := a.client.callAPI(r)
	if err != nil || lHttpResponse == nil {
		return lHttpResponse, err
	}

	lBody, err := ioutil.ReadAll(lHttpResponse.Body)
	lHttpResponse.Body.Close()
	if err != nil {
		return lHttpResponse, err
	}

	if lHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  lBody,
			error: lHttpResponse.Status,
		}
		return lHttpResponse, newErr
	}

	return lHttpResponse, nil
}

/*
SpoofguardApiService spoofGuardPolicyIPRead
Retrieve IP addresses for the specified state.   Parameters:  policyID: SpoofGuard policy ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param policyID
 * @param optional nil or *SpoofguardApiServicesSpoofguardPolicyIDGetOpts - Optional Parameters:
     * @param "List" (optional.String) -

*/

type SpoofguardApiServicesSpoofguardPolicyIDGetOpts struct {
	List optional.String
}

func (a *SpoofguardApiService) ServicesSpoofguardPolicyIDGet(ctx context.Context, policyID string, lOptionals *SpoofguardApiServicesSpoofguardPolicyIDGetOpts) (*http.Response, error) {
	var (
		lHttpMethod = strings.ToUpper("Get")
		lPostBody   interface{}
		lFileName   string
		lFileBytes  []byte
	)

	// create path and map variables
	lPath := a.client.cfg.BasePath + "/api/4.0/services/spoofguard/{policyID}"
	lPath = strings.Replace(lPath, "{"+"policyID"+"}", fmt.Sprintf("%v", policyID), -1)

	lHeaderParams := make(map[string]string)
	lQueryParams := url.Values{}
	lFormParams := url.Values{}

	if lOptionals != nil && lOptionals.List.IsSet() {
		lQueryParams.Add("list", parameterToString(lOptionals.List.Value(), ""))
	}
	// to determine the Content-Type header
	lHttpContentTypes := []string{}

	// set Content-Type header
	lHttpContentType := selectHeaderContentType(lHttpContentTypes)
	if lHttpContentType != "" {
		lHeaderParams["Content-Type"] = lHttpContentType
	}

	// to determine the Accept header
	lHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	lHttpHeaderAccept := selectHeaderAccept(lHttpHeaderAccepts)
	if lHttpHeaderAccept != "" {
		lHeaderParams["Accept"] = lHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, lPath, lHttpMethod, lPostBody, lHeaderParams, lQueryParams, lFormParams, lFileName, lFileBytes)
	if err != nil {
		return nil, err
	}

	lHttpResponse, err := a.client.callAPI(r)
	if err != nil || lHttpResponse == nil {
		return lHttpResponse, err
	}

	lBody, err := ioutil.ReadAll(lHttpResponse.Body)
	lHttpResponse.Body.Close()
	if err != nil {
		return lHttpResponse, err
	}

	if lHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  lBody,
			error: lHttpResponse.Status,
		}
		return lHttpResponse, newErr
	}

	return lHttpResponse, nil
}

/*
SpoofguardApiService spoofGuardPolicyIPAction
Approve or publish IP addresses.  Parameters:  policyID: SpoofGuard policy ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param policyID
 * @param optional nil or *SpoofguardApiServicesSpoofguardPolicyIDPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -
     * @param "VnicId" (optional.String) -
     * @param "Action" (optional.String) -

*/

type SpoofguardApiServicesSpoofguardPolicyIDPostOpts struct {
	Body        optional.Interface
	ContentType optional.String
	VnicId      optional.String
	Action      optional.String
}

func (a *SpoofguardApiService) ServicesSpoofguardPolicyIDPost(ctx context.Context, policyID string, lOptionals *SpoofguardApiServicesSpoofguardPolicyIDPostOpts) (*http.Response, error) {
	var (
		lHttpMethod = strings.ToUpper("Post")
		lPostBody   interface{}
		lFileName   string
		lFileBytes  []byte
	)

	// create path and map variables
	lPath := a.client.cfg.BasePath + "/api/4.0/services/spoofguard/{policyID}"
	lPath = strings.Replace(lPath, "{"+"policyID"+"}", fmt.Sprintf("%v", policyID), -1)

	lHeaderParams := make(map[string]string)
	lQueryParams := url.Values{}
	lFormParams := url.Values{}

	if lOptionals != nil && lOptionals.VnicId.IsSet() {
		lQueryParams.Add("vnicId", parameterToString(lOptionals.VnicId.Value(), ""))
	}
	if lOptionals != nil && lOptionals.Action.IsSet() {
		lQueryParams.Add("action", parameterToString(lOptionals.Action.Value(), ""))
	}
	// to determine the Content-Type header
	lHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	lHttpContentType := selectHeaderContentType(lHttpContentTypes)
	if lHttpContentType != "" {
		lHeaderParams["Content-Type"] = lHttpContentType
	}

	// to determine the Accept header
	lHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	lHttpHeaderAccept := selectHeaderAccept(lHttpHeaderAccepts)
	if lHttpHeaderAccept != "" {
		lHeaderParams["Accept"] = lHttpHeaderAccept
	}
	if lOptionals != nil && lOptionals.ContentType.IsSet() {
		lHeaderParams["Content-Type"] = parameterToString(lOptionals.ContentType.Value(), "")
	}
	// body params
	if lOptionals != nil && lOptionals.Body.IsSet() {

		lOptionalBody := lOptionals.Body.Value()
		lPostBody = &lOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, lPath, lHttpMethod, lPostBody, lHeaderParams, lQueryParams, lFormParams, lFileName, lFileBytes)
	if err != nil {
		return nil, err
	}

	lHttpResponse, err := a.client.callAPI(r)
	if err != nil || lHttpResponse == nil {
		return lHttpResponse, err
	}

	lBody, err := ioutil.ReadAll(lHttpResponse.Body)
	lHttpResponse.Body.Close()
	if err != nil {
		return lHttpResponse, err
	}

	if lHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  lBody,
			error: lHttpResponse.Status,
		}
		return lHttpResponse, newErr
	}

	return lHttpResponse, nil
}
