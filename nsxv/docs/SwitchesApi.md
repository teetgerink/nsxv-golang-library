# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VdnSwitchesDatacenterDatacenterIDGet**](SwitchesApi.md#VdnSwitchesDatacenterDatacenterIDGet) | **Get** /api/2.0/vdn/switches/datacenter/{datacenterID} | vdsListInDc
[**VdnSwitchesGet**](SwitchesApi.md#VdnSwitchesGet) | **Get** /api/2.0/vdn/switches | vdsList
[**VdnSwitchesPost**](SwitchesApi.md#VdnSwitchesPost) | **Post** /api/2.0/vdn/switches | vdsPrepare
[**VdnSwitchesVdsIdDelete**](SwitchesApi.md#VdnSwitchesVdsIdDelete) | **Delete** /api/2.0/vdn/switches/{vdsId} | vdsDelete
[**VdnSwitchesVdsIdGet**](SwitchesApi.md#VdnSwitchesVdsIdGet) | **Get** /api/2.0/vdn/switches/{vdsId} | vdsShow

# **VdnSwitchesDatacenterDatacenterIDGet**
> VdnSwitchesDatacenterDatacenterIDGet(ctx, datacenterID)
vdsListInDc

Retrieve information about all vSphere Distributed Switches in the specified datacenter.   Parameters:  datacenterID: A valid datacenter ID (e.g. datacenter-21)   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datacenterID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnSwitchesGet**
> VdnSwitchesGet(ctx, )
vdsList

Retrieve information about all vSphere Distributed Switches.   Parameters:  

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

# **VdnSwitchesPost**
> VdnSwitchesPost(ctx, optional)
vdsPrepare

Prepare a vSphere Distributed Switch.  The MTU is the maximum amount of data that can be transmitted in one packet before it is divided into smaller packets. VXLAN frames are slightly larger in size because of the traffic encapsulation, so the MTU required is higher than the standard MTU. You must set the MTU for each switch to 1602 or higher.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwitchesApiVdnSwitchesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwitchesApiVdnSwitchesPostOpts struct
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

# **VdnSwitchesVdsIdDelete**
> VdnSwitchesVdsIdDelete(ctx, vdsId)
vdsDelete

Delete the specified vSphere Distributed Switch.   Parameters:  vdsId: A valid vSphere Distributed Switch ID (e.g. dvs-35)   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdsId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnSwitchesVdsIdGet**
> VdnSwitchesVdsIdGet(ctx, vdsId)
vdsShow

Retrieve information about the specified vSphere Distributed Switch.   Parameters:  vdsId: A valid vSphere Distributed Switch ID (e.g. dvs-35)   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdsId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

