# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndpointsecurityActivationGet**](EndpointsecurityApi.md#EndpointsecurityActivationGet) | **Get** /api/2.0/endpointsecurity/activation | guestIntrospectionSolutionActivationRead
[**EndpointsecurityActivationVendorIDAltitudeMoidDelete**](EndpointsecurityApi.md#EndpointsecurityActivationVendorIDAltitudeMoidDelete) | **Delete** /api/2.0/endpointsecurity/activation/{vendorID}/{altitude}/{moid} | guestIntrospectionSolutionDeactivate
[**EndpointsecurityActivationVendorIDAltitudeMoidGet**](EndpointsecurityApi.md#EndpointsecurityActivationVendorIDAltitudeMoidGet) | **Get** /api/2.0/endpointsecurity/activation/{vendorID}/{altitude}/{moid} | guestIntrospectionSolutionActivationStatusRead
[**EndpointsecurityActivationVendorIDAltitudePost**](EndpointsecurityApi.md#EndpointsecurityActivationVendorIDAltitudePost) | **Post** /api/2.0/endpointsecurity/activation/{vendorID}/{altitude} | guestIntrospectionSolutionActivateRegistered
[**EndpointsecurityActivationVendorIDSolutionIDGet**](EndpointsecurityApi.md#EndpointsecurityActivationVendorIDSolutionIDGet) | **Get** /api/2.0/endpointsecurity/activation/{vendorID}/{solutionID} | guestIntrospectionSolutionActivatedVMsRead
[**EndpointsecurityRegistrationPost**](EndpointsecurityApi.md#EndpointsecurityRegistrationPost) | **Post** /api/2.0/endpointsecurity/registration | guestIntrospectionVendorCreate
[**EndpointsecurityRegistrationVendorIDAltitudeDelete**](EndpointsecurityApi.md#EndpointsecurityRegistrationVendorIDAltitudeDelete) | **Delete** /api/2.0/endpointsecurity/registration/{vendorID}/{altitude} | guestIntrospectionSolutionDelete
[**EndpointsecurityRegistrationVendorIDAltitudeGet**](EndpointsecurityApi.md#EndpointsecurityRegistrationVendorIDAltitudeGet) | **Get** /api/2.0/endpointsecurity/registration/{vendorID}/{altitude} | guestIntrospectionSolutionInfoRead
[**EndpointsecurityRegistrationVendorIDAltitudeLocationDelete**](EndpointsecurityApi.md#EndpointsecurityRegistrationVendorIDAltitudeLocationDelete) | **Delete** /api/2.0/endpointsecurity/registration/{vendorID}/{altitude}/location | guestIntrospectionSolutionIPPortDelete
[**EndpointsecurityRegistrationVendorIDAltitudeLocationGet**](EndpointsecurityApi.md#EndpointsecurityRegistrationVendorIDAltitudeLocationGet) | **Get** /api/2.0/endpointsecurity/registration/{vendorID}/{altitude}/location | guestIntrospectionSolutionIPPortRead
[**EndpointsecurityRegistrationVendorIDAltitudeLocationPost**](EndpointsecurityApi.md#EndpointsecurityRegistrationVendorIDAltitudeLocationPost) | **Post** /api/2.0/endpointsecurity/registration/{vendorID}/{altitude}/location | guestIntrospectionSolutionIPPortUpdate
[**EndpointsecurityRegistrationVendorIDDelete**](EndpointsecurityApi.md#EndpointsecurityRegistrationVendorIDDelete) | **Delete** /api/2.0/endpointsecurity/registration/{vendorID} | guestIntrospectionVendorDelete
[**EndpointsecurityRegistrationVendorIDGet**](EndpointsecurityApi.md#EndpointsecurityRegistrationVendorIDGet) | **Get** /api/2.0/endpointsecurity/registration/{vendorID} | guestIntrospectionVendorInfoRead
[**EndpointsecurityRegistrationVendorIDPost**](EndpointsecurityApi.md#EndpointsecurityRegistrationVendorIDPost) | **Post** /api/2.0/endpointsecurity/registration/{vendorID} | guestIntrospectionSolutionCreate
[**EndpointsecurityRegistrationVendorIDSolutionsGet**](EndpointsecurityApi.md#EndpointsecurityRegistrationVendorIDSolutionsGet) | **Get** /api/2.0/endpointsecurity/registration/{vendorID}/solutions | guestIntrospectionSolutionsInfoRead
[**EndpointsecurityRegistrationVendorsGet**](EndpointsecurityApi.md#EndpointsecurityRegistrationVendorsGet) | **Get** /api/2.0/endpointsecurity/registration/vendors | guestIntrospectionVendorsInfoList

# **EndpointsecurityActivationGet**
> EndpointsecurityActivationGet(ctx, optional)
guestIntrospectionSolutionActivationRead

Retrieve activation information for all activated security VMs on the specified host.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EndpointsecurityApiEndpointsecurityActivationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsecurityApiEndpointsecurityActivationGetOpts struct
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

# **EndpointsecurityActivationVendorIDAltitudeMoidDelete**
> EndpointsecurityActivationVendorIDAltitudeMoidDelete(ctx, vendorID, altitude, moid)
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

# **EndpointsecurityActivationVendorIDAltitudeMoidGet**
> EndpointsecurityActivationVendorIDAltitudeMoidGet(ctx, vendorID, altitude, moid)
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

# **EndpointsecurityActivationVendorIDAltitudePost**
> EndpointsecurityActivationVendorIDAltitudePost(ctx, vendorID, altitude, optional)
guestIntrospectionSolutionActivateRegistered

Activate an endpoint protection solution that has been registered and located. Specify the following parameter in the request body.  | Name            | Comments | |-----------------|------------| |**svmMoid**     | Managed object ID of the virtual machine of the activated endpoint protection solution. |   Parameters:  vendorID: VMware-assigned ID for the vendor.  altitude: VMware-assigned number to uniquely identify a solution. Describes the type of solution and the order in which the solution receives events relative to other solutions on the same host.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
  **altitude** | **string**|  | 
 **optional** | ***EndpointsecurityApiEndpointsecurityActivationVendorIDAltitudePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsecurityApiEndpointsecurityActivationVendorIDAltitudePostOpts struct
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

# **EndpointsecurityActivationVendorIDSolutionIDGet**
> EndpointsecurityActivationVendorIDSolutionIDGet(ctx, vendorID, solutionID)
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

# **EndpointsecurityRegistrationPost**
> EndpointsecurityRegistrationPost(ctx, optional)
guestIntrospectionVendorCreate

Register the vendor of an endpoint protection solution. Specify the following parameters in the request.  | Name            | Comments | |-----------------|------------| |**vendorId**     | VMware-assigned ID for the vendor. | |**vendorTitle**  | Vendor-specified title. | |**vendorDescription** | Vendor-specified description. |   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EndpointsecurityApiEndpointsecurityRegistrationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsecurityApiEndpointsecurityRegistrationPostOpts struct
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

# **EndpointsecurityRegistrationVendorIDAltitudeDelete**
> EndpointsecurityRegistrationVendorIDAltitudeDelete(ctx, vendorID, altitude)
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

# **EndpointsecurityRegistrationVendorIDAltitudeGet**
> EndpointsecurityRegistrationVendorIDAltitudeGet(ctx, vendorID, altitude)
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

# **EndpointsecurityRegistrationVendorIDAltitudeLocationDelete**
> EndpointsecurityRegistrationVendorIDAltitudeLocationDelete(ctx, vendorID, altitude)
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

# **EndpointsecurityRegistrationVendorIDAltitudeLocationGet**
> EndpointsecurityRegistrationVendorIDAltitudeLocationGet(ctx, vendorID, altitude)
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

# **EndpointsecurityRegistrationVendorIDAltitudeLocationPost**
> EndpointsecurityRegistrationVendorIDAltitudeLocationPost(ctx, vendorID, altitude, optional)
guestIntrospectionSolutionIPPortUpdate

Set the IP address and port on the vNIC host for an endpoint protection solution.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
  **altitude** | **string**|  | 
 **optional** | ***EndpointsecurityApiEndpointsecurityRegistrationVendorIDAltitudeLocationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsecurityApiEndpointsecurityRegistrationVendorIDAltitudeLocationPostOpts struct
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

# **EndpointsecurityRegistrationVendorIDDelete**
> EndpointsecurityRegistrationVendorIDDelete(ctx, vendorID)
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

# **EndpointsecurityRegistrationVendorIDGet**
> EndpointsecurityRegistrationVendorIDGet(ctx, vendorID)
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

# **EndpointsecurityRegistrationVendorIDPost**
> EndpointsecurityRegistrationVendorIDPost(ctx, vendorID, optional)
guestIntrospectionSolutionCreate

Register an endpoint protection solution. Specify the following parameters in the request.  | Name            | Comments | |-----------------|------------| |**solutionAltitude**     | VMware-assigned altitude for the solution. *Altitude* is a number that VMware assigns to uniquely identify the solution. The altitude describes the type of solution and the order in which the solution receives events relative to other solutions on the same host. | |**solutionTitle**  | Vendor-specified title for the solution. | |**solutionDescription** | Vendor-specified description of the solution. |   Parameters:  vendorID: VMware-assigned ID for the vendor.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorID** | **string**|  | 
 **optional** | ***EndpointsecurityApiEndpointsecurityRegistrationVendorIDPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsecurityApiEndpointsecurityRegistrationVendorIDPostOpts struct
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

# **EndpointsecurityRegistrationVendorIDSolutionsGet**
> EndpointsecurityRegistrationVendorIDSolutionsGet(ctx, vendorID)
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

# **EndpointsecurityRegistrationVendorsGet**
> EndpointsecurityRegistrationVendorsGet(ctx, )
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

