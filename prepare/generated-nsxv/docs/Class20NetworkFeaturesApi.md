# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20XvsNetworksIDFeaturesGet**](Class20NetworkFeaturesApi.md#Api20XvsNetworksIDFeaturesGet) | **Get** /api/2.0/xvs/networks/{ID}/features | arpMACRead
[**Api20XvsNetworksIDFeaturesPut**](Class20NetworkFeaturesApi.md#Api20XvsNetworksIDFeaturesPut) | **Put** /api/2.0/xvs/networks/{ID}/features | arpMACUpdate

# **Api20XvsNetworksIDFeaturesGet**
> Api20XvsNetworksIDFeaturesGet(ctx, iD)
arpMACRead

Retrieve IP discovery and MAC learning information.  Parameters:  ID: dvPortGroup MOID or logical switch (virtual wire) ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20XvsNetworksIDFeaturesPut**
> Api20XvsNetworksIDFeaturesPut(ctx, iD, optional)
arpMACUpdate

Enable or disable IP discovery and MAC learning.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method updated. IP discovery can be disabled on secondary NSX Managers.   Parameters:  ID: dvPortGroup MOID or logical switch (virtual wire) ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 
 **optional** | ***Class20NetworkFeaturesApiApi20XvsNetworksIDFeaturesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20NetworkFeaturesApiApi20XvsNetworksIDFeaturesPutOpts struct
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

