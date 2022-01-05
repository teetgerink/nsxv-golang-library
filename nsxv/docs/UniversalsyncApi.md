# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UniversalsyncConfigurationNsxmanagersDelete**](UniversalsyncApi.md#UniversalsyncConfigurationNsxmanagersDelete) | **Delete** /api/2.0/universalsync/configuration/nsxmanagers | universalSyncConfigurationNsxManagersDelete
[**UniversalsyncConfigurationNsxmanagersGet**](UniversalsyncApi.md#UniversalsyncConfigurationNsxmanagersGet) | **Get** /api/2.0/universalsync/configuration/nsxmanagers | universalSyncConfigurationNsxManagersList
[**UniversalsyncConfigurationNsxmanagersNsxManagerIDDelete**](UniversalsyncApi.md#UniversalsyncConfigurationNsxmanagersNsxManagerIDDelete) | **Delete** /api/2.0/universalsync/configuration/nsxmanagers/{nsxManagerID} | universalSyncConfigurationManagersDelete
[**UniversalsyncConfigurationNsxmanagersNsxManagerIDGet**](UniversalsyncApi.md#UniversalsyncConfigurationNsxmanagersNsxManagerIDGet) | **Get** /api/2.0/universalsync/configuration/nsxmanagers/{nsxManagerID} | universalSyncConfigurationManagersRead
[**UniversalsyncConfigurationNsxmanagersNsxManagerIDPut**](UniversalsyncApi.md#UniversalsyncConfigurationNsxmanagersNsxManagerIDPut) | **Put** /api/2.0/universalsync/configuration/nsxmanagers/{nsxManagerID} | universalSyncConfigurationManagersUpdate
[**UniversalsyncConfigurationNsxmanagersPost**](UniversalsyncApi.md#UniversalsyncConfigurationNsxmanagersPost) | **Post** /api/2.0/universalsync/configuration/nsxmanagers | universalSyncConfigurationNsxManagersCreate
[**UniversalsyncConfigurationRoleGet**](UniversalsyncApi.md#UniversalsyncConfigurationRoleGet) | **Get** /api/2.0/universalsync/configuration/role | universalSyncConfigurationRoleRead
[**UniversalsyncConfigurationRolePost**](UniversalsyncApi.md#UniversalsyncConfigurationRolePost) | **Post** /api/2.0/universalsync/configuration/role | universalSyncConfigurationRoleSet
[**UniversalsyncEntitystatusGet**](UniversalsyncApi.md#UniversalsyncEntitystatusGet) | **Get** /api/2.0/universalsync/entitystatus | universalSyncEntityStatusRead
[**UniversalsyncStatusGet**](UniversalsyncApi.md#UniversalsyncStatusGet) | **Get** /api/2.0/universalsync/status | universalSyncStatusRead
[**UniversalsyncSyncPost**](UniversalsyncApi.md#UniversalsyncSyncPost) | **Post** /api/2.0/universalsync/sync | universalSyncSyncAction

# **UniversalsyncConfigurationNsxmanagersDelete**
> UniversalsyncConfigurationNsxmanagersDelete(ctx, )
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

# **UniversalsyncConfigurationNsxmanagersGet**
> UniversalsyncConfigurationNsxmanagersGet(ctx, )
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

# **UniversalsyncConfigurationNsxmanagersNsxManagerIDDelete**
> UniversalsyncConfigurationNsxmanagersNsxManagerIDDelete(ctx, nsxManagerID, optional)
universalSyncConfigurationManagersDelete

Delete the specified secondary NSX Manager.  Parameters:  nsxManagerID: NSX Manager UUID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsxManagerID** | **string**|  | 
 **optional** | ***UniversalsyncApiUniversalsyncConfigurationNsxmanagersNsxManagerIDDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniversalsyncApiUniversalsyncConfigurationNsxmanagersNsxManagerIDDeleteOpts struct
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

# **UniversalsyncConfigurationNsxmanagersNsxManagerIDGet**
> UniversalsyncConfigurationNsxmanagersNsxManagerIDGet(ctx, nsxManagerID)
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

# **UniversalsyncConfigurationNsxmanagersNsxManagerIDPut**
> UniversalsyncConfigurationNsxmanagersNsxManagerIDPut(ctx, nsxManagerID, optional)
universalSyncConfigurationManagersUpdate

Update the the specified secondary NSX manager IP or thumbprint in the universal sync configuration.   Parameters:  nsxManagerID: NSX Manager UUID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsxManagerID** | **string**|  | 
 **optional** | ***UniversalsyncApiUniversalsyncConfigurationNsxmanagersNsxManagerIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniversalsyncApiUniversalsyncConfigurationNsxmanagersNsxManagerIDPutOpts struct
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

# **UniversalsyncConfigurationNsxmanagersPost**
> UniversalsyncConfigurationNsxmanagersPost(ctx, optional)
universalSyncConfigurationNsxManagersCreate

Add a secondary NSX manager.  Run this method on the primary NSX Manager, providing details of the secondary NSX Manager.  Retrieve the certificate thumbprint of the secondary NSX Manager using the `GET /api/1.0/appliance-management/certificatemanager/certificates/nsx` method. The **sha1Hash** parameter contains the thumbprint.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniversalsyncApiUniversalsyncConfigurationNsxmanagersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniversalsyncApiUniversalsyncConfigurationNsxmanagersPostOpts struct
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

# **UniversalsyncConfigurationRoleGet**
> UniversalsyncConfigurationRoleGet(ctx, )
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

# **UniversalsyncConfigurationRolePost**
> UniversalsyncConfigurationRolePost(ctx, optional)
universalSyncConfigurationRoleSet

Set the universal sync configuration role.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniversalsyncApiUniversalsyncConfigurationRolePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniversalsyncApiUniversalsyncConfigurationRolePostOpts struct
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

# **UniversalsyncEntitystatusGet**
> UniversalsyncEntitystatusGet(ctx, optional)
universalSyncEntityStatusRead

Retrieve the status of a universal sync entity.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniversalsyncApiUniversalsyncEntitystatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniversalsyncApiUniversalsyncEntitystatusGetOpts struct
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

# **UniversalsyncStatusGet**
> UniversalsyncStatusGet(ctx, )
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

# **UniversalsyncSyncPost**
> UniversalsyncSyncPost(ctx, optional)
universalSyncSyncAction

Sync all objects on the NSX Manager.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniversalsyncApiUniversalsyncSyncPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniversalsyncApiUniversalsyncSyncPostOpts struct
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

