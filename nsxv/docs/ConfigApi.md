# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VdnConfigHostHostIdVxlanVtepsPost**](ConfigApi.md#VdnConfigHostHostIdVxlanVtepsPost) | **Post** /api/2.0/vdn/config/host/{hostId}/vxlan/vteps | hostVtepResolveAction
[**VdnConfigMulticastsGet**](ConfigApi.md#VdnConfigMulticastsGet) | **Get** /api/2.0/vdn/config/multicasts | vdnMulticastPoolList
[**VdnConfigMulticastsMulticastAddresssRangeIdDelete**](ConfigApi.md#VdnConfigMulticastsMulticastAddresssRangeIdDelete) | **Delete** /api/2.0/vdn/config/multicasts/{multicastAddresssRangeId} | vdnMulticastPoolDelete
[**VdnConfigMulticastsMulticastAddresssRangeIdGet**](ConfigApi.md#VdnConfigMulticastsMulticastAddresssRangeIdGet) | **Get** /api/2.0/vdn/config/multicasts/{multicastAddresssRangeId} | vdnMulticastPoolShow
[**VdnConfigMulticastsMulticastAddresssRangeIdPut**](ConfigApi.md#VdnConfigMulticastsMulticastAddresssRangeIdPut) | **Put** /api/2.0/vdn/config/multicasts/{multicastAddresssRangeId} | vdnMulticastPoolChange
[**VdnConfigMulticastsPost**](ConfigApi.md#VdnConfigMulticastsPost) | **Post** /api/2.0/vdn/config/multicasts | vdnMulticastPoolAdd
[**VdnConfigResourcesAllocatedGet**](ConfigApi.md#VdnConfigResourcesAllocatedGet) | **Get** /api/2.0/vdn/config/resources/allocated | allocatedResourcesRead
[**VdnConfigSegmentsGet**](ConfigApi.md#VdnConfigSegmentsGet) | **Get** /api/2.0/vdn/config/segments | vdnSegmentPoolList
[**VdnConfigSegmentsPost**](ConfigApi.md#VdnConfigSegmentsPost) | **Post** /api/2.0/vdn/config/segments | vdnSegmentPoolAdd
[**VdnConfigSegmentsSegmentPoolIdDelete**](ConfigApi.md#VdnConfigSegmentsSegmentPoolIdDelete) | **Delete** /api/2.0/vdn/config/segments/{segmentPoolId} | vdnSegmentPoolDelete
[**VdnConfigSegmentsSegmentPoolIdGet**](ConfigApi.md#VdnConfigSegmentsSegmentPoolIdGet) | **Get** /api/2.0/vdn/config/segments/{segmentPoolId} | vdnSegmentPoolShow
[**VdnConfigSegmentsSegmentPoolIdPut**](ConfigApi.md#VdnConfigSegmentsSegmentPoolIdPut) | **Put** /api/2.0/vdn/config/segments/{segmentPoolId} | vdnSegmentPoolChange
[**VdnConfigVxlanUdpPortGet**](ConfigApi.md#VdnConfigVxlanUdpPortGet) | **Get** /api/2.0/vdn/config/vxlan/udp/port | vdnConfigUDPShow
[**VdnConfigVxlanUdpPortPortNumberPut**](ConfigApi.md#VdnConfigVxlanUdpPortPortNumberPut) | **Put** /api/2.0/vdn/config/vxlan/udp/port/{portNumber} | vdnConfigUDPUpdateUpdate
[**VdnConfigVxlanUdpPortTaskStatusGet**](ConfigApi.md#VdnConfigVxlanUdpPortTaskStatusGet) | **Get** /api/2.0/vdn/config/vxlan/udp/port/taskStatus | vdnConfigUDPUpdateStatusRead

# **VdnConfigHostHostIdVxlanVtepsPost**
> VdnConfigHostHostIdVxlanVtepsPost(ctx, hostId, optional)
hostVtepResolveAction

Resolve missing VXLAN VMKernel adapters.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  hostId:   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**|  | 
 **optional** | ***ConfigApiVdnConfigHostHostIdVxlanVtepsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiVdnConfigHostHostIdVxlanVtepsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnConfigMulticastsGet**
> VdnConfigMulticastsGet(ctx, )
vdnMulticastPoolList

Retrieve information about all configured multicast address ranges.  Universal multicast address ranges have the property isUniversal set to *true*.   Parameters:  

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

# **VdnConfigMulticastsMulticastAddresssRangeIdDelete**
> VdnConfigMulticastsMulticastAddresssRangeIdDelete(ctx, multicastAddresssRangeId)
vdnMulticastPoolDelete

Delete the specified multicast address range.  If the multicast address range is universal you must send the API request to the primary NSX Manager.   Parameters:  multicastAddresssRangeId: A valid multicast address range ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **multicastAddresssRangeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnConfigMulticastsMulticastAddresssRangeIdGet**
> VdnConfigMulticastsMulticastAddresssRangeIdGet(ctx, multicastAddresssRangeId)
vdnMulticastPoolShow

Retrieve information about the specified multicast address range.   Parameters:  multicastAddresssRangeId: A valid multicast address range ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **multicastAddresssRangeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnConfigMulticastsMulticastAddresssRangeIdPut**
> VdnConfigMulticastsMulticastAddresssRangeIdPut(ctx, multicastAddresssRangeId, optional)
vdnMulticastPoolChange

Update the specified multicast address range.  If the multicast address range is universal you must send the API request to the primary NSX Manager.   Parameters:  multicastAddresssRangeId: A valid multicast address range ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **multicastAddresssRangeId** | **string**|  | 
 **optional** | ***ConfigApiVdnConfigMulticastsMulticastAddresssRangeIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiVdnConfigMulticastsMulticastAddresssRangeIdPutOpts struct
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

# **VdnConfigMulticastsPost**
> VdnConfigMulticastsPost(ctx, optional)
vdnMulticastPoolAdd

Add a multicast address range for logical switches.  The address range includes the beginning and ending addresses.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiVdnConfigMulticastsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiVdnConfigMulticastsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **isUniversal** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnConfigResourcesAllocatedGet**
> VdnConfigResourcesAllocatedGet(ctx, optional)
allocatedResourcesRead

Retrieve information about allocated segment IDs or multicast addresses.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiVdnConfigResourcesAllocatedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiVdnConfigResourcesAllocatedGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**|  | 
 **pagesize** | **optional.String**|  | 
 **startindex** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnConfigSegmentsGet**
> VdnConfigSegmentsGet(ctx, )
vdnSegmentPoolList

Retrieve information about all segment ID pools.   Parameters:  

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

# **VdnConfigSegmentsPost**
> VdnConfigSegmentsPost(ctx, optional)
vdnSegmentPoolAdd

Add a segment ID pool.  * **name** - Required property. * **desc** - Optional property. * **begin** - Required property. Minimum value is *5000* * **end** - Required property. Maximum value is *16777216*   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiVdnConfigSegmentsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiVdnConfigSegmentsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **isUniversal** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnConfigSegmentsSegmentPoolIdDelete**
> VdnConfigSegmentsSegmentPoolIdDelete(ctx, segmentPoolId)
vdnSegmentPoolDelete

Delete the specified segment ID pool.  If the segment ID pool is universal you must send the API request to the primary NSX Manager.   Parameters:  segmentPoolId: A valid *segmentPoolId*  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentPoolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnConfigSegmentsSegmentPoolIdGet**
> VdnConfigSegmentsSegmentPoolIdGet(ctx, segmentPoolId)
vdnSegmentPoolShow

Retrieve information about the specified segment ID pool.   Parameters:  segmentPoolId: A valid *segmentPoolId*  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentPoolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnConfigSegmentsSegmentPoolIdPut**
> VdnConfigSegmentsSegmentPoolIdPut(ctx, segmentPoolId, optional)
vdnSegmentPoolChange

Update the specified segment ID pool.  If the segment ID pool is universal you must send the API request to the primary NSX Manager.   Parameters:  segmentPoolId: A valid *segmentPoolId*  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentPoolId** | **string**|  | 
 **optional** | ***ConfigApiVdnConfigSegmentsSegmentPoolIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiVdnConfigSegmentsSegmentPoolIdPutOpts struct
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

# **VdnConfigVxlanUdpPortGet**
> VdnConfigVxlanUdpPortGet(ctx, )
vdnConfigUDPShow

Retrieve the UDP port configured for VXLAN traffic.   Parameters:  

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

# **VdnConfigVxlanUdpPortPortNumberPut**
> VdnConfigVxlanUdpPortPortNumberPut(ctx, portNumber, optional)
vdnConfigUDPUpdateUpdate

Update the VXLAN port configuration to use port *portNumber*.  This method changes the VXLAN port in a three phrase process, avoiding disruption of VXLAN traffic. In a cross-vCenter NSX environment, change the VXLAN port on the primary NSX Manager to propagate this change on all NSX Managers and hosts in the cross-vCenter NSX environment.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method updated. Port change is now non-disruptive, and propagates to secondary NSX Managers if performed on the primary NSX Manager. Force parameter added.   Parameters:  portNumber: A valid UDP port for VXLAN  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portNumber** | **string**|  | 
 **optional** | ***ConfigApiVdnConfigVxlanUdpPortPortNumberPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiVdnConfigVxlanUdpPortPortNumberPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnConfigVxlanUdpPortTaskStatusGet**
> VdnConfigVxlanUdpPortTaskStatusGet(ctx, )
vdnConfigUDPUpdateStatusRead

Retrieve the status of the VXLAN port configuration update.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

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

