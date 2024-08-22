---
title: Listen Functionality Overview
---
This document explains the 'Listen' functionality, which is crucial for initiating the intake loop for the <SwmToken path="comp/dogstatsd/listeners/uds_stream.go" pos="67:6:6" line-data="func (l *UDSStreamListener) Listen() {">`UDSStreamListener`</SwmToken>. It describes how 'Listen' operates, including starting the connection tracker and handling incoming connections.

The 'Listen' functionality starts by initiating the intake loop for the <SwmToken path="comp/dogstatsd/listeners/uds_stream.go" pos="67:6:6" line-data="func (l *UDSStreamListener) Listen() {">`UDSStreamListener`</SwmToken>. This process runs in its own separate thread. Once started, it begins to accept incoming Unix connections. For each connection, a new thread is created to handle it. This ensures that multiple connections can be managed simultaneously without blocking the main process. The connection tracker is also started to keep track of all active connections.

# Flow drill down

```mermaid
graph TD;
      subgraph compdogstatsdlisteners["comp/dogstatsd/listeners"]
7e7a973f13afccd57ccb2e129cf12a1a8187eef689f11039f1e437c5acd11632(Listen):::mainFlowStyle --> c5febf5855bc9f3c4f7e7f62b60da50a1821e0c0987de509d8814a7f0a3a8ef6(listen):::mainFlowStyle
end

subgraph compdogstatsdlisteners["comp/dogstatsd/listeners"]
c5febf5855bc9f3c4f7e7f62b60da50a1821e0c0987de509d8814a7f0a3a8ef6(listen):::mainFlowStyle --> 9d8b392773aa1f35c00f8d5118c008d0b944cf59548d59fd5783d6556e0841f3(handleConnection):::mainFlowStyle
end

subgraph compdogstatsdpacketsbuffergo["comp/dogstatsd/packets/buffer.go"]
9d8b392773aa1f35c00f8d5118c008d0b944cf59548d59fd5783d6556e0841f3(handleConnection):::mainFlowStyle --> e7eeba59ec71d19cf0d7eed7ed31a71a15935c1f8157ef08648cc1bef6aa9012(NewBuffer)
end

subgraph compdogstatsdlisteners["comp/dogstatsd/listeners"]
9d8b392773aa1f35c00f8d5118c008d0b944cf59548d59fd5783d6556e0841f3(handleConnection):::mainFlowStyle --> 1e25234225a9ad9c1fc7f342f1b84eb0cb28c0b58e09da3f2698fdd68eb00e53(BuildMemBasedRateLimiter)
end

subgraph compdogstatsdpacketsbuffergo["comp/dogstatsd/packets/buffer.go"]
9d8b392773aa1f35c00f8d5118c008d0b944cf59548d59fd5783d6556e0841f3(handleConnection):::mainFlowStyle --> 3b655feecf68724512fb40c5625a86ade79849ee0b2abfeb794210b4631b7892(Append):::mainFlowStyle
end

subgraph compdogstatsdpacketsbuffergo["comp/dogstatsd/packets/buffer.go"]
3b655feecf68724512fb40c5625a86ade79849ee0b2abfeb794210b4631b7892(Append):::mainFlowStyle --> c29ad8d1eac7d4b10fec07ecade427d6bda819b5f391940c38fb6ec1ff5457af(flush):::mainFlowStyle
end

1e25234225a9ad9c1fc7f342f1b84eb0cb28c0b58e09da3f2698fdd68eb00e53(BuildMemBasedRateLimiter) --> a1f7fb2be87c281ef9fae7874e5210d631b5e88a7730fe6f36903694e1d0be72(make)

a1f7fb2be87c281ef9fae7874e5210d631b5e88a7730fe6f36903694e1d0be72(make) --> 6065ff794d55a7a0c679bef8dcba778c98ea1d8dad1f3857f74439581b936d21(clear_cmake_cache)

a1f7fb2be87c281ef9fae7874e5210d631b5e88a7730fe6f36903694e1d0be72(make) --> 623dfead172f1b4318a325c2d05ba841a1a749721395ee03d3430ad4123b8f2a(run_make_command)

a1f7fb2be87c281ef9fae7874e5210d631b5e88a7730fe6f36903694e1d0be72(make) --> 5dd57fd7b09ac752033eed1235d0a59b195a71880bf845b67c8f95089d35c59a(run)

subgraph compdogstatsdpacketsbuffergo["comp/dogstatsd/packets/buffer.go"]
e7eeba59ec71d19cf0d7eed7ed31a71a15935c1f8157ef08648cc1bef6aa9012(NewBuffer) --> d2a29bf69c97576dfb20857557f1ad5c0f93da61af17eefc48cc17e14b545a49(flushLoop)
end

e7eeba59ec71d19cf0d7eed7ed31a71a15935c1f8157ef08648cc1bef6aa9012(NewBuffer) --> a1f7fb2be87c281ef9fae7874e5210d631b5e88a7730fe6f36903694e1d0be72(make)


      classDef mainFlowStyle color:#000000,fill:#7CB9F4
classDef rootsStyle color:#000000,fill:#00FFF4
classDef Style1 color:#000000,fill:#00FFAA
classDef Style2 color:#000000,fill:#FFFF00
classDef Style3 color:#000000,fill:#AA7CB9
```

