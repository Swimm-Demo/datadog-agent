---
title: Overview of Flare in Packages
---
```mermaid
graph TD;
 A[Component] -->|Registers Callback| B[Flare]
 B -->|Uses| C[Fx Groups]
 B -->|Migrates Code| D[pkg/flare]
 D -->|Calls| E[CompleteFlare]
 E -->|Uses| F[FlareBuilder]

 graph TD;
 A[GetConfigState] -->|Calls| B[cli.GetConfigState];
 C[GetConfigStateHA] -->|Calls| D[cli.GetConfigStateHA];
 B -->|Stores response in| E[s];
 D -->|Stores response in| F[haState];
```

# Overview of Flare in Packages

Flare refers to a mechanism used to collect and package diagnostic information from the Datadog Agent. It is designed to help with troubleshooting by gathering relevant data into a single archive.

## Registering a Callback

The general idea is to register a callback within your component to be called each time a flare is created. This uses Fx groups under the hood, but helpers are there to abstract all the complexity.

## Example Usage

Here is an example of how to use the <SwmToken path="pkg/flare/archive.go" pos="47:8:8" line-data="func CompleteFlare(fb flaretypes.FlareBuilder, diagnoseDeps diagnose.SuitesDeps) error {">`FlareBuilder`</SwmToken> to create a new file within a flare: `import (flare `<SwmPath>[comp/aggregator/demultiplexerendpoint/def/](comp/aggregator/demultiplexerendpoint/def/)</SwmPath>`) `<SwmToken path="pkg/flare/archive.go" pos="47:0:0" line-data="func CompleteFlare(fb flaretypes.FlareBuilder, diagnoseDeps diagnose.SuitesDeps) error {">`func`</SwmToken>` `<SwmToken path="pkg/flare/archive.go" pos="318:11:12" line-data="	r, err := apiutil.DoGet(c, remoteURL, apiutil.LeaveConnectionOpen)">`(c`</SwmToken>` *myComponent) fillFlare(fb flare.FlareBuilder) error{fb.AddFile("runtime_config_dump.yaml", []byte("content `<SwmToken path="pkg/flare/archive.go" pos="56:27:27" line-data="		fb.AddFile(&quot;status.log&quot;, []byte(&quot;unable to get the status of the agent, is it running?&quot;))           //nolint:errcheck">`of`</SwmToken>` my file")) `<SwmToken path="pkg/flare/archive.go" pos="56:44:47" line-data="		fb.AddFile(&quot;status.log&quot;, []byte(&quot;unable to get the status of the agent, is it running?&quot;))           //nolint:errcheck">`//nolint:errcheck`</SwmToken>` }`

## Testing

The flare component offers a <SwmToken path="pkg/flare/archive.go" pos="47:8:8" line-data="func CompleteFlare(fb flaretypes.FlareBuilder, diagnoseDeps diagnose.SuitesDeps) error {">`FlareBuilder`</SwmToken> mock to test your callback. This allows you to ensure that your flare creation logic works as expected.

## Main Functions

There are several main functions in this folder. Some of them are <SwmToken path="pkg/flare/archive.go" pos="45:2:2" line-data="// CompleteFlare packages up the files with an already created builder. This is aimed to be used by the flare">`CompleteFlare`</SwmToken>, <SwmToken path="pkg/flare/flare.go" pos="15:2:2" line-data="// SendFlare sends a flare and returns the message returned by the backend. This entry point is deprecated in favor of">`SendFlare`</SwmToken>, <SwmToken path="pkg/flare/archive.go" pos="73:1:1" line-data="	addSystemProbePlatformSpecificEntries(fb)">`addSystemProbePlatformSpecificEntries`</SwmToken>, getLinuxKernelSymbols, getLinuxKprobeEvents, and getLinuxTracingAvailableFilterFunctions. We will dive a little into <SwmToken path="pkg/flare/archive.go" pos="45:2:2" line-data="// CompleteFlare packages up the files with an already created builder. This is aimed to be used by the flare">`CompleteFlare`</SwmToken> and <SwmToken path="pkg/flare/flare.go" pos="15:2:2" line-data="// SendFlare sends a flare and returns the message returned by the backend. This entry point is deprecated in favor of">`SendFlare`</SwmToken>.

<SwmSnippet path="/pkg/flare/archive.go" line="45">

---

### <SwmToken path="pkg/flare/archive.go" pos="45:2:2" line-data="// CompleteFlare packages up the files with an already created builder. This is aimed to be used by the flare">`CompleteFlare`</SwmToken>

The <SwmToken path="pkg/flare/archive.go" pos="45:2:2" line-data="// CompleteFlare packages up the files with an already created builder. This is aimed to be used by the flare">`CompleteFlare`</SwmToken> function packages up the files with an already created builder. It ensures that no credentials or unnecessary <SwmToken path="pkg/flare/archive.go" pos="51:7:9" line-data="	 * or unnecessary user-specific data. The FlareBuilder scrubs secrets that match pre-programmed patterns, but it">`user-specific`</SwmToken> data are included in the flare. This function is aimed to be used by the flare component while migrating to a component architecture.

