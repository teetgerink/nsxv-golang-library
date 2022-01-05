# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesSpoofguardPoliciesGet**](SpoofguardApi.md#ServicesSpoofguardPoliciesGet) | **Get** /api/4.0/services/spoofguard/policies | spoofGuardPoliciesList
[**ServicesSpoofguardPoliciesPolicyIDDelete**](SpoofguardApi.md#ServicesSpoofguardPoliciesPolicyIDDelete) | **Delete** /api/4.0/services/spoofguard/policies/{policyID} | spoofGuardPolicyDelete
[**ServicesSpoofguardPoliciesPolicyIDGet**](SpoofguardApi.md#ServicesSpoofguardPoliciesPolicyIDGet) | **Get** /api/4.0/services/spoofguard/policies/{policyID} | spoofGuardPolicyRead
[**ServicesSpoofguardPoliciesPolicyIDPut**](SpoofguardApi.md#ServicesSpoofguardPoliciesPolicyIDPut) | **Put** /api/4.0/services/spoofguard/policies/{policyID} | spoofGuardPolicyUpdate
[**ServicesSpoofguardPoliciesPost**](SpoofguardApi.md#ServicesSpoofguardPoliciesPost) | **Post** /api/4.0/services/spoofguard/policies | spoofGuardPoliciesCreate
[**ServicesSpoofguardPolicyIDGet**](SpoofguardApi.md#ServicesSpoofguardPolicyIDGet) | **Get** /api/4.0/services/spoofguard/{policyID} | spoofGuardPolicyIPRead
[**ServicesSpoofguardPolicyIDPost**](SpoofguardApi.md#ServicesSpoofguardPolicyIDPost) | **Post** /api/4.0/services/spoofguard/{policyID} | spoofGuardPolicyIPAction

# **ServicesSpoofguardPoliciesGet**
> ServicesSpoofguardPoliciesGet(ctx, )
spoofGuardPoliciesList

Retrieve information about all SpoofGuard policies.  Parameters:  

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

# **ServicesSpoofguardPoliciesPolicyIDDelete**
> ServicesSpoofguardPoliciesPolicyIDDelete(ctx, policyID)
spoofGuardPolicyDelete

Delete the specified SpoofGuard policy.  Parameters:  policyID: SpoofGuard policy ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSpoofguardPoliciesPolicyIDGet**
> ServicesSpoofguardPoliciesPolicyIDGet(ctx, policyID)
spoofGuardPolicyRead

Retrieve information about the specified SpoofGuard policy.   Parameters:  policyID: SpoofGuard policy ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSpoofguardPoliciesPolicyIDPut**
> ServicesSpoofguardPoliciesPolicyIDPut(ctx, policyID, optional)
spoofGuardPolicyUpdate

Modify the specified SpoofGuard policy.  Parameters:  policyID: SpoofGuard policy ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyID** | **string**|  | 
 **optional** | ***SpoofguardApiServicesSpoofguardPoliciesPolicyIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoofguardApiServicesSpoofguardPoliciesPolicyIDPutOpts struct
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

# **ServicesSpoofguardPoliciesPost**
> ServicesSpoofguardPoliciesPost(ctx, optional)
spoofGuardPoliciesCreate

Create a SpoofGuard policy to specify the operation mode for networks.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpoofguardApiServicesSpoofguardPoliciesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoofguardApiServicesSpoofguardPoliciesPostOpts struct
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

# **ServicesSpoofguardPolicyIDGet**
> ServicesSpoofguardPolicyIDGet(ctx, policyID, optional)
spoofGuardPolicyIPRead

Retrieve IP addresses for the specified state.   Parameters:  policyID: SpoofGuard policy ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyID** | **string**|  | 
 **optional** | ***SpoofguardApiServicesSpoofguardPolicyIDGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoofguardApiServicesSpoofguardPolicyIDGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **list** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicesSpoofguardPolicyIDPost**
> ServicesSpoofguardPolicyIDPost(ctx, policyID, optional)
spoofGuardPolicyIPAction

Approve or publish IP addresses.  Parameters:  policyID: SpoofGuard policy ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyID** | **string**|  | 
 **optional** | ***SpoofguardApiServicesSpoofguardPolicyIDPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoofguardApiServicesSpoofguardPolicyIDPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **vnicId** | **optional.**|  | 
 **action** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

