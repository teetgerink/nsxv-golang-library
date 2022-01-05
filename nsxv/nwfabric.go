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

type NwfabricApiService service

/*
NwfabricApiService nwfabricClustersDelete
Delete locale ID for the specified cluster.  Parameters:  clusterID: Cluster ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterID

*/
func (a *NwfabricApiService) Api20NwfabricClustersClusterIDDelete(ctx context.Context, clusterID string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/clusters/{clusterID}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterID"+"}", fmt.Sprintf("%v", clusterID), -1)

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
NwfabricApiService nwfabricClustersRead
Retrieve the locale ID for the specified cluster.  Parameters:  clusterID: Cluster ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterID

*/
func (a *NwfabricApiService) Api20NwfabricClustersClusterIDGet(ctx context.Context, clusterID string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/clusters/{clusterID}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterID"+"}", fmt.Sprintf("%v", clusterID), -1)

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
NwfabricApiService nwfabricClustersUpdate
Update the locale ID for the specified cluster.  Parameters:  clusterID: Cluster ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterID
 * @param optional nil or *NwfabricApiApi20NwfabricClustersClusterIDPutOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -

*/

type NwfabricApiApi20NwfabricClustersClusterIDPutOpts struct {
	Body        optional.Interface
	ContentType optional.String
}

func (a *NwfabricApiService) Api20NwfabricClustersClusterIDPut(ctx context.Context, clusterID string, localVarOptionals *NwfabricApiApi20NwfabricClustersClusterIDPutOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/clusters/{clusterID}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterID"+"}", fmt.Sprintf("%v", clusterID), -1)

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
NwfabricApiService nwfabricConfigDelete
Remove VXLAN or network virtualization components.  Removing network virtualization components removes previously installed VIBs, tears down NSX Manager to ESXi messaging, and removes any other network fabric dependent features such as logical switches. If a feature such as logical switches is being used in your environment, this call fails.  Removing VXLAN does not remove the network virtualization components from the cluster.  | Name | Comments | |------|----------| |**resourceId** | vCenter MOB ID of cluster. For example, domain-7.| |**featureId** | Feature to act upon. Omit for network virtualization components operations. Use *com.vmware.vshield.vsm.vxlan* for VXLAN operations.|  ### Remove Network Virtualization Components  &#x60;&#x60;&#x60; &lt;nwFabricFeatureConfig&gt;   &lt;resourceConfig&gt;     &lt;resourceId&gt;CLUSTER MOID&lt;/resourceId&gt;   &lt;/resourceConfig&gt; &lt;/nwFabricFeatureConfig&gt; &#x60;&#x60;&#x60;  ### Remove VXLAN  &#x60;&#x60;&#x60; &lt;nwFabricFeatureConfig&gt;   &lt;featureId&gt;com.vmware.vshield.vsm.vxlan&lt;/featureId&gt;   &lt;resourceConfig&gt;     &lt;resourceId&gt;CLUSTER MOID&lt;/resourceId&gt;    &lt;/resourceConfig&gt; &lt;/nwFabricFeatureConfig&gt; &#x60;&#x60;&#x60;  ### Remove VXLAN with vDS context  &#x60;&#x60;&#x60; &lt;nwFabricFeatureConfig&gt;   &lt;featureId&gt;com.vmware.vshield.vsm.vxlan&lt;/featureId&gt;   &lt;resourceConfig&gt;     &lt;resourceId&gt;CLUSTER MOID&lt;/resourceId&gt;     &lt;configSpec class&#x3D;\&quot;map\&quot;&gt;       &lt;entry&gt;         &lt;keyclass&#x3D;\&quot;java.lang.String\&quot;&gt;vxlan&lt;/key&gt;         &lt;valueclass&#x3D;\&quot;java.lang.String\&quot;&gt;cascadeDeleteVdsContext&lt;/value&gt;       &lt;/entry&gt;     &lt;/configSpec&gt;   &lt;/resourceConfig&gt; &lt;/nwFabricFeatureConfig&gt; &#x60;&#x60;&#x60;   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *NwfabricApiApi20NwfabricConfigureDeleteOpts - Optional Parameters:
     * @param "ContentType" (optional.String) -

*/

type NwfabricApiApi20NwfabricConfigureDeleteOpts struct {
	ContentType optional.String
}

