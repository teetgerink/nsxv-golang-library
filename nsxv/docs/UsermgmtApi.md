# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesUsermgmtEnablestateValuePut**](UsermgmtApi.md#ServicesUsermgmtEnablestateValuePut) | **Put** /api/2.0/services/usermgmt/enablestate/{value} | userMgmtStateUpdate
[**ServicesUsermgmtRoleUserIdDelete**](UsermgmtApi.md#ServicesUsermgmtRoleUserIdDelete) | **Delete** /api/2.0/services/usermgmt/role/{userId} | userMgmtRoleDelete
[**ServicesUsermgmtRoleUserIdGet**](UsermgmtApi.md#ServicesUsermgmtRoleUserIdGet) | **Get** /api/2.0/services/usermgmt/role/{userId} | userMgmtRoleRead
[**ServicesUsermgmtRoleUserIdPost**](UsermgmtApi.md#ServicesUsermgmtRoleUserIdPost) | **Post** /api/2.0/services/usermgmt/role/{userId} | userMgmtRoleCreate
[**ServicesUsermgmtRoleUserIdPut**](UsermgmtApi.md#ServicesUsermgmtRoleUserIdPut) | **Put** /api/2.0/services/usermgmt/role/{userId} | userMgmtRoleUpdate
[**ServicesUsermgmtRolesGet**](UsermgmtApi.md#ServicesUsermgmtRolesGet) | **Get** /api/2.0/services/usermgmt/roles | userMgmtRolesRead
[**ServicesUsermgmtScopingobjectsGet**](UsermgmtApi.md#ServicesUsermgmtScopingobjectsGet) | **Get** /api/2.0/services/usermgmt/scopingobjects | userScopingObjectsRead
[**ServicesUsermgmtUserUserIdDelete**](UsermgmtApi.md#ServicesUsermgmtUserUserIdDelete) | **Delete** /api/2.0/services/usermgmt/user/{userId} | userMgmtUserDelete
[**ServicesUsermgmtUserUserIdGet**](UsermgmtApi.md#ServicesUsermgmtUserUserIdGet) | **Get** /api/2.0/services/usermgmt/user/{userId} | userMgmtUserRead
[**ServicesUsermgmtUsersVsmGet**](UsermgmtApi.md#ServicesUsermgmtUsersVsmGet) | **Get** /api/2.0/services/usermgmt/users/vsm | userNSXManagerInfoRead

# **ServicesUsermgmtEnablestateValuePut**
> ServicesUsermgmtEnablestateValuePut(ctx, value)
userMgmtStateUpdate

Enable or disable a user account.  Parameters:  value: value can be 0 to disable, or 1 to enable.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **value** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesUsermgmtRoleUserIdDelete**
> ServicesUsermgmtRoleUserIdDelete(ctx, userId)
userMgmtRoleDelete

Delete the role assignment for specified vCenter user. Once this role is deleted, the user is removed from NSX Manager. You cannot delete the role for a local user.   Parameters:  userId: User to retrieve role information from.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesUsermgmtRoleUserIdGet**
> ServicesUsermgmtRoleUserIdGet(ctx, userId)
userMgmtRoleRead

Retrieve a user's role (possible roles are super_user, vshield_admin, enterprise_admin, security_admin, and audit).   Parameters:  userId: User to retrieve role information from.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesUsermgmtRoleUserIdPost**
> ServicesUsermgmtRoleUserIdPost(ctx, userId, optional)
userMgmtRoleCreate

Add role and resources for a user.  Parameters:  userId: User to retrieve role information from.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UsermgmtApiServicesUsermgmtRoleUserIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsermgmtApiServicesUsermgmtRoleUserIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **isGroup** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesUsermgmtRoleUserIdPut**
> ServicesUsermgmtRoleUserIdPut(ctx, userId, optional)
userMgmtRoleUpdate

Change a user's role.  Parameters:  userId: User to retrieve role information from.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UsermgmtApiServicesUsermgmtRoleUserIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsermgmtApiServicesUsermgmtRoleUserIdPutOpts struct
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

# **ServicesUsermgmtRolesGet**
> ServicesUsermgmtRolesGet(ctx, )
userMgmtRolesRead

Read all possible roles in NSX Manager  Parameters:  

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

# **ServicesUsermgmtScopingobjectsGet**
> ServicesUsermgmtScopingobjectsGet(ctx, )
userScopingObjectsRead

Retrieve a list of objects that can be used to define a user's access scope.   Parameters:  

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

# **ServicesUsermgmtUserUserIdDelete**
> ServicesUsermgmtUserUserIdDelete(ctx, userId)
userMgmtUserDelete

Remove the NSX role for a vCenter user.  Parameters:  userId: user ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesUsermgmtUserUserIdGet**
> ServicesUsermgmtUserUserIdGet(ctx, userId)
userMgmtUserRead

Get information about a user.  Parameters:  userId: user ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesUsermgmtUsersVsmGet**
> ServicesUsermgmtUsersVsmGet(ctx, )
userNSXManagerInfoRead

Get information about users who have been assigned a NSX Manager role (local users as well as vCenter users with NSX Manager role).   Parameters:  

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

