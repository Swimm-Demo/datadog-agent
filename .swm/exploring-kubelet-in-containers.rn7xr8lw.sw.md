---
title: Exploring Kubelet in Containers
---
```mermaid
classDiagram
 KubeletCheck --> Provider
 KubeletCheck : +core.CheckBase CheckBase
 KubeletCheck : +common.KubeletConfig instance
 KubeletCheck : +containers.Filter filter
 KubeletCheck : +Provider[] providers
 KubeletCheck : +common.PodUtils podUtils
 KubeletCheck : +workloadmeta.Component store
 Provider : +Provide(kubelet.KubeUtilInterface, sender.Sender) error
graph TD;
 A[Kubelet] -->|/stats/summary| B[Summary Provider];
 A -->|/metrics/probes| C[Probes Provider];
 B --> D[Process System Stats];
 B --> E[Process Pod Stats];
 B --> F[Process Container Stats];
 C --> G[Process Probe Metrics];

%% Swimm:
%% classDiagram
%%  <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> --> Provider
%%  <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> : +core.CheckBase <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="44:3:3" line-data="	core.CheckBase">`CheckBase`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> : +common.KubeletConfig instance
%%  <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> : +containers.Filter filter
%%  <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> : +Provider[] providers
%%  <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> : +common.PodUtils <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="48:1:1" line-data="	podUtils  *common.PodUtils">`podUtils`</SwmToken>
%%  <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> : +workloadmeta.Component store
%%  Provider : +Provide(kubelet.KubeUtilInterface, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="39:8:10" line-data="	Provide(kubelet.KubeUtilInterface, sender.Sender) error">`sender.Sender`</SwmToken>) error
%% graph TD;
%%  A[Kubelet] -->|/stats/summary| B[Summary Provider];
%%  A -->|/metrics/probes| C[Probes Provider];
%%  B --> D[Process System Stats];
%%  B --> E[Process Pod Stats];
%%  B --> F[Process Container Stats];
%%  C --> G[Process Probe Metrics];
```

# Overview

Kubelet is a core component responsible for managing the lifecycle of containers on a Kubernetes node. It ensures that the containers are running as expected and reports their status to the Kubernetes control plane.

# <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> Struct

In the Datadog Agent, the <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> struct is used to wrap the configuration and metric stores needed to perform checks on the Kubelet. This struct includes members such as <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="45:1:1" line-data="	instance  *common.KubeletConfig">`instance`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="46:1:1" line-data="	filter    *containers.Filter">`filter`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="47:1:1" line-data="	providers []Provider">`providers`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="48:1:1" line-data="	podUtils  *common.PodUtils">`podUtils`</SwmToken>, and <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="49:1:1" line-data="	store     workloadmeta.Component">`store`</SwmToken>.

<SwmSnippet path="/pkg/collector/corechecks/containers/kubelet/kubelet.go" line="42">

---

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> struct wraps the configuration and metric stores needed to run the check. It includes members such as <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="45:1:1" line-data="	instance  *common.KubeletConfig">`instance`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="46:1:1" line-data="	filter    *containers.Filter">`filter`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="47:1:1" line-data="	providers []Provider">`providers`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="48:1:1" line-data="	podUtils  *common.PodUtils">`podUtils`</SwmToken>, and <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="49:1:1" line-data="	store     workloadmeta.Component">`store`</SwmToken>.

```go
// KubeletCheck wraps the config and the metric stores needed to run the check
type KubeletCheck struct {
	core.CheckBase
	instance  *common.KubeletConfig
	filter    *containers.Filter
	providers []Provider
	podUtils  *common.PodUtils
	store     workloadmeta.Component
}
```

---

</SwmSnippet>

# Initialization

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> struct is initialized using the <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="52:2:2" line-data="// NewKubeletCheck returns a new KubeletCheck">`NewKubeletCheck`</SwmToken> function, which sets up the base check, instance configuration, and metric store. This setup is essential for collecting and processing metrics from the Kubelet.

