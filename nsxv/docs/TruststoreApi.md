# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20ServicesTruststoreCertificateCertificateIdDelete**](TruststoreApi.md#Api20ServicesTruststoreCertificateCertificateIdDelete) | **Delete** /api/2.0/services/truststore/certificate/{certificateId} | certificateIdDelete
[**Api20ServicesTruststoreCertificateCertificateIdGet**](TruststoreApi.md#Api20ServicesTruststoreCertificateCertificateIdGet) | **Get** /api/2.0/services/truststore/certificate/{certificateId} | certificateIdRead
[**Api20ServicesTruststoreCertificatePost**](TruststoreApi.md#Api20ServicesTruststoreCertificatePost) | **Post** /api/2.0/services/truststore/certificate | certificateCreate
[**Api20ServicesTruststoreCertificateScopeIdPost**](TruststoreApi.md#Api20ServicesTruststoreCertificateScopeIdPost) | **Post** /api/2.0/services/truststore/certificate/{scopeId} | certificateSelfSignedCreate
[**Api20ServicesTruststoreCertificateScopeScopeIdGet**](TruststoreApi.md#Api20ServicesTruststoreCertificateScopeScopeIdGet) | **Get** /api/2.0/services/truststore/certificate/scope/{scopeId} | certificateScopeRead
[**Api20ServicesTruststoreCrlCrlIdDelete**](TruststoreApi.md#Api20ServicesTruststoreCrlCrlIdDelete) | **Delete** /api/2.0/services/truststore/crl/{crlId} | crlIDDelete
[**Api20ServicesTruststoreCrlCrlIdGet**](TruststoreApi.md#Api20ServicesTruststoreCrlCrlIdGet) | **Get** /api/2.0/services/truststore/crl/{crlId} | crlIDRead
[**Api20ServicesTruststoreCrlScopeIdPost**](TruststoreApi.md#Api20ServicesTruststoreCrlScopeIdPost) | **Post** /api/2.0/services/truststore/crl/{scopeId} | crlCreate
[**Api20ServicesTruststoreCrlScopeScopeIdGet**](TruststoreApi.md#Api20ServicesTruststoreCrlScopeScopeIdGet) | **Get** /api/2.0/services/truststore/crl/scope/{scopeId} | crlScopeRead
[**Api20ServicesTruststoreCsrCsrIdGet**](TruststoreApi.md#Api20ServicesTruststoreCsrCsrIdGet) | **Get** /api/2.0/services/truststore/csr/{csrId} | csrSelfSignedRead
[**Api20ServicesTruststoreCsrCsrIdPut**](TruststoreApi.md#Api20ServicesTruststoreCsrCsrIdPut) | **Put** /api/2.0/services/truststore/csr/{csrId} | csrSelfSignedUpdate
[**Api20ServicesTruststoreCsrScopeIdPost**](TruststoreApi.md#Api20ServicesTruststoreCsrScopeIdPost) | **Post** /api/2.0/services/truststore/csr/{scopeId} | csrCreate
[**Api20ServicesTruststoreCsrScopeScopeIdGet**](TruststoreApi.md#Api20ServicesTruststoreCsrScopeScopeIdGet) | **Get** /api/2.0/services/truststore/csr/scope/{scopeId} | csrScopeRead

# **Api20ServicesTruststoreCertificateCertificateIdDelete**
> Api20ServicesTruststoreCertificateCertificateIdDelete(ctx, certificateId)
certificateIdDelete

Delete the specified certificate.  Parameters:  certificateId: Certificate ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesTruststoreCertificateCertificateIdGet**
> Api20ServicesTruststoreCertificateCertificateIdGet(ctx, certificateId)
certificateIdRead

Retrieve the certificate object specified by ID. If the ID specifies a chain, multiple certificate objects are retrieved.   Parameters:  certificateId: Certificate ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesTruststoreCertificatePost**
> Api20ServicesTruststoreCertificatePost(ctx, optional)
certificateCreate

Import a certificate or a certificate chain against a certificate signing request.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TruststoreApiApi20ServicesTruststoreCertificatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TruststoreApiApi20ServicesTruststoreCertificatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **csrId** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesTruststoreCertificateScopeIdPost**
> Api20ServicesTruststoreCertificateScopeIdPost(ctx, scopeId, optional)
certificateSelfSignedCreate

Create a single certificate  You can create a certificate for a specific NSX Edge, or if you specify a scope of *globalroot-0* you can create a global certificate in NSX Manager which is available to all NSX Edges.   Parameters:  scopeId: Scope ID. Specify the ID of an NSX Edge, e.g. *edge-5*, or *globalroot-0*.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***TruststoreApiApi20ServicesTruststoreCertificateScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TruststoreApiApi20ServicesTruststoreCertificateScopeIdPostOpts struct
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

# **Api20ServicesTruststoreCertificateScopeScopeIdGet**
> Api20ServicesTruststoreCertificateScopeScopeIdGet(ctx, scopeId)
certificateScopeRead

Retrieve all certificates on the specified scope.  Parameters:  scopeId: Scope ID. Specify the ID of an NSX Edge, e.g. *edge-5*, or *globalroot-0*.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesTruststoreCrlCrlIdDelete**
> Api20ServicesTruststoreCrlCrlIdDelete(ctx, crlId)
crlIDDelete

Delete the specified certificate revocation list (CRL).  Parameters:  crlId: CRL ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crlId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesTruststoreCrlCrlIdGet**
> Api20ServicesTruststoreCrlCrlIdGet(ctx, crlId)
crlIDRead

Retrieve certificate object for the specified certificate revocation list (CRL).   Parameters:  crlId: CRL ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crlId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesTruststoreCrlScopeIdPost**
> Api20ServicesTruststoreCrlScopeIdPost(ctx, scopeId, optional)
crlCreate

Create a certificate revocation list (CRL) on the specified scope.   Parameters:  scopeId: Specified scope.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***TruststoreApiApi20ServicesTruststoreCrlScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TruststoreApiApi20ServicesTruststoreCrlScopeIdPostOpts struct
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

# **Api20ServicesTruststoreCrlScopeScopeIdGet**
> Api20ServicesTruststoreCrlScopeScopeIdGet(ctx, scopeId)
crlScopeRead

Retrieve all certificates for the specified scope.  Parameters:  scopeId: Specified scope  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesTruststoreCsrCsrIdGet**
> Api20ServicesTruststoreCsrCsrIdGet(ctx, csrId)
csrSelfSignedRead

Retrieve the specified certificate signing request (CSR).   Parameters:  csrId: CSR ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **csrId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesTruststoreCsrCsrIdPut**
> Api20ServicesTruststoreCsrCsrIdPut(ctx, csrId, optional)
csrSelfSignedUpdate

Create a self-signed certificate for CSR.   Parameters:  csrId: CSR ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **csrId** | **string**|  | 
 **optional** | ***TruststoreApiApi20ServicesTruststoreCsrCsrIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TruststoreApiApi20ServicesTruststoreCsrCsrIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noOfDays** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20ServicesTruststoreCsrScopeIdPost**
> Api20ServicesTruststoreCsrScopeIdPost(ctx, scopeId, optional)
csrCreate

Create a certificate signing request (CSR).  Parameters:  scopeId: Specified scope ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***TruststoreApiApi20ServicesTruststoreCsrScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TruststoreApiApi20ServicesTruststoreCsrScopeIdPostOpts struct
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

# **Api20ServicesTruststoreCsrScopeScopeIdGet**
> Api20ServicesTruststoreCsrScopeScopeIdGet(ctx, scopeId)
csrScopeRead

Retrieve certificate signing requests (CSR) on the specified scope.  Parameters:  scopeId: Specified scope.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

