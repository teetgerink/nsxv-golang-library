# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesAlarmsSourceIdGet**](AlarmsApi.md#Api20ServicesAlarmsSourceIdGet) | **Get** /api/2.0/services/alarms/{sourceId} | servicesAlarmsSourceList
[**Api20ServicesAlarmsSourceIdPost**](AlarmsApi.md#Api20ServicesAlarmsSourceIdPost) | **Post** /api/2.0/services/alarms/{sourceId} | servicesAlarmsSourceUpdate

# **Api20ServicesAlarmsSourceIdGet**
> Api20ServicesAlarmsSourceIdGet(ctx, sourceId)
servicesAlarmsSourceList

Retrive all alarms from the specified source.   Parameters:  sourceId: ID of the object for which you want to manage alarms. *sourceId* can be the ID of a cluster, host, resource pool, security group, or edge.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesAlarmsSourceIdPost**
> Api20ServicesAlarmsSourceIdPost(ctx, sourceId, optional)
servicesAlarmsSourceUpdate

Resolve all alarms for the specified source.  Alarms will resolve automatically when the cause of the alarm is resolved.  For example, if an NSX Edge appliance is powered off, this will trigger an alarm. If you power the NSX Edge appliance back on, the alarm will resolve. If however, you delete the NSX Edge appliance, the alarm will persist, because the alarm cause was never resolved. In this case, you may want to manually resolve the alarm. Resolving the alarms will clear them from the NSX dashboard.  Use `GET /api/2.0/services/alarms/{sourceId}` to retrieve the list of alarms for the source. Use this response as the request body for the `POST` call.   Parameters:  sourceId: ID of the object for which you want to manage alarms. *sourceId* can be the ID of a cluster, host, resource pool, security group, or edge.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceId** | **string**|  | 
 **optional** | ***AlarmsApiApi20ServicesAlarmsSourceIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlarmsApiApi20ServicesAlarmsSourceIdPostOpts struct
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

