# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VdnScopesScopeIdVirtualwiresGet**](VirtualwiresApi.md#VdnScopesScopeIdVirtualwiresGet) | **Get** /api/2.0/vdn/scopes/{scopeId}/virtualwires | logicalSwitchList
[**VdnScopesScopeIdVirtualwiresPost**](VirtualwiresApi.md#VdnScopesScopeIdVirtualwiresPost) | **Post** /api/2.0/vdn/scopes/{scopeId}/virtualwires | logicalSwitchCreate
[**VdnVirtualwiresGet**](VirtualwiresApi.md#VdnVirtualwiresGet) | **Get** /api/2.0/vdn/virtualwires | logicalSwitchesRead
[**VdnVirtualwiresVirtualWireIDBackingPost**](VirtualwiresApi.md#VdnVirtualwiresVirtualWireIDBackingPost) | **Post** /api/2.0/vdn/virtualwires/{virtualWireID}/backing | logicalSwitchPortGroupFixAction
[**VdnVirtualwiresVirtualWireIDConnCheckMulticastPost**](VirtualwiresApi.md#VdnVirtualwiresVirtualWireIDConnCheckMulticastPost) | **Post** /api/2.0/vdn/virtualwires/{virtualWireID}/conn-check/multicast | logicalSwitchConnCheckExecute
[**VdnVirtualwiresVirtualWireIDConnCheckP2pPost**](VirtualwiresApi.md#VdnVirtualwiresVirtualWireIDConnCheckP2pPost) | **Post** /api/2.0/vdn/virtualwires/{virtualWireID}/conn-check/p2p | logicalSwitchPingExecute
[**VdnVirtualwiresVirtualWireIDDelete**](VirtualwiresApi.md#VdnVirtualwiresVirtualWireIDDelete) | **Delete** /api/2.0/vdn/virtualwires/{virtualWireID} | logicalSwitchDelete
[**VdnVirtualwiresVirtualWireIDGet**](VirtualwiresApi.md#VdnVirtualwiresVirtualWireIDGet) | **Get** /api/2.0/vdn/virtualwires/{virtualWireID} | logicalSwitchShow
[**VdnVirtualwiresVirtualWireIDHardwaregatewaysGet**](VirtualwiresApi.md#VdnVirtualwiresVirtualWireIDHardwaregatewaysGet) | **Get** /api/2.0/vdn/virtualwires/{virtualWireID}/hardwaregateways | logicalSwitchHardwareGatewayBindingsList
[**VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPost**](VirtualwiresApi.md#VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPost) | **Post** /api/2.0/vdn/virtualwires/{virtualWireID}/hardwaregateways/{hardwareGatewayBindingId} | logicalSwitchHardwareGatewayBindingCreate
[**VdnVirtualwiresVirtualWireIDPut**](VirtualwiresApi.md#VdnVirtualwiresVirtualWireIDPut) | **Put** /api/2.0/vdn/virtualwires/{virtualWireID} | logicalSwitchUpdate
[**VdnVirtualwiresVmVnicPost**](VirtualwiresApi.md#VdnVirtualwiresVmVnicPost) | **Post** /api/2.0/vdn/virtualwires/vm/vnic | logicalSwitchVmAttachCreate

# **VdnScopesScopeIdVirtualwiresGet**
> VdnScopesScopeIdVirtualwiresGet(ctx, scopeId, optional)
logicalSwitchList

Retrieve information about all logical switches in the specified transport zone (network scope).   Parameters:  scopeId: A valid transport zone ID (vdnScope objectId).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***VirtualwiresApiVdnScopesScopeIdVirtualwiresGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiVdnScopesScopeIdVirtualwiresGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startindex** | **optional.String**|  | 
 **pagesize** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnScopesScopeIdVirtualwiresPost**
> VdnScopesScopeIdVirtualwiresPost(ctx, scopeId, optional)
logicalSwitchCreate

Create a logical switch.  To create a universal logical switch use *universalvdnscope* as the scopeId in the URI and send the request to the primary NSX Manager. Request body parameters:   * **name** - Optional. The name of the logical switch.   * **description** - Optional. Description of the logical switch.   * **tenantId** - Required.   * **controlPlaneMode** - Optional. The control plane mode. If not     specified, the **controlPlaneMode** of the transport zone is used. It     can be one of the following:       * *UNICAST_MODE*       * *HYBRID_MODE*       * *MULTICAST_MODE*   * **guestVlanAllowed** - Optional. Default is *false*.   Parameters:  scopeId: A valid transport zone ID (vdnScope objectId).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***VirtualwiresApiVdnScopesScopeIdVirtualwiresPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiVdnScopesScopeIdVirtualwiresPostOpts struct
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

# **VdnVirtualwiresGet**
> VdnVirtualwiresGet(ctx, optional)
logicalSwitchesRead

Retrieve information about all logical switches in all transport zones.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualwiresApiVdnVirtualwiresGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiVdnVirtualwiresGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startindex** | **optional.String**|  | 
 **pagesize** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnVirtualwiresVirtualWireIDBackingPost**
> VdnVirtualwiresVirtualWireIDBackingPost(ctx, virtualWireID, optional)
logicalSwitchPortGroupFixAction

For every logical switch created, NSX creates a corresponding port group in vCenter. If the port group is missing, the logical switch will stop functioning.  If the port group backing a logical switch is deleted, you can recreate a new backing port group for the logical switch.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 
 **optional** | ***VirtualwiresApiVdnVirtualwiresVirtualWireIDBackingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiVdnVirtualwiresVirtualWireIDBackingPostOpts struct
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

# **VdnVirtualwiresVirtualWireIDConnCheckMulticastPost**
> VdnVirtualwiresVirtualWireIDConnCheckMulticastPost(ctx, virtualWireID, optional)
logicalSwitchConnCheckExecute

Test multicast group connectivity.  Test multicast group connectivity between two hosts connected to the specified logical switch.  Parameter **packetSizeMode** has one of the following values: * *0* - VXLAN standard packet size * *1* - minimum packet size * *2* - customized packet size. If you set **packetSizeMode** to *2*, you must specify the size using the **packetSize** parameter.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 
 **optional** | ***VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckMulticastPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckMulticastPostOpts struct
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

# **VdnVirtualwiresVirtualWireIDConnCheckP2pPost**
> VdnVirtualwiresVirtualWireIDConnCheckP2pPost(ctx, virtualWireID, optional)
logicalSwitchPingExecute

Test point-to-point connectivity.  Test point-to-point connectivity between two hosts connected to the specified logical switch.  Parameter **packetSizeMode** has one of the following values: * *0* - VXLAN standard packet size * *1* - minimum packet size * *2* - customized packet size. If you set **packetSizeMode** to *2*, you must specify the size using the **packetSize** parameter.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 
 **optional** | ***VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckP2pPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckP2pPostOpts struct
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

# **VdnVirtualwiresVirtualWireIDDelete**
> VdnVirtualwiresVirtualWireIDDelete(ctx, virtualWireID)
logicalSwitchDelete

Delete the specified logical switch.  Parameters:  virtualWireID: A logical switch id, e.g. virtualwire-1002  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnVirtualwiresVirtualWireIDGet**
> VdnVirtualwiresVirtualWireIDGet(ctx, virtualWireID)
logicalSwitchShow

Retrieve information about the specified logical switch.  If the switch is a universal logical switch the **isUniversal** parameter is set to true in the response body.   Parameters:  virtualWireID: A logical switch id, e.g. virtualwire-1002  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnVirtualwiresVirtualWireIDHardwaregatewaysGet**
> VdnVirtualwiresVirtualWireIDHardwaregatewaysGet(ctx, virtualWireID)
logicalSwitchHardwareGatewayBindingsList

Retrieve hardware gateway bindings for the specified logical switch.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPost**
> VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPost(ctx, virtualWireID, hardwareGatewayBindingId, optional)
logicalSwitchHardwareGatewayBindingCreate

Manage the connection between a hardware gateway and a logical switch.  ### Attach a hardware gateway to a logical switch and create a new binding with the information provided  `POST /api/2.0/vdn/virtualwires/{virtualwireid}/hardwaregateways`  ``` <hardwareGatewayBinding>   <hardwareGatewayId>hardwarewgateway1</hardwareGatewayId>   <vlan>v1</vlan>   <switchName>s1</switchName>   <portName>s1</portName> </hardwareGatewayBinding>  ```  ### Attach a hardware gateway to a logical switch, specifying an existing binding by ID  `POST /api/2.0/vdn/virtualwires/<virtualwireId>/hardwaregateways/{bindingId}?action=attach`  ``` <virtualWire>   ...   <hardwareGatewayBindings>     <hardwareGatewayBinding>       <id>binding id</id>     </hardwareGatewayBinding>   </hardwareGatewayBindings> </virtualWire> ```  ### Detach a hardware gateway from a logical switch  `POST /api/2.0/vdn/virtualwires/<virtualwireId>/hardwaregateways/{bindingId}?action=detach`  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  hardwareGatewayBindingId: Hardware Gateway Binding ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 
  **hardwareGatewayBindingId** | **string**|  | 
 **optional** | ***VirtualwiresApiVdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiVdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **action** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnVirtualwiresVirtualWireIDPut**
> VdnVirtualwiresVirtualWireIDPut(ctx, virtualWireID, optional)
logicalSwitchUpdate

Update the specified logical switch.  For example, you can update the name, description, or control plane mode.   Parameters:  virtualWireID: A logical switch id, e.g. virtualwire-1002  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 
 **optional** | ***VirtualwiresApiVdnVirtualwiresVirtualWireIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiVdnVirtualwiresVirtualWireIDPutOpts struct
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

# **VdnVirtualwiresVmVnicPost**
> VdnVirtualwiresVmVnicPost(ctx, optional)
logicalSwitchVmAttachCreate

Attach a VM vNIC to, or detach a VM vNIC from a logical switch.  Specify the logical switch ID in the **portgroupId** parameter. To detach a VM vNIC from a logical switch, leave the **portgroupId** parameter empty.  To find the ID of a VM vNIC, do the following: 1. In the vSphere MOB, navigate to the VM you want to connect or disconnect. 2. Click **config** and take note of the **instanceUuid**. 3. Click **hardware** and take note of the last three digits of the appropriate network interface device.  Use these two values to form the VM vNIC ID.  For example, if the **instanceUuid** is *502e71fa-1a00-759b-e40f-ce778e915f16* and the appropriate **device** value is *device[4000]*, the **objectId** and **vnicUuid** are both *502e71fa-1a00-759b-e40f-ce778e915f16.000*.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualwiresApiVdnVirtualwiresVmVnicPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiVdnVirtualwiresVmVnicPostOpts struct
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

