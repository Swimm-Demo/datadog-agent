---
title: Exploring SNMP Check Configuration
---
```mermaid
classDiagram
 CheckConfig <|-- InitConfig
 CheckConfig <|-- InstanceConfig
 CheckConfig : +string Name
 CheckConfig : +string IPAddress
 CheckConfig : +uint16 Port
 CheckConfig : +string SnmpVersion
 CheckConfig : +int Timeout
 CheckConfig : +int Retries
 CheckConfig : +string User
 CheckConfig : +string AuthProtocol
 CheckConfig : +string AuthKey
 CheckConfig : +string PrivProtocol
 CheckConfig : +string PrivKey
 CheckConfig : +string ContextName
 CheckConfig : +OidConfig OidConfig
 CheckConfig : +[]MetricsConfig RequestedMetrics
 CheckConfig : +[]MetricTagConfig RequestedMetricTags
 CheckConfig : +[]MetricsConfig Metrics
 CheckConfig : +MetadataConfig Metadata
 CheckConfig : +[]MetricTagConfig MetricTags
 CheckConfig : +int OidBatchSize
 CheckConfig : +uint32 BulkMaxRepetitions
 CheckConfig : +ProfileConfigMap Profiles
 CheckConfig : +[]string ProfileTags
 CheckConfig : +string Profile
 CheckConfig : +ProfileDefinition ProfileDef
 CheckConfig : +[]string ExtraTags
 CheckConfig : +[]string InstanceTags
 CheckConfig : +bool CollectDeviceMetadata
 CheckConfig : +bool CollectTopology
 CheckConfig : +bool UseDeviceIDAsHostname
 CheckConfig : +string DeviceID
 CheckConfig : +[]string DeviceIDTags
 CheckConfig : +string ResolvedSubnetName
 CheckConfig : +string Namespace
 CheckConfig : +bool AutodetectProfile
 CheckConfig : +time.Duration MinCollectionInterval
 CheckConfig : +bool DetectMetricsEnabled
 CheckConfig : +int DetectMetricsRefreshInterval
 CheckConfig : +string Network
 CheckConfig : +int DiscoveryWorkers
 CheckConfig : +int Workers
 CheckConfig : +int DiscoveryInterval
 CheckConfig : +map IgnoredIPAddresses
 CheckConfig : +int DiscoveryAllowedFailures
 CheckConfig : +[]InterfaceConfig InterfaceConfigs
 CheckConfig : +bool PingEnabled
 CheckConfig : +Config PingConfig
 InitConfig : +ProfileConfigMap Profiles
 InitConfig : +[]MetricsConfig GlobalMetrics
 InitConfig : +Number OidBatchSize
 InitConfig : +Number BulkMaxRepetitions
 InitConfig : +Boolean CollectDeviceMetadata
 InitConfig : +Boolean CollectTopology
 InitConfig : +Boolean UseDeviceIDAsHostname
 InitConfig : +int MinCollectionInterval
 InitConfig : +string Namespace
 InitConfig : +PackedPingConfig PingConfig
 InitConfig : +Boolean DetectMetricsEnabled
 InitConfig : +int DetectMetricsRefreshInterval
 InstanceConfig : +string Name
 InstanceConfig : +string IPAddress
 InstanceConfig : +Number Port
 InstanceConfig : +string CommunityString
 InstanceConfig : +string SnmpVersion
 InstanceConfig : +Number Timeout
 InstanceConfig : +Number Retries
 InstanceConfig : +string User
 InstanceConfig : +string AuthProtocol
 InstanceConfig : +string AuthKey
 InstanceConfig : +string PrivProtocol
 InstanceConfig : +string PrivKey
 InstanceConfig : +string ContextName
 InstanceConfig : +[]MetricsConfig Metrics
 InstanceConfig : +[]MetricTagConfig MetricTags
 InstanceConfig : +string Profile
 InstanceConfig : +bool UseGlobalMetrics
 InstanceConfig : +Boolean CollectDeviceMetadata
 InstanceConfig : +Boolean CollectTopology
 InstanceConfig : +Boolean UseDeviceIDAsHostname
 InstanceConfig : +PackedPingConfig PingConfig
 InstanceConfig : +string ExtraTags

%% Swimm:
%% classDiagram
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> <|-- <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> <|-- <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string Name
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="101:1:1" line-data="	IPAddress             string                              `yaml:&quot;ip_address&quot;`">`IPAddress`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +uint16 Port
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="104:1:1" line-data="	SnmpVersion           string                              `yaml:&quot;snmp_version&quot;`">`SnmpVersion`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +int Timeout
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +int Retries
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string User
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="108:1:1" line-data="	AuthProtocol          string                              `yaml:&quot;authProtocol&quot;`">`AuthProtocol`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="109:1:1" line-data="	AuthKey               string                              `yaml:&quot;authKey&quot;`">`AuthKey`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="110:1:1" line-data="	PrivProtocol          string                              `yaml:&quot;privProtocol&quot;`">`PrivProtocol`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="111:1:1" line-data="	PrivKey               string                              `yaml:&quot;privKey&quot;`">`PrivKey`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="112:1:1" line-data="	ContextName           string                              `yaml:&quot;context_name&quot;`">`ContextName`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +OidConfig <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="176:1:1" line-data="	OidConfig       OidConfig">`OidConfig`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +[]MetricsConfig <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="177:3:3" line-data="	// RequestedMetrics are the metrics explicitly requested by config.">`RequestedMetrics`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +[]MetricTagConfig <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="179:3:3" line-data="	// RequestedMetricTags are the tags explicitly requested by config.">`RequestedMetricTags`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +[]MetricsConfig Metrics
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +MetadataConfig Metadata
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +[]MetricTagConfig <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="114:1:1" line-data="	MetricTags            []profiledefinition.MetricTagConfig `yaml:&quot;metric_tags&quot;` // SNMP metric tags definition">`MetricTags`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +int <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="86:1:1" line-data="	OidBatchSize                 Number                            `yaml:&quot;oid_batch_size&quot;`">`OidBatchSize`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +uint32 <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="87:1:1" line-data="	BulkMaxRepetitions           Number                            `yaml:&quot;bulk_max_repetitions&quot;`">`BulkMaxRepetitions`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +ProfileConfigMap Profiles
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +[]string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="239:3:3" line-data="	c.ProfileTags = tags">`ProfileTags`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string Profile
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +ProfileDefinition <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="227:3:3" line-data="	c.ProfileDef = &amp;definition">`ProfileDef`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +[]string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="122:3:3" line-data="	// ExtraTags is a workaround to pass tags from snmp listener to snmp integration via AD template">`ExtraTags`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +[]string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="193:1:1" line-data="	InstanceTags          []string">`InstanceTags`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +bool <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="88:1:1" line-data="	CollectDeviceMetadata        Boolean                           `yaml:&quot;collect_device_metadata&quot;`">`CollectDeviceMetadata`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +bool <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="89:1:1" line-data="	CollectTopology              Boolean                           `yaml:&quot;collect_topology&quot;`">`CollectTopology`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +bool <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="90:1:1" line-data="	UseDeviceIDAsHostname        Boolean                           `yaml:&quot;use_device_id_as_hostname&quot;`">`UseDeviceIDAsHostname`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="275:6:6" line-data="// UpdateDeviceIDAndTags updates DeviceID and DeviceIDTags">`DeviceID`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +[]string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="275:10:10" line-data="// UpdateDeviceIDAndTags updates DeviceID and DeviceIDTags">`DeviceIDTags`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="199:1:1" line-data="	ResolvedSubnetName    string">`ResolvedSubnetName`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string Namespace
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +bool <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="201:1:1" line-data="	AutodetectProfile     bool">`AutodetectProfile`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +time.Duration <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="91:1:1" line-data="	MinCollectionInterval        int                               `yaml:&quot;min_collection_interval&quot;`">`MinCollectionInterval`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +bool <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="94:1:1" line-data="	DetectMetricsEnabled         Boolean                           `yaml:&quot;experimental_detect_metrics_enabled&quot;`">`DetectMetricsEnabled`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +int <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="95:1:1" line-data="	DetectMetricsRefreshInterval int                               `yaml:&quot;experimental_detect_metrics_refresh_interval&quot;`">`DetectMetricsRefreshInterval`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +string Network
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +int <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="147:1:1" line-data="	DiscoveryWorkers         int      `yaml:&quot;discovery_workers&quot;`">`DiscoveryWorkers`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +int Workers
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +int <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="145:1:1" line-data="	DiscoveryInterval        int      `yaml:&quot;discovery_interval&quot;`">`DiscoveryInterval`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +map <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="144:1:1" line-data="	IgnoredIPAddresses       []string `yaml:&quot;ignored_ip_addresses&quot;`">`IgnoredIPAddresses`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +int <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="146:1:1" line-data="	DiscoveryAllowedFailures int      `yaml:&quot;discovery_allowed_failures&quot;`">`DiscoveryAllowedFailures`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +[]InterfaceConfig <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="158:1:1" line-data="	InterfaceConfigs InterfaceConfigs `yaml:&quot;interface_configs&quot;`">`InterfaceConfigs`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +bool <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="215:1:1" line-data="	PingEnabled bool">`PingEnabled`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> : +Config <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="93:1:1" line-data="	PingConfig                   snmpintegration.PackedPingConfig  `yaml:&quot;ping&quot;`">`PingConfig`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +ProfileConfigMap Profiles
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +[]MetricsConfig <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="85:1:1" line-data="	GlobalMetrics                []profiledefinition.MetricsConfig `yaml:&quot;global_metrics&quot;`">`GlobalMetrics`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +Number <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="86:1:1" line-data="	OidBatchSize                 Number                            `yaml:&quot;oid_batch_size&quot;`">`OidBatchSize`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +Number <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="87:1:1" line-data="	BulkMaxRepetitions           Number                            `yaml:&quot;bulk_max_repetitions&quot;`">`BulkMaxRepetitions`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +Boolean <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="88:1:1" line-data="	CollectDeviceMetadata        Boolean                           `yaml:&quot;collect_device_metadata&quot;`">`CollectDeviceMetadata`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +Boolean <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="89:1:1" line-data="	CollectTopology              Boolean                           `yaml:&quot;collect_topology&quot;`">`CollectTopology`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +Boolean <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="90:1:1" line-data="	UseDeviceIDAsHostname        Boolean                           `yaml:&quot;use_device_id_as_hostname&quot;`">`UseDeviceIDAsHostname`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +int <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="91:1:1" line-data="	MinCollectionInterval        int                               `yaml:&quot;min_collection_interval&quot;`">`MinCollectionInterval`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +string Namespace
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +PackedPingConfig <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="93:1:1" line-data="	PingConfig                   snmpintegration.PackedPingConfig  `yaml:&quot;ping&quot;`">`PingConfig`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +Boolean <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="94:1:1" line-data="	DetectMetricsEnabled         Boolean                           `yaml:&quot;experimental_detect_metrics_enabled&quot;`">`DetectMetricsEnabled`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> : +int <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="95:1:1" line-data="	DetectMetricsRefreshInterval int                               `yaml:&quot;experimental_detect_metrics_refresh_interval&quot;`">`DetectMetricsRefreshInterval`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string Name
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="101:1:1" line-data="	IPAddress             string                              `yaml:&quot;ip_address&quot;`">`IPAddress`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +Number Port
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="103:1:1" line-data="	CommunityString       string                              `yaml:&quot;community_string&quot;`">`CommunityString`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="104:1:1" line-data="	SnmpVersion           string                              `yaml:&quot;snmp_version&quot;`">`SnmpVersion`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +Number Timeout
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +Number Retries
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string User
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="108:1:1" line-data="	AuthProtocol          string                              `yaml:&quot;authProtocol&quot;`">`AuthProtocol`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="109:1:1" line-data="	AuthKey               string                              `yaml:&quot;authKey&quot;`">`AuthKey`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="110:1:1" line-data="	PrivProtocol          string                              `yaml:&quot;privProtocol&quot;`">`PrivProtocol`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="111:1:1" line-data="	PrivKey               string                              `yaml:&quot;privKey&quot;`">`PrivKey`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="112:1:1" line-data="	ContextName           string                              `yaml:&quot;context_name&quot;`">`ContextName`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +[]MetricsConfig Metrics
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +[]MetricTagConfig <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="114:1:1" line-data="	MetricTags            []profiledefinition.MetricTagConfig `yaml:&quot;metric_tags&quot;` // SNMP metric tags definition">`MetricTags`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string Profile
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +bool <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="116:1:1" line-data="	UseGlobalMetrics      bool                                `yaml:&quot;use_global_metrics&quot;`">`UseGlobalMetrics`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +Boolean <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="88:1:1" line-data="	CollectDeviceMetadata        Boolean                           `yaml:&quot;collect_device_metadata&quot;`">`CollectDeviceMetadata`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +Boolean <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="89:1:1" line-data="	CollectTopology              Boolean                           `yaml:&quot;collect_topology&quot;`">`CollectTopology`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +Boolean <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="90:1:1" line-data="	UseDeviceIDAsHostname        Boolean                           `yaml:&quot;use_device_id_as_hostname&quot;`">`UseDeviceIDAsHostname`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +PackedPingConfig <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="93:1:1" line-data="	PingConfig                   snmpintegration.PackedPingConfig  `yaml:&quot;ping&quot;`">`PingConfig`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> : +string <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="122:3:3" line-data="	// ExtraTags is a workaround to pass tags from snmp listener to snmp integration via AD template">`ExtraTags`</SwmToken>
```