<SwmSnippet path="/comp/dogstatsd/listeners/uds_stream.go" line="66">

---

## Listen

The <SwmToken path="comp/dogstatsd/listeners/uds_stream.go" pos="66:2:2" line-data="// Listen runs the intake loop. Should be called in its own goroutine">`Listen`</SwmToken> function initiates the intake loop for the <SwmToken path="comp/dogstatsd/listeners/uds_stream.go" pos="67:6:6" line-data="func (l *UDSStreamListener) Listen() {">`UDSStreamListener`</SwmToken>. It runs in its own goroutine and calls the <SwmToken path="comp/dogstatsd/listeners/uds_stream.go" pos="66:2:2" line-data="// Listen runs the intake loop. Should be called in its own goroutine">`Listen`</SwmToken> function to start accepting connections.

```go
// Listen runs the intake loop. Should be called in its own goroutine
func (l *UDSStreamListener) Listen() {
	l.listenWg.Add(1)
	go func() {
		defer l.listenWg.Done()
		l.listen()
	}()
}
```

---

</SwmSnippet>

<SwmSnippet path="/comp/dogstatsd/listeners/uds_stream.go" line="75">

---

## listen

The <SwmToken path="comp/dogstatsd/listeners/uds_stream.go" pos="75:9:9" line-data="func (l *UDSStreamListener) listen() {">`listen`</SwmToken> function starts the connection tracker and enters a loop to accept incoming Unix connections. For each accepted connection, it spawns a new goroutine to handle the connection using the <SwmToken path="comp/dogstatsd/listeners/uds_stream.go" pos="89:7:7" line-data="			_ = l.handleConnection(conn, func(c *net.UnixConn) error {">`handleConnection`</SwmToken> function.

```go
func (l *UDSStreamListener) listen() {

	l.connTracker.Start()
	log.Infof("dogstatsd-uds-stream: starting to listen on %s", l.conn.Addr())
	for {
		conn, err := l.conn.AcceptUnix()
		if err != nil {
			if !strings.HasSuffix(err.Error(), " use of closed network connection") {
				log.Errorf("dogstatsd-uds: error accepting connection: %v", err)
			}
			break
		}
		go func() {
			l.connTracker.Track(conn)
			_ = l.handleConnection(conn, func(c *net.UnixConn) error {
				l.connTracker.Close(c)
				return nil
			})
			if err != nil {
				log.Errorf("dogstatsd-uds-stream: error handling connection: %v", err)
			}
```

---

</SwmSnippet>

# Where is this flow used?

This flow is used multiple times in the codebase as represented in the following diagram:

(Note - these are only some of the entry points of this flow)

