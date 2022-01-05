# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesSecuritygroupBulkObjectIdPut**](SecuritygroupApi.md#ServicesSecuritygroupBulkObjectIdPut) | **Put** /api/2.0/services/securitygroup/bulk/{objectId} | secGroupBulkObjectUpdate
[**ServicesSecuritygroupBulkScopeIdPost**](SecuritygroupApi.md#ServicesSecuritygroupBulkScopeIdPost) | **Post** /api/2.0/services/securitygroup/bulk/{scopeId} | secGroupBulkCreate
[**ServicesSecuritygroupInternalScopeScopeIdGet**](SecuritygroupApi.md#ServicesSecuritygroupInternalScopeScopeIdGet) | **Get** /api/2.0/services/securitygroup/internal/scope/{scopeId} | secGroupInternalRead
[**ServicesSecuritygroupLookupVirtualmachineVirtualMachineIdGet**](SecuritygroupApi.md#ServicesSecuritygroupLookupVirtualmachineVirtualMachineIdGet) | **Get** /api/2.0/services/securitygroup/lookup/virtualmachine/{virtualMachineId} | secGroupLookupVMRead
[**ServicesSecuritygroupObjectIdDelete**](SecuritygroupApi.md#ServicesSecuritygroupObjectIdDelete) | **Delete** /api/2.0/services/securitygroup/{objectId} | secGroupObjectDelete
[**ServicesSecuritygroupObjectIdGet**](SecuritygroupApi.md#ServicesSecuritygroupObjectIdGet) | **Get** /api/2.0/services/securitygroup/{objectId} | secGroupObjectRead
[**ServicesSecuritygroupObjectIdMembersMemberIdDelete**](SecuritygroupApi.md#ServicesSecuritygroupObjectIdMembersMemberIdDelete) | **Delete** /api/2.0/services/securitygroup/{objectId}/members/{memberId} | secGroupMemberDelete
[**ServicesSecuritygroupObjectIdMembersMemberIdPut**](SecuritygroupApi.md#ServicesSecuritygroupObjectIdMembersMemberIdPut) | **Put** /api/2.0/services/securitygroup/{objectId}/members/{memberId} | secGroupMemberUpdate
[**ServicesSecuritygroupObjectIdPut**](SecuritygroupApi.md#ServicesSecuritygroupObjectIdPut) | **Put** /api/2.0/services/securitygroup/{objectId} | secGroupObjectUpdate
[**ServicesSecuritygroupObjectIdTranslationIpaddressesGet**](SecuritygroupApi.md#ServicesSecuritygroupObjectIdTranslationIpaddressesGet) | **Get** /api/2.0/services/securitygroup/{objectId}/translation/ipaddresses | secGroupIPNodesRead
[**ServicesSecuritygroupObjectIdTranslationMacaddressesGet**](SecuritygroupApi.md#ServicesSecuritygroupObjectIdTranslationMacaddressesGet) | **Get** /api/2.0/services/securitygroup/{objectId}/translation/macaddresses | secGroupMacNodesRead
[**ServicesSecuritygroupObjectIdTranslationVirtualmachinesGet**](SecuritygroupApi.md#ServicesSecuritygroupObjectIdTranslationVirtualmachinesGet) | **Get** /api/2.0/services/securitygroup/{objectId}/translation/virtualmachines | secGroupVMNodesRead
[**ServicesSecuritygroupObjectIdTranslationVnicsGet**](SecuritygroupApi.md#ServicesSecuritygroupObjectIdTranslationVnicsGet) | **Get** /api/2.0/services/securitygroup/{objectId}/translation/vnics | secGroupVnicNodesRead
[**ServicesSecuritygroupScopeIdPost**](SecuritygroupApi.md#ServicesSecuritygroupScopeIdPost) | **Post** /api/2.0/services/securitygroup/{scopeId} | secGroupScopeIdCreate
[**ServicesSecuritygroupScopeScopeIdGet**](SecuritygroupApi.md#ServicesSecuritygroupScopeScopeIdGet) | **Get** /api/2.0/services/securitygroup/scope/{scopeId} | secGroupScopeRead
[**ServicesSecuritygroupScopeScopeIdMemberTypesGet**](SecuritygroupApi.md#ServicesSecuritygroupScopeScopeIdMemberTypesGet) | **Get** /api/2.0/services/securitygroup/scope/{scopeId}/memberTypes | secGroupScopeMembersRead
[**ServicesSecuritygroupScopeScopeIdMembersMemberTypeGet**](SecuritygroupApi.md#ServicesSecuritygroupScopeScopeIdMembersMemberTypeGet) | **Get** /api/2.0/services/securitygroup/scope/{scopeId}/members/{memberType} | secGroupScopeMemberTypeRead

# **ServicesSecuritygroupBulkObjectIdPut**
> ServicesSecuritygroupBulkObjectIdPut(ctx, objectId, optional)
secGroupBulkObjectUpdate

Update configuration for the specified security group, including membership information.   Parameters:  objectId: Security group ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 
 **optional** | ***SecuritygroupApiServicesSecuritygroupBulkObjectIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritygroupApiServicesSecuritygroupBulkObjectIdPutOpts struct
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

# **ServicesSecuritygroupBulkScopeIdPost**
> ServicesSecuritygroupBulkScopeIdPost(ctx, scopeId, optional)
secGroupBulkCreate

Create a new security group on a global scope or universal scope with membership information.  Universal security groups are read-only when querying a secondary NSX manager.  When you create a universal security group (on scope *universalroot-0*) by default **localMembersOnly** is set to *false* which indicates that the universal security group will contain members across the cross-vCenter NSX environment.  This is the case in an active active environment. You can add the following objects to a universal security group with *localMembersOnly=false* (active active): * IP Address Set * MAC Address Set * Universal Security Groups with *localMembersOnly=false*  When you create a universal security group (on scope *universalroot-0*) you can set the extendedAttribute **localMembersOnly** to *true* to indicate that the universal security group will contain members local to that NSX Manager only.  This is the case in an active standby environment, because only one NSX environment is active at a time, and the same VMs are present in each NSX environment. You can add the following objects to a universal security group with *localMembersOnly=true* (active standby): * Universal Security Tag * IP Address Set * MAC Address Set * Universal Security Groups with *localMembersOnly=true* * Dynamic criteria using VM name  You can set the **localMembersOnly** attribute only when the universal security group is created, it cannot be modified afterwards.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Extended attribute **localMembersOnly** introduced.   Parameters:  scopeId: For the scopeId use *globalroot-0* for non-universal security groups and *universalroot-0* for universal security groups.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***SecuritygroupApiServicesSecuritygroupBulkScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritygroupApiServicesSecuritygroupBulkScopeIdPostOpts struct
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

# **ServicesSecuritygroupInternalScopeScopeIdGet**
> ServicesSecuritygroupInternalScopeScopeIdGet(ctx, scopeId)
secGroupInternalRead

Retrieve all internal security groups on the NSX Manager. These are used  internally by the system and should not be created or modified by end users.   Parameters:  scopeId: Specified transport zone (scope)  

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

# **ServicesSecuritygroupLookupVirtualmachineVirtualMachineIdGet**
> ServicesSecuritygroupLookupVirtualmachineVirtualMachineIdGet(ctx, virtualMachineId)
secGroupLookupVMRead

Retrieve list of security groups that the specified virtual machine belongs to.   Parameters:  virtualMachineId: Specified virtual machine  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualMachineId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritygroupObjectIdDelete**
> ServicesSecuritygroupObjectIdDelete(ctx, objectId, optional)
secGroupObjectDelete

Delete an existing security group.  If *force=true* is specified, the object is deleted even if used in other configurations, such as firewall rules. If *force=true* is not specified, the object is deleted only if it is not used by other configuration; otherwise the delete fails.   Parameters:  objectId: Security group ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 
 **optional** | ***SecuritygroupApiServicesSecuritygroupObjectIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritygroupApiServicesSecuritygroupObjectIdDeleteOpts struct
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

# **ServicesSecuritygroupObjectIdGet**
> ServicesSecuritygroupObjectIdGet(ctx, objectId)
secGroupObjectRead

Retrieve all members of the specified security group.  Parameters:  objectId: Security group ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritygroupObjectIdMembersMemberIdDelete**
> ServicesSecuritygroupObjectIdMembersMemberIdDelete(ctx, objectId, memberId, optional)
secGroupMemberDelete

Delete member from the specified security group.  Parameters:  memberId: Security group member, can be a vSphere managed object ID or NSX object ID.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 
  **memberId** | **string**|  | 
 **optional** | ***SecuritygroupApiServicesSecuritygroupObjectIdMembersMemberIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritygroupApiServicesSecuritygroupObjectIdMembersMemberIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **failIfAbsent** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritygroupObjectIdMembersMemberIdPut**
> ServicesSecuritygroupObjectIdMembersMemberIdPut(ctx, objectId, memberId, optional)
secGroupMemberUpdate

Add a new member to the specified security group.   Parameters:  memberId: Security group member, can be a vSphere managed object ID or NSX object ID.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 
  **memberId** | **string**|  | 
 **optional** | ***SecuritygroupApiServicesSecuritygroupObjectIdMembersMemberIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritygroupApiServicesSecuritygroupObjectIdMembersMemberIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **failIfExists** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritygroupObjectIdPut**
> ServicesSecuritygroupObjectIdPut(ctx, objectId, optional)
secGroupObjectUpdate

Update configuration for the specified security group. Members are not updated. You must use `PUT /2.0/services/securitygroup/bulk/{objectId}` to update a security group membership.   Parameters:  objectId: Security group ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 
 **optional** | ***SecuritygroupApiServicesSecuritygroupObjectIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritygroupApiServicesSecuritygroupObjectIdPutOpts struct
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

# **ServicesSecuritygroupObjectIdTranslationIpaddressesGet**
> ServicesSecuritygroupObjectIdTranslationIpaddressesGet(ctx, objectId)
secGroupIPNodesRead

Retrieve list of IP addresses that belong to a specific security group.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritygroupObjectIdTranslationMacaddressesGet**
> ServicesSecuritygroupObjectIdTranslationMacaddressesGet(ctx, objectId)
secGroupMacNodesRead

Retrieve list of MAC addresses that belong to a specific security group.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritygroupObjectIdTranslationVirtualmachinesGet**
> ServicesSecuritygroupObjectIdTranslationVirtualmachinesGet(ctx, objectId)
secGroupVMNodesRead

Retrieve list of virtual machine entities that belong to a specific security group.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritygroupObjectIdTranslationVnicsGet**
> ServicesSecuritygroupObjectIdTranslationVnicsGet(ctx, objectId)
secGroupVnicNodesRead

Retrieve list of vNICs that belong to a specific security group.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSecuritygroupScopeIdPost**
> ServicesSecuritygroupScopeIdPost(ctx, scopeId, optional)
secGroupScopeIdCreate

Create a new security group, with no membership information specified. You can add members later with `PUT /2.0/services/securitygroup/bulk/{objectId}`  When you create a universal security group (on scope *universalroot-0*) by default **localMembersOnly** is set to *false* which indicates that the universal security group will contain members across the cross-vCenter NSX environment.  This is the case in an active active environment. You can add the following objects to a universal security group with *localMembersOnly=false* (active active): * IP Address Set * MAC Address Set * Universal Security Groups with *localMembersOnly=false*  When you create a universal security group (on scope *universalroot-0*) you can set the extendedAttribute **localMembersOnly** to *true* to indicate that the universal security group will contain members local to that NSX Manager only.  This is the case in an active standby environment, because only one NSX environment is active at a time, and the same VMs are present in each NSX environment. You can add the following objects to a universal security group with *localMembersOnly=true* (active standby): * Universal Security Tag * IP Address Set * MAC Address Set * Universal Security Groups with *localMembersOnly=true* * Dynamic criteria using VM name  You can set the **localMembersOnly** attribute only when the universal security group is created, it cannot be modified afterwards.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Extended attribute **localMembersOnly** introduced.   Parameters:  scopeId: For the scopeId use *globalroot-0* for non-universal security groups and *universalroot-0* for universal security groups.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***SecuritygroupApiServicesSecuritygroupScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecuritygroupApiServicesSecuritygroupScopeIdPostOpts struct
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

# **ServicesSecuritygroupScopeScopeIdGet**
> ServicesSecuritygroupScopeScopeIdGet(ctx, scopeId)
secGroupScopeRead

List all the security groups created on a specific scope.  Parameters:  scopeId: scopeId can be \"globalroot-0\", \"universalroot-0\" or datacenterID / portgroupID in upgrade use cases   

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

# **ServicesSecuritygroupScopeScopeIdMemberTypesGet**
> ServicesSecuritygroupScopeScopeIdMemberTypesGet(ctx, scopeId)
secGroupScopeMembersRead

Retrieve a list of valid elements that can be added to a security group.   Parameters:  

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

# **ServicesSecuritygroupScopeScopeIdMembersMemberTypeGet**
> ServicesSecuritygroupScopeScopeIdMembersMemberTypeGet(ctx, scopeId, memberType)
secGroupScopeMemberTypeRead

Retrieve members of a specific type in the specified scope.  Parameters:  memberType: Specific member type  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
  **memberType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

