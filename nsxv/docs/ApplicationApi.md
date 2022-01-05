# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesApplicationApplicationIdDelete**](ApplicationApi.md#ServicesApplicationApplicationIdDelete) | **Delete** /api/2.0/services/application/{applicationId} | serviceAppDelete
[**ServicesApplicationApplicationIdGet**](ApplicationApi.md#ServicesApplicationApplicationIdGet) | **Get** /api/2.0/services/application/{applicationId} | serviceAppRead
[**ServicesApplicationApplicationIdPut**](ApplicationApi.md#ServicesApplicationApplicationIdPut) | **Put** /api/2.0/services/application/{applicationId} | serviceAppUpdate
[**ServicesApplicationScopeIdPost**](ApplicationApi.md#ServicesApplicationScopeIdPost) | **Post** /api/2.0/services/application/{scopeId} | servicesAppsScopeCreate
[**ServicesApplicationScopeScopeIdGet**](ApplicationApi.md#ServicesApplicationScopeScopeIdGet) | **Get** /api/2.0/services/application/scope/{scopeId} | servicesAppsScopeScopeRead

# **ServicesApplicationApplicationIdDelete**
> ServicesApplicationApplicationIdDelete(ctx, applicationId, optional)
serviceAppDelete

Delete the specified service.  Parameters:  applicationId: Application ID. You can get a list of application IDs from `GET /api/2.0/services/application/scope/{scopeId}`.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
 **optional** | ***ApplicationApiServicesApplicationApplicationIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiServicesApplicationApplicationIdDeleteOpts struct
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

# **ServicesApplicationApplicationIdGet**
> ServicesApplicationApplicationIdGet(ctx, applicationId)
serviceAppRead

Retrieve details about the specified service.  Parameters:  applicationId: Application ID. You can get a list of application IDs from `GET /api/2.0/services/application/scope/{scopeId}`.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesApplicationApplicationIdPut**
> ServicesApplicationApplicationIdPut(ctx, applicationId, optional)
serviceAppUpdate

Modify the name, description, applicationProtocol, or port value of a service.   Parameters:  applicationId: Application ID. You can get a list of application IDs from `GET /api/2.0/services/application/scope/{scopeId}`.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
 **optional** | ***ApplicationApiServicesApplicationApplicationIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiServicesApplicationApplicationIdPutOpts struct
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

# **ServicesApplicationScopeIdPost**
> ServicesApplicationScopeIdPost(ctx, scopeId, optional)
servicesAppsScopeCreate

Create a new service on the specified scope.   Parameters:  scopeId:   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***ApplicationApiServicesApplicationScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiServicesApplicationScopeIdPostOpts struct
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

# **ServicesApplicationScopeScopeIdGet**
> ServicesApplicationScopeScopeIdGet(ctx, scopeId)
servicesAppsScopeScopeRead

Retrieve services that have been created on the specified scope.  Parameters:  scopeId: Can be \"globalroot-0\", \"universalroot-0\" or datacenterId in upgrade use cases.   

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

