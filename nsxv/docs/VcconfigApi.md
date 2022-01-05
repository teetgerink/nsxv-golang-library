# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesVcconfigGet**](VcconfigApi.md#ServicesVcconfigGet) | **Get** /api/2.0/services/vcconfig | vCenterConfigRead
[**ServicesVcconfigPut**](VcconfigApi.md#ServicesVcconfigPut) | **Put** /api/2.0/services/vcconfig | vCenterConfigUpdate
[**ServicesVcconfigStatusGet**](VcconfigApi.md#ServicesVcconfigStatusGet) | **Get** /api/2.0/services/vcconfig/status | vCenterStatusRead

# **ServicesVcconfigGet**
> ServicesVcconfigGet(ctx, )
vCenterConfigRead

Get vCenter Server configuration details on NSX Manager.  Parameters:  

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

# **ServicesVcconfigPut**
> ServicesVcconfigPut(ctx, optional)
vCenterConfigUpdate

Synchronize NSX Manager with vCenter server.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VcconfigApiServicesVcconfigPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VcconfigApiServicesVcconfigPutOpts struct
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

# **ServicesVcconfigStatusGet**
> ServicesVcconfigStatusGet(ctx, )
vCenterStatusRead

Get default vCenter Server connection status.  Parameters:  

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

