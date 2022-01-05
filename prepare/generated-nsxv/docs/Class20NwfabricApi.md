# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20NwfabricClustersClusterIDDelete**](Class20NwfabricApi.md#Api20NwfabricClustersClusterIDDelete) | **Delete** /api/2.0/nwfabric/clusters/{clusterID} | nwfabricClustersDelete
[**Api20NwfabricClustersClusterIDGet**](Class20NwfabricApi.md#Api20NwfabricClustersClusterIDGet) | **Get** /api/2.0/nwfabric/clusters/{clusterID} | nwfabricClustersRead
[**Api20NwfabricClustersClusterIDPut**](Class20NwfabricApi.md#Api20NwfabricClustersClusterIDPut) | **Put** /api/2.0/nwfabric/clusters/{clusterID} | nwfabricClustersUpdate
[**Api20NwfabricConfigureDelete**](Class20NwfabricApi.md#Api20NwfabricConfigureDelete) | **Delete** /api/2.0/nwfabric/configure | nwfabricConfigDelete
[**Api20NwfabricConfigurePost**](Class20NwfabricApi.md#Api20NwfabricConfigurePost) | **Post** /api/2.0/nwfabric/configure | nwfabricConfigCreate
[**Api20NwfabricConfigurePut**](Class20NwfabricApi.md#Api20NwfabricConfigurePut) | **Put** /api/2.0/nwfabric/configure | nwfabricConfigUpdate
[**Api20NwfabricFeaturesGet**](Class20NwfabricApi.md#Api20NwfabricFeaturesGet) | **Get** /api/2.0/nwfabric/features | nwfabricFeaturesList
[**Api20NwfabricHostsHostIDDelete**](Class20NwfabricApi.md#Api20NwfabricHostsHostIDDelete) | **Delete** /api/2.0/nwfabric/hosts/{hostID} | nwfabricHostsDelete
[**Api20NwfabricHostsHostIDGet**](Class20NwfabricApi.md#Api20NwfabricHostsHostIDGet) | **Get** /api/2.0/nwfabric/hosts/{hostID} | nwfabricHostsRead
[**Api20NwfabricHostsHostIDPut**](Class20NwfabricApi.md#Api20NwfabricHostsHostIDPut) | **Put** /api/2.0/nwfabric/hosts/{hostID} | nwfabricHostsUpdate
[**Api20NwfabricStatusAlleligibleResourceTypeGet**](Class20NwfabricApi.md#Api20NwfabricStatusAlleligibleResourceTypeGet) | **Get** /api/2.0/nwfabric/status/alleligible/{resourceType} | statusResourceTypeRead
[**Api20NwfabricStatusChildParentResourceIDGet**](Class20NwfabricApi.md#Api20NwfabricStatusChildParentResourceIDGet) | **Get** /api/2.0/nwfabric/status/child/{parentResourceID} | childStatusRead
[**Api20NwfabricStatusGet**](Class20NwfabricApi.md#Api20NwfabricStatusGet) | **Get** /api/2.0/nwfabric/status | nwfabricStatusRead

# **Api20NwfabricClustersClusterIDDelete**
> Api20NwfabricClustersClusterIDDelete(ctx, clusterID)
nwfabricClustersDelete

Delete locale ID for the specified cluster.  Parameters:  clusterID: Cluster ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20NwfabricClustersClusterIDGet**
> Api20NwfabricClustersClusterIDGet(ctx, clusterID)
nwfabricClustersRead

Retrieve the locale ID for the specified cluster.  Parameters:  clusterID: Cluster ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20NwfabricClustersClusterIDPut**
> Api20NwfabricClustersClusterIDPut(ctx, clusterID, optional)
nwfabricClustersUpdate

Update the locale ID for the specified cluster.  Parameters:  clusterID: Cluster ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterID** | **string**|  | 
 **optional** | ***Class20NwfabricApiApi20NwfabricClustersClusterIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20NwfabricApiApi20NwfabricClustersClusterIDPutOpts struct
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

# **Api20NwfabricConfigureDelete**
> Api20NwfabricConfigureDelete(ctx, optional)
nwfabricConfigDelete

Remove VXLAN or network virtualization components.  Removing network virtualization components removes previously installed VIBs, tears down NSX Manager to ESXi messaging, and removes any other network fabric dependent features such as logical switches. If a feature such as logical switches is being used in your environment, this call fails.  Removing VXLAN does not remove the network virtualization components from the cluster.  | Name | Comments | |------|----------| |**resourceId** | vCenter MOB ID of cluster. For example, domain-7.| |**featureId** | Feature to act upon. Omit for network virtualization components operations. Use *com.vmware.vshield.vsm.vxlan* for VXLAN operations.|  ### Remove Network Virtualization Components  ``` <nwFabricFeatureConfig>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>   </resourceConfig> </nwFabricFeatureConfig> ```  ### Remove VXLAN  ``` <nwFabricFeatureConfig>   <featureId>com.vmware.vshield.vsm.vxlan</featureId>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>    </resourceConfig> </nwFabricFeatureConfig> ```  ### Remove VXLAN with vDS context  ``` <nwFabricFeatureConfig>   <featureId>com.vmware.vshield.vsm.vxlan</featureId>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>     <configSpec class=\"map\">       <entry>         <keyclass=\"java.lang.String\">vxlan</key>         <valueclass=\"java.lang.String\">cascadeDeleteVdsContext</value>       </entry>     </configSpec>   </resourceConfig> </nwFabricFeatureConfig> ```   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20NwfabricApiApi20NwfabricConfigureDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20NwfabricApiApi20NwfabricConfigureDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20NwfabricConfigurePost**
> Api20NwfabricConfigurePost(ctx, optional)
nwfabricConfigCreate

Install network fabric or VXLAN.  This method can be used to perform the following tasks:  * Install Network Virtualization Components * Configure VXLAN * Configure VXLAN with LACPv2 * Reset Communication Between NSX Manager and a Host or Cluster  **Parameter Information**  | Name | Comments | |------|----------| |**resourceId** | vCenter MOB ID of cluster. For example, *domain-7*. A host can be specified when resetting communication. For example, *host-24*. | |**featureId** | Feature to act upon. Omit for network virtualization components operations. Use *com.vmware.vshield.vsm.vxlan* for VXLAN operations, *com.vmware.vshield.vsm.messagingInfra* for message bus operations.| |**ipPoolId** | Used for VXLAN installation. If not specified, DHCP is used for VTEP address assignment.| |**teaming** | Used for VXLAN installation. Options are *FAILOVER_ORDER*, *ETHER_CHANNEL*, *LACP_ACTIVE*, *LACP_PASSIVE*, *LOADBALANCE_LOADBASED*, *LOADBALANCE_SRCID*, *LOADBALANCE_SRCMAC*, *LACP_V2*| |**uplinkPortName** | The *uplinkPortName* as specified in vCenter.|  ### Install Network Virtualization Components  `POST /api/2.0/nwfabric/configure`  ``` <nwFabricFeatureConfig>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>   </resourceConfig> </nwFabricFeatureConfig> ```  ### Configure VXLAN  `POST /api/2.0/nwfabric/configure`  ``` <nwFabricFeatureConfig>   <featureId>com.vmware.vshield.vsm.vxlan</featureId>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>     <configSpec class=\"clusterMappingSpec\">       <switch>         <objectId>DVS MOID</objectId></switch>         <vlanId>0</vlanId>         <vmknicCount>1</vmknicCount>         <ipPoolId>IPADDRESSPOOL ID</ipPoolId>     </configSpec>   </resourceConfig>   <resourceConfig>     <resourceId>DVS MOID</resourceId>     <configSpec class=\"vdsContext\">       <switch>           <objectId>DVS MOID</objectId>       </switch>       <mtu>1600</mtu>       <teaming>ETHER_CHANNEL</teaming>     </configSpec>   </resourceConfig> </nwFabricFeatureConfig> ```  ### Configure VXLAN with LACPv2  `POST /api/2.0/nwfabric/configure`  ``` <nwFabricFeatureConfig>   <featureId>com.vmware.vshield.nsxmgr.vxlan</featureId>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>     <configSpec class=\"clusterMappingSpec\">       <switch>         <objectId>DVS MOID</objectId>       </switch>       <vlanId>0</vlanId>       <vmknicCount>1</vmknicCount>     </configSpec>   </resourceConfig>   <resourceConfig>     <resourceId>DVS MOID</resourceId>     <configSpec class=\"vdsContext\">       <switch>         <objectId>DVS MOID</objectId>       </switch>       <mtu>1600</mtu>       <teaming>LACP_V2</teaming>       <uplinkPortName>LAG NAME</uplinkPortName>     </configSpec>   </resourceConfig> </nwFabricFeatureConfig> ```  ### Reset Communication Between NSX Manager and a Host or Cluster  `POST /api/2.0/nwfabric/configure?action=synchronize`  ```  <nwFabricFeatureConfig>   <featureId>com.vmware.vshield.vsm.messagingInfra</featureId>   <resourceConfig>     <resourceId>resourceId</resourceId>   </resourceConfig> </nwFabricFeatureConfig>  ```   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20NwfabricApiApi20NwfabricConfigurePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20NwfabricApiApi20NwfabricConfigurePostOpts struct
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

# **Api20NwfabricConfigurePut**
> Api20NwfabricConfigurePut(ctx, optional)
nwfabricConfigUpdate

Upgrade Network virtualization components. _ This API call can be used to upgrade network virtualization components. After NSX Manager is upgraded, previously prepared clusters must have the 6.x network virtualization components installed.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20NwfabricApiApi20NwfabricConfigurePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20NwfabricApiApi20NwfabricConfigurePutOpts struct
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

# **Api20NwfabricFeaturesGet**
> Api20NwfabricFeaturesGet(ctx, )
nwfabricFeaturesList

Retrieves all network fabric features available on the cluster. Multiple **featureInfo** sections may be returned.   Parameters:  

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

# **Api20NwfabricHostsHostIDDelete**
> Api20NwfabricHostsHostIDDelete(ctx, hostID)
nwfabricHostsDelete

Delete the locale ID for the specified host.  Parameters:  hostID: Host ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20NwfabricHostsHostIDGet**
> Api20NwfabricHostsHostIDGet(ctx, hostID)
nwfabricHostsRead

Retrieve the locale ID for the specified host.  Parameters:  hostID: Host ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20NwfabricHostsHostIDPut**
> Api20NwfabricHostsHostIDPut(ctx, hostID, optional)
nwfabricHostsUpdate

Update the locale ID for the specified host.  Parameters:  hostID: Host ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostID** | **string**|  | 
 **optional** | ***Class20NwfabricApiApi20NwfabricHostsHostIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20NwfabricApiApi20NwfabricHostsHostIDPutOpts struct
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

# **Api20NwfabricStatusAlleligibleResourceTypeGet**
> Api20NwfabricStatusAlleligibleResourceTypeGet(ctx, resourceType)
statusResourceTypeRead

Retrieve status of resources by criterion.   Parameters:  resourceType: Valid resource type  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20NwfabricStatusChildParentResourceIDGet**
> Api20NwfabricStatusChildParentResourceIDGet(ctx, parentResourceID)
childStatusRead

Retrieve the network fabric status of child resources of the specified resource.   Parameters:  parentResourceID: Parent resource ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentResourceID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20NwfabricStatusGet**
> Api20NwfabricStatusGet(ctx, optional)
nwfabricStatusRead

Retrieve the network fabric status of the specified resource.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20NwfabricApiApi20NwfabricStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20NwfabricApiApi20NwfabricStatusGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resource** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

