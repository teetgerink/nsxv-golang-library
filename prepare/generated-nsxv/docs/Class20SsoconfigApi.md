# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesSsoconfigDelete**](Class20SsoconfigApi.md#Api20ServicesSsoconfigDelete) | **Delete** /api/2.0/services/ssoconfig | ssoConfigDelete
[**Api20ServicesSsoconfigGet**](Class20SsoconfigApi.md#Api20ServicesSsoconfigGet) | **Get** /api/2.0/services/ssoconfig | ssoConfigRead
[**Api20ServicesSsoconfigPost**](Class20SsoconfigApi.md#Api20ServicesSsoconfigPost) | **Post** /api/2.0/services/ssoconfig | ssoConfigSet
[**Api20ServicesSsoconfigStatusGet**](Class20SsoconfigApi.md#Api20ServicesSsoconfigStatusGet) | **Get** /api/2.0/services/ssoconfig/status | ssoStatusRead

# **Api20ServicesSsoconfigDelete**
> Api20ServicesSsoconfigDelete(ctx, )
ssoConfigDelete

Deletes the NSX Manager SSO Configuration.  Parameters:  

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

# **Api20ServicesSsoconfigGet**
> Api20ServicesSsoconfigGet(ctx, )
ssoConfigRead

Retrieve SSO Configuration.  Parameters:  

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

# **Api20ServicesSsoconfigPost**
> Api20ServicesSsoconfigPost(ctx, optional)
ssoConfigSet

Register NSX Manager to SSO Services.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20SsoconfigApiApi20ServicesSsoconfigPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20SsoconfigApiApi20ServicesSsoconfigPostOpts struct
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

# **Api20ServicesSsoconfigStatusGet**
> Api20ServicesSsoconfigStatusGet(ctx, )
ssoStatusRead

Retrieve the SSO configuration status of NSX Manager.  Parameters:  

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

