# {{classname}}

All URIs are relative to *https://{{nsxmanager}}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FirewallConfigGlobalconfigurationGet**](FirewallApi.md#FirewallConfigGlobalconfigurationGet) | **Get** /api/4.0/firewall/config/globalconfiguration | dfwPerformanceRead
[**FirewallConfigGlobalconfigurationPut**](FirewallApi.md#FirewallConfigGlobalconfigurationPut) | **Put** /api/4.0/firewall/config/globalconfiguration | dfwPerformanceUpdate
[**FirewallContextIdConfigIpfixDelete**](FirewallApi.md#FirewallContextIdConfigIpfixDelete) | **Delete** /api/4.0/firewall/{contextId}/config/ipfix | dfwIPFixDelete
[**FirewallContextIdConfigIpfixGet**](FirewallApi.md#FirewallContextIdConfigIpfixGet) | **Get** /api/4.0/firewall/{contextId}/config/ipfix | dfwIPFixRead
[**FirewallContextIdConfigIpfixPut**](FirewallApi.md#FirewallContextIdConfigIpfixPut) | **Put** /api/4.0/firewall/{contextId}/config/ipfix | dfwIPFixUpdate
[**FirewallDomainIDEnableTruefalsePost**](FirewallApi.md#FirewallDomainIDEnableTruefalsePost) | **Post** /api/4.0/firewall/{domainID}/enable/{truefalse} | dfwEnableDisableToggle
[**FirewallForceSyncIDPost**](FirewallApi.md#FirewallForceSyncIDPost) | **Post** /api/4.0/firewall/forceSync/{ID} | dfwSyncExecute
[**FirewallGlobalroot0ConfigDelete**](FirewallApi.md#FirewallGlobalroot0ConfigDelete) | **Delete** /api/4.0/firewall/globalroot-0/config | dfwConfigDelete
[**FirewallGlobalroot0ConfigGet**](FirewallApi.md#FirewallGlobalroot0ConfigGet) | **Get** /api/4.0/firewall/globalroot-0/config | dfwConfigShow
[**FirewallGlobalroot0ConfigLayer2sectionsGet**](FirewallApi.md#FirewallGlobalroot0ConfigLayer2sectionsGet) | **Get** /api/4.0/firewall/globalroot-0/config/layer2sections | dfwL2SectionRead
[**FirewallGlobalroot0ConfigLayer2sectionsPost**](FirewallApi.md#FirewallGlobalroot0ConfigLayer2sectionsPost) | **Post** /api/4.0/firewall/globalroot-0/config/layer2sections | dfwL2SectionCreate
[**FirewallGlobalroot0ConfigLayer2sectionsSectionIdDelete**](FirewallApi.md#FirewallGlobalroot0ConfigLayer2sectionsSectionIdDelete) | **Delete** /api/4.0/firewall/globalroot-0/config/layer2sections/{sectionId} | dfwL2SectionIdDelete
[**FirewallGlobalroot0ConfigLayer2sectionsSectionIdGet**](FirewallApi.md#FirewallGlobalroot0ConfigLayer2sectionsSectionIdGet) | **Get** /api/4.0/firewall/globalroot-0/config/layer2sections/{sectionId} | dfwL2SectionIdRead
[**FirewallGlobalroot0ConfigLayer2sectionsSectionIdPost**](FirewallApi.md#FirewallGlobalroot0ConfigLayer2sectionsSectionIdPost) | **Post** /api/4.0/firewall/globalroot-0/config/layer2sections/{sectionId} | dfwL2SectionIdAction
[**FirewallGlobalroot0ConfigLayer2sectionsSectionIdPut**](FirewallApi.md#FirewallGlobalroot0ConfigLayer2sectionsSectionIdPut) | **Put** /api/4.0/firewall/globalroot-0/config/layer2sections/{sectionId} | dfwL2SectionIdUpdate
[**FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesPost**](FirewallApi.md#FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesPost) | **Post** /api/4.0/firewall/globalroot-0/config/layer2sections/{sectionId}/rules | dfwL2RulesAdd
[**FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdDelete**](FirewallApi.md#FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdDelete) | **Delete** /api/4.0/firewall/globalroot-0/config/layer2sections/{sectionId}/rules/{ruleId} | dfwL2RuleDelete
[**FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdGet**](FirewallApi.md#FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdGet) | **Get** /api/4.0/firewall/globalroot-0/config/layer2sections/{sectionId}/rules/{ruleId} | dfwL2RuleRead
[**FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdPut**](FirewallApi.md#FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdPut) | **Put** /api/4.0/firewall/globalroot-0/config/layer2sections/{sectionId}/rules/{ruleId} | dfwL2RuleUpdate
[**FirewallGlobalroot0ConfigLayer3redirectProfilesGet**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3redirectProfilesGet) | **Get** /api/4.0/firewall/globalroot-0/config/layer3redirect/profiles | serviceInsertionProfilesRead
[**FirewallGlobalroot0ConfigLayer3redirectsectionsPost**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3redirectsectionsPost) | **Post** /api/4.0/firewall/globalroot-0/config/layer3redirectsections | layer3RedirectSectionsCreate
[**FirewallGlobalroot0ConfigLayer3redirectsectionsSectionDelete**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3redirectsectionsSectionDelete) | **Delete** /api/4.0/firewall/globalroot-0/config/layer3redirectsections/{section} | sectionDelete
[**FirewallGlobalroot0ConfigLayer3redirectsectionsSectionGet**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3redirectsectionsSectionGet) | **Get** /api/4.0/firewall/globalroot-0/config/layer3redirectsections/{section} | sectionRead
[**FirewallGlobalroot0ConfigLayer3redirectsectionsSectionPut**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3redirectsectionsSectionPut) | **Put** /api/4.0/firewall/globalroot-0/config/layer3redirectsections/{section} | sectionUpdate
[**FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesPost**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesPost) | **Post** /api/4.0/firewall/globalroot-0/config/layer3redirectsections/{section}/rules | rulesCreate
[**FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDDelete**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDDelete) | **Delete** /api/4.0/firewall/globalroot-0/config/layer3redirectsections/{section}/rules/{ruleID} | ruleDelete
[**FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDGet**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDGet) | **Get** /api/4.0/firewall/globalroot-0/config/layer3redirectsections/{section}/rules/{ruleID} | ruleRead
[**FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDPut**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDPut) | **Put** /api/4.0/firewall/globalroot-0/config/layer3redirectsections/{section}/rules/{ruleID} | ruleUpdate
[**FirewallGlobalroot0ConfigLayer3sectionsGet**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3sectionsGet) | **Get** /api/4.0/firewall/globalroot-0/config/layer3sections | dfwL3SectionRead
[**FirewallGlobalroot0ConfigLayer3sectionsPost**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3sectionsPost) | **Post** /api/4.0/firewall/globalroot-0/config/layer3sections | dfwL3SectionCreate
[**FirewallGlobalroot0ConfigLayer3sectionsSectionIdDelete**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3sectionsSectionIdDelete) | **Delete** /api/4.0/firewall/globalroot-0/config/layer3sections/{sectionId} | dfwL3SectionIdDelete
[**FirewallGlobalroot0ConfigLayer3sectionsSectionIdGet**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3sectionsSectionIdGet) | **Get** /api/4.0/firewall/globalroot-0/config/layer3sections/{sectionId} | dfwL3SectionIdRead
[**FirewallGlobalroot0ConfigLayer3sectionsSectionIdPost**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3sectionsSectionIdPost) | **Post** /api/4.0/firewall/globalroot-0/config/layer3sections/{sectionId} | dfwL3SectionIdAction
[**FirewallGlobalroot0ConfigLayer3sectionsSectionIdPut**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3sectionsSectionIdPut) | **Put** /api/4.0/firewall/globalroot-0/config/layer3sections/{sectionId} | dfwL3SectionIdUpdate
[**FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesPost**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesPost) | **Post** /api/4.0/firewall/globalroot-0/config/layer3sections/{sectionId}/rules | dfwL3RulesAdd
[**FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdDelete**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdDelete) | **Delete** /api/4.0/firewall/globalroot-0/config/layer3sections/{sectionId}/rules/{ruleId} | dfwL3RuleDelete
[**FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdGet**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdGet) | **Get** /api/4.0/firewall/globalroot-0/config/layer3sections/{sectionId}/rules/{ruleId} | dfwL3RuleRead
[**FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdPut**](FirewallApi.md#FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdPut) | **Put** /api/4.0/firewall/globalroot-0/config/layer3sections/{sectionId}/rules/{ruleId} | dfwL3RuleUpdate
[**FirewallGlobalroot0ConfigPut**](FirewallApi.md#FirewallGlobalroot0ConfigPut) | **Put** /api/4.0/firewall/globalroot-0/config | dfwConfigUpdate
[**FirewallGlobalroot0DefaultconfigGet**](FirewallApi.md#FirewallGlobalroot0DefaultconfigGet) | **Get** /api/4.0/firewall/globalroot-0/defaultconfig | dfwConfigDefaultRead
[**FirewallGlobalroot0DraftsDraftIDActionExportGet**](FirewallApi.md#FirewallGlobalroot0DraftsDraftIDActionExportGet) | **Get** /api/4.0/firewall/globalroot-0/drafts/{draftID}/action/export | dfwConfigExportRead
[**FirewallGlobalroot0DraftsDraftIDActionImportPost**](FirewallApi.md#FirewallGlobalroot0DraftsDraftIDActionImportPost) | **Post** /api/4.0/firewall/globalroot-0/drafts/{draftID}/action/import | dfwConfigImportCreate
[**FirewallGlobalroot0DraftsDraftIDDelete**](FirewallApi.md#FirewallGlobalroot0DraftsDraftIDDelete) | **Delete** /api/4.0/firewall/globalroot-0/drafts/{draftID} | dfwDraftDelete
[**FirewallGlobalroot0DraftsDraftIDGet**](FirewallApi.md#FirewallGlobalroot0DraftsDraftIDGet) | **Get** /api/4.0/firewall/globalroot-0/drafts/{draftID} | dfwDraftRead
[**FirewallGlobalroot0DraftsDraftIDPut**](FirewallApi.md#FirewallGlobalroot0DraftsDraftIDPut) | **Put** /api/4.0/firewall/globalroot-0/drafts/{draftID} | dfwDraftUpdate
[**FirewallGlobalroot0DraftsGet**](FirewallApi.md#FirewallGlobalroot0DraftsGet) | **Get** /api/4.0/firewall/globalroot-0/drafts | dfwSaveConfigRead
[**FirewallGlobalroot0DraftsPost**](FirewallApi.md#FirewallGlobalroot0DraftsPost) | **Post** /api/4.0/firewall/globalroot-0/drafts | dfwDraftsCreate
[**FirewallGlobalroot0StateGet**](FirewallApi.md#FirewallGlobalroot0StateGet) | **Get** /api/4.0/firewall/globalroot-0/state | dfwUpgradeRead
[**FirewallGlobalroot0StatePut**](FirewallApi.md#FirewallGlobalroot0StatePut) | **Put** /api/4.0/firewall/globalroot-0/state | dfwUpgradeEnable
[**FirewallGlobalroot0StatusGet**](FirewallApi.md#FirewallGlobalroot0StatusGet) | **Get** /api/4.0/firewall/globalroot-0/status | dfwStatusRead
[**FirewallGlobalroot0StatusLayer2sectionsSectionIDGet**](FirewallApi.md#FirewallGlobalroot0StatusLayer2sectionsSectionIDGet) | **Get** /api/4.0/firewall/globalroot-0/status/layer2sections/{sectionID} | L2SectionStatusRead
[**FirewallGlobalroot0StatusLayer3sectionsSectionIDGet**](FirewallApi.md#FirewallGlobalroot0StatusLayer3sectionsSectionIDGet) | **Get** /api/4.0/firewall/globalroot-0/status/layer3sections/{sectionID} | L3SectionStatusRead
[**FirewallGlobalroot0TimeoutsConfigIdDelete**](FirewallApi.md#FirewallGlobalroot0TimeoutsConfigIdDelete) | **Delete** /api/4.0/firewall/globalroot-0/timeouts/{configId} | dfwTimeoutConfigIdDelete
[**FirewallGlobalroot0TimeoutsConfigIdGet**](FirewallApi.md#FirewallGlobalroot0TimeoutsConfigIdGet) | **Get** /api/4.0/firewall/globalroot-0/timeouts/{configId} | dfwTimeoutConfigIdRead
[**FirewallGlobalroot0TimeoutsConfigIdPut**](FirewallApi.md#FirewallGlobalroot0TimeoutsConfigIdPut) | **Put** /api/4.0/firewall/globalroot-0/timeouts/{configId} | dfwTimeoutConfigIdUpdate
[**FirewallGlobalroot0TimeoutsGet**](FirewallApi.md#FirewallGlobalroot0TimeoutsGet) | **Get** /api/4.0/firewall/globalroot-0/timeouts | dfwTimeoutsList
[**FirewallGlobalroot0TimeoutsPost**](FirewallApi.md#FirewallGlobalroot0TimeoutsPost) | **Post** /api/4.0/firewall/globalroot-0/timeouts | dfwTimeoutCreate
[**FirewallStatsEventthresholdsGet**](FirewallApi.md#FirewallStatsEventthresholdsGet) | **Get** /api/4.0/firewall/stats/eventthresholds | dfwThresholdsRead
[**FirewallStatsEventthresholdsPut**](FirewallApi.md#FirewallStatsEventthresholdsPut) | **Put** /api/4.0/firewall/stats/eventthresholds | dfwThresholdsUpdate

# **FirewallConfigGlobalconfigurationGet**
> FirewallConfigGlobalconfigurationGet(ctx, )
dfwPerformanceRead

Retrieve performance configuration for distributed firewall.  Parameters:  

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

# **FirewallConfigGlobalconfigurationPut**
> FirewallConfigGlobalconfigurationPut(ctx, optional)
dfwPerformanceUpdate

Update the distributed firewall performance configuration.  **Method history:**  Release | Modification --------|------------- 6.2.3 | Method updated. **autoDraftDisabled** parameter added.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallApiFirewallConfigGlobalconfigurationPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallConfigGlobalconfigurationPutOpts struct
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

# **FirewallContextIdConfigIpfixDelete**
> FirewallContextIdConfigIpfixDelete(ctx, contextId)
dfwIPFixDelete

Deleting IPFIX configuration resets the config to default values   Parameters:  contextId: Specified context  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallContextIdConfigIpfixGet**
> FirewallContextIdConfigIpfixGet(ctx, contextId)
dfwIPFixRead

Query IPFIX configuration.  Parameters:  contextId: Specified context  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallContextIdConfigIpfixPut**
> FirewallContextIdConfigIpfixPut(ctx, contextId, optional)
dfwIPFixUpdate

Configure IPFIX.  Parameters:  contextId: Specified context  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***FirewallApiFirewallContextIdConfigIpfixPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallContextIdConfigIpfixPutOpts struct
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

# **FirewallDomainIDEnableTruefalsePost**
> FirewallDomainIDEnableTruefalsePost(ctx, domainID, truefalse)
dfwEnableDisableToggle

Enable or disable firewall components on a cluster  Parameters:  domainID: Specified cluster  truefalse: Set parameter to true/false to enable/disable  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainID** | **string**|  | 
  **truefalse** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallForceSyncIDPost**
> FirewallForceSyncIDPost(ctx, iD)
dfwSyncExecute

Force sync host or cluster.  Parameters:  ID: Specified host or cluster to synchronize  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iD** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigDelete**
> FirewallGlobalroot0ConfigDelete(ctx, )
dfwConfigDelete

Restores default configuration, which means one defaultLayer3 section with three default allow rules and one defaultLayer2Section with one default allow rule.   Parameters:  

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

# **FirewallGlobalroot0ConfigGet**
> FirewallGlobalroot0ConfigGet(ctx, optional)
dfwConfigShow

Retrieve distributed firewall rule configuration.  If no query parameters are used, all rule configuration is retrieved. Use the query parameters to filter the rule configuration information.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleType** | **optional.String**|  | 
 **source** | **optional.String**|  | 
 **destination** | **optional.String**|  | 
 **ruleId** | **optional.String**|  | 
 **comment** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **siProfile** | **optional.String**|  | 
 **edgeId** | **optional.String**|  | 
 **action** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer2sectionsGet**
> FirewallGlobalroot0ConfigLayer2sectionsGet(ctx, optional)
dfwL2SectionRead

Retrieve rules from the layer 2 section specified by section **name**.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer2sectionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer2sectionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer2sectionsPost**
> FirewallGlobalroot0ConfigLayer2sectionsPost(ctx, optional)
dfwL2SectionCreate

Create a layer 2 distributed firewall section.  By default, the section is created at the top of the firewall table. You can specify a location for the section with the **operation** and **anchorId** query parameters.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer2sectionsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer2sectionsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **operation** | **optional.**|  | 
 **anchorId** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer2sectionsSectionIdDelete**
> FirewallGlobalroot0ConfigLayer2sectionsSectionIdDelete(ctx, sectionId)
dfwL2SectionIdDelete

Delete the specified layer 2 section and its contents.  If the default layer 2 firewall section is selected, the request is rejected. See `GET /api/4.0/firewall/globalroot-0/defaultconfig` for information on resetting the default firewall section.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method updated. When deleting the default firewall rule section, the method previously removed all rules except for the default rule. The method now returns status 400 and the message `Cannot delete default section <sectionId>`.   Parameters:  sectionId: The ID of the section to modify.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer2sectionsSectionIdGet**
> FirewallGlobalroot0ConfigLayer2sectionsSectionIdGet(ctx, sectionId)
dfwL2SectionIdRead

Retrieve information about the specified layer 2 section.  Parameters:  sectionId: The ID of the section to modify.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer2sectionsSectionIdPost**
> FirewallGlobalroot0ConfigLayer2sectionsSectionIdPost(ctx, sectionId, optional)
dfwL2SectionIdAction

Move the specified layer 2 section.  Use the **action**, **operation**, and optionally **achorId** query parameters to specify the destination for the section.  `POST /api/4.0/firewall/globalroot-0/config/layer2sections/1009 ?action=revise&operation=insert_before&anchorId=1008`  `If-Match: 1478307787160`  ``` <section id=\"1009\" name=\"Test Section\" generationNumber=\"1478307787160\" timestamp=\"1478307787160\" type=\"LAYER2\">   ... </section> ```   Parameters:  sectionId: The ID of the section to modify.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **ifMatch** | **optional.**|  | 
 **contentType** | **optional.**|  | 
 **action** | **optional.**|  | 
 **operation** | **optional.**|  | 
 **anchorId** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer2sectionsSectionIdPut**
> FirewallGlobalroot0ConfigLayer2sectionsSectionIdPut(ctx, sectionId, optional)
dfwL2SectionIdUpdate

Update the specified layer 2 section in distributed firewall.  * Retrieve the configuration for the specified section. * Retrieve the Etag value from the response headers. * Extract and modify the configuration from the response body as needed. * Set the If-Match header to the Etag value, and submit the request.  Not all fields are required while sending the request. All the optional fields are safe to be ignored while sending the configuration to server. For example, if an IP set is referenced in the rule only IPSet and Type is needed in the Source/Destination objects and not Name and isValid tags.  When updating the firewall configuration: * IDs for new objects (rule/section) should be removed or set to zero. * If new entities (sections/rules) have been sent in the request, the response   will contain the system-generated IDs, which are assigned to these new   entities. * **appliedTo** can be any valid firewall rule element. * **action** can be *ALLOW*, *BLOCK*, or *REJECT*. REJECT sends reject message for   unaccepted packets; RST packets are sent for TCP connections and ICMP   unreachable code packets are sent for UDP, ICMP, and other IP connections * source and destination can have an exclude flag. For example, if you add an   exclude tag for 1.1.1.1 in the source parameter, the rule looks for traffic   originating from all IPs other than 1.1.1.1.  When Distributed Firewall is used with Service Composer, firewall sections created by Service Composer contain an additional attribute in the XML called managedBy. You should not modify Service Composer firewall sections using Distributed Firewall REST APIs. If you do, you must synchronize firewall rules from Service Composer using the `GET /api/2.0/services/policy/serviceprovider/firewall` API.   Parameters:  sectionId: The ID of the section to modify.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **ifMatch** | **optional.**|  | 
 **contentType** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesPost**
> FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesPost(ctx, sectionId, optional)
dfwL2RulesAdd

Add rules to the specified layer 2 section in distributed firewall.  You add firewall rules at the global scope. You can then narrow down the scope (datacenter, cluster, distributed virtual port group, network, virtual machine, vNIC, or logical switch) at which you want to apply the rule. Firewall allows you to add multiple objects at the source and destination levels for each rule, which helps reduce the total number of firewall rules to be added.  To add a identity based firewall rule, first create a security group based on Directory Group objects. Then create a firewall rule with the security group as the source or destination.  Rules that direct traffic to a third part service are referred to as layer3 redirect rules, and are displayed in the layer3 redirect tab.  When Distributed Firewall is used with Service Composer, firewall rules created by Service Composer contain an additional attribute in the XML called managedBy.  Follow this procedure to add a rule:  * Retrieve the configuration for the specified section. * Retrieve the Etag value from the response headers.   **Note**: Each section contains its own Etag, generationNumber, and   timestamp. When adding a new rule, you must use the Etag value of the   firewall section to which you wish to add the rule. * Extract and modify the configuration from the response body as needed. * Set the If-Match header to the section Etag value, and submit the request.  Not all fields are required while sending the request. All the optional fields are safe to be ignored while sending the configuration to server. For example, if an IP set is referenced in the rule only IPSet and Type is needed in the Source/Destination objects and not Name and isValid tags.  When updating the firewall configuration:  * IDs for new rules should be removed or set to zero. * If new rules have been sent in the request, the response   will contain the system-generated IDs, which are assigned to these new   entities. * **appliedTo** can be any valid firewall rule element. * **action** can be *ALLOW*, *BLOCK*, or *REJECT*. REJECT sends reject message for   unaccepted packets; RST packets are sent for TCP connections and ICMP   unreachable code packets are sent for UDP, ICMP, and other IP connections * source and destination can have an exclude flag. For example, if you add an   exclude tag for 1.1.1.1 in the source parameter, the rule looks for traffic   originating from all IPs other than 1.1.1.1.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **ifMatch** | **optional.**|  | 
 **contentType** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdDelete**
> FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdDelete(ctx, sectionId, ruleId, optional)
dfwL2RuleDelete

Delete the specified distributed firewall rule.  Parameters:  ruleId: The ID of the rule.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdGet**
> FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdGet(ctx, sectionId, ruleId)
dfwL2RuleRead

Retrieve the configuration of the specified rule.   Parameters:  ruleId: The ID of the rule.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdPut**
> FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdPut(ctx, sectionId, ruleId, optional)
dfwL2RuleUpdate

Update a distributed firewall rule in a layer 2 section.  * Retrieve the configuration for the section that contains the rule you want   to modify. * Retrieve the Etag value from the response headers.   **Note**: This is the Etag value of the firewall section to which you want   to add the rule. If you are keeping this rule in the same section, you must   keep the same Etag number. * Extract and modify the rule configuration from the response body as needed. * Set the If-Match header to the section Etag value, and submit the request.  Not all fields are required while sending the request. All the optional fields are safe to be ignored while sending the configuration to server. For example, if an IP set is referenced in the rule only IPSet and Type is needed in the Source/Destination objects and not Name and isValid tags.   Parameters:  ruleId: The ID of the rule.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **ifMatch** | **optional.**|  | 
 **contentType** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3redirectProfilesGet**
> FirewallGlobalroot0ConfigLayer3redirectProfilesGet(ctx, )
serviceInsertionProfilesRead

Retrieve the Service Insertion profiles that can be applied to layer3 redirect rules.   Parameters:  

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

# **FirewallGlobalroot0ConfigLayer3redirectsectionsPost**
> FirewallGlobalroot0ConfigLayer3redirectsectionsPost(ctx, optional)
layer3RedirectSectionsCreate

Add L3 redirect section  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsPostOpts struct
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

# **FirewallGlobalroot0ConfigLayer3redirectsectionsSectionDelete**
> FirewallGlobalroot0ConfigLayer3redirectsectionsSectionDelete(ctx, section)
sectionDelete

Delete specified L3 redirect section  Parameters:  section: Specify section by ID or name  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **section** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3redirectsectionsSectionGet**
> FirewallGlobalroot0ConfigLayer3redirectsectionsSectionGet(ctx, section)
sectionRead

Get L3 redirect section configuration  Parameters:  section: Specify section by ID or name  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **section** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3redirectsectionsSectionPut**
> FirewallGlobalroot0ConfigLayer3redirectsectionsSectionPut(ctx, section, optional)
sectionUpdate

Modify layer 3 redirect section. You will need to get the Etag value out of the GET first. Then pass the modified version of the whole redirect section configuration in the GET body.   Parameters:  section: Specify section by ID or name  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **section** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsSectionPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsSectionPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **ifMatch** | **optional.**|  | 
 **contentType** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesPost**
> FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesPost(ctx, section, optional)
rulesCreate

Add L3 redirect rule  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **section** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesPostOpts struct
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

# **FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDDelete**
> FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDDelete(ctx, section, ruleID)
ruleDelete

Delete specified L3 redirect rule  Parameters:  ruleID: Specified redirect rule  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **section** | **string**|  | 
  **ruleID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDGet**
> FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDGet(ctx, section, ruleID)
ruleRead

Get L3 redirect rule  Parameters:  ruleID: Specified redirect rule  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **section** | **string**|  | 
  **ruleID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDPut**
> FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDPut(ctx, section, ruleID, optional)
ruleUpdate

Modify L3 redirect rule. You will need Etag value from the response header of GET call. Then, pass Etag value as the if-match header in PUT call   Parameters:  ruleID: Specified redirect rule  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **section** | **string**|  | 
  **ruleID** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **ifMatch** | **optional.**|  | 
 **contentType** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3sectionsGet**
> FirewallGlobalroot0ConfigLayer3sectionsGet(ctx, optional)
dfwL3SectionRead

Retrieve rules from the layer 3 section specified by section **name**.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer3sectionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer3sectionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3sectionsPost**
> FirewallGlobalroot0ConfigLayer3sectionsPost(ctx, optional)
dfwL3SectionCreate

Create a layer 3 distributed firewall section.  By default, the section is created at the top of the firewall table. You can specify a location for the section with the **operation** and **anchorId** query parameters.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer3sectionsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer3sectionsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **contentType** | **optional.**|  | 
 **operation** | **optional.**|  | 
 **anchorId** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3sectionsSectionIdDelete**
> FirewallGlobalroot0ConfigLayer3sectionsSectionIdDelete(ctx, sectionId)
dfwL3SectionIdDelete

Delete the specified layer 3 distributed firewall section.  If the default layer 3 firewall section is selected, the request is rejected. See `GET /api/4.0/firewall/globalroot-0/defaultconfig` for information on resetting the default firewall section.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method updated. When deleting the default firewall rule section, the method previously removed all rules except for the default rule. The method now returns status 400 and the message `Cannot delete default section <sectionId>`.   Parameters:  sectionId: The ID of the section to modify.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3sectionsSectionIdGet**
> FirewallGlobalroot0ConfigLayer3sectionsSectionIdGet(ctx, sectionId)
dfwL3SectionIdRead

Retrieve information about the specified layer 3 section.  Parameters:  sectionId: The ID of the section to modify.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3sectionsSectionIdPost**
> FirewallGlobalroot0ConfigLayer3sectionsSectionIdPost(ctx, sectionId, optional)
dfwL3SectionIdAction

Move the specified layer 3 section.  Use the **action**, **operation**, and optionally **achorId** query parameters to specify the destination for the section.  `POST /api/4.0/firewall/globalroot-0/config/layer3sections/1007 ?action=revise&operation=insert_before&anchorId=1006`  `If-Match: 1477989118875`   ``` <section id=\"1007\" name=\"Web Section\" generationNumber=\"1477989118875\" timestamp=\"1477989118875\" type=\"LAYER3\">   ... </section> ```   Parameters:  sectionId: The ID of the section to modify.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **ifMatch** | **optional.**|  | 
 **contentType** | **optional.**|  | 
 **action** | **optional.**|  | 
 **operation** | **optional.**|  | 
 **anchorId** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3sectionsSectionIdPut**
> FirewallGlobalroot0ConfigLayer3sectionsSectionIdPut(ctx, sectionId, optional)
dfwL3SectionIdUpdate

Update the specified layer 3 section in distributed firewall.  * Retrieve the configuration for the specified section. * Retrieve the Etag value from the response headers. * Extract and modify the configuration from the response body as needed. * Set the If-Match header to the Etag value, and submit the request.  Not all fields are required while sending the request. All the optional fields are safe to be ignored while sending the configuration to server. For example, if an IP set is referenced in the rule only IPSet and Type is needed in the Source/Destination objects and not Name and isValid tags.  When updating the firewall configuration: * IDs for new objects (rule/section) should be removed or set to zero. * If new entities (sections/rules) have been sent in the request, the response   will contain the system-generated IDs, which are assigned to these new   entities. * **appliedTo** can be any valid firewall rule element. * **action** can be *ALLOW*, *BLOCK*, or *REJECT*. REJECT sends reject message for   unaccepted packets; RST packets are sent for TCP connections and ICMP   unreachable code packets are sent for UDP, ICMP, and other IP connections * source and destination can have an exclude flag. For example, if you add an   exclude tag for 1.1.1.1 in the source parameter, the rule looks for traffic   originating from all IPs other than 1.1.1.1.  When Distributed Firewall is used with Service Composer, firewall sections created by Service Composer contain an additional attribute in the XML called managedBy. You should not modify Service Composer firewall sections using Distributed Firewall REST APIs. If you do, you must synchronize firewall rules from Service Composer using the `GET /api/2.0/services/policy/serviceprovider/firewall` API.   Parameters:  sectionId: The ID of the section to modify.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **ifMatch** | **optional.**|  | 
 **contentType** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesPost**
> FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesPost(ctx, sectionId, optional)
dfwL3RulesAdd

Add rules to the specified layer 2 section in distributed firewall.  You add firewall rules at the global scope. You can then narrow down the scope (datacenter, cluster, distributed virtual port group, network, virtual machine, vNIC, or logical switch) at which you want to apply the rule. Firewall allows you to add multiple objects at the source and destination levels for each rule, which helps reduce the total number of firewall rules to be added.  To add a identity based firewall rule, first create a security group based on Directory Group objects. Then create a firewall rule with the security group as the source or destination.  Rules that direct traffic to a third part service are referred to as layer3 redirect rules, and are displayed in the layer3 redirect tab.  When Distributed Firewall is used with Service Composer, firewall rules created by Service Composer contain an additional attribute in the XML called managedBy.  Follow this procedure to add a rule:  * Retrieve the configuration for the specified section. * Retrieve the Etag value from the response headers.   **Note**: Each section contains its own Etag, generationNumber, and   timestamp. When adding a new rule, you must use the Etag value of the   firewall section to which you wish to add the rule. * Extract and modify the configuration from the response body as needed. * Set the If-Match header to the section Etag value, and submit the request.  Not all fields are required while sending the request. All the optional fields are safe to be ignored while sending the configuration to server. For example, if an IP set is referenced in the rule only IPSet and Type is needed in the Source/Destination objects and not Name and isValid tags.  When updating the firewall configuration:  * IDs for new rules should be removed or set to zero. * If new rules have been sent in the request, the response   will contain the system-generated IDs, which are assigned to these new   entities. * **appliedTo** can be any valid firewall rule element. * **action** can be *ALLOW*, *BLOCK*, or *REJECT*. REJECT sends reject message for   unaccepted packets; RST packets are sent for TCP connections and ICMP   unreachable code packets are sent for UDP, ICMP, and other IP connections * source and destination can have an exclude flag. For example, if you add an   exclude tag for 1.1.1.1 in the source parameter, the rule looks for traffic   originating from all IPs other than 1.1.1.1.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **ifMatch** | **optional.**|  | 
 **contentType** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdDelete**
> FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdDelete(ctx, sectionId, ruleId, optional)
dfwL3RuleDelete

Delete the specified distributed firewall rule.  Parameters:  ruleId: The ID of the rule beeing read, updated or deleted  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdGet**
> FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdGet(ctx, sectionId, ruleId)
dfwL3RuleRead

Retrieve information about the specified distributed firewall rule.   Parameters:  ruleId: The ID of the rule beeing read, updated or deleted  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdPut**
> FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdPut(ctx, sectionId, ruleId, optional)
dfwL3RuleUpdate

Update a distributed firewall rule in a layer 3 section.  * Retrieve the configuration for the section that contains the rule you want   to modify. * Retrieve the Etag value from the response headers.   **Note**: This is the Etag value of the firewall section to which you want   to add the rule. If you are keeping this rule in the same section, you must   keep the same Etag number. * Extract and modify the rule configuration from the response body as needed. * Set the If-Match header to the section Etag value, and submit the request.  Not all fields are required while sending the request. All the optional fields are safe to be ignored while sending the configuration to server. For example, if an IP set is referenced in the rule only IPSet and Type is needed in the Source/Destination objects and not Name and isValid tags.   Parameters:  ruleId: The ID of the rule beeing read, updated or deleted  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **ifMatch** | **optional.**|  | 
 **contentType** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0ConfigPut**
> FirewallGlobalroot0ConfigPut(ctx, optional)
dfwConfigUpdate

Update the complete firewall configuration in all sections.  * Retrieve the configuration with `GET /api/4.0/firewall/globalroot-0/config`. * Retrieve the Etag value from the response headers. * Extract and modify the configuration from the response body as needed. * Set the If-Match header to the Etag value, and submit the request.  Not all fields are required while sending the request. All the optional fields are safe to be ignored while sending the configuration to server. For example, if an IP set is referenced in the rule only IPSet and Type is needed in the Source/Destination objects and not Name and isValid tags.  When updating the firewall configuration: * IDs for new objects (rule/section) should be removed or set to zero. * If new entities (sections/rules) have been sent in the request, the response   will contain the system-generated IDs, which are assigned to these new   entities. * **appliedTo** can be any valid firewall rule element. * **action** can be *ALLOW*, *BLOCK*, or *REJECT*. REJECT sends reject message for   unaccepted packets; RST packets are sent for TCP connections and ICMP   unreachable code packets are sent for UDP, ICMP, and other IP connections * source and destination can have an exclude flag. For example, if you add an   exclude tag for 1.1.1.1 in the source parameter, the rule looks for traffic   originating from all IPs other than 1.1.1.1.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallApiFirewallGlobalroot0ConfigPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0ConfigPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **ifMatch** | **optional.**|  | 
 **contentType** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0DefaultconfigGet**
> FirewallGlobalroot0DefaultconfigGet(ctx, )
dfwConfigDefaultRead

Retrieve the default firewall configuration.  The output of this method can be used to restore the firewall config back to default. For example, to replace the layer 2 or layer 3 default section, use the relevant default section from the `GET /api/4.0/firewall/globalroot-0/defaultconfig` response body to create the request body of `PUT /api/4.0/firewall/globalroot-0/config/layer2sections|layer3sections/{sectionId}`.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  

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

# **FirewallGlobalroot0DraftsDraftIDActionExportGet**
> FirewallGlobalroot0DraftsDraftIDActionExportGet(ctx, draftID, optional)
dfwConfigExportRead

Export a configuration.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftID** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0DraftsDraftIDActionExportGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0DraftsDraftIDActionExportGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getLatestForUniversal** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0DraftsDraftIDActionImportPost**
> FirewallGlobalroot0DraftsDraftIDActionImportPost(ctx, draftID, optional)
dfwConfigImportCreate

Import a configuration.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftID** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0DraftsDraftIDActionImportPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0DraftsDraftIDActionImportPostOpts struct
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

# **FirewallGlobalroot0DraftsDraftIDDelete**
> FirewallGlobalroot0DraftsDraftIDDelete(ctx, draftID)
dfwDraftDelete

Delete a configuration.  Parameters:  draftID: Specified draft ID. Use `GET /4.0/firewall/globalroot-0/drafts` to retrieve all drafts.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0DraftsDraftIDGet**
> FirewallGlobalroot0DraftsDraftIDGet(ctx, draftID)
dfwDraftRead

Get a saved firewall configuration.  Parameters:  draftID: Specified draft ID. Use `GET /4.0/firewall/globalroot-0/drafts` to retrieve all drafts.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0DraftsDraftIDPut**
> FirewallGlobalroot0DraftsDraftIDPut(ctx, draftID, optional)
dfwDraftUpdate

Update a saved firewall configuration.  Parameters:  draftID: Specified draft ID. Use `GET /4.0/firewall/globalroot-0/drafts` to retrieve all drafts.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **draftID** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0DraftsDraftIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0DraftsDraftIDPutOpts struct
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

# **FirewallGlobalroot0DraftsGet**
> FirewallGlobalroot0DraftsGet(ctx, )
dfwSaveConfigRead

Displays the draft IDs of all saved configurations.  Parameters:  

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

# **FirewallGlobalroot0DraftsPost**
> FirewallGlobalroot0DraftsPost(ctx, optional)
dfwDraftsCreate

Save a firewall configuration.  Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallApiFirewallGlobalroot0DraftsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0DraftsPostOpts struct
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

# **FirewallGlobalroot0StateGet**
> FirewallGlobalroot0StateGet(ctx, )
dfwUpgradeRead

Retrieve current state of firewall functioning after NSX upgrade.   Parameters:  

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

# **FirewallGlobalroot0StatePut**
> FirewallGlobalroot0StatePut(ctx, )
dfwUpgradeEnable

Enable distributed firewall.  Parameters:  

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0StatusGet**
> FirewallGlobalroot0StatusGet(ctx, )
dfwStatusRead

Get firewall configuration status  **Method history:**  Release | Modification --------|------------- 6.2.4 | Method updated. Parameter **generationNumberObjects** added. Clusters not configured for firewall are excluded from the status output.   Parameters:  

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

# **FirewallGlobalroot0StatusLayer2sectionsSectionIDGet**
> FirewallGlobalroot0StatusLayer2sectionsSectionIDGet(ctx, sectionID)
L2SectionStatusRead

Retrieve status of the last publish action for the specified layer 2 section.  **Method history:**  Release | Modification --------|------------- 6.2.4 | Method updated. Parameter **generationNumberObjects** added. Clusters not configured for firewall are excluded from the status output.   Parameters:  sectionID: Section ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0StatusLayer3sectionsSectionIDGet**
> FirewallGlobalroot0StatusLayer3sectionsSectionIDGet(ctx, sectionID)
L3SectionStatusRead

Retrieve status of the last publish action for the specified layer 3 section.  **Method history:**  Release | Modification --------|------------- 6.2.4 | Method updated. Parameter **generationNumberObjects** added. Clusters not configured for firewall are excluded from the status output.   Parameters:  sectionID: Section ID  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionID** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0TimeoutsConfigIdDelete**
> FirewallGlobalroot0TimeoutsConfigIdDelete(ctx, configId)
dfwTimeoutConfigIdDelete

Delete the specified Distributed Firewall session timer configuration.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  configId: Session timer configuration ID (**firewallTimeoutConfiguration** id). For example, *1004*.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0TimeoutsConfigIdGet**
> FirewallGlobalroot0TimeoutsConfigIdGet(ctx, configId)
dfwTimeoutConfigIdRead

Retrieve the specified Distributed Firewall session timer configuration.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  configId: Session timer configuration ID (**firewallTimeoutConfiguration** id). For example, *1004*.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirewallGlobalroot0TimeoutsConfigIdPut**
> FirewallGlobalroot0TimeoutsConfigIdPut(ctx, configId, optional)
dfwTimeoutConfigIdUpdate

Update the specified Distributed Firewall session timer configuration.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  configId: Session timer configuration ID (**firewallTimeoutConfiguration** id). For example, *1004*.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | **string**|  | 
 **optional** | ***FirewallApiFirewallGlobalroot0TimeoutsConfigIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0TimeoutsConfigIdPutOpts struct
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

# **FirewallGlobalroot0TimeoutsGet**
> FirewallGlobalroot0TimeoutsGet(ctx, )
dfwTimeoutsList

Retrieve Distributed Firewall session timer configuration.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  

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

# **FirewallGlobalroot0TimeoutsPost**
> FirewallGlobalroot0TimeoutsPost(ctx, optional)
dfwTimeoutCreate

Create a Distributed Firewall session timer configuration.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallApiFirewallGlobalroot0TimeoutsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallGlobalroot0TimeoutsPostOpts struct
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

# **FirewallStatsEventthresholdsGet**
> FirewallStatsEventthresholdsGet(ctx, )
dfwThresholdsRead

Retrieve threshold configuration for distributed firewall.   Parameters:  

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

# **FirewallStatsEventthresholdsPut**
> FirewallStatsEventthresholdsPut(ctx, optional)
dfwThresholdsUpdate

Update threshold configuration for distributed firewall.   Parameters:  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirewallApiFirewallStatsEventthresholdsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FirewallApiFirewallStatsEventthresholdsPutOpts struct
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

