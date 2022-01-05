# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20EndpointsecurityActivationGet**](EndpointsecurityApi.md#Api20EndpointsecurityActivationGet) | **Get** /api/2.0/endpointsecurity/activation | guestIntrospectionSolutionActivationRead
[**Api20EndpointsecurityActivationVendorIDAltitudeMoidDelete**](EndpointsecurityApi.md#Api20EndpointsecurityActivationVendorIDAltitudeMoidDelete) | **Delete** /api/2.0/endpointsecurity/activation/{vendorID}/{altitude}/{moid} | guestIntrospectionSolutionDeactivate
[**Api20EndpointsecurityActivationVendorIDAltitudeMoidGet**](EndpointsecurityApi.md#Api20EndpointsecurityActivationVendorIDAltitudeMoidGet) | **Get** /api/2.0/endpointsecurity/activation/{vendorID}/{altitude}/{moid} | guestIntrospectionSolutionActivationStatusRead
[**Api20EndpointsecurityActivationVendorIDAltitudePost**](EndpointsecurityApi.md#Api20EndpointsecurityActivationVendorIDAltitudePost) | **Post** /api/2.0/endpointsecurity/activation/{vendorID}/{altitude} | guestIntrospectionSolutionActivateRegistered
[**Api20EndpointsecurityActivationVendorIDSolutionIDGet**](EndpointsecurityApi.md#Api20EndpointsecurityActivationVendorIDSolutionIDGet) | **Get** /api/2.0/endpointsecurity/activation/{vendorID}/{solutionID} | guestIntrospectionSolutionActivatedVMsRead
[**Api20EndpointsecurityRegistrationPost**](EndpointsecurityApi.md#Api20EndpointsecurityRegistrationPost) | **Post** /api/2.0/endpointsecurity/registration | guestIntrospectionVendorCreate
[**Api20EndpointsecurityRegistrationVendorIDAltitudeDelete**](EndpointsecurityApi.md#Api20EndpointsecurityRegistrationVendorIDAltitudeDelete) | **Delete** /api/2.0/endpointsecurity/registration/{vendorID}/{altitude} | guestIntrospectionSolutionDelete
[**Api20EndpointsecurityRegistrationVendorIDAltitudeGet**](EndpointsecurityApi.md#Api20EndpointsecurityRegistrationVendorIDAltitudeGet) | **Get** /api/2.0/endpointsecurity/registration/{vendorID}/{altitude} | guestIntrospectionSolutionInfoRead
[**Api20EndpointsecurityRegistrationVendorIDAltitudeLocationDelete**](EndpointsecurityApi.md#Api20EndpointsecurityRegistrationVendorIDAltitudeLocationDelete) | **Delete** /api/2.0/endpointsecurity/registration/{vendorID}/{altitude}/location | guestIntrospectionSolutionIPPortDelete
[**Api20EndpointsecurityRegistrationVendorIDAltitudeLocationGet**](EndpointsecurityApi.md#Api20EndpointsecurityRegistrationVendorIDAltitudeLocationGet) | **Get** /api/2.0/endpointsecurity/registration/{vendorID}/{altitude}/location | guestIntrospectionSolutionIPPortRead
[**Api20EndpointsecurityRegistrationVendorIDAltitudeLocationPost**](EndpointsecurityApi.md#Api20EndpointsecurityRegistrationVendorIDAltitudeLocationPost) | **Post** /api/2.0/endpointsecurity/registration/{vendorID}/{altitude}/location | guestIntrospectionSolutionIPPortUpdate
[**Api20EndpointsecurityRegistrationVendorIDDelete**](EndpointsecurityApi.md#Api20EndpointsecurityRegistrationVendorIDDelete) | **Delete** /api/2.0/endpointsecurity/registration/{vendorID} | guestIntrospectionVendorDelete
[**Api20EndpointsecurityRegistrationVendorIDGet**](EndpointsecurityApi.md#Api20EndpointsecurityRegistrationVendorIDGet) | **Get** /api/2.0/endpointsecurity/registration/{vendorID} | guestIntrospectionVendorInfoRead
[**Api20EndpointsecurityRegistrationVendorIDPost**](EndpointsecurityApi.md#Api20EndpointsecurityRegistrationVendorIDPost) | **Post** /api/2.0/endpointsecurity/registration/{vendorID} | guestIntrospectionSolutionCreate
[**Api20EndpointsecurityRegistrationVendorIDSolutionsGet**](EndpointsecurityApi.md#Api20EndpointsecurityRegistrationVendorIDSolutionsGet) | **Get** /api/2.0/endpointsecurity/registration/{vendorID}/solutions | guestIntrospectionSolutionsInfoRead
[**Api20EndpointsecurityRegistrationVendorsGet**](EndpointsecurityApi.md#Api20EndpointsecurityRegistrationVendorsGet) | **Get** /api/2.0/endpointsecurity/registration/vendors | guestIntrospectionVendorsInfoList

# **Api20EndpointsecurityActivationGet**
> Api20EndpointsecurityActivationGet(ctx, optional)
guestIntrospectionSolutionActivationRead

Retrieve activation information for all activated security VMs on the specified host.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EndpointsecurityApiApi20EndpointsecurityActivationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsecurityApiApi20EndpointsecurityActivationGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostId** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20EndpointsecurityActivationVendorIDAltitudeMoidDelete**
> Api20EndpointsecurityActivationVendorIDAltitudeMoidDelete(ctx, vendorID, altitude, moid)
guestIntrospectionSolutionDeactivate

Deactivate an endpoint protection solution on a host.  Parameters:  moid: Managed object reference of a VM.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
  **altitude** | **string**|  | 
  **moid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20EndpointsecurityActivationVendorIDAltitudeMoidGet**
> Api20EndpointsecurityActivationVendorIDAltitudeMoidGet(ctx, vendorID, altitude, moid)
guestIntrospectionSolutionActivationStatusRead

Retrieve the endpoint protection solution activation status, either true (activated) or false (not activated).  Parameters:  moid: Managed object reference of a VM.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
  **altitude** | **string**|  | 
  **moid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20EndpointsecurityActivationVendorIDAltitudePost**
> Api20EndpointsecurityActivationVendorIDAltitudePost(ctx, vendorID, altitude, optional)
guestIntrospectionSolutionActivateRegistered

Activate an endpoint protection solution that has been registered and located. Specify the following parameter in the request body.  | Name            | Comments | |-----------------|------------| |**svmMoid**     | Managed object ID of the virtual machine of the activated endpoint protection solution. |   Parameters:  vendorID: VMware-assigned ID for the vendor.  altitude: VMware-assigned number to uniquely identify a solution. Describes the type of solution and the order in which the solution receives events relative to other solutions on the same host.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
  **altitude** | **string**|  | 
 **optional** | ***EndpointsecurityApiApi20EndpointsecurityActivationVendorIDAltitudePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsecurityApiApi20EndpointsecurityActivationVendorIDAltitudePostOpts struct
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

# **Api20EndpointsecurityActivationVendorIDSolutionIDGet**
> Api20EndpointsecurityActivationVendorIDSolutionIDGet(ctx, vendorID, solutionID)
guestIntrospectionSolutionActivatedVMsRead

Retrieve a list of activated security VMs for an endpoint protection solution.   Parameters:  vendorID: VMware-assigned ID for the vendor.  solutionID: solution ID for the endpoint protection solution.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
  **solutionID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20EndpointsecurityRegistrationPost**
> Api20EndpointsecurityRegistrationPost(ctx, optional)
guestIntrospectionVendorCreate

Register the vendor of an endpoint protection solution. Specify the following parameters in the request.  | Name            | Comments | |-----------------|------------| |**vendorId**     | VMware-assigned ID for the vendor. | |**vendorTitle**  | Vendor-specified title. | |**vendorDescription** | Vendor-specified description. |   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EndpointsecurityApiApi20EndpointsecurityRegistrationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsecurityApiApi20EndpointsecurityRegistrationPostOpts struct
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

# **Api20EndpointsecurityRegistrationVendorIDAltitudeDelete**
> Api20EndpointsecurityRegistrationVendorIDAltitudeDelete(ctx, vendorID, altitude)
guestIntrospectionSolutionDelete

Unregister an endpoint protection solution.  Parameters:  altitude: VMware-assigned number that uniquely identifies a solution. Describes the type of solution and the order in which the solution receives events relative to other solutions on the same host.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
  **altitude** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20EndpointsecurityRegistrationVendorIDAltitudeGet**
> Api20EndpointsecurityRegistrationVendorIDAltitudeGet(ctx, vendorID, altitude)
guestIntrospectionSolutionInfoRead

Get registration information for an endpoint protection solution.  Parameters:  altitude: VMware-assigned number that uniquely identifies a solution. Describes the type of solution and the order in which the solution receives events relative to other solutions on the same host.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
  **altitude** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20EndpointsecurityRegistrationVendorIDAltitudeLocationDelete**
> Api20EndpointsecurityRegistrationVendorIDAltitudeLocationDelete(ctx, vendorID, altitude)
guestIntrospectionSolutionIPPortDelete

Unset the IP address and port for an endpoint protection solution.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
  **altitude** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20EndpointsecurityRegistrationVendorIDAltitudeLocationGet**
> Api20EndpointsecurityRegistrationVendorIDAltitudeLocationGet(ctx, vendorID, altitude)
guestIntrospectionSolutionIPPortRead

Get the IP address and port on the vNIC host for an endpoint protection solution.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
  **altitude** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20EndpointsecurityRegistrationVendorIDAltitudeLocationPost**
> Api20EndpointsecurityRegistrationVendorIDAltitudeLocationPost(ctx, vendorID, altitude, optional)
guestIntrospectionSolutionIPPortUpdate

Set the IP address and port on the vNIC host for an endpoint protection solution.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
  **altitude** | **string**|  | 
 **optional** | ***EndpointsecurityApiApi20EndpointsecurityRegistrationVendorIDAltitudeLocationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsecurityApiApi20EndpointsecurityRegistrationVendorIDAltitudeLocationPostOpts struct
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

# **Api20EndpointsecurityRegistrationVendorIDDelete**
> Api20EndpointsecurityRegistrationVendorIDDelete(ctx, vendorID)
guestIntrospectionVendorDelete

Unregister a Guest Introspection vendor.  Parameters:  vendorID: VMware-assigned ID for the vendor.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20EndpointsecurityRegistrationVendorIDGet**
> Api20EndpointsecurityRegistrationVendorIDGet(ctx, vendorID)
guestIntrospectionVendorInfoRead

Retrieve registration information for a Guest Introspection vendor.  Parameters:  vendorID: VMware-assigned ID for the vendor.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20EndpointsecurityRegistrationVendorIDPost**
> Api20EndpointsecurityRegistrationVendorIDPost(ctx, vendorID, optional)
guestIntrospectionSolutionCreate

Register an endpoint protection solution. Specify the following parameters in the request.  | Name            | Comments | |-----------------|------------| |**solutionAltitude**     | VMware-assigned altitude for the solution. *Altitude* is a number that VMware assigns to uniquely identify the solution. The altitude describes the type of solution and the order in which the solution receives events relative to other solutions on the same host. | |**solutionTitle**  | Vendor-specified title for the solution. | |**solutionDescription** | Vendor-specified description of the solution. |   Parameters:  vendorID: VMware-assigned ID for the vendor.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
 **optional** | ***EndpointsecurityApiApi20EndpointsecurityRegistrationVendorIDPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsecurityApiApi20EndpointsecurityRegistrationVendorIDPostOpts struct
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

# **Api20EndpointsecurityRegistrationVendorIDSolutionsGet**
> Api20EndpointsecurityRegistrationVendorIDSolutionsGet(ctx, vendorID)
guestIntrospectionSolutionsInfoRead

Get registration information for all endpoint protection solutions for a Guest Introspection vendor.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20EndpointsecurityRegistrationVendorsGet**
> Api20EndpointsecurityRegistrationVendorsGet(ctx, )
guestIntrospectionVendorsInfoList

Retrieve the list of all registered Guest Introspection vendors.  Parameters:  

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

