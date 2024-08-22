---
title: Adding Host Metadata to Flare Archive
---
This document explains the process of adding host metadata to a flare archive. The process involves converting the metadata into a JSON format and then including it in the archive.

The flow starts with the <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="135:9:9" line-data="func (h *host) fillFlare(fb flaretypes.FlareBuilder) error {">`fillFlare`</SwmToken> function, which is responsible for adding host metadata to a flare archive. It calls the <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="131:9:9" line-data="func (h *host) GetPayloadAsJSON(ctx context.Context) ([]byte, error) {">`GetPayloadAsJSON`</SwmToken> function to convert the metadata into a JSON format. The <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="131:9:9" line-data="func (h *host) GetPayloadAsJSON(ctx context.Context) ([]byte, error) {">`GetPayloadAsJSON`</SwmToken> function, in turn, calls the <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="132:9:9" line-data="	return json.MarshalIndent(h.getPayload(ctx), &quot;&quot;, &quot;    &quot;)">`getPayload`</SwmToken> function to gather all the necessary metadata. The <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="132:9:9" line-data="	return json.MarshalIndent(h.getPayload(ctx), &quot;&quot;, &quot;    &quot;)">`getPayload`</SwmToken> function collects various pieces of information, such as common payload data, resource payload, and optional gohai payload, and returns a complete metadata payload. This payload is then converted into a JSON-formatted byte array by <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="131:9:9" line-data="func (h *host) GetPayloadAsJSON(ctx context.Context) ([]byte, error) {">`GetPayloadAsJSON`</SwmToken> and added to the flare archive by <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="135:9:9" line-data="func (h *host) fillFlare(fb flaretypes.FlareBuilder) error {">`fillFlare`</SwmToken>.

# Flow drill down

```mermaid
graph TD;
      subgraph compmetadatahosthostimpl["comp/metadata/host/hostimpl"]
8af9c381cbdfa24cb23bb8779dbad5eedf8675009585161163a89de049ca84ac(fillFlare):::mainFlowStyle --> e80d17541405e498d3e4eef38bdf52df2e7b0121571c2cfc792c13ecf7a62c2a(GetPayloadAsJSON):::mainFlowStyle
end

subgraph compmetadatahosthostimpl["comp/metadata/host/hostimpl"]
e80d17541405e498d3e4eef38bdf52df2e7b0121571c2cfc792c13ecf7a62c2a(GetPayloadAsJSON):::mainFlowStyle --> 3087b269fe3e8cc008400b0bf3dbef3f1c35acd190dfe9cddcd74cde98552f87(getPayload):::mainFlowStyle
end

subgraph compmetadatahosthostimpl["comp/metadata/host/hostimpl"]
3087b269fe3e8cc008400b0bf3dbef3f1c35acd190dfe9cddcd74cde98552f87(getPayload):::mainFlowStyle --> 903ffcafeef3459956f73bf8fceb0628c871ea10b7c01303db100de307074b34(GetPayload):::mainFlowStyle
end


      classDef mainFlowStyle color:#000000,fill:#7CB9F4
classDef rootsStyle color:#000000,fill:#00FFF4
classDef Style1 color:#000000,fill:#00FFAA
classDef Style2 color:#000000,fill:#FFFF00
classDef Style3 color:#000000,fill:#AA7CB9
```

<SwmSnippet path="/comp/metadata/host/hostimpl/host.go" line="135">

---

## <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="135:9:9" line-data="func (h *host) fillFlare(fb flaretypes.FlareBuilder) error {">`fillFlare`</SwmToken>

The <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="135:9:9" line-data="func (h *host) fillFlare(fb flaretypes.FlareBuilder) error {">`fillFlare`</SwmToken> function is responsible for adding a file to the flare archive. It uses the <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="136:5:5" line-data="	return fb.AddFileFromFunc(filepath.Join(&quot;metadata&quot;, &quot;host.json&quot;), func() ([]byte, error) { return h.GetPayloadAsJSON(context.Background()) })">`AddFileFromFunc`</SwmToken> method to include the host metadata in JSON format by calling <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="136:43:43" line-data="	return fb.AddFileFromFunc(filepath.Join(&quot;metadata&quot;, &quot;host.json&quot;), func() ([]byte, error) { return h.GetPayloadAsJSON(context.Background()) })">`GetPayloadAsJSON`</SwmToken>.

