---
title: Exploring DNS Resolution
---
```mermaid
graph TD;
 A[DNS Query] --> B[DNS Cache];
 B --> C[Resolve Method];
 C --> D[DNS Hostname];
 B --> E[Generate BPF Filter];
 E --> F[Capture DNS Packets];

 graph TD;
 A[NewReverseDNS] --> B[Initialize eBPF Program];
 B --> C[Create Packet Source];
 C --> D[Start Monitoring DNS Traffic];
 D --> E[WaitForDomain];
 E --> F[Detect Specified Domain];

%% Swimm:
%% graph TD;
%%  A[DNS Query] --> B[DNS Cache];
%%  B --> C[Resolve Method];
%%  C --> D[DNS Hostname];
%%  B --> E[Generate BPF Filter];
%%  E --> F[Capture DNS Packets];
%% 
%%  graph TD;
%%  A[NewReverseDNS] --> B[Initialize <SwmToken path="pkg/network/dns/snooper.go" pos="46:24:24" line-data="// socketFilterSnooper is a DNS traffic snooper built on top of an eBPF SOCKET_FILTER">`eBPF`</SwmToken> Program];
%%  B --> C[Create Packet Source];
%%  C --> D[Start Monitoring DNS Traffic];
%%  D --> E[WaitForDomain];
%%  E --> F[Detect Specified Domain];
```

# Overview

DNS resolution is the process of translating domain names into IP addresses and vice versa. This is essential for network communication, allowing devices and services to be identified by human-readable names rather than numerical IP addresses.

# DNS Package Components

The <SwmToken path="pkg/network/dns/snooper.go" pos="120:8:8" line-data="// Resolve IPs to DNS addresses">`DNS`</SwmToken> package includes various components such as <SwmPath>[pkg/network/dns/tcp_dns_layer.go](pkg/network/dns/tcp_dns_layer.go)</SwmPath>, <SwmPath>[comp/core/autodiscovery/autodiscoveryimpl/stats.go](comp/core/autodiscovery/autodiscoveryimpl/stats.go)</SwmPath>, <SwmPath>[pkg/network/dns/snooper.go](pkg/network/dns/snooper.go)</SwmPath>, and <SwmPath>[pkg/network/dns/cache.go](pkg/network/dns/cache.go)</SwmPath>, which handle different aspects of DNS resolution and caching. These components work together to monitor and resolve DNS queries efficiently.

# DNS Resolution Process

The <SwmToken path="pkg/network/dns/snooper.go" pos="120:2:2" line-data="// Resolve IPs to DNS addresses">`Resolve`</SwmToken> method in <SwmPath>[pkg/network/dns/snooper.go](pkg/network/dns/snooper.go)</SwmPath> is used to map IP addresses to DNS hostnames by querying the DNS cache. This method ensures that the DNS resolution process is optimized by leveraging cached data.

<SwmSnippet path="/pkg/network/dns/snooper.go" line="120">

---

The <SwmToken path="pkg/network/dns/snooper.go" pos="120:2:2" line-data="// Resolve IPs to DNS addresses">`Resolve`</SwmToken> function maps IP addresses to DNS hostnames by querying the DNS cache, optimizing the DNS resolution process.

```go
// Resolve IPs to DNS addresses
func (s *socketFilterSnooper) Resolve(ips map[util.Address]struct{}) map[util.Address][]Hostname {
	return s.cache.Get(ips)
}
```

---

</SwmSnippet>

# DNS Caching

The <SwmToken path="pkg/network/dns/cache.go" pos="19:2:2" line-data="const dnsCacheModuleName = &quot;network_tracer__dns_cache&quot;">`dnsCacheModuleName`</SwmToken> constant in <SwmPath>[pkg/network/dns/cache.go](pkg/network/dns/cache.go)</SwmPath> indicates the module responsible for DNS caching, which helps in reducing the latency of DNS lookups by storing previously resolved addresses.

<SwmSnippet path="/pkg/network/dns/cache.go" line="19">

---

The <SwmToken path="pkg/network/dns/cache.go" pos="19:2:2" line-data="const dnsCacheModuleName = &quot;network_tracer__dns_cache&quot;">`dnsCacheModuleName`</SwmToken> constant indicates the module responsible for DNS caching, reducing the latency of DNS lookups.

