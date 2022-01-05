# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VdnControllerClusterGet**](ControllerApi.md#VdnControllerClusterGet) | **Get** /api/2.0/vdn/controller/cluster | nsxControllerClusterRead
[**VdnControllerClusterPut**](ControllerApi.md#VdnControllerClusterPut) | **Put** /api/2.0/vdn/controller/cluster | nsxControllerClusterUpdate
[**VdnControllerControllerIdDelete**](ControllerApi.md#VdnControllerControllerIdDelete) | **Delete** /api/2.0/vdn/controller/{controllerId} | nsxControllerDelete
[**VdnControllerControllerIdPost**](ControllerApi.md#VdnControllerControllerIdPost) | **Post** /api/2.0/vdn/controller/{controllerId} | nsxControllerRemediateAction
[**VdnControllerControllerIdSnapshotGet**](ControllerApi.md#VdnControllerControllerIdSnapshotGet) | **Get** /api/2.0/vdn/controller/{controllerId}/snapshot | nsxControllerSnapshotRead
[**VdnControllerControllerIdSyslogDelete**](ControllerApi.md#VdnControllerControllerIdSyslogDelete) | **Delete** /api/2.0/vdn/controller/{controllerId}/syslog | nsxControllerSyslogDelete
[**VdnControllerControllerIdSyslogGet**](ControllerApi.md#VdnControllerControllerIdSyslogGet) | **Get** /api/2.0/vdn/controller/{controllerId}/syslog | nsxControllerSyslogRead
[**VdnControllerControllerIdSyslogPost**](ControllerApi.md#VdnControllerControllerIdSyslogPost) | **Post** /api/2.0/vdn/controller/{controllerId}/syslog | nsxControllerSyslogCreate
[**VdnControllerControllerIdSystemStatsGet**](ControllerApi.md#VdnControllerControllerIdSystemStatsGet) | **Get** /api/2.0/vdn/controller/{controllerId}/systemStats | nsxControllerStatsList
[**VdnControllerControllerIdTechsupportlogsGet**](ControllerApi.md#VdnControllerControllerIdTechsupportlogsGet) | **Get** /api/2.0/vdn/controller/{controllerId}/techsupportlogs | nsxControllerLogsRead
[**VdnControllerCredentialPut**](ControllerApi.md#VdnControllerCredentialPut) | **Put** /api/2.0/vdn/controller/credential | nsxControllerPasswordUpdate
[**VdnControllerGet**](ControllerApi.md#VdnControllerGet) | **Get** /api/2.0/vdn/controller | nsxControllersRead
[**VdnControllerPost**](ControllerApi.md#VdnControllerPost) | **Post** /api/2.0/vdn/controller | nsxControllerCreate
[**VdnControllerProgressJobIdGet**](ControllerApi.md#VdnControllerProgressJobIdGet) | **Get** /api/2.0/vdn/controller/progress/{jobId} | nsxControllerJobRead
[**VdnControllerUpgradeAvailableGet**](ControllerApi.md#VdnControllerUpgradeAvailableGet) | **Get** /api/2.0/vdn/controller/upgrade-available | nsxControllerUpgradeAvailabilityRead

# **VdnControllerClusterGet**
> VdnControllerClusterGet(ctx, )
nsxControllerClusterRead

Retrieve cluster wide configuration information for controller.   Parameters:  

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

# **VdnControllerClusterPut**
> VdnControllerClusterPut(ctx, optional)
nsxControllerClusterUpdate

Modify cluster wide configuration information for controller.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ControllerApiVdnControllerClusterPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiVdnControllerClusterPutOpts struct
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

# **VdnControllerControllerIdDelete**
> VdnControllerControllerIdDelete(ctx, controllerId, optional)
nsxControllerDelete

Delete the NSX controller.   Parameters:  controllerId: Specified controller ID.  To retrieve the controller IDs, log in to the vSphere Web Client. Navigate to Networking & Security > Installation > Management, and view the **NSX Controller nodes** section. The controller ID is listed in the **Controller ID** or **Controller Node** column, depending on NSX version. An example controller ID is *controller-1*.   In a cross-vCenter NSX environment, retrieve the controller IDs from rows where the NSX Manager column contains the primary NSX Manager IP address.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **controllerId** | **string**|  | 
 **optional** | ***ControllerApiVdnControllerControllerIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiVdnControllerControllerIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceRemoval** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnControllerControllerIdPost**
> VdnControllerControllerIdPost(ctx, controllerId, optional)
nsxControllerRemediateAction

If you power off or delete a controller from vCenter, NSX Manager detects the change in controller status. You can remediate the controller, which will power on a powered off controller, or remove the controller from the NSX Manager database if the controller is deleted.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  controllerId: Specified controller ID.  To retrieve the controller IDs, log in to the vSphere Web Client. Navigate to Networking & Security > Installation > Management, and view the **NSX Controller nodes** section. The controller ID is listed in the **Controller ID** or **Controller Node** column, depending on NSX version. An example controller ID is *controller-1*.   In a cross-vCenter NSX environment, retrieve the controller IDs from rows where the NSX Manager column contains the primary NSX Manager IP address.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **controllerId** | **string**|  | 
 **optional** | ***ControllerApiVdnControllerControllerIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiVdnControllerControllerIdPostOpts struct
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

# **VdnControllerControllerIdSnapshotGet**
> VdnControllerControllerIdSnapshotGet(ctx, controllerId)
nsxControllerSnapshotRead

Take a snapshot of the control cluster from the specified controller node.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **controllerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnControllerControllerIdSyslogDelete**
> VdnControllerControllerIdSyslogDelete(ctx, controllerId)
nsxControllerSyslogDelete

Deletes syslog exporter on the specified controller node.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **controllerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnControllerControllerIdSyslogGet**
> VdnControllerControllerIdSyslogGet(ctx, controllerId)
nsxControllerSyslogRead

Retrieve details about the syslog exporter on the controller.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **controllerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnControllerControllerIdSyslogPost**
> VdnControllerControllerIdSyslogPost(ctx, controllerId, optional)
nsxControllerSyslogCreate

Add controller syslog exporter on the controller.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **controllerId** | **string**|  | 
 **optional** | ***ControllerApiVdnControllerControllerIdSyslogPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiVdnControllerControllerIdSyslogPostOpts struct
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

# **VdnControllerControllerIdSystemStatsGet**
> VdnControllerControllerIdSystemStatsGet(ctx, controllerId)
nsxControllerStatsList

Retrieve NSX Controller system statistics.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **controllerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnControllerControllerIdTechsupportlogsGet**
> VdnControllerControllerIdTechsupportlogsGet(ctx, controllerId)
nsxControllerLogsRead

Retrieve controller logs. Response content type is application/octet-stream and response header is filename. This streams a fairly large bundle back (possibly hundreds of MB).   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **controllerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnControllerCredentialPut**
> VdnControllerCredentialPut(ctx, optional)
nsxControllerPasswordUpdate

Change the NSX controller password.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ControllerApiVdnControllerCredentialPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiVdnControllerCredentialPutOpts struct
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

# **VdnControllerGet**
> VdnControllerGet(ctx, )
nsxControllersRead

Retrieves details and runtime status for all controllers.  Runtime status can be one of the following:    * **Deploying** - controller is being deployed and the procedure has not   completed yet.   * **Removing** - controller is being removed and the procedure has not   completed yet.   * **Running** - controller has been deployed and can respond to API   invocation.   * **Unknown** - controller has been deployed but fails to respond to API   invocation.   Parameters:  

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

# **VdnControllerPost**
> VdnControllerPost(ctx, optional)
nsxControllerCreate

Adds a new NSX controller on the specified given cluster. The *hostId* parameter is optional. The *resourcePoolId* can be either the *clusterId* or *resourcePoolId*.  The IP address of the controller node will be allocated from the specified IP pool. The *deployType* property determines the controller node memory size and can be small, medium, or large. However, different controller deployment types are not currently supported because the OVF overrides it and different OVF types require changes in the manager build scripts. Despite not being supported, an arbitrary *deployType* size must still be specified or an error will be returned. Request without body to upgrade controller cluster.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ControllerApiVdnControllerPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiVdnControllerPostOpts struct
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

# **VdnControllerProgressJobIdGet**
> VdnControllerProgressJobIdGet(ctx, jobId)
nsxControllerJobRead

Retrieves status of controller creation or removal. The progress gives a percentage indication of current deploy / remove procedure.   Parameters:  jobId: Specified job Id  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VdnControllerUpgradeAvailableGet**
> VdnControllerUpgradeAvailableGet(ctx, )
nsxControllerUpgradeAvailabilityRead

Retrieve controller upgrade availability.  Parameters:  

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

