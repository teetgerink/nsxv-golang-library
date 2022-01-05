# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesIpsetIpsetIdDelete**](IpsetApi.md#ServicesIpsetIpsetIdDelete) | **Delete** /api/2.0/services/ipset/{ipsetId} | ipsetDelete
[**ServicesIpsetIpsetIdGet**](IpsetApi.md#ServicesIpsetIpsetIdGet) | **Get** /api/2.0/services/ipset/{ipsetId} | ipsetRead
[**ServicesIpsetIpsetIdPut**](IpsetApi.md#ServicesIpsetIpsetIdPut) | **Put** /api/2.0/services/ipset/{ipsetId} | ipsetUpdate
[**ServicesIpsetScopeMorefPost**](IpsetApi.md#ServicesIpsetScopeMorefPost) | **Post** /api/2.0/services/ipset/{scopeMoref} | ipsetCreateCreate
[**ServicesIpsetScopeScopeMorefGet**](IpsetApi.md#ServicesIpsetScopeScopeMorefGet) | **Get** /api/2.0/services/ipset/scope/{scopeMoref} | ipsetListList

# **ServicesIpsetIpsetIdDelete**
> ServicesIpsetIpsetIdDelete(ctx, ipsetId, optional)
ipsetDelete

Delete an IP set.  Parameters:  ipsetId: The IP set to be queried or changed.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsetId** | **string**|  | 
 **optional** | ***IpsetApiServicesIpsetIpsetIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IpsetApiServicesIpsetIpsetIdDeleteOpts struct
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

# **ServicesIpsetIpsetIdGet**
> ServicesIpsetIpsetIdGet(ctx, ipsetId)
ipsetRead

Retrieve an individual IP set.  Parameters:  ipsetId: The IP set to be queried or changed.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesIpsetIpsetIdPut**
> ServicesIpsetIpsetIdPut(ctx, ipsetId, optional)
ipsetUpdate

Modify an existing IP set.  Parameters:  ipsetId: The IP set to be queried or changed.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsetId** | **string**|  | 
 **optional** | ***IpsetApiServicesIpsetIpsetIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IpsetApiServicesIpsetIpsetIdPutOpts struct
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

# **ServicesIpsetScopeMorefPost**
> ServicesIpsetScopeMorefPost(ctx, scopeMoref, optional)
ipsetCreateCreate

Create a new IP set.  Parameters:  scopeMoref: For scopeMoref use \"globalroot-0\" for non-universal IP sets and use \"universalroot-0\" for universal IP sets.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeMoref** | **string**|  | 
 **optional** | ***IpsetApiServicesIpsetScopeMorefPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IpsetApiServicesIpsetScopeMorefPostOpts struct
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

# **ServicesIpsetScopeScopeMorefGet**
> ServicesIpsetScopeScopeMorefGet(ctx, scopeMoref)
ipsetListList

Retrieve all configured IPSets  Parameters:  scopeMoref: For scopeMoref use \"globalroot-0\" for non-universal IP sets and use \"universalroot-0\" for universal IP sets.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeMoref** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

