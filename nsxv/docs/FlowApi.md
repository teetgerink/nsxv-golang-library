# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppFlowConfigGet**](FlowApi.md#AppFlowConfigGet) | **Get** /api/2.1/app/flow/config | flowConfigRead
[**AppFlowConfigPut**](FlowApi.md#AppFlowConfigPut) | **Put** /api/2.1/app/flow/config | flowsConfigCreate
[**AppFlowContextIdDelete**](FlowApi.md#AppFlowContextIdDelete) | **Delete** /api/2.1/app/flow/{contextId} | flowContextDelete
[**AppFlowFlowstatsGet**](FlowApi.md#AppFlowFlowstatsGet) | **Get** /api/2.1/app/flow/flowstats | flowStatsRead
[**AppFlowFlowstatsInfoGet**](FlowApi.md#AppFlowFlowstatsInfoGet) | **Get** /api/2.1/app/flow/flowstats/info | flowStatsInfoRead

# **AppFlowConfigGet**
> AppFlowConfigGet(ctx, )
flowConfigRead

Retrieve flow monitoring configuration.  Parameters:  

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

# **AppFlowConfigPut**
> AppFlowConfigPut(ctx, optional)
flowsConfigCreate

Update flow monitoring configuration.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiAppFlowConfigPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiAppFlowConfigPutOpts struct
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

# **AppFlowContextIdDelete**
> AppFlowContextIdDelete(ctx, contextId)
flowContextDelete

Delete flow records for the specified context.  Parameters:  contextId: Context ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppFlowFlowstatsGet**
> AppFlowFlowstatsGet(ctx, optional)
flowStatsRead

Retrieve flow statistics for a datacenter, port group, VM, or vNIC.  Response values for flow statistics: * **blocked** - indicates whether traffic is blocked:   * 0 - flow allowed   * 1 - flow blocked   * 2 - flow blocked by SpoofGuard * **protocol** - protocol in flow:   * 0 - TCP   * 1 - UDP   * 2 - ICMP * **direction** - direction of flow:   * 0 - to virtual machine   * 1 - from virtual machine * **controlDirection** - control direction for dynamic TCP traffic:   * 0 - source -> destination   * 1 - destination -> source   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiAppFlowFlowstatsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiAppFlowFlowstatsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextId** | **optional.String**|  | 
 **flowType** | **optional.String**|  | 
 **startTime** | **optional.String**|  | 
 **endTime** | **optional.String**|  | 
 **startIndex** | **optional.String**|  | 
 **pageSize** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppFlowFlowstatsInfoGet**
> AppFlowFlowstatsInfoGet(ctx, )
flowStatsInfoRead

Retrieve flow statistics meta-data.  This method retrieves the following information for each flow type: * minimum start time * maximum end time * total flow count   Parameters:  

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

