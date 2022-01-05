# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api40ServicesSpoofguardPoliciesGet**](Class40SpoofguardApi.md#Api40ServicesSpoofguardPoliciesGet) | **Get** /api/4.0/services/spoofguard/policies | spoofGuardPoliciesList
[**Api40ServicesSpoofguardPoliciesPolicyIDDelete**](Class40SpoofguardApi.md#Api40ServicesSpoofguardPoliciesPolicyIDDelete) | **Delete** /api/4.0/services/spoofguard/policies/{policyID} | spoofGuardPolicyDelete
[**Api40ServicesSpoofguardPoliciesPolicyIDGet**](Class40SpoofguardApi.md#Api40ServicesSpoofguardPoliciesPolicyIDGet) | **Get** /api/4.0/services/spoofguard/policies/{policyID} | spoofGuardPolicyRead
[**Api40ServicesSpoofguardPoliciesPolicyIDPut**](Class40SpoofguardApi.md#Api40ServicesSpoofguardPoliciesPolicyIDPut) | **Put** /api/4.0/services/spoofguard/policies/{policyID} | spoofGuardPolicyUpdate
[**Api40ServicesSpoofguardPoliciesPost**](Class40SpoofguardApi.md#Api40ServicesSpoofguardPoliciesPost) | **Post** /api/4.0/services/spoofguard/policies | spoofGuardPoliciesCreate
[**Api40ServicesSpoofguardPolicyIDGet**](Class40SpoofguardApi.md#Api40ServicesSpoofguardPolicyIDGet) | **Get** /api/4.0/services/spoofguard/{policyID} | spoofGuardPolicyIPRead
[**Api40ServicesSpoofguardPolicyIDPost**](Class40SpoofguardApi.md#Api40ServicesSpoofguardPolicyIDPost) | **Post** /api/4.0/services/spoofguard/{policyID} | spoofGuardPolicyIPAction

# **Api40ServicesSpoofguardPoliciesGet**
> Api40ServicesSpoofguardPoliciesGet(ctx, )
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

# **Api40ServicesSpoofguardPoliciesPolicyIDDelete**
> Api40ServicesSpoofguardPoliciesPolicyIDDelete(ctx, policyID)
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

# **Api40ServicesSpoofguardPoliciesPolicyIDGet**
> Api40ServicesSpoofguardPoliciesPolicyIDGet(ctx, policyID)
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

# **Api40ServicesSpoofguardPoliciesPolicyIDPut**
> Api40ServicesSpoofguardPoliciesPolicyIDPut(ctx, policyID, optional)
spoofGuardPolicyUpdate

Modify the specified SpoofGuard policy.  Parameters:  policyID: SpoofGuard policy ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyID** | **string**|  | 
 **optional** | ***Class40SpoofguardApiApi40ServicesSpoofguardPoliciesPolicyIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class40SpoofguardApiApi40ServicesSpoofguardPoliciesPolicyIDPutOpts struct
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

# **Api40ServicesSpoofguardPoliciesPost**
> Api40ServicesSpoofguardPoliciesPost(ctx, optional)
spoofGuardPoliciesCreate

Create a SpoofGuard policy to specify the operation mode for networks.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Class40SpoofguardApiApi40ServicesSpoofguardPoliciesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class40SpoofguardApiApi40ServicesSpoofguardPoliciesPostOpts struct
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

# **Api40ServicesSpoofguardPolicyIDGet**
> Api40ServicesSpoofguardPolicyIDGet(ctx, policyID, optional)
spoofGuardPolicyIPRead

Retrieve IP addresses for the specified state.   Parameters:  policyID: SpoofGuard policy ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyID** | **string**|  | 
 **optional** | ***Class40SpoofguardApiApi40ServicesSpoofguardPolicyIDGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class40SpoofguardApiApi40ServicesSpoofguardPolicyIDGetOpts struct
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

# **Api40ServicesSpoofguardPolicyIDPost**
> Api40ServicesSpoofguardPolicyIDPost(ctx, policyID, optional)
spoofGuardPolicyIPAction

Approve or publish IP addresses.  Parameters:  policyID: SpoofGuard policy ID.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyID** | **string**|  | 
 **optional** | ***Class40SpoofguardApiApi40ServicesSpoofguardPolicyIDPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Class40SpoofguardApiApi40ServicesSpoofguardPolicyIDPostOpts struct
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

