# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api40EdgePublishTuningConfigurationGet**](Class40EdgePublishApi.md#Api40EdgePublishTuningConfigurationGet) | **Get** /api/4.0/edgePublish/tuningConfiguration | nsxEdgePublishTuningRead
[**Api40EdgePublishTuningConfigurationPut**](Class40EdgePublishApi.md#Api40EdgePublishTuningConfigurationPut) | **Put** /api/4.0/edgePublish/tuningConfiguration | nsxEdgePublishTuningUpdate

# **Api40EdgePublishTuningConfigurationGet**
> Api40EdgePublishTuningConfigurationGet(ctx, )
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

# **Api40EdgePublishTuningConfigurationPut**
> Api40EdgePublishTuningConfigurationPut(ctx, optional)
nsxEdgePublishTuningUpdate

Update the NSX Edge tuning configuration.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class40EdgePublishApiApi40EdgePublishTuningConfigurationPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class40EdgePublishApiApi40EdgePublishTuningConfigurationPutOpts struct
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

