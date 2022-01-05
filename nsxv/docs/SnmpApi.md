# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesSnmpManagerGet**](SnmpApi.md#ServicesSnmpManagerGet) | **Get** /api/2.0/services/snmp/manager | snmpManagersList
[**ServicesSnmpManagerManagerIdDelete**](SnmpApi.md#ServicesSnmpManagerManagerIdDelete) | **Delete** /api/2.0/services/snmp/manager/{managerId} | snmpManagerDelete
[**ServicesSnmpManagerManagerIdGet**](SnmpApi.md#ServicesSnmpManagerManagerIdGet) | **Get** /api/2.0/services/snmp/manager/{managerId} | snmpManagerRead
[**ServicesSnmpManagerManagerIdPut**](SnmpApi.md#ServicesSnmpManagerManagerIdPut) | **Put** /api/2.0/services/snmp/manager/{managerId} | snmpManagerUpdate
[**ServicesSnmpManagerPost**](SnmpApi.md#ServicesSnmpManagerPost) | **Post** /api/2.0/services/snmp/manager | snmpManagerCreate
[**ServicesSnmpStatusGet**](SnmpApi.md#ServicesSnmpStatusGet) | **Get** /api/2.0/services/snmp/status | snmpStatusRead
[**ServicesSnmpStatusPut**](SnmpApi.md#ServicesSnmpStatusPut) | **Put** /api/2.0/services/snmp/status | snmpStatusUpdate
[**ServicesSnmpTrapGet**](SnmpApi.md#ServicesSnmpTrapGet) | **Get** /api/2.0/services/snmp/trap | snmpTrapsList
[**ServicesSnmpTrapOidGet**](SnmpApi.md#ServicesSnmpTrapOidGet) | **Get** /api/2.0/services/snmp/trap/{oid} | snmpTrapRead
[**ServicesSnmpTrapOidPut**](SnmpApi.md#ServicesSnmpTrapOidPut) | **Put** /api/2.0/services/snmp/trap/{oid} | snmpTrapUpdate

# **ServicesSnmpManagerGet**
> ServicesSnmpManagerGet(ctx, )
snmpManagersList

Retrieve information about SNMP managers.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

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

# **ServicesSnmpManagerManagerIdDelete**
> ServicesSnmpManagerManagerIdDelete(ctx, managerId)
snmpManagerDelete

Delete an SNMP manager configuration.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  managerId: ID of the SNMP manager.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **managerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSnmpManagerManagerIdGet**
> ServicesSnmpManagerManagerIdGet(ctx, managerId)
snmpManagerRead

Retrieve information about the specified SNMP manager.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  managerId: ID of the SNMP manager.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **managerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSnmpManagerManagerIdPut**
> ServicesSnmpManagerManagerIdPut(ctx, managerId, optional)
snmpManagerUpdate

Update an SNMP manager configuration.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  managerId: ID of the SNMP manager.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **managerId** | **string**|  | 
 **optional** | ***SnmpApiServicesSnmpManagerManagerIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnmpApiServicesSnmpManagerManagerIdPutOpts struct
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

# **ServicesSnmpManagerPost**
> ServicesSnmpManagerPost(ctx, optional)
snmpManagerCreate

Add an SNMP manager.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnmpApiServicesSnmpManagerPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnmpApiServicesSnmpManagerPostOpts struct
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

# **ServicesSnmpStatusGet**
> ServicesSnmpStatusGet(ctx, )
snmpStatusRead

Retrieve SNMP status settings.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

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

# **ServicesSnmpStatusPut**
> ServicesSnmpStatusPut(ctx, optional)
snmpStatusUpdate

Update SNMP status settings.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnmpApiServicesSnmpStatusPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnmpApiServicesSnmpStatusPutOpts struct
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

# **ServicesSnmpTrapGet**
> ServicesSnmpTrapGet(ctx, )
snmpTrapsList

Retrieve information about SNMP traps.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

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

# **ServicesSnmpTrapOidGet**
> ServicesSnmpTrapOidGet(ctx, oid)
snmpTrapRead

Retrieve information about the specified SNMP trap.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  oid: SNMP object identifier.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSnmpTrapOidPut**
> ServicesSnmpTrapOidPut(ctx, oid, optional)
snmpTrapUpdate

Update the specified SNMP trap.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  oid: SNMP object identifier.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oid** | **string**|  | 
 **optional** | ***SnmpApiServicesSnmpTrapOidPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnmpApiServicesSnmpTrapOidPutOpts struct
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

