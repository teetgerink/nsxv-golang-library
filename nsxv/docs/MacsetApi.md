# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesMacsetMacsetIdDelete**](MacsetApi.md#ServicesMacsetMacsetIdDelete) | **Delete** /api/2.0/services/macset/{macsetId} | macsetDelete
[**ServicesMacsetMacsetIdGet**](MacsetApi.md#ServicesMacsetMacsetIdGet) | **Get** /api/2.0/services/macset/{macsetId} | macsetRead
[**ServicesMacsetMacsetIdPut**](MacsetApi.md#ServicesMacsetMacsetIdPut) | **Put** /api/2.0/services/macset/{macsetId} | macsetUpdate
[**ServicesMacsetScopeScopeIdGet**](MacsetApi.md#ServicesMacsetScopeScopeIdGet) | **Get** /api/2.0/services/macset/scope/{scopeId} | macsetScopesRead
[**ServicesMacsetScopeScopeIdPost**](MacsetApi.md#ServicesMacsetScopeScopeIdPost) | **Post** /api/2.0/services/macset/scope/{scopeId} | macsetScopesCreate

# **ServicesMacsetMacsetIdDelete**
> ServicesMacsetMacsetIdDelete(ctx, macsetId, optional)
macsetDelete

Delete a MAC address set.  Parameters:  macsetId: Specified MAC address set ID (can be retrieved by listing the MAC address set on a scope).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **macsetId** | **string**|  | 
 **optional** | ***MacsetApiServicesMacsetMacsetIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MacsetApiServicesMacsetMacsetIdDeleteOpts struct
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

# **ServicesMacsetMacsetIdGet**
> ServicesMacsetMacsetIdGet(ctx, macsetId)
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

# **ServicesMacsetMacsetIdPut**
> ServicesMacsetMacsetIdPut(ctx, macsetId, optional)
macsetUpdate

Modify an existing MAC address set.  Parameters:  macsetId: Specified MAC address set ID (can be retrieved by listing the MAC address set on a scope).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **macsetId** | **string**|  | 
 **optional** | ***MacsetApiServicesMacsetMacsetIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MacsetApiServicesMacsetMacsetIdPutOpts struct
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

# **ServicesMacsetScopeScopeIdGet**
> ServicesMacsetScopeScopeIdGet(ctx, scopeId)
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

# **ServicesMacsetScopeScopeIdPost**
> ServicesMacsetScopeScopeIdPost(ctx, scopeId, optional)
macsetScopesCreate

Create a MAC address set on the specified scope.  Parameters:  scopeId: Can be \"globalroot-0\", \"universalroot-0\" or datacenterId in upgrade use cases. Universal MAC address sets are read-only from secondary managers.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***MacsetApiServicesMacsetScopeScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MacsetApiServicesMacsetScopeScopeIdPostOpts struct
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

