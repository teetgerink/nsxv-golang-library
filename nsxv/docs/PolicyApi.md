# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesPolicySecurityactionCategoryVirtualmachinesGet**](PolicyApi.md#Api20ServicesPolicySecurityactionCategoryVirtualmachinesGet) | **Get** /api/2.0/services/policy/securityaction/category/virtualmachines | securityActionVMRead
[**Api20ServicesPolicySecuritygroupIDSecurityactionsGet**](PolicyApi.md#Api20ServicesPolicySecuritygroupIDSecurityactionsGet) | **Get** /api/2.0/services/policy/securitygroup/{ID}/securityactions | securityActionRead
[**Api20ServicesPolicySecuritygroupIDSecuritypoliciesGet**](PolicyApi.md#Api20ServicesPolicySecuritygroupIDSecuritypoliciesGet) | **Get** /api/2.0/services/policy/securitygroup/{ID}/securitypolicies | secGroupPoliciesRead
[**Api20ServicesPolicySecuritypolicyAlarmsAllGet**](PolicyApi.md#Api20ServicesPolicySecuritypolicyAlarmsAllGet) | **Get** /api/2.0/services/policy/securitypolicy/alarms/all | securityPolicyAlarmsAllList
[**Api20ServicesPolicySecuritypolicyHierarchyGet**](PolicyApi.md#Api20ServicesPolicySecuritypolicyHierarchyGet) | **Get** /api/2.0/services/policy/securitypolicy/hierarchy | hierarchyRead
[**Api20ServicesPolicySecuritypolicyHierarchyPost**](PolicyApi.md#Api20ServicesPolicySecuritypolicyHierarchyPost) | **Post** /api/2.0/services/policy/securitypolicy/hierarchy | hierarchyCreate
[**Api20ServicesPolicySecuritypolicyIDDelete**](PolicyApi.md#Api20ServicesPolicySecuritypolicyIDDelete) | **Delete** /api/2.0/services/policy/securitypolicy/{ID} | securityPolicyIDDelete
[**Api20ServicesPolicySecuritypolicyIDGet**](PolicyApi.md#Api20ServicesPolicySecuritypolicyIDGet) | **Get** /api/2.0/services/policy/securitypolicy/{ID} | securityPolicyIDRead
[**Api20ServicesPolicySecuritypolicyIDPut**](PolicyApi.md#Api20ServicesPolicySecuritypolicyIDPut) | **Put** /api/2.0/services/policy/securitypolicy/{ID} | securityPolicyIDUpdate
[**Api20ServicesPolicySecuritypolicyIDSecurityactionsGet**](PolicyApi.md#Api20ServicesPolicySecuritypolicyIDSecurityactionsGet) | **Get** /api/2.0/services/policy/securitypolicy/{ID}/securityactions | securityActionsRead
[**Api20ServicesPolicySecuritypolicyPost**](PolicyApi.md#Api20ServicesPolicySecuritypolicyPost) | **Post** /api/2.0/services/policy/securitypolicy | securityPolicyCreate
[**Api20ServicesPolicySecuritypolicyServiceproviderFirewallGet**](PolicyApi.md#Api20ServicesPolicySecuritypolicyServiceproviderFirewallGet) | **Get** /api/2.0/services/policy/securitypolicy/serviceprovider/firewall | serviceComposerFirewallAppliedToRead
[**Api20ServicesPolicySecuritypolicyServiceproviderFirewallPut**](PolicyApi.md#Api20ServicesPolicySecuritypolicyServiceproviderFirewallPut) | **Put** /api/2.0/services/policy/securitypolicy/serviceprovider/firewall | serviceComposerFirewallAppliedToUpdate
[**Api20ServicesPolicySecuritypolicyStatusGet**](PolicyApi.md#Api20ServicesPolicySecuritypolicyStatusGet) | **Get** /api/2.0/services/policy/securitypolicy/status | securityPolicyStatusRead
[**Api20ServicesPolicyServiceproviderFirewallGet**](PolicyApi.md#Api20ServicesPolicyServiceproviderFirewallGet) | **Get** /api/2.0/services/policy/serviceprovider/firewall | serviceComposerDFWRead
[**Api20ServicesPolicyVirtualmachineIDSecurityactionsGet**](PolicyApi.md#Api20ServicesPolicyVirtualmachineIDSecurityactionsGet) | **Get** /api/2.0/services/policy/virtualmachine/{ID}/securityactions | vmApplicableSecurityActionRead

# **Api20ServicesPolicySecurityactionCategoryVirtualmachinesGet**
> Api20ServicesPolicySecurityactionCategoryVirtualmachinesGet(ctx, optional)
securityActionVMRead

Retrieve all VirtualMachine objects on which security action of a given category and attribute has been applied.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiApi20ServicesPolicySecurityactionCategoryVirtualmachinesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiApi20ServicesPolicySecurityactionCategoryVirtualmachinesGetOpts struct
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

# **Api20ServicesPolicySecuritygroupIDSecurityactionsGet**
> Api20ServicesPolicySecuritygroupIDSecurityactionsGet(ctx, iD)
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

# **Api20ServicesPolicySecuritygroupIDSecuritypoliciesGet**
> Api20ServicesPolicySecuritygroupIDSecuritypoliciesGet(ctx, iD)
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

# **Api20ServicesPolicySecuritypolicyAlarmsAllGet**
> Api20ServicesPolicySecuritypolicyAlarmsAllGet(ctx, )
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

# **Api20ServicesPolicySecuritypolicyHierarchyGet**
> Api20ServicesPolicySecuritypolicyHierarchyGet(ctx, optional)
hierarchyRead

Export a Service Composer configuration (along with the security groups to which the security policies are mapped). You can save the response to a file.  The saved configuration can be used as a backup for situations where you may accidentally delete a policy configuration, or it can be exported for use in another NSX Manager environment.  If a prefix is specified, it is added before the names of the security policy, security action, and security group objects in the exported XML. The prefix can thus be used to indicate the remote source from where the hierarchy was exported.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiApi20ServicesPolicySecuritypolicyHierarchyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiApi20ServicesPolicySecuritypolicyHierarchyGetOpts struct
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

# **Api20ServicesPolicySecuritypolicyHierarchyPost**
> Api20ServicesPolicySecuritypolicyHierarchyPost(ctx, optional)
hierarchyCreate

Import a security policy configuration  You can create multiple security policies and parent-child hierarchies using the data fetched through export. All objects including security policies, security groups and security actions are created on a global scope.  The policy that is being imported needs to be included in the request body.  If a suffix is specified, it is added after the names of the security policy, security action, and security group objects in the exported XML. The suffix can thus be used to differentiate locally created objects from imported ones.  The location of the newly created security policy objects (multiple locations are separated by commas) is populated in the Location header of the response.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiApi20ServicesPolicySecuritypolicyHierarchyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiApi20ServicesPolicySecuritypolicyHierarchyPostOpts struct
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

# **Api20ServicesPolicySecuritypolicyIDDelete**
> Api20ServicesPolicySecuritypolicyIDDelete(ctx, iD, optional)
securityPolicyIDDelete

Delete a security policy.  When you delete a security policy, its child security policies and all the actions in it are deleted as well.   Parameters:  ID: ID of desired security policy (or 'all' to for all security policies).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 
 **optional** | ***PolicyApiApi20ServicesPolicySecuritypolicyIDDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiApi20ServicesPolicySecuritypolicyIDDeleteOpts struct
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

# **Api20ServicesPolicySecuritypolicyIDGet**
> Api20ServicesPolicySecuritypolicyIDGet(ctx, iD)
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

# **Api20ServicesPolicySecuritypolicyIDPut**
> Api20ServicesPolicySecuritypolicyIDPut(ctx, iD, optional)
securityPolicyIDUpdate

Edit a security policy.  To update a security policy, you must first fetch it. Then edit the received XML and pass it back as the input. The specified configuration replaces the current configuration.  Security group mappings provided in the PUT call replaces the security group mappings for the security policy. To remove all mappings, delete the securityGroupBindings parameter.  You can add or update actions for the security policy by editing the actionsByCategory parameter. To remove all actions (belonging to all categories), delete the actionsByCategory parameter. To remove actions belonging to a specific category, delete the block for that category.   Parameters:  ID: ID of desired security policy (or 'all' to for all security policies).   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 
 **optional** | ***PolicyApiApi20ServicesPolicySecuritypolicyIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiApi20ServicesPolicySecuritypolicyIDPutOpts struct
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

# **Api20ServicesPolicySecuritypolicyIDSecurityactionsGet**
> Api20ServicesPolicySecuritypolicyIDSecurityactionsGet(ctx, iD)
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

# **Api20ServicesPolicySecuritypolicyPost**
> Api20ServicesPolicySecuritypolicyPost(ctx, optional)
securityPolicyCreate

Create a security policy.  When creating a security policy, a parent security policy can be specified if required. The security policy inherits services from the parent security policy. Security group bindings and actions can also be specified while creating the policy. Note that execution order of actions in a category is implied by their order in the list. The response of the call has Location header populated with the URI using which the created object can be fetched.  Ensure that: * the required VMware built in services (such as Distributed Firewall   and Endpoint) are installed. See *NSX Installation Guide*. * the required partner services have been registered with NSX Manager. * the required security groups have been created.  Tags related to Service Composer, security policies, and security groups: Common Tags * **actionType** - Defines the type of action belonging to a given executionOrderCategory * **executionOrderCategory** - Category to which the action belongs to (endpoint, firewall or traffic_steering) * **isActive** - In a security policy hierarchy, an action within a policy may or may not be active based on the precedence of the policy or usage of isActionEnforced flag in that hierarchy * **isActionEnforced** - Enforces an action of a parent policy on its child policies for a given actionType and executionOrderCategory. Note that in a policy hierarchy, for a given actionType and executionOrderCategory, there can be only one action which can be marked as enforced. * **isEnabled** - Indicates whether an action is enabled * **secondarySecurityGroup** - Applicable for actions which need secondary security groups, say a source-destination firewall rule * **securityPolicy** - Parent policy in an action Output only Tags * **executionOrder** - Defines the sequence in which actions belonging to an executionOrderCategory are executed. Note that this is not an input parameter and its value is implied by the index in the list. Firewall Category Tags * **action** - Allow or block the traffic * **applications** - Applications / application groups on which the rules are to be applied * **direction** - Direction of traffic towards primary security group. Possible values: inbound, outbound, intra * **logged** - Flag to enable logging of the traffic that is hit by this rule * **outsideSecondaryContainer** - Flag to specify outside i.e. outside securitygroup-3 Endpoint Category Tags * **serviceId** - ID of the service (as registered with the service insertion module). If this tag is null, the functionality type (as defined in actionType tag) is not applied which will also result in blocking the actions (of given functionality type) that are inherited from the parent security policy. This is true if there is no action of enforce type. * **invalidServiceId** - Flag to indicate that the service that was referenced in this rule is deleted, which make the rule ineffective (or deviate from the original intent that existed while configuring the rule). You must either modify this rule by adding correct Service or delete this rule. * **serviceName** -Name of the service * **serviceProfile** - Profile to be referenced in Endpoint rule. * **invalidServiceProfile** - Flag to indicate that the service profile that was referenced in this rule is deleted, which makes the rule ineffective (or deviate from the original intent that existed while configuring the rule). You must either modify this rule by adding correct Service Profile or delete this rule. The following tags are deprecated: * **vendorTemplateId** * **invalidVendorTemplateId** * **vendorTemplateName** Traffic Steering/NetX Category Tags * **redirect** - Flag to indicate whether to redirect the traffic or not * **serviceProfile** - Service profile for which redirection is being configured * **logged** - Flag to enable logging of the traffic that is hit by this rule   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiApi20ServicesPolicySecuritypolicyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiApi20ServicesPolicySecuritypolicyPostOpts struct
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

# **Api20ServicesPolicySecuritypolicyServiceproviderFirewallGet**
> Api20ServicesPolicySecuritypolicyServiceproviderFirewallGet(ctx, )
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

# **Api20ServicesPolicySecuritypolicyServiceproviderFirewallPut**
> Api20ServicesPolicySecuritypolicyServiceproviderFirewallPut(ctx, optional)
serviceComposerFirewallAppliedToUpdate

Update the Service Composer firewall applied to setting.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiApi20ServicesPolicySecuritypolicyServiceproviderFirewallPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiApi20ServicesPolicySecuritypolicyServiceproviderFirewallPutOpts struct
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

# **Api20ServicesPolicySecuritypolicyStatusGet**
> Api20ServicesPolicySecuritypolicyStatusGet(ctx, )
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

# **Api20ServicesPolicyServiceproviderFirewallGet**
> Api20ServicesPolicyServiceproviderFirewallGet(ctx, optional)
serviceComposerDFWRead

If Service Composer goes out of sync with Distributed Firewall, you must re-synchronize Service Composer rules with firewall rules. If Service Composer stays out of sync, firewall configuration may not stay enforced as expected.  This GET method can perform the following functions, depending on the request body provided. **Note:** Some REST clients do not allow you to specify a request body with a GET request.  ### Check if Service Composer firewall and Distributed Firewall are in sync  **Note: Deprecated.** Use `GET /2.0/services/policy/securitypolicy/status` instead.   * If they are in sync, the response body does not contain any data.   * If they are out of sync, the response body contains the unix timestamp representing the time since when Service Composer firewall is out of sync.  ``` <keyValues>   <keyValue>     <key>getServiceComposerFirewallOutOfSyncTimestamp</key>   </keyValue> </keyValues> ```  ### Synchronize Service Composer firewall with Distributed Firewall  ``` <keyValues>   <keyValue>     <key>forceSync</key>   </keyValue> </keyValues> ```  ### Retrieve the state of the auto save draft property in Service Composer  Retrieve the state of the auto save draft property in Service Composer. Response is true or false.  ``` <keyValues>   <keyValue>     <key>getAutoSaveDraft</key>   </keyValue> </keyValues> ```  ### Change the state of the auto save draft property in Service Composer  **Note: Deprecated.**  Change the state of the auto save draft property in Service Composer. Provide request body value of true or false.  ``` <keyValues>   <keyValue>     <key>autoSaveDraft</key>     <value>false</value>   </keyValue> </keyValues>  ```  **Method history:**    Release | Modification   --------|-------------   6.2.3 | Method updated and some functions deprecated. Changing auto save draft with the **autoSaveDraft** parameter is deprecated, and will be removed in a future release.  <br>The default setting of **autoSaveDraft** is changed from *true* to *false*.<br>Method to check if Service Composer and Distributed Firewall are in sync is deprecated, and will be removed in a future release. Use `GET /2.0/services/policy/securitypolicy/status` instead.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyApiApi20ServicesPolicyServiceproviderFirewallGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyApiApi20ServicesPolicyServiceproviderFirewallGetOpts struct
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

# **Api20ServicesPolicyVirtualmachineIDSecurityactionsGet**
> Api20ServicesPolicyVirtualmachineIDSecurityactionsGet(ctx, iD)
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

