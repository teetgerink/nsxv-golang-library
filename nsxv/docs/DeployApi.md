# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SiDeployClusterClusterIDDelete**](DeployApi.md#SiDeployClusterClusterIDDelete) | **Delete** /api/2.0/si/deploy/cluster/{clusterID} | serviceClusterDelete
[**SiDeployClusterClusterIDGet**](DeployApi.md#SiDeployClusterClusterIDGet) | **Get** /api/2.0/si/deploy/cluster/{clusterID} | serviceClusterRead
[**SiDeployClusterClusterIDServiceServiceIDGet**](DeployApi.md#SiDeployClusterClusterIDServiceServiceIDGet) | **Get** /api/2.0/si/deploy/cluster/{clusterID}/service/{serviceID} | serviceDetailsRead
[**SiDeployPost**](DeployApi.md#SiDeployPost) | **Post** /api/2.0/si/deploy | securityFabricCreate
[**SiDeployPut**](DeployApi.md#SiDeployPut) | **Put** /api/2.0/si/deploy | serviceUpgrade
[**SiDeployServiceServiceIDDelete**](DeployApi.md#SiDeployServiceServiceIDDelete) | **Delete** /api/2.0/si/deploy/service/{serviceID} | serviceDelete
[**SiDeployServiceServiceIDDependsOnGet**](DeployApi.md#SiDeployServiceServiceIDDependsOnGet) | **Get** /api/2.0/si/deploy/service/{serviceID}/dependsOn | serviceDependencyRead
[**SiDeployServiceServiceIDGet**](DeployApi.md#SiDeployServiceServiceIDGet) | **Get** /api/2.0/si/deploy/service/{serviceID} | serviceRead

# **SiDeployClusterClusterIDDelete**
> SiDeployClusterClusterIDDelete(ctx, clusterID, optional)
serviceClusterDelete

Uninstall a service. Fails if you try to remove a service that another service depends on.  In order to uninstall services in any order, set parameter ignoreDependency to true.   Parameters:  clusterID: Cluster ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterID** | **string**|  | 
 **optional** | ***DeployApiSiDeployClusterClusterIDDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiSiDeployClusterClusterIDDeleteOpts struct
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

# **SiDeployClusterClusterIDGet**
> SiDeployClusterClusterIDGet(ctx, clusterID)
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

# **SiDeployClusterClusterIDServiceServiceIDGet**
> SiDeployClusterClusterIDServiceServiceIDGet(ctx, clusterID, serviceID)
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

# **SiDeployPost**
> SiDeployPost(ctx, optional)
securityFabricCreate

Deploy security fabric.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeployApiSiDeployPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiSiDeployPostOpts struct
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

# **SiDeployPut**
> SiDeployPut(ctx, optional)
serviceUpgrade

Upgrade service to recent version.  The datastore, dvPortGroup, and ipPool variables should either not be specified or have same value as provided at time of deployment.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeployApiSiDeployPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiSiDeployPutOpts struct
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

# **SiDeployServiceServiceIDDelete**
> SiDeployServiceServiceIDDelete(ctx, serviceID, optional)
serviceDelete

Uninstall specified service from specified clusters.  Parameters:  serviceID: Specified service.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceID** | **string**|  | 
 **optional** | ***DeployApiSiDeployServiceServiceIDDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployApiSiDeployServiceServiceIDDeleteOpts struct
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

# **SiDeployServiceServiceIDDependsOnGet**
> SiDeployServiceServiceIDDependsOnGet(ctx, serviceID)
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

# **SiDeployServiceServiceIDGet**
> SiDeployServiceServiceIDGet(ctx, serviceID)
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

