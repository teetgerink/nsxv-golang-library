# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesVcconfigGet**](Class20VcconfigApi.md#Api20ServicesVcconfigGet) | **Get** /api/2.0/services/vcconfig | vCenterConfigRead
[**Api20ServicesVcconfigPut**](Class20VcconfigApi.md#Api20ServicesVcconfigPut) | **Put** /api/2.0/services/vcconfig | vCenterConfigUpdate
[**Api20ServicesVcconfigStatusGet**](Class20VcconfigApi.md#Api20ServicesVcconfigStatusGet) | **Get** /api/2.0/services/vcconfig/status | vCenterStatusRead

# **Api20ServicesVcconfigGet**
> Api20ServicesVcconfigGet(ctx, )
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

# **Api20ServicesVcconfigPut**
> Api20ServicesVcconfigPut(ctx, optional)
vCenterConfigUpdate

Synchronize NSX Manager with vCenter server.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20VcconfigApiApi20ServicesVcconfigPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20VcconfigApiApi20ServicesVcconfigPutOpts struct
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

# **Api20ServicesVcconfigStatusGet**
> Api20ServicesVcconfigStatusGet(ctx, )
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