```go
// CompleteFlare packages up the files with an already created builder. This is aimed to be used by the flare
// component while we migrate to a component architecture.
func CompleteFlare(fb flaretypes.FlareBuilder, diagnoseDeps diagnose.SuitesDeps) error {
	/** WARNING
	 *
	 * When adding data to flares, carefully analyze what is being added and ensure that it contains no credentials
	 * or unnecessary user-specific data. The FlareBuilder scrubs secrets that match pre-programmed patterns, but it
	 * is always better to not capture data containing secrets, than to scrub that data.
	 */
	if fb.IsLocal() {
		// Can't reach the agent, mention it in those two files
		fb.AddFile("status.log", []byte("unable to get the status of the agent, is it running?"))           //nolint:errcheck
		fb.AddFile("config-check.log", []byte("unable to get loaded checks config, is the agent running?")) //nolint:errcheck
	} else {
		fb.AddFileFromFunc("tagger-list.json", getAgentTaggerList)                      //nolint:errcheck
		fb.AddFileFromFunc("workload-list.log", getAgentWorkloadList)                   //nolint:errcheck
		fb.AddFileFromFunc("process-agent_tagger-list.json", getProcessAgentTaggerList) //nolint:errcheck
		if !config.Datadog().GetBool("process_config.run_in_core_agent.enabled") {
			getChecksFromProcessAgent(fb, config.GetProcessAPIAddressPort)
		}
	}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/flare/flare.go" line="15">

---

### <SwmToken path="pkg/flare/flare.go" pos="15:2:2" line-data="// SendFlare sends a flare and returns the message returned by the backend. This entry point is deprecated in favor of">`SendFlare`</SwmToken>

The <SwmToken path="pkg/flare/flare.go" pos="15:2:2" line-data="// SendFlare sends a flare and returns the message returned by the backend. This entry point is deprecated in favor of">`SendFlare`</SwmToken> function sends a flare and returns the message returned by the backend. This entry point is deprecated in favor of the 'Send' method of the flare component.

```go
// SendFlare sends a flare and returns the message returned by the backend. This entry point is deprecated in favor of
// the 'Send' method of the flare component.
func SendFlare(cfg pkgconfigmodel.Reader, archivePath string, caseID string, email string, source helpers.FlareSource) (string, error) {
	return helpers.SendTo(cfg, archivePath, caseID, email, config.Datadog().GetString("api_key"), utils.GetInfraEndpoint(config.Datadog()), source)
}
```

---

</SwmSnippet>

## Flare Endpoints

Flare Endpoints

<SwmSnippet path="/pkg/flare/remote_config.go" line="61">

---

### <SwmToken path="pkg/flare/remote_config.go" pos="61:10:10" line-data="	s, err := cli.GetConfigState(ctx, in)">`GetConfigState`</SwmToken>

The <SwmToken path="pkg/flare/remote_config.go" pos="61:10:10" line-data="	s, err := cli.GetConfigState(ctx, in)">`GetConfigState`</SwmToken> endpoint is used to retrieve the current state of the configuration repository. This is done by calling the <SwmToken path="pkg/flare/remote_config.go" pos="61:10:10" line-data="	s, err := cli.GetConfigState(ctx, in)">`GetConfigState`</SwmToken> method on the <SwmToken path="pkg/flare/remote_config.go" pos="61:8:8" line-data="	s, err := cli.GetConfigState(ctx, in)">`cli`</SwmToken> object, which is an instance of the secure client obtained from <SwmToken path="pkg/flare/remote_config.go" pos="55:8:10" line-data="	cli, err := agentgrpc.GetDDAgentSecureClient(ctx, ipcAddress, config.GetIPCPort())">`agentgrpc.GetDDAgentSecureClient`</SwmToken>. The response is stored in the variable <SwmToken path="pkg/flare/remote_config.go" pos="61:1:1" line-data="	s, err := cli.GetConfigState(ctx, in)">`s`</SwmToken>.

```go
	s, err := cli.GetConfigState(ctx, in)
	if err != nil {
		return fmt.Errorf("couldn't get the repositories state: %v", err)
	}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/flare/remote_config.go" line="67">

---

### <SwmToken path="pkg/flare/remote_config.go" pos="68:12:12" line-data="		if haState, err = cli.GetConfigStateHA(ctx, in); err != nil {">`GetConfigStateHA`</SwmToken>

The <SwmToken path="pkg/flare/remote_config.go" pos="68:12:12" line-data="		if haState, err = cli.GetConfigStateHA(ctx, in); err != nil {">`GetConfigStateHA`</SwmToken> endpoint is used to retrieve the state of the configuration repository in a high-availability setup. This is conditional on the <SwmToken path="pkg/flare/remote_config.go" pos="67:12:14" line-data="	if config.Datadog().GetBool(&quot;multi_region_failover.enabled&quot;) {">`multi_region_failover.enabled`</SwmToken> configuration being set to true. The method <SwmToken path="pkg/flare/remote_config.go" pos="68:12:12" line-data="		if haState, err = cli.GetConfigStateHA(ctx, in); err != nil {">`GetConfigStateHA`</SwmToken> is called on the <SwmToken path="pkg/flare/remote_config.go" pos="68:10:10" line-data="		if haState, err = cli.GetConfigStateHA(ctx, in); err != nil {">`cli`</SwmToken> object, and the response is stored in the variable <SwmToken path="pkg/flare/remote_config.go" pos="68:3:3" line-data="		if haState, err = cli.GetConfigStateHA(ctx, in); err != nil {">`haState`</SwmToken>.

```go
	if config.Datadog().GetBool("multi_region_failover.enabled") {
		if haState, err = cli.GetConfigStateHA(ctx, in); err != nil {
			return fmt.Errorf("couldn't get the MRF repositories state: %v", err)
		}
	}
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
