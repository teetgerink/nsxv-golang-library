# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesSecuritytagsSelectionCriteriaGet**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsSelectionCriteriaGet) | **Get** /api/2.0/services/securitytags/selection-criteria | securitytagsCriteriaRead
[**Api20ServicesSecuritytagsSelectionCriteriaPut**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsSelectionCriteriaPut) | **Put** /api/2.0/services/securitytags/selection-criteria | securitytagsCriteriaUpdate
[**Api20ServicesSecuritytagsTagGet**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsTagGet) | **Get** /api/2.0/services/securitytags/tag | securityTagRead
[**Api20ServicesSecuritytagsTagPost**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsTagPost) | **Post** /api/2.0/services/securitytags/tag | securityTagCreate
[**Api20ServicesSecuritytagsTagTagIdDelete**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsTagTagIdDelete) | **Delete** /api/2.0/services/securitytags/tag/{tagId} | securityTagDeleteDelete
[**Api20ServicesSecuritytagsTagTagIdVmDetailGet**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsTagTagIdVmDetailGet) | **Get** /api/2.0/services/securitytags/tag/{tagId}/vmDetail | securityTagVMDetailList
[**Api20ServicesSecuritytagsTagTagIdVmGet**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsTagTagIdVmGet) | **Get** /api/2.0/services/securitytags/tag/{tagId}/vm | securityTagVMsListList
[**Api20ServicesSecuritytagsTagTagIdVmPost**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsTagTagIdVmPost) | **Post** /api/2.0/services/securitytags/tag/{tagId}/vm | securityTagVMsAction
[**Api20ServicesSecuritytagsTagTagIdVmVmIdDelete**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsTagTagIdVmVmIdDelete) | **Delete** /api/2.0/services/securitytags/tag/{tagId}/vm/{vmId} | securityTagVMDetach
[**Api20ServicesSecuritytagsTagTagIdVmVmIdPut**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsTagTagIdVmVmIdPut) | **Put** /api/2.0/services/securitytags/tag/{tagId}/vm/{vmId} | securityTagVMAttach
[**Api20ServicesSecuritytagsVmVmIdGet**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsVmVmIdGet) | **Get** /api/2.0/services/securitytags/vm/{vmId} | securitytagVmIdList
[**Api20ServicesSecuritytagsVmVmIdPost**](Class20SecuritytagsApi.md#Api20ServicesSecuritytagsVmVmIdPost) | **Post** /api/2.0/services/securitytags/vm/{vmId} | securitytagVmMoidAction

# **Api20ServicesSecuritytagsSelectionCriteriaGet**
> Api20ServicesSecuritytagsSelectionCriteriaGet(ctx, )
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

# **Api20ServicesSecuritytagsSelectionCriteriaPut**
> Api20ServicesSecuritytagsSelectionCriteriaPut(ctx, optional)
securitytagsCriteriaUpdate

Configure the unique ID section criteria configuration.  If you set the selection criteria and assign security tags to VMs, you must remove all security tags from VMs before you can change the selection criteria.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20SecuritytagsApiApi20ServicesSecuritytagsSelectionCriteriaPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20SecuritytagsApiApi20ServicesSecuritytagsSelectionCriteriaPutOpts struct
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

# **Api20ServicesSecuritytagsTagGet**
> Api20ServicesSecuritytagsTagGet(ctx, optional)
securityTagRead

Retrieve all security tags.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method updated. Added **isUniversal** query parameter to filter universal security tags.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20SecuritytagsApiApi20ServicesSecuritytagsTagGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20SecuritytagsApiApi20ServicesSecuritytagsTagGetOpts struct
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

# **Api20ServicesSecuritytagsTagPost**
> Api20ServicesSecuritytagsTagPost(ctx, optional)
securityTagCreate

Create a new security tag.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method updated. **isUniversal** parameter can be set to create a universal security tag.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class20SecuritytagsApiApi20ServicesSecuritytagsTagPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20SecuritytagsApiApi20ServicesSecuritytagsTagPostOpts struct
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

# **Api20ServicesSecuritytagsTagTagIdDelete**
> Api20ServicesSecuritytagsTagTagIdDelete(ctx, tagId)
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

# **Api20ServicesSecuritytagsTagTagIdVmDetailGet**
> Api20ServicesSecuritytagsTagTagIdVmDetailGet(ctx, tagId)
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

# **Api20ServicesSecuritytagsTagTagIdVmGet**
> Api20ServicesSecuritytagsTagTagIdVmGet(ctx, tagId)
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

# **Api20ServicesSecuritytagsTagTagIdVmPost**
> Api20ServicesSecuritytagsTagTagIdVmPost(ctx, tagId, optional)
securityTagVMsAction

Attach or detach a security tag to a virtual machine.  This operation does not check that the virtual machine exists in the local inventory. This allows you to attach a universal security tag to a virtual machine that is connected to a secondary NSX Manager (and therefore is not connected to the primary NSX Manager where the call is sent).  Possible keys for the tagParameter are: * instance_uuid * bios_uuid * vmname  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **string**|  | 
 **optional** | ***Class20SecuritytagsApiApi20ServicesSecuritytagsTagTagIdVmPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20SecuritytagsApiApi20ServicesSecuritytagsTagTagIdVmPostOpts struct
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

# **Api20ServicesSecuritytagsTagTagIdVmVmIdDelete**
> Api20ServicesSecuritytagsTagTagIdVmVmIdDelete(ctx, tagId, vmId)
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

# **Api20ServicesSecuritytagsTagTagIdVmVmIdPut**
> Api20ServicesSecuritytagsTagTagIdVmVmIdPut(ctx, tagId, vmId)
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

# **Api20ServicesSecuritytagsVmVmIdGet**
> Api20ServicesSecuritytagsVmVmIdGet(ctx, vmId)
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

# **Api20ServicesSecuritytagsVmVmIdPost**
> Api20ServicesSecuritytagsVmVmIdPost(ctx, vmId, optional)
securitytagVmMoidAction

Update security tags associated with the specified virtual machine.  You can assign multiple tags at a time to the specified VM, or clear all assigned tags from the specified VM.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  vmId: Specify VM using VM managed object ID or VM instance UUID.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**|  | 
 **optional** | ***Class20SecuritytagsApiApi20ServicesSecuritytagsVmVmIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20SecuritytagsApiApi20ServicesSecuritytagsVmVmIdPostOpts struct
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

