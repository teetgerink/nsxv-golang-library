# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VdnInventoryHostHostIdConnectionStatusGet**](InventoryApi.md#VdnInventoryHostHostIdConnectionStatusGet) | **Get** /api/2.0/vdn/inventory/host/{hostId}/connection/status | inventoryStatusHostRead
[**VdnInventoryHostsConnectionStatusGet**](InventoryApi.md#VdnInventoryHostsConnectionStatusGet) | **Get** /api/2.0/vdn/inventory/hosts/connection/status | inventoryStatusHostsList

# **VdnInventoryHostHostIdConnectionStatusGet**
> VdnInventoryHostHostIdConnectionStatusGet(ctx, hostId)
inventoryStatusHostRead

Retrieve the status of the specified host.  History:  Release | Modification --------|------------- 6.2.3 | Method updated. Introduced **hostToControllerConnectionErrors** array.<br>Deprecated **fullSyncCount** parameter. Parameter is still present, but always has value of -1.   Parameters:  hostId: ID of the host to check.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnInventoryHostsConnectionStatusGet**
> VdnInventoryHostsConnectionStatusGet(ctx, optional)
inventoryStatusHostsList

Retrieve the status of a list of hosts.  Release | Modification --------|------------- 6.2.3 | Method updated. Introduced **hostToControllerConnectionErrors** array.<br>Deprecated **fullSyncCount** parameter. Parameter is still present, but always has value of -1.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InventoryApiVdnInventoryHostsConnectionStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoryApiVdnInventoryHostsConnectionStatusGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostId** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

