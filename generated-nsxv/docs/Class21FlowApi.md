# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api21AppFlowConfigGet**](Class21FlowApi.md#Api21AppFlowConfigGet) | **Get** /api/2.1/app/flow/config | flowConfigRead
[**Api21AppFlowConfigPut**](Class21FlowApi.md#Api21AppFlowConfigPut) | **Put** /api/2.1/app/flow/config | flowsConfigCreate
[**Api21AppFlowContextIdDelete**](Class21FlowApi.md#Api21AppFlowContextIdDelete) | **Delete** /api/2.1/app/flow/{contextId} | flowContextDelete
[**Api21AppFlowFlowstatsGet**](Class21FlowApi.md#Api21AppFlowFlowstatsGet) | **Get** /api/2.1/app/flow/flowstats | flowStatsRead
[**Api21AppFlowFlowstatsInfoGet**](Class21FlowApi.md#Api21AppFlowFlowstatsInfoGet) | **Get** /api/2.1/app/flow/flowstats/info | flowStatsInfoRead

# **Api21AppFlowConfigGet**
> Api21AppFlowConfigGet(ctx, )
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

# **Api21AppFlowConfigPut**
> Api21AppFlowConfigPut(ctx, optional)
flowsConfigCreate

Update flow monitoring configuration.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class21FlowApiApi21AppFlowConfigPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class21FlowApiApi21AppFlowConfigPutOpts struct
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

# **Api21AppFlowContextIdDelete**
> Api21AppFlowContextIdDelete(ctx, contextId)
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

# **Api21AppFlowFlowstatsGet**
> Api21AppFlowFlowstatsGet(ctx, optional)
flowStatsRead

Retrieve flow statistics for a datacenter, port group, VM, or vNIC.  Response values for flow statistics: * **blocked** - indicates whether traffic is blocked:   * 0 - flow allowed   * 1 - flow blocked   * 2 - flow blocked by SpoofGuard * **protocol** - protocol in flow:   * 0 - TCP   * 1 - UDP   * 2 - ICMP * **direction** - direction of flow:   * 0 - to virtual machine   * 1 - from virtual machine * **controlDirection** - control direction for dynamic TCP traffic:   * 0 - source -> destination   * 1 - destination -> source   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class21FlowApiApi21AppFlowFlowstatsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class21FlowApiApi21AppFlowFlowstatsGetOpts struct
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

# **Api21AppFlowFlowstatsInfoGet**
> Api21AppFlowFlowstatsInfoGet(ctx, )
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

