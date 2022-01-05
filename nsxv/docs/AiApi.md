# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiAppAppIDGet**](AiApi.md#AiAppAppIDGet) | **Get** /api/3.0/ai/app/{appID} | specificAppDetailsRead
[**AiAppGet**](AiApi.md#AiAppGet) | **Get** /api/3.0/ai/app | appDetailsRead
[**AiDesktoppoolDesktoppoolIDGet**](AiApi.md#AiDesktoppoolDesktoppoolIDGet) | **Get** /api/3.0/ai/desktoppool/{desktoppoolID} | specificDesktopPoolRead
[**AiDesktoppoolGet**](AiApi.md#AiDesktoppoolGet) | **Get** /api/3.0/ai/desktoppool | desktopPoolRead
[**AiDirectorygroupDirectorygroupIDGet**](AiApi.md#AiDirectorygroupDirectorygroupIDGet) | **Get** /api/3.0/ai/directorygroup/{directorygroupID} | specificDirGroupRead
[**AiDirectorygroupGet**](AiApi.md#AiDirectorygroupGet) | **Get** /api/3.0/ai/directorygroup | dirGroupRead
[**AiDirectorygroupUserUserIDGet**](AiApi.md#AiDirectorygroupUserUserIDGet) | **Get** /api/3.0/ai/directorygroup/user/{userID} | userDirGroupRead
[**AiHostGet**](AiApi.md#AiHostGet) | **Get** /api/3.0/ai/host | hostDetailsRead
[**AiHostHostIDGet**](AiApi.md#AiHostHostIDGet) | **Get** /api/3.0/ai/host/{hostID} | specificHostDetailsRead
[**AiRecordsGet**](AiApi.md#AiRecordsGet) | **Get** /api/3.0/ai/records | userActivityRead
[**AiSecuritygroupGet**](AiApi.md#AiSecuritygroupGet) | **Get** /api/3.0/ai/securitygroup | secgroupDetailsRead
[**AiSecuritygroupSecgroupIDGet**](AiApi.md#AiSecuritygroupSecgroupIDGet) | **Get** /api/3.0/ai/securitygroup/{secgroupID} | specificSecgroupDetailsRead
[**AiUserUserIDGet**](AiApi.md#AiUserUserIDGet) | **Get** /api/3.0/ai/user/{userID} | specificUserDetailsRead
[**AiUserdetailsGet**](AiApi.md#AiUserdetailsGet) | **Get** /api/3.0/ai/userdetails | userDetailsRead
[**AiVmGet**](AiApi.md#AiVmGet) | **Get** /api/3.0/ai/vm | vmDetailsRead
[**AiVmVmIDGet**](AiApi.md#AiVmVmIDGet) | **Get** /api/3.0/ai/vm/{vmID} | specificVMDetailsRead

# **AiAppAppIDGet**
> AiAppAppIDGet(ctx, appID)
specificAppDetailsRead

Retrieve details for specific app.  Parameters:  appID: Specified app ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AiAppGet**
> AiAppGet(ctx, )
appDetailsRead

Retrieve app details.  Parameters:  

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

# **AiDesktoppoolDesktoppoolIDGet**
> AiDesktoppoolDesktoppoolIDGet(ctx, desktoppoolID)
specificDesktopPoolRead

Retrieve specific desktop pool details.  Parameters:  desktoppoolID: Specified desktop pool.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **desktoppoolID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AiDesktoppoolGet**
> AiDesktoppoolGet(ctx, )
desktopPoolRead

Retrieve list of all discovered desktop pools by agent introspection.   Parameters:  

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

# **AiDirectorygroupDirectorygroupIDGet**
> AiDirectorygroupDirectorygroupIDGet(ctx, directorygroupID)
specificDirGroupRead

Retrieve details about a specific directory group.  Parameters:  directorygroupID: Specified directory group.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **directorygroupID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AiDirectorygroupGet**
> AiDirectorygroupGet(ctx, )
dirGroupRead

Retrieve list of all discovered (and configured) LDAP directory groups.   Parameters:  

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

# **AiDirectorygroupUserUserIDGet**
> AiDirectorygroupUserUserIDGet(ctx, userID)
userDirGroupRead

Retrieve Active Directory groups that user belongs to.  Parameters:  userID: User ID.  

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

# **AiHostGet**
> AiHostGet(ctx, )
hostDetailsRead

Retrieve list of all discovered hosts (both by agent introspection and LDAP Sync) and their detail.   Parameters:  

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

# **AiHostHostIDGet**
> AiHostHostIDGet(ctx, hostID)
specificHostDetailsRead

Get host details.  Parameters:  hostID: Specified host ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AiRecordsGet**
> AiRecordsGet(ctx, optional)
userActivityRead

### View Outbound Activity  You can view what applications are being run by a security group or desktop pool and then drill down into the report to find out which client applications are making outbound connections by a particular group of users. You can also discover all user groups and users who are accessing a particular application, which can help you determine if you need to adjust identity firewall in your environment.  * query=*resource* * param=&lt;param-name&gt;:&lt;param-type&gt;:&lt;comma-separated-values&gt;:&lt;operator&gt;, where:   * &lt;param-name&gt; is one of:     * *src* (required)     * *dest* (required)     * *app*   * &lt;param-type&gt; is one of:     * for src: *SECURITY_GROUP*, *DIRECTORY_GROUP*, *DESKTOP_POOL*     * for dest: *VIRTUAL_MACHINE*     * for app: *SRC_APP*   * &lt;comma-separated-values&gt; is a comma-separated numbers (optional). If none specified then no filter is applied.   * &lt;operator&gt; is one of *INCLUDE*, *EXCLUDE* (default is *INCLUDE*).  **Example:** View user activities to VM ID 1 originating from application ID 1   `GET /api/3.0/ai/records?query=resource&interval=60m&param=src:DIRECTORY_GROUP`   `&param=dest:VIRTUAL_MACHINE:1&param=app:SRC_APP:1`  ### View Inbound Activity  You can view all inbound activity to a server by desktop pool, security group, or AD group.  * query=*sam* * param=&lt;param-name&gt;:&lt;param-type&gt;:&lt;comma-separated-values&gt;:&lt;operator&gt;, where:   * &lt;param-name&gt; is one of:     * *src* (required)     * *dest* (required)     * *app*   * &lt;param-type&gt; is one of:     * for src: *SECURITY_GROUP*, *DIRECTORY_GROUP*, *DESKTOP_POOL*     * for dest: *VIRTUAL_MACHINE*     * for app: *DEST_APP*   * &lt;comma-separated-values&gt; is a comma-separated numbers (optional). If none specified then no filter is applied.   * &lt;operator&gt; is one of *INCLUDE*, *EXCLUDE*, *NOT* (default is *INCLUDE*).  **Example:** View user activities to VM ID 1 originating from application ID 1   `GET /api/3.0/ai/records?query=containers&interval=60m&param=dest:SECURITY_GROUP:1:EXCLUDE`   `&param=src:SECURITY_GROUP:1`  ### View Interaction between Inventory Containers You can view the traffic passing between defined containers such as AD groups, security groups and/or desktop pools. This can help you identify and configure access to shared services and to resolve misconfigured relationships between Inventory container definitions, desktop pools and AD groups.  * query=*containers* * param=&lt;param-name&gt;:&lt;param-type&gt;:&lt;comma-separated-values&gt;:&lt;operator&gt;, where:   * &lt;param-name&gt; is one of:     * *src* (required)     * *dest* (required)   * &lt;param-type&gt; is one of:     * for src: *SECURITY_GROUP*, *DIRECTORY_GROUP*, *DESKTOP_POOL*     * for dest: *SECURITY_GROUP*, * *DESKTOP_POOL*    * &lt;comma-separated-values&gt; is a comma-separated numbers (optional). If none specified then no filter is applied.   * &lt;operator&gt; is one of *INCLUDE*, *EXCLUDE*, or *NOT* (default * is *INCLUDE*).  **Example:** View interaction between inventory containers   `GET /api/3.0/ai/records?query=containers&interval=60m&param=dest:SECURITY_GROUP:1:EXCLUDE`   `&param=src:SECURITY_GROUP:1`  ### View Outbound AD Group Activity  You can view the traffic between members of defined Active Directory groups and can use this data to fine tune your firewall rules.  * query=*adg* * param=&lt;param-name&gt;:&lt;param-type&gt;:&lt;comma-separated-values&gt;:&lt;operator&gt;, where:   * &lt;param-name&gt; is one of:     * *src* (required)     * *adg*   * &lt;param-type&gt; is one of:     * for src: *SECURITY_GROUP*, *DESKTOP_POOL*     * for adg: *USER*   * &lt;comma-separated-values&gt; is a comma-separated numbers (optional). If none specified then no filter is applied.   * &lt;operator&gt; is one of *INCLUDE*, *EXCLUDE* (default * is *INCLUDE*).  **Example:** View outbound AD group activity     `GET https://NSX-Manager-IP-Address/api/3.0/ai/records?query=adg&interval=24h&param=adg:USER:1:INCLUDE`   `&param=src:SECURITY_GROUP:1:EXCLUDE`   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AiApiAiRecordsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AiApiAiRecordsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**|  | 
 **interval** | **optional.String**|  | 
 **stime** | **optional.String**|  | 
 **etime** | **optional.String**|  | 
 **param** | **optional.String**|  | 
 **pagesize** | **optional.String**|  | 
 **startindex** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AiSecuritygroupGet**
> AiSecuritygroupGet(ctx, )
secgroupDetailsRead

Retrieve list of all observed security groups.  Observed entities are the ones that are reported by the agents. For example, if a host activity is reported by an agent and if that host belongs to a security group then that security group would reported as observed in SAM database.   Parameters:  

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

# **AiSecuritygroupSecgroupIDGet**
> AiSecuritygroupSecgroupIDGet(ctx, secgroupID)
specificSecgroupDetailsRead

Retrieve details about specific security group.  Parameters:  secgroupID: Specified security group.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **secgroupID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AiUserUserIDGet**
> AiUserUserIDGet(ctx, userID)
specificUserDetailsRead

Retrieve details for a specific user.  Parameters:  userID: User ID  

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

# **AiUserdetailsGet**
> AiUserdetailsGet(ctx, optional)
userDetailsRead

### View Outbound Activity You can view what applications are being run by a security group or desktop pool and then drill down into the report to find out which client applications are making outbound connections by a particular group of users. You can also discover all user groups and users who are accessing a particular application, which can help you determine if you need to adjust identity firewall in your environment.  * query=*resource* * param=&lt;param-name&gt;&lt;param-type&gt;&lt;comma-separated-values&gt;&lt;operator&gt;, where:   * &lt;param-name&gt; is one of:     * *src* (required)     * *dest* (required)     * *app*   * &lt;param-type&gt; is one of:     * for src: *SECURITY_GROUP*, *DIRECTORY_GROUP*, *DESKTOP_POOL*     * for dest: *IP* - a valid IP address in dot notation, xx.xx.xx.xx     * for app: *SRC_APP*   * &lt;comma-separated-values&gt; is a comma-separated numbers (optional). If none specified then no filter is applied.   * &lt;operator&gt; is one of *INCLUDE*, *EXCLUDE* (default is *INCLUDE*).  **Example:** View user activities to VM ID 1 originating from application ID 1   `GET /api/3.0/ai/userdetails?query=resource&stime=2012-10-15T00:00:00&etime=2012-10-20T00:00:00`   `&param=src:DIRECTORY_GROUP:2&param=app:SRC_APP:16&param=dest:IP:172.16.4.52`  ### View Inbound Activity  You can view all inbound activity to a server by desktop pool, security group, or AD group.  * query=*sam* * param=&lt;param-name&gt;&lt;param-type&gt;&lt;comma-separated-values&gt;&lt;operator&gt;, where:   * &lt;param-name&gt; is one of:     * *src* (required)     * *dest* (required)     * *app* (required)   * &lt;param-type&gt; is one of:     * for src: *SECURITY_GROUP*, *DIRECTORY_GROUP*, *DESKTOP_POOL*     * for dest: *VIRTUAL_MACHINE*     * for app: *DEST_APP*   * &lt;comma-separated-values&gt; is a comma-separated numbers (optional). If none specified then no filter is applied.   * &lt;operator&gt; is one of *INCLUDE*, *EXCLUDE*, *NOT* (default is *INCLUDE*).  **Example:** View user activities to VM ID 1 originating from application ID 1   `GET /api/3.0/userdetails?query=sam&interval=60m&param=app:DEST_APP:1:EXCLUDE`   `&param=dest:IP:1:EXCLUDE&param=src:SECURITY_GROUP:1:EXCLUDE`  ### View Interaction between Inventory Containers You can view the traffic passing between defined containers such as AD groups, security groups and/or desktop pools. This can help you identify and configure access to shared services and to resolve misconfigured relationships between Inventory container definitions, desktop pools and AD groups.  * query=*containers* * param=&lt;param-name&gt;&lt;param-type&gt;&lt;comma-separated-values&gt;&lt;operator&gt;, where:   * &lt;param-name&gt; is one of:     * *src* (required)     * *dest* (required)   * &lt;param-type&gt; is one of:     * for src: *SECURITY_GROUP*, *DIRECTORY_GROUP*, *DESKTOP_POOL*     * for dest: *SECURITY_GROUP*, * *DESKTOP_POOL*    * &lt;comma-separated-values&gt; is a comma-separated numbers (optional). If none specified then no filter is applied.   * &lt;operator&gt; is one of *INCLUDE*, *EXCLUDE*, or *NOT* (default * is *INCLUDE*).  **Example:** View interaction between inventory containers   `GET /api/3.0/ai/userdetails?query=containers&interval=60m&param=dest:SECURITY_GROUP:1:EXCLUDE`   `&param=src:SECURITY_GROUP:1`  ### View Outbound AD Group Activity  You can view the traffic between members of defined Active Directory groups and can use this data to fine tune your firewall rules.  * query=*adg* * param=&lt;param-name&gt;&lt;param-type&gt;&lt;comma-separated-values&gt;&lt;operator&gt;, where:   * &lt;param-name&gt; is one of:     * *src* (required)     * *adg*   * &lt;param-type&gt; is one of:     * for src: *SECURITY_GROUP*, *DESKTOP_POOL*     * for adg: *USER*   * &lt;comma-separated-values&gt; is a comma-separated numbers (optional). If none specified then no filter is applied.   * &lt;operator&gt; is one of *INCLUDE*, *EXCLUDE* (default is *INCLUDE*).  **Example:** View outbound AD group activity     `GET /api/3.0/ai/userdetails?query=adg&interval=24h&param=adg:USER:1:INCLUDE`   `&param=src:SECURITY_GROUP:1:EXCLUDE`  ### View Virtual Machine Activity Report  * query=*vma* * param=&lt;param-name&gt;&lt;param-type&gt;&lt;comma-separated-values&gt;&lt;operator&gt;, where:   * &lt;param-name&gt; is one of:     * *src*     * *dst*     * *app*     * If no parameters are passed, then this would show all SAM     activities   * &lt;param-type&gt; is one of:     * for src: *SECURITY_GROUP*, *DESKTOP_POOL*     * for dst: *VIRTUAL_MACHINE*, *VM_UUID*     * for app - *SRC_APP* or *DEST_APP*   * &lt;comma-separated-values&gt; is a comma-separated numbers (optional). If none specified then no filter is applied.   * &lt;operator&gt; is one of *INCLUDE*, *EXCLUDE* (default is *INCLUDE*).  **Example:** View outbound AD group activity     `GET /api/3.0/ai/userdetails?query=vma&interval=60m&param=dest:VIRTUAL_MACHINE:1 &param=app:DEST_APP:16`   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AiApiAiUserdetailsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AiApiAiUserdetailsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**|  | 
 **interval** | **optional.String**|  | 
 **stime** | **optional.String**|  | 
 **etime** | **optional.String**|  | 
 **param** | **optional.String**|  | 
 **pagesize** | **optional.String**|  | 
 **startindex** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AiVmGet**
> AiVmGet(ctx, )
vmDetailsRead

Retrieve list of all discovered VMs.  Parameters:  

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

# **AiVmVmIDGet**
> AiVmVmIDGet(ctx, vmID)
specificVMDetailsRead

Retrieve details about a specific virtual machine.  Parameters:  vmID: VM ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

