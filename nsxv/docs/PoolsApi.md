# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesIpamPoolsPoolIdDelete**](PoolsApi.md#ServicesIpamPoolsPoolIdDelete) | **Delete** /api/2.0/services/ipam/pools/{poolId} | ipPoolDelete
[**ServicesIpamPoolsPoolIdGet**](PoolsApi.md#ServicesIpamPoolsPoolIdGet) | **Get** /api/2.0/services/ipam/pools/{poolId} | ipPoolRead
[**ServicesIpamPoolsPoolIdIpaddressesGet**](PoolsApi.md#ServicesIpamPoolsPoolIdIpaddressesGet) | **Get** /api/2.0/services/ipam/pools/{poolId}/ipaddresses | ipPoolIpsRead
[**ServicesIpamPoolsPoolIdIpaddressesIpAddressDelete**](PoolsApi.md#ServicesIpamPoolsPoolIdIpaddressesIpAddressDelete) | **Delete** /api/2.0/services/ipam/pools/{poolId}/ipaddresses/{ipAddress} | ipAddressReleaseDelete
[**ServicesIpamPoolsPoolIdIpaddressesPost**](PoolsApi.md#ServicesIpamPoolsPoolIdIpaddressesPost) | **Post** /api/2.0/services/ipam/pools/{poolId}/ipaddresses | ipPoolAllocateIp
[**ServicesIpamPoolsPoolIdPut**](PoolsApi.md#ServicesIpamPoolsPoolIdPut) | **Put** /api/2.0/services/ipam/pools/{poolId} | ipPoolUpdate
[**ServicesIpamPoolsScopeScopeIdGet**](PoolsApi.md#ServicesIpamPoolsScopeScopeIdGet) | **Get** /api/2.0/services/ipam/pools/scope/{scopeId} | ipPoolsList
[**ServicesIpamPoolsScopeScopeIdPost**](PoolsApi.md#ServicesIpamPoolsScopeScopeIdPost) | **Post** /api/2.0/services/ipam/pools/scope/{scopeId} | ipPoolCreate

# **ServicesIpamPoolsPoolIdDelete**
> ServicesIpamPoolsPoolIdDelete(ctx, poolId)
ipPoolDelete

Delete an IP pool.  Parameters:  poolId: Specifiy the pool ID as *poolId* in the URI.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesIpamPoolsPoolIdGet**
> ServicesIpamPoolsPoolIdGet(ctx, poolId)
ipPoolRead

Retrieve details about a specific IP pool.  Parameters:  poolId: Specifiy the pool ID as *poolId* in the URI.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesIpamPoolsPoolIdIpaddressesGet**
> ServicesIpamPoolsPoolIdIpaddressesGet(ctx, poolId)
ipPoolIpsRead

Retrieves all allocated IP addresses from the specified pool.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesIpamPoolsPoolIdIpaddressesIpAddressDelete**
> ServicesIpamPoolsPoolIdIpaddressesIpAddressDelete(ctx, poolId, ipAddress)
ipAddressReleaseDelete

Release an IP address allocation in the pool.  Parameters:  ipAddress: The IP address to release, e.g. '192.168.10.10'   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 
  **ipAddress** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesIpamPoolsPoolIdIpaddressesPost**
> ServicesIpamPoolsPoolIdIpaddressesPost(ctx, poolId, optional)
ipPoolAllocateIp

Allocate an IP Address from the pool. Use *ALLOCATE* in the **allocationMode** field in the body to allocate the next available IP. To allocate a specific one use *RESERVE* and pass the IP to reserve in the **ipAddress** fields in the body.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 
 **optional** | ***PoolsApiServicesIpamPoolsPoolIdIpaddressesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoolsApiServicesIpamPoolsPoolIdIpaddressesPostOpts struct
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

# **ServicesIpamPoolsPoolIdPut**
> ServicesIpamPoolsPoolIdPut(ctx, poolId, optional)
ipPoolUpdate

To modify an IP pool, query the IP pool first. Then modify the output and send it back as the request body.   Parameters:  poolId: Specifiy the pool ID as *poolId* in the URI.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 
 **optional** | ***PoolsApiServicesIpamPoolsPoolIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoolsApiServicesIpamPoolsPoolIdPutOpts struct
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

# **ServicesIpamPoolsScopeScopeIdGet**
> ServicesIpamPoolsScopeScopeIdGet(ctx, scopeId)
ipPoolsList

Retrieves all IP pools on the specified scope where the *scopeID* is the reference to the desired scope. An example of the *scopeID* is globalroot-0.   Parameters:  scopeId: For *scopeID* use globalroot-0 or *datacenterId* in upgrade use cases.   

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

# **ServicesIpamPoolsScopeScopeIdPost**
> ServicesIpamPoolsScopeScopeIdPost(ctx, scopeId, optional)
ipPoolCreate

Create a pool of IP addresses. For *scopeId* use globalroot-0 or the *datacenterId* in upgrade use cases.   Parameters:  scopeId: For *scopeID* use globalroot-0 or *datacenterId* in upgrade use cases.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***PoolsApiServicesIpamPoolsScopeScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoolsApiServicesIpamPoolsScopeScopeIdPostOpts struct
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

