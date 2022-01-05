package nsxv

import (
	"context"
	"net/http"
)

// Compile-time interface checks
var _ IApiClient = &ApiClient{}

// IApiClient is the main interface for the library
type IApiClient interface {
	GetContext() context.Context
	Getcfg() Icfg
	Getcommon() Icommon
	GetAiApi() IAiApi
	GetAlarmsApi() IAlarmsApi
	GetApplianceManagementApi() IApplianceManagementApi
	GetApplicationApi() IApplicationApi
	GetApplicationgroupApi() IApplicationgroupApi
	GetAuditlogApi() IAuditlogApi
	GetCapacityApi() ICapacityApi
	GetCliApi() ICliApi
	GetConfigApi() IConfigApi
	GetControllerApi() IControllerApi
	GetDeployApi() IDeployApi
	GetDirectoryApi() IDirectoryApi
	GetEdgePublishApi() IEdgePublishApi
	GetEdgesApi() IEdgesApi
	GetEndpointsecurityApi() IEndpointsecurityApi
	GetEventcontrolApi() IEventcontrolApi
	GetExcludelistApi() IExcludelistApi
	GetFirewallApi() IFirewallApi
	GetFlowApi() IFlowApi
	GetHardwaregatewaysApi() IHardwaregatewaysApi
	GetIdentityApi() IIdentityApi
	GetInventoryApi() IInventoryApi
	GetIpsetApi() IIpsetApi
	GetJobApi() IJobApi
	GetMacsetApi() IMacsetApi
	GetNetworkFeaturesApi() INetworkFeaturesApi
	GetNwfabricApi() INwfabricApi
	GetPolicyApi() IPolicyApi
	GetPoolsApi() IPoolsApi
	GetScopesApi() IScopesApi
	GetSecuritygroupApi() ISecuritygroupApi
	GetSecuritytagsApi() ISecuritytagsApi
	GetSiApi() ISiApi
	GetSnmpApi() ISnmpApi
	GetSpoofguardApi() ISpoofguardApi
	GetSsoconfigApi() ISsoconfigApi
	GetSwitchesApi() ISwitchesApi
	GetSyslogApi() ISyslogApi
	GetSystemalarmsApi() ISystemalarmsApi
	GetSystemeventApi() ISystemeventApi
	GetTraceflowApi() ITraceflowApi
	GetTruststoreApi() ITruststoreApi
	GetUniversalsyncApi() IUniversalsyncApi
	GetUsermgmtApi() IUsermgmtApi
	GetVcconfigApi() IVcconfigApi
	GetVirtualwiresApi() IVirtualwiresApi
}

type ApiClient struct {
	APIClient
	Context context.Context
}

func (a *ApiClient) GetContext() context.Context {
	return a.Context
}

func (a *ApiClient) Getcfg() Icfg {
	return a.cfg
}

func (a *ApiClient) Getcommon() Icommon {
	return a.common
}

func (a *ApiClient) GetAiApi() IAiApi {
	return a.AiApi
}

func (a *ApiClient) GetAlarmsApi() IAlarmsApi {
	return a.AlarmsApi
}

func (a *ApiClient) GetApplianceManagementApi() IApplianceManagementApi {
	return a.ApplianceManagementApi
}

func (a *ApiClient) GetApplicationApi() IApplicationApi {
	return a.ApplicationApi
}

func (a *ApiClient) GetApplicationgroupApi() IApplicationgroupApi {
	return a.ApplicationgroupApi
}

func (a *ApiClient) GetAuditlogApi() IAuditlogApi {
	return a.AuditlogApi
}

func (a *ApiClient) GetCapacityApi() ICapacityApi {
	return a.CapacityApi
}

func (a *ApiClient) GetCliApi() ICliApi {
	return a.CliApi
}

func (a *ApiClient) GetConfigApi() IConfigApi {
	return a.ConfigApi
}

func (a *ApiClient) GetControllerApi() IControllerApi {
	return a.ControllerApi
}

func (a *ApiClient) GetDeployApi() IDeployApi {
	return a.DeployApi
}

func (a *ApiClient) GetDirectoryApi() IDirectoryApi {
	return a.DirectoryApi
}

func (a *ApiClient) GetEdgePublishApi() IEdgePublishApi {
	return a.EdgePublishApi
}

func (a *ApiClient) GetEdgesApi() IEdgesApi {
	return a.EdgesApi
}

func (a *ApiClient) GetEndpointsecurityApi() IEndpointsecurityApi {
	return a.EndpointsecurityApi
}

func (a *ApiClient) GetEventcontrolApi() IEventcontrolApi {
	return a.EventcontrolApi
}

func (a *ApiClient) GetExcludelistApi() IExcludelistApi {
	return a.ExcludelistApi
}

func (a *ApiClient) GetFirewallApi() IFirewallApi {
	return a.FirewallApi
}

func (a *ApiClient) GetFlowApi() IFlowApi {
	return a.FlowApi
}

func (a *ApiClient) GetHardwaregatewaysApi() IHardwaregatewaysApi {
	return a.HardwaregatewaysApi
}

func (a *ApiClient) GetIdentityApi() IIdentityApi {
	return a.IdentityApi
}

func (a *ApiClient) GetInventoryApi() IInventoryApi {
	return a.InventoryApi
}

func (a *ApiClient) GetIpsetApi() IIpsetApi {
	return a.IpsetApi
}

func (a *ApiClient) GetJobApi() IJobApi {
	return a.JobApi
}

func (a *ApiClient) GetMacsetApi() IMacsetApi {
	return a.MacsetApi
}

func (a *ApiClient) GetNetworkFeaturesApi() INetworkFeaturesApi {
	return a.NetworkFeaturesApi
}

func (a *ApiClient) GetNwfabricApi() INwfabricApi {
	return a.NwfabricApi
}

func (a *ApiClient) GetPolicyApi() IPolicyApi {
	return a.PolicyApi
}

func (a *ApiClient) GetPoolsApi() IPoolsApi {
	return a.PoolsApi
}

func (a *ApiClient) GetScopesApi() IScopesApi {
	return a.ScopesApi
}

func (a *ApiClient) GetSecuritygroupApi() ISecuritygroupApi {
	return a.SecuritygroupApi
}

func (a *ApiClient) GetSecuritytagsApi() ISecuritytagsApi {
	return a.SecuritytagsApi
}

func (a *ApiClient) GetSiApi() ISiApi {
	return a.SiApi
}

func (a *ApiClient) GetSnmpApi() ISnmpApi {
	return a.SnmpApi
}

func (a *ApiClient) GetSpoofguardApi() ISpoofguardApi {
	return a.SpoofguardApi
}

func (a *ApiClient) GetSsoconfigApi() ISsoconfigApi {
	return a.SsoconfigApi
}

func (a *ApiClient) GetSwitchesApi() ISwitchesApi {
	return a.SwitchesApi
}

func (a *ApiClient) GetSyslogApi() ISyslogApi {
	return a.SyslogApi
}

func (a *ApiClient) GetSystemalarmsApi() ISystemalarmsApi {
	return a.SystemalarmsApi
}

func (a *ApiClient) GetSystemeventApi() ISystemeventApi {
	return a.SystemeventApi
}

func (a *ApiClient) GetTraceflowApi() ITraceflowApi {
	return a.TraceflowApi
}

func (a *ApiClient) GetTruststoreApi() ITruststoreApi {
	return a.TruststoreApi
}

func (a *ApiClient) GetUniversalsyncApi() IUniversalsyncApi {
	return a.UniversalsyncApi
}

func (a *ApiClient) GetUsermgmtApi() IUsermgmtApi {
	return a.UsermgmtApi
}

func (a *ApiClient) GetVcconfigApi() IVcconfigApi {
	return a.VcconfigApi
}

func (a *ApiClient) GetVirtualwiresApi() IVirtualwiresApi {
	return a.VirtualwiresApi
}

type Icfg interface {
	AddDefaultHeader(string, string)
}

type Icommon interface {
}

type IAiApi interface {
	AiAppAppIDGet(context.Context, string) (*http.Response, error)
	AiAppGet(context.Context) (*http.Response, error)
	AiDesktoppoolDesktoppoolIDGet(context.Context, string) (*http.Response, error)
	AiDesktoppoolGet(context.Context) (*http.Response, error)
	AiDirectorygroupDirectorygroupIDGet(context.Context, string) (*http.Response, error)
	AiDirectorygroupGet(context.Context) (*http.Response, error)
	AiDirectorygroupUserUserIDGet(context.Context, string) (*http.Response, error)
	AiHostGet(context.Context) (*http.Response, error)
	AiHostHostIDGet(context.Context, string) (*http.Response, error)
	AiRecordsGet(context.Context, *AiApiAiRecordsGetOpts) (*http.Response, error)
	AiSecuritygroupGet(context.Context) (*http.Response, error)
	AiSecuritygroupSecgroupIDGet(context.Context, string) (*http.Response, error)
	AiUserUserIDGet(context.Context, string) (*http.Response, error)
	AiUserdetailsGet(context.Context, *AiApiAiUserdetailsGetOpts) (*http.Response, error)
	AiVmGet(context.Context) (*http.Response, error)
	AiVmVmIDGet(context.Context, string) (*http.Response, error)
}

