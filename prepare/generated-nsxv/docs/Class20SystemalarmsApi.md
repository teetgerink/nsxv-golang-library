# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesSystemalarmsAlarmIdGet**](Class20SystemalarmsApi.md#Api20ServicesSystemalarmsAlarmIdGet) | **Get** /api/2.0/services/systemalarms/{alarmId} | servicesSystemAlarmsIdRead
[**Api20ServicesSystemalarmsAlarmIdPost**](Class20SystemalarmsApi.md#Api20ServicesSystemalarmsAlarmIdPost) | **Post** /api/2.0/services/systemalarms/{alarmId} | servicesSystemAlarmsIdAction

# **Api20ServicesSystemalarmsAlarmIdGet**
> Api20ServicesSystemalarmsAlarmIdGet(ctx, alarmId)
servicesSystemAlarmsIdRead

Retrieve information about the specified alarm.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  alarmId: The alarm ID you want to manage. Find the alarm ID using the `GET /api/2.0/services/alarms/{source-Id}` method.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alarmId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesSystemalarmsAlarmIdPost**
> Api20ServicesSystemalarmsAlarmIdPost(ctx, alarmId, optional)
servicesSystemAlarmsIdAction

Resolve the specified alarm.  Alarms will resolve automatically when the cause of the alarm is resolved.  For example, if an NSX Edge appliance is powered off, this will trigger an alarm. If you power the NSX Edge appliance back on, the alarm will resolve. If however, you delete the NSX Edge appliance, the alarm will persist, because the alarm cause was never resolved. In this case, you may want to manually resolve the alarm. Resolving the alarm  will clear it from the NSX dashboard.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  alarmId: The alarm ID you want to manage. Find the alarm ID using the `GET /api/2.0/services/alarms/{source-Id}` method.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alarmId** | **string**|  | 
 **optional** | ***Class20SystemalarmsApiApi20ServicesSystemalarmsAlarmIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class20SystemalarmsApiApi20ServicesSystemalarmsAlarmIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

