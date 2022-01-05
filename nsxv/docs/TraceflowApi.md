# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiVdnTraceflowPost**](TraceflowApi.md#ApiVdnTraceflowPost) | **Post** /api/api/2.0/vdn/traceflow | traceflowCreate
[**ApiVdnTraceflowTraceflowIdGet**](TraceflowApi.md#ApiVdnTraceflowTraceflowIdGet) | **Get** /api/api/2.0/vdn/traceflow/{traceflowId} | traceflowRead
[**ApiVdnTraceflowTraceflowIdObservationsGet**](TraceflowApi.md#ApiVdnTraceflowTraceflowIdObservationsGet) | **Get** /api/api/2.0/vdn/traceflow/{traceflowId}/observations | traceflowObservationsList

# **ApiVdnTraceflowPost**
> ApiVdnTraceflowPost(ctx, optional)
traceflowCreate

Create a traceflow.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TraceflowApiApiVdnTraceflowPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TraceflowApiApiVdnTraceflowPostOpts struct
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

# **ApiVdnTraceflowTraceflowIdGet**
> ApiVdnTraceflowTraceflowIdGet(ctx, traceflowId)
traceflowRead

Query a specific traceflow by *tracflowId* which is the value returned after executing the create Traceflow API call.   Parameters:  traceflowId: Traceflow ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **traceflowId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiVdnTraceflowTraceflowIdObservationsGet**
> ApiVdnTraceflowTraceflowIdObservationsGet(ctx, traceflowId)
traceflowObservationsList

Retrieve traceflow observations.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **traceflowId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

