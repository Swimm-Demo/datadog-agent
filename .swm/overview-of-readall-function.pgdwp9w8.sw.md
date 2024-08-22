---
title: Overview of ReadAll Function
---
This document explains the <SwmToken path="pkg/network/protocols/events/batch_reader.go" pos="58:2:2" line-data="// ReadAll batches from eBPF (concurrently) and execute the given">`ReadAll`</SwmToken> function, which is responsible for batching data from <SwmToken path="pkg/network/protocols/events/batch_reader.go" pos="58:8:8" line-data="// ReadAll batches from eBPF (concurrently) and execute the given">`eBPF`</SwmToken> and executing a callback function for each batch. It uses a worker pool to process batches concurrently across multiple <SwmToken path="tasks/gotest.py" pos="247:1:1" line-data="    cpus=None,">`cpus`</SwmToken>, ensuring synchronization and proper termination of goroutines.

The <SwmToken path="pkg/network/protocols/events/batch_reader.go" pos="58:2:2" line-data="// ReadAll batches from eBPF (concurrently) and execute the given">`ReadAll`</SwmToken> function starts by locking to ensure synchronization. It then checks if the process has been stopped. If not, it initializes a wait group to manage multiple <SwmToken path="tasks/gotest.py" pos="247:1:1" line-data="    cpus=None,">`cpus`</SwmToken>. For each CPU, it enqueues a job in the worker pool to process a batch of data. Once all jobs are done, it waits for all goroutines to complete before returning.

# Flow drill down

```mermaid
graph TD;
      subgraph pkgnetwork["pkg/network"]
d1d7307a3ff2a899a181fcd8f271dd8a3ee23fe0daf3727ef39610863bab1cbc(ReadAll):::mainFlowStyle --> 076934ab6ab23dfab6df3ee4beb219940d53d7e880b1974eff249a583745d89b(Lookup):::mainFlowStyle
end

subgraph pkgnetwork["pkg/network"]
076934ab6ab23dfab6df3ee4beb219940d53d7e880b1974eff249a583745d89b(Lookup):::mainFlowStyle --> 5129fc632dc217bfd094f958a7f9074647e928cf8c4a90a503a2fab7094fb557(LookupWithIPs):::mainFlowStyle
end


      classDef mainFlowStyle color:#000000,fill:#7CB9F4
classDef rootsStyle color:#000000,fill:#00FFF4
classDef Style1 color:#000000,fill:#00FFAA
classDef Style2 color:#000000,fill:#FFFF00
classDef Style3 color:#000000,fill:#AA7CB9
```

<SwmSnippet path="/pkg/network/protocols/events/batch_reader.go" line="58">

---

## <SwmToken path="pkg/network/protocols/events/batch_reader.go" pos="58:2:2" line-data="// ReadAll batches from eBPF (concurrently) and execute the given">`ReadAll`</SwmToken>

The <SwmToken path="pkg/network/protocols/events/batch_reader.go" pos="58:2:2" line-data="// ReadAll batches from eBPF (concurrently) and execute the given">`ReadAll`</SwmToken> function is responsible for batching data from <SwmToken path="pkg/network/protocols/events/batch_reader.go" pos="58:8:8" line-data="// ReadAll batches from eBPF (concurrently) and execute the given">`eBPF`</SwmToken> and executing a callback function for each batch. It uses a worker pool to process batches concurrently across multiple <SwmToken path="tasks/gotest.py" pos="247:1:1" line-data="    cpus=None,">`cpus`</SwmToken>. The function ensures synchronization using a lock and waits for all goroutines to complete before returning.

```go
// ReadAll batches from eBPF (concurrently) and execute the given
// callback function for each batch
func (r *batchReader) ReadAll(f func(cpu int, b *batch)) {
	// This lock is used only for the purposes of synchronizing termination
	// and it's only held while *enqueing* the jobs.
	r.Lock()
	if r.stopped {
		r.Unlock()
		return
	}

	var wg sync.WaitGroup
	wg.Add(r.numCPUs)

	for i := 0; i < r.numCPUs; i++ {
		cpu := i // required to properly capture this variable in the function closure
		r.workerPool.Do(func() {
			defer wg.Done()
			batchID, key := r.generateBatchKey(cpu)

			b := batchPool.Get()
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/gateway_lookup_linux.go" line="114">

---

## Lookup

The <SwmToken path="pkg/network/gateway_lookup_linux.go" pos="114:2:2" line-data="// Lookup performs a gateway lookup for connection stats">`Lookup`</SwmToken> function performs a gateway lookup for connection statistics. It determines the destination IP address and then calls <SwmToken path="pkg/network/gateway_lookup_linux.go" pos="121:5:5" line-data="	return g.LookupWithIPs(cs.Source, dest, cs.NetNS)">`LookupWithIPs`</SwmToken> to perform the actual lookup using the source and destination <SwmToken path="tasks/kernel_matrix_testing/ci.py" pos="126:1:1" line-data="        ips: set[str] = set()">`ips`</SwmToken> along with the network namespace.

```go
// Lookup performs a gateway lookup for connection stats
func (g *gatewayLookup) Lookup(cs *ConnectionStats) *Via {
	dest := cs.Dest
	if cs.IPTranslation != nil {
		dest = cs.IPTranslation.ReplSrcIP
	}

	return g.LookupWithIPs(cs.Source, dest, cs.NetNS)
}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/gateway_lookup_linux.go" line="124">

