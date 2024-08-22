---
title: Getting Started with HTTP/2 Protocol
---
```mermaid
classDiagram
 class Protocol {
 +Config cfg
 +Manager mgr
 +Telemetry telemetry
 +StatKeeper statkeeper
 +MapCleaner http2InFlightMapCleaner
 +Consumer eventsConsumer
 +KernelTelemetry http2Telemetry
 +Channel kernelTelemetryStopChannel
 +DynamicTable dynamicTable
 }
 class kernelTelemetry {
 +MetricGroup metricGroup
 +Counter http2requests
 +Counter http2responses
 +Counter endOfStream
 +Counter endOfStreamRST
 +Array pathSizeBucket
 +Counter literalValueExceedsFrame
 +Counter exceedingMaxInterestingFrames
 +Counter exceedingMaxFramesToFilter
 +Counter fragmentedFrameCountRST
 +Counter fragmentedHeadersFrameEOSCount
 +Counter fragmentedHeadersFrameCount
 +Counter fragmentedDataFrameEOSCount
 +HTTP2Telemetry telemetryLastState
 }
graph TD;
 A[HTTP/2 Connection] --> B[InFlightMap];
 A --> C[TelemetryMap];
 B --> D[Track Streams];
 C --> E[Collect Metrics];

%% Swimm:
%% classDiagram
%%  class Protocol {
%%  +Config cfg
%%  +Manager mgr
%%  +Telemetry telemetry
%%  +StatKeeper statkeeper
%%  +MapCleaner <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="37:1:1" line-data="	http2InFlightMapCleaner *ddebpf.MapCleaner[HTTP2StreamKey, HTTP2Stream]">`http2InFlightMapCleaner`</SwmToken>
%%  +Consumer <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="38:1:1" line-data="	eventsConsumer          *events.Consumer[EbpfTx]">`eventsConsumer`</SwmToken>
%%  +KernelTelemetry <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="40:3:3" line-data="	// http2Telemetry is used to retrieve metrics from the kernel">`http2Telemetry`</SwmToken>
%%  +Channel <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="42:1:1" line-data="	kernelTelemetryStopChannel chan struct{}">`kernelTelemetryStopChannel`</SwmToken>
%%  +DynamicTable <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="44:1:1" line-data="	dynamicTable *DynamicTable">`dynamicTable`</SwmToken>
%%  }
%%  class <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="41:4:4" line-data="	http2Telemetry             *kernelTelemetry">`kernelTelemetry`</SwmToken> {
%%  +MetricGroup <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="50:3:3" line-data="	// metricGroup is used here mostly for building the log message below">`metricGroup`</SwmToken>
%%  +Counter <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="53:3:3" line-data="	// http2requests Count of HTTP/2 requests seen">`http2requests`</SwmToken>
%%  +Counter <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="55:3:3" line-data="	// http2responses Count of HTTP/2 responses seen">`http2responses`</SwmToken>
%%  +Counter <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="57:3:3" line-data="	// endOfStream Count of END_OF_STREAM flags seen">`endOfStream`</SwmToken>
%%  +Counter <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="59:3:3" line-data="	// endOfStreamRST Count of RST flags seen">`endOfStreamRST`</SwmToken>
%%  +Array <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="61:3:3" line-data="	// pathSizeBucket Count of path sizes divided into buckets.">`pathSizeBucket`</SwmToken>
%%  +Counter <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="63:3:3" line-data="	// literalValueExceedsFrame Count of times we couldn&#39;t retrieve the literal value due to reaching the end of the frame.">`literalValueExceedsFrame`</SwmToken>
%%  +Counter <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="65:3:3" line-data="	// exceedingMaxInterestingFrames Count of times we reached the max number of frames per iteration.">`exceedingMaxInterestingFrames`</SwmToken>
%%  +Counter <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="67:3:3" line-data="	// exceedingMaxFramesToFilter Count of times we have left with more frames to filter than the max number of frames to filter.">`exceedingMaxFramesToFilter`</SwmToken>
%%  +Counter <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="69:3:3" line-data="	// fragmentedFrameCountRST Count of times we have seen a fragmented RST frame.">`fragmentedFrameCountRST`</SwmToken>
%%  +Counter <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="71:3:3" line-data="	// fragmentedHeadersFrameEOSCount Count of times we have seen a fragmented headers frame with EOS.">`fragmentedHeadersFrameEOSCount`</SwmToken>
%%  +Counter <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="73:3:3" line-data="	// fragmentedHeadersFrameCount Count of times we have seen a fragmented headers frame.">`fragmentedHeadersFrameCount`</SwmToken>
%%  +Counter <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="75:3:3" line-data="	// fragmentedDataFrameEOSCount Count of times we have seen a fragmented data frame with EOS.">`fragmentedDataFrameEOSCount`</SwmToken>
%%  +HTTP2Telemetry <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="77:3:3" line-data="	// telemetryLastState represents the latest HTTP2 eBPF Kernel telemetry observed from the kernel">`telemetryLastState`</SwmToken>
%%  }
%% graph TD;
%%  A[HTTP/2 Connection] --> B[InFlightMap];
%%  A --> C[TelemetryMap];
%%  B --> D[Track Streams];
%%  C --> E[Collect Metrics];
```

