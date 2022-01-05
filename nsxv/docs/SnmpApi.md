# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesSnmpManagerGet**](SnmpApi.md#Api20ServicesSnmpManagerGet) | **Get** /api/2.0/services/snmp/manager | snmpManagersList
[**Api20ServicesSnmpManagerManagerIdDelete**](SnmpApi.md#Api20ServicesSnmpManagerManagerIdDelete) | **Delete** /api/2.0/services/snmp/manager/{managerId} | snmpManagerDelete
[**Api20ServicesSnmpManagerManagerIdGet**](SnmpApi.md#Api20ServicesSnmpManagerManagerIdGet) | **Get** /api/2.0/services/snmp/manager/{managerId} | snmpManagerRead
[**Api20ServicesSnmpManagerManagerIdPut**](SnmpApi.md#Api20ServicesSnmpManagerManagerIdPut) | **Put** /api/2.0/services/snmp/manager/{managerId} | snmpManagerUpdate
[**Api20ServicesSnmpManagerPost**](SnmpApi.md#Api20ServicesSnmpManagerPost) | **Post** /api/2.0/services/snmp/manager | snmpManagerCreate
[**Api20ServicesSnmpStatusGet**](SnmpApi.md#Api20ServicesSnmpStatusGet) | **Get** /api/2.0/services/snmp/status | snmpStatusRead
[**Api20ServicesSnmpStatusPut**](SnmpApi.md#Api20ServicesSnmpStatusPut) | **Put** /api/2.0/services/snmp/status | snmpStatusUpdate
[**Api20ServicesSnmpTrapGet**](SnmpApi.md#Api20ServicesSnmpTrapGet) | **Get** /api/2.0/services/snmp/trap | snmpTrapsList
[**Api20ServicesSnmpTrapOidGet**](SnmpApi.md#Api20ServicesSnmpTrapOidGet) | **Get** /api/2.0/services/snmp/trap/{oid} | snmpTrapRead
[**Api20ServicesSnmpTrapOidPut**](SnmpApi.md#Api20ServicesSnmpTrapOidPut) | **Put** /api/2.0/services/snmp/trap/{oid} | snmpTrapUpdate

# **Api20ServicesSnmpManagerGet**
> Api20ServicesSnmpManagerGet(ctx, )
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

# **Api20ServicesSnmpManagerManagerIdDelete**
> Api20ServicesSnmpManagerManagerIdDelete(ctx, managerId)
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

# **Api20ServicesSnmpManagerManagerIdGet**
> Api20ServicesSnmpManagerManagerIdGet(ctx, managerId)
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

# **Api20ServicesSnmpManagerManagerIdPut**
> Api20ServicesSnmpManagerManagerIdPut(ctx, managerId, optional)
snmpManagerUpdate

Update an SNMP manager configuration.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  managerId: ID of the SNMP manager.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **managerId** | **string**|  | 
 **optional** | ***SnmpApiApi20ServicesSnmpManagerManagerIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnmpApiApi20ServicesSnmpManagerManagerIdPutOpts struct
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

# **Api20ServicesSnmpManagerPost**
> Api20ServicesSnmpManagerPost(ctx, optional)
snmpManagerCreate

Add an SNMP manager.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnmpApiApi20ServicesSnmpManagerPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnmpApiApi20ServicesSnmpManagerPostOpts struct
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

# **Api20ServicesSnmpStatusGet**
> Api20ServicesSnmpStatusGet(ctx, )
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

# **Api20ServicesSnmpStatusPut**
> Api20ServicesSnmpStatusPut(ctx, optional)
snmpStatusUpdate

Update SNMP status settings.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnmpApiApi20ServicesSnmpStatusPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnmpApiApi20ServicesSnmpStatusPutOpts struct
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

# **Api20ServicesSnmpTrapGet**
> Api20ServicesSnmpTrapGet(ctx, )
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

# **Api20ServicesSnmpTrapOidGet**
> Api20ServicesSnmpTrapOidGet(ctx, oid)
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

# **Api20ServicesSnmpTrapOidPut**
> Api20ServicesSnmpTrapOidPut(ctx, oid, optional)
snmpTrapUpdate

Update the specified SNMP trap.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method introduced.   Parameters:  oid: SNMP object identifier.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oid** | **string**|  | 
 **optional** | ***SnmpApiApi20ServicesSnmpTrapOidPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnmpApiApi20ServicesSnmpTrapOidPutOpts struct
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

