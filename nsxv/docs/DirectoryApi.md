# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api10DirectoryDeleteDomainIDDelete**](DirectoryApi.md#Api10DirectoryDeleteDomainIDDelete) | **Delete** /api/1.0/directory/deleteDomain/{ID} | domainDeleteDelete
[**Api10DirectoryDeleteEventLogServerServerIDDelete**](DirectoryApi.md#Api10DirectoryDeleteEventLogServerServerIDDelete) | **Delete** /api/1.0/directory/deleteEventLogServer/{serverID} | eLogServerDeleteDelete
[**Api10DirectoryDeleteLdapServerServerIDDelete**](DirectoryApi.md#Api10DirectoryDeleteLdapServerServerIDDelete) | **Delete** /api/1.0/directory/deleteLdapServer/{serverID} | ldapServerDeleteDelete
[**Api10DirectoryDeltaSyncDomainIDPut**](DirectoryApi.md#Api10DirectoryDeltaSyncDomainIDPut) | **Put** /api/1.0/directory/deltaSync/{domainID} | ldapServerDeltaSyncExecute
[**Api10DirectoryFullSyncDomainIDPut**](DirectoryApi.md#Api10DirectoryFullSyncDomainIDPut) | **Put** /api/1.0/directory/fullSync/{domainID} | ldapServerSyncExecute
[**Api10DirectoryListDomainsGet**](DirectoryApi.md#Api10DirectoryListDomainsGet) | **Get** /api/1.0/directory/listDomains | domainListRead
[**Api10DirectoryListEventLogServersForDomainDomainIDGet**](DirectoryApi.md#Api10DirectoryListEventLogServersForDomainDomainIDGet) | **Get** /api/1.0/directory/listEventLogServersForDomain/{domainID} | eLogServerDomainReadRead
[**Api10DirectoryListLdapServersForDomainDomainIDGet**](DirectoryApi.md#Api10DirectoryListLdapServersForDomainDomainIDGet) | **Get** /api/1.0/directory/listLdapServersForDomain/{domainID} | ldapServerDomainRead
[**Api10DirectoryUpdateDomainPost**](DirectoryApi.md#Api10DirectoryUpdateDomainPost) | **Post** /api/1.0/directory/updateDomain | domainCreate
[**Api10DirectoryUpdateEventLogServerPost**](DirectoryApi.md#Api10DirectoryUpdateEventLogServerPost) | **Post** /api/1.0/directory/updateEventLogServer | eLogServerCreateCreate
[**Api10DirectoryUpdateLdapServerPost**](DirectoryApi.md#Api10DirectoryUpdateLdapServerPost) | **Post** /api/1.0/directory/updateLdapServer | ldapServerCreateCreate

# **Api10DirectoryDeleteDomainIDDelete**
> Api10DirectoryDeleteDomainIDDelete(ctx, iD)
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

# **Api10DirectoryDeleteEventLogServerServerIDDelete**
> Api10DirectoryDeleteEventLogServerServerIDDelete(ctx, serverID)
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

# **Api10DirectoryDeleteLdapServerServerIDDelete**
> Api10DirectoryDeleteLdapServerServerIDDelete(ctx, serverID)
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

# **Api10DirectoryDeltaSyncDomainIDPut**
> Api10DirectoryDeltaSyncDomainIDPut(ctx, domainID)
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

# **Api10DirectoryFullSyncDomainIDPut**
> Api10DirectoryFullSyncDomainIDPut(ctx, domainID)
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

# **Api10DirectoryListDomainsGet**
> Api10DirectoryListDomainsGet(ctx, )
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

# **Api10DirectoryListEventLogServersForDomainDomainIDGet**
> Api10DirectoryListEventLogServersForDomainDomainIDGet(ctx, domainID)
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

# **Api10DirectoryListLdapServersForDomainDomainIDGet**
> Api10DirectoryListLdapServersForDomainDomainIDGet(ctx, domainID)
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

# **Api10DirectoryUpdateDomainPost**
> Api10DirectoryUpdateDomainPost(ctx, optional)
domainCreate

Register or update a domain with NSX Manager  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoryApiApi10DirectoryUpdateDomainPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DirectoryApiApi10DirectoryUpdateDomainPostOpts struct
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

# **Api10DirectoryUpdateEventLogServerPost**
> Api10DirectoryUpdateEventLogServerPost(ctx, optional)
eLogServerCreateCreate

Create EventLog server.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoryApiApi10DirectoryUpdateEventLogServerPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DirectoryApiApi10DirectoryUpdateEventLogServerPostOpts struct
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

# **Api10DirectoryUpdateLdapServerPost**
> Api10DirectoryUpdateLdapServerPost(ctx, optional)
ldapServerCreateCreate

Create LDAP server.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DirectoryApiApi10DirectoryUpdateLdapServerPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DirectoryApiApi10DirectoryUpdateLdapServerPostOpts struct
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

