# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VdnHardwaregatewaysBfdConfigGet**](HardwaregatewaysApi.md#VdnHardwaregatewaysBfdConfigGet) | **Get** /api/2.0/vdn/hardwaregateways/bfd/config | hardwareGatewayBfdConfigRead
[**VdnHardwaregatewaysBfdConfigPut**](HardwaregatewaysApi.md#VdnHardwaregatewaysBfdConfigPut) | **Put** /api/2.0/vdn/hardwaregateways/bfd/config | hardwareGatewayBfdConfigUpdate
[**VdnHardwaregatewaysBfdStatusGet**](HardwaregatewaysApi.md#VdnHardwaregatewaysBfdStatusGet) | **Get** /api/2.0/vdn/hardwaregateways/bfd/status | hardwareGatewayBfdStatusRead
[**VdnHardwaregatewaysBindingsBindingIdDelete**](HardwaregatewaysApi.md#VdnHardwaregatewaysBindingsBindingIdDelete) | **Delete** /api/2.0/vdn/hardwaregateways/bindings/{bindingId} | hardwareGatewayBindingDelete
[**VdnHardwaregatewaysBindingsBindingIdGet**](HardwaregatewaysApi.md#VdnHardwaregatewaysBindingsBindingIdGet) | **Get** /api/2.0/vdn/hardwaregateways/bindings/{bindingId} | hardwareGatewayBindingRead
[**VdnHardwaregatewaysBindingsBindingIdPut**](HardwaregatewaysApi.md#VdnHardwaregatewaysBindingsBindingIdPut) | **Put** /api/2.0/vdn/hardwaregateways/bindings/{bindingId} | hardwareGatewayBindingUpdate
[**VdnHardwaregatewaysBindingsBindingIdStatisticGet**](HardwaregatewaysApi.md#VdnHardwaregatewaysBindingsBindingIdStatisticGet) | **Get** /api/2.0/vdn/hardwaregateways/bindings/{bindingId}/statistic | hardwareGatewayBindingStatisticRead
[**VdnHardwaregatewaysBindingsGet**](HardwaregatewaysApi.md#VdnHardwaregatewaysBindingsGet) | **Get** /api/2.0/vdn/hardwaregateways/bindings | hardwareGatewayBindingsList
[**VdnHardwaregatewaysBindingsManagePost**](HardwaregatewaysApi.md#VdnHardwaregatewaysBindingsManagePost) | **Post** /api/2.0/vdn/hardwaregateways/bindings/manage | hardwareGatewayBindingsManageOperations
[**VdnHardwaregatewaysBindingsPost**](HardwaregatewaysApi.md#VdnHardwaregatewaysBindingsPost) | **Post** /api/2.0/vdn/hardwaregateways/bindings | hardwareGatewayBindingCreate
[**VdnHardwaregatewaysGet**](HardwaregatewaysApi.md#VdnHardwaregatewaysGet) | **Get** /api/2.0/vdn/hardwaregateways | hardwareGatewaysList
[**VdnHardwaregatewaysHardwareGatewayIdDelete**](HardwaregatewaysApi.md#VdnHardwaregatewaysHardwareGatewayIdDelete) | **Delete** /api/2.0/vdn/hardwaregateways/{hardwareGatewayId} | hardwareGatewayDelete
[**VdnHardwaregatewaysHardwareGatewayIdGet**](HardwaregatewaysApi.md#VdnHardwaregatewaysHardwareGatewayIdGet) | **Get** /api/2.0/vdn/hardwaregateways/{hardwareGatewayId} | hardwareGatewayRead
[**VdnHardwaregatewaysHardwareGatewayIdPut**](HardwaregatewaysApi.md#VdnHardwaregatewaysHardwareGatewayIdPut) | **Put** /api/2.0/vdn/hardwaregateways/{hardwareGatewayId} | hardwareGatewayUpdate
[**VdnHardwaregatewaysHardwareGatewayIdSwitchesGet**](HardwaregatewaysApi.md#VdnHardwaregatewaysHardwareGatewayIdSwitchesGet) | **Get** /api/2.0/vdn/hardwaregateways/{hardwareGatewayId}/switches | hardwareGatewaySwitchesList
[**VdnHardwaregatewaysHardwareGatewayIdSwitchesSwitchNameSwitchportsGet**](HardwaregatewaysApi.md#VdnHardwaregatewaysHardwareGatewayIdSwitchesSwitchNameSwitchportsGet) | **Get** /api/2.0/vdn/hardwaregateways/{hardwareGatewayId}/switches/{switchName}/switchports | hardwareGatewaySwitchPortsList
[**VdnHardwaregatewaysPost**](HardwaregatewaysApi.md#VdnHardwaregatewaysPost) | **Post** /api/2.0/vdn/hardwaregateways | hardwareGatewayCreate
[**VdnHardwaregatewaysReplicationclusterGet**](HardwaregatewaysApi.md#VdnHardwaregatewaysReplicationclusterGet) | **Get** /api/2.0/vdn/hardwaregateways/replicationcluster | hardwareGatewayReplicationClusterRead
[**VdnHardwaregatewaysReplicationclusterPut**](HardwaregatewaysApi.md#VdnHardwaregatewaysReplicationclusterPut) | **Put** /api/2.0/vdn/hardwaregateways/replicationcluster | hardwareGatewayReplicationClusterUpdate

# **VdnHardwaregatewaysBfdConfigGet**
> VdnHardwaregatewaysBfdConfigGet(ctx, )
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

# **VdnHardwaregatewaysBfdConfigPut**
> VdnHardwaregatewaysBfdConfigPut(ctx, optional)
hardwareGatewayBfdConfigUpdate

Update global hardware gateway BFD configuration.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiVdnHardwaregatewaysBfdConfigPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiVdnHardwaregatewaysBfdConfigPutOpts struct
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

# **VdnHardwaregatewaysBfdStatusGet**
> VdnHardwaregatewaysBfdStatusGet(ctx, )
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

# **VdnHardwaregatewaysBindingsBindingIdDelete**
> VdnHardwaregatewaysBindingsBindingIdDelete(ctx, bindingId)
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

# **VdnHardwaregatewaysBindingsBindingIdGet**
> VdnHardwaregatewaysBindingsBindingIdGet(ctx, bindingId)
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

# **VdnHardwaregatewaysBindingsBindingIdPut**
> VdnHardwaregatewaysBindingsBindingIdPut(ctx, bindingId, optional)
hardwareGatewayBindingUpdate

Update the specified hardware gateway binding.  You can update the binding parameters. This API will fail if: * the specified *hardwareGatewayId* does not exist. * the specified logical switch (*virtualWire*) is not present or there is a software   gateway on the binding. * the new binding value is a duplicate of an existing binding.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  bindingId: hardware gateway binding ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bindingId** | **string**|  | 
 **optional** | ***HardwaregatewaysApiVdnHardwaregatewaysBindingsBindingIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiVdnHardwaregatewaysBindingsBindingIdPutOpts struct
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

# **VdnHardwaregatewaysBindingsBindingIdStatisticGet**
> VdnHardwaregatewaysBindingsBindingIdStatisticGet(ctx, bindingId)
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

# **VdnHardwaregatewaysBindingsGet**
> VdnHardwaregatewaysBindingsGet(ctx, optional)
hardwareGatewayBindingsList

Retrieve information about hardware gateway bindings.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiVdnHardwaregatewaysBindingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiVdnHardwaregatewaysBindingsGetOpts struct
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

# **VdnHardwaregatewaysBindingsManagePost**
> VdnHardwaregatewaysBindingsManagePost(ctx, optional)
hardwareGatewayBindingsManageOperations

Manage hardware gateway binding objects.  Use this API to attach, detach, and update multiple bindings in a single API call.  This API accepts three lists for add, update, and delete. Each list accepts a hardwareGatewayManageBindingsItem with a full description of the new binding with its objectID. This API handles a maximum of 100 HardwareGatewayManageBindingsItem objects for each of the Add/Update/Delete lists.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiVdnHardwaregatewaysBindingsManagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiVdnHardwaregatewaysBindingsManagePostOpts struct
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

# **VdnHardwaregatewaysBindingsPost**
> VdnHardwaregatewaysBindingsPost(ctx, optional)
hardwareGatewayBindingCreate

Create a hardware gateway binding.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiVdnHardwaregatewaysBindingsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiVdnHardwaregatewaysBindingsPostOpts struct
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

# **VdnHardwaregatewaysGet**
> VdnHardwaregatewaysGet(ctx, )
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

# **VdnHardwaregatewaysHardwareGatewayIdDelete**
> VdnHardwaregatewaysHardwareGatewayIdDelete(ctx, hardwareGatewayId)
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

# **VdnHardwaregatewaysHardwareGatewayIdGet**
> VdnHardwaregatewaysHardwareGatewayIdGet(ctx, hardwareGatewayId)
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

# **VdnHardwaregatewaysHardwareGatewayIdPut**
> VdnHardwaregatewaysHardwareGatewayIdPut(ctx, hardwareGatewayId, optional)
hardwareGatewayUpdate

Update the specified hardware gateway.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  hardwareGatewayId: Object ID of the hardware gateway.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hardwareGatewayId** | **string**|  | 
 **optional** | ***HardwaregatewaysApiVdnHardwaregatewaysHardwareGatewayIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiVdnHardwaregatewaysHardwareGatewayIdPutOpts struct
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

# **VdnHardwaregatewaysHardwareGatewayIdSwitchesGet**
> VdnHardwaregatewaysHardwareGatewayIdSwitchesGet(ctx, hardwareGatewayId)
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

# **VdnHardwaregatewaysHardwareGatewayIdSwitchesSwitchNameSwitchportsGet**
> VdnHardwaregatewaysHardwareGatewayIdSwitchesSwitchNameSwitchportsGet(ctx, hardwareGatewayId, switchName)
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

# **VdnHardwaregatewaysPost**
> VdnHardwaregatewaysPost(ctx, optional)
hardwareGatewayCreate

Install a hardware gateway.  **bfdEnabled** is true by default.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiVdnHardwaregatewaysPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiVdnHardwaregatewaysPostOpts struct
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

# **VdnHardwaregatewaysReplicationclusterGet**
> VdnHardwaregatewaysReplicationclusterGet(ctx, )
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

# **VdnHardwaregatewaysReplicationclusterPut**
> VdnHardwaregatewaysReplicationclusterPut(ctx, optional)
hardwareGatewayReplicationClusterUpdate

Update the hardware gateway replication cluster.  Add or remove hosts on a replication cluster.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HardwaregatewaysApiVdnHardwaregatewaysReplicationclusterPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HardwaregatewaysApiVdnHardwaregatewaysReplicationclusterPutOpts struct
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

