# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Api10NsxCliPost**](CliApi.md#Api10NsxCliPost) | **Post** /api/1.0/nsx/cli | nsxCliExecute

# **Api10NsxCliPost**
> Api10NsxCliPost(ctx, optional)
nsxCliExecute

The central command-line interface (central CLI) commands are run from the NSX Manager command line, and retrieve information from the NSX Manager and other devices. These commands can also be executed in the API.  You can insert any valid Central CLI command as the **command** parameter. For a complete list of the Central CLI commands executable through the API, please see the Central CLI chapter of the *NSX Command Line Interface Reference*.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CliApiApi10NsxCliPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CliApiApi10NsxCliPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **accept** | **optional.**|  | 
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