---

## <SwmToken path="pkg/network/gateway_lookup_linux.go" pos="124:2:2" line-data="// LookupWithIPs performs a gateway lookup given the">`LookupWithIPs`</SwmToken>

The <SwmToken path="pkg/network/gateway_lookup_linux.go" pos="124:2:2" line-data="// LookupWithIPs performs a gateway lookup given the">`LookupWithIPs`</SwmToken> function performs a gateway lookup given the source, destination, and namespace. It first checks the route cache for an existing entry. If no gateway is found, it returns nil. Otherwise, it looks up the subnet information, caches the result, and returns the gateway information. The function also handles various error conditions and updates telemetry metrics.

```go
// LookupWithIPs performs a gateway lookup given the
// source, destination, and namespace
func (g *gatewayLookup) LookupWithIPs(source util.Address, dest util.Address, netns uint32) *Via {
	r, ok := g.routeCache.Get(source, dest, netns)
	if !ok {
		return nil
	}

	// if there is no gateway, we don't need to add subnet info
	// for gateway resolution in the backend
	if r.Gateway.IsZero() || r.Gateway.IsUnspecified() {
		return nil
	}

	gatewayLookupTelemetry.subnetCacheLookups.Inc()
	v, ok := g.subnetCache.Get(r.IfIndex)
	if !ok {
		gatewayLookupTelemetry.subnetCacheMisses.Inc()

		var s Subnet
		var err error
```

---

</SwmSnippet>

# Where is this flow used?

This flow is used multiple times in the codebase as represented in the following diagram:

(Note - these are only some of the entry points of this flow)