type IAlarmsApi interface {
	ServicesAlarmsSourceIdGet(context.Context, string) (*http.Response, error)
	ServicesAlarmsSourceIdPost(context.Context, string, *AlarmsApiServicesAlarmsSourceIdPostOpts) (*http.Response, error)
}

type IApplianceManagementApi interface {
	ApplianceManagementBackuprestoreBackupPost(context.Context) (*http.Response, error)
	ApplianceManagementBackuprestoreBackupsGet(context.Context) (*http.Response, error)
	ApplianceManagementBackuprestoreBackupsettingsDelete(context.Context) (*http.Response, error)
	ApplianceManagementBackuprestoreBackupsettingsExcludedataPut(context.Context) (*http.Response, error)
	ApplianceManagementBackuprestoreBackupsettingsFtpsettingsPut(context.Context) (*http.Response, error)
	ApplianceManagementBackuprestoreBackupsettingsGet(context.Context) (*http.Response, error)
	ApplianceManagementBackuprestoreBackupsettingsPut(context.Context, *ApplianceManagementApiApplianceManagementBackuprestoreBackupsettingsPutOpts) (*http.Response, error)
	ApplianceManagementBackuprestoreBackupsettingsScheduleDelete(context.Context) (*http.Response, error)
	ApplianceManagementBackuprestoreBackupsettingsSchedulePut(context.Context) (*http.Response, error)
	ApplianceManagementBackuprestoreRestorePost(context.Context, *ApplianceManagementApiApplianceManagementBackuprestoreRestorePostOpts) (*http.Response, error)
	ApplianceManagementCertificatemanagerCertificatesNsxGet(context.Context) (*http.Response, error)
	ApplianceManagementCertificatemanagerCsrNsxGet(context.Context) (*http.Response, error)
	ApplianceManagementCertificatemanagerCsrNsxPost(context.Context, *ApplianceManagementApiApplianceManagementCertificatemanagerCsrNsxPostOpts) (*http.Response, error)
	ApplianceManagementCertificatemanagerPkcs12keystoreNsxPost(context.Context, *ApplianceManagementApiApplianceManagementCertificatemanagerPkcs12keystoreNsxPostOpts) (*http.Response, error)
	ApplianceManagementCertificatemanagerUploadchainNsxPost(context.Context) (*http.Response, error)
	ApplianceManagementComponentsComponentAPPMGMTRestartPost(context.Context) (*http.Response, error)
	ApplianceManagementComponentsComponentComponentIDDependenciesGet(context.Context, string) (*http.Response, error)
	ApplianceManagementComponentsComponentComponentIDDependentsGet(context.Context, string) (*http.Response, error)
	ApplianceManagementComponentsComponentComponentIDGet(context.Context, string) (*http.Response, error)
	ApplianceManagementComponentsComponentComponentIDStatusGet(context.Context, string) (*http.Response, error)
	ApplianceManagementComponentsComponentComponentIDToggleStatusCommandPost(context.Context, string, string) (*http.Response, error)
	ApplianceManagementComponentsGet(context.Context) (*http.Response, error)
	ApplianceManagementGlobalInfoGet(context.Context) (*http.Response, error)
	ApplianceManagementNotificationsDelete(context.Context) (*http.Response, error)
	ApplianceManagementNotificationsGet(context.Context) (*http.Response, error)
	ApplianceManagementNotificationsIDAcknowledgePost(context.Context, string) (*http.Response, error)
	ApplianceManagementSummaryComponentsGet(context.Context) (*http.Response, error)
	ApplianceManagementSummarySystemGet(context.Context) (*http.Response, error)
	ApplianceManagementSystemCpuinfoGet(context.Context) (*http.Response, error)
	ApplianceManagementSystemLocaleGet(context.Context) (*http.Response, error)
	ApplianceManagementSystemLocalePut(context.Context, *ApplianceManagementApiApplianceManagementSystemLocalePutOpts) (*http.Response, error)
	ApplianceManagementSystemMeminfoGet(context.Context) (*http.Response, error)
	ApplianceManagementSystemNetworkDnsDelete(context.Context) (*http.Response, error)
	ApplianceManagementSystemNetworkDnsPut(context.Context, *ApplianceManagementApiApplianceManagementSystemNetworkDnsPutOpts) (*http.Response, error)
	ApplianceManagementSystemNetworkGet(context.Context) (*http.Response, error)
	ApplianceManagementSystemNetworkPut(context.Context, *ApplianceManagementApiApplianceManagementSystemNetworkPutOpts) (*http.Response, error)
	ApplianceManagementSystemRestartPost(context.Context) (*http.Response, error)
	ApplianceManagementSystemSecuritysettingsGet(context.Context) (*http.Response, error)
	ApplianceManagementSystemSecuritysettingsPost(context.Context, *ApplianceManagementApiApplianceManagementSystemSecuritysettingsPostOpts) (*http.Response, error)
	ApplianceManagementSystemStorageinfoGet(context.Context) (*http.Response, error)
	ApplianceManagementSystemSyslogserverDelete(context.Context) (*http.Response, error)
	ApplianceManagementSystemSyslogserverGet(context.Context) (*http.Response, error)
	ApplianceManagementSystemSyslogserverPut(context.Context, *ApplianceManagementApiApplianceManagementSystemSyslogserverPutOpts) (*http.Response, error)
	ApplianceManagementSystemTimesettingsGet(context.Context) (*http.Response, error)
	ApplianceManagementSystemTimesettingsNtpDelete(context.Context) (*http.Response, error)
	ApplianceManagementSystemTimesettingsPut(context.Context, *ApplianceManagementApiApplianceManagementSystemTimesettingsPutOpts) (*http.Response, error)
	ApplianceManagementSystemTlssettingsGet(context.Context) (*http.Response, error)
	ApplianceManagementSystemTlssettingsPost(context.Context, *ApplianceManagementApiApplianceManagementSystemTlssettingsPostOpts) (*http.Response, error)
	ApplianceManagementSystemUptimeGet(context.Context) (*http.Response, error)
	ApplianceManagementTechsupportlogsComponentIDPost(context.Context, string) (*http.Response, error)
	ApplianceManagementTechsupportlogsFilenameGet(context.Context, string) (*http.Response, error)
	ApplianceManagementUpgradeInformationComponentIDGet(context.Context, string) (*http.Response, error)
	ApplianceManagementUpgradeStartComponentIDPost(context.Context, string, *ApplianceManagementApiApplianceManagementUpgradeStartComponentIDPostOpts) (*http.Response, error)
	ApplianceManagementUpgradeStatusComponentIDGet(context.Context, string) (*http.Response, error)
	ApplianceManagementUpgradeUploadbundleComponentIDPost(context.Context, string) (*http.Response, error)
}

type IApplicationApi interface {
	ServicesApplicationApplicationIdDelete(context.Context, string, *ApplicationApiServicesApplicationApplicationIdDeleteOpts) (*http.Response, error)
	ServicesApplicationApplicationIdGet(context.Context, string) (*http.Response, error)
	ServicesApplicationApplicationIdPut(context.Context, string, *ApplicationApiServicesApplicationApplicationIdPutOpts) (*http.Response, error)
	ServicesApplicationScopeIdPost(context.Context, string, *ApplicationApiServicesApplicationScopeIdPostOpts) (*http.Response, error)
	ServicesApplicationScopeScopeIdGet(context.Context, string) (*http.Response, error)
}

type IApplicationgroupApi interface {
	ServicesApplicationgroupApplicationgroupIdDelete(context.Context, string, *ApplicationgroupApiServicesApplicationgroupApplicationgroupIdDeleteOpts) (*http.Response, error)
	ServicesApplicationgroupApplicationgroupIdGet(context.Context, string) (*http.Response, error)
	ServicesApplicationgroupApplicationgroupIdMembersMorefDelete(context.Context, string, string) (*http.Response, error)
	ServicesApplicationgroupApplicationgroupIdMembersMorefPut(context.Context, string, string) (*http.Response, error)
	ServicesApplicationgroupApplicationgroupIdPut(context.Context, string, *ApplicationgroupApiServicesApplicationgroupApplicationgroupIdPutOpts) (*http.Response, error)
	ServicesApplicationgroupScopeScopeIdGet(context.Context, string) (*http.Response, error)
	ServicesApplicationgroupScopeScopeIdMembersGet(context.Context, string) (*http.Response, error)
	ServicesApplicationgroupScopeScopeIdPost(context.Context, string, *ApplicationgroupApiServicesApplicationgroupScopeScopeIdPostOpts) (*http.Response, error)
}

type IAuditlogApi interface {
	AuditlogGet(context.Context, *AuditlogApiAuditlogGetOpts) (*http.Response, error)
}

type ICapacityApi interface {
	ServicesLicensingCapacityusageGet(context.Context) (*http.Response, error)
}

type ICliApi interface {
	NsxCliPost(context.Context, *CliApiNsxCliPostOpts) (*http.Response, error)
}

