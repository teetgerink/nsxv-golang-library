# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesTruststoreCertificateCertificateIdDelete**](TruststoreApi.md#ServicesTruststoreCertificateCertificateIdDelete) | **Delete** /api/2.0/services/truststore/certificate/{certificateId} | certificateIdDelete
[**ServicesTruststoreCertificateCertificateIdGet**](TruststoreApi.md#ServicesTruststoreCertificateCertificateIdGet) | **Get** /api/2.0/services/truststore/certificate/{certificateId} | certificateIdRead
[**ServicesTruststoreCertificatePost**](TruststoreApi.md#ServicesTruststoreCertificatePost) | **Post** /api/2.0/services/truststore/certificate | certificateCreate
[**ServicesTruststoreCertificateScopeIdPost**](TruststoreApi.md#ServicesTruststoreCertificateScopeIdPost) | **Post** /api/2.0/services/truststore/certificate/{scopeId} | certificateSelfSignedCreate
[**ServicesTruststoreCertificateScopeScopeIdGet**](TruststoreApi.md#ServicesTruststoreCertificateScopeScopeIdGet) | **Get** /api/2.0/services/truststore/certificate/scope/{scopeId} | certificateScopeRead
[**ServicesTruststoreCrlCrlIdDelete**](TruststoreApi.md#ServicesTruststoreCrlCrlIdDelete) | **Delete** /api/2.0/services/truststore/crl/{crlId} | crlIDDelete
[**ServicesTruststoreCrlCrlIdGet**](TruststoreApi.md#ServicesTruststoreCrlCrlIdGet) | **Get** /api/2.0/services/truststore/crl/{crlId} | crlIDRead
[**ServicesTruststoreCrlScopeIdPost**](TruststoreApi.md#ServicesTruststoreCrlScopeIdPost) | **Post** /api/2.0/services/truststore/crl/{scopeId} | crlCreate
[**ServicesTruststoreCrlScopeScopeIdGet**](TruststoreApi.md#ServicesTruststoreCrlScopeScopeIdGet) | **Get** /api/2.0/services/truststore/crl/scope/{scopeId} | crlScopeRead
[**ServicesTruststoreCsrCsrIdGet**](TruststoreApi.md#ServicesTruststoreCsrCsrIdGet) | **Get** /api/2.0/services/truststore/csr/{csrId} | csrSelfSignedRead
[**ServicesTruststoreCsrCsrIdPut**](TruststoreApi.md#ServicesTruststoreCsrCsrIdPut) | **Put** /api/2.0/services/truststore/csr/{csrId} | csrSelfSignedUpdate
[**ServicesTruststoreCsrScopeIdPost**](TruststoreApi.md#ServicesTruststoreCsrScopeIdPost) | **Post** /api/2.0/services/truststore/csr/{scopeId} | csrCreate
[**ServicesTruststoreCsrScopeScopeIdGet**](TruststoreApi.md#ServicesTruststoreCsrScopeScopeIdGet) | **Get** /api/2.0/services/truststore/csr/scope/{scopeId} | csrScopeRead

# **ServicesTruststoreCertificateCertificateIdDelete**
> ServicesTruststoreCertificateCertificateIdDelete(ctx, certificateId)
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

# **ServicesTruststoreCertificateCertificateIdGet**
> ServicesTruststoreCertificateCertificateIdGet(ctx, certificateId)
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

# **ServicesTruststoreCertificatePost**
> ServicesTruststoreCertificatePost(ctx, optional)
certificateCreate

Import a certificate or a certificate chain against a certificate signing request.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TruststoreApiServicesTruststoreCertificatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TruststoreApiServicesTruststoreCertificatePostOpts struct
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

# **ServicesTruststoreCertificateScopeIdPost**
> ServicesTruststoreCertificateScopeIdPost(ctx, scopeId, optional)
certificateSelfSignedCreate

Create a single certificate  You can create a certificate for a specific NSX Edge, or if you specify a scope of *globalroot-0* you can create a global certificate in NSX Manager which is available to all NSX Edges.   Parameters:  scopeId: Scope ID. Specify the ID of an NSX Edge, e.g. *edge-5*, or *globalroot-0*.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***TruststoreApiServicesTruststoreCertificateScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TruststoreApiServicesTruststoreCertificateScopeIdPostOpts struct
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

# **ServicesTruststoreCertificateScopeScopeIdGet**
> ServicesTruststoreCertificateScopeScopeIdGet(ctx, scopeId)
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

# **ServicesTruststoreCrlCrlIdDelete**
> ServicesTruststoreCrlCrlIdDelete(ctx, crlId)
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

# **ServicesTruststoreCrlCrlIdGet**
> ServicesTruststoreCrlCrlIdGet(ctx, crlId)
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

# **ServicesTruststoreCrlScopeIdPost**
> ServicesTruststoreCrlScopeIdPost(ctx, scopeId, optional)
crlCreate

Create a certificate revocation list (CRL) on the specified scope.   Parameters:  scopeId: Specified scope.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***TruststoreApiServicesTruststoreCrlScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TruststoreApiServicesTruststoreCrlScopeIdPostOpts struct
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

# **ServicesTruststoreCrlScopeScopeIdGet**
> ServicesTruststoreCrlScopeScopeIdGet(ctx, scopeId)
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

# **ServicesTruststoreCsrCsrIdGet**
> ServicesTruststoreCsrCsrIdGet(ctx, csrId)
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

# **ServicesTruststoreCsrCsrIdPut**
> ServicesTruststoreCsrCsrIdPut(ctx, csrId, optional)
csrSelfSignedUpdate

Create a self-signed certificate for CSR.   Parameters:  csrId: CSR ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **csrId** | **string**|  | 
 **optional** | ***TruststoreApiServicesTruststoreCsrCsrIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TruststoreApiServicesTruststoreCsrCsrIdPutOpts struct
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

# **ServicesTruststoreCsrScopeIdPost**
> ServicesTruststoreCsrScopeIdPost(ctx, scopeId, optional)
csrCreate

Create a certificate signing request (CSR).  Parameters:  scopeId: Specified scope ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | **string**|  | 
 **optional** | ***TruststoreApiServicesTruststoreCsrScopeIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TruststoreApiServicesTruststoreCsrScopeIdPostOpts struct
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

# **ServicesTruststoreCsrScopeScopeIdGet**
> ServicesTruststoreCsrScopeScopeIdGet(ctx, scopeId)
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

