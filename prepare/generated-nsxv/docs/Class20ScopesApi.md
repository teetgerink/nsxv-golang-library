# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20VdnScopesGet**](Class20ScopesApi.md#Api20VdnScopesGet) | **Get** /api/2.0/vdn/scopes | vdnScopesList
[**Api20VdnScopesPost**](Class20ScopesApi.md#Api20VdnScopesPost) | **Post** /api/2.0/vdn/scopes | vdnScopeCreate
[**Api20VdnScopesScopeIdAttributesPut**](Class20ScopesApi.md#Api20VdnScopesScopeIdAttributesPut) | **Put** /api/2.0/vdn/scopes/{scopeId}/attributes | vdnScopeAttribUpdateUpdate
[**Api20VdnScopesScopeIdConnCheckMulticastPost**](Class20ScopesApi.md#Api20VdnScopesScopeIdConnCheckMulticastPost) | **Post** /api/2.0/vdn/scopes/{scopeId}/conn-check/multicast | vdnScopeConnCheckExecute
[**Api20VdnScopesScopeIdDelete**](Class20ScopesApi.md#Api20VdnScopesScopeIdDelete) | **Delete** /api/2.0/vdn/scopes/{scopeId} | vdnScopeDelete
[**Api20VdnScopesScopeIdGet**](Class20ScopesApi.md#Api20VdnScopesScopeIdGet) | **Get** /api/2.0/vdn/scopes/{scopeId} | vdnScopeProperties
[**Api20VdnScopesScopeIdPost**](Class20ScopesApi.md#Api20VdnScopesScopeIdPost) | **Post** /api/2.0/vdn/scopes/{scopeId} | vdnScopeChange

# **Api20VdnScopesGet**
> Api20VdnScopesGet(ctx, )
vdnScopesList

Retrieve information about all transport zones (also known as network scopes).   Parameters:  

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnScopesPost**
> Api20VdnScopesPost(ctx, optional)
vdnScopeCreate

Create a transport zone.  Request body parameters:    * **name** - Required. The name of the transport zone.   * **description** - Optional. Description of the transport zone.   * **objectId** - Required. The cluster object ID from vSphere. One or more are     required.   * **controlPlaneMode** - Optional. The control plane mode. It can be     one of the following:       * *UNICAST_MODE*       * *HYBRID_MODE*       * *MULTICAST_MODE*   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20ScopesApiApi20VdnScopesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20ScopesApiApi20VdnScopesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **isUniversal** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnScopesScopeIdAttributesPut**
> Api20VdnScopesScopeIdAttributesPut(ctx, scopeId, optional)
vdnScopeAttribUpdateUpdate

Update the attributes of a transport zone.  For example, you can update the name, description, or control plane mode. You must include the cluster object IDs for the transport zone in the request body.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***Class20ScopesApiApi20VdnScopesScopeIdAttributesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20ScopesApiApi20VdnScopesScopeIdAttributesPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnScopesScopeIdConnCheckMulticastPost**
> Api20VdnScopesScopeIdConnCheckMulticastPost(ctx, scopeId, optional)
vdnScopeConnCheckExecute

Test multicast group connectivity.  Test multicast group connectivity between two hosts connected to the specified transport zone.  Parameter **packetSizeMode** has one of the following values: * *0* - VXLAN standard packet size * *1* - minimum packet size * *2* - customized packet size. If you set **packetSizeMode** to *2*, you must specify the size using the **packetSize** parameter.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***Class20ScopesApiApi20VdnScopesScopeIdConnCheckMulticastPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20ScopesApiApi20VdnScopesScopeIdConnCheckMulticastPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnScopesScopeIdDelete**
> Api20VdnScopesScopeIdDelete(ctx, scopeId)
vdnScopeDelete

Delete the specified transport zone.   Parameters:  scopeId: A valid transport zone ID (vdnScope objectId)  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnScopesScopeIdGet**
> Api20VdnScopesScopeIdGet(ctx, scopeId)
vdnScopeProperties

Retrieve information about the specified transport zone.   Parameters:  scopeId: A valid transport zone ID (vdnScope objectId)  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnScopesScopeIdPost**
> Api20VdnScopesScopeIdPost(ctx, scopeId, optional)
vdnScopeChange

Update the specified transport zone.  You can add a cluster to or delete a cluster from a transport zone.  You can also repair missing portgroups. For every logical switch created, NSX creates a corresponding portgroup in vCenter. If the portgroup is lost for any reason, the logical switch will stop functioning. The repair action recreates any missing portgroups.   Parameters:  scopeId: A valid transport zone ID (vdnScope objectId)  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***Class20ScopesApiApi20VdnScopesScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20ScopesApiApi20VdnScopesScopeIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **action** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

