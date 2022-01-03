# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesApplicationgroupApplicationgroupIdDelete**](Class20ApplicationgroupApi.md#Api20ServicesApplicationgroupApplicationgroupIdDelete) | **Delete** /api/2.0/services/applicationgroup/{applicationgroupId} | serviceGroupDelete
[**Api20ServicesApplicationgroupApplicationgroupIdGet**](Class20ApplicationgroupApi.md#Api20ServicesApplicationgroupApplicationgroupIdGet) | **Get** /api/2.0/services/applicationgroup/{applicationgroupId} | serviceGroupRead
[**Api20ServicesApplicationgroupApplicationgroupIdMembersMorefDelete**](Class20ApplicationgroupApi.md#Api20ServicesApplicationgroupApplicationgroupIdMembersMorefDelete) | **Delete** /api/2.0/services/applicationgroup/{applicationgroupId}/members/{moref} | serviceGroupMembersDelete
[**Api20ServicesApplicationgroupApplicationgroupIdMembersMorefPut**](Class20ApplicationgroupApi.md#Api20ServicesApplicationgroupApplicationgroupIdMembersMorefPut) | **Put** /api/2.0/services/applicationgroup/{applicationgroupId}/members/{moref} | serviceGroupMembersCreate
[**Api20ServicesApplicationgroupApplicationgroupIdPut**](Class20ApplicationgroupApi.md#Api20ServicesApplicationgroupApplicationgroupIdPut) | **Put** /api/2.0/services/applicationgroup/{applicationgroupId} | serviceGroupUpdate
[**Api20ServicesApplicationgroupScopeScopeIdGet**](Class20ApplicationgroupApi.md#Api20ServicesApplicationgroupScopeScopeIdGet) | **Get** /api/2.0/services/applicationgroup/scope/{scopeId} | serviceGroupsRead
[**Api20ServicesApplicationgroupScopeScopeIdMembersGet**](Class20ApplicationgroupApi.md#Api20ServicesApplicationgroupScopeScopeIdMembersGet) | **Get** /api/2.0/services/applicationgroup/scope/{scopeId}/members | serviceGroupMembersScopeList
[**Api20ServicesApplicationgroupScopeScopeIdPost**](Class20ApplicationgroupApi.md#Api20ServicesApplicationgroupScopeScopeIdPost) | **Post** /api/2.0/services/applicationgroup/scope/{scopeId} | serviceGroupsCreate

# **Api20ServicesApplicationgroupApplicationgroupIdDelete**
> Api20ServicesApplicationgroupApplicationgroupIdDelete(ctx, applicationgroupId, optional)
serviceGroupDelete

Delete the specified service group from a scope.  Parameters:  applicationgroupId: Application group ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationgroupId** | **string**|  | 
 **optional** | ***Class20ApplicationgroupApiApi20ServicesApplicationgroupApplicationgroupIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20ApplicationgroupApiApi20ServicesApplicationgroupApplicationgroupIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesApplicationgroupApplicationgroupIdGet**
> Api20ServicesApplicationgroupApplicationgroupIdGet(ctx, applicationgroupId)
serviceGroupRead

Retrieve details about the specified service group.  Parameters:  applicationgroupId: Application group ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationgroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesApplicationgroupApplicationgroupIdMembersMorefDelete**
> Api20ServicesApplicationgroupApplicationgroupIdMembersMorefDelete(ctx, applicationgroupId, moref)
serviceGroupMembersDelete

Delete a member from the service group.  Parameters:  moref: Managed object reference to the member.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationgroupId** | **string**|  | 
  **moref** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesApplicationgroupApplicationgroupIdMembersMorefPut**
> Api20ServicesApplicationgroupApplicationgroupIdMembersMorefPut(ctx, applicationgroupId, moref)
serviceGroupMembersCreate

Add a member to the service group.  Parameters:  moref: Managed object reference to the member.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationgroupId** | **string**|  | 
  **moref** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesApplicationgroupApplicationgroupIdPut**
> Api20ServicesApplicationgroupApplicationgroupIdPut(ctx, applicationgroupId, optional)
serviceGroupUpdate

Modify the name, description, applicationProtocol, or port value of the specified service group.   Parameters:  applicationgroupId: Application group ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationgroupId** | **string**|  | 
 **optional** | ***Class20ApplicationgroupApiApi20ServicesApplicationgroupApplicationgroupIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20ApplicationgroupApiApi20ServicesApplicationgroupApplicationgroupIdPutOpts struct
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

# **Api20ServicesApplicationgroupScopeScopeIdGet**
> Api20ServicesApplicationgroupScopeScopeIdGet(ctx, scopeId)
serviceGroupsRead

Retrieve a list of service groups that have been created on the scope.   Parameters:  scopeId: The scopeId can be \"globalroot-0\", \"universalroot-0\" or datacenterId in upgrade use cases   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesApplicationgroupScopeScopeIdMembersGet**
> Api20ServicesApplicationgroupScopeScopeIdMembersGet(ctx, scopeId)
serviceGroupMembersScopeList

Get a list of member elements that can be added to the service groups created on a particular scope.  Since service group allows only either services or other service groups as members to be added, this helps you get a list of all possible valid elements that can be added to the service.   Parameters:  scopeId: globalroot-0 or datacenterId in upgrade use cases  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesApplicationgroupScopeScopeIdPost**
> Api20ServicesApplicationgroupScopeScopeIdPost(ctx, scopeId, optional)
serviceGroupsCreate

Create a new service group on the specified scope.  Parameters:  scopeId: The scopeId can be \"globalroot-0\", \"universalroot-0\" or datacenterId in upgrade use cases   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***Class20ApplicationgroupApiApi20ServicesApplicationgroupScopeScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20ApplicationgroupApiApi20ServicesApplicationgroupScopeScopeIdPostOpts struct
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