<SwmSnippet path="/pkg/collector/corechecks/containers/kubelet/kubelet.go" line="52">

---

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="52:2:2" line-data="// NewKubeletCheck returns a new KubeletCheck">`NewKubeletCheck`</SwmToken> function initializes the <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="52:10:10" line-data="// NewKubeletCheck returns a new KubeletCheck">`KubeletCheck`</SwmToken> struct by setting up the base check, instance configuration, and metric store.

```go
// NewKubeletCheck returns a new KubeletCheck
func NewKubeletCheck(base core.CheckBase, instance *common.KubeletConfig, store workloadmeta.Component) *KubeletCheck {
	return &KubeletCheck{
		CheckBase: base,
		instance:  instance,
		store:     store,
	}
}
```

---

</SwmSnippet>

# Providers Initialization

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="61:2:2" line-data="func initProviders(filter *containers.Filter, config *common.KubeletConfig, podUtils *common.PodUtils, store workloadmeta.Component) []Provider {">`initProviders`</SwmToken> function within the <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> struct initializes various providers that collect metrics from different Kubelet endpoints. These providers include <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="62:1:1" line-data="	podProvider := pod.NewProvider(filter, config, podUtils)">`podProvider`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="63:3:3" line-data="	// nodeProvider collects from the /spec endpoint, which was hidden by default in k8s 1.18 and removed in k8s 1.19.">`nodeProvider`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="66:1:1" line-data="	healthProvider := health.NewProvider(config)">`healthProvider`</SwmToken>, and <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="67:1:1" line-data="	summaryProvider := summary.NewProvider(filter, config, store)">`summaryProvider`</SwmToken>.

<SwmSnippet path="/pkg/collector/corechecks/containers/kubelet/kubelet.go" line="61">

---

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="61:2:2" line-data="func initProviders(filter *containers.Filter, config *common.KubeletConfig, podUtils *common.PodUtils, store workloadmeta.Component) []Provider {">`initProviders`</SwmToken> function initializes various providers that collect metrics from different Kubelet endpoints, such as <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="62:1:1" line-data="	podProvider := pod.NewProvider(filter, config, podUtils)">`podProvider`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="63:3:3" line-data="	// nodeProvider collects from the /spec endpoint, which was hidden by default in k8s 1.18 and removed in k8s 1.19.">`nodeProvider`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="66:1:1" line-data="	healthProvider := health.NewProvider(config)">`healthProvider`</SwmToken>, and <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="67:1:1" line-data="	summaryProvider := summary.NewProvider(filter, config, store)">`summaryProvider`</SwmToken>.

```go
func initProviders(filter *containers.Filter, config *common.KubeletConfig, podUtils *common.PodUtils, store workloadmeta.Component) []Provider {
	podProvider := pod.NewProvider(filter, config, podUtils)
	// nodeProvider collects from the /spec endpoint, which was hidden by default in k8s 1.18 and removed in k8s 1.19.
	// It is here for backwards compatibility.
	nodeProvider := node.NewProvider(config)
	healthProvider := health.NewProvider(config)
	summaryProvider := summary.NewProvider(filter, config, store)

	sliProvider, err := slis.NewProvider(filter, config, store)
	if err != nil {
		log.Warnf("Can't get sli provider: %v", err)
	}

	probeProvider, err := probe.NewProvider(filter, config, store)
	if err != nil {
		log.Warnf("Can't get probe provider: %v", err)
	}

	kubeletProvider, err := kube.NewProvider(filter, config, store, podUtils)
	if err != nil {
		log.Warnf("Can't get kubelet provider: %v", err)
```

---

</SwmSnippet>

# Provider Interface

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="37:2:2" line-data="// Provider provides the metrics related to a given Kubelet endpoint">`Provider`</SwmToken> interface defines the methods required to collect metrics from a given Kubelet endpoint. Implementations of this interface are responsible for fetching and sending metrics to the Datadog platform.

