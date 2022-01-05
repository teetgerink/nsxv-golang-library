# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20VdnScopesScopeIdVirtualwiresGet**](VirtualwiresApi.md#Api20VdnScopesScopeIdVirtualwiresGet) | **Get** /api/2.0/vdn/scopes/{scopeId}/virtualwires | logicalSwitchList
[**Api20VdnScopesScopeIdVirtualwiresPost**](VirtualwiresApi.md#Api20VdnScopesScopeIdVirtualwiresPost) | **Post** /api/2.0/vdn/scopes/{scopeId}/virtualwires | logicalSwitchCreate
[**Api20VdnVirtualwiresGet**](VirtualwiresApi.md#Api20VdnVirtualwiresGet) | **Get** /api/2.0/vdn/virtualwires | logicalSwitchesRead
[**Api20VdnVirtualwiresVirtualWireIDBackingPost**](VirtualwiresApi.md#Api20VdnVirtualwiresVirtualWireIDBackingPost) | **Post** /api/2.0/vdn/virtualwires/{virtualWireID}/backing | logicalSwitchPortGroupFixAction
[**Api20VdnVirtualwiresVirtualWireIDConnCheckMulticastPost**](VirtualwiresApi.md#Api20VdnVirtualwiresVirtualWireIDConnCheckMulticastPost) | **Post** /api/2.0/vdn/virtualwires/{virtualWireID}/conn-check/multicast | logicalSwitchConnCheckExecute
[**Api20VdnVirtualwiresVirtualWireIDConnCheckP2pPost**](VirtualwiresApi.md#Api20VdnVirtualwiresVirtualWireIDConnCheckP2pPost) | **Post** /api/2.0/vdn/virtualwires/{virtualWireID}/conn-check/p2p | logicalSwitchPingExecute
[**Api20VdnVirtualwiresVirtualWireIDDelete**](VirtualwiresApi.md#Api20VdnVirtualwiresVirtualWireIDDelete) | **Delete** /api/2.0/vdn/virtualwires/{virtualWireID} | logicalSwitchDelete
[**Api20VdnVirtualwiresVirtualWireIDGet**](VirtualwiresApi.md#Api20VdnVirtualwiresVirtualWireIDGet) | **Get** /api/2.0/vdn/virtualwires/{virtualWireID} | logicalSwitchShow
[**Api20VdnVirtualwiresVirtualWireIDHardwaregatewaysGet**](VirtualwiresApi.md#Api20VdnVirtualwiresVirtualWireIDHardwaregatewaysGet) | **Get** /api/2.0/vdn/virtualwires/{virtualWireID}/hardwaregateways | logicalSwitchHardwareGatewayBindingsList
[**Api20VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPost**](VirtualwiresApi.md#Api20VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPost) | **Post** /api/2.0/vdn/virtualwires/{virtualWireID}/hardwaregateways/{hardwareGatewayBindingId} | logicalSwitchHardwareGatewayBindingCreate
[**Api20VdnVirtualwiresVirtualWireIDPut**](VirtualwiresApi.md#Api20VdnVirtualwiresVirtualWireIDPut) | **Put** /api/2.0/vdn/virtualwires/{virtualWireID} | logicalSwitchUpdate
[**Api20VdnVirtualwiresVmVnicPost**](VirtualwiresApi.md#Api20VdnVirtualwiresVmVnicPost) | **Post** /api/2.0/vdn/virtualwires/vm/vnic | logicalSwitchVmAttachCreate

# **Api20VdnScopesScopeIdVirtualwiresGet**
> Api20VdnScopesScopeIdVirtualwiresGet(ctx, scopeId, optional)
logicalSwitchList

Retrieve information about all logical switches in the specified transport zone (network scope).   Parameters:  scopeId: A valid transport zone ID (vdnScope objectId).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***VirtualwiresApiApi20VdnScopesScopeIdVirtualwiresGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiApi20VdnScopesScopeIdVirtualwiresGetOpts struct
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

# **Api20VdnScopesScopeIdVirtualwiresPost**
> Api20VdnScopesScopeIdVirtualwiresPost(ctx, scopeId, optional)
logicalSwitchCreate

Create a logical switch.  To create a universal logical switch use *universalvdnscope* as the scopeId in the URI and send the request to the primary NSX Manager. Request body parameters:   * **name** - Optional. The name of the logical switch.   * **description** - Optional. Description of the logical switch.   * **tenantId** - Required.   * **controlPlaneMode** - Optional. The control plane mode. If not     specified, the **controlPlaneMode** of the transport zone is used. It     can be one of the following:       * *UNICAST_MODE*       * *HYBRID_MODE*       * *MULTICAST_MODE*   * **guestVlanAllowed** - Optional. Default is *false*.   Parameters:  scopeId: A valid transport zone ID (vdnScope objectId).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***VirtualwiresApiApi20VdnScopesScopeIdVirtualwiresPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiApi20VdnScopesScopeIdVirtualwiresPostOpts struct
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

# **Api20VdnVirtualwiresGet**
> Api20VdnVirtualwiresGet(ctx, optional)
logicalSwitchesRead

Retrieve information about all logical switches in all transport zones.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualwiresApiApi20VdnVirtualwiresGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiApi20VdnVirtualwiresGetOpts struct
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

# **Api20VdnVirtualwiresVirtualWireIDBackingPost**
> Api20VdnVirtualwiresVirtualWireIDBackingPost(ctx, virtualWireID, optional)
logicalSwitchPortGroupFixAction

For every logical switch created, NSX creates a corresponding port group in vCenter. If the port group is missing, the logical switch will stop functioning.  If the port group backing a logical switch is deleted, you can recreate a new backing port group for the logical switch.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 
 **optional** | ***VirtualwiresApiApi20VdnVirtualwiresVirtualWireIDBackingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiApi20VdnVirtualwiresVirtualWireIDBackingPostOpts struct
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

# **Api20VdnVirtualwiresVirtualWireIDConnCheckMulticastPost**
> Api20VdnVirtualwiresVirtualWireIDConnCheckMulticastPost(ctx, virtualWireID, optional)
logicalSwitchConnCheckExecute

Test multicast group connectivity.  Test multicast group connectivity between two hosts connected to the specified logical switch.  Parameter **packetSizeMode** has one of the following values: * *0* - VXLAN standard packet size * *1* - minimum packet size * *2* - customized packet size. If you set **packetSizeMode** to *2*, you must specify the size using the **packetSize** parameter.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 
 **optional** | ***VirtualwiresApiApi20VdnVirtualwiresVirtualWireIDConnCheckMulticastPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiApi20VdnVirtualwiresVirtualWireIDConnCheckMulticastPostOpts struct
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

# **Api20VdnVirtualwiresVirtualWireIDConnCheckP2pPost**
> Api20VdnVirtualwiresVirtualWireIDConnCheckP2pPost(ctx, virtualWireID, optional)
logicalSwitchPingExecute

Test point-to-point connectivity.  Test point-to-point connectivity between two hosts connected to the specified logical switch.  Parameter **packetSizeMode** has one of the following values: * *0* - VXLAN standard packet size * *1* - minimum packet size * *2* - customized packet size. If you set **packetSizeMode** to *2*, you must specify the size using the **packetSize** parameter.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 
 **optional** | ***VirtualwiresApiApi20VdnVirtualwiresVirtualWireIDConnCheckP2pPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiApi20VdnVirtualwiresVirtualWireIDConnCheckP2pPostOpts struct
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

# **Api20VdnVirtualwiresVirtualWireIDDelete**
> Api20VdnVirtualwiresVirtualWireIDDelete(ctx, virtualWireID)
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

# **Api20VdnVirtualwiresVirtualWireIDGet**
> Api20VdnVirtualwiresVirtualWireIDGet(ctx, virtualWireID)
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

# **Api20VdnVirtualwiresVirtualWireIDHardwaregatewaysGet**
> Api20VdnVirtualwiresVirtualWireIDHardwaregatewaysGet(ctx, virtualWireID)
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

# **Api20VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPost**
> Api20VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPost(ctx, virtualWireID, hardwareGatewayBindingId, optional)
logicalSwitchHardwareGatewayBindingCreate

Manage the connection between a hardware gateway and a logical switch.  ### Attach a hardware gateway to a logical switch and create a new binding with the information provided  `POST /api/2.0/vdn/virtualwires/{virtualwireid}/hardwaregateways`  ``` <hardwareGatewayBinding>   <hardwareGatewayId>hardwarewgateway1</hardwareGatewayId>   <vlan>v1</vlan>   <switchName>s1</switchName>   <portName>s1</portName> </hardwareGatewayBinding>  ```  ### Attach a hardware gateway to a logical switch, specifying an existing binding by ID  `POST /api/2.0/vdn/virtualwires/<virtualwireId>/hardwaregateways/{bindingId}?action=attach`  ``` <virtualWire>   ...   <hardwareGatewayBindings>     <hardwareGatewayBinding>       <id>binding id</id>     </hardwareGatewayBinding>   </hardwareGatewayBindings> </virtualWire> ```  ### Detach a hardware gateway from a logical switch  `POST /api/2.0/vdn/virtualwires/<virtualwireId>/hardwaregateways/{bindingId}?action=detach`  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  hardwareGatewayBindingId: Hardware Gateway Binding ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 
  **hardwareGatewayBindingId** | **string**|  | 
 **optional** | ***VirtualwiresApiApi20VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiApi20VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPostOpts struct
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

# **Api20VdnVirtualwiresVirtualWireIDPut**
> Api20VdnVirtualwiresVirtualWireIDPut(ctx, virtualWireID, optional)
logicalSwitchUpdate

Update the specified logical switch.  For example, you can update the name, description, or control plane mode.   Parameters:  virtualWireID: A logical switch id, e.g. virtualwire-1002  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualWireID** | **string**|  | 
 **optional** | ***VirtualwiresApiApi20VdnVirtualwiresVirtualWireIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiApi20VdnVirtualwiresVirtualWireIDPutOpts struct
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

# **Api20VdnVirtualwiresVmVnicPost**
> Api20VdnVirtualwiresVmVnicPost(ctx, optional)
logicalSwitchVmAttachCreate

Attach a VM vNIC to, or detach a VM vNIC from a logical switch.  Specify the logical switch ID in the **portgroupId** parameter. To detach a VM vNIC from a logical switch, leave the **portgroupId** parameter empty.  To find the ID of a VM vNIC, do the following: 1. In the vSphere MOB, navigate to the VM you want to connect or disconnect. 2. Click **config** and take note of the **instanceUuid**. 3. Click **hardware** and take note of the last three digits of the appropriate network interface device.  Use these two values to form the VM vNIC ID.  For example, if the **instanceUuid** is *502e71fa-1a00-759b-e40f-ce778e915f16* and the appropriate **device** value is *device[4000]*, the **objectId** and **vnicUuid** are both *502e71fa-1a00-759b-e40f-ce778e915f16.000*.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualwiresApiApi20VdnVirtualwiresVmVnicPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualwiresApiApi20VdnVirtualwiresVmVnicPostOpts struct
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

