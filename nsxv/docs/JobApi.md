# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesTaskserviceJobGet**](JobApi.md#ServicesTaskserviceJobGet) | **Get** /api/2.0/services/taskservice/job | taskFrameworkCriteria
[**ServicesTaskserviceJobJobIdGet**](JobApi.md#ServicesTaskserviceJobJobIdGet) | **Get** /api/2.0/services/taskservice/job/{jobId} | taskFrameworkJobsRead

# **ServicesTaskserviceJobGet**
> ServicesTaskserviceJobGet(ctx, optional)
taskFrameworkCriteria

Query job instances by criterion.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobApiServicesTaskserviceJobGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiServicesTaskserviceJobGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **optional.String**|  | 
 **pageSize** | **optional.String**|  | 
 **sortBy** | **optional.String**|  | 
 **sortOrderAscending** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesTaskserviceJobJobIdGet**
> ServicesTaskserviceJobJobIdGet(ctx, jobId)
taskFrameworkJobsRead

Retrieve all job instances for the specified job ID.  Parameters:  jobId: Specified job ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

