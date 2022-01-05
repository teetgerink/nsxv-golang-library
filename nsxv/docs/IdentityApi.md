# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityDirectoryGroupsForUserGet**](IdentityApi.md#IdentityDirectoryGroupsForUserGet) | **Get** /api/1.0/identity/directoryGroupsForUser | userDirGroupsRead
[**IdentityHostIpMappingGet**](IdentityApi.md#IdentityHostIpMappingGet) | **Get** /api/1.0/identity/hostIpMapping | hostIpMappingRead
[**IdentityIpToUserMappingGet**](IdentityApi.md#IdentityIpToUserMappingGet) | **Get** /api/1.0/identity/ipToUserMapping | ipToUserMappingRead
[**IdentityStaticUserMappingUserIDIPPost**](IdentityApi.md#IdentityStaticUserMappingUserIDIPPost) | **Post** /api/1.0/identity/staticUserMapping/{userID}/{IP} | staticUserMappingCreate
[**IdentityStaticUserMappingsGet**](IdentityApi.md#IdentityStaticUserMappingsGet) | **Get** /api/1.0/identity/staticUserMappings | staticUserMappingsRead
[**IdentityStaticUserMappingsbyIPIPDelete**](IdentityApi.md#IdentityStaticUserMappingsbyIPIPDelete) | **Delete** /api/1.0/identity/staticUserMappingsbyIP/{IP} | staticUserMappingsbyIPDelete
[**IdentityStaticUserMappingsbyIPIPGet**](IdentityApi.md#IdentityStaticUserMappingsbyIPIPGet) | **Get** /api/1.0/identity/staticUserMappingsbyIP/{IP} | staticUserMappingsbyIPRead
[**IdentityStaticUserMappingsbyUserUserIDDelete**](IdentityApi.md#IdentityStaticUserMappingsbyUserUserIDDelete) | **Delete** /api/1.0/identity/staticUserMappingsbyUser/{userID} | staticUserMappingsbyUserDelete
[**IdentityStaticUserMappingsbyUserUserIDGet**](IdentityApi.md#IdentityStaticUserMappingsbyUserUserIDGet) | **Get** /api/1.0/identity/staticUserMappingsbyUser/{userID} | staticUserMappingsbyUserRead
[**IdentityUserIpMappingGet**](IdentityApi.md#IdentityUserIpMappingGet) | **Get** /api/1.0/identity/userIpMapping | userIpMappingRead

# **IdentityDirectoryGroupsForUserGet**
> IdentityDirectoryGroupsForUserGet(ctx, )
userDirGroupsRead

Query set of Windows Domain Groups (AD Groups) to which the specified user belongs.   Parameters:  

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

# **IdentityHostIpMappingGet**
> IdentityHostIpMappingGet(ctx, )
hostIpMappingRead

Query host-to-ip mapping list from database.  Parameters:  

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

# **IdentityIpToUserMappingGet**
> IdentityIpToUserMappingGet(ctx, )
ipToUserMappingRead

Retrieve set of users associated with a given set of IP addresses during a specified time period. Since more than one user can be associated with a single IP address during the specified time period, each IP address can be associated with zero or more (i.e a SET of) users.   Parameters:  

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

# **IdentityStaticUserMappingUserIDIPPost**
> IdentityStaticUserMappingUserIDIPPost(ctx, userID, iP)
staticUserMappingCreate

Create static user IP mapping.  Parameters:  userID: User ID  IP: IP address  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userID** | **string**|  | 
  **iP** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IdentityStaticUserMappingsGet**
> IdentityStaticUserMappingsGet(ctx, )
staticUserMappingsRead

Query static user IP mapping list.  Parameters:  

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

# **IdentityStaticUserMappingsbyIPIPDelete**
> IdentityStaticUserMappingsbyIPIPDelete(ctx, iP)
staticUserMappingsbyIPDelete

Delete static user IP mapping for specified IP.  Parameters:  IP: IP address  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iP** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IdentityStaticUserMappingsbyIPIPGet**
> IdentityStaticUserMappingsbyIPIPGet(ctx, iP)
staticUserMappingsbyIPRead

Query static user IP mapping for specified IP.  Parameters:  IP: IP address  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iP** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IdentityStaticUserMappingsbyUserUserIDDelete**
> IdentityStaticUserMappingsbyUserUserIDDelete(ctx, userID)
staticUserMappingsbyUserDelete

Delete static user IP mapping for specified user.  Parameters:  userID: User ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IdentityStaticUserMappingsbyUserUserIDGet**
> IdentityStaticUserMappingsbyUserUserIDGet(ctx, userID)
staticUserMappingsbyUserRead

Query static user IP mapping for specified user.  Parameters:  userID: User ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IdentityUserIpMappingGet**
> IdentityUserIpMappingGet(ctx, )
userIpMappingRead

Query user-to-ip mapping list from database.  Parameters:  

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