# Introduction

<SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> is a protocol used for transferring data over the web. It is designed to improve the performance of web applications by allowing multiple requests and responses to be multiplexed over a single connection.

# <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> in Datadog Agent

In the Datadog Agent, the <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="239:14:14" line-data="// ConfigureOptions add the necessary options for http2 monitoring to work,">`http2`</SwmToken> package is responsible for handling the <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> protocol. This includes managing the state of <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> streams, parsing <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> frames, and collecting telemetry data.

<SwmSnippet path="/pkg/network/protocols/http2/protocol.go" line="31">

---

## Protocol Struct

The <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:2:2" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`Protocol`</SwmToken> struct in the <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="239:14:14" line-data="// ConfigureOptions add the necessary options for http2 monitoring to work,">`http2`</SwmToken> package contains various fields and methods for managing <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> connections. It includes a <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="35:1:1" line-data="	telemetry               *http.Telemetry">`telemetry`</SwmToken> field for collecting metrics from the kernel, a <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="44:1:1" line-data="	dynamicTable *DynamicTable">`dynamicTable`</SwmToken> for managing the dynamic table used in <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> header compression, and several constants for managing different aspects of the <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> protocol.

```go
// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.
type Protocol struct {
	cfg                     *config.Config
	mgr                     *manager.Manager
	telemetry               *http.Telemetry
	statkeeper              *http.StatKeeper
	http2InFlightMapCleaner *ddebpf.MapCleaner[HTTP2StreamKey, HTTP2Stream]
	eventsConsumer          *events.Consumer[EbpfTx]

	// http2Telemetry is used to retrieve metrics from the kernel
	http2Telemetry             *kernelTelemetry
	kernelTelemetryStopChannel chan struct{}

	dynamicTable *DynamicTable
}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/http2/telemetry.go" line="49">

---

## Kernel Telemetry

The <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="49:2:2" line-data="type kernelTelemetry struct {">`kernelTelemetry`</SwmToken> struct in the <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="239:14:14" line-data="// ConfigureOptions add the necessary options for http2 monitoring to work,">`http2`</SwmToken> package is used to collect various metrics related to <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="53:9:11" line-data="	// http2requests Count of HTTP/2 requests seen">`HTTP/2`</SwmToken> traffic. This includes counts of <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="53:9:11" line-data="	// http2requests Count of HTTP/2 requests seen">`HTTP/2`</SwmToken> requests and responses, end-of-stream flags, and other metrics that help in monitoring the performance and behavior of <SwmToken path="pkg/network/protocols/http2/telemetry.go" pos="53:9:11" line-data="	// http2requests Count of HTTP/2 requests seen">`HTTP/2`</SwmToken> connections.

```go
type kernelTelemetry struct {
	// metricGroup is used here mostly for building the log message below
	metricGroup *libtelemetry.MetricGroup

	// http2requests Count of HTTP/2 requests seen
	http2requests *tlsAwareCounter
	// http2responses Count of HTTP/2 responses seen
	http2responses *tlsAwareCounter
	// endOfStream Count of END_OF_STREAM flags seen
	endOfStream *tlsAwareCounter
	// endOfStreamRST Count of RST flags seen
	endOfStreamRST *tlsAwareCounter
	// pathSizeBucket Count of path sizes divided into buckets.
	pathSizeBucket [http2PathBuckets + 1]*tlsAwareCounter
	// literalValueExceedsFrame Count of times we couldn't retrieve the literal value due to reaching the end of the frame.
	literalValueExceedsFrame *tlsAwareCounter
	// exceedingMaxInterestingFrames Count of times we reached the max number of frames per iteration.
	exceedingMaxInterestingFrames *tlsAwareCounter
	// exceedingMaxFramesToFilter Count of times we have left with more frames to filter than the max number of frames to filter.
	exceedingMaxFramesToFilter *tlsAwareCounter
	// fragmentedFrameCountRST Count of times we have seen a fragmented RST frame.
```

---

</SwmSnippet>

# Configuring <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> Monitoring

To enable <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> monitoring, several configuration steps are required. These include setting up options, enabling specific features, and configuring event processing.

<SwmSnippet path="/pkg/network/protocols/http2/protocol.go" line="239">

---

## <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="239:2:2" line-data="// ConfigureOptions add the necessary options for http2 monitoring to work,">`ConfigureOptions`</SwmToken> Method

The <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="239:2:2" line-data="// ConfigureOptions add the necessary options for http2 monitoring to work,">`ConfigureOptions`</SwmToken> method is used to add the necessary options for <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> monitoring. It sets the size of the <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="241:9:9" line-data="// - Set the `http2_in_flight` map size to the value of the `max_tracked_connection` configuration variable.">`http2_in_flight`</SwmToken> map and configures the <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> event stream with the manager and its options.

```go
// ConfigureOptions add the necessary options for http2 monitoring to work,
// to be used by the manager. These are:
// - Set the `http2_in_flight` map size to the value of the `max_tracked_connection` configuration variable.
//
// We also configure the http2 event stream with the manager and its options.
func (p *Protocol) ConfigureOptions(mgr *manager.Manager, opts *manager.Options) {
	opts.MapSpecEditors[InFlightMap] = manager.MapSpecEditor{
		MaxEntries: p.cfg.MaxUSMConcurrentRequests,
		EditorFlag: manager.EditMaxEntries,
	}
	opts.MapSpecEditors[remainderTable] = manager.MapSpecEditor{
		MaxEntries: p.cfg.MaxUSMConcurrentRequests,
		EditorFlag: manager.EditMaxEntries,
	}
	opts.MapSpecEditors[dynamicTable] = manager.MapSpecEditor{
		MaxEntries: p.cfg.MaxUSMConcurrentRequests,
		EditorFlag: manager.EditMaxEntries,
	}
	opts.MapSpecEditors[dynamicTableCounter] = manager.MapSpecEditor{
		MaxEntries: p.cfg.MaxUSMConcurrentRequests,
		EditorFlag: manager.EditMaxEntries,
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/usm/utils/ebpfoptions.go" line="32">

---

## <SwmToken path="pkg/network/usm/utils/ebpfoptions.go" pos="32:2:2" line-data="// EnableOption adds a constant editor to the options with the given name and value true">`EnableOption`</SwmToken> Function

The <SwmToken path="pkg/network/usm/utils/ebpfoptions.go" pos="32:2:2" line-data="// EnableOption adds a constant editor to the options with the given name and value true">`EnableOption`</SwmToken> function adds a constant editor to the options with the given name and value set to true. This is used within <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="239:2:2" line-data="// ConfigureOptions add the necessary options for http2 monitoring to work,">`ConfigureOptions`</SwmToken> to enable <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> monitoring options.

```go
// EnableOption adds a constant editor to the options with the given name and value true
func EnableOption(options *manager.Options, name string) {
	options.ConstantEditors = append(options.ConstantEditors,
		manager.ConstantEditor{
			Name:  name,
			Value: enabled,
		},
	)
}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/events/configuration.go" line="38">

---

## Configure Function

The <SwmToken path="pkg/network/protocols/events/configuration.go" pos="38:2:2" line-data="// Configure a given `*manager.Manager` for event processing">`Configure`</SwmToken> function sets up the manager for event processing, instantiating the perf <SwmToken path="pkg/network/protocols/events/configuration.go" pos="39:12:14" line-data="// This essentially instantiates the perf map/ring buffers and configure the">`map/ring`</SwmToken> buffers and configuring the <SwmToken path="pkg/network/protocols/events/configuration.go" pos="40:2:2" line-data="// eBPF maps where events are enqueued.">`eBPF`</SwmToken> maps where events are enqueued. This is called within <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="239:2:2" line-data="// ConfigureOptions add the necessary options for http2 monitoring to work,">`ConfigureOptions`</SwmToken> to configure the event stream.

```go
// Configure a given `*manager.Manager` for event processing
// This essentially instantiates the perf map/ring buffers and configure the
// eBPF maps where events are enqueued.
// Note this must be called *before* manager.InitWithOptions
func Configure(cfg *config.Config, proto string, m *manager.Manager, o *manager.Options) {
	if alreadySetUp(proto, m) {
		return
	}

	numCPUs, err := kernel.PossibleCPUs()
	if err != nil {
		numCPUs = 96
		log.Error("unable to detect number of CPUs. assuming 96 cores")
	}

	configureBatchMaps(proto, o, numCPUs)

	useRingBuffer := cfg.EnableUSMRingBuffers && features.HaveMapType(ebpf.RingBuf) == nil
	utils.AddBoolConst(o, useRingBuffer, "use_ring_buffer")

	if useRingBuffer {
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/http2/dynamic_table.go" line="49">

---

## <SwmToken path="pkg/network/protocols/http2/dynamic_table.go" pos="50:6:6" line-data="func (dt *DynamicTable) configureOptions(mgr *manager.Manager, opts *manager.Options) {">`DynamicTable`</SwmToken> Configuration

The <SwmToken path="pkg/network/protocols/http2/dynamic_table.go" pos="49:2:2" line-data="// configureOptions configures the perf handler options for the map cleaner.">`configureOptions`</SwmToken> method in the <SwmToken path="pkg/network/protocols/http2/dynamic_table.go" pos="50:6:6" line-data="func (dt *DynamicTable) configureOptions(mgr *manager.Manager, opts *manager.Options) {">`DynamicTable`</SwmToken> struct configures the perf handler options for the map cleaner. This is part of the overall configuration process in <SwmToken path="pkg/network/protocols/http2/dynamic_table.go" pos="49:2:2" line-data="// configureOptions configures the perf handler options for the map cleaner.">`configureOptions`</SwmToken>.

```go
// configureOptions configures the perf handler options for the map cleaner.
func (dt *DynamicTable) configureOptions(mgr *manager.Manager, opts *manager.Options) {
	events.Configure(dt.cfg, terminatedConnectionsEventStream, mgr, opts)
}
```

---

</SwmSnippet>

# <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> Endpoints

<SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> endpoints are used to store and retrieve data related to <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> streams and telemetry.

<SwmSnippet path="/pkg/network/protocols/http2/protocol.go" line="47">

---

## <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="48:3:3" line-data="	// InFlightMap is the name of the map used to store in-flight HTTP/2 streams">`InFlightMap`</SwmToken>

The <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="48:3:3" line-data="	// InFlightMap is the name of the map used to store in-flight HTTP/2 streams">`InFlightMap`</SwmToken> is used to store <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="48:23:25" line-data="	// InFlightMap is the name of the map used to store in-flight HTTP/2 streams">`in-flight`</SwmToken> <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="48:27:29" line-data="	// InFlightMap is the name of the map used to store in-flight HTTP/2 streams">`HTTP/2`</SwmToken> streams. This map helps in tracking the state of ongoing <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="48:27:29" line-data="	// InFlightMap is the name of the map used to store in-flight HTTP/2 streams">`HTTP/2`</SwmToken> streams, which is crucial for monitoring and managing <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="48:27:29" line-data="	// InFlightMap is the name of the map used to store in-flight HTTP/2 streams">`HTTP/2`</SwmToken> connections.

```go
const (
	// InFlightMap is the name of the map used to store in-flight HTTP/2 streams
	InFlightMap               = "http2_in_flight"
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/http2/protocol.go" line="62">

---

## <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="62:3:3" line-data="	// TelemetryMap is the name of the map used to retrieve plaintext metrics from the kernel">`TelemetryMap`</SwmToken>

The <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="62:3:3" line-data="	// TelemetryMap is the name of the map used to retrieve plaintext metrics from the kernel">`TelemetryMap`</SwmToken> is used to retrieve plaintext metrics from the kernel. This map is essential for collecting various metrics related to <SwmToken path="pkg/network/protocols/http2/protocol.go" pos="31:26:28" line-data="// Protocol implements the interface that represents a protocol supported by USM for HTTP/2.">`HTTP/2`</SwmToken> traffic, such as counts of requests and responses, which are used for performance monitoring.

```go
	// TelemetryMap is the name of the map used to retrieve plaintext metrics from the kernel
	TelemetryMap = "http2_telemetry"
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
