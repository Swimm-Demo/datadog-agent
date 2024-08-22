---
title: Handling Network Namespace Cache Dump Requests
---
In this document, we will explain the process of handling network namespace cache dump requests. The process involves creating a runtime security client, sending a dump request, processing the request, and creating dump and graph files.

The flow starts by creating a runtime security client and sending a request to dump the network namespace cache. If the request is successful, the server processes it by delegating to the <SwmToken path="pkg/security/module/server_linux.go" pos="193:7:7" line-data="	return p.Resolvers.NamespaceResolver.DumpNetworkNamespaces(params), nil">`NamespaceResolver`</SwmToken>. The resolver then creates temporary files for the dump and graph, encodes the dump to JSON, generates a dot graph, and handles any errors that occur during these processes. Finally, the filenames of the dump and graph are printed.

# Flow drill down

```mermaid
graph TD;
      subgraph pkgsecurity["pkg/security"]
720a16366f90197428cbfe673cc4b6fa6ecca976eda26f60c5ad0fce43e7a65c(dumpNetworkNamespace):::mainFlowStyle --> b7c4339d463227708f74641101525601362d2f220778975429f3b866c07ad4b3(DumpNetworkNamespace):::mainFlowStyle
end

subgraph pkgsecurity["pkg/security"]
b7c4339d463227708f74641101525601362d2f220778975429f3b866c07ad4b3(DumpNetworkNamespace):::mainFlowStyle --> eb9cd1adfc562e92ecb69e38d268c3deb24fcc05b5973673ce674787aa6124fe(DumpNetworkNamespaces):::mainFlowStyle
end


      classDef mainFlowStyle color:#000000,fill:#7CB9F4
classDef rootsStyle color:#000000,fill:#00FFF4
classDef Style1 color:#000000,fill:#00FFAA
classDef Style2 color:#000000,fill:#FFFF00
classDef Style3 color:#000000,fill:#AA7CB9
```

<SwmSnippet path="/cmd/system-probe/subcommands/runtime/command.go" line="336">

---

## Handling network namespace cache dump requests

The function <SwmToken path="cmd/system-probe/subcommands/runtime/command.go" pos="336:2:2" line-data="func dumpNetworkNamespace(_ log.Component, _ config.Component, _ secrets.Component, dumpNetworkNamespaceArgs *dumpNetworkNamespaceCliParams) error {">`dumpNetworkNamespace`</SwmToken> initiates the process by creating a runtime security client and sending a network namespace cache dump request. It handles errors related to client creation and request sending, and prints the dump and graph filenames upon success.

```go
func dumpNetworkNamespace(_ log.Component, _ config.Component, _ secrets.Component, dumpNetworkNamespaceArgs *dumpNetworkNamespaceCliParams) error {
	client, err := secagent.NewRuntimeSecurityClient()
	if err != nil {
		return fmt.Errorf("unable to create a runtime security client instance: %w", err)
	}
	defer client.Close()

	resp, err := client.DumpNetworkNamespace(dumpNetworkNamespaceArgs.snapshotInterfaces)
	if err != nil {
		return fmt.Errorf("couldn't send network namespace cache dump request: %w", err)
	}

	if len(resp.GetError()) > 0 {
		return fmt.Errorf("couldn't dump network namespaces: %w", err)
	}

	fmt.Printf("Network namespace dump: %s\n", resp.GetDumpFilename())
	fmt.Printf("Network namespace dump graph: %s\n", resp.GetGraphFilename())
	return nil
}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/security/module/server_linux.go" line="187">

---

## Processing the dump request

The method <SwmToken path="pkg/security/module/server_linux.go" pos="187:9:9" line-data="func (a *APIServer) DumpNetworkNamespace(_ context.Context, params *api.DumpNetworkNamespaceParams) (*api.DumpNetworkNamespaceMessage, error) {">`DumpNetworkNamespace`</SwmToken> in <SwmToken path="pkg/security/module/server_linux.go" pos="187:6:6" line-data="func (a *APIServer) DumpNetworkNamespace(_ context.Context, params *api.DumpNetworkNamespaceParams) (*api.DumpNetworkNamespaceMessage, error) {">`APIServer`</SwmToken> processes the network namespace cache dump request by delegating it to the <SwmToken path="pkg/security/module/server_linux.go" pos="193:7:7" line-data="	return p.Resolvers.NamespaceResolver.DumpNetworkNamespaces(params), nil">`NamespaceResolver`</SwmToken>. It checks if the platform probe is supported and returns the dump result.

```go
func (a *APIServer) DumpNetworkNamespace(_ context.Context, params *api.DumpNetworkNamespaceParams) (*api.DumpNetworkNamespaceMessage, error) {
	p, ok := a.probe.PlatformProbe.(*probe.EBPFProbe)
	if !ok {
		return nil, fmt.Errorf("not supported")
	}

	return p.Resolvers.NamespaceResolver.DumpNetworkNamespaces(params), nil
}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/security/resolvers/netns/resolver.go" line="651">

---

### Creating dump and graph files

The method <SwmToken path="pkg/security/resolvers/netns/resolver.go" pos="651:2:2" line-data="// DumpNetworkNamespaces dumps the network namespaces held by the namespace resolver">`DumpNetworkNamespaces`</SwmToken> in <SwmToken path="pkg/security/resolvers/netns/resolver.go" pos="651:20:20" line-data="// DumpNetworkNamespaces dumps the network namespaces held by the namespace resolver">`resolver`</SwmToken> handles the actual dumping of network namespaces. It creates temporary files for the dump and graph, encodes the dump to JSON, generates a dot graph, and handles any errors that occur during these processes.

```go
// DumpNetworkNamespaces dumps the network namespaces held by the namespace resolver
func (nr *Resolver) DumpNetworkNamespaces(params *api.DumpNetworkNamespaceParams) *api.DumpNetworkNamespaceMessage {
	resp := &api.DumpNetworkNamespaceMessage{}
	dump := nr.dump(params)

	// create the dump file
	dumpFile, err := newTmpFile("network-namespace-dump-*.json")
	if err != nil {
		resp.Error = fmt.Sprintf("couldn't create temporary file: %v", err)
		seclog.Warnf(resp.Error)
		return resp
	}
	defer dumpFile.Close()
	resp.DumpFilename = dumpFile.Name()

	// dump to JSON file
	encoder := json.NewEncoder(dumpFile)
	if err = encoder.Encode(dump); err != nil {
		resp.Error = fmt.Sprintf("couldn't encode list of network namespace: %v", err)
		seclog.Warnf(resp.Error)
		return resp
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