```mermaid
graph TD;
      subgraph compotelcolextensionimpl["comp/otelcol/extension/impl"]
940ce0ca3c7183c4182fe22e25b91f768a5bf075a3651564af70308073610b2f(CreateExtension):::rootsStyle --> 5a834b92196659d6ef7f3a547d90f933a907c9a112ffa91f1d114b27a37ae535(NewExtension)
end

subgraph compotelcolextensionimpl["comp/otelcol/extension/impl"]
5a834b92196659d6ef7f3a547d90f933a907c9a112ffa91f1d114b27a37ae535(NewExtension) --> a481d20997d51bf498cad908e69743efa093df1d522d510b3c61653e01357645(buildHTTPServer)
end

a481d20997d51bf498cad908e69743efa093df1d522d510b3c61653e01357645(buildHTTPServer) --> 7e7a973f13afccd57ccb2e129cf12a1a8187eef689f11039f1e437c5acd11632(Listen):::mainFlowStyle

9b4c4c177dd31c7e2e8be90f4fcf04b527919ab45e1595c31e42104cb7e31681(newTrapListener):::rootsStyle --> 634c56f06a614ad3668ffb81cf38dca63943ae71dc7819b8047a4c18e03f0758(start)

634c56f06a614ad3668ffb81cf38dca63943ae71dc7819b8047a4c18e03f0758(start) --> 94e9eacd866cf31c6f762d25a189558aea6758b218e8caf844698396da6c9f1b(run)

94e9eacd866cf31c6f762d25a189558aea6758b218e8caf844698396da6c9f1b(run) --> 7e7a973f13afccd57ccb2e129cf12a1a8187eef689f11039f1e437c5acd11632(Listen):::mainFlowStyle

subgraph pkgnetworkprotocolshttpgotlslookupinternaltestprogramprogramgo["pkg/network/protocols/http/gotls/lookup/internal/testprogram/program.go"]
227eaee48f790cb276fab7c3bb5c4aed07dc34a153c81931559b3ec413b1cb32(main):::rootsStyle --> 2b6bb518510a1f6262e78849b61924dcb73904c9a3844454236f8bada4523e17(run)
end

subgraph pkgnetworkprotocolshttpgotlslookupinternaltestprogramprogramgo["pkg/network/protocols/http/gotls/lookup/internal/testprogram/program.go"]
2b6bb518510a1f6262e78849b61924dcb73904c9a3844454236f8bada4523e17(run) --> b4f17448b949d5d4c447bdd07d7ee966d7e39b0fdf11be203f95a05cc06522d5(createTCPListener)
end

b4f17448b949d5d4c447bdd07d7ee966d7e39b0fdf11be203f95a05cc06522d5(createTCPListener) --> 7e7a973f13afccd57ccb2e129cf12a1a8187eef689f11039f1e437c5acd11632(Listen):::mainFlowStyle

subgraph pkgeventmonitor["pkg/eventmonitor"]
f86c986e2224ae116c5ed70a12caf5b16befef99e23af5e400bcbbaeed01ea08(OpenMsiLog):::rootsStyle --> 7f51827570143625203991a40f0b04a9fb5d621d9684d7cfd8a5bd2bff96a55c(Start)
end

subgraph pkgeventmonitor["pkg/eventmonitor"]
7f51827570143625203991a40f0b04a9fb5d621d9684d7cfd8a5bd2bff96a55c(Start) --> 6435a88d2356dc8d2d8947d14e104b987e828a98b06e01d9c6de2296e582eea2(getListener)
end

6435a88d2356dc8d2d8947d14e104b987e828a98b06e01d9c6de2296e582eea2(getListener) --> 7e7a973f13afccd57ccb2e129cf12a1a8187eef689f11039f1e437c5acd11632(Listen):::mainFlowStyle

subgraph pkgeventmonitor["pkg/eventmonitor"]
73869e0db9797f1e7329fc5769f02cd87b36f1855957ea53e3fc57ecbdb2952b(StartService):::rootsStyle --> 7f51827570143625203991a40f0b04a9fb5d621d9684d7cfd8a5bd2bff96a55c(Start)
end


      classDef mainFlowStyle color:#000000,fill:#7CB9F4
classDef rootsStyle color:#000000,fill:#00FFF4
classDef Style1 color:#000000,fill:#00FFAA
classDef Style2 color:#000000,fill:#FFFF00
classDef Style3 color:#000000,fill:#AA7CB9
```

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
