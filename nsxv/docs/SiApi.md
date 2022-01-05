# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20SiAgentAgentIDGet**](SiApi.md#Api20SiAgentAgentIDGet) | **Get** /api/2.0/si/agent/{agentID} | agentInformationRead
[**Api20SiDeploymentDeploymentunitIDAgentsGet**](SiApi.md#Api20SiDeploymentDeploymentunitIDAgentsGet) | **Get** /api/2.0/si/deployment/{deploymentunitID}/agents | deploymentAgentsRead
[**Api20SiFabricSyncConflictsGet**](SiApi.md#Api20SiFabricSyncConflictsGet) | **Get** /api/2.0/si/fabric/sync/conflicts | agentConflictsRead
[**Api20SiFabricSyncConflictsPut**](SiApi.md#Api20SiFabricSyncConflictsPut) | **Put** /api/2.0/si/fabric/sync/conflicts | agentConflictsUpdate
[**Api20SiHostHostIDAgentsGet**](SiApi.md#Api20SiHostHostIDAgentsGet) | **Get** /api/2.0/si/host/{hostID}/agents | hostAgentsRead

# **Api20SiAgentAgentIDGet**
> Api20SiAgentAgentIDGet(ctx, agentID)
agentInformationRead

Retrieve agent (host components and appliances) details.   Parameters:  agentID: Specified agent  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20SiDeploymentDeploymentunitIDAgentsGet**
> Api20SiDeploymentDeploymentunitIDAgentsGet(ctx, deploymentunitID)
deploymentAgentsRead

Retrieve all agents for the specified deployment.  Parameters:  deploymentunitID: Specified deployment.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentunitID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20SiFabricSyncConflictsGet**
> Api20SiFabricSyncConflictsGet(ctx, )
agentConflictsRead

Retrieve conflicting deployment units and EAM agencies, if any, and the allowed operations on them.   Parameters:  

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

# **Api20SiFabricSyncConflictsPut**
> Api20SiFabricSyncConflictsPut(ctx, optional)
agentConflictsUpdate

Create deployment units for conflicting EAM Agencies, delete conflicting EAM agencies, or delete deployment units for conflicting EAM agencies.  ### Create deployment units for conflicting EAM agencies  ``` <conflictResolverInfo>   <agencyAction>RESTORE</agencyAction> </conflictResolverInfo> ```  ### Delete conflicting EAM agencies  ``` <conflictResolverInfo>   <agencyAction>DELETE</agencyAction> </conflictResolverInfo> ```  ### Delete deployment units for conflicting EAM agencies  ``` <conflictResolverInfo>   <deploymentUnitAction>DELETE</deploymentUnitAction> </conflictResolverInfo> ```   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SiApiApi20SiFabricSyncConflictsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SiApiApi20SiFabricSyncConflictsPutOpts struct
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

# **Api20SiHostHostIDAgentsGet**
> Api20SiHostHostIDAgentsGet(ctx, hostID)
hostAgentsRead

Retrieves all agents on the specified host. The response body contains agent IDs for each agent, which you can use to retrieve details about that agent.   Parameters:  hostID: Specified host  

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

