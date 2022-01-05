# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesMacsetMacsetIdDelete**](MacsetApi.md#Api20ServicesMacsetMacsetIdDelete) | **Delete** /api/2.0/services/macset/{macsetId} | macsetDelete
[**Api20ServicesMacsetMacsetIdGet**](MacsetApi.md#Api20ServicesMacsetMacsetIdGet) | **Get** /api/2.0/services/macset/{macsetId} | macsetRead
[**Api20ServicesMacsetMacsetIdPut**](MacsetApi.md#Api20ServicesMacsetMacsetIdPut) | **Put** /api/2.0/services/macset/{macsetId} | macsetUpdate
[**Api20ServicesMacsetScopeScopeIdGet**](MacsetApi.md#Api20ServicesMacsetScopeScopeIdGet) | **Get** /api/2.0/services/macset/scope/{scopeId} | macsetScopesRead
[**Api20ServicesMacsetScopeScopeIdPost**](MacsetApi.md#Api20ServicesMacsetScopeScopeIdPost) | **Post** /api/2.0/services/macset/scope/{scopeId} | macsetScopesCreate

# **Api20ServicesMacsetMacsetIdDelete**
> Api20ServicesMacsetMacsetIdDelete(ctx, macsetId, optional)
macsetDelete

Delete a MAC address set.  Parameters:  macsetId: Specified MAC address set ID (can be retrieved by listing the MAC address set on a scope).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **macsetId** | **string**|  | 
 **optional** | ***MacsetApiApi20ServicesMacsetMacsetIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MacsetApiApi20ServicesMacsetMacsetIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesMacsetMacsetIdGet**
> Api20ServicesMacsetMacsetIdGet(ctx, macsetId)
macsetRead

Retrieve details about a MAC address set.  Parameters:  macsetId: Specified MAC address set ID (can be retrieved by listing the MAC address set on a scope).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **macsetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesMacsetMacsetIdPut**
> Api20ServicesMacsetMacsetIdPut(ctx, macsetId, optional)
macsetUpdate

Modify an existing MAC address set.  Parameters:  macsetId: Specified MAC address set ID (can be retrieved by listing the MAC address set on a scope).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **macsetId** | **string**|  | 
 **optional** | ***MacsetApiApi20ServicesMacsetMacsetIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MacsetApiApi20ServicesMacsetMacsetIdPutOpts struct
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

# **Api20ServicesMacsetScopeScopeIdGet**
> Api20ServicesMacsetScopeScopeIdGet(ctx, scopeId)
macsetScopesRead

List MAC address sets on the specified scope.  Parameters:  scopeId: Can be \"globalroot-0\", \"universalroot-0\" or datacenterId in upgrade use cases. Universal MAC address sets are read-only from secondary managers.   

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

# **Api20ServicesMacsetScopeScopeIdPost**
> Api20ServicesMacsetScopeScopeIdPost(ctx, scopeId, optional)
macsetScopesCreate

Create a MAC address set on the specified scope.  Parameters:  scopeId: Can be \"globalroot-0\", \"universalroot-0\" or datacenterId in upgrade use cases. Universal MAC address sets are read-only from secondary managers.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***MacsetApiApi20ServicesMacsetScopeScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MacsetApiApi20ServicesMacsetScopeScopeIdPostOpts struct
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