```go
const dnsCacheModuleName = "network_tracer__dns_cache"

// Telemetry
```

---

</SwmSnippet>

# Capturing DNS Packets

The <SwmToken path="pkg/network/dns/bpf.go" pos="20:2:2" line-data="func generateBPFFilter(c *config.Config) ([]bpf.RawInstruction, error) {">`generateBPFFilter`</SwmToken> function in <SwmPath>[pkg/network/dns/bpf.go](pkg/network/dns/bpf.go)</SwmPath> creates a BPF filter to capture DNS packets based on specific criteria such as port numbers and protocols. This is essential for monitoring DNS traffic and collecting relevant statistics.

<SwmSnippet path="/pkg/network/dns/bpf.go" line="20">

---

The <SwmToken path="pkg/network/dns/bpf.go" pos="20:2:2" line-data="func generateBPFFilter(c *config.Config) ([]bpf.RawInstruction, error) {">`generateBPFFilter`</SwmToken> function creates a BPF filter to capture DNS packets based on specific criteria, essential for monitoring DNS traffic.

```go
func generateBPFFilter(c *config.Config) ([]bpf.RawInstruction, error) {
	allowedDestPort := uint32(math.MaxUint32)
	if c.CollectDNSStats {
		allowedDestPort = port
	}

	return bpf.Assemble([]bpf.Instruction{
		//(000) ldh      [12] -- load Ethertype
		bpf.LoadAbsolute{Size: 2, Off: 12},
		//(001) jeq      #0x86dd          jt 2	jf 9 -- if IPv6, goto 2, else 9
		bpf.JumpIf{Cond: bpf.JumpEqual, Val: 0x86dd, SkipTrue: 0, SkipFalse: 7},
		//(002) ldb      [20] -- load IPv6 Next Header
		bpf.LoadAbsolute{Size: 1, Off: 20},
		//(003) jeq      #0x6             jt 5	jf 4 -- IPv6 Next Header: if TCP, goto 5
		bpf.JumpIf{Cond: bpf.JumpEqual, Val: 0x6, SkipTrue: 1, SkipFalse: 0},
		//(004) jeq      #0x11            jt 5	jf 21 -- IPv6 Next Header: if UDP, goto 5, else drop
		bpf.JumpIf{Cond: bpf.JumpEqual, Val: 0x11, SkipTrue: 0, SkipFalse: 16},
		//(005) ldh      [54] -- load source port
		bpf.LoadAbsolute{Size: 2, Off: 54},
		//(006) jeq      #0x35            jt 20	jf 7 -- if 53, capture
		bpf.JumpIf{Cond: bpf.JumpEqual, Val: port, SkipTrue: 13, SkipFalse: 0},
```

---

</SwmSnippet>

# Main Functions

Several main functions in the DNS package include <SwmToken path="pkg/network/dns/snooper.go" pos="120:2:2" line-data="// Resolve IPs to DNS addresses">`Resolve`</SwmToken>, <SwmToken path="pkg/network/dns/monitor_linux.go" pos="34:2:2" line-data="// NewReverseDNS starts snooping on DNS traffic to allow IP -&gt; domain reverse resolution">`NewReverseDNS`</SwmToken>, and <SwmToken path="pkg/network/dns/monitor_linux.go" pos="97:9:9" line-data="func (m *dnsMonitor) WaitForDomain(domain string) error {">`WaitForDomain`</SwmToken>. We will explore <SwmToken path="pkg/network/dns/snooper.go" pos="120:2:2" line-data="// Resolve IPs to DNS addresses">`Resolve`</SwmToken> and <SwmToken path="pkg/network/dns/monitor_linux.go" pos="34:2:2" line-data="// NewReverseDNS starts snooping on DNS traffic to allow IP -&gt; domain reverse resolution">`NewReverseDNS`</SwmToken> in detail.

## Resolve

