# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesSecuritytagsSelectionCriteriaGet**](SecuritytagsApi.md#ServicesSecuritytagsSelectionCriteriaGet) | **Get** /api/2.0/services/securitytags/selection-criteria | securitytagsCriteriaRead
[**ServicesSecuritytagsSelectionCriteriaPut**](SecuritytagsApi.md#ServicesSecuritytagsSelectionCriteriaPut) | **Put** /api/2.0/services/securitytags/selection-criteria | securitytagsCriteriaUpdate
[**ServicesSecuritytagsTagGet**](SecuritytagsApi.md#ServicesSecuritytagsTagGet) | **Get** /api/2.0/services/securitytags/tag | securityTagRead
[**ServicesSecuritytagsTagPost**](SecuritytagsApi.md#ServicesSecuritytagsTagPost) | **Post** /api/2.0/services/securitytags/tag | securityTagCreate
[**ServicesSecuritytagsTagTagIdDelete**](SecuritytagsApi.md#ServicesSecuritytagsTagTagIdDelete) | **Delete** /api/2.0/services/securitytags/tag/{tagId} | securityTagDeleteDelete
[**ServicesSecuritytagsTagTagIdVmDetailGet**](SecuritytagsApi.md#ServicesSecuritytagsTagTagIdVmDetailGet) | **Get** /api/2.0/services/securitytags/tag/{tagId}/vmDetail | securityTagVMDetailList
[**ServicesSecuritytagsTagTagIdVmGet**](SecuritytagsApi.md#ServicesSecuritytagsTagTagIdVmGet) | **Get** /api/2.0/services/securitytags/tag/{tagId}/vm | securityTagVMsListList
[**ServicesSecuritytagsTagTagIdVmPost**](SecuritytagsApi.md#ServicesSecuritytagsTagTagIdVmPost) | **Post** /api/2.0/services/securitytags/tag/{tagId}/vm | securityTagVMsAction
[**ServicesSecuritytagsTagTagIdVmVmIdDelete**](SecuritytagsApi.md#ServicesSecuritytagsTagTagIdVmVmIdDelete) | **Delete** /api/2.0/services/securitytags/tag/{tagId}/vm/{vmId} | securityTagVMDetach
[**ServicesSecuritytagsTagTagIdVmVmIdPut**](SecuritytagsApi.md#ServicesSecuritytagsTagTagIdVmVmIdPut) | **Put** /api/2.0/services/securitytags/tag/{tagId}/vm/{vmId} | securityTagVMAttach
[**ServicesSecuritytagsVmVmIdGet**](SecuritytagsApi.md#ServicesSecuritytagsVmVmIdGet) | **Get** /api/2.0/services/securitytags/vm/{vmId} | securitytagVmIdList
[**ServicesSecuritytagsVmVmIdPost**](SecuritytagsApi.md#ServicesSecuritytagsVmVmIdPost) | **Post** /api/2.0/services/securitytags/vm/{vmId} | securitytagVmMoidAction

# **ServicesSecuritytagsSelectionCriteriaGet**
> ServicesSecuritytagsSelectionCriteriaGet(ctx, )
securitytagsCriteriaRead

Retrieve unique ID section criteria configuration.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  

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

# **ServicesSecuritytagsSelectionCriteriaPut**
> ServicesSecuritytagsSelectionCriteriaPut(ctx, optional)
securitytagsCriteriaUpdate

Configure the unique ID section criteria configuration.  If you set the selection criteria and assign security tags to VMs, you must remove all security tags from VMs before you can change the selection criteria.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecuritytagsApiServicesSecuritytagsSelectionCriteriaPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritytagsApiServicesSecuritytagsSelectionCriteriaPutOpts struct
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

# **ServicesSecuritytagsTagGet**
> ServicesSecuritytagsTagGet(ctx, optional)
securityTagRead

Retrieve all security tags.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method updated. Added **isUniversal** query parameter to filter universal security tags.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecuritytagsApiServicesSecuritytagsTagGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritytagsApiServicesSecuritytagsTagGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isUniversal** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritytagsTagPost**
> ServicesSecuritytagsTagPost(ctx, optional)
securityTagCreate

Create a new security tag.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method updated. **isUniversal** parameter can be set to create a universal security tag.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecuritytagsApiServicesSecuritytagsTagPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritytagsApiServicesSecuritytagsTagPostOpts struct
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

# **ServicesSecuritytagsTagTagIdDelete**
> ServicesSecuritytagsTagTagIdDelete(ctx, tagId)
securityTagDeleteDelete

Delete the specified security tag.  Parameters:  tagId: Specified security tag.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritytagsTagTagIdVmDetailGet**
> ServicesSecuritytagsTagTagIdVmDetailGet(ctx, tagId)
securityTagVMDetailList

Retrieve details about the VMs that are attached to the specified security tag.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritytagsTagTagIdVmGet**
> ServicesSecuritytagsTagTagIdVmGet(ctx, tagId)
securityTagVMsListList

Retrieve the list of VMs that have the specified tag attached to them.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritytagsTagTagIdVmPost**
> ServicesSecuritytagsTagTagIdVmPost(ctx, tagId, optional)
securityTagVMsAction

Attach or detach a security tag to a virtual machine.  This operation does not check that the virtual machine exists in the local inventory. This allows you to attach a universal security tag to a virtual machine that is connected to a secondary NSX Manager (and therefore is not connected to the primary NSX Manager where the call is sent).  Possible keys for the tagParameter are: * instance_uuid * bios_uuid * vmname  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **string**|  | 
 **optional** | ***SecuritytagsApiServicesSecuritytagsTagTagIdVmPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritytagsApiServicesSecuritytagsTagTagIdVmPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **action** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritytagsTagTagIdVmVmIdDelete**
> ServicesSecuritytagsTagTagIdVmVmIdDelete(ctx, tagId, vmId)
securityTagVMDetach

Detach a security tag from the specified virtual machine.   Parameters:  vmId: Specify VM using VM managed object ID or VM instance UUID.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **string**|  | 
  **vmId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritytagsTagTagIdVmVmIdPut**
> ServicesSecuritytagsTagTagIdVmVmIdPut(ctx, tagId, vmId)
securityTagVMAttach

Apply a security tag to the specified virtual machine.  **Note:** this method can attach a universal security tag to a virtual machine. However, this method checks that the VM exists on the NSX Manager to which the API call is sent. In a cross-vCenter active active environment, the VM might exist on a secondary NSX Manager, and so the call would fail.   You can instead use the `POST /api/2.0/services/securitytags/tag/{tagId}/vm?action=attach` method to attach universal security tags to a VM that is not local to the primary NSX Manager. This method does not check that the VM is local to the NSX Manager.   Parameters:  vmId: Specify VM using VM managed object ID or VM instance UUID.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **string**|  | 
  **vmId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritytagsVmVmIdGet**
> ServicesSecuritytagsVmVmIdGet(ctx, vmId)
securitytagVmIdList

Retrieve all security tags associated with the specified virtual machine.   Parameters:  vmId: Specify VM using VM managed object ID or VM instance UUID.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritytagsVmVmIdPost**
> ServicesSecuritytagsVmVmIdPost(ctx, vmId, optional)
securitytagVmMoidAction

Update security tags associated with the specified virtual machine.  You can assign multiple tags at a time to the specified VM, or clear all assigned tags from the specified VM.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  vmId: Specify VM using VM managed object ID or VM instance UUID.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**|  | 
 **optional** | ***SecuritytagsApiServicesSecuritytagsVmVmIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritytagsApiServicesSecuritytagsVmVmIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **action** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

