---
title: Exploring Trace Components
---
```mermaid
graph TD;
    A[Trace Agent Configuration] --> B[Trace Agent Status];
    B --> C[Trace Agent Components];
    C --> D[Trace Endpoints];
    D --> E[/v0.7/config];
    D --> F[/config/set];
```

<SwmSnippet path="/comp/trace/agent/impl/run.go" line="37">

---

# Trace Agent Configuration

The function <SwmToken path="comp/trace/agent/impl/run.go" pos="37:2:2" line-data="// runAgentSidekicks is the entrypoint for running non-components that run along the agent.">`runAgentSidekicks`</SwmToken> initializes and configures the <SwmToken path="comp/trace/agent/impl/run.go" pos="17:16:18" line-data="	remotecfg &quot;github.com/DataDog/datadog-agent/cmd/trace-agent/config/remote&quot;">`trace-agent`</SwmToken>, setting up necessary endpoints and runtime parameters. This is crucial for ensuring the <SwmToken path="comp/trace/agent/impl/run.go" pos="17:16:18" line-data="	remotecfg &quot;github.com/DataDog/datadog-agent/cmd/trace-agent/config/remote&quot;">`trace-agent`</SwmToken> operates correctly and efficiently.

```go
// runAgentSidekicks is the entrypoint for running non-components that run along the agent.
func runAgentSidekicks(ag component) error {
	tracecfg := ag.config.Object()
	err := info.InitInfo(tracecfg) // for expvar & -info option
	if err != nil {
		return err
	}

	defer watchdog.LogOnPanic(ag.Statsd)

	if err := util.SetupCoreDump(coreconfig.Datadog()); err != nil {
		log.Warnf("Can't setup core dumps: %v, core dumps might not be available after a crash", err)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	if coreconfig.IsRemoteConfigEnabled(coreconfig.Datadog()) {
		cf, err := newConfigFetcher()
		if err != nil {
			ag.telemetryCollector.SendStartupError(telemetry.CantCreateRCCLient, err)
			return fmt.Errorf("could not instantiate the tracer remote config client: %v", err)
```

---

</SwmSnippet>

# Trace Agent Components

The trace component bundle provides essential components for the Trace Agent, including configuration handlers, status providers, and compression utilities. These components are integral to the trace-agent's functionality, ensuring it can handle various tasks efficiently.

# Trace Endpoints

Trace endpoints are used to handle configuration requests and set runtime configurations dynamically. These endpoints are crucial for interacting with the <SwmToken path="comp/trace/agent/impl/run.go" pos="17:16:18" line-data="	remotecfg &quot;github.com/DataDog/datadog-agent/cmd/trace-agent/config/remote&quot;">`trace-agent`</SwmToken> and modifying its behavior as needed.

<SwmSnippet path="/comp/trace/agent/impl/run.go" line="60">

---

## <SwmPath>[cmd/agent/subcommands/config/](cmd/agent/subcommands/config/)</SwmPath>

The <SwmPath>[cmd/agent/subcommands/config/](cmd/agent/subcommands/config/)</SwmPath> endpoint is used to handle configuration requests. It attaches a handler that processes configuration data and interacts with the remote configuration client.

```go
		api.AttachEndpoint(api.Endpoint{
			Pattern: "/v0.7/config",
			Handler: func(r *api.HTTPReceiver) http.Handler {
				return remotecfg.ConfigHandler(r, cf, tracecfg, ag.Statsd, ag.Timing)
			},
		})
```

---

</SwmSnippet>

<SwmSnippet path="/comp/trace/agent/impl/run.go" line="78">

---

## <SwmToken path="comp/trace/agent/impl/run.go" pos="79:5:8" line-data="		Pattern: &quot;/config/set&quot;,">`/config/set`</SwmToken>

The <SwmToken path="comp/trace/agent/impl/run.go" pos="79:5:8" line-data="		Pattern: &quot;/config/set&quot;,">`/config/set`</SwmToken> endpoint allows for setting runtime configuration. It attaches a handler that processes POST requests to change configuration settings dynamically.

```go
	api.AttachEndpoint(api.Endpoint{
		Pattern: "/config/set",
		Handler: func(r *api.HTTPReceiver) http.Handler {
			return ag.config.SetHandler()
		},
	})
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