```mermaid
graph TD;
      2f0941aeb7b33abf9c5268d1def464982656fc1d5d275fa4eefd5d4886a619ac(downloadPolicy):::rootsStyle --> b8e2edbec3bb09f2d4dfc05aa36da0cd894b665d7fd504a2a86ffd1e55b105e5(checkPolicies)

b8e2edbec3bb09f2d4dfc05aa36da0cd894b665d7fd504a2a86ffd1e55b105e5(checkPolicies) --> cdc9b92e40105115448a670413234f51b311bff9a40426b6ad745c6d21c2a42e(checkPoliciesLocal)

cdc9b92e40105115448a670413234f51b311bff9a40426b6ad745c6d21c2a42e(checkPoliciesLocal) --> 7e87265e601842ff16e35c9ddd607ae2ae152ef5b6ee3d4758ffbbc7f430b421(LoadPolicies)

subgraph pkgcompliancek8sconfig["pkg/compliance/k8sconfig"]
7e87265e601842ff16e35c9ddd607ae2ae152ef5b6ee3d4758ffbbc7f430b421(LoadPolicies) --> 118b874f940326fa861845d04d562c0ceee33aafec142b4cac7cfa1e2ff733f7(load)
end

subgraph pkgcompliancek8sconfig["pkg/compliance/k8sconfig"]
118b874f940326fa861845d04d562c0ceee33aafec142b4cac7cfa1e2ff733f7(load) --> c6476b48f7b7c2bb74965fc56d92849e4c08406be8a4bc1d0171d069b75d4908(newK8sKubeletConfig)
end

subgraph pkgcompliancek8sconfig["pkg/compliance/k8sconfig"]
c6476b48f7b7c2bb74965fc56d92849e4c08406be8a4bc1d0171d069b75d4908(newK8sKubeletConfig) --> 10d6f8182d9cd02950b7c2c54b85356caffd081e055ad349b6ee26423e5f034b(loadKubeconfigMeta)
end

subgraph pkgcompliancek8sconfig["pkg/compliance/k8sconfig"]
10d6f8182d9cd02950b7c2c54b85356caffd081e055ad349b6ee26423e5f034b(loadKubeconfigMeta) --> def3eed72fa88af23eade7cd706ec5e58556ebfc8b45943d23881a4cd1e83166(loadCertFileMeta)
end

subgraph pkgcompliancek8sconfig["pkg/compliance/k8sconfig"]
def3eed72fa88af23eade7cd706ec5e58556ebfc8b45943d23881a4cd1e83166(loadCertFileMeta) --> 81bd27540d14c9ed237b9595b1818b30efc49789350eac7ed81d2a099d254c40(loadMeta)
end

81bd27540d14c9ed237b9595b1818b30efc49789350eac7ed81d2a099d254c40(loadMeta) --> d1d7307a3ff2a899a181fcd8f271dd8a3ee23fe0daf3727ef39610863bab1cbc(ReadAll):::mainFlowStyle

2f0941aeb7b33abf9c5268d1def464982656fc1d5d275fa4eefd5d4886a619ac(downloadPolicy):::rootsStyle --> d4eff7cac999d49a6ba6fe57be911ed0d612182b5e17807554142211a2c6a0a3(checkPolicies)

d4eff7cac999d49a6ba6fe57be911ed0d612182b5e17807554142211a2c6a0a3(checkPolicies) --> 1807bd7c0ba446b35fe483d1c325f65a36abb586705480f32f5524eb086e2696(checkPoliciesLocal)

1807bd7c0ba446b35fe483d1c325f65a36abb586705480f32f5524eb086e2696(checkPoliciesLocal) --> 7e87265e601842ff16e35c9ddd607ae2ae152ef5b6ee3d4758ffbbc7f430b421(LoadPolicies)

0391dbe9516912675240da12e2f1531aa5f04d963187480708bc1f99498a0add(evalRule):::rootsStyle --> 7e87265e601842ff16e35c9ddd607ae2ae152ef5b6ee3d4758ffbbc7f430b421(LoadPolicies)

0391dbe9516912675240da12e2f1531aa5f04d963187480708bc1f99498a0add(evalRule):::rootsStyle --> 7e87265e601842ff16e35c9ddd607ae2ae152ef5b6ee3d4758ffbbc7f430b421(LoadPolicies)

subgraph pkgcollectorcorechecksnetworkdevicesciscosdwanclient["pkg/collector/corechecks/network-devices/cisco-sdwan/client"]
d8e0993349c6c674e58105a7d4d4077647a15f77b7dea09b58cb72cd6f837cab(Run):::rootsStyle --> 3e340b5280c13da2835c4caba3132014fc75910d7c4d04e45eeaf2297327a46d(GetDeviceHardwareMetrics)
end

subgraph pkgcollectorcorechecksnetworkdevicesciscosdwanclient["pkg/collector/corechecks/network-devices/cisco-sdwan/client"]
3e340b5280c13da2835c4caba3132014fc75910d7c4d04e45eeaf2297327a46d(GetDeviceHardwareMetrics) --> 1077c372561fabe5ebd3e1ba4bb1a11caf9c83b394cd1a2ba46e19efdd2eaeea(getAllEntries)
end

subgraph pkgcollectorcorechecksnetworkdevicesciscosdwanclient["pkg/collector/corechecks/network-devices/cisco-sdwan/client"]
1077c372561fabe5ebd3e1ba4bb1a11caf9c83b394cd1a2ba46e19efdd2eaeea(getAllEntries) --> edbcb5a0fc92dbcc08cc5a682066244415b41ae08e3d9014d2d210bc9e2790ca(getMoreEntries)
end

subgraph pkgcollectorcorechecksnetworkdevicesciscosdwanclient["pkg/collector/corechecks/network-devices/cisco-sdwan/client"]
edbcb5a0fc92dbcc08cc5a682066244415b41ae08e3d9014d2d210bc9e2790ca(getMoreEntries) --> 096b46eb5cdf1e6da53475d1260b54dc19f895a0a40136827c59470b8a70e7a5(get)
end

subgraph pkgcollectorcorechecksnetworkdevicesciscosdwanclient["pkg/collector/corechecks/network-devices/cisco-sdwan/client"]
096b46eb5cdf1e6da53475d1260b54dc19f895a0a40136827c59470b8a70e7a5(get) --> 625a663a8202862ae4c933cefae2cbd58bd9a8bb29f69501820097f9d9f74f28(authenticate)
end

subgraph pkgcollectorcorechecksnetworkdevicesciscosdwanclient["pkg/collector/corechecks/network-devices/cisco-sdwan/client"]
625a663a8202862ae4c933cefae2cbd58bd9a8bb29f69501820097f9d9f74f28(authenticate) --> 9c90ad94eb1a37a97da5641a3c5db778c29d6d3dceb4b17cf61f807461edc9c7(login)
end

9c90ad94eb1a37a97da5641a3c5db778c29d6d3dceb4b17cf61f807461edc9c7(login) --> d1d7307a3ff2a899a181fcd8f271dd8a3ee23fe0daf3727ef39610863bab1cbc(ReadAll):::mainFlowStyle


      classDef mainFlowStyle color:#000000,fill:#7CB9F4
classDef rootsStyle color:#000000,fill:#00FFF4
classDef Style1 color:#000000,fill:#00FFAA
classDef Style2 color:#000000,fill:#FFFF00
classDef Style3 color:#000000,fill:#AA7CB9
```

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
