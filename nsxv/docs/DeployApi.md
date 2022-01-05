# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api20SiDeployClusterClusterIDDelete**](DeployApi.md#Api20SiDeployClusterClusterIDDelete) | **Delete** /api/2.0/si/deploy/cluster/{clusterID} | serviceClusterDelete
[**Api20SiDeployClusterClusterIDGet**](DeployApi.md#Api20SiDeployClusterClusterIDGet) | **Get** /api/2.0/si/deploy/cluster/{clusterID} | serviceClusterRead
[**Api20SiDeployClusterClusterIDServiceServiceIDGet**](DeployApi.md#Api20SiDeployClusterClusterIDServiceServiceIDGet) | **Get** /api/2.0/si/deploy/cluster/{clusterID}/service/{serviceID} | serviceDetailsRead
[**Api20SiDeployPost**](DeployApi.md#Api20SiDeployPost) | **Post** /api/2.0/si/deploy | securityFabricCreate
[**Api20SiDeployPut**](DeployApi.md#Api20SiDeployPut) | **Put** /api/2.0/si/deploy | serviceUpgrade
[**Api20SiDeployServiceServiceIDDelete**](DeployApi.md#Api20SiDeployServiceServiceIDDelete) | **Delete** /api/2.0/si/deploy/service/{serviceID} | serviceDelete
[**Api20SiDeployServiceServiceIDDependsOnGet**](DeployApi.md#Api20SiDeployServiceServiceIDDependsOnGet) | **Get** /api/2.0/si/deploy/service/{serviceID}/dependsOn | serviceDependencyRead
[**Api20SiDeployServiceServiceIDGet**](DeployApi.md#Api20SiDeployServiceServiceIDGet) | **Get** /api/2.0/si/deploy/service/{serviceID} | serviceRead

# **Api20SiDeployClusterClusterIDDelete**
> Api20SiDeployClusterClusterIDDelete(ctx, clusterID, optional)
serviceClusterDelete

Uninstall a service. Fails if you try to remove a service that another service depends on.  In order to uninstall services in any order, set parameter ignoreDependency to true.   Parameters:  clusterID: Cluster ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterID** | **string**|  | 
 **optional** | ***DeployApiApi20SiDeployClusterClusterIDDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiApi20SiDeployClusterClusterIDDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **services** | **optional.String**|  | 
 **startTime** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20SiDeployClusterClusterIDGet**
> Api20SiDeployClusterClusterIDGet(ctx, clusterID)
serviceClusterRead

Retrieve all services deployed along with their status.  Parameters:  clusterID: Cluster ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20SiDeployClusterClusterIDServiceServiceIDGet**
> Api20SiDeployClusterClusterIDServiceServiceIDGet(ctx, clusterID, serviceID)
serviceDetailsRead

Retrieve detailed information about the service.  Parameters:  serviceID: Service ID on cluster  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterID** | **string**|  | 
  **serviceID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20SiDeployPost**
> Api20SiDeployPost(ctx, optional)
securityFabricCreate

Deploy security fabric.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeployApiApi20SiDeployPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiApi20SiDeployPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **startTime** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20SiDeployPut**
> Api20SiDeployPut(ctx, optional)
serviceUpgrade

Upgrade service to recent version.  The datastore, dvPortGroup, and ipPool variables should either not be specified or have same value as provided at time of deployment.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeployApiApi20SiDeployPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiApi20SiDeployPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **startTime** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20SiDeployServiceServiceIDDelete**
> Api20SiDeployServiceServiceIDDelete(ctx, serviceID, optional)
serviceDelete

Uninstall specified service from specified clusters.  Parameters:  serviceID: Specified service.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceID** | **string**|  | 
 **optional** | ***DeployApiApi20SiDeployServiceServiceIDDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiApi20SiDeployServiceServiceIDDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusters** | **optional.String**|  | 
 **startTime** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20SiDeployServiceServiceIDDependsOnGet**
> Api20SiDeployServiceServiceIDDependsOnGet(ctx, serviceID)
serviceDependencyRead

Retrieve service on which the specified service depends.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Api20SiDeployServiceServiceIDGet**
> Api20SiDeployServiceServiceIDGet(ctx, serviceID)
serviceRead

Retrieve all clusters on which the service is installed.  Parameters:  serviceID: Specified service.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

