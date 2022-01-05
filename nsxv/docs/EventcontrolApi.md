# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventcontrolConfigVmVmIDGet**](EventcontrolApi.md#EventcontrolConfigVmVmIDGet) | **Get** /api/1.0/eventcontrol/config/vm/{vmID} | dataCollectionVMInfoRead
[**EventcontrolEventcontrolRootRequestPost**](EventcontrolApi.md#EventcontrolEventcontrolRootRequestPost) | **Post** /api/1.0/eventcontrol/eventcontrol-root/request | dataCollectionKillSwitchToggle
[**EventcontrolVmVmIDRequestPost**](EventcontrolApi.md#EventcontrolVmVmIDRequestPost) | **Post** /api/1.0/eventcontrol/vm/{vmID}/request | dataCollectionVMCreate

# **EventcontrolConfigVmVmIDGet**
> EventcontrolConfigVmVmIDGet(ctx, vmID)
dataCollectionVMInfoRead

Retrieve per VM configuration for data collection.   Parameters:  vmID: MOID of the guest vm  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventcontrolEventcontrolRootRequestPost**
> EventcontrolEventcontrolRootRequestPost(ctx, optional)
dataCollectionKillSwitchToggle

Turn data collection on or off at the global level.  In case of an emergency such as a network overload, you can turn off data collection at a global level (kill switch). This overrides all other data collection settings.  Set **value** to *enabled* or *disabled*.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventcontrolApiEventcontrolEventcontrolRootRequestPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventcontrolApiEventcontrolEventcontrolRootRequestPostOpts struct
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

# **EventcontrolVmVmIDRequestPost**
> EventcontrolVmVmIDRequestPost(ctx, vmID, optional)
dataCollectionVMCreate

Enable or disable data collection on a virtual machine  Set **value** to *enabled* or *disabled*.   Parameters:  vmID: MOID of the guest vm  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmID** | **string**|  | 
 **optional** | ***EventcontrolApiEventcontrolVmVmIDRequestPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventcontrolApiEventcontrolVmVmIDRequestPostOpts struct
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

