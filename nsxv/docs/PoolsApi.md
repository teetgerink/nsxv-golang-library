# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesIpamPoolsPoolIdDelete**](PoolsApi.md#Api20ServicesIpamPoolsPoolIdDelete) | **Delete** /api/2.0/services/ipam/pools/{poolId} | ipPoolDelete
[**Api20ServicesIpamPoolsPoolIdGet**](PoolsApi.md#Api20ServicesIpamPoolsPoolIdGet) | **Get** /api/2.0/services/ipam/pools/{poolId} | ipPoolRead
[**Api20ServicesIpamPoolsPoolIdIpaddressesGet**](PoolsApi.md#Api20ServicesIpamPoolsPoolIdIpaddressesGet) | **Get** /api/2.0/services/ipam/pools/{poolId}/ipaddresses | ipPoolIpsRead
[**Api20ServicesIpamPoolsPoolIdIpaddressesIpAddressDelete**](PoolsApi.md#Api20ServicesIpamPoolsPoolIdIpaddressesIpAddressDelete) | **Delete** /api/2.0/services/ipam/pools/{poolId}/ipaddresses/{ipAddress} | ipAddressReleaseDelete
[**Api20ServicesIpamPoolsPoolIdIpaddressesPost**](PoolsApi.md#Api20ServicesIpamPoolsPoolIdIpaddressesPost) | **Post** /api/2.0/services/ipam/pools/{poolId}/ipaddresses | ipPoolAllocateIp
[**Api20ServicesIpamPoolsPoolIdPut**](PoolsApi.md#Api20ServicesIpamPoolsPoolIdPut) | **Put** /api/2.0/services/ipam/pools/{poolId} | ipPoolUpdate
[**Api20ServicesIpamPoolsScopeScopeIdGet**](PoolsApi.md#Api20ServicesIpamPoolsScopeScopeIdGet) | **Get** /api/2.0/services/ipam/pools/scope/{scopeId} | ipPoolsList
[**Api20ServicesIpamPoolsScopeScopeIdPost**](PoolsApi.md#Api20ServicesIpamPoolsScopeScopeIdPost) | **Post** /api/2.0/services/ipam/pools/scope/{scopeId} | ipPoolCreate

# **Api20ServicesIpamPoolsPoolIdDelete**
> Api20ServicesIpamPoolsPoolIdDelete(ctx, poolId)
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

# **Api20ServicesIpamPoolsPoolIdGet**
> Api20ServicesIpamPoolsPoolIdGet(ctx, poolId)
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

# **Api20ServicesIpamPoolsPoolIdIpaddressesGet**
> Api20ServicesIpamPoolsPoolIdIpaddressesGet(ctx, poolId)
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

# **Api20ServicesIpamPoolsPoolIdIpaddressesIpAddressDelete**
> Api20ServicesIpamPoolsPoolIdIpaddressesIpAddressDelete(ctx, poolId, ipAddress)
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

# **Api20ServicesIpamPoolsPoolIdIpaddressesPost**
> Api20ServicesIpamPoolsPoolIdIpaddressesPost(ctx, poolId, optional)
ipPoolAllocateIp

Allocate an IP Address from the pool. Use *ALLOCATE* in the **allocationMode** field in the body to allocate the next available IP. To allocate a specific one use *RESERVE* and pass the IP to reserve in the **ipAddress** fields in the body.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 
 **optional** | ***PoolsApiApi20ServicesIpamPoolsPoolIdIpaddressesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoolsApiApi20ServicesIpamPoolsPoolIdIpaddressesPostOpts struct
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

# **Api20ServicesIpamPoolsPoolIdPut**
> Api20ServicesIpamPoolsPoolIdPut(ctx, poolId, optional)
ipPoolUpdate

To modify an IP pool, query the IP pool first. Then modify the output and send it back as the request body.   Parameters:  poolId: Specifiy the pool ID as *poolId* in the URI.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 
 **optional** | ***PoolsApiApi20ServicesIpamPoolsPoolIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoolsApiApi20ServicesIpamPoolsPoolIdPutOpts struct
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

# **Api20ServicesIpamPoolsScopeScopeIdGet**
> Api20ServicesIpamPoolsScopeScopeIdGet(ctx, scopeId)
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

# **Api20ServicesIpamPoolsScopeScopeIdPost**
> Api20ServicesIpamPoolsScopeScopeIdPost(ctx, scopeId, optional)
ipPoolCreate

Create a pool of IP addresses. For *scopeId* use globalroot-0 or the *datacenterId* in upgrade use cases.   Parameters:  scopeId: For *scopeID* use globalroot-0 or *datacenterId* in upgrade use cases.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***PoolsApiApi20ServicesIpamPoolsScopeScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoolsApiApi20ServicesIpamPoolsScopeScopeIdPostOpts struct
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

