# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgePublishTuningConfigurationGet**](EdgePublishApi.md#EdgePublishTuningConfigurationGet) | **Get** /api/4.0/edgePublish/tuningConfiguration | nsxEdgePublishTuningRead
[**EdgePublishTuningConfigurationPut**](EdgePublishApi.md#EdgePublishTuningConfigurationPut) | **Put** /api/4.0/edgePublish/tuningConfiguration | nsxEdgePublishTuningUpdate

# **EdgePublishTuningConfigurationGet**
> EdgePublishTuningConfigurationGet(ctx, )
nsxEdgePublishTuningRead

Retrieve the NSX Edge tuning configuration.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

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

# **EdgePublishTuningConfigurationPut**
> EdgePublishTuningConfigurationPut(ctx, optional)
nsxEdgePublishTuningUpdate

Update the NSX Edge tuning configuration.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EdgePublishApiEdgePublishTuningConfigurationPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgePublishApiEdgePublishTuningConfigurationPutOpts struct
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

