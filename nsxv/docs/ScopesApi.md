# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VdnScopesGet**](ScopesApi.md#VdnScopesGet) | **Get** /api/2.0/vdn/scopes | vdnScopesList
[**VdnScopesPost**](ScopesApi.md#VdnScopesPost) | **Post** /api/2.0/vdn/scopes | vdnScopeCreate
[**VdnScopesScopeIdAttributesPut**](ScopesApi.md#VdnScopesScopeIdAttributesPut) | **Put** /api/2.0/vdn/scopes/{scopeId}/attributes | vdnScopeAttribUpdateUpdate
[**VdnScopesScopeIdConnCheckMulticastPost**](ScopesApi.md#VdnScopesScopeIdConnCheckMulticastPost) | **Post** /api/2.0/vdn/scopes/{scopeId}/conn-check/multicast | vdnScopeConnCheckExecute
[**VdnScopesScopeIdDelete**](ScopesApi.md#VdnScopesScopeIdDelete) | **Delete** /api/2.0/vdn/scopes/{scopeId} | vdnScopeDelete
[**VdnScopesScopeIdGet**](ScopesApi.md#VdnScopesScopeIdGet) | **Get** /api/2.0/vdn/scopes/{scopeId} | vdnScopeProperties
[**VdnScopesScopeIdPost**](ScopesApi.md#VdnScopesScopeIdPost) | **Post** /api/2.0/vdn/scopes/{scopeId} | vdnScopeChange

# **VdnScopesGet**
> VdnScopesGet(ctx, )
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

# **VdnScopesPost**
> VdnScopesPost(ctx, optional)
vdnScopeCreate

Create a transport zone.  Request body parameters:    * **name** - Required. The name of the transport zone.   * **description** - Optional. Description of the transport zone.   * **objectId** - Required. The cluster object ID from vSphere. One or more are     required.   * **controlPlaneMode** - Optional. The control plane mode. It can be     one of the following:       * *UNICAST_MODE*       * *HYBRID_MODE*       * *MULTICAST_MODE*   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScopesApiVdnScopesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScopesApiVdnScopesPostOpts struct
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

# **VdnScopesScopeIdAttributesPut**
> VdnScopesScopeIdAttributesPut(ctx, scopeId, optional)
vdnScopeAttribUpdateUpdate

Update the attributes of a transport zone.  For example, you can update the name, description, or control plane mode. You must include the cluster object IDs for the transport zone in the request body.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***ScopesApiVdnScopesScopeIdAttributesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScopesApiVdnScopesScopeIdAttributesPutOpts struct
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

# **VdnScopesScopeIdConnCheckMulticastPost**
> VdnScopesScopeIdConnCheckMulticastPost(ctx, scopeId, optional)
vdnScopeConnCheckExecute

Test multicast group connectivity.  Test multicast group connectivity between two hosts connected to the specified transport zone.  Parameter **packetSizeMode** has one of the following values: * *0* - VXLAN standard packet size * *1* - minimum packet size * *2* - customized packet size. If you set **packetSizeMode** to *2*, you must specify the size using the **packetSize** parameter.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***ScopesApiVdnScopesScopeIdConnCheckMulticastPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScopesApiVdnScopesScopeIdConnCheckMulticastPostOpts struct
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

# **VdnScopesScopeIdDelete**
> VdnScopesScopeIdDelete(ctx, scopeId)
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

# **VdnScopesScopeIdGet**
> VdnScopesScopeIdGet(ctx, scopeId)
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

# **VdnScopesScopeIdPost**
> VdnScopesScopeIdPost(ctx, scopeId, optional)
vdnScopeChange

Update the specified transport zone.  You can add a cluster to or delete a cluster from a transport zone.  You can also repair missing portgroups. For every logical switch created, NSX creates a corresponding portgroup in vCenter. If the portgroup is lost for any reason, the logical switch will stop functioning. The repair action recreates any missing portgroups.   Parameters:  scopeId: A valid transport zone ID (vdnScope objectId)  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***ScopesApiVdnScopesScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScopesApiVdnScopesScopeIdPostOpts struct
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

