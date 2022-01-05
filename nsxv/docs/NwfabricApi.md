# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NwfabricClustersClusterIDDelete**](NwfabricApi.md#NwfabricClustersClusterIDDelete) | **Delete** /api/2.0/nwfabric/clusters/{clusterID} | nwfabricClustersDelete
[**NwfabricClustersClusterIDGet**](NwfabricApi.md#NwfabricClustersClusterIDGet) | **Get** /api/2.0/nwfabric/clusters/{clusterID} | nwfabricClustersRead
[**NwfabricClustersClusterIDPut**](NwfabricApi.md#NwfabricClustersClusterIDPut) | **Put** /api/2.0/nwfabric/clusters/{clusterID} | nwfabricClustersUpdate
[**NwfabricConfigureDelete**](NwfabricApi.md#NwfabricConfigureDelete) | **Delete** /api/2.0/nwfabric/configure | nwfabricConfigDelete
[**NwfabricConfigurePost**](NwfabricApi.md#NwfabricConfigurePost) | **Post** /api/2.0/nwfabric/configure | nwfabricConfigCreate
[**NwfabricConfigurePut**](NwfabricApi.md#NwfabricConfigurePut) | **Put** /api/2.0/nwfabric/configure | nwfabricConfigUpdate
[**NwfabricFeaturesGet**](NwfabricApi.md#NwfabricFeaturesGet) | **Get** /api/2.0/nwfabric/features | nwfabricFeaturesList
[**NwfabricHostsHostIDDelete**](NwfabricApi.md#NwfabricHostsHostIDDelete) | **Delete** /api/2.0/nwfabric/hosts/{hostID} | nwfabricHostsDelete
[**NwfabricHostsHostIDGet**](NwfabricApi.md#NwfabricHostsHostIDGet) | **Get** /api/2.0/nwfabric/hosts/{hostID} | nwfabricHostsRead
[**NwfabricHostsHostIDPut**](NwfabricApi.md#NwfabricHostsHostIDPut) | **Put** /api/2.0/nwfabric/hosts/{hostID} | nwfabricHostsUpdate
[**NwfabricStatusAlleligibleResourceTypeGet**](NwfabricApi.md#NwfabricStatusAlleligibleResourceTypeGet) | **Get** /api/2.0/nwfabric/status/alleligible/{resourceType} | statusResourceTypeRead
[**NwfabricStatusChildParentResourceIDGet**](NwfabricApi.md#NwfabricStatusChildParentResourceIDGet) | **Get** /api/2.0/nwfabric/status/child/{parentResourceID} | childStatusRead
[**NwfabricStatusGet**](NwfabricApi.md#NwfabricStatusGet) | **Get** /api/2.0/nwfabric/status | nwfabricStatusRead

# **NwfabricClustersClusterIDDelete**
> NwfabricClustersClusterIDDelete(ctx, clusterID)
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

# **NwfabricClustersClusterIDGet**
> NwfabricClustersClusterIDGet(ctx, clusterID)
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

# **NwfabricClustersClusterIDPut**
> NwfabricClustersClusterIDPut(ctx, clusterID, optional)
nwfabricClustersUpdate

Update the locale ID for the specified cluster.  Parameters:  clusterID: Cluster ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterID** | **string**|  | 
 **optional** | ***NwfabricApiNwfabricClustersClusterIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NwfabricApiNwfabricClustersClusterIDPutOpts struct
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

# **NwfabricConfigureDelete**
> NwfabricConfigureDelete(ctx, optional)
nwfabricConfigDelete

Remove VXLAN or network virtualization components.  Removing network virtualization components removes previously installed VIBs, tears down NSX Manager to ESXi messaging, and removes any other network fabric dependent features such as logical switches. If a feature such as logical switches is being used in your environment, this call fails.  Removing VXLAN does not remove the network virtualization components from the cluster.  | Name | Comments | |------|----------| |**resourceId** | vCenter MOB ID of cluster. For example, domain-7.| |**featureId** | Feature to act upon. Omit for network virtualization components operations. Use *com.vmware.vshield.vsm.vxlan* for VXLAN operations.|  ### Remove Network Virtualization Components  ``` <nwFabricFeatureConfig>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>   </resourceConfig> </nwFabricFeatureConfig> ```  ### Remove VXLAN  ``` <nwFabricFeatureConfig>   <featureId>com.vmware.vshield.vsm.vxlan</featureId>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>    </resourceConfig> </nwFabricFeatureConfig> ```  ### Remove VXLAN with vDS context  ``` <nwFabricFeatureConfig>   <featureId>com.vmware.vshield.vsm.vxlan</featureId>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>     <configSpec class=\"map\">       <entry>         <keyclass=\"java.lang.String\">vxlan</key>         <valueclass=\"java.lang.String\">cascadeDeleteVdsContext</value>       </entry>     </configSpec>   </resourceConfig> </nwFabricFeatureConfig> ```   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NwfabricApiNwfabricConfigureDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NwfabricApiNwfabricConfigureDeleteOpts struct
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

# **NwfabricConfigurePost**
> NwfabricConfigurePost(ctx, optional)
nwfabricConfigCreate

Install network fabric or VXLAN.  This method can be used to perform the following tasks:  * Install Network Virtualization Components * Configure VXLAN * Configure VXLAN with LACPv2 * Reset Communication Between NSX Manager and a Host or Cluster  **Parameter Information**  | Name | Comments | |------|----------| |**resourceId** | vCenter MOB ID of cluster. For example, *domain-7*. A host can be specified when resetting communication. For example, *host-24*. | |**featureId** | Feature to act upon. Omit for network virtualization components operations. Use *com.vmware.vshield.vsm.vxlan* for VXLAN operations, *com.vmware.vshield.vsm.messagingInfra* for message bus operations.| |**ipPoolId** | Used for VXLAN installation. If not specified, DHCP is used for VTEP address assignment.| |**teaming** | Used for VXLAN installation. Options are *FAILOVER_ORDER*, *ETHER_CHANNEL*, *LACP_ACTIVE*, *LACP_PASSIVE*, *LOADBALANCE_LOADBASED*, *LOADBALANCE_SRCID*, *LOADBALANCE_SRCMAC*, *LACP_V2*| |**uplinkPortName** | The *uplinkPortName* as specified in vCenter.|  ### Install Network Virtualization Components  `POST /api/2.0/nwfabric/configure`  ``` <nwFabricFeatureConfig>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>   </resourceConfig> </nwFabricFeatureConfig> ```  ### Configure VXLAN  `POST /api/2.0/nwfabric/configure`  ``` <nwFabricFeatureConfig>   <featureId>com.vmware.vshield.vsm.vxlan</featureId>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>     <configSpec class=\"clusterMappingSpec\">       <switch>         <objectId>DVS MOID</objectId></switch>         <vlanId>0</vlanId>         <vmknicCount>1</vmknicCount>         <ipPoolId>IPADDRESSPOOL ID</ipPoolId>     </configSpec>   </resourceConfig>   <resourceConfig>     <resourceId>DVS MOID</resourceId>     <configSpec class=\"vdsContext\">       <switch>           <objectId>DVS MOID</objectId>       </switch>       <mtu>1600</mtu>       <teaming>ETHER_CHANNEL</teaming>     </configSpec>   </resourceConfig> </nwFabricFeatureConfig> ```  ### Configure VXLAN with LACPv2  `POST /api/2.0/nwfabric/configure`  ``` <nwFabricFeatureConfig>   <featureId>com.vmware.vshield.nsxmgr.vxlan</featureId>   <resourceConfig>     <resourceId>CLUSTER MOID</resourceId>     <configSpec class=\"clusterMappingSpec\">       <switch>         <objectId>DVS MOID</objectId>       </switch>       <vlanId>0</vlanId>       <vmknicCount>1</vmknicCount>     </configSpec>   </resourceConfig>   <resourceConfig>     <resourceId>DVS MOID</resourceId>     <configSpec class=\"vdsContext\">       <switch>         <objectId>DVS MOID</objectId>       </switch>       <mtu>1600</mtu>       <teaming>LACP_V2</teaming>       <uplinkPortName>LAG NAME</uplinkPortName>     </configSpec>   </resourceConfig> </nwFabricFeatureConfig> ```  ### Reset Communication Between NSX Manager and a Host or Cluster  `POST /api/2.0/nwfabric/configure?action=synchronize`  ```  <nwFabricFeatureConfig>   <featureId>com.vmware.vshield.vsm.messagingInfra</featureId>   <resourceConfig>     <resourceId>resourceId</resourceId>   </resourceConfig> </nwFabricFeatureConfig>  ```   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NwfabricApiNwfabricConfigurePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NwfabricApiNwfabricConfigurePostOpts struct
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

# **NwfabricConfigurePut**
> NwfabricConfigurePut(ctx, optional)
nwfabricConfigUpdate

Upgrade Network virtualization components. _ This API call can be used to upgrade network virtualization components. After NSX Manager is upgraded, previously prepared clusters must have the 6.x network virtualization components installed.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NwfabricApiNwfabricConfigurePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NwfabricApiNwfabricConfigurePutOpts struct
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

# **NwfabricFeaturesGet**
> NwfabricFeaturesGet(ctx, )
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

# **NwfabricHostsHostIDDelete**
> NwfabricHostsHostIDDelete(ctx, hostID)
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

# **NwfabricHostsHostIDGet**
> NwfabricHostsHostIDGet(ctx, hostID)
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

# **NwfabricHostsHostIDPut**
> NwfabricHostsHostIDPut(ctx, hostID, optional)
nwfabricHostsUpdate

Update the locale ID for the specified host.  Parameters:  hostID: Host ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostID** | **string**|  | 
 **optional** | ***NwfabricApiNwfabricHostsHostIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NwfabricApiNwfabricHostsHostIDPutOpts struct
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

# **NwfabricStatusAlleligibleResourceTypeGet**
> NwfabricStatusAlleligibleResourceTypeGet(ctx, resourceType)
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

# **NwfabricStatusChildParentResourceIDGet**
> NwfabricStatusChildParentResourceIDGet(ctx, parentResourceID)
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

# **NwfabricStatusGet**
> NwfabricStatusGet(ctx, optional)
nwfabricStatusRead

Retrieve the network fabric status of the specified resource.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NwfabricApiNwfabricStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NwfabricApiNwfabricStatusGetOpts struct
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

