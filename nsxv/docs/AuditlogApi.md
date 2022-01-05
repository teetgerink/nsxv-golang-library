# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditlogGet**](AuditlogApi.md#AuditlogGet) | **Get** /api/2.0/auditlog | auditLogsRead

# **AuditlogGet**
> AuditlogGet(ctx, optional)
auditLogsRead

Get NSX Manager audit logs  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuditlogApiAuditlogGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditlogApiAuditlogGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **optional.String**|  | 
 **pageSize** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

