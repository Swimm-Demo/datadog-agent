---
title: Introduction to Trace in Packages
---
```mermaid
graph TD;
    A[Trace Collection] --> B[Trace Agent];
    B --> C[Trace Sampling];
    B --> D[Trace Processing];
    D --> E[Trace Utilities];
    D --> F[Trace Core Functionality];
    F --> G[Trace Endpoints];
    G --> H[/v0.3/traces];
    G --> I[/v0.4/traces];
```

# Introduction to Trace in Packages

Trace refers to the collection of APM (Application Performance Monitoring) traces. These traces are used to monitor and analyze the performance of applications by tracking the flow of requests through various services.

# Trace Agent

The <SwmToken path="pkg/trace/api/endpoints.go" pos="36:16:18" line-data="// AttachEndpoint attaches an additional endpoint to the trace-agent. It is not thread-safe">`trace-agent`</SwmToken> is responsible for collecting APM traces. It processes and samples traces before sending them to the Datadog platform.

<SwmSnippet path="/pkg/trace/writer/trace.go" line="42">

---

# Trace Sampling

The <SwmToken path="pkg/trace/writer/trace.go" pos="42:2:2" line-data="// SampledChunks represents the result of a trace sampling operation.">`SampledChunks`</SwmToken> type represents the result of a trace sampling operation, containing sampled chunks, their size, and the number of spans and events. This is crucial for understanding how traces are processed and stored.

```go
// SampledChunks represents the result of a trace sampling operation.
type SampledChunks struct {
	// TracerPayload contains all the chunks that were sampled as part of processing a payload.
	TracerPayload *pb.TracerPayload
	// Size represents the approximated message size in bytes.
	Size int
	// SpanCount specifies the number of spans that were sampled as part of a trace inside the TracerPayload.
	SpanCount int64
	// EventCount specifies the total number of events found in Traces.
	EventCount int64
}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/trace/traceutil/doc.go" line="6">

---

# Trace Utilities

The <SwmToken path="pkg/trace/traceutil/doc.go" pos="6:4:4" line-data="// Package traceutil contains functions for extracting and processing traces. It should">`traceutil`</SwmToken> package contains functions for extracting and processing traces. It should only import payload and nothing else, ensuring minimal dependencies.

```go
// Package traceutil contains functions for extracting and processing traces. It should
// only import payload and nothing else.
package traceutil
```

---

</SwmSnippet>

# Trace Core Functionality

The <SwmToken path="pkg/trace/writer/trace.go" pos="42:14:14" line-data="// SampledChunks represents the result of a trace sampling operation.">`trace`</SwmToken> package provides the core functionality for handling traces, including the API for trace operations. This package is essential for the overall trace handling mechanism.

# Trace Endpoints

Trace endpoints are used to handle trace data. They are registered with the HTTP receiver and process incoming trace payloads.

<SwmSnippet path="/pkg/trace/api/endpoints.go" line="74">

---

## <SwmToken path="pkg/trace/api/endpoints.go" pos="74:5:10" line-data="		Pattern: &quot;/v0.3/traces&quot;,">`/v0.3/traces`</SwmToken>

The <SwmToken path="pkg/trace/api/endpoints.go" pos="74:5:10" line-data="		Pattern: &quot;/v0.3/traces&quot;,">`/v0.3/traces`</SwmToken> endpoint is used to handle trace data. It is registered with the HTTP receiver and processes incoming trace payloads.

```go
		Pattern: "/v0.3/traces",
		Handler: func(r *HTTPReceiver) http.Handler { return r.handleWithVersion(v03, r.handleTraces) },
	},
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
