# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20VdnHardwaregatewaysBfdConfigGet**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysBfdConfigGet) | **Get** /api/2.0/vdn/hardwaregateways/bfd/config | hardwareGatewayBfdConfigRead
[**Api20VdnHardwaregatewaysBfdConfigPut**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysBfdConfigPut) | **Put** /api/2.0/vdn/hardwaregateways/bfd/config | hardwareGatewayBfdConfigUpdate
[**Api20VdnHardwaregatewaysBfdStatusGet**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysBfdStatusGet) | **Get** /api/2.0/vdn/hardwaregateways/bfd/status | hardwareGatewayBfdStatusRead
[**Api20VdnHardwaregatewaysBindingsBindingIdDelete**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysBindingsBindingIdDelete) | **Delete** /api/2.0/vdn/hardwaregateways/bindings/{bindingId} | hardwareGatewayBindingDelete
[**Api20VdnHardwaregatewaysBindingsBindingIdGet**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysBindingsBindingIdGet) | **Get** /api/2.0/vdn/hardwaregateways/bindings/{bindingId} | hardwareGatewayBindingRead
[**Api20VdnHardwaregatewaysBindingsBindingIdPut**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysBindingsBindingIdPut) | **Put** /api/2.0/vdn/hardwaregateways/bindings/{bindingId} | hardwareGatewayBindingUpdate
[**Api20VdnHardwaregatewaysBindingsBindingIdStatisticGet**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysBindingsBindingIdStatisticGet) | **Get** /api/2.0/vdn/hardwaregateways/bindings/{bindingId}/statistic | hardwareGatewayBindingStatisticRead
[**Api20VdnHardwaregatewaysBindingsGet**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysBindingsGet) | **Get** /api/2.0/vdn/hardwaregateways/bindings | hardwareGatewayBindingsList
[**Api20VdnHardwaregatewaysBindingsManagePost**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysBindingsManagePost) | **Post** /api/2.0/vdn/hardwaregateways/bindings/manage | hardwareGatewayBindingsManageOperations
[**Api20VdnHardwaregatewaysBindingsPost**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysBindingsPost) | **Post** /api/2.0/vdn/hardwaregateways/bindings | hardwareGatewayBindingCreate
[**Api20VdnHardwaregatewaysGet**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysGet) | **Get** /api/2.0/vdn/hardwaregateways | hardwareGatewaysList
[**Api20VdnHardwaregatewaysHardwareGatewayIdDelete**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysHardwareGatewayIdDelete) | **Delete** /api/2.0/vdn/hardwaregateways/{hardwareGatewayId} | hardwareGatewayDelete
[**Api20VdnHardwaregatewaysHardwareGatewayIdGet**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysHardwareGatewayIdGet) | **Get** /api/2.0/vdn/hardwaregateways/{hardwareGatewayId} | hardwareGatewayRead
[**Api20VdnHardwaregatewaysHardwareGatewayIdPut**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysHardwareGatewayIdPut) | **Put** /api/2.0/vdn/hardwaregateways/{hardwareGatewayId} | hardwareGatewayUpdate
[**Api20VdnHardwaregatewaysHardwareGatewayIdSwitchesGet**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysHardwareGatewayIdSwitchesGet) | **Get** /api/2.0/vdn/hardwaregateways/{hardwareGatewayId}/switches | hardwareGatewaySwitchesList
[**Api20VdnHardwaregatewaysHardwareGatewayIdSwitchesSwitchNameSwitchportsGet**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysHardwareGatewayIdSwitchesSwitchNameSwitchportsGet) | **Get** /api/2.0/vdn/hardwaregateways/{hardwareGatewayId}/switches/{switchName}/switchports | hardwareGatewaySwitchPortsList
[**Api20VdnHardwaregatewaysPost**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysPost) | **Post** /api/2.0/vdn/hardwaregateways | hardwareGatewayCreate
[**Api20VdnHardwaregatewaysReplicationclusterGet**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysReplicationclusterGet) | **Get** /api/2.0/vdn/hardwaregateways/replicationcluster | hardwareGatewayReplicationClusterRead
[**Api20VdnHardwaregatewaysReplicationclusterPut**](HardwaregatewaysApi.md#Api20VdnHardwaregatewaysReplicationclusterPut) | **Put** /api/2.0/vdn/hardwaregateways/replicationcluster | hardwareGatewayReplicationClusterUpdate

# **Api20VdnHardwaregatewaysBfdConfigGet**
> Api20VdnHardwaregatewaysBfdConfigGet(ctx, )
hardwareGatewayBfdConfigRead

Retrieve global hardware gateway BFD configuration.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

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

# **Api20VdnHardwaregatewaysBfdConfigPut**
> Api20VdnHardwaregatewaysBfdConfigPut(ctx, optional)
hardwareGatewayBfdConfigUpdate

Update global hardware gateway BFD configuration.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiApi20VdnHardwaregatewaysBfdConfigPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiApi20VdnHardwaregatewaysBfdConfigPutOpts struct
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

# **Api20VdnHardwaregatewaysBfdStatusGet**
> Api20VdnHardwaregatewaysBfdStatusGet(ctx, )
hardwareGatewayBfdStatusRead

Retrieve hardware gateway BFD tunnel status for all tunnel endpoints, including hosts and hardware gateways.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

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

# **Api20VdnHardwaregatewaysBindingsBindingIdDelete**
> Api20VdnHardwaregatewaysBindingsBindingIdDelete(ctx, bindingId)
hardwareGatewayBindingDelete

Delete the specified hardware gateway binding.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  bindingId: hardware gateway binding ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnHardwaregatewaysBindingsBindingIdGet**
> Api20VdnHardwaregatewaysBindingsBindingIdGet(ctx, bindingId)
hardwareGatewayBindingRead

Retrieve information about the specified hardware gateway binding.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  bindingId: hardware gateway binding ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnHardwaregatewaysBindingsBindingIdPut**
> Api20VdnHardwaregatewaysBindingsBindingIdPut(ctx, bindingId, optional)
hardwareGatewayBindingUpdate

Update the specified hardware gateway binding.  You can update the binding parameters. This API will fail if: * the specified *hardwareGatewayId* does not exist. * the specified logical switch (*virtualWire*) is not present or there is a software   gateway on the binding. * the new binding value is a duplicate of an existing binding.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  bindingId: hardware gateway binding ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bindingId** | **string**|  | 
 **optional** | ***HardwaregatewaysApiApi20VdnHardwaregatewaysBindingsBindingIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiApi20VdnHardwaregatewaysBindingsBindingIdPutOpts struct
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

# **Api20VdnHardwaregatewaysBindingsBindingIdStatisticGet**
> Api20VdnHardwaregatewaysBindingsBindingIdStatisticGet(ctx, bindingId)
hardwareGatewayBindingStatisticRead

Retrieve statistics for the specified hardware gateway binding.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnHardwaregatewaysBindingsGet**
> Api20VdnHardwaregatewaysBindingsGet(ctx, optional)
hardwareGatewayBindingsList

Retrieve information about hardware gateway bindings.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiApi20VdnHardwaregatewaysBindingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiApi20VdnHardwaregatewaysBindingsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hardwareGatewayId** | **optional.String**|  | 
 **vni** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnHardwaregatewaysBindingsManagePost**
> Api20VdnHardwaregatewaysBindingsManagePost(ctx, optional)
hardwareGatewayBindingsManageOperations

Manage hardware gateway binding objects.  Use this API to attach, detach, and update multiple bindings in a single API call.  This API accepts three lists for add, update, and delete. Each list accepts a hardwareGatewayManageBindingsItem with a full description of the new binding with its objectID. This API handles a maximum of 100 HardwareGatewayManageBindingsItem objects for each of the Add/Update/Delete lists.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiApi20VdnHardwaregatewaysBindingsManagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiApi20VdnHardwaregatewaysBindingsManagePostOpts struct
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

# **Api20VdnHardwaregatewaysBindingsPost**
> Api20VdnHardwaregatewaysBindingsPost(ctx, optional)
hardwareGatewayBindingCreate

Create a hardware gateway binding.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiApi20VdnHardwaregatewaysBindingsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiApi20VdnHardwaregatewaysBindingsPostOpts struct
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

# **Api20VdnHardwaregatewaysGet**
> Api20VdnHardwaregatewaysGet(ctx, )
hardwareGatewaysList

Retrieve information about all hardware gateways.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

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

# **Api20VdnHardwaregatewaysHardwareGatewayIdDelete**
> Api20VdnHardwaregatewaysHardwareGatewayIdDelete(ctx, hardwareGatewayId)
hardwareGatewayDelete

Delete the specified hardware gateway.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  hardwareGatewayId: Object ID of the hardware gateway.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hardwareGatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnHardwaregatewaysHardwareGatewayIdGet**
> Api20VdnHardwaregatewaysHardwareGatewayIdGet(ctx, hardwareGatewayId)
hardwareGatewayRead

Retrieve information about the specified hardware gateway.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  hardwareGatewayId: Object ID of the hardware gateway.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hardwareGatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnHardwaregatewaysHardwareGatewayIdPut**
> Api20VdnHardwaregatewaysHardwareGatewayIdPut(ctx, hardwareGatewayId, optional)
hardwareGatewayUpdate

Update the specified hardware gateway.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  hardwareGatewayId: Object ID of the hardware gateway.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hardwareGatewayId** | **string**|  | 
 **optional** | ***HardwaregatewaysApiApi20VdnHardwaregatewaysHardwareGatewayIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiApi20VdnHardwaregatewaysHardwareGatewayIdPutOpts struct
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

# **Api20VdnHardwaregatewaysHardwareGatewayIdSwitchesGet**
> Api20VdnHardwaregatewaysHardwareGatewayIdSwitchesGet(ctx, hardwareGatewayId)
hardwareGatewaySwitchesList

Retrieve information about switches on the specified hardware gateway.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hardwareGatewayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnHardwaregatewaysHardwareGatewayIdSwitchesSwitchNameSwitchportsGet**
> Api20VdnHardwaregatewaysHardwareGatewayIdSwitchesSwitchNameSwitchportsGet(ctx, hardwareGatewayId, switchName)
hardwareGatewaySwitchPortsList

Retrive information about the hardware gateway switch ports for the specified switch and hardware gateway.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hardwareGatewayId** | **string**|  | 
  **switchName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20VdnHardwaregatewaysPost**
> Api20VdnHardwaregatewaysPost(ctx, optional)
hardwareGatewayCreate

Install a hardware gateway.  **bfdEnabled** is true by default.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiApi20VdnHardwaregatewaysPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiApi20VdnHardwaregatewaysPostOpts struct
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

# **Api20VdnHardwaregatewaysReplicationclusterGet**
> Api20VdnHardwaregatewaysReplicationclusterGet(ctx, )
hardwareGatewayReplicationClusterRead

Retrieve information about the hardware gateway replication cluster.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

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

# **Api20VdnHardwaregatewaysReplicationclusterPut**
> Api20VdnHardwaregatewaysReplicationclusterPut(ctx, optional)
hardwareGatewayReplicationClusterUpdate

Update the hardware gateway replication cluster.  Add or remove hosts on a replication cluster.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiApi20VdnHardwaregatewaysReplicationclusterPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiApi20VdnHardwaregatewaysReplicationclusterPutOpts struct
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

