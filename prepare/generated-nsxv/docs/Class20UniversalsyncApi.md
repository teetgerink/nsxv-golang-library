# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20UniversalsyncConfigurationNsxmanagersDelete**](Class20UniversalsyncApi.md#Api20UniversalsyncConfigurationNsxmanagersDelete) | **Delete** /api/2.0/universalsync/configuration/nsxmanagers | universalSyncConfigurationNsxManagersDelete
[**Api20UniversalsyncConfigurationNsxmanagersGet**](Class20UniversalsyncApi.md#Api20UniversalsyncConfigurationNsxmanagersGet) | **Get** /api/2.0/universalsync/configuration/nsxmanagers | universalSyncConfigurationNsxManagersList
[**Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDDelete**](Class20UniversalsyncApi.md#Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDDelete) | **Delete** /api/2.0/universalsync/configuration/nsxmanagers/{nsxManagerID} | universalSyncConfigurationManagersDelete
[**Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDGet**](Class20UniversalsyncApi.md#Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDGet) | **Get** /api/2.0/universalsync/configuration/nsxmanagers/{nsxManagerID} | universalSyncConfigurationManagersRead
[**Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDPut**](Class20UniversalsyncApi.md#Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDPut) | **Put** /api/2.0/universalsync/configuration/nsxmanagers/{nsxManagerID} | universalSyncConfigurationManagersUpdate
[**Api20UniversalsyncConfigurationNsxmanagersPost**](Class20UniversalsyncApi.md#Api20UniversalsyncConfigurationNsxmanagersPost) | **Post** /api/2.0/universalsync/configuration/nsxmanagers | universalSyncConfigurationNsxManagersCreate
[**Api20UniversalsyncConfigurationRoleGet**](Class20UniversalsyncApi.md#Api20UniversalsyncConfigurationRoleGet) | **Get** /api/2.0/universalsync/configuration/role | universalSyncConfigurationRoleRead
[**Api20UniversalsyncConfigurationRolePost**](Class20UniversalsyncApi.md#Api20UniversalsyncConfigurationRolePost) | **Post** /api/2.0/universalsync/configuration/role | universalSyncConfigurationRoleSet
[**Api20UniversalsyncEntitystatusGet**](Class20UniversalsyncApi.md#Api20UniversalsyncEntitystatusGet) | **Get** /api/2.0/universalsync/entitystatus | universalSyncEntityStatusRead
[**Api20UniversalsyncStatusGet**](Class20UniversalsyncApi.md#Api20UniversalsyncStatusGet) | **Get** /api/2.0/universalsync/status | universalSyncStatusRead
[**Api20UniversalsyncSyncPost**](Class20UniversalsyncApi.md#Api20UniversalsyncSyncPost) | **Post** /api/2.0/universalsync/sync | universalSyncSyncAction

# **Api20UniversalsyncConfigurationNsxmanagersDelete**
> Api20UniversalsyncConfigurationNsxmanagersDelete(ctx, )
universalSyncConfigurationNsxManagersDelete

Delete secondary NSX manager configuration.  Parameters:  

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

# **Api20UniversalsyncConfigurationNsxmanagersGet**
> Api20UniversalsyncConfigurationNsxmanagersGet(ctx, )
universalSyncConfigurationNsxManagersList

If run on a primary NSX Manager, it will list secondary NSX Managers configured on the primary NSX Manager.  If run on a secondary NSX Manager, it will list information about the secondary NSX Manager and the primary NSX Manager it is associated with.   Parameters:  

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

# **Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDDelete**
> Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDDelete(ctx, nsxManagerID, optional)
universalSyncConfigurationManagersDelete

Delete the specified secondary NSX Manager.  Parameters:  nsxManagerID: NSX Manager UUID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsxManagerID** | **string**|  | 
 **optional** | ***Class20UniversalsyncApiApi20UniversalsyncConfigurationNsxmanagersNsxManagerIDDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20UniversalsyncApiApi20UniversalsyncConfigurationNsxmanagersNsxManagerIDDeleteOpts struct
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

# **Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDGet**
> Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDGet(ctx, nsxManagerID)
universalSyncConfigurationManagersRead

Retrieve information about the specified secondary NSX Manager.   Parameters:  nsxManagerID: NSX Manager UUID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsxManagerID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDPut**
> Api20UniversalsyncConfigurationNsxmanagersNsxManagerIDPut(ctx, nsxManagerID, optional)
universalSyncConfigurationManagersUpdate

Update the the specified secondary NSX manager IP or thumbprint in the universal sync configuration.   Parameters:  nsxManagerID: NSX Manager UUID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsxManagerID** | **string**|  | 
 **optional** | ***Class20UniversalsyncApiApi20UniversalsyncConfigurationNsxmanagersNsxManagerIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20UniversalsyncApiApi20UniversalsyncConfigurationNsxmanagersNsxManagerIDPutOpts struct
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

# **Api20UniversalsyncConfigurationNsxmanagersPost**
> Api20UniversalsyncConfigurationNsxmanagersPost(ctx, optional)
universalSyncConfigurationNsxManagersCreate

Add a secondary NSX manager.  Run this method on the primary NSX Manager, providing details of the secondary NSX Manager.  Retrieve the certificate thumbprint of the secondary NSX Manager using the `GET /api/1.0/appliance-management/certificatemanager/certificates/nsx` method. The **sha1Hash** parameter contains the thumbprint.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20UniversalsyncApiApi20UniversalsyncConfigurationNsxmanagersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20UniversalsyncApiApi20UniversalsyncConfigurationNsxmanagersPostOpts struct
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

# **Api20UniversalsyncConfigurationRoleGet**
> Api20UniversalsyncConfigurationRoleGet(ctx, )
universalSyncConfigurationRoleRead

Retrieve the universal sync configuration role.  Parameters:  

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

# **Api20UniversalsyncConfigurationRolePost**
> Api20UniversalsyncConfigurationRolePost(ctx, optional)
universalSyncConfigurationRoleSet

Set the universal sync configuration role.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20UniversalsyncApiApi20UniversalsyncConfigurationRolePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20UniversalsyncApiApi20UniversalsyncConfigurationRolePostOpts struct
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

# **Api20UniversalsyncEntitystatusGet**
> Api20UniversalsyncEntitystatusGet(ctx, optional)
universalSyncEntityStatusRead

Retrieve the status of a universal sync entity.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20UniversalsyncApiApi20UniversalsyncEntitystatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20UniversalsyncApiApi20UniversalsyncEntitystatusGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectType** | **optional.String**|  | 
 **objectId** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20UniversalsyncStatusGet**
> Api20UniversalsyncStatusGet(ctx, )
universalSyncStatusRead

Retrieve the universal sync status.  Parameters:  

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

# **Api20UniversalsyncSyncPost**
> Api20UniversalsyncSyncPost(ctx, optional)
universalSyncSyncAction

Sync all objects on the NSX Manager.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20UniversalsyncApiApi20UniversalsyncSyncPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20UniversalsyncApiApi20UniversalsyncSyncPostOpts struct
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

