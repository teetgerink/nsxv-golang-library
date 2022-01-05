# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api10IdentityDirectoryGroupsForUserGet**](Class10IdentityApi.md#Api10IdentityDirectoryGroupsForUserGet) | **Get** /api/1.0/identity/directoryGroupsForUser | userDirGroupsRead
[**Api10IdentityHostIpMappingGet**](Class10IdentityApi.md#Api10IdentityHostIpMappingGet) | **Get** /api/1.0/identity/hostIpMapping | hostIpMappingRead
[**Api10IdentityIpToUserMappingGet**](Class10IdentityApi.md#Api10IdentityIpToUserMappingGet) | **Get** /api/1.0/identity/ipToUserMapping | ipToUserMappingRead
[**Api10IdentityStaticUserMappingUserIDIPPost**](Class10IdentityApi.md#Api10IdentityStaticUserMappingUserIDIPPost) | **Post** /api/1.0/identity/staticUserMapping/{userID}/{IP} | staticUserMappingCreate
[**Api10IdentityStaticUserMappingsGet**](Class10IdentityApi.md#Api10IdentityStaticUserMappingsGet) | **Get** /api/1.0/identity/staticUserMappings | staticUserMappingsRead
[**Api10IdentityStaticUserMappingsbyIPIPDelete**](Class10IdentityApi.md#Api10IdentityStaticUserMappingsbyIPIPDelete) | **Delete** /api/1.0/identity/staticUserMappingsbyIP/{IP} | staticUserMappingsbyIPDelete
[**Api10IdentityStaticUserMappingsbyIPIPGet**](Class10IdentityApi.md#Api10IdentityStaticUserMappingsbyIPIPGet) | **Get** /api/1.0/identity/staticUserMappingsbyIP/{IP} | staticUserMappingsbyIPRead
[**Api10IdentityStaticUserMappingsbyUserUserIDDelete**](Class10IdentityApi.md#Api10IdentityStaticUserMappingsbyUserUserIDDelete) | **Delete** /api/1.0/identity/staticUserMappingsbyUser/{userID} | staticUserMappingsbyUserDelete
[**Api10IdentityStaticUserMappingsbyUserUserIDGet**](Class10IdentityApi.md#Api10IdentityStaticUserMappingsbyUserUserIDGet) | **Get** /api/1.0/identity/staticUserMappingsbyUser/{userID} | staticUserMappingsbyUserRead
[**Api10IdentityUserIpMappingGet**](Class10IdentityApi.md#Api10IdentityUserIpMappingGet) | **Get** /api/1.0/identity/userIpMapping | userIpMappingRead

# **Api10IdentityDirectoryGroupsForUserGet**
> Api10IdentityDirectoryGroupsForUserGet(ctx, )
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

# **Api10IdentityHostIpMappingGet**
> Api10IdentityHostIpMappingGet(ctx, )
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

# **Api10IdentityIpToUserMappingGet**
> Api10IdentityIpToUserMappingGet(ctx, )
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

# **Api10IdentityStaticUserMappingUserIDIPPost**
> Api10IdentityStaticUserMappingUserIDIPPost(ctx, userID, iP)
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

# **Api10IdentityStaticUserMappingsGet**
> Api10IdentityStaticUserMappingsGet(ctx, )
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

# **Api10IdentityStaticUserMappingsbyIPIPDelete**
> Api10IdentityStaticUserMappingsbyIPIPDelete(ctx, iP)
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

# **Api10IdentityStaticUserMappingsbyIPIPGet**
> Api10IdentityStaticUserMappingsbyIPIPGet(ctx, iP)
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

# **Api10IdentityStaticUserMappingsbyUserUserIDDelete**
> Api10IdentityStaticUserMappingsbyUserUserIDDelete(ctx, userID)
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

# **Api10IdentityStaticUserMappingsbyUserUserIDGet**
> Api10IdentityStaticUserMappingsbyUserUserIDGet(ctx, userID)
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

# **Api10IdentityUserIpMappingGet**
> Api10IdentityUserIpMappingGet(ctx, )
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