<SwmSnippet path="/pkg/collector/corechecks/containers/kubelet/kubelet.go" line="37">

---

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="37:2:2" line-data="// Provider provides the metrics related to a given Kubelet endpoint">`Provider`</SwmToken> interface defines the methods required to collect metrics from a given Kubelet endpoint.

```go
// Provider provides the metrics related to a given Kubelet endpoint
type Provider interface {
	Provide(kubelet.KubeUtilInterface, sender.Sender) error
}
```

---

</SwmSnippet>

# Main Functions

There are several main functions related to Kubelet. Some of them are <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="52:2:2" line-data="// NewKubeletCheck returns a new KubeletCheck">`NewKubeletCheck`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="61:2:2" line-data="func initProviders(filter *containers.Filter, config *common.KubeletConfig, podUtils *common.PodUtils, store workloadmeta.Component) []Provider {">`initProviders`</SwmToken>, and <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:22:22" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`run`</SwmToken>. We will dive a little into these functions.

## <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="52:2:2" line-data="// NewKubeletCheck returns a new KubeletCheck">`NewKubeletCheck`</SwmToken>

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="52:2:2" line-data="// NewKubeletCheck returns a new KubeletCheck">`NewKubeletCheck`</SwmToken> function initializes a new <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:2:2" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`KubeletCheck`</SwmToken> struct. It sets up the base check, instance configuration, and metric store, which are essential for collecting and processing metrics from the Kubelet.

<SwmSnippet path="/pkg/collector/corechecks/containers/kubelet/kubelet.go" line="52">

---

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="52:2:2" line-data="// NewKubeletCheck returns a new KubeletCheck">`NewKubeletCheck`</SwmToken> function initializes the <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="52:10:10" line-data="// NewKubeletCheck returns a new KubeletCheck">`KubeletCheck`</SwmToken> struct by setting up the base check, instance configuration, and metric store.

```go
// NewKubeletCheck returns a new KubeletCheck
func NewKubeletCheck(base core.CheckBase, instance *common.KubeletConfig, store workloadmeta.Component) *KubeletCheck {
	return &KubeletCheck{
		CheckBase: base,
		instance:  instance,
		store:     store,
	}
}
```

---

</SwmSnippet>

## <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="61:2:2" line-data="func initProviders(filter *containers.Filter, config *common.KubeletConfig, podUtils *common.PodUtils, store workloadmeta.Component) []Provider {">`initProviders`</SwmToken>

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="61:2:2" line-data="func initProviders(filter *containers.Filter, config *common.KubeletConfig, podUtils *common.PodUtils, store workloadmeta.Component) []Provider {">`initProviders`</SwmToken> function initializes various providers that collect metrics from different Kubelet endpoints. These providers include <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="62:1:1" line-data="	podProvider := pod.NewProvider(filter, config, podUtils)">`podProvider`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="63:3:3" line-data="	// nodeProvider collects from the /spec endpoint, which was hidden by default in k8s 1.18 and removed in k8s 1.19.">`nodeProvider`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="66:1:1" line-data="	healthProvider := health.NewProvider(config)">`healthProvider`</SwmToken>, and <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="67:1:1" line-data="	summaryProvider := summary.NewProvider(filter, config, store)">`summaryProvider`</SwmToken>.

<SwmSnippet path="/pkg/collector/corechecks/containers/kubelet/kubelet.go" line="61">

---

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="61:2:2" line-data="func initProviders(filter *containers.Filter, config *common.KubeletConfig, podUtils *common.PodUtils, store workloadmeta.Component) []Provider {">`initProviders`</SwmToken> function initializes various providers that collect metrics from different Kubelet endpoints, such as <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="62:1:1" line-data="	podProvider := pod.NewProvider(filter, config, podUtils)">`podProvider`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="63:3:3" line-data="	// nodeProvider collects from the /spec endpoint, which was hidden by default in k8s 1.18 and removed in k8s 1.19.">`nodeProvider`</SwmToken>, <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="66:1:1" line-data="	healthProvider := health.NewProvider(config)">`healthProvider`</SwmToken>, and <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="67:1:1" line-data="	summaryProvider := summary.NewProvider(filter, config, store)">`summaryProvider`</SwmToken>.