The <SwmToken path="pkg/network/dns/snooper.go" pos="120:2:2" line-data="// Resolve IPs to DNS addresses">`Resolve`</SwmToken> function is used to map IP addresses to DNS hostnames by querying the DNS cache. This method ensures that the DNS resolution process is optimized by leveraging cached data.

## <SwmToken path="pkg/network/dns/monitor_linux.go" pos="34:2:2" line-data="// NewReverseDNS starts snooping on DNS traffic to allow IP -&gt; domain reverse resolution">`NewReverseDNS`</SwmToken>

The <SwmToken path="pkg/network/dns/monitor_linux.go" pos="34:2:2" line-data="// NewReverseDNS starts snooping on DNS traffic to allow IP -&gt; domain reverse resolution">`NewReverseDNS`</SwmToken> function starts snooping on DNS traffic to allow IP to domain reverse resolution. It initializes the necessary <SwmToken path="pkg/network/dns/snooper.go" pos="46:24:24" line-data="// socketFilterSnooper is a DNS traffic snooper built on top of an eBPF SOCKET_FILTER">`eBPF`</SwmToken> programs and packet sources to capture DNS packets and process them.

<SwmSnippet path="/pkg/network/dns/monitor_linux.go" line="34">

---

The <SwmToken path="pkg/network/dns/monitor_linux.go" pos="34:2:2" line-data="// NewReverseDNS starts snooping on DNS traffic to allow IP -&gt; domain reverse resolution">`NewReverseDNS`</SwmToken> function initializes a new DNS monitor that snoops on DNS traffic to allow IP to domain reverse resolution, setting up the necessary <SwmToken path="pkg/network/dns/snooper.go" pos="46:24:24" line-data="// socketFilterSnooper is a DNS traffic snooper built on top of an eBPF SOCKET_FILTER">`eBPF`</SwmToken> programs and socket filters.

```go
// NewReverseDNS starts snooping on DNS traffic to allow IP -> domain reverse resolution
func NewReverseDNS(cfg *config.Config, _ telemetry.Component) (ReverseDNS, error) {
	currKernelVersion, err := kernel.HostVersion()
	if err != nil {
		// if the platform couldn't be determined, treat it as new kernel case
		log.Warn("could not detect the platform, will use kprobes from kernel version >= 4.1.0")
		currKernelVersion = math.MaxUint32
	}
	pre410Kernel := currKernelVersion < kernel.VersionCode(4, 1, 0)

	var p *ebpfProgram
	var filter *manager.Probe
	var bpfFilter []bpf.RawInstruction
	if pre410Kernel {
		bpfFilter, err = generateBPFFilter(cfg)
		if err != nil {
			return nil, fmt.Errorf("error creating bpf classic filter: %w", err)
		}
	} else {
		p, err = newEBPFProgram(cfg)
		if err != nil {
```

---

</SwmSnippet>

## <SwmToken path="pkg/network/dns/monitor_linux.go" pos="97:9:9" line-data="func (m *dnsMonitor) WaitForDomain(domain string) error {">`WaitForDomain`</SwmToken>

The <SwmToken path="pkg/network/dns/monitor_linux.go" pos="97:9:9" line-data="func (m *dnsMonitor) WaitForDomain(domain string) error {">`WaitForDomain`</SwmToken> function waits for a specific domain to appear in the DNS traffic. It uses the <SwmToken path="pkg/network/dns/monitor_linux.go" pos="98:5:5" line-data="	return m.statKeeper.WaitForDomain(domain)">`statKeeper`</SwmToken> to monitor and detect the specified domain.

<SwmSnippet path="/pkg/network/dns/monitor_linux.go" line="97">

---

The <SwmToken path="pkg/network/dns/monitor_linux.go" pos="97:9:9" line-data="func (m *dnsMonitor) WaitForDomain(domain string) error {">`WaitForDomain`</SwmToken> function waits for a specific domain to appear in the DNS traffic, using the <SwmToken path="pkg/network/dns/monitor_linux.go" pos="98:5:5" line-data="	return m.statKeeper.WaitForDomain(domain)">`statKeeper`</SwmToken> to monitor and detect the specified domain.

```go
func (m *dnsMonitor) WaitForDomain(domain string) error {
	return m.statKeeper.WaitForDomain(domain)
}
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