type IConfigApi interface {
	VdnConfigHostHostIdVxlanVtepsPost(context.Context, string, *ConfigApiVdnConfigHostHostIdVxlanVtepsPostOpts) (*http.Response, error)
	VdnConfigMulticastsGet(context.Context) (*http.Response, error)
	VdnConfigMulticastsMulticastAddresssRangeIdDelete(context.Context, string) (*http.Response, error)
	VdnConfigMulticastsMulticastAddresssRangeIdGet(context.Context, string) (*http.Response, error)
	VdnConfigMulticastsMulticastAddresssRangeIdPut(context.Context, string, *ConfigApiVdnConfigMulticastsMulticastAddresssRangeIdPutOpts) (*http.Response, error)
	VdnConfigMulticastsPost(context.Context, *ConfigApiVdnConfigMulticastsPostOpts) (*http.Response, error)
	VdnConfigResourcesAllocatedGet(context.Context, *ConfigApiVdnConfigResourcesAllocatedGetOpts) (*http.Response, error)
	VdnConfigSegmentsGet(context.Context) (*http.Response, error)
	VdnConfigSegmentsPost(context.Context, *ConfigApiVdnConfigSegmentsPostOpts) (*http.Response, error)
	VdnConfigSegmentsSegmentPoolIdDelete(context.Context, string) (*http.Response, error)
	VdnConfigSegmentsSegmentPoolIdGet(context.Context, string) (*http.Response, error)
	VdnConfigSegmentsSegmentPoolIdPut(context.Context, string, *ConfigApiVdnConfigSegmentsSegmentPoolIdPutOpts) (*http.Response, error)
	VdnConfigVxlanUdpPortGet(context.Context) (*http.Response, error)
	VdnConfigVxlanUdpPortPortNumberPut(context.Context, string, *ConfigApiVdnConfigVxlanUdpPortPortNumberPutOpts) (*http.Response, error)
	VdnConfigVxlanUdpPortTaskStatusGet(context.Context) (*http.Response, error)
}

type IControllerApi interface {
	VdnControllerClusterGet(context.Context) (*http.Response, error)
	VdnControllerClusterPut(context.Context, *ControllerApiVdnControllerClusterPutOpts) (*http.Response, error)
	VdnControllerControllerIdDelete(context.Context, string, *ControllerApiVdnControllerControllerIdDeleteOpts) (*http.Response, error)
	VdnControllerControllerIdPost(context.Context, string, *ControllerApiVdnControllerControllerIdPostOpts) (*http.Response, error)
	VdnControllerControllerIdSnapshotGet(context.Context, string) (*http.Response, error)
	VdnControllerControllerIdSyslogDelete(context.Context, string) (*http.Response, error)
	VdnControllerControllerIdSyslogGet(context.Context, string) (*http.Response, error)
	VdnControllerControllerIdSyslogPost(context.Context, string, *ControllerApiVdnControllerControllerIdSyslogPostOpts) (*http.Response, error)
	VdnControllerControllerIdSystemStatsGet(context.Context, string) (*http.Response, error)
	VdnControllerControllerIdTechsupportlogsGet(context.Context, string) (*http.Response, error)
	VdnControllerCredentialPut(context.Context, *ControllerApiVdnControllerCredentialPutOpts) (*http.Response, error)
	VdnControllerGet(context.Context) (*http.Response, error)
	VdnControllerPost(context.Context, *ControllerApiVdnControllerPostOpts) (*http.Response, error)
	VdnControllerProgressJobIdGet(context.Context, string) (*http.Response, error)
	VdnControllerUpgradeAvailableGet(context.Context) (*http.Response, error)
}

type IDeployApi interface {
	SiDeployClusterClusterIDDelete(context.Context, string, *DeployApiSiDeployClusterClusterIDDeleteOpts) (*http.Response, error)
	SiDeployClusterClusterIDGet(context.Context, string) (*http.Response, error)
	SiDeployClusterClusterIDServiceServiceIDGet(context.Context, string, string) (*http.Response, error)
	SiDeployPost(context.Context, *DeployApiSiDeployPostOpts) (*http.Response, error)
	SiDeployPut(context.Context, *DeployApiSiDeployPutOpts) (*http.Response, error)
	SiDeployServiceServiceIDDelete(context.Context, string, *DeployApiSiDeployServiceServiceIDDeleteOpts) (*http.Response, error)
	SiDeployServiceServiceIDDependsOnGet(context.Context, string) (*http.Response, error)
	SiDeployServiceServiceIDGet(context.Context, string) (*http.Response, error)
}

type IDirectoryApi interface {
	DirectoryDeleteDomainIDDelete(context.Context, string) (*http.Response, error)
	DirectoryDeleteEventLogServerServerIDDelete(context.Context, string) (*http.Response, error)
	DirectoryDeleteLdapServerServerIDDelete(context.Context, string) (*http.Response, error)
	DirectoryDeltaSyncDomainIDPut(context.Context, string) (*http.Response, error)
	DirectoryFullSyncDomainIDPut(context.Context, string) (*http.Response, error)
	DirectoryListDomainsGet(context.Context) (*http.Response, error)
	DirectoryListEventLogServersForDomainDomainIDGet(context.Context, string) (*http.Response, error)
	DirectoryListLdapServersForDomainDomainIDGet(context.Context, string) (*http.Response, error)
	DirectoryUpdateDomainPost(context.Context, *DirectoryApiDirectoryUpdateDomainPostOpts) (*http.Response, error)
	DirectoryUpdateEventLogServerPost(context.Context, *DirectoryApiDirectoryUpdateEventLogServerPostOpts) (*http.Response, error)
	DirectoryUpdateLdapServerPost(context.Context, *DirectoryApiDirectoryUpdateLdapServerPostOpts) (*http.Response, error)
}

type IEdgePublishApi interface {
	EdgePublishTuningConfigurationGet(context.Context) (*http.Response, error)
	EdgePublishTuningConfigurationPut(context.Context, *EdgePublishApiEdgePublishTuningConfigurationPutOpts) (*http.Response, error)
}