```go
func initProviders(filter *containers.Filter, config *common.KubeletConfig, podUtils *common.PodUtils, store workloadmeta.Component) []Provider {
	podProvider := pod.NewProvider(filter, config, podUtils)
	// nodeProvider collects from the /spec endpoint, which was hidden by default in k8s 1.18 and removed in k8s 1.19.
	// It is here for backwards compatibility.
	nodeProvider := node.NewProvider(config)
	healthProvider := health.NewProvider(config)
	summaryProvider := summary.NewProvider(filter, config, store)

	sliProvider, err := slis.NewProvider(filter, config, store)
	if err != nil {
		log.Warnf("Can't get sli provider: %v", err)
	}

	probeProvider, err := probe.NewProvider(filter, config, store)
	if err != nil {
		log.Warnf("Can't get probe provider: %v", err)
	}

	kubeletProvider, err := kube.NewProvider(filter, config, store, podUtils)
	if err != nil {
		log.Warnf("Can't get kubelet provider: %v", err)
```

---

</SwmSnippet>

## Run

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="42:22:22" line-data="// KubeletCheck wraps the config and the metric stores needed to run the check">`run`</SwmToken> function executes the Kubelet check. It retrieves the Kubelet client, iterates through the providers, and collects metrics, which are then sent to the Datadog platform.

<SwmSnippet path="/pkg/collector/corechecks/containers/kubelet/kubelet.go" line="138">

---

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="138:2:2" line-data="// Run runs the check">`Run`</SwmToken> function executes the Kubelet check, retrieves the Kubelet client, iterates through the providers, and collects metrics.

```go
// Run runs the check
func (k *KubeletCheck) Run() error {
	sender, err := k.GetSender()
	if err != nil {
		return err
	}
	defer sender.Commit()
	defer k.podUtils.Reset()

	// Get client
	kc, err := kubelet.GetKubeUtil()
	if err != nil {
		_ = k.Warnf("Error initialising check: %s", err)
		return err
	}

	for _, provider := range k.providers {
		if provider != nil {
			err = provider.Provide(kc, sender)
			if err != nil {
				_ = k.Warnf("Error reporting metrics: %s", err)
```

---

</SwmSnippet>

# Kubelet Endpoints

Kubelet exposes several endpoints for collecting metrics. Two important endpoints are <SwmPath>[pkg/collector/corechecks/containers/kubelet/provider/summary/](pkg/collector/corechecks/containers/kubelet/provider/summary/)</SwmPath> and <SwmPath>[pkg/network/ebpf/probes/](pkg/network/ebpf/probes/)</SwmPath>.

## <SwmPath>[pkg/collector/corechecks/containers/kubelet/provider/summary/](pkg/collector/corechecks/containers/kubelet/provider/summary/)</SwmPath>

The <SwmPath>[pkg/collector/corechecks/containers/kubelet/provider/summary/](pkg/collector/corechecks/containers/kubelet/provider/summary/)</SwmPath> endpoint provides a summary of the resource usage statistics for pods and containers. The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="37:2:2" line-data="// Provider provides the metrics related to a given Kubelet endpoint">`Provider`</SwmToken> struct in <SwmPath>[pkg/collector/corechecks/containers/kubelet/provider/probe/provider.go](pkg/collector/corechecks/containers/kubelet/provider/probe/provider.go)</SwmPath> uses this endpoint to collect and process metrics, which are then reported to the Datadog platform. The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="39:1:1" line-data="	Provide(kubelet.KubeUtilInterface, sender.Sender) error">`Provide`</SwmToken> method fetches the stats summary and processes system stats, pod stats, and container stats.

