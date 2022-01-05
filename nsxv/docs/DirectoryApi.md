# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryDeleteDomainIDDelete**](DirectoryApi.md#DirectoryDeleteDomainIDDelete) | **Delete** /api/1.0/directory/deleteDomain/{ID} | domainDeleteDelete
[**DirectoryDeleteEventLogServerServerIDDelete**](DirectoryApi.md#DirectoryDeleteEventLogServerServerIDDelete) | **Delete** /api/1.0/directory/deleteEventLogServer/{serverID} | eLogServerDeleteDelete
[**DirectoryDeleteLdapServerServerIDDelete**](DirectoryApi.md#DirectoryDeleteLdapServerServerIDDelete) | **Delete** /api/1.0/directory/deleteLdapServer/{serverID} | ldapServerDeleteDelete
[**DirectoryDeltaSyncDomainIDPut**](DirectoryApi.md#DirectoryDeltaSyncDomainIDPut) | **Put** /api/1.0/directory/deltaSync/{domainID} | ldapServerDeltaSyncExecute
[**DirectoryFullSyncDomainIDPut**](DirectoryApi.md#DirectoryFullSyncDomainIDPut) | **Put** /api/1.0/directory/fullSync/{domainID} | ldapServerSyncExecute
[**DirectoryListDomainsGet**](DirectoryApi.md#DirectoryListDomainsGet) | **Get** /api/1.0/directory/listDomains | domainListRead
[**DirectoryListEventLogServersForDomainDomainIDGet**](DirectoryApi.md#DirectoryListEventLogServersForDomainDomainIDGet) | **Get** /api/1.0/directory/listEventLogServersForDomain/{domainID} | eLogServerDomainReadRead
[**DirectoryListLdapServersForDomainDomainIDGet**](DirectoryApi.md#DirectoryListLdapServersForDomainDomainIDGet) | **Get** /api/1.0/directory/listLdapServersForDomain/{domainID} | ldapServerDomainRead
[**DirectoryUpdateDomainPost**](DirectoryApi.md#DirectoryUpdateDomainPost) | **Post** /api/1.0/directory/updateDomain | domainCreate
[**DirectoryUpdateEventLogServerPost**](DirectoryApi.md#DirectoryUpdateEventLogServerPost) | **Post** /api/1.0/directory/updateEventLogServer | eLogServerCreateCreate
[**DirectoryUpdateLdapServerPost**](DirectoryApi.md#DirectoryUpdateLdapServerPost) | **Post** /api/1.0/directory/updateLdapServer | ldapServerCreateCreate

# **DirectoryDeleteDomainIDDelete**
> DirectoryDeleteDomainIDDelete(ctx, iD)
domainDeleteDelete

Delete domain.  Parameters:  ID: Domain ID.  

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

# **DirectoryDeleteEventLogServerServerIDDelete**
> DirectoryDeleteEventLogServerServerIDDelete(ctx, serverID)
eLogServerDeleteDelete

Delete EventLog server.  Parameters:  serverID: Specified EventLog server ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryDeleteLdapServerServerIDDelete**
> DirectoryDeleteLdapServerServerIDDelete(ctx, serverID)
ldapServerDeleteDelete

Delete LDAP server.  Parameters:  serverID: Specified LDAP server.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryDeltaSyncDomainIDPut**
> DirectoryDeltaSyncDomainIDPut(ctx, domainID)
ldapServerDeltaSyncExecute

Start LDAP delta sync.  Parameters:  domainID: Specified domain.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryFullSyncDomainIDPut**
> DirectoryFullSyncDomainIDPut(ctx, domainID)
ldapServerSyncExecute

Start LDAP full sync.  Parameters:  domainID: Specified domain.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryListDomainsGet**
> DirectoryListDomainsGet(ctx, )
domainListRead

Retrieve all agent discovered (or configured) LDAP domains.  Parameters:  

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

# **DirectoryListEventLogServersForDomainDomainIDGet**
> DirectoryListEventLogServersForDomainDomainIDGet(ctx, domainID)
eLogServerDomainReadRead

Query EventLog servers for a domain.  Parameters:  domainID: Specified domain.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryListLdapServersForDomainDomainIDGet**
> DirectoryListLdapServersForDomainDomainIDGet(ctx, domainID)
ldapServerDomainRead

Query LDAP servers for a domain.  Parameters:  domainID: Specified domain.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DirectoryUpdateDomainPost**
> DirectoryUpdateDomainPost(ctx, optional)
domainCreate

Register or update a domain with NSX Manager  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoryApiDirectoryUpdateDomainPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DirectoryApiDirectoryUpdateDomainPostOpts struct
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

# **DirectoryUpdateEventLogServerPost**
> DirectoryUpdateEventLogServerPost(ctx, optional)
eLogServerCreateCreate

Create EventLog server.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoryApiDirectoryUpdateEventLogServerPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DirectoryApiDirectoryUpdateEventLogServerPostOpts struct
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

# **DirectoryUpdateLdapServerPost**
> DirectoryUpdateLdapServerPost(ctx, optional)
ldapServerCreateCreate

Create LDAP server.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoryApiDirectoryUpdateLdapServerPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DirectoryApiDirectoryUpdateLdapServerPostOpts struct
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

