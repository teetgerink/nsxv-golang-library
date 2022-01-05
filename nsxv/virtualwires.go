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

type VirtualwiresApiService service

/*
VirtualwiresApiService logicalSwitchList
Retrieve information about all logical switches in the specified transport zone (network scope).   Parameters:  scopeId: A valid transport zone ID (vdnScope objectId).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param scopeId
 * @param optional nil or *VirtualwiresApiVdnScopesScopeIdVirtualwiresGetOpts - Optional Parameters:
     * @param "Startindex" (optional.String) -
     * @param "Pagesize" (optional.String) -

*/

type VirtualwiresApiVdnScopesScopeIdVirtualwiresGetOpts struct {
	Startindex optional.String
	Pagesize   optional.String
}

func (a *VirtualwiresApiService) VdnScopesScopeIdVirtualwiresGet(ctx context.Context, scopeId string, localVarOptionals *VirtualwiresApiVdnScopesScopeIdVirtualwiresGetOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/scopes/{scopeId}/virtualwires"
	localVarPath = strings.Replace(localVarPath, "{"+"scopeId"+"}", fmt.Sprintf("%v", scopeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Startindex.IsSet() {
		localVarQueryParams.Add("startindex", parameterToString(localVarOptionals.Startindex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pagesize.IsSet() {
		localVarQueryParams.Add("pagesize", parameterToString(localVarOptionals.Pagesize.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
VirtualwiresApiService logicalSwitchCreate
Create a logical switch.  To create a universal logical switch use *universalvdnscope* as the scopeId in the URI and send the request to the primary NSX Manager. Request body parameters:   * **name** - Optional. The name of the logical switch.   * **description** - Optional. Description of the logical switch.   * **tenantId** - Required.   * **controlPlaneMode** - Optional. The control plane mode. If not     specified, the **controlPlaneMode** of the transport zone is used. It     can be one of the following:       * *UNICAST_MODE*       * *HYBRID_MODE*       * *MULTICAST_MODE*   * **guestVlanAllowed** - Optional. Default is *false*.   Parameters:  scopeId: A valid transport zone ID (vdnScope objectId).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param scopeId
 * @param optional nil or *VirtualwiresApiVdnScopesScopeIdVirtualwiresPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -

*/

type VirtualwiresApiVdnScopesScopeIdVirtualwiresPostOpts struct {
	Body        optional.Interface
	ContentType optional.String
}

func (a *VirtualwiresApiService) VdnScopesScopeIdVirtualwiresPost(ctx context.Context, scopeId string, localVarOptionals *VirtualwiresApiVdnScopesScopeIdVirtualwiresPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/scopes/{scopeId}/virtualwires"
	localVarPath = strings.Replace(localVarPath, "{"+"scopeId"+"}", fmt.Sprintf("%v", scopeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentType.IsSet() {
		localVarHeaderParams["Content-Type"] = parameterToString(localVarOptionals.ContentType.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
VirtualwiresApiService logicalSwitchesRead
Retrieve information about all logical switches in all transport zones.   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *VirtualwiresApiVdnVirtualwiresGetOpts - Optional Parameters:
     * @param "Startindex" (optional.String) -
     * @param "Pagesize" (optional.String) -

*/

type VirtualwiresApiVdnVirtualwiresGetOpts struct {
	Startindex optional.String
	Pagesize   optional.String
}

func (a *VirtualwiresApiService) VdnVirtualwiresGet(ctx context.Context, localVarOptionals *VirtualwiresApiVdnVirtualwiresGetOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/virtualwires"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Startindex.IsSet() {
		localVarQueryParams.Add("startindex", parameterToString(localVarOptionals.Startindex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pagesize.IsSet() {
		localVarQueryParams.Add("pagesize", parameterToString(localVarOptionals.Pagesize.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
VirtualwiresApiService logicalSwitchPortGroupFixAction
For every logical switch created, NSX creates a corresponding port group in vCenter. If the port group is missing, the logical switch will stop functioning.  If the port group backing a logical switch is deleted, you can recreate a new backing port group for the logical switch.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param virtualWireID
 * @param optional nil or *VirtualwiresApiVdnVirtualwiresVirtualWireIDBackingPostOpts - Optional Parameters:
     * @param "Action" (optional.String) -

*/

type VirtualwiresApiVdnVirtualwiresVirtualWireIDBackingPostOpts struct {
	Action optional.String
}

func (a *VirtualwiresApiService) VdnVirtualwiresVirtualWireIDBackingPost(ctx context.Context, virtualWireID string, localVarOptionals *VirtualwiresApiVdnVirtualwiresVirtualWireIDBackingPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/virtualwires/{virtualWireID}/backing"
	localVarPath = strings.Replace(localVarPath, "{"+"virtualWireID"+"}", fmt.Sprintf("%v", virtualWireID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Action.IsSet() {
		localVarQueryParams.Add("action", parameterToString(localVarOptionals.Action.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
VirtualwiresApiService logicalSwitchConnCheckExecute
Test multicast group connectivity.  Test multicast group connectivity between two hosts connected to the specified logical switch.  Parameter **packetSizeMode** has one of the following values: * *0* - VXLAN standard packet size * *1* - minimum packet size * *2* - customized packet size. If you set **packetSizeMode** to *2*, you must specify the size using the **packetSize** parameter.   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param virtualWireID
 * @param optional nil or *VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckMulticastPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -

*/

type VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckMulticastPostOpts struct {
	Body        optional.Interface
	ContentType optional.String
}

func (a *VirtualwiresApiService) VdnVirtualwiresVirtualWireIDConnCheckMulticastPost(ctx context.Context, virtualWireID string, localVarOptionals *VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckMulticastPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/virtualwires/{virtualWireID}/conn-check/multicast"
	localVarPath = strings.Replace(localVarPath, "{"+"virtualWireID"+"}", fmt.Sprintf("%v", virtualWireID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentType.IsSet() {
		localVarHeaderParams["Content-Type"] = parameterToString(localVarOptionals.ContentType.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
VirtualwiresApiService logicalSwitchPingExecute
Test point-to-point connectivity.  Test point-to-point connectivity between two hosts connected to the specified logical switch.  Parameter **packetSizeMode** has one of the following values: * *0* - VXLAN standard packet size * *1* - minimum packet size * *2* - customized packet size. If you set **packetSizeMode** to *2*, you must specify the size using the **packetSize** parameter.   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param virtualWireID
 * @param optional nil or *VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckP2pPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -

*/

type VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckP2pPostOpts struct {
	Body        optional.Interface
	ContentType optional.String
}

func (a *VirtualwiresApiService) VdnVirtualwiresVirtualWireIDConnCheckP2pPost(ctx context.Context, virtualWireID string, localVarOptionals *VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckP2pPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/virtualwires/{virtualWireID}/conn-check/p2p"
	localVarPath = strings.Replace(localVarPath, "{"+"virtualWireID"+"}", fmt.Sprintf("%v", virtualWireID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentType.IsSet() {
		localVarHeaderParams["Content-Type"] = parameterToString(localVarOptionals.ContentType.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
VirtualwiresApiService logicalSwitchDelete
Delete the specified logical switch.  Parameters:  virtualWireID: A logical switch id, e.g. virtualwire-1002
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param virtualWireID

*/
func (a *VirtualwiresApiService) VdnVirtualwiresVirtualWireIDDelete(ctx context.Context, virtualWireID string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/virtualwires/{virtualWireID}"
	localVarPath = strings.Replace(localVarPath, "{"+"virtualWireID"+"}", fmt.Sprintf("%v", virtualWireID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
VirtualwiresApiService logicalSwitchShow
Retrieve information about the specified logical switch.  If the switch is a universal logical switch the **isUniversal** parameter is set to true in the response body.   Parameters:  virtualWireID: A logical switch id, e.g. virtualwire-1002
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param virtualWireID

*/
func (a *VirtualwiresApiService) VdnVirtualwiresVirtualWireIDGet(ctx context.Context, virtualWireID string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/virtualwires/{virtualWireID}"
	localVarPath = strings.Replace(localVarPath, "{"+"virtualWireID"+"}", fmt.Sprintf("%v", virtualWireID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
VirtualwiresApiService logicalSwitchHardwareGatewayBindingsList
Retrieve hardware gateway bindings for the specified logical switch.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param virtualWireID

*/
func (a *VirtualwiresApiService) VdnVirtualwiresVirtualWireIDHardwaregatewaysGet(ctx context.Context, virtualWireID string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/virtualwires/{virtualWireID}/hardwaregateways"
	localVarPath = strings.Replace(localVarPath, "{"+"virtualWireID"+"}", fmt.Sprintf("%v", virtualWireID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
VirtualwiresApiService logicalSwitchHardwareGatewayBindingCreate
Manage the connection between a hardware gateway and a logical switch.  ### Attach a hardware gateway to a logical switch and create a new binding with the information provided  &#x60;POST /api/2.0/vdn/virtualwires/{virtualwireid}/hardwaregateways&#x60;  &#x60;&#x60;&#x60; &lt;hardwareGatewayBinding&gt;   &lt;hardwareGatewayId&gt;hardwarewgateway1&lt;/hardwareGatewayId&gt;   &lt;vlan&gt;v1&lt;/vlan&gt;   &lt;switchName&gt;s1&lt;/switchName&gt;   &lt;portName&gt;s1&lt;/portName&gt; &lt;/hardwareGatewayBinding&gt;  &#x60;&#x60;&#x60;  ### Attach a hardware gateway to a logical switch, specifying an existing binding by ID  &#x60;POST /api/2.0/vdn/virtualwires/&lt;virtualwireId&gt;/hardwaregateways/{bindingId}?action&#x3D;attach&#x60;  &#x60;&#x60;&#x60; &lt;virtualWire&gt;   ...   &lt;hardwareGatewayBindings&gt;     &lt;hardwareGatewayBinding&gt;       &lt;id&gt;binding id&lt;/id&gt;     &lt;/hardwareGatewayBinding&gt;   &lt;/hardwareGatewayBindings&gt; &lt;/virtualWire&gt; &#x60;&#x60;&#x60;  ### Detach a hardware gateway from a logical switch  &#x60;POST /api/2.0/vdn/virtualwires/&lt;virtualwireId&gt;/hardwaregateways/{bindingId}?action&#x3D;detach&#x60;  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  hardwareGatewayBindingId: Hardware Gateway Binding ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param virtualWireID
 * @param hardwareGatewayBindingId
 * @param optional nil or *VirtualwiresApiVdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -
     * @param "Action" (optional.String) -

*/

type VirtualwiresApiVdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPostOpts struct {
	Body        optional.Interface
	ContentType optional.String
	Action      optional.String
}

func (a *VirtualwiresApiService) VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPost(ctx context.Context, virtualWireID string, hardwareGatewayBindingId string, localVarOptionals *VirtualwiresApiVdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/virtualwires/{virtualWireID}/hardwaregateways/{hardwareGatewayBindingId}"
	localVarPath = strings.Replace(localVarPath, "{"+"virtualWireID"+"}", fmt.Sprintf("%v", virtualWireID), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"hardwareGatewayBindingId"+"}", fmt.Sprintf("%v", hardwareGatewayBindingId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Action.IsSet() {
		localVarQueryParams.Add("action", parameterToString(localVarOptionals.Action.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentType.IsSet() {
		localVarHeaderParams["Content-Type"] = parameterToString(localVarOptionals.ContentType.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
VirtualwiresApiService logicalSwitchUpdate
Update the specified logical switch.  For example, you can update the name, description, or control plane mode.   Parameters:  virtualWireID: A logical switch id, e.g. virtualwire-1002
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param virtualWireID
 * @param optional nil or *VirtualwiresApiVdnVirtualwiresVirtualWireIDPutOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -

*/

type VirtualwiresApiVdnVirtualwiresVirtualWireIDPutOpts struct {
	Body        optional.Interface
	ContentType optional.String
}

func (a *VirtualwiresApiService) VdnVirtualwiresVirtualWireIDPut(ctx context.Context, virtualWireID string, localVarOptionals *VirtualwiresApiVdnVirtualwiresVirtualWireIDPutOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/virtualwires/{virtualWireID}"
	localVarPath = strings.Replace(localVarPath, "{"+"virtualWireID"+"}", fmt.Sprintf("%v", virtualWireID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentType.IsSet() {
		localVarHeaderParams["Content-Type"] = parameterToString(localVarOptionals.ContentType.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
VirtualwiresApiService logicalSwitchVmAttachCreate
Attach a VM vNIC to, or detach a VM vNIC from a logical switch.  Specify the logical switch ID in the **portgroupId** parameter. To detach a VM vNIC from a logical switch, leave the **portgroupId** parameter empty.  To find the ID of a VM vNIC, do the following: 1. In the vSphere MOB, navigate to the VM you want to connect or disconnect. 2. Click **config** and take note of the **instanceUuid**. 3. Click **hardware** and take note of the last three digits of the appropriate network interface device.  Use these two values to form the VM vNIC ID.  For example, if the **instanceUuid** is *502e71fa-1a00-759b-e40f-ce778e915f16* and the appropriate **device** value is *device[4000]*, the **objectId** and **vnicUuid** are both *502e71fa-1a00-759b-e40f-ce778e915f16.000*.   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *VirtualwiresApiVdnVirtualwiresVmVnicPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -

*/

type VirtualwiresApiVdnVirtualwiresVmVnicPostOpts struct {
	Body        optional.Interface
	ContentType optional.String
}

func (a *VirtualwiresApiService) VdnVirtualwiresVmVnicPost(ctx context.Context, localVarOptionals *VirtualwiresApiVdnVirtualwiresVmVnicPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/vdn/virtualwires/vm/vnic"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentType.IsSet() {
		localVarHeaderParams["Content-Type"] = parameterToString(localVarOptionals.ContentType.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