<SwmSnippet path="/pkg/collector/corechecks/containers/kubelet/provider/summary/provider.go" line="34">

---

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/provider/summary/provider.go" pos="34:2:2" line-data="// Provider provides the data collected from the `/stats/summary` Kubelet endpoint">`Provider`</SwmToken> struct uses the <SwmPath>[pkg/collector/corechecks/containers/kubelet/provider/summary/](pkg/collector/corechecks/containers/kubelet/provider/summary/)</SwmPath> endpoint to collect and process metrics, which are then reported to the Datadog platform.

```go
// Provider provides the data collected from the `/stats/summary` Kubelet endpoint
type Provider struct {
	filter                *containers.Filter
	config                *common.KubeletConfig
	store                 workloadmeta.Component
	defaultRateFilterList []*regexp.Regexp
}

// NewProvider is created by filter, config and workloadmeta
func NewProvider(filter *containers.Filter,
	config *common.KubeletConfig,
	store workloadmeta.Component) *Provider {
	defaultRateFilterList := []*regexp.Regexp{
		regexp.MustCompile("diskio[.]io_service_bytes[.]stats[.]total"),
		regexp.MustCompile("network[.].._bytes"),
		regexp.MustCompile("cpu[.].*[.]total"),
	} //default enabled_rates

	return &Provider{
		filter:                filter,
		config:                config,
```

---

</SwmSnippet>

## <SwmPath>[pkg/network/ebpf/probes/](pkg/network/ebpf/probes/)</SwmPath>

The <SwmPath>[pkg/network/ebpf/probes/](pkg/network/ebpf/probes/)</SwmPath> endpoint provides metrics related to the health probes of containers. The <SwmToken path="pkg/collector/corechecks/containers/kubelet/kubelet.go" pos="37:2:2" line-data="// Provider provides the metrics related to a given Kubelet endpoint">`Provider`</SwmToken> struct in <SwmPath>[pkg/collector/corechecks/containers/kubelet/provider/probe/provider.go](pkg/collector/corechecks/containers/kubelet/provider/probe/provider.go)</SwmPath> uses this endpoint to collect probe metrics such as liveness, readiness, and startup probes. The <SwmToken path="pkg/collector/corechecks/containers/kubelet/provider/probe/provider.go" pos="40:8:8" line-data="		&quot;prober_probe_total&quot;: provider.proberProbeTotal,">`proberProbeTotal`</SwmToken> method processes these metrics and sends them to the Datadog platform.

<SwmSnippet path="/pkg/collector/corechecks/containers/kubelet/provider/probe/provider.go" line="25">

---

The <SwmToken path="pkg/collector/corechecks/containers/kubelet/provider/probe/provider.go" pos="25:2:2" line-data="// Provider provides the metrics related to data collected from the `/metrics/probes` Kubelet endpoint">`Provider`</SwmToken> struct uses the <SwmPath>[pkg/network/ebpf/probes/](pkg/network/ebpf/probes/)</SwmPath> endpoint to collect probe metrics such as liveness, readiness, and startup probes.

```go
// Provider provides the metrics related to data collected from the `/metrics/probes` Kubelet endpoint
type Provider struct {
	filter *containers.Filter
	store  workloadmeta.Component
	prometheus.Provider
}

// NewProvider returns a metrics prometheus kubelet provider and an error
func NewProvider(filter *containers.Filter, config *common.KubeletConfig, store workloadmeta.Component) (*Provider, error) {
	provider := &Provider{
		filter: filter,
		store:  store,
	}

	transformers := prometheus.Transformers{
		"prober_probe_total": provider.proberProbeTotal,
	}

	scraperConfig := &prometheus.ScraperConfig{AllowNotFound: true}
	if config.ProbesMetricsEndpoint == nil || *config.ProbesMetricsEndpoint != "" {
		scraperConfig.Path = "/metrics/probes"
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
