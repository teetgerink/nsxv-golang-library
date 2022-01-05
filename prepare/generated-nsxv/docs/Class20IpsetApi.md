# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesIpsetIpsetIdDelete**](Class20IpsetApi.md#Api20ServicesIpsetIpsetIdDelete) | **Delete** /api/2.0/services/ipset/{ipsetId} | ipsetDelete
[**Api20ServicesIpsetIpsetIdGet**](Class20IpsetApi.md#Api20ServicesIpsetIpsetIdGet) | **Get** /api/2.0/services/ipset/{ipsetId} | ipsetRead
[**Api20ServicesIpsetIpsetIdPut**](Class20IpsetApi.md#Api20ServicesIpsetIpsetIdPut) | **Put** /api/2.0/services/ipset/{ipsetId} | ipsetUpdate
[**Api20ServicesIpsetScopeMorefPost**](Class20IpsetApi.md#Api20ServicesIpsetScopeMorefPost) | **Post** /api/2.0/services/ipset/{scopeMoref} | ipsetCreateCreate
[**Api20ServicesIpsetScopeScopeMorefGet**](Class20IpsetApi.md#Api20ServicesIpsetScopeScopeMorefGet) | **Get** /api/2.0/services/ipset/scope/{scopeMoref} | ipsetListList

# **Api20ServicesIpsetIpsetIdDelete**
> Api20ServicesIpsetIpsetIdDelete(ctx, ipsetId, optional)
ipsetDelete

Delete an IP set.  Parameters:  ipsetId: The IP set to be queried or changed.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsetId** | **string**|  | 
 **optional** | ***Class20IpsetApiApi20ServicesIpsetIpsetIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20IpsetApiApi20ServicesIpsetIpsetIdDeleteOpts struct
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

# **Api20ServicesIpsetIpsetIdGet**
> Api20ServicesIpsetIpsetIdGet(ctx, ipsetId)
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

# **Api20ServicesIpsetIpsetIdPut**
> Api20ServicesIpsetIpsetIdPut(ctx, ipsetId, optional)
ipsetUpdate

Modify an existing IP set.  Parameters:  ipsetId: The IP set to be queried or changed.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsetId** | **string**|  | 
 **optional** | ***Class20IpsetApiApi20ServicesIpsetIpsetIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20IpsetApiApi20ServicesIpsetIpsetIdPutOpts struct
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

# **Api20ServicesIpsetScopeMorefPost**
> Api20ServicesIpsetScopeMorefPost(ctx, scopeMoref, optional)
ipsetCreateCreate

Create a new IP set.  Parameters:  scopeMoref: For scopeMoref use \"globalroot-0\" for non-universal IP sets and use \"universalroot-0\" for universal IP sets.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeMoref** | **string**|  | 
 **optional** | ***Class20IpsetApiApi20ServicesIpsetScopeMorefPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20IpsetApiApi20ServicesIpsetScopeMorefPostOpts struct
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

# **Api20ServicesIpsetScopeScopeMorefGet**
> Api20ServicesIpsetScopeScopeMorefGet(ctx, scopeMoref)
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