```go
func (h *host) fillFlare(fb flaretypes.FlareBuilder) error {
	return fb.AddFileFromFunc(filepath.Join("metadata", "host.json"), func() ([]byte, error) { return h.GetPayloadAsJSON(context.Background()) })
}
```

---

</SwmSnippet>

<SwmSnippet path="/comp/metadata/host/hostimpl/host.go" line="131">

---

## <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="131:9:9" line-data="func (h *host) GetPayloadAsJSON(ctx context.Context) ([]byte, error) {">`GetPayloadAsJSON`</SwmToken>

The <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="131:9:9" line-data="func (h *host) GetPayloadAsJSON(ctx context.Context) ([]byte, error) {">`GetPayloadAsJSON`</SwmToken> function converts the host metadata payload into a JSON-formatted byte array. It calls the <SwmToken path="comp/metadata/host/hostimpl/host.go" pos="132:9:9" line-data="	return json.MarshalIndent(h.getPayload(ctx), &quot;&quot;, &quot;    &quot;)">`getPayload`</SwmToken> function to retrieve the metadata and then marshals it into JSON.

```go
func (h *host) GetPayloadAsJSON(ctx context.Context) ([]byte, error) {
	return json.MarshalIndent(h.getPayload(ctx), "", "    ")
}
```

---

</SwmSnippet>

<SwmSnippet path="/comp/metadata/host/hostimpl/payload.go" line="44">

---

## <SwmToken path="comp/metadata/host/hostimpl/payload.go" pos="44:2:2" line-data="// getPayload returns the complete metadata payload as seen in Agent v5">`getPayload`</SwmToken>

The <SwmToken path="comp/metadata/host/hostimpl/payload.go" pos="44:2:2" line-data="// getPayload returns the complete metadata payload as seen in Agent v5">`getPayload`</SwmToken> function constructs the complete metadata payload. It gathers various pieces of information such as common payload data, resource payload, and optional gohai payload, and returns a <SwmToken path="comp/metadata/host/hostimpl/payload.go" pos="44:12:12" line-data="// getPayload returns the complete metadata payload as seen in Agent v5">`payload`</SwmToken> object.

```go
// getPayload returns the complete metadata payload as seen in Agent v5
func (h *host) getPayload(ctx context.Context) *Payload {
	p := &Payload{
		CommonPayload: *utils.GetCommonPayload(h.hostname, h.config),
		Payload:       *utils.GetPayload(ctx, h.config),
	}

	if r := h.resources.Get(); r != nil {
		p.ResourcesPayload = r["resources"]
	}

	if h.config.GetBool("enable_gohai") {
		gohaiPayload, err := gohai.GetPayloadAsString(pkgconfig.IsContainerized())
		if err != nil {
			h.log.Errorf("Could not serialize gohai payload: %s", err)
		} else {
			p.GohaiPayload = gohaiPayload
		}
	}
	return p
}
```

---

</SwmSnippet>

<SwmSnippet path="/comp/metadata/host/hostimpl/utils/host.go" line="169">

---

### <SwmToken path="comp/metadata/host/hostimpl/utils/host.go" pos="169:2:2" line-data="// GetPayload builds a metadata payload every time is called.">`GetPayload`</SwmToken>

The <SwmToken path="comp/metadata/host/hostimpl/utils/host.go" pos="169:2:2" line-data="// GetPayload builds a metadata payload every time is called.">`GetPayload`</SwmToken> function builds the metadata payload by collecting data from different sources. It includes hostname data, system stats, host tags, and other metadata, and caches the result for future use.

```go
// GetPayload builds a metadata payload every time is called.
// Some data is collected only once, some is cached, some is collected at every call.
func GetPayload(ctx context.Context, conf config.Reader) *Payload {
	hostnameData, err := hostname.GetWithProvider(ctx)
	if err != nil {
		log.Errorf("Error grabbing hostname for status: %v", err)
		hostnameData = hostname.Data{Hostname: "unknown", Provider: "unknown"}
	}

	meta := GetMeta(ctx, conf)
	meta.Hostname = hostnameData.Hostname

	p := &Payload{
		Os:            osName,
		AgentFlavor:   flavor.GetFlavor(),
		PythonVersion: python.GetPythonInfo(),
		SystemStats:   getSystemStats(),
		Meta:          meta,
		HostTags:      hosttags.Get(ctx, false, conf),
		ContainerMeta: containerMetadata.Get(1 * time.Second),
		NetworkMeta:   getNetworkMeta(ctx),
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
