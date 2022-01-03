# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api21AppExcludelistGet**](Class21ExcludelistApi.md#Api21AppExcludelistGet) | **Get** /api/2.1/app/excludelist | dfwExclusionRead
[**Api21AppExcludelistMemberIDDelete**](Class21ExcludelistApi.md#Api21AppExcludelistMemberIDDelete) | **Delete** /api/2.1/app/excludelist/{memberID} | dfwExclusionDelete
[**Api21AppExcludelistMemberIDPut**](Class21ExcludelistApi.md#Api21AppExcludelistMemberIDPut) | **Put** /api/2.1/app/excludelist/{memberID} | dfwExclusionUpdate

# **Api21AppExcludelistGet**
> Api21AppExcludelistGet(ctx, )
dfwExclusionRead

Retrieve the set of VMs in the exclusion list.  Parameters:  

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

# **Api21AppExcludelistMemberIDDelete**
> Api21AppExcludelistMemberIDDelete(ctx, memberID)
dfwExclusionDelete

Delete a vm from exclusion list.  Parameters:  memberID: vc-moref-id of a virtual machine.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **memberID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api21AppExcludelistMemberIDPut**
> Api21AppExcludelistMemberIDPut(ctx, memberID)
dfwExclusionUpdate

Add a vm to the exclusion list.  Parameters:  memberID: vc-moref-id of a virtual machine.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **memberID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

