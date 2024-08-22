---
title: Building the NodeJS Test Data with Docker Compose
---
# Intro

This document explains how Docker is used in the <SwmPath>[pkg/network/protocols/tls/nodejs/testdata/](pkg/network/protocols/tls/nodejs/testdata/)</SwmPath> directory. It will go through the configuration steps in the <SwmPath>[pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml](pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml)</SwmPath> file step by step.

<SwmSnippet path="/pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" line="1">

---

## Version and Name

The <SwmPath>[pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml](pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml)</SwmPath> file starts by specifying the version of the Docker Compose file format and the name of the project. Here, version '3' is used, and the project is named 'node'.

```yaml
version: '3'
name: node
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" line="3">

---

## Services Definition

The <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="3:0:0" line-data="services:">`services`</SwmToken> section defines the services that will be part of this Docker Compose application. In this case, there is a single service named 'node'.

```yaml
services:
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" line="4">

---

## Node Service Configuration

The 'node' service uses the <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="5:4:10" line-data="    image: node:lts-alpine3.19">`node:lts-alpine3.19`</SwmToken> image. The <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="6:1:1" line-data="    command: [&quot;node&quot;, &quot;/v/server.js&quot;]">`command`</SwmToken> specifies that the container should run <SwmPath>[pkg/network/protocols/tls/nodejs/testdata/server.js](pkg/network/protocols/tls/nodejs/testdata/server.js)</SwmPath> when it starts.

```yaml
  node:
    image: node:lts-alpine3.19
    command: ["node", "/v/server.js"]
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" line="7">

---

## Port Mapping

The <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="7:1:1" line-data="    ports:">`ports`</SwmToken> section maps a port on the host machine to a port in the container. Here, the environment variable <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="8:3:6" line-data="      - ${PORT}:4141">`${PORT}`</SwmToken> is mapped to port 4141 in the container.

```yaml
    ports:
      - ${PORT}:4141
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" line="9">

---

## Environment Variables

The <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="9:1:1" line-data="    environment:">`environment`</SwmToken> section sets environment variables inside the container. In this case, <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="10:3:3" line-data="      - ADDR">`ADDR`</SwmToken> and <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="11:3:3" line-data="      - CERTS_DIR">`CERTS_DIR`</SwmToken> are set.

```yaml
    environment:
      - ADDR
      - CERTS_DIR
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" line="12">

---

## Volume Mapping

The <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="12:1:1" line-data="    volumes:">`volumes`</SwmToken> section mounts a directory from the host machine into the container. Here, the environment variable <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="13:3:6" line-data="      - ${TESTDIR}:/v:z">`${TESTDIR}`</SwmToken> is mounted to <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="6:11:12" line-data="    command: [&quot;node&quot;, &quot;/v/server.js&quot;]">`/v`</SwmToken> in the container with the <SwmToken path="pkg/network/protocols/tls/nodejs/testdata/docker-compose.yml" pos="13:9:10" line-data="      - ${TESTDIR}:/v:z">`:z`</SwmToken> option for SELinux compatibility.

```yaml
    volumes:
      - ${TESTDIR}:/v:z
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