func (a *NwfabricApiService) Api20NwfabricConfigureDelete(ctx context.Context, localVarOptionals *NwfabricApiApi20NwfabricConfigureDeleteOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/configure"

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
	if localVarOptionals != nil && localVarOptionals.ContentType.IsSet() {
		localVarHeaderParams["Content-Type"] = parameterToString(localVarOptionals.ContentType.Value(), "")
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
NwfabricApiService nwfabricConfigCreate
Install network fabric or VXLAN.  This method can be used to perform the following tasks:  * Install Network Virtualization Components * Configure VXLAN * Configure VXLAN with LACPv2 * Reset Communication Between NSX Manager and a Host or Cluster  **Parameter Information**  | Name | Comments | |------|----------| |**resourceId** | vCenter MOB ID of cluster. For example, *domain-7*. A host can be specified when resetting communication. For example, *host-24*. | |**featureId** | Feature to act upon. Omit for network virtualization components operations. Use *com.vmware.vshield.vsm.vxlan* for VXLAN operations, *com.vmware.vshield.vsm.messagingInfra* for message bus operations.| |**ipPoolId** | Used for VXLAN installation. If not specified, DHCP is used for VTEP address assignment.| |**teaming** | Used for VXLAN installation. Options are *FAILOVER_ORDER*, *ETHER_CHANNEL*, *LACP_ACTIVE*, *LACP_PASSIVE*, *LOADBALANCE_LOADBASED*, *LOADBALANCE_SRCID*, *LOADBALANCE_SRCMAC*, *LACP_V2*| |**uplinkPortName** | The *uplinkPortName* as specified in vCenter.|  ### Install Network Virtualization Components  &#x60;POST /api/2.0/nwfabric/configure&#x60;  &#x60;&#x60;&#x60; &lt;nwFabricFeatureConfig&gt;   &lt;resourceConfig&gt;     &lt;resourceId&gt;CLUSTER MOID&lt;/resourceId&gt;   &lt;/resourceConfig&gt; &lt;/nwFabricFeatureConfig&gt; &#x60;&#x60;&#x60;  ### Configure VXLAN  &#x60;POST /api/2.0/nwfabric/configure&#x60;  &#x60;&#x60;&#x60; &lt;nwFabricFeatureConfig&gt;   &lt;featureId&gt;com.vmware.vshield.vsm.vxlan&lt;/featureId&gt;   &lt;resourceConfig&gt;     &lt;resourceId&gt;CLUSTER MOID&lt;/resourceId&gt;     &lt;configSpec class&#x3D;\&quot;clusterMappingSpec\&quot;&gt;       &lt;switch&gt;         &lt;objectId&gt;DVS MOID&lt;/objectId&gt;&lt;/switch&gt;         &lt;vlanId&gt;0&lt;/vlanId&gt;         &lt;vmknicCount&gt;1&lt;/vmknicCount&gt;         &lt;ipPoolId&gt;IPADDRESSPOOL ID&lt;/ipPoolId&gt;     &lt;/configSpec&gt;   &lt;/resourceConfig&gt;   &lt;resourceConfig&gt;     &lt;resourceId&gt;DVS MOID&lt;/resourceId&gt;     &lt;configSpec class&#x3D;\&quot;vdsContext\&quot;&gt;       &lt;switch&gt;           &lt;objectId&gt;DVS MOID&lt;/objectId&gt;       &lt;/switch&gt;       &lt;mtu&gt;1600&lt;/mtu&gt;       &lt;teaming&gt;ETHER_CHANNEL&lt;/teaming&gt;     &lt;/configSpec&gt;   &lt;/resourceConfig&gt; &lt;/nwFabricFeatureConfig&gt; &#x60;&#x60;&#x60;  ### Configure VXLAN with LACPv2  &#x60;POST /api/2.0/nwfabric/configure&#x60;  &#x60;&#x60;&#x60; &lt;nwFabricFeatureConfig&gt;   &lt;featureId&gt;com.vmware.vshield.nsxmgr.vxlan&lt;/featureId&gt;   &lt;resourceConfig&gt;     &lt;resourceId&gt;CLUSTER MOID&lt;/resourceId&gt;     &lt;configSpec class&#x3D;\&quot;clusterMappingSpec\&quot;&gt;       &lt;switch&gt;         &lt;objectId&gt;DVS MOID&lt;/objectId&gt;       &lt;/switch&gt;       &lt;vlanId&gt;0&lt;/vlanId&gt;       &lt;vmknicCount&gt;1&lt;/vmknicCount&gt;     &lt;/configSpec&gt;   &lt;/resourceConfig&gt;   &lt;resourceConfig&gt;     &lt;resourceId&gt;DVS MOID&lt;/resourceId&gt;     &lt;configSpec class&#x3D;\&quot;vdsContext\&quot;&gt;       &lt;switch&gt;         &lt;objectId&gt;DVS MOID&lt;/objectId&gt;       &lt;/switch&gt;       &lt;mtu&gt;1600&lt;/mtu&gt;       &lt;teaming&gt;LACP_V2&lt;/teaming&gt;       &lt;uplinkPortName&gt;LAG NAME&lt;/uplinkPortName&gt;     &lt;/configSpec&gt;   &lt;/resourceConfig&gt; &lt;/nwFabricFeatureConfig&gt; &#x60;&#x60;&#x60;  ### Reset Communication Between NSX Manager and a Host or Cluster  &#x60;POST /api/2.0/nwfabric/configure?action&#x3D;synchronize&#x60;  &#x60;&#x60;&#x60;  &lt;nwFabricFeatureConfig&gt;   &lt;featureId&gt;com.vmware.vshield.vsm.messagingInfra&lt;/featureId&gt;   &lt;resourceConfig&gt;     &lt;resourceId&gt;resourceId&lt;/resourceId&gt;   &lt;/resourceConfig&gt; &lt;/nwFabricFeatureConfig&gt;  &#x60;&#x60;&#x60;   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *NwfabricApiApi20NwfabricConfigurePostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -
     * @param "Action" (optional.String) -

*/

type NwfabricApiApi20NwfabricConfigurePostOpts struct {
	Body        optional.Interface
	ContentType optional.String
	Action      optional.String
}

func (a *NwfabricApiService) Api20NwfabricConfigurePost(ctx context.Context, localVarOptionals *NwfabricApiApi20NwfabricConfigurePostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/configure"

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
NwfabricApiService nwfabricConfigUpdate
Upgrade Network virtualization components. _ This API call can be used to upgrade network virtualization components. After NSX Manager is upgraded, previously prepared clusters must have the 6.x network virtualization components installed.   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *NwfabricApiApi20NwfabricConfigurePutOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -

*/

type NwfabricApiApi20NwfabricConfigurePutOpts struct {
	Body        optional.Interface
	ContentType optional.String
}

func (a *NwfabricApiService) Api20NwfabricConfigurePut(ctx context.Context, localVarOptionals *NwfabricApiApi20NwfabricConfigurePutOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/configure"

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
NwfabricApiService nwfabricFeaturesList
Retrieves all network fabric features available on the cluster. Multiple **featureInfo** sections may be returned.   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

*/
func (a *NwfabricApiService) Api20NwfabricFeaturesGet(ctx context.Context) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/features"

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
NwfabricApiService nwfabricHostsDelete
Delete the locale ID for the specified host.  Parameters:  hostID: Host ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param hostID

*/
func (a *NwfabricApiService) Api20NwfabricHostsHostIDDelete(ctx context.Context, hostID string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/hosts/{hostID}"
	localVarPath = strings.Replace(localVarPath, "{"+"hostID"+"}", fmt.Sprintf("%v", hostID), -1)

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
NwfabricApiService nwfabricHostsRead
Retrieve the locale ID for the specified host.  Parameters:  hostID: Host ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param hostID

*/
func (a *NwfabricApiService) Api20NwfabricHostsHostIDGet(ctx context.Context, hostID string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/hosts/{hostID}"
	localVarPath = strings.Replace(localVarPath, "{"+"hostID"+"}", fmt.Sprintf("%v", hostID), -1)

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
NwfabricApiService nwfabricHostsUpdate
Update the locale ID for the specified host.  Parameters:  hostID: Host ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param hostID
 * @param optional nil or *NwfabricApiApi20NwfabricHostsHostIDPutOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -

*/

type NwfabricApiApi20NwfabricHostsHostIDPutOpts struct {
	Body        optional.Interface
	ContentType optional.String
}

func (a *NwfabricApiService) Api20NwfabricHostsHostIDPut(ctx context.Context, hostID string, localVarOptionals *NwfabricApiApi20NwfabricHostsHostIDPutOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/hosts/{hostID}"
	localVarPath = strings.Replace(localVarPath, "{"+"hostID"+"}", fmt.Sprintf("%v", hostID), -1)

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
NwfabricApiService statusResourceTypeRead
Retrieve status of resources by criterion.   Parameters:  resourceType: Valid resource type
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param resourceType

*/
func (a *NwfabricApiService) Api20NwfabricStatusAlleligibleResourceTypeGet(ctx context.Context, resourceType string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/status/alleligible/{resourceType}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceType"+"}", fmt.Sprintf("%v", resourceType), -1)

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
NwfabricApiService childStatusRead
Retrieve the network fabric status of child resources of the specified resource.   Parameters:  parentResourceID: Parent resource ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param parentResourceID

*/
func (a *NwfabricApiService) Api20NwfabricStatusChildParentResourceIDGet(ctx context.Context, parentResourceID string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/status/child/{parentResourceID}"
	localVarPath = strings.Replace(localVarPath, "{"+"parentResourceID"+"}", fmt.Sprintf("%v", parentResourceID), -1)

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
NwfabricApiService nwfabricStatusRead
Retrieve the network fabric status of the specified resource.   Parameters:
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *NwfabricApiApi20NwfabricStatusGetOpts - Optional Parameters:
     * @param "Resource" (optional.String) -

*/

type NwfabricApiApi20NwfabricStatusGetOpts struct {
	Resource optional.String
}

func (a *NwfabricApiService) Api20NwfabricStatusGet(ctx context.Context, localVarOptionals *NwfabricApiApi20NwfabricStatusGetOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/nwfabric/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Resource.IsSet() {
		localVarQueryParams.Add("resource", parameterToString(localVarOptionals.Resource.Value(), ""))
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
