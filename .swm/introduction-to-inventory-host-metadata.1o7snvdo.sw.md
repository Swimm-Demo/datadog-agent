---
title: Introduction to Inventory Host Metadata
---
```mermaid
classDiagram
 inventoryhost --|> hostMetadata
 inventoryhost : +string Hostname
 inventoryhost : +int64 Timestamp
 inventoryhost : +hostMetadata Metadata
 inventoryhost : +string UUID
 class hostMetadata{
 +uint64 CPUCores
 +uint64 CPULogicalProcessors
 +string CPUVendor
 +string CPUModel
 +string CPUModelID
 +string CPUFamily
 +string CPUStepping
 +float64 CPUFrequency
 +uint64 CPUCacheSize
 +string KernelName
 +string KernelRelease
 +string KernelVersion
 +string OS
 +string CPUArchitecture
 +uint64 MemoryTotalKb
 +uint64 MemorySwapTotalKb
 +string IPAddress
 +string IPv6Address
 +string MacAddress
 +string AgentVersion
 +string CloudProvider
 +string CloudProviderSource
 +string CloudProviderAccountID
 +string CloudProviderHostID
 +string osversion
 +string HypervisorGuestUUID
 +string DmiProductUUID
 +string DmiBoardAssetTag
 +string DmiBoardVendor
 +bool LinuxPackageSigningEnabled
 +bool RPMGlobalRepoGPGCheckEnabled
 }
graph TD;
 A[HTTP Request] -->|calls| B[writePayloadAsJSON];
 B -->|calls| C[getPayload];
 C -->|fills| D[Inventory Host Data];
 D -->|writes| E[HTTP Response];

%% Swimm:
%% classDiagram
%%  <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="24:16:16" line-data="	&quot;github.com/DataDog/datadog-agent/comp/metadata/inventoryhost&quot;">`inventoryhost`</SwmToken> --|> <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="58:2:2" line-data="// hostMetadata contains metadata about the host">`hostMetadata`</SwmToken>
%%  <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="24:16:16" line-data="	&quot;github.com/DataDog/datadog-agent/comp/metadata/inventoryhost&quot;">`inventoryhost`</SwmToken> : +string Hostname
%%  <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="24:16:16" line-data="	&quot;github.com/DataDog/datadog-agent/comp/metadata/inventoryhost&quot;">`inventoryhost`</SwmToken> : +int64 Timestamp
%%  <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="24:16:16" line-data="	&quot;github.com/DataDog/datadog-agent/comp/metadata/inventoryhost&quot;">`inventoryhost`</SwmToken> : +hostMetadata Metadata
%%  <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="24:16:16" line-data="	&quot;github.com/DataDog/datadog-agent/comp/metadata/inventoryhost&quot;">`inventoryhost`</SwmToken> : +string UUID
%%  class <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="159:5:6" line-data="		data:     &amp;hostMetadata{},">`hostMetadata{`</SwmToken>
%%  +uint64 <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="61:1:1" line-data="	CPUCores             uint64  `json:&quot;cpu_cores&quot;`">`CPUCores`</SwmToken>
%%  +uint64 <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="62:1:1" line-data="	CPULogicalProcessors uint64  `json:&quot;cpu_logical_processors&quot;`">`CPULogicalProcessors`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="63:1:1" line-data="	CPUVendor            string  `json:&quot;cpu_vendor&quot;`">`CPUVendor`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="64:1:1" line-data="	CPUModel             string  `json:&quot;cpu_model&quot;`">`CPUModel`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="65:1:1" line-data="	CPUModelID           string  `json:&quot;cpu_model_id&quot;`">`CPUModelID`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="66:1:1" line-data="	CPUFamily            string  `json:&quot;cpu_family&quot;`">`CPUFamily`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="67:1:1" line-data="	CPUStepping          string  `json:&quot;cpu_stepping&quot;`">`CPUStepping`</SwmToken>
%%  +float64 <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="68:1:1" line-data="	CPUFrequency         float64 `json:&quot;cpu_frequency&quot;`">`CPUFrequency`</SwmToken>
%%  +uint64 <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="69:1:1" line-data="	CPUCacheSize         uint64  `json:&quot;cpu_cache_size&quot;`">`CPUCacheSize`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="72:1:1" line-data="	KernelName      string `json:&quot;kernel_name&quot;`">`KernelName`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="73:1:1" line-data="	KernelRelease   string `json:&quot;kernel_release&quot;`">`KernelRelease`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="74:1:1" line-data="	KernelVersion   string `json:&quot;kernel_version&quot;`">`KernelVersion`</SwmToken>
%%  +string OS
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="76:1:1" line-data="	CPUArchitecture string `json:&quot;cpu_architecture&quot;`">`CPUArchitecture`</SwmToken>
%%  +uint64 <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="79:1:1" line-data="	MemoryTotalKb     uint64 `json:&quot;memory_total_kb&quot;`">`MemoryTotalKb`</SwmToken>
%%  +uint64 <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="80:1:1" line-data="	MemorySwapTotalKb uint64 `json:&quot;memory_swap_total_kb&quot;`">`MemorySwapTotalKb`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="83:1:1" line-data="	IPAddress   string `json:&quot;ip_address&quot;`">`IPAddress`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="84:1:1" line-data="	IPv6Address string `json:&quot;ipv6_address&quot;`">`IPv6Address`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="85:1:1" line-data="	MacAddress  string `json:&quot;mac_address&quot;`">`MacAddress`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="88:1:1" line-data="	AgentVersion           string `json:&quot;agent_version&quot;`">`AgentVersion`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="89:1:1" line-data="	CloudProvider          string `json:&quot;cloud_provider&quot;`">`CloudProvider`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="90:1:1" line-data="	CloudProviderSource    string `json:&quot;cloud_provider_source&quot;`">`CloudProviderSource`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="91:1:1" line-data="	CloudProviderAccountID string `json:&quot;cloud_provider_account_id&quot;`">`CloudProviderAccountID`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="92:1:1" line-data="	CloudProviderHostID    string `json:&quot;cloud_provider_host_id&quot;`">`CloudProviderHostID`</SwmToken>
%%  +string <SwmToken path="tasks/new_e2e_tests.py" pos="48:1:1" line-data="    osversion=&quot;&quot;,">`osversion`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="96:1:1" line-data="	HypervisorGuestUUID string `json:&quot;hypervisor_guest_uuid&quot;`">`HypervisorGuestUUID`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="97:1:1" line-data="	DmiProductUUID      string `json:&quot;dmi_product_uuid&quot;`">`DmiProductUUID`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="98:1:1" line-data="	DmiBoardAssetTag    string `json:&quot;dmi_board_asset_tag&quot;`">`DmiBoardAssetTag`</SwmToken>
%%  +string <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="99:1:1" line-data="	DmiBoardVendor      string `json:&quot;dmi_board_vendor&quot;`">`DmiBoardVendor`</SwmToken>
%%  +bool <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="102:1:1" line-data="	LinuxPackageSigningEnabled   bool `json:&quot;linux_package_signing_enabled&quot;`">`LinuxPackageSigningEnabled`</SwmToken>
%%  +bool <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="103:1:1" line-data="	RPMGlobalRepoGPGCheckEnabled bool `json:&quot;rpm_global_repo_gpg_check_enabled&quot;`">`RPMGlobalRepoGPGCheckEnabled`</SwmToken>
%%  }
%% graph TD;
%%  A[HTTP Request] -->|calls| B[writePayloadAsJSON];
%%  B -->|calls| C[getPayload];
%%  C -->|fills| D[Inventory Host Data];
%%  D -->|writes| E[HTTP Response];
```

# Introduction

Inventory Host refers to the collection of metadata about the host system where the Datadog Agent is running. This metadata includes details about the CPU, memory, network, operating system, and other system attributes.

# Metadata Collection

The Inventory Host data is populated into the <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="124:17:17" line-data="	return nil, fmt.Errorf(&quot;could not split inventories host payload any more, payload is too big for intake&quot;)">`inventories`</SwmToken> product in Datadog, specifically into the `host_agent` table. This process is enabled by default but can be turned off using the `inventories_enabled` configuration.

# Metadata Details

The metadata collected includes information such as CPU cores, logical processors, vendor, model, kernel details, total memory, IP addresses, and more. This data is essential for monitoring and analyzing the performance and health of the host system.

# Data Transmission

The payload containing this metadata is sent to Datadog every 10 minutes or whenever it is updated, with a minimum interval of 5 minutes between updates. This ensures that the data remains current and accurate.

<SwmSnippet path="/comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" line="58">

---

The <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="58:2:2" line-data="// hostMetadata contains metadata about the host">`hostMetadata`</SwmToken> struct defines the various fields that are collected as part of the Inventory Host data. These fields include CPU details, memory information, network addresses, agent version, cloud provider details, and more.

```go
// hostMetadata contains metadata about the host
type hostMetadata struct {
	// from gohai/cpu
	CPUCores             uint64  `json:"cpu_cores"`
	CPULogicalProcessors uint64  `json:"cpu_logical_processors"`
	CPUVendor            string  `json:"cpu_vendor"`
	CPUModel             string  `json:"cpu_model"`
	CPUModelID           string  `json:"cpu_model_id"`
	CPUFamily            string  `json:"cpu_family"`
	CPUStepping          string  `json:"cpu_stepping"`
	CPUFrequency         float64 `json:"cpu_frequency"`
	CPUCacheSize         uint64  `json:"cpu_cache_size"`

	// from gohai/platform
	KernelName      string `json:"kernel_name"`
	KernelRelease   string `json:"kernel_release"`
	KernelVersion   string `json:"kernel_version"`
	OS              string `json:"os"`
	CPUArchitecture string `json:"cpu_architecture"`

	// from gohai/memory
```

---

</SwmSnippet>

# Example Payload

Here is an example of an inventory payload: {"host_metadata":{"cpu_architecture": "unknown", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="69:9:9" line-data="	CPUCacheSize         uint64  `json:&quot;cpu_cache_size&quot;`">`cpu_cache_size`</SwmToken>: 9437184, <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="61:9:9" line-data="	CPUCores             uint64  `json:&quot;cpu_cores&quot;`">`cpu_cores`</SwmToken>: 6, <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="66:9:9" line-data="	CPUFamily            string  `json:&quot;cpu_family&quot;`">`cpu_family`</SwmToken>: "6", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="68:9:9" line-data="	CPUFrequency         float64 `json:&quot;cpu_frequency&quot;`">`cpu_frequency`</SwmToken>: 2208.007, <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="62:9:9" line-data="	CPULogicalProcessors uint64  `json:&quot;cpu_logical_processors&quot;`">`cpu_logical_processors`</SwmToken>: 6, <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="64:9:9" line-data="	CPUModel             string  `json:&quot;cpu_model&quot;`">`cpu_model`</SwmToken>: "Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="65:9:9" line-data="	CPUModelID           string  `json:&quot;cpu_model_id&quot;`">`cpu_model_id`</SwmToken>: "158", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="67:9:9" line-data="	CPUStepping          string  `json:&quot;cpu_stepping&quot;`">`cpu_stepping`</SwmToken>: "10", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="63:9:9" line-data="	CPUVendor            string  `json:&quot;cpu_vendor&quot;`">`cpu_vendor`</SwmToken>: "GenuineIntel", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="72:9:9" line-data="	KernelName      string `json:&quot;kernel_name&quot;`">`kernel_name`</SwmToken>: "Linux", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="73:9:9" line-data="	KernelRelease   string `json:&quot;kernel_release&quot;`">`kernel_release`</SwmToken>: "5.16.0-6-amd64", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="74:9:9" line-data="	KernelVersion   string `json:&quot;kernel_version&quot;`">`kernel_version`</SwmToken>: "#1 SMP PREEMPT Debian 5.16.18-1 (2022-03-29)", "os": "GNU/Linux", <SwmToken path="tasks/kernel_matrix_testing/types.py" pos="26:1:1" line-data="    os_version: str  # Version  # noqa: F841">`os_version`</SwmToken>: "debian bookworm/sid", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="80:9:9" line-data="	MemorySwapTotalKb uint64 `json:&quot;memory_swap_total_kb&quot;`">`memory_swap_total_kb`</SwmToken>: 10237948, <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="79:9:9" line-data="	MemoryTotalKb     uint64 `json:&quot;memory_total_kb&quot;`">`memory_total_kb`</SwmToken>: 12227556, <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="83:9:9" line-data="	IPAddress   string `json:&quot;ip_address&quot;`">`ip_address`</SwmToken>: "192.168.24.138", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="84:9:9" line-data="	IPv6Address string `json:&quot;ipv6_address&quot;`">`ipv6_address`</SwmToken>: "fe80::1ff:fe23:4567:890a", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="85:9:9" line-data="	MacAddress  string `json:&quot;mac_address&quot;`">`mac_address`</SwmToken>: "01:23:45:67:89:AB", <SwmToken path="pkg/config/legacy/tests/config.py" pos="29:0:0" line-data="AGENT_VERSION = &quot;5.18.0&quot;">`AGENT_VERSION`</SwmToken>: "7.37.0-devel+git.198.68a5b69", <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="89:9:9" line-data="	CloudProvider          string `json:&quot;cloud_provider&quot;`">`cloud_provider`</SwmToken>: "AWS" }

# Main Functions

There are several main functions in this folder. Some of them are <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="171:9:9" line-data="func (ih *invHost) fillData() {">`fillData`</SwmToken> and <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="254:9:9" line-data="func (ih *invHost) getPayload() marshaler.JSONMarshaler {">`getPayload`</SwmToken>. We will dive a little into <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="171:9:9" line-data="func (ih *invHost) fillData() {">`fillData`</SwmToken> and <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="254:9:9" line-data="func (ih *invHost) getPayload() marshaler.JSONMarshaler {">`getPayload`</SwmToken>.

## <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="171:9:9" line-data="func (ih *invHost) fillData() {">`fillData`</SwmToken>

The <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="171:9:9" line-data="func (ih *invHost) fillData() {">`fillData`</SwmToken> function is responsible for collecting various pieces of metadata about the host system, such as CPU information, platform details, memory statistics, and network information. This data is then stored in the <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="58:2:2" line-data="// hostMetadata contains metadata about the host">`hostMetadata`</SwmToken> struct.

<SwmSnippet path="/comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" line="171">

---

The <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="171:9:9" line-data="func (ih *invHost) fillData() {">`fillData`</SwmToken> function collects various pieces of metadata about the host system and stores them in the <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="58:2:2" line-data="// hostMetadata contains metadata about the host">`hostMetadata`</SwmToken> struct.

```go
func (ih *invHost) fillData() {
	logWarnings := func(warnings []string) {
		for _, w := range warnings {
			ih.log.Infof("gohai: %s", w)
		}
	}

	cpuInfo := cpuGet()
	_, warnings, err := cpuInfo.AsJSON()
	if err != nil {
		ih.log.Errorf("Failed to retrieve cpu metadata from gohai: %s", err) //nolint:errcheck
	} else {
		logWarnings(warnings)

		ih.data.CPUCores = cpuInfo.CPUCores.ValueOrDefault()
		ih.data.CPULogicalProcessors = cpuInfo.CPULogicalProcessors.ValueOrDefault()
		ih.data.CPUVendor = cpuInfo.VendorID.ValueOrDefault()
		ih.data.CPUModel = cpuInfo.ModelName.ValueOrDefault()
		ih.data.CPUModelID = cpuInfo.Model.ValueOrDefault()
		ih.data.CPUFamily = cpuInfo.Family.ValueOrDefault()
		ih.data.CPUStepping = cpuInfo.Stepping.ValueOrDefault()
```

---

</SwmSnippet>

## <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="254:9:9" line-data="func (ih *invHost) getPayload() marshaler.JSONMarshaler {">`getPayload`</SwmToken>

The <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="254:9:9" line-data="func (ih *invHost) getPayload() marshaler.JSONMarshaler {">`getPayload`</SwmToken> function calls <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="171:9:9" line-data="func (ih *invHost) fillData() {">`fillData`</SwmToken> to populate the metadata and then returns it as a JSON marshaler. This function is crucial for generating the payload that will be sent to Datadog.

<SwmSnippet path="/comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" line="254">

---

The <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="254:9:9" line-data="func (ih *invHost) getPayload() marshaler.JSONMarshaler {">`getPayload`</SwmToken> function calls <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="255:3:3" line-data="	ih.fillData()">`fillData`</SwmToken> to populate the metadata and then returns it as a JSON marshaler.

```go
func (ih *invHost) getPayload() marshaler.JSONMarshaler {
	ih.fillData()
```

---

</SwmSnippet>

# Inventory Host Endpoints

The Inventory Host Endpoints provide HTTP handlers for interacting with the Inventory Host data.

## <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="265:9:9" line-data="func (ih *invHost) writePayloadAsJSON(w http.ResponseWriter, _ *http.Request) {">`writePayloadAsJSON`</SwmToken>

The <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="265:9:9" line-data="func (ih *invHost) writePayloadAsJSON(w http.ResponseWriter, _ *http.Request) {">`writePayloadAsJSON`</SwmToken> function is an HTTP handler that writes the Inventory Host payload as a JSON response. It calls the <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="254:9:9" line-data="func (ih *invHost) getPayload() marshaler.JSONMarshaler {">`getPayload`</SwmToken> method to fill the data and then writes it to the HTTP response writer.

<SwmSnippet path="/comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" line="265">

---

The <SwmToken path="comp/metadata/inventoryhost/inventoryhostimpl/inventoryhost.go" pos="265:9:9" line-data="func (ih *invHost) writePayloadAsJSON(w http.ResponseWriter, _ *http.Request) {">`writePayloadAsJSON`</SwmToken> function writes the Inventory Host payload as a JSON response.

```go
func (ih *invHost) writePayloadAsJSON(w http.ResponseWriter, _ *http.Request) {
	// GetAsJSON already return scrubbed data
	scrubbed, err := ih.GetAsJSON()
	if err != nil {
		httputils.SetJSONError(w, err, 500)
		return
	}
	w.Write(scrubbed)
}
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
