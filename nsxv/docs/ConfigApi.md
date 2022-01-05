# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20VdnConfigHostHostIdVxlanVtepsPost**](ConfigApi.md#Api20VdnConfigHostHostIdVxlanVtepsPost) | **Post** /api/2.0/vdn/config/host/{hostId}/vxlan/vteps | hostVtepResolveAction
[**Api20VdnConfigMulticastsGet**](ConfigApi.md#Api20VdnConfigMulticastsGet) | **Get** /api/2.0/vdn/config/multicasts | vdnMulticastPoolList
[**Api20VdnConfigMulticastsMulticastAddresssRangeIdDelete**](ConfigApi.md#Api20VdnConfigMulticastsMulticastAddresssRangeIdDelete) | **Delete** /api/2.0/vdn/config/multicasts/{multicastAddresssRangeId} | vdnMulticastPoolDelete
[**Api20VdnConfigMulticastsMulticastAddresssRangeIdGet**](ConfigApi.md#Api20VdnConfigMulticastsMulticastAddresssRangeIdGet) | **Get** /api/2.0/vdn/config/multicasts/{multicastAddresssRangeId} | vdnMulticastPoolShow
[**Api20VdnConfigMulticastsMulticastAddresssRangeIdPut**](ConfigApi.md#Api20VdnConfigMulticastsMulticastAddresssRangeIdPut) | **Put** /api/2.0/vdn/config/multicasts/{multicastAddresssRangeId} | vdnMulticastPoolChange
[**Api20VdnConfigMulticastsPost**](ConfigApi.md#Api20VdnConfigMulticastsPost) | **Post** /api/2.0/vdn/config/multicasts | vdnMulticastPoolAdd
[**Api20VdnConfigResourcesAllocatedGet**](ConfigApi.md#Api20VdnConfigResourcesAllocatedGet) | **Get** /api/2.0/vdn/config/resources/allocated | allocatedResourcesRead
[**Api20VdnConfigSegmentsGet**](ConfigApi.md#Api20VdnConfigSegmentsGet) | **Get** /api/2.0/vdn/config/segments | vdnSegmentPoolList
[**Api20VdnConfigSegmentsPost**](ConfigApi.md#Api20VdnConfigSegmentsPost) | **Post** /api/2.0/vdn/config/segments | vdnSegmentPoolAdd
[**Api20VdnConfigSegmentsSegmentPoolIdDelete**](ConfigApi.md#Api20VdnConfigSegmentsSegmentPoolIdDelete) | **Delete** /api/2.0/vdn/config/segments/{segmentPoolId} | vdnSegmentPoolDelete
[**Api20VdnConfigSegmentsSegmentPoolIdGet**](ConfigApi.md#Api20VdnConfigSegmentsSegmentPoolIdGet) | **Get** /api/2.0/vdn/config/segments/{segmentPoolId} | vdnSegmentPoolShow
[**Api20VdnConfigSegmentsSegmentPoolIdPut**](ConfigApi.md#Api20VdnConfigSegmentsSegmentPoolIdPut) | **Put** /api/2.0/vdn/config/segments/{segmentPoolId} | vdnSegmentPoolChange
[**Api20VdnConfigVxlanUdpPortGet**](ConfigApi.md#Api20VdnConfigVxlanUdpPortGet) | **Get** /api/2.0/vdn/config/vxlan/udp/port | vdnConfigUDPShow
[**Api20VdnConfigVxlanUdpPortPortNumberPut**](ConfigApi.md#Api20VdnConfigVxlanUdpPortPortNumberPut) | **Put** /api/2.0/vdn/config/vxlan/udp/port/{portNumber} | vdnConfigUDPUpdateUpdate
[**Api20VdnConfigVxlanUdpPortTaskStatusGet**](ConfigApi.md#Api20VdnConfigVxlanUdpPortTaskStatusGet) | **Get** /api/2.0/vdn/config/vxlan/udp/port/taskStatus | vdnConfigUDPUpdateStatusRead

# **Api20VdnConfigHostHostIdVxlanVtepsPost**
> Api20VdnConfigHostHostIdVxlanVtepsPost(ctx, hostId, optional)
hostVtepResolveAction

Resolve missing VXLAN VMKernel adapters.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  hostId:   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**|  | 
 **optional** | ***ConfigApiApi20VdnConfigHostHostIdVxlanVtepsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiApi20VdnConfigHostHostIdVxlanVtepsPostOpts struct
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

# **Api20VdnConfigMulticastsGet**
> Api20VdnConfigMulticastsGet(ctx, )
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

# **Api20VdnConfigMulticastsMulticastAddresssRangeIdDelete**
> Api20VdnConfigMulticastsMulticastAddresssRangeIdDelete(ctx, multicastAddresssRangeId)
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

# **Api20VdnConfigMulticastsMulticastAddresssRangeIdGet**
> Api20VdnConfigMulticastsMulticastAddresssRangeIdGet(ctx, multicastAddresssRangeId)
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

# **Api20VdnConfigMulticastsMulticastAddresssRangeIdPut**
> Api20VdnConfigMulticastsMulticastAddresssRangeIdPut(ctx, multicastAddresssRangeId, optional)
vdnMulticastPoolChange

Update the specified multicast address range.  If the multicast address range is universal you must send the API request to the primary NSX Manager.   Parameters:  multicastAddresssRangeId: A valid multicast address range ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **multicastAddresssRangeId** | **string**|  | 
 **optional** | ***ConfigApiApi20VdnConfigMulticastsMulticastAddresssRangeIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiApi20VdnConfigMulticastsMulticastAddresssRangeIdPutOpts struct
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

# **Api20VdnConfigMulticastsPost**
> Api20VdnConfigMulticastsPost(ctx, optional)
vdnMulticastPoolAdd

Add a multicast address range for logical switches.  The address range includes the beginning and ending addresses.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiApi20VdnConfigMulticastsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiApi20VdnConfigMulticastsPostOpts struct
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

# **Api20VdnConfigResourcesAllocatedGet**
> Api20VdnConfigResourcesAllocatedGet(ctx, optional)
allocatedResourcesRead

Retrieve information about allocated segment IDs or multicast addresses.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiApi20VdnConfigResourcesAllocatedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiApi20VdnConfigResourcesAllocatedGetOpts struct
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

# **Api20VdnConfigSegmentsGet**
> Api20VdnConfigSegmentsGet(ctx, )
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

# **Api20VdnConfigSegmentsPost**
> Api20VdnConfigSegmentsPost(ctx, optional)
vdnSegmentPoolAdd

Add a segment ID pool.  * **name** - Required property. * **desc** - Optional property. * **begin** - Required property. Minimum value is *5000* * **end** - Required property. Maximum value is *16777216*   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiApi20VdnConfigSegmentsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiApi20VdnConfigSegmentsPostOpts struct
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

# **Api20VdnConfigSegmentsSegmentPoolIdDelete**
> Api20VdnConfigSegmentsSegmentPoolIdDelete(ctx, segmentPoolId)
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

# **Api20VdnConfigSegmentsSegmentPoolIdGet**
> Api20VdnConfigSegmentsSegmentPoolIdGet(ctx, segmentPoolId)
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

# **Api20VdnConfigSegmentsSegmentPoolIdPut**
> Api20VdnConfigSegmentsSegmentPoolIdPut(ctx, segmentPoolId, optional)
vdnSegmentPoolChange

Update the specified segment ID pool.  If the segment ID pool is universal you must send the API request to the primary NSX Manager.   Parameters:  segmentPoolId: A valid *segmentPoolId*  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentPoolId** | **string**|  | 
 **optional** | ***ConfigApiApi20VdnConfigSegmentsSegmentPoolIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiApi20VdnConfigSegmentsSegmentPoolIdPutOpts struct
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

# **Api20VdnConfigVxlanUdpPortGet**
> Api20VdnConfigVxlanUdpPortGet(ctx, )
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

# **Api20VdnConfigVxlanUdpPortPortNumberPut**
> Api20VdnConfigVxlanUdpPortPortNumberPut(ctx, portNumber, optional)
vdnConfigUDPUpdateUpdate

Update the VXLAN port configuration to use port *portNumber*.  This method changes the VXLAN port in a three phrase process, avoiding disruption of VXLAN traffic. In a cross-vCenter NSX environment, change the VXLAN port on the primary NSX Manager to propagate this change on all NSX Managers and hosts in the cross-vCenter NSX environment.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method updated. Port change is now non-disruptive, and propagates to secondary NSX Managers if performed on the primary NSX Manager. Force parameter added.   Parameters:  portNumber: A valid UDP port for VXLAN  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portNumber** | **string**|  | 
 **optional** | ***ConfigApiApi20VdnConfigVxlanUdpPortPortNumberPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiApi20VdnConfigVxlanUdpPortPortNumberPutOpts struct
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

# **Api20VdnConfigVxlanUdpPortTaskStatusGet**
> Api20VdnConfigVxlanUdpPortTaskStatusGet(ctx, )
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