type IEdgesApi interface {
	EdgesEdgeIdAesniPost(context.Context, string, *EdgesApiEdgesEdgeIdAesniPostOpts) (*http.Response, error)
	EdgesEdgeIdAppliancesGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdAppliancesHaIndexDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdAppliancesHaIndexGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdAppliancesHaIndexPost(context.Context, string, string, *EdgesApiEdgesEdgeIdAppliancesHaIndexPostOpts) (*http.Response, error)
	EdgesEdgeIdAppliancesHaIndexPut(context.Context, string, string, *EdgesApiEdgesEdgeIdAppliancesHaIndexPutOpts) (*http.Response, error)
	EdgesEdgeIdAppliancesPost(context.Context, string, *EdgesApiEdgesEdgeIdAppliancesPostOpts) (*http.Response, error)
	EdgesEdgeIdAppliancesPut(context.Context, string) (*http.Response, error)
	EdgesEdgeIdAutoconfigurationGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdAutoconfigurationPut(context.Context, string, *EdgesApiEdgesEdgeIdAutoconfigurationPutOpts) (*http.Response, error)
	EdgesEdgeIdBridgingConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdBridgingConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdBridgingConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdBridgingConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdCliremoteaccessPost(context.Context, string, *EdgesApiEdgesEdgeIdCliremoteaccessPostOpts) (*http.Response, error)
	EdgesEdgeIdClisettingsPut(context.Context, string, *EdgesApiEdgesEdgeIdClisettingsPutOpts) (*http.Response, error)
	EdgesEdgeIdCoredumpPost(context.Context, string) (*http.Response, error)
	EdgesEdgeIdDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdDhcpConfigBindingsBindingIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdDhcpConfigBindingsPost(context.Context, string, *EdgesApiEdgesEdgeIdDhcpConfigBindingsPostOpts) (*http.Response, error)
	EdgesEdgeIdDhcpConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdDhcpConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdDhcpConfigIppoolsPoolIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdDhcpConfigIppoolsPost(context.Context, string, *EdgesApiEdgesEdgeIdDhcpConfigIppoolsPostOpts) (*http.Response, error)
	EdgesEdgeIdDhcpConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdDhcpConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdDhcpConfigRelayDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdDhcpConfigRelayGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdDhcpConfigRelayPut(context.Context, string, *EdgesApiEdgesEdgeIdDhcpConfigRelayPutOpts) (*http.Response, error)
	EdgesEdgeIdDhcpLeaseInfoGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdDnsConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdDnsConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdDnsConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdDnsConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdDnsStatisticsGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdDnsclientPut(context.Context, string, *EdgesApiEdgesEdgeIdDnsclientPutOpts) (*http.Response, error)
	EdgesEdgeIdFipsPost(context.Context, string, *EdgesApiEdgesEdgeIdFipsPostOpts) (*http.Response, error)
	EdgesEdgeIdFirewallConfigDefaultpolicyGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdFirewallConfigDefaultpolicyPut(context.Context, string, *EdgesApiEdgesEdgeIdFirewallConfigDefaultpolicyPutOpts) (*http.Response, error)
	EdgesEdgeIdFirewallConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdFirewallConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdFirewallConfigGlobalGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdFirewallConfigGlobalPut(context.Context, string, *EdgesApiEdgesEdgeIdFirewallConfigGlobalPutOpts) (*http.Response, error)
	EdgesEdgeIdFirewallConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdFirewallConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdFirewallConfigRulesPost(context.Context, string, *EdgesApiEdgesEdgeIdFirewallConfigRulesPostOpts) (*http.Response, error)
	EdgesEdgeIdFirewallConfigRulesRuleIdDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdFirewallConfigRulesRuleIdGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdFirewallConfigRulesRuleIdPut(context.Context, string, string, *EdgesApiEdgesEdgeIdFirewallConfigRulesRuleIdPutOpts) (*http.Response, error)
	EdgesEdgeIdFirewallStatisticsFirewallGet(context.Context, string, *EdgesApiEdgesEdgeIdFirewallStatisticsFirewallGetOpts) (*http.Response, error)
	EdgesEdgeIdFirewallStatisticsRuleIdGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdGet(context.Context, string, *EdgesApiEdgesEdgeIdGetOpts) (*http.Response, error)
	EdgesEdgeIdHighavailabilityConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdHighavailabilityConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdHighavailabilityConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdHighavailabilityConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdInterfacesDelete(context.Context, string, *EdgesApiEdgesEdgeIdInterfacesDeleteOpts) (*http.Response, error)
	EdgesEdgeIdInterfacesGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdInterfacesIndexDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdInterfacesIndexGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdInterfacesIndexPut(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdInterfacesPost(context.Context, string, *EdgesApiEdgesEdgeIdInterfacesPostOpts) (*http.Response, error)
	EdgesEdgeIdIpsecConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdIpsecConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdIpsecConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdIpsecConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdIpsecStatisticsGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdL2vpnConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdL2vpnConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdL2vpnConfigPost(context.Context, string, *EdgesApiEdgesEdgeIdL2vpnConfigPostOpts) (*http.Response, error)
	EdgesEdgeIdL2vpnConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdL2vpnConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdL2vpnConfigStatisticsGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerAccelerationPost(context.Context, string, *EdgesApiEdgesEdgeIdLoadbalancerAccelerationPostOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationprofilesAppProfileIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationprofilesAppProfileIDGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationprofilesAppProfileIDPut(context.Context, string, string, *EdgesApiEdgesEdgeIdLoadbalancerConfigApplicationprofilesAppProfileIDPutOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationprofilesDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationprofilesGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationprofilesPost(context.Context, string, *EdgesApiEdgesEdgeIdLoadbalancerConfigApplicationprofilesPostOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationrulesAppruleIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationrulesAppruleIDGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationrulesAppruleIDPut(context.Context, string, string, *EdgesApiEdgesEdgeIdLoadbalancerConfigApplicationrulesAppruleIDPutOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationrulesDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationrulesGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigApplicationrulesPost(context.Context, string, *EdgesApiEdgesEdgeIdLoadbalancerConfigApplicationrulesPostOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigMembersMemberIDPost(context.Context, string, string, *EdgesApiEdgesEdgeIdLoadbalancerConfigMembersMemberIDPostOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigMonitorsDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigMonitorsGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigMonitorsMonitorIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigMonitorsMonitorIDGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigMonitorsMonitorIDPut(context.Context, string, string, *EdgesApiEdgesEdgeIdLoadbalancerConfigMonitorsMonitorIDPutOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigMonitorsPost(context.Context, string, *EdgesApiEdgesEdgeIdLoadbalancerConfigMonitorsPostOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigPoolsDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigPoolsGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigPoolsPoolIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigPoolsPoolIDGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigPoolsPoolIDPut(context.Context, string, string, *EdgesApiEdgesEdgeIdLoadbalancerConfigPoolsPoolIDPutOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigPoolsPost(context.Context, string, *EdgesApiEdgesEdgeIdLoadbalancerConfigPoolsPostOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdLoadbalancerConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigVirtualserversDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigVirtualserversGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigVirtualserversPost(context.Context, string, *EdgesApiEdgesEdgeIdLoadbalancerConfigVirtualserversPostOpts) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigVirtualserversVirtualserverIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerConfigVirtualserversVirtualserverIDGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdLoadbalancerStatisticsGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdLoggingPost(context.Context, string, *EdgesApiEdgesEdgeIdLoggingPostOpts) (*http.Response, error)
	EdgesEdgeIdMgmtinterfaceGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdMgmtinterfacePut(context.Context, string, *EdgesApiEdgesEdgeIdMgmtinterfacePutOpts) (*http.Response, error)
	EdgesEdgeIdNatConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdNatConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdNatConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdNatConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdNatConfigRulesPost(context.Context, string, *EdgesApiEdgesEdgeIdNatConfigRulesPostOpts) (*http.Response, error)
	EdgesEdgeIdNatConfigRulesRuleIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdNatConfigRulesRuleIDPut(context.Context, string, string, *EdgesApiEdgesEdgeIdNatConfigRulesRuleIDPutOpts) (*http.Response, error)
	EdgesEdgeIdPost(context.Context, string, *EdgesApiEdgesEdgeIdPostOpts) (*http.Response, error)
	EdgesEdgeIdPut(context.Context, string, *EdgesApiEdgesEdgeIdPutOpts) (*http.Response, error)
	EdgesEdgeIdRoutingConfigBgpDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdRoutingConfigBgpGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdRoutingConfigBgpPut(context.Context, string, *EdgesApiEdgesEdgeIdRoutingConfigBgpPutOpts) (*http.Response, error)
	EdgesEdgeIdRoutingConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdRoutingConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdRoutingConfigGlobalGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdRoutingConfigGlobalPut(context.Context, string, *EdgesApiEdgesEdgeIdRoutingConfigGlobalPutOpts) (*http.Response, error)
	EdgesEdgeIdRoutingConfigOspfDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdRoutingConfigOspfGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdRoutingConfigOspfPut(context.Context, string, *EdgesApiEdgesEdgeIdRoutingConfigOspfPutOpts) (*http.Response, error)
	EdgesEdgeIdRoutingConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdRoutingConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdRoutingConfigStaticDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdRoutingConfigStaticGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdRoutingConfigStaticPut(context.Context, string, *EdgesApiEdgesEdgeIdRoutingConfigStaticPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnActivesessionsGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnActivesessionsSessionIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnAuthLocalusersUsersPut(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnAuthLocalusersUsersPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigAdvancedconfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigAdvancedconfigPut(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigAdvancedconfigPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigAuthLocalserverUsersDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigAuthLocalserverUsersPost(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigAuthLocalserverUsersPostOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigAuthLocalserverUsersPut(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigAuthLocalserverUsersPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigAuthLocalserverUsersUserIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigAuthLocalserverUsersUserIDGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigAuthSettingsGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigAuthSettingsPut(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigAuthSettingsPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigAuthSettingsRsaconfigfilePost(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionClientconfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionClientconfigPut(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigClientNetworkextensionClientconfigPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionInstallpackagesDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionInstallpackagesGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionInstallpackagesPackageIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionInstallpackagesPackageIDGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionInstallpackagesPackageIDPut(context.Context, string, string, *EdgesApiEdgesEdgeIdSslvpnConfigClientNetworkextensionInstallpackagesPackageIDPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionInstallpackagesPost(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigClientNetworkextensionInstallpackagesPostOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionInstallpackagesPut(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionIppoolsDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionIppoolsGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionIppoolsIppoolIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionIppoolsIppoolIDGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionIppoolsIppoolIDPut(context.Context, string, string, *EdgesApiEdgesEdgeIdSslvpnConfigClientNetworkextensionIppoolsIppoolIDPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionIppoolsPost(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigClientNetworkextensionIppoolsPostOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionIppoolsPut(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigClientNetworkextensionIppoolsPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionPrivatenetworksDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionPrivatenetworksGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionPrivatenetworksNetworkIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionPrivatenetworksNetworkIDGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionPrivatenetworksNetworkIDPut(context.Context, string, string, *EdgesApiEdgesEdgeIdSslvpnConfigClientNetworkextensionPrivatenetworksNetworkIDPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionPrivatenetworksPost(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigClientNetworkextensionPrivatenetworksPostOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigClientNetworkextensionPrivatenetworksPut(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigLayoutGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigLayoutImagesImageTypePost(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigLayoutPut(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigLayoutPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigPost(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigPostOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigPut(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigScriptDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigScriptFileIDDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigScriptFileIDGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigScriptFileIDPut(context.Context, string, string, *EdgesApiEdgesEdgeIdSslvpnConfigScriptFileIDPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigScriptFilePost(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigScriptGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigScriptPost(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigScriptPostOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigScriptPut(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigScriptPutOpts) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigServerGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSslvpnConfigServerPut(context.Context, string, *EdgesApiEdgesEdgeIdSslvpnConfigServerPutOpts) (*http.Response, error)
	EdgesEdgeIdStatisticsDashboardInterfaceGet(context.Context, string, *EdgesApiEdgesEdgeIdStatisticsDashboardInterfaceGetOpts) (*http.Response, error)
	EdgesEdgeIdStatisticsDashboardIpsecGet(context.Context, string, *EdgesApiEdgesEdgeIdStatisticsDashboardIpsecGetOpts) (*http.Response, error)
	EdgesEdgeIdStatisticsDashboardSslvpnGet(context.Context, string, *EdgesApiEdgesEdgeIdStatisticsDashboardSslvpnGetOpts) (*http.Response, error)
	EdgesEdgeIdStatisticsInterfacesGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdStatisticsInterfacesInternalGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdStatisticsInterfacesUplinkGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdStatusGet(context.Context, string, *EdgesApiEdgesEdgeIdStatusGetOpts) (*http.Response, error)
	EdgesEdgeIdSummaryGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSyslogConfigDelete(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSyslogConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSyslogConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdSyslogConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdSystemcontrolConfigDelete(context.Context, string, *EdgesApiEdgesEdgeIdSystemcontrolConfigDeleteOpts) (*http.Response, error)
	EdgesEdgeIdSystemcontrolConfigGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdSystemcontrolConfigPut(context.Context, string, *EdgesApiEdgesEdgeIdSystemcontrolConfigPutOpts) (*http.Response, error)
	EdgesEdgeIdTechsupportlogsGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdVnicsGet(context.Context, string) (*http.Response, error)
	EdgesEdgeIdVnicsIndexDelete(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdVnicsIndexGet(context.Context, string, string) (*http.Response, error)
	EdgesEdgeIdVnicsIndexPut(context.Context, string, string, *EdgesApiEdgesEdgeIdVnicsIndexPutOpts) (*http.Response, error)
	EdgesEdgeIdVnicsPost(context.Context, string, *EdgesApiEdgesEdgeIdVnicsPostOpts) (*http.Response, error)
	EdgesGet(context.Context, *EdgesApiEdgesGetOpts) (*http.Response, error)
	EdgesJobsGet(context.Context, *EdgesApiEdgesJobsGetOpts) (*http.Response, error)
	EdgesJobsJobIdGet(context.Context, string) (*http.Response, error)
	EdgesPost(context.Context, *EdgesApiEdgesPostOpts) (*http.Response, error)
}

type IEndpointsecurityApi interface {
	EndpointsecurityActivationGet(context.Context, *EndpointsecurityApiEndpointsecurityActivationGetOpts) (*http.Response, error)
	EndpointsecurityActivationVendorIDAltitudeMoidDelete(context.Context, string, string, string) (*http.Response, error)
	EndpointsecurityActivationVendorIDAltitudeMoidGet(context.Context, string, string, string) (*http.Response, error)
	EndpointsecurityActivationVendorIDAltitudePost(context.Context, string, string, *EndpointsecurityApiEndpointsecurityActivationVendorIDAltitudePostOpts) (*http.Response, error)
	EndpointsecurityActivationVendorIDSolutionIDGet(context.Context, string, string) (*http.Response, error)
	EndpointsecurityRegistrationPost(context.Context, *EndpointsecurityApiEndpointsecurityRegistrationPostOpts) (*http.Response, error)
	EndpointsecurityRegistrationVendorIDAltitudeDelete(context.Context, string, string) (*http.Response, error)
	EndpointsecurityRegistrationVendorIDAltitudeGet(context.Context, string, string) (*http.Response, error)
	EndpointsecurityRegistrationVendorIDAltitudeLocationDelete(context.Context, string, string) (*http.Response, error)
	EndpointsecurityRegistrationVendorIDAltitudeLocationGet(context.Context, string, string) (*http.Response, error)
	EndpointsecurityRegistrationVendorIDAltitudeLocationPost(context.Context, string, string, *EndpointsecurityApiEndpointsecurityRegistrationVendorIDAltitudeLocationPostOpts) (*http.Response, error)
	EndpointsecurityRegistrationVendorIDDelete(context.Context, string) (*http.Response, error)
	EndpointsecurityRegistrationVendorIDGet(context.Context, string) (*http.Response, error)
	EndpointsecurityRegistrationVendorIDPost(context.Context, string, *EndpointsecurityApiEndpointsecurityRegistrationVendorIDPostOpts) (*http.Response, error)
	EndpointsecurityRegistrationVendorIDSolutionsGet(context.Context, string) (*http.Response, error)
	EndpointsecurityRegistrationVendorsGet(context.Context) (*http.Response, error)
}

type IEventcontrolApi interface {
	EventcontrolConfigVmVmIDGet(context.Context, string) (*http.Response, error)
	EventcontrolEventcontrolRootRequestPost(context.Context, *EventcontrolApiEventcontrolEventcontrolRootRequestPostOpts) (*http.Response, error)
	EventcontrolVmVmIDRequestPost(context.Context, string, *EventcontrolApiEventcontrolVmVmIDRequestPostOpts) (*http.Response, error)
}

type IExcludelistApi interface {
	AppExcludelistGet(context.Context) (*http.Response, error)
	AppExcludelistMemberIDDelete(context.Context, string) (*http.Response, error)
	AppExcludelistMemberIDPut(context.Context, string) (*http.Response, error)
}

type IFirewallApi interface {
	FirewallConfigGlobalconfigurationGet(context.Context) (*http.Response, error)
	FirewallConfigGlobalconfigurationPut(context.Context, *FirewallApiFirewallConfigGlobalconfigurationPutOpts) (*http.Response, error)
	FirewallContextIdConfigIpfixDelete(context.Context, string) (*http.Response, error)
	FirewallContextIdConfigIpfixGet(context.Context, string) (*http.Response, error)
	FirewallContextIdConfigIpfixPut(context.Context, string, *FirewallApiFirewallContextIdConfigIpfixPutOpts) (*http.Response, error)
	FirewallDomainIDEnableTruefalsePost(context.Context, string, string) (*http.Response, error)
	FirewallForceSyncIDPost(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0ConfigDelete(context.Context) (*http.Response, error)
	FirewallGlobalroot0ConfigGet(context.Context, *FirewallApiFirewallGlobalroot0ConfigGetOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer2sectionsGet(context.Context, *FirewallApiFirewallGlobalroot0ConfigLayer2sectionsGetOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer2sectionsPost(context.Context, *FirewallApiFirewallGlobalroot0ConfigLayer2sectionsPostOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer2sectionsSectionIdDelete(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer2sectionsSectionIdGet(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer2sectionsSectionIdPost(context.Context, string, *FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdPostOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer2sectionsSectionIdPut(context.Context, string, *FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdPutOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesPost(context.Context, string, *FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesPostOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdDelete(context.Context, string, string, *FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdDeleteOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdGet(context.Context, string, string) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdPut(context.Context, string, string, *FirewallApiFirewallGlobalroot0ConfigLayer2sectionsSectionIdRulesRuleIdPutOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3redirectProfilesGet(context.Context) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3redirectsectionsPost(context.Context, *FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsPostOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3redirectsectionsSectionDelete(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3redirectsectionsSectionGet(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3redirectsectionsSectionPut(context.Context, string, *FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsSectionPutOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesPost(context.Context, string, *FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesPostOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDDelete(context.Context, string, string) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDGet(context.Context, string, string) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDPut(context.Context, string, string, *FirewallApiFirewallGlobalroot0ConfigLayer3redirectsectionsSectionRulesRuleIDPutOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3sectionsGet(context.Context, *FirewallApiFirewallGlobalroot0ConfigLayer3sectionsGetOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3sectionsPost(context.Context, *FirewallApiFirewallGlobalroot0ConfigLayer3sectionsPostOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3sectionsSectionIdDelete(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3sectionsSectionIdGet(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3sectionsSectionIdPost(context.Context, string, *FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdPostOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3sectionsSectionIdPut(context.Context, string, *FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdPutOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesPost(context.Context, string, *FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesPostOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdDelete(context.Context, string, string, *FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdDeleteOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdGet(context.Context, string, string) (*http.Response, error)
	FirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdPut(context.Context, string, string, *FirewallApiFirewallGlobalroot0ConfigLayer3sectionsSectionIdRulesRuleIdPutOpts) (*http.Response, error)
	FirewallGlobalroot0ConfigPut(context.Context, *FirewallApiFirewallGlobalroot0ConfigPutOpts) (*http.Response, error)
	FirewallGlobalroot0DefaultconfigGet(context.Context) (*http.Response, error)
	FirewallGlobalroot0DraftsDraftIDActionExportGet(context.Context, string, *FirewallApiFirewallGlobalroot0DraftsDraftIDActionExportGetOpts) (*http.Response, error)
	FirewallGlobalroot0DraftsDraftIDActionImportPost(context.Context, string, *FirewallApiFirewallGlobalroot0DraftsDraftIDActionImportPostOpts) (*http.Response, error)
	FirewallGlobalroot0DraftsDraftIDDelete(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0DraftsDraftIDGet(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0DraftsDraftIDPut(context.Context, string, *FirewallApiFirewallGlobalroot0DraftsDraftIDPutOpts) (*http.Response, error)
	FirewallGlobalroot0DraftsGet(context.Context) (*http.Response, error)
	FirewallGlobalroot0DraftsPost(context.Context, *FirewallApiFirewallGlobalroot0DraftsPostOpts) (*http.Response, error)
	FirewallGlobalroot0StateGet(context.Context) (*http.Response, error)
	FirewallGlobalroot0StatePut(context.Context) (*http.Response, error)
	FirewallGlobalroot0StatusGet(context.Context) (*http.Response, error)
	FirewallGlobalroot0StatusLayer2sectionsSectionIDGet(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0StatusLayer3sectionsSectionIDGet(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0TimeoutsConfigIdDelete(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0TimeoutsConfigIdGet(context.Context, string) (*http.Response, error)
	FirewallGlobalroot0TimeoutsConfigIdPut(context.Context, string, *FirewallApiFirewallGlobalroot0TimeoutsConfigIdPutOpts) (*http.Response, error)
	FirewallGlobalroot0TimeoutsGet(context.Context) (*http.Response, error)
	FirewallGlobalroot0TimeoutsPost(context.Context, *FirewallApiFirewallGlobalroot0TimeoutsPostOpts) (*http.Response, error)
	FirewallStatsEventthresholdsGet(context.Context) (*http.Response, error)
	FirewallStatsEventthresholdsPut(context.Context, *FirewallApiFirewallStatsEventthresholdsPutOpts) (*http.Response, error)
}

type IFlowApi interface {
	AppFlowConfigGet(context.Context) (*http.Response, error)
	AppFlowConfigPut(context.Context, *FlowApiAppFlowConfigPutOpts) (*http.Response, error)
	AppFlowContextIdDelete(context.Context, string) (*http.Response, error)
	AppFlowFlowstatsGet(context.Context, *FlowApiAppFlowFlowstatsGetOpts) (*http.Response, error)
	AppFlowFlowstatsInfoGet(context.Context) (*http.Response, error)
}

type IHardwaregatewaysApi interface {
	VdnHardwaregatewaysBfdConfigGet(context.Context) (*http.Response, error)
	VdnHardwaregatewaysBfdConfigPut(context.Context, *HardwaregatewaysApiVdnHardwaregatewaysBfdConfigPutOpts) (*http.Response, error)
	VdnHardwaregatewaysBfdStatusGet(context.Context) (*http.Response, error)
	VdnHardwaregatewaysBindingsBindingIdDelete(context.Context, string) (*http.Response, error)
	VdnHardwaregatewaysBindingsBindingIdGet(context.Context, string) (*http.Response, error)
	VdnHardwaregatewaysBindingsBindingIdPut(context.Context, string, *HardwaregatewaysApiVdnHardwaregatewaysBindingsBindingIdPutOpts) (*http.Response, error)
	VdnHardwaregatewaysBindingsBindingIdStatisticGet(context.Context, string) (*http.Response, error)
	VdnHardwaregatewaysBindingsGet(context.Context, *HardwaregatewaysApiVdnHardwaregatewaysBindingsGetOpts) (*http.Response, error)
	VdnHardwaregatewaysBindingsManagePost(context.Context, *HardwaregatewaysApiVdnHardwaregatewaysBindingsManagePostOpts) (*http.Response, error)
	VdnHardwaregatewaysBindingsPost(context.Context, *HardwaregatewaysApiVdnHardwaregatewaysBindingsPostOpts) (*http.Response, error)
	VdnHardwaregatewaysGet(context.Context) (*http.Response, error)
	VdnHardwaregatewaysHardwareGatewayIdDelete(context.Context, string) (*http.Response, error)
	VdnHardwaregatewaysHardwareGatewayIdGet(context.Context, string) (*http.Response, error)
	VdnHardwaregatewaysHardwareGatewayIdPut(context.Context, string, *HardwaregatewaysApiVdnHardwaregatewaysHardwareGatewayIdPutOpts) (*http.Response, error)
	VdnHardwaregatewaysHardwareGatewayIdSwitchesGet(context.Context, string) (*http.Response, error)
	VdnHardwaregatewaysHardwareGatewayIdSwitchesSwitchNameSwitchportsGet(context.Context, string, string) (*http.Response, error)
	VdnHardwaregatewaysPost(context.Context, *HardwaregatewaysApiVdnHardwaregatewaysPostOpts) (*http.Response, error)
	VdnHardwaregatewaysReplicationclusterGet(context.Context) (*http.Response, error)
	VdnHardwaregatewaysReplicationclusterPut(context.Context, *HardwaregatewaysApiVdnHardwaregatewaysReplicationclusterPutOpts) (*http.Response, error)
}

type IIdentityApi interface {
	IdentityDirectoryGroupsForUserGet(context.Context) (*http.Response, error)
	IdentityHostIpMappingGet(context.Context) (*http.Response, error)
	IdentityIpToUserMappingGet(context.Context) (*http.Response, error)
	IdentityStaticUserMappingUserIDIPPost(context.Context, string, string) (*http.Response, error)
	IdentityStaticUserMappingsGet(context.Context) (*http.Response, error)
	IdentityStaticUserMappingsbyIPIPDelete(context.Context, string) (*http.Response, error)
	IdentityStaticUserMappingsbyIPIPGet(context.Context, string) (*http.Response, error)
	IdentityStaticUserMappingsbyUserUserIDDelete(context.Context, string) (*http.Response, error)
	IdentityStaticUserMappingsbyUserUserIDGet(context.Context, string) (*http.Response, error)
	IdentityUserIpMappingGet(context.Context) (*http.Response, error)
}

type IInventoryApi interface {
	VdnInventoryHostHostIdConnectionStatusGet(context.Context, string) (*http.Response, error)
	VdnInventoryHostsConnectionStatusGet(context.Context, *InventoryApiVdnInventoryHostsConnectionStatusGetOpts) (*http.Response, error)
}

type IIpsetApi interface {
	ServicesIpsetIpsetIdDelete(context.Context, string, *IpsetApiServicesIpsetIpsetIdDeleteOpts) (*http.Response, error)
	ServicesIpsetIpsetIdGet(context.Context, string) (*http.Response, error)
	ServicesIpsetIpsetIdPut(context.Context, string, *IpsetApiServicesIpsetIpsetIdPutOpts) (*http.Response, error)
	ServicesIpsetScopeMorefPost(context.Context, string, *IpsetApiServicesIpsetScopeMorefPostOpts) (*http.Response, error)
	ServicesIpsetScopeScopeMorefGet(context.Context, string) (*http.Response, error)
}

type IJobApi interface {
	ServicesTaskserviceJobGet(context.Context, *JobApiServicesTaskserviceJobGetOpts) (*http.Response, error)
	ServicesTaskserviceJobJobIdGet(context.Context, string) (*http.Response, error)
}

type IMacsetApi interface {
	ServicesMacsetMacsetIdDelete(context.Context, string, *MacsetApiServicesMacsetMacsetIdDeleteOpts) (*http.Response, error)
	ServicesMacsetMacsetIdGet(context.Context, string) (*http.Response, error)
	ServicesMacsetMacsetIdPut(context.Context, string, *MacsetApiServicesMacsetMacsetIdPutOpts) (*http.Response, error)
	ServicesMacsetScopeScopeIdGet(context.Context, string) (*http.Response, error)
	ServicesMacsetScopeScopeIdPost(context.Context, string, *MacsetApiServicesMacsetScopeScopeIdPostOpts) (*http.Response, error)
}

type INetworkFeaturesApi interface {
	XvsNetworksIDFeaturesGet(context.Context, string) (*http.Response, error)
	XvsNetworksIDFeaturesPut(context.Context, string, *NetworkFeaturesApiXvsNetworksIDFeaturesPutOpts) (*http.Response, error)
}

type INwfabricApi interface {
	NwfabricClustersClusterIDDelete(context.Context, string) (*http.Response, error)
	NwfabricClustersClusterIDGet(context.Context, string) (*http.Response, error)
	NwfabricClustersClusterIDPut(context.Context, string, *NwfabricApiNwfabricClustersClusterIDPutOpts) (*http.Response, error)
	NwfabricConfigureDelete(context.Context, *NwfabricApiNwfabricConfigureDeleteOpts) (*http.Response, error)
	NwfabricConfigurePost(context.Context, *NwfabricApiNwfabricConfigurePostOpts) (*http.Response, error)
	NwfabricConfigurePut(context.Context, *NwfabricApiNwfabricConfigurePutOpts) (*http.Response, error)
	NwfabricFeaturesGet(context.Context) (*http.Response, error)
	NwfabricHostsHostIDDelete(context.Context, string) (*http.Response, error)
	NwfabricHostsHostIDGet(context.Context, string) (*http.Response, error)
	NwfabricHostsHostIDPut(context.Context, string, *NwfabricApiNwfabricHostsHostIDPutOpts) (*http.Response, error)
	NwfabricStatusAlleligibleResourceTypeGet(context.Context, string) (*http.Response, error)
	NwfabricStatusChildParentResourceIDGet(context.Context, string) (*http.Response, error)
	NwfabricStatusGet(context.Context, *NwfabricApiNwfabricStatusGetOpts) (*http.Response, error)
}

type IPolicyApi interface {
	ServicesPolicySecurityactionCategoryVirtualmachinesGet(context.Context, *PolicyApiServicesPolicySecurityactionCategoryVirtualmachinesGetOpts) (*http.Response, error)
	ServicesPolicySecuritygroupIDSecurityactionsGet(context.Context, string) (*http.Response, error)
	ServicesPolicySecuritygroupIDSecuritypoliciesGet(context.Context, string) (*http.Response, error)
	ServicesPolicySecuritypolicyAlarmsAllGet(context.Context) (*http.Response, error)
	ServicesPolicySecuritypolicyHierarchyGet(context.Context, *PolicyApiServicesPolicySecuritypolicyHierarchyGetOpts) (*http.Response, error)
	ServicesPolicySecuritypolicyHierarchyPost(context.Context, *PolicyApiServicesPolicySecuritypolicyHierarchyPostOpts) (*http.Response, error)
	ServicesPolicySecuritypolicyIDDelete(context.Context, string, *PolicyApiServicesPolicySecuritypolicyIDDeleteOpts) (*http.Response, error)
	ServicesPolicySecuritypolicyIDGet(context.Context, string) (*http.Response, error)
	ServicesPolicySecuritypolicyIDPut(context.Context, string, *PolicyApiServicesPolicySecuritypolicyIDPutOpts) (*http.Response, error)
	ServicesPolicySecuritypolicyIDSecurityactionsGet(context.Context, string) (*http.Response, error)
	ServicesPolicySecuritypolicyPost(context.Context, *PolicyApiServicesPolicySecuritypolicyPostOpts) (*http.Response, error)
	ServicesPolicySecuritypolicyServiceproviderFirewallGet(context.Context) (*http.Response, error)
	ServicesPolicySecuritypolicyServiceproviderFirewallPut(context.Context, *PolicyApiServicesPolicySecuritypolicyServiceproviderFirewallPutOpts) (*http.Response, error)
	ServicesPolicySecuritypolicyStatusGet(context.Context) (*http.Response, error)
	ServicesPolicyServiceproviderFirewallGet(context.Context, *PolicyApiServicesPolicyServiceproviderFirewallGetOpts) (*http.Response, error)
	ServicesPolicyVirtualmachineIDSecurityactionsGet(context.Context, string) (*http.Response, error)
}

type IPoolsApi interface {
	ServicesIpamPoolsPoolIdDelete(context.Context, string) (*http.Response, error)
	ServicesIpamPoolsPoolIdGet(context.Context, string) (*http.Response, error)
	ServicesIpamPoolsPoolIdIpaddressesGet(context.Context, string) (*http.Response, error)
	ServicesIpamPoolsPoolIdIpaddressesIpAddressDelete(context.Context, string, string) (*http.Response, error)
	ServicesIpamPoolsPoolIdIpaddressesPost(context.Context, string, *PoolsApiServicesIpamPoolsPoolIdIpaddressesPostOpts) (*http.Response, error)
	ServicesIpamPoolsPoolIdPut(context.Context, string, *PoolsApiServicesIpamPoolsPoolIdPutOpts) (*http.Response, error)
	ServicesIpamPoolsScopeScopeIdGet(context.Context, string) (*http.Response, error)
	ServicesIpamPoolsScopeScopeIdPost(context.Context, string, *PoolsApiServicesIpamPoolsScopeScopeIdPostOpts) (*http.Response, error)
}

type IScopesApi interface {
	VdnScopesGet(context.Context) (*http.Response, error)
	VdnScopesPost(context.Context, *ScopesApiVdnScopesPostOpts) (*http.Response, error)
	VdnScopesScopeIdAttributesPut(context.Context, string, *ScopesApiVdnScopesScopeIdAttributesPutOpts) (*http.Response, error)
	VdnScopesScopeIdConnCheckMulticastPost(context.Context, string, *ScopesApiVdnScopesScopeIdConnCheckMulticastPostOpts) (*http.Response, error)
	VdnScopesScopeIdDelete(context.Context, string) (*http.Response, error)
	VdnScopesScopeIdGet(context.Context, string) (*http.Response, error)
	VdnScopesScopeIdPost(context.Context, string, *ScopesApiVdnScopesScopeIdPostOpts) (*http.Response, error)
}

type ISecuritygroupApi interface {
	ServicesSecuritygroupBulkObjectIdPut(context.Context, string, *SecuritygroupApiServicesSecuritygroupBulkObjectIdPutOpts) (*http.Response, error)
	ServicesSecuritygroupBulkScopeIdPost(context.Context, string, *SecuritygroupApiServicesSecuritygroupBulkScopeIdPostOpts) (*http.Response, error)
	ServicesSecuritygroupInternalScopeScopeIdGet(context.Context, string) (*http.Response, error)
	ServicesSecuritygroupLookupVirtualmachineVirtualMachineIdGet(context.Context, string) (*http.Response, error)
	ServicesSecuritygroupObjectIdDelete(context.Context, string, *SecuritygroupApiServicesSecuritygroupObjectIdDeleteOpts) (*http.Response, error)
	ServicesSecuritygroupObjectIdGet(context.Context, string) (*http.Response, error)
	ServicesSecuritygroupObjectIdMembersMemberIdDelete(context.Context, string, string, *SecuritygroupApiServicesSecuritygroupObjectIdMembersMemberIdDeleteOpts) (*http.Response, error)
	ServicesSecuritygroupObjectIdMembersMemberIdPut(context.Context, string, string, *SecuritygroupApiServicesSecuritygroupObjectIdMembersMemberIdPutOpts) (*http.Response, error)
	ServicesSecuritygroupObjectIdPut(context.Context, string, *SecuritygroupApiServicesSecuritygroupObjectIdPutOpts) (*http.Response, error)
	ServicesSecuritygroupObjectIdTranslationIpaddressesGet(context.Context, string) (*http.Response, error)
	ServicesSecuritygroupObjectIdTranslationMacaddressesGet(context.Context, string) (*http.Response, error)
	ServicesSecuritygroupObjectIdTranslationVirtualmachinesGet(context.Context, string) (*http.Response, error)
	ServicesSecuritygroupObjectIdTranslationVnicsGet(context.Context, string) (*http.Response, error)
	ServicesSecuritygroupScopeIdPost(context.Context, string, *SecuritygroupApiServicesSecuritygroupScopeIdPostOpts) (*http.Response, error)
	ServicesSecuritygroupScopeScopeIdGet(context.Context, string) (*http.Response, error)
	ServicesSecuritygroupScopeScopeIdMemberTypesGet(context.Context, string) (*http.Response, error)
	ServicesSecuritygroupScopeScopeIdMembersMemberTypeGet(context.Context, string, string) (*http.Response, error)
}

type ISecuritytagsApi interface {
	ServicesSecuritytagsSelectionCriteriaGet(context.Context) (*http.Response, error)
	ServicesSecuritytagsSelectionCriteriaPut(context.Context, *SecuritytagsApiServicesSecuritytagsSelectionCriteriaPutOpts) (*http.Response, error)
	ServicesSecuritytagsTagGet(context.Context, *SecuritytagsApiServicesSecuritytagsTagGetOpts) (*http.Response, error)
	ServicesSecuritytagsTagPost(context.Context, *SecuritytagsApiServicesSecuritytagsTagPostOpts) (*http.Response, error)
	ServicesSecuritytagsTagTagIdDelete(context.Context, string) (*http.Response, error)
	ServicesSecuritytagsTagTagIdVmDetailGet(context.Context, string) (*http.Response, error)
	ServicesSecuritytagsTagTagIdVmGet(context.Context, string) (*http.Response, error)
	ServicesSecuritytagsTagTagIdVmPost(context.Context, string, *SecuritytagsApiServicesSecuritytagsTagTagIdVmPostOpts) (*http.Response, error)
	ServicesSecuritytagsTagTagIdVmVmIdDelete(context.Context, string, string) (*http.Response, error)
	ServicesSecuritytagsTagTagIdVmVmIdPut(context.Context, string, string) (*http.Response, error)
	ServicesSecuritytagsVmVmIdGet(context.Context, string) (*http.Response, error)
	ServicesSecuritytagsVmVmIdPost(context.Context, string, *SecuritytagsApiServicesSecuritytagsVmVmIdPostOpts) (*http.Response, error)
}

type ISiApi interface {
	SiAgentAgentIDGet(context.Context, string) (*http.Response, error)
	SiDeploymentDeploymentunitIDAgentsGet(context.Context, string) (*http.Response, error)
	SiFabricSyncConflictsGet(context.Context) (*http.Response, error)
	SiFabricSyncConflictsPut(context.Context, *SiApiSiFabricSyncConflictsPutOpts) (*http.Response, error)
	SiHostHostIDAgentsGet(context.Context, string) (*http.Response, error)
}

type ISnmpApi interface {
	ServicesSnmpManagerGet(context.Context) (*http.Response, error)
	ServicesSnmpManagerManagerIdDelete(context.Context, string) (*http.Response, error)
	ServicesSnmpManagerManagerIdGet(context.Context, string) (*http.Response, error)
	ServicesSnmpManagerManagerIdPut(context.Context, string, *SnmpApiServicesSnmpManagerManagerIdPutOpts) (*http.Response, error)
	ServicesSnmpManagerPost(context.Context, *SnmpApiServicesSnmpManagerPostOpts) (*http.Response, error)
	ServicesSnmpStatusGet(context.Context) (*http.Response, error)
	ServicesSnmpStatusPut(context.Context, *SnmpApiServicesSnmpStatusPutOpts) (*http.Response, error)
	ServicesSnmpTrapGet(context.Context) (*http.Response, error)
	ServicesSnmpTrapOidGet(context.Context, string) (*http.Response, error)
	ServicesSnmpTrapOidPut(context.Context, string, *SnmpApiServicesSnmpTrapOidPutOpts) (*http.Response, error)
}

type ISpoofguardApi interface {
	ServicesSpoofguardPoliciesGet(context.Context) (*http.Response, error)
	ServicesSpoofguardPoliciesPolicyIDDelete(context.Context, string) (*http.Response, error)
	ServicesSpoofguardPoliciesPolicyIDGet(context.Context, string) (*http.Response, error)
	ServicesSpoofguardPoliciesPolicyIDPut(context.Context, string, *SpoofguardApiServicesSpoofguardPoliciesPolicyIDPutOpts) (*http.Response, error)
	ServicesSpoofguardPoliciesPost(context.Context, *SpoofguardApiServicesSpoofguardPoliciesPostOpts) (*http.Response, error)
	ServicesSpoofguardPolicyIDGet(context.Context, string, *SpoofguardApiServicesSpoofguardPolicyIDGetOpts) (*http.Response, error)
	ServicesSpoofguardPolicyIDPost(context.Context, string, *SpoofguardApiServicesSpoofguardPolicyIDPostOpts) (*http.Response, error)
}

type ISsoconfigApi interface {
	ServicesSsoconfigDelete(context.Context) (*http.Response, error)
	ServicesSsoconfigGet(context.Context) (*http.Response, error)
	ServicesSsoconfigPost(context.Context, *SsoconfigApiServicesSsoconfigPostOpts) (*http.Response, error)
	ServicesSsoconfigStatusGet(context.Context) (*http.Response, error)
}

type ISwitchesApi interface {
	VdnSwitchesDatacenterDatacenterIDGet(context.Context, string) (*http.Response, error)
	VdnSwitchesGet(context.Context) (*http.Response, error)
	VdnSwitchesPost(context.Context, *SwitchesApiVdnSwitchesPostOpts) (*http.Response, error)
	VdnSwitchesVdsIdDelete(context.Context, string) (*http.Response, error)
	VdnSwitchesVdsIdGet(context.Context, string) (*http.Response, error)
}

type ISyslogApi interface {
	SamSyslogDisablePost(context.Context) (*http.Response, error)
	SamSyslogEnablePost(context.Context) (*http.Response, error)
}

type ISystemalarmsApi interface {
	ServicesSystemalarmsAlarmIdGet(context.Context, string) (*http.Response, error)
	ServicesSystemalarmsAlarmIdPost(context.Context, string, *SystemalarmsApiServicesSystemalarmsAlarmIdPostOpts) (*http.Response, error)
}

type ISystemeventApi interface {
	SystemeventGet(context.Context, *SystemeventApiSystemeventGetOpts) (*http.Response, error)
}

type ITraceflowApi interface {
	ApiVdnTraceflowPost(context.Context, *TraceflowApiApiVdnTraceflowPostOpts) (*http.Response, error)
	ApiVdnTraceflowTraceflowIdGet(context.Context, string) (*http.Response, error)
	ApiVdnTraceflowTraceflowIdObservationsGet(context.Context, string) (*http.Response, error)
}

type ITruststoreApi interface {
	ServicesTruststoreCertificateCertificateIdDelete(context.Context, string) (*http.Response, error)
	ServicesTruststoreCertificateCertificateIdGet(context.Context, string) (*http.Response, error)
	ServicesTruststoreCertificatePost(context.Context, *TruststoreApiServicesTruststoreCertificatePostOpts) (*http.Response, error)
	ServicesTruststoreCertificateScopeIdPost(context.Context, string, *TruststoreApiServicesTruststoreCertificateScopeIdPostOpts) (*http.Response, error)
	ServicesTruststoreCertificateScopeScopeIdGet(context.Context, string) (*http.Response, error)
	ServicesTruststoreCrlCrlIdDelete(context.Context, string) (*http.Response, error)
	ServicesTruststoreCrlCrlIdGet(context.Context, string) (*http.Response, error)
	ServicesTruststoreCrlScopeIdPost(context.Context, string, *TruststoreApiServicesTruststoreCrlScopeIdPostOpts) (*http.Response, error)
	ServicesTruststoreCrlScopeScopeIdGet(context.Context, string) (*http.Response, error)
	ServicesTruststoreCsrCsrIdGet(context.Context, string) (*http.Response, error)
	ServicesTruststoreCsrCsrIdPut(context.Context, string, *TruststoreApiServicesTruststoreCsrCsrIdPutOpts) (*http.Response, error)
	ServicesTruststoreCsrScopeIdPost(context.Context, string, *TruststoreApiServicesTruststoreCsrScopeIdPostOpts) (*http.Response, error)
	ServicesTruststoreCsrScopeScopeIdGet(context.Context, string) (*http.Response, error)
}

type IUniversalsyncApi interface {
	UniversalsyncConfigurationNsxmanagersDelete(context.Context) (*http.Response, error)
	UniversalsyncConfigurationNsxmanagersGet(context.Context) (*http.Response, error)
	UniversalsyncConfigurationNsxmanagersNsxManagerIDDelete(context.Context, string, *UniversalsyncApiUniversalsyncConfigurationNsxmanagersNsxManagerIDDeleteOpts) (*http.Response, error)
	UniversalsyncConfigurationNsxmanagersNsxManagerIDGet(context.Context, string) (*http.Response, error)
	UniversalsyncConfigurationNsxmanagersNsxManagerIDPut(context.Context, string, *UniversalsyncApiUniversalsyncConfigurationNsxmanagersNsxManagerIDPutOpts) (*http.Response, error)
	UniversalsyncConfigurationNsxmanagersPost(context.Context, *UniversalsyncApiUniversalsyncConfigurationNsxmanagersPostOpts) (*http.Response, error)
	UniversalsyncConfigurationRoleGet(context.Context) (*http.Response, error)
	UniversalsyncConfigurationRolePost(context.Context, *UniversalsyncApiUniversalsyncConfigurationRolePostOpts) (*http.Response, error)
	UniversalsyncEntitystatusGet(context.Context, *UniversalsyncApiUniversalsyncEntitystatusGetOpts) (*http.Response, error)
	UniversalsyncStatusGet(context.Context) (*http.Response, error)
	UniversalsyncSyncPost(context.Context, *UniversalsyncApiUniversalsyncSyncPostOpts) (*http.Response, error)
}

type IUsermgmtApi interface {
	ServicesUsermgmtEnablestateValuePut(context.Context, string) (*http.Response, error)
	ServicesUsermgmtRoleUserIdDelete(context.Context, string) (*http.Response, error)
	ServicesUsermgmtRoleUserIdGet(context.Context, string) (*http.Response, error)
	ServicesUsermgmtRoleUserIdPost(context.Context, string, *UsermgmtApiServicesUsermgmtRoleUserIdPostOpts) (*http.Response, error)
	ServicesUsermgmtRoleUserIdPut(context.Context, string, *UsermgmtApiServicesUsermgmtRoleUserIdPutOpts) (*http.Response, error)
	ServicesUsermgmtRolesGet(context.Context) (*http.Response, error)
	ServicesUsermgmtScopingobjectsGet(context.Context) (*http.Response, error)
	ServicesUsermgmtUserUserIdDelete(context.Context, string) (*http.Response, error)
	ServicesUsermgmtUserUserIdGet(context.Context, string) (*http.Response, error)
	ServicesUsermgmtUsersVsmGet(context.Context) (*http.Response, error)
}

type IVcconfigApi interface {
	ServicesVcconfigGet(context.Context) (*http.Response, error)
	ServicesVcconfigPut(context.Context, *VcconfigApiServicesVcconfigPutOpts) (*http.Response, error)
	ServicesVcconfigStatusGet(context.Context) (*http.Response, error)
}

type IVirtualwiresApi interface {
	VdnScopesScopeIdVirtualwiresGet(context.Context, string, *VirtualwiresApiVdnScopesScopeIdVirtualwiresGetOpts) (*http.Response, error)
	VdnScopesScopeIdVirtualwiresPost(context.Context, string, *VirtualwiresApiVdnScopesScopeIdVirtualwiresPostOpts) (*http.Response, error)
	VdnVirtualwiresGet(context.Context, *VirtualwiresApiVdnVirtualwiresGetOpts) (*http.Response, error)
	VdnVirtualwiresVirtualWireIDBackingPost(context.Context, string, *VirtualwiresApiVdnVirtualwiresVirtualWireIDBackingPostOpts) (*http.Response, error)
	VdnVirtualwiresVirtualWireIDConnCheckMulticastPost(context.Context, string, *VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckMulticastPostOpts) (*http.Response, error)
	VdnVirtualwiresVirtualWireIDConnCheckP2pPost(context.Context, string, *VirtualwiresApiVdnVirtualwiresVirtualWireIDConnCheckP2pPostOpts) (*http.Response, error)
	VdnVirtualwiresVirtualWireIDDelete(context.Context, string) (*http.Response, error)
	VdnVirtualwiresVirtualWireIDGet(context.Context, string) (*http.Response, error)
	VdnVirtualwiresVirtualWireIDHardwaregatewaysGet(context.Context, string) (*http.Response, error)
	VdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPost(context.Context, string, string, *VirtualwiresApiVdnVirtualwiresVirtualWireIDHardwaregatewaysHardwareGatewayBindingIdPostOpts) (*http.Response, error)
	VdnVirtualwiresVirtualWireIDPut(context.Context, string, *VirtualwiresApiVdnVirtualwiresVirtualWireIDPutOpts) (*http.Response, error)
	VdnVirtualwiresVmVnicPost(context.Context, *VirtualwiresApiVdnVirtualwiresVmVnicPostOpts) (*http.Response, error)
}
