# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesPolicySecurityactionCategoryVirtualmachinesGet**](PolicyApi.md#ServicesPolicySecurityactionCategoryVirtualmachinesGet) | **Get** /api/2.0/services/policy/securityaction/category/virtualmachines | securityActionVMRead
[**ServicesPolicySecuritygroupIDSecurityactionsGet**](PolicyApi.md#ServicesPolicySecuritygroupIDSecurityactionsGet) | **Get** /api/2.0/services/policy/securitygroup/{ID}/securityactions | securityActionRead
[**ServicesPolicySecuritygroupIDSecuritypoliciesGet**](PolicyApi.md#ServicesPolicySecuritygroupIDSecuritypoliciesGet) | **Get** /api/2.0/services/policy/securitygroup/{ID}/securitypolicies | secGroupPoliciesRead
[**ServicesPolicySecuritypolicyAlarmsAllGet**](PolicyApi.md#ServicesPolicySecuritypolicyAlarmsAllGet) | **Get** /api/2.0/services/policy/securitypolicy/alarms/all | securityPolicyAlarmsAllList
[**ServicesPolicySecuritypolicyHierarchyGet**](PolicyApi.md#ServicesPolicySecuritypolicyHierarchyGet) | **Get** /api/2.0/services/policy/securitypolicy/hierarchy | hierarchyRead
[**ServicesPolicySecuritypolicyHierarchyPost**](PolicyApi.md#ServicesPolicySecuritypolicyHierarchyPost) | **Post** /api/2.0/services/policy/securitypolicy/hierarchy | hierarchyCreate
[**ServicesPolicySecuritypolicyIDDelete**](PolicyApi.md#ServicesPolicySecuritypolicyIDDelete) | **Delete** /api/2.0/services/policy/securitypolicy/{ID} | securityPolicyIDDelete
[**ServicesPolicySecuritypolicyIDGet**](PolicyApi.md#ServicesPolicySecuritypolicyIDGet) | **Get** /api/2.0/services/policy/securitypolicy/{ID} | securityPolicyIDRead
[**ServicesPolicySecuritypolicyIDPut**](PolicyApi.md#ServicesPolicySecuritypolicyIDPut) | **Put** /api/2.0/services/policy/securitypolicy/{ID} | securityPolicyIDUpdate
[**ServicesPolicySecuritypolicyIDSecurityactionsGet**](PolicyApi.md#ServicesPolicySecuritypolicyIDSecurityactionsGet) | **Get** /api/2.0/services/policy/securitypolicy/{ID}/securityactions | securityActionsRead
[**ServicesPolicySecuritypolicyPost**](PolicyApi.md#ServicesPolicySecuritypolicyPost) | **Post** /api/2.0/services/policy/securitypolicy | securityPolicyCreate
[**ServicesPolicySecuritypolicyServiceproviderFirewallGet**](PolicyApi.md#ServicesPolicySecuritypolicyServiceproviderFirewallGet) | **Get** /api/2.0/services/policy/securitypolicy/serviceprovider/firewall | serviceComposerFirewallAppliedToRead
[**ServicesPolicySecuritypolicyServiceproviderFirewallPut**](PolicyApi.md#ServicesPolicySecuritypolicyServiceproviderFirewallPut) | **Put** /api/2.0/services/policy/securitypolicy/serviceprovider/firewall | serviceComposerFirewallAppliedToUpdate
[**ServicesPolicySecuritypolicyStatusGet**](PolicyApi.md#ServicesPolicySecuritypolicyStatusGet) | **Get** /api/2.0/services/policy/securitypolicy/status | securityPolicyStatusRead
[**ServicesPolicyServiceproviderFirewallGet**](PolicyApi.md#ServicesPolicyServiceproviderFirewallGet) | **Get** /api/2.0/services/policy/serviceprovider/firewall | serviceComposerDFWRead
[**ServicesPolicyVirtualmachineIDSecurityactionsGet**](PolicyApi.md#ServicesPolicyVirtualmachineIDSecurityactionsGet) | **Get** /api/2.0/services/policy/virtualmachine/{ID}/securityactions | vmApplicableSecurityActionRead

# **ServicesPolicySecurityactionCategoryVirtualmachinesGet**
> ServicesPolicySecurityactionCategoryVirtualmachinesGet(ctx, optional)
securityActionVMRead

Retrieve all VirtualMachine objects on which security action of a given category and attribute has been applied.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiServicesPolicySecurityactionCategoryVirtualmachinesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiServicesPolicySecurityactionCategoryVirtualmachinesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeKey** | **optional.String**|  | 
 **attributeValue** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPolicySecuritygroupIDSecurityactionsGet**
> ServicesPolicySecuritygroupIDSecurityactionsGet(ctx, iD)
securityActionRead

Retrieve all security actions applicable on a security group.  Retrieve all security actions applicable on a security group for all ExecutionOrderCategories. The list is sorted based on the weight of security actions in descending order.  The **isActive** tag indicates if a securityaction will be applied (by the enforcement engine) on the security group.   Parameters:  ID: Specified security group.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPolicySecuritygroupIDSecuritypoliciesGet**
> ServicesPolicySecuritygroupIDSecuritypoliciesGet(ctx, iD)
secGroupPoliciesRead

Retrieve security policies mapped to a security group.  Parameters:  ID: Specified security group ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPolicySecuritypolicyAlarmsAllGet**
> ServicesPolicySecuritypolicyAlarmsAllGet(ctx, )
securityPolicyAlarmsAllList

Retrieve all system alarms that are raised at Service Composer level and policy level.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

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

# **ServicesPolicySecuritypolicyHierarchyGet**
> ServicesPolicySecuritypolicyHierarchyGet(ctx, optional)
hierarchyRead

Export a Service Composer configuration (along with the security groups to which the security policies are mapped). You can save the response to a file.  The saved configuration can be used as a backup for situations where you may accidentally delete a policy configuration, or it can be exported for use in another NSX Manager environment.  If a prefix is specified, it is added before the names of the security policy, security action, and security group objects in the exported XML. The prefix can thus be used to indicate the remote source from where the hierarchy was exported.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiServicesPolicySecuritypolicyHierarchyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiServicesPolicySecuritypolicyHierarchyGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyIds** | **optional.String**|  | 
 **prefix** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPolicySecuritypolicyHierarchyPost**
> ServicesPolicySecuritypolicyHierarchyPost(ctx, optional)
hierarchyCreate

Import a security policy configuration  You can create multiple security policies and parent-child hierarchies using the data fetched through export. All objects including security policies, security groups and security actions are created on a global scope.  The policy that is being imported needs to be included in the request body.  If a suffix is specified, it is added after the names of the security policy, security action, and security group objects in the exported XML. The suffix can thus be used to differentiate locally created objects from imported ones.  The location of the newly created security policy objects (multiple locations are separated by commas) is populated in the Location header of the response.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiServicesPolicySecuritypolicyHierarchyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiServicesPolicySecuritypolicyHierarchyPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **suffix** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPolicySecuritypolicyIDDelete**
> ServicesPolicySecuritypolicyIDDelete(ctx, iD, optional)
securityPolicyIDDelete

Delete a security policy.  When you delete a security policy, its child security policies and all the actions in it are deleted as well.   Parameters:  ID: ID of desired security policy (or 'all' to for all security policies).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 
 **optional** | ***PolicyApiServicesPolicySecuritypolicyIDDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiServicesPolicySecuritypolicyIDDeleteOpts struct
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

# **ServicesPolicySecuritypolicyIDGet**
> ServicesPolicySecuritypolicyIDGet(ctx, iD)
securityPolicyIDRead

Retrieve security policy information.   Parameters:  ID: ID of desired security policy (or 'all' to for all security policies).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPolicySecuritypolicyIDPut**
> ServicesPolicySecuritypolicyIDPut(ctx, iD, optional)
securityPolicyIDUpdate

Edit a security policy.  To update a security policy, you must first fetch it. Then edit the received XML and pass it back as the input. The specified configuration replaces the current configuration.  Security group mappings provided in the PUT call replaces the security group mappings for the security policy. To remove all mappings, delete the securityGroupBindings parameter.  You can add or update actions for the security policy by editing the actionsByCategory parameter. To remove all actions (belonging to all categories), delete the actionsByCategory parameter. To remove actions belonging to a specific category, delete the block for that category.   Parameters:  ID: ID of desired security policy (or 'all' to for all security policies).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 
 **optional** | ***PolicyApiServicesPolicySecuritypolicyIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiServicesPolicySecuritypolicyIDPutOpts struct
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

# **ServicesPolicySecuritypolicyIDSecurityactionsGet**
> ServicesPolicySecuritypolicyIDSecurityactionsGet(ctx, iD)
securityActionsRead

Retrieve all security actions applicable on a security policy.  This list includes security actions from associated parent security policies, if any. Security actions per Execution Order Category are sorted based on the weight of security actions in descending order.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPolicySecuritypolicyPost**
> ServicesPolicySecuritypolicyPost(ctx, optional)
securityPolicyCreate

Create a security policy.  When creating a security policy, a parent security policy can be specified if required. The security policy inherits services from the parent security policy. Security group bindings and actions can also be specified while creating the policy. Note that execution order of actions in a category is implied by their order in the list. The response of the call has Location header populated with the URI using which the created object can be fetched.  Ensure that: * the required VMware built in services (such as Distributed Firewall   and Endpoint) are installed. See *NSX Installation Guide*. * the required partner services have been registered with NSX Manager. * the required security groups have been created.  Tags related to Service Composer, security policies, and security groups: Common Tags * **actionType** - Defines the type of action belonging to a given executionOrderCategory * **executionOrderCategory** - Category to which the action belongs to (endpoint, firewall or traffic_steering) * **isActive** - In a security policy hierarchy, an action within a policy may or may not be active based on the precedence of the policy or usage of isActionEnforced flag in that hierarchy * **isActionEnforced** - Enforces an action of a parent policy on its child policies for a given actionType and executionOrderCategory. Note that in a policy hierarchy, for a given actionType and executionOrderCategory, there can be only one action which can be marked as enforced. * **isEnabled** - Indicates whether an action is enabled * **secondarySecurityGroup** - Applicable for actions which need secondary security groups, say a source-destination firewall rule * **securityPolicy** - Parent policy in an action Output only Tags * **executionOrder** - Defines the sequence in which actions belonging to an executionOrderCategory are executed. Note that this is not an input parameter and its value is implied by the index in the list. Firewall Category Tags * **action** - Allow or block the traffic * **applications** - Applications / application groups on which the rules are to be applied * **direction** - Direction of traffic towards primary security group. Possible values: inbound, outbound, intra * **logged** - Flag to enable logging of the traffic that is hit by this rule * **outsideSecondaryContainer** - Flag to specify outside i.e. outside securitygroup-3 Endpoint Category Tags * **serviceId** - ID of the service (as registered with the service insertion module). If this tag is null, the functionality type (as defined in actionType tag) is not applied which will also result in blocking the actions (of given functionality type) that are inherited from the parent security policy. This is true if there is no action of enforce type. * **invalidServiceId** - Flag to indicate that the service that was referenced in this rule is deleted, which make the rule ineffective (or deviate from the original intent that existed while configuring the rule). You must either modify this rule by adding correct Service or delete this rule. * **serviceName** -Name of the service * **serviceProfile** - Profile to be referenced in Endpoint rule. * **invalidServiceProfile** - Flag to indicate that the service profile that was referenced in this rule is deleted, which makes the rule ineffective (or deviate from the original intent that existed while configuring the rule). You must either modify this rule by adding correct Service Profile or delete this rule. The following tags are deprecated: * **vendorTemplateId** * **invalidVendorTemplateId** * **vendorTemplateName** Traffic Steering/NetX Category Tags * **redirect** - Flag to indicate whether to redirect the traffic or not * **serviceProfile** - Service profile for which redirection is being configured * **logged** - Flag to enable logging of the traffic that is hit by this rule   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiServicesPolicySecuritypolicyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiServicesPolicySecuritypolicyPostOpts struct
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

# **ServicesPolicySecuritypolicyServiceproviderFirewallGet**
> ServicesPolicySecuritypolicyServiceproviderFirewallGet(ctx, )
serviceComposerFirewallAppliedToRead

Retrieve the Service Composer firewall applied to setting.   Parameters:  

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

# **ServicesPolicySecuritypolicyServiceproviderFirewallPut**
> ServicesPolicySecuritypolicyServiceproviderFirewallPut(ctx, optional)
serviceComposerFirewallAppliedToUpdate

Update the Service Composer firewall applied to setting.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiServicesPolicySecuritypolicyServiceproviderFirewallPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiServicesPolicySecuritypolicyServiceproviderFirewallPutOpts struct
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

# **ServicesPolicySecuritypolicyStatusGet**
> ServicesPolicySecuritypolicyStatusGet(ctx, )
securityPolicyStatusRead

Retrieve the consolidated status of Service Composer.  The possible return of value for status are: *in_sync*, *in_progress*, *out_of_sync*, and *pending*.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

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

# **ServicesPolicyServiceproviderFirewallGet**
> ServicesPolicyServiceproviderFirewallGet(ctx, optional)
serviceComposerDFWRead

If Service Composer goes out of sync with Distributed Firewall, you must re-synchronize Service Composer rules with firewall rules. If Service Composer stays out of sync, firewall configuration may not stay enforced as expected.  This GET method can perform the following functions, depending on the request body provided. **Note:** Some REST clients do not allow you to specify a request body with a GET request.  ### Check if Service Composer firewall and Distributed Firewall are in sync  **Note: Deprecated.** Use `GET /2.0/services/policy/securitypolicy/status` instead.   * If they are in sync, the response body does not contain any data.   * If they are out of sync, the response body contains the unix timestamp representing the time since when Service Composer firewall is out of sync.  ``` <keyValues>   <keyValue>     <key>getServiceComposerFirewallOutOfSyncTimestamp</key>   </keyValue> </keyValues> ```  ### Synchronize Service Composer firewall with Distributed Firewall  ``` <keyValues>   <keyValue>     <key>forceSync</key>   </keyValue> </keyValues> ```  ### Retrieve the state of the auto save draft property in Service Composer  Retrieve the state of the auto save draft property in Service Composer. Response is true or false.  ``` <keyValues>   <keyValue>     <key>getAutoSaveDraft</key>   </keyValue> </keyValues> ```  ### Change the state of the auto save draft property in Service Composer  **Note: Deprecated.**  Change the state of the auto save draft property in Service Composer. Provide request body value of true or false.  ``` <keyValues>   <keyValue>     <key>autoSaveDraft</key>     <value>false</value>   </keyValue> </keyValues>  ```  **Method history:**    Release | Modification   --------|-------------   6.2.3 | Method updated and some functions deprecated. Changing auto save draft with the **autoSaveDraft** parameter is deprecated, and will be removed in a future release.  <br>The default setting of **autoSaveDraft** is changed from *true* to *false*.<br>Method to check if Service Composer and Distributed Firewall are in sync is deprecated, and will be removed in a future release. Use `GET /2.0/services/policy/securitypolicy/status` instead.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiServicesPolicyServiceproviderFirewallGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiServicesPolicyServiceproviderFirewallGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesPolicyVirtualmachineIDSecurityactionsGet**
> ServicesPolicyVirtualmachineIDSecurityactionsGet(ctx, iD)
vmApplicableSecurityActionRead

Retrieve the security actions applicable on a virtual machine.   Parameters:  ID: VM ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