# Overview

This document provides an in-depth look at the SNMP check configuration in the Datadog Agent. The configuration is essential for the SNMP check to run properly and includes various parameters such as IP address, port, SNMP version, and authentication details.

# Configuration Structure

The configuration is divided into two main parts: <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> and <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken>. <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> contains global settings like profiles, global metrics, and batch sizes, while <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> holds instance-specific settings such as IP address, port, and community string.

<SwmSnippet path="/pkg/collector/corechecks/snmp/internal/checkconfig/config.go" line="82">

---

<SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> is used to deserialize integration init config, containing global settings like profiles, global metrics, and batch sizes.

```go
// InitConfig is used to deserialize integration init config
type InitConfig struct {
	Profiles                     profile.ProfileConfigMap          `yaml:"profiles"`
	GlobalMetrics                []profiledefinition.MetricsConfig `yaml:"global_metrics"`
	OidBatchSize                 Number                            `yaml:"oid_batch_size"`
	BulkMaxRepetitions           Number                            `yaml:"bulk_max_repetitions"`
	CollectDeviceMetadata        Boolean                           `yaml:"collect_device_metadata"`
	CollectTopology              Boolean                           `yaml:"collect_topology"`
	UseDeviceIDAsHostname        Boolean                           `yaml:"use_device_id_as_hostname"`
	MinCollectionInterval        int                               `yaml:"min_collection_interval"`
	Namespace                    string                            `yaml:"namespace"`
	PingConfig                   snmpintegration.PackedPingConfig  `yaml:"ping"`
	DetectMetricsEnabled         Boolean                           `yaml:"experimental_detect_metrics_enabled"`
	DetectMetricsRefreshInterval int                               `yaml:"experimental_detect_metrics_refresh_interval"`
}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/collector/corechecks/snmp/internal/checkconfig/config.go" line="98">

---

<SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> is used to deserialize integration instance config, holding instance-specific settings such as IP address, port, and community string.

```go
// InstanceConfig is used to deserialize integration instance config
type InstanceConfig struct {
	Name                  string                              `yaml:"name"`
	IPAddress             string                              `yaml:"ip_address"`
	Port                  Number                              `yaml:"port"`
	CommunityString       string                              `yaml:"community_string"`
	SnmpVersion           string                              `yaml:"snmp_version"`
	Timeout               Number                              `yaml:"timeout"`
	Retries               Number                              `yaml:"retries"`
	User                  string                              `yaml:"user"`
	AuthProtocol          string                              `yaml:"authProtocol"`
	AuthKey               string                              `yaml:"authKey"`
	PrivProtocol          string                              `yaml:"privProtocol"`
	PrivKey               string                              `yaml:"privKey"`
	ContextName           string                              `yaml:"context_name"`
	Metrics               []profiledefinition.MetricsConfig   `yaml:"metrics"`     // SNMP metrics definition
	MetricTags            []profiledefinition.MetricTagConfig `yaml:"metric_tags"` // SNMP metric tags definition
	Profile               string                              `yaml:"profile"`
	UseGlobalMetrics      bool                                `yaml:"use_global_metrics"`
	CollectDeviceMetadata *Boolean                            `yaml:"collect_device_metadata"`
	CollectTopology       *Boolean                            `yaml:"collect_topology"`
```

---

</SwmSnippet>

# <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> Struct

The <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> struct consolidates these configurations and includes additional settings like requested metrics, metric tags, and device metadata collection options. It ensures that all necessary parameters are available for the SNMP check to execute.

<SwmSnippet path="/pkg/collector/corechecks/snmp/internal/checkconfig/config.go" line="161">

---

<SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> consolidates <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="82:2:2" line-data="// InitConfig is used to deserialize integration init config">`InitConfig`</SwmToken> and <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="98:2:2" line-data="// InstanceConfig is used to deserialize integration instance config">`InstanceConfig`</SwmToken> and includes additional settings like requested metrics, metric tags, and device metadata collection options.

```go
// CheckConfig holds config needed for an integration instance to run
type CheckConfig struct {
	Name            string
	IPAddress       string
	Port            uint16
	CommunityString string
	SnmpVersion     string
	Timeout         int
	Retries         int
	User            string
	AuthProtocol    string
	AuthKey         string
	PrivProtocol    string
	PrivKey         string
	ContextName     string
	OidConfig       OidConfig
	// RequestedMetrics are the metrics explicitly requested by config.
	RequestedMetrics []profiledefinition.MetricsConfig
	// RequestedMetricTags are the tags explicitly requested by config.
	RequestedMetricTags []profiledefinition.MetricTagConfig
	// Metrics combines RequestedMetrics with profile metrics.
```

---

</SwmSnippet>

# <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="344:2:2" line-data="// NewCheckConfig builds a new check config">`NewCheckConfig`</SwmToken> Function

The <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="344:2:2" line-data="// NewCheckConfig builds a new check config">`NewCheckConfig`</SwmToken> function is responsible for creating a new <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="161:2:2" line-data="// CheckConfig holds config needed for an integration instance to run">`CheckConfig`</SwmToken> instance by unmarshalling raw configuration data and setting default values where necessary. This function validates the configuration and ensures that either an IP address or a network is provided.

<SwmSnippet path="/pkg/collector/corechecks/snmp/internal/checkconfig/config.go" line="344">

---

The <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="344:2:2" line-data="// NewCheckConfig builds a new check config">`NewCheckConfig`</SwmToken> function creates a new <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="345:20:20" line-data="func NewCheckConfig(rawInstance integration.Data, rawInitConfig integration.Data) (*CheckConfig, error) {">`CheckConfig`</SwmToken> instance by unmarshalling raw configuration data and setting default values where necessary.

```go
// NewCheckConfig builds a new check config
func NewCheckConfig(rawInstance integration.Data, rawInitConfig integration.Data) (*CheckConfig, error) {
	instance := InstanceConfig{}
	initConfig := InitConfig{}

	// Set defaults before unmarshalling
	instance.UseGlobalMetrics = true
	initConfig.CollectDeviceMetadata = true
	initConfig.CollectTopology = true

	err := yaml.Unmarshal(rawInitConfig, &initConfig)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(rawInstance, &instance)
	if err != nil {
		return nil, err
	}

	c := &CheckConfig{}
```

---

</SwmSnippet>

# <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="219:2:2" line-data="// SetProfile refreshes config based on profile">`SetProfile`</SwmToken> Method

The <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="219:2:2" line-data="// SetProfile refreshes config based on profile">`SetProfile`</SwmToken> method refreshes the configuration based on the provided profile. It updates the profile definition and tags, ensuring that the configuration is aligned with the selected profile.

<SwmSnippet path="/pkg/collector/corechecks/snmp/internal/checkconfig/config.go" line="219">

---

The <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="219:2:2" line-data="// SetProfile refreshes config based on profile">`SetProfile`</SwmToken> method refreshes the configuration based on the provided profile.

```go
// SetProfile refreshes config based on profile
func (c *CheckConfig) SetProfile(profile string) error {
	if _, ok := c.Profiles[profile]; !ok {
		return fmt.Errorf("unknown profile `%s`", profile)
	}
	log.Debugf("Refreshing with profile `%s`", profile)
	tags := []string{"snmp_profile:" + profile}
	definition := c.Profiles[profile].Definition
	c.ProfileDef = &definition
	c.Profile = profile

	if log.ShouldLog(seelog.DebugLvl) {
		profileDefJSON, _ := json.Marshal(definition)
		log.Debugf("Profile content `%s`: %s", profile, string(profileDefJSON))
	}

	if definition.Device.Vendor != "" {
		tags = append(tags, "device_vendor:"+definition.Device.Vendor)
	}
	tags = append(tags, definition.StaticTags...)
	c.ProfileTags = tags
```

---

</SwmSnippet>

# <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="257:2:2" line-data="// RebuildMetadataMetricsAndTags rebuilds c.Metrics, c.Metadata, c.MetricTags,">`RebuildMetadataMetricsAndTags`</SwmToken> Method

The <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="257:2:2" line-data="// RebuildMetadataMetricsAndTags rebuilds c.Metrics, c.Metadata, c.MetricTags,">`RebuildMetadataMetricsAndTags`</SwmToken> method rebuilds the metrics, metadata, and metric tags by merging data from requested <SwmPath>[comp/core/tagger/tags/](comp/core/tagger/tags/)</SwmPath> and the current profile. This ensures that the configuration is up-to-date with the latest profile settings.

<SwmSnippet path="/pkg/collector/corechecks/snmp/internal/checkconfig/config.go" line="257">

---

The <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="257:2:2" line-data="// RebuildMetadataMetricsAndTags rebuilds c.Metrics, c.Metadata, c.MetricTags,">`RebuildMetadataMetricsAndTags`</SwmToken> method rebuilds the metrics, metadata, and metric tags by merging data from requested <SwmPath>[comp/core/tagger/tags/](comp/core/tagger/tags/)</SwmPath> and the current profile.

```go
// RebuildMetadataMetricsAndTags rebuilds c.Metrics, c.Metadata, c.MetricTags,
// and c.OidConfig by merging data from requested metrics/tags and the current
// profile.
func (c *CheckConfig) RebuildMetadataMetricsAndTags() {
	c.Metrics = c.RequestedMetrics
	c.MetricTags = c.RequestedMetricTags
	if c.ProfileDef != nil {
		c.Metadata = updateMetadataDefinitionWithDefaults(c.ProfileDef.Metadata, c.CollectTopology)
		c.Metrics = append(c.Metrics, c.ProfileDef.Metrics...)
		c.MetricTags = append(c.MetricTags, c.ProfileDef.MetricTags...)
	} else {
		c.Metadata = updateMetadataDefinitionWithDefaults(nil, c.CollectTopology)
	}
	c.OidConfig.clean()
	c.OidConfig.addScalarOids(c.parseScalarOids(c.Metrics, c.MetricTags, c.Metadata))
	c.OidConfig.addColumnOids(c.parseColumnOids(c.Metrics, c.Metadata))
}
```

---

</SwmSnippet>

# <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="275:2:2" line-data="// UpdateDeviceIDAndTags updates DeviceID and DeviceIDTags">`UpdateDeviceIDAndTags`</SwmToken> Method

The <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="275:2:2" line-data="// UpdateDeviceIDAndTags updates DeviceID and DeviceIDTags">`UpdateDeviceIDAndTags`</SwmToken> method updates the device ID and device ID tags. It sorts and ensures the uniqueness of the device ID tags, which are used for generating the device ID.

<SwmSnippet path="/pkg/collector/corechecks/snmp/internal/checkconfig/config.go" line="275">

---

The <SwmToken path="pkg/collector/corechecks/snmp/internal/checkconfig/config.go" pos="275:2:2" line-data="// UpdateDeviceIDAndTags updates DeviceID and DeviceIDTags">`UpdateDeviceIDAndTags`</SwmToken> method updates the device ID and device ID tags.

```go
// UpdateDeviceIDAndTags updates DeviceID and DeviceIDTags
func (c *CheckConfig) UpdateDeviceIDAndTags() {
	c.DeviceIDTags = coreutil.SortUniqInPlace(c.getDeviceIDTags())
	c.DeviceID = c.Namespace + ":" + c.IPAddress
}
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
