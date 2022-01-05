# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesUsermgmtEnablestateValuePut**](UsermgmtApi.md#Api20ServicesUsermgmtEnablestateValuePut) | **Put** /api/2.0/services/usermgmt/enablestate/{value} | userMgmtStateUpdate
[**Api20ServicesUsermgmtRoleUserIdDelete**](UsermgmtApi.md#Api20ServicesUsermgmtRoleUserIdDelete) | **Delete** /api/2.0/services/usermgmt/role/{userId} | userMgmtRoleDelete
[**Api20ServicesUsermgmtRoleUserIdGet**](UsermgmtApi.md#Api20ServicesUsermgmtRoleUserIdGet) | **Get** /api/2.0/services/usermgmt/role/{userId} | userMgmtRoleRead
[**Api20ServicesUsermgmtRoleUserIdPost**](UsermgmtApi.md#Api20ServicesUsermgmtRoleUserIdPost) | **Post** /api/2.0/services/usermgmt/role/{userId} | userMgmtRoleCreate
[**Api20ServicesUsermgmtRoleUserIdPut**](UsermgmtApi.md#Api20ServicesUsermgmtRoleUserIdPut) | **Put** /api/2.0/services/usermgmt/role/{userId} | userMgmtRoleUpdate
[**Api20ServicesUsermgmtRolesGet**](UsermgmtApi.md#Api20ServicesUsermgmtRolesGet) | **Get** /api/2.0/services/usermgmt/roles | userMgmtRolesRead
[**Api20ServicesUsermgmtScopingobjectsGet**](UsermgmtApi.md#Api20ServicesUsermgmtScopingobjectsGet) | **Get** /api/2.0/services/usermgmt/scopingobjects | userScopingObjectsRead
[**Api20ServicesUsermgmtUserUserIdDelete**](UsermgmtApi.md#Api20ServicesUsermgmtUserUserIdDelete) | **Delete** /api/2.0/services/usermgmt/user/{userId} | userMgmtUserDelete
[**Api20ServicesUsermgmtUserUserIdGet**](UsermgmtApi.md#Api20ServicesUsermgmtUserUserIdGet) | **Get** /api/2.0/services/usermgmt/user/{userId} | userMgmtUserRead
[**Api20ServicesUsermgmtUsersVsmGet**](UsermgmtApi.md#Api20ServicesUsermgmtUsersVsmGet) | **Get** /api/2.0/services/usermgmt/users/vsm | userNSXManagerInfoRead

# **Api20ServicesUsermgmtEnablestateValuePut**
> Api20ServicesUsermgmtEnablestateValuePut(ctx, value)
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

# **Api20ServicesUsermgmtRoleUserIdDelete**
> Api20ServicesUsermgmtRoleUserIdDelete(ctx, userId)
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

# **Api20ServicesUsermgmtRoleUserIdGet**
> Api20ServicesUsermgmtRoleUserIdGet(ctx, userId)
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

# **Api20ServicesUsermgmtRoleUserIdPost**
> Api20ServicesUsermgmtRoleUserIdPost(ctx, userId, optional)
userMgmtRoleCreate

Add role and resources for a user.  Parameters:  userId: User to retrieve role information from.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UsermgmtApiApi20ServicesUsermgmtRoleUserIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsermgmtApiApi20ServicesUsermgmtRoleUserIdPostOpts struct
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

# **Api20ServicesUsermgmtRoleUserIdPut**
> Api20ServicesUsermgmtRoleUserIdPut(ctx, userId, optional)
userMgmtRoleUpdate

Change a user's role.  Parameters:  userId: User to retrieve role information from.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UsermgmtApiApi20ServicesUsermgmtRoleUserIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsermgmtApiApi20ServicesUsermgmtRoleUserIdPutOpts struct
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

# **Api20ServicesUsermgmtRolesGet**
> Api20ServicesUsermgmtRolesGet(ctx, )
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

# **Api20ServicesUsermgmtScopingobjectsGet**
> Api20ServicesUsermgmtScopingobjectsGet(ctx, )
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

# **Api20ServicesUsermgmtUserUserIdDelete**
> Api20ServicesUsermgmtUserUserIdDelete(ctx, userId)
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

# **Api20ServicesUsermgmtUserUserIdGet**
> Api20ServicesUsermgmtUserUserIdGet(ctx, userId)
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

# **Api20ServicesUsermgmtUsersVsmGet**
> Api20ServicesUsermgmtUsersVsmGet(ctx, )
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

