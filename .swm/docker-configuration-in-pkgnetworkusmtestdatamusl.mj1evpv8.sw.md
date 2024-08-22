---
title: Docker Configuration in pkg/network/usm/testdata/musl
---
# Intro

This document explains how Docker is used in the <SwmPath>[pkg/network/usm/testdata/musl/](pkg/network/usm/testdata/musl/)</SwmPath> directory. It will go through the configuration steps in the <SwmPath>[pkg/network/usm/testdata/musl/docker-compose.yml](pkg/network/usm/testdata/musl/docker-compose.yml)</SwmPath> file.

<SwmSnippet path="/pkg/network/usm/testdata/musl/docker-compose.yml" line="1">

---

## Docker-compose configuration

The <SwmPath>[pkg/network/usm/testdata/musl/docker-compose.yml](pkg/network/usm/testdata/musl/docker-compose.yml)</SwmPath> file starts by specifying the version of Docker Compose being used. In this case, it is version 3.

```yaml
version: '3'
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/usm/testdata/musl/docker-compose.yml" line="2">

---

## Service Configuration

The services section defines the services that will be run by Docker Compose. Here, a single service named <SwmToken path="pkg/network/usm/testdata/musl/docker-compose.yml" pos="3:1:1" line-data="  alpine:">`alpine`</SwmToken> is defined.

```yaml
services:
  alpine:
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/usm/testdata/musl/docker-compose.yml" line="4">

---

## Image Configuration

The <SwmToken path="pkg/network/usm/testdata/musl/docker-compose.yml" pos="4:6:6" line-data="    image: byrnedo/alpine-curl:3.17">`alpine`</SwmToken> service uses the <SwmToken path="pkg/network/usm/testdata/musl/docker-compose.yml" pos="4:4:12" line-data="    image: byrnedo/alpine-curl:3.17">`byrnedo/alpine-curl:3.17`</SwmToken> image. This image is based on Alpine Linux and includes the <SwmToken path="pkg/network/usm/testdata/musl/docker-compose.yml" pos="4:8:8" line-data="    image: byrnedo/alpine-curl:3.17">`curl`</SwmToken> utility.

```yaml
    image: byrnedo/alpine-curl:3.17
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/usm/testdata/musl/docker-compose.yml" line="5">

---

## Interactive Terminal

The <SwmToken path="pkg/network/usm/testdata/musl/docker-compose.yml" pos="5:1:1" line-data="    stdin_open: true">`stdin_open`</SwmToken> and <SwmToken path="pkg/network/usm/testdata/musl/docker-compose.yml" pos="6:1:1" line-data="    tty: true">`tty`</SwmToken> options are set to `true`, which allows the container to run with an interactive terminal. This is useful for debugging and running commands interactively.

```yaml
    stdin_open: true
    tty: true
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/usm/testdata/musl/docker-compose.yml" line="9">

---

## Entrypoint Configuration

The <SwmToken path="pkg/network/usm/testdata/musl/docker-compose.yml" pos="9:1:1" line-data="    entrypoint: /bin/sh -c &quot;echo started &amp;&amp; /bin/cat&quot;">`entrypoint`</SwmToken> is set to <SwmToken path="pkg/network/usm/testdata/musl/docker-compose.yml" pos="9:4:23" line-data="    entrypoint: /bin/sh -c &quot;echo started &amp;&amp; /bin/cat&quot;">`/bin/sh -c "echo started && /bin/cat"`</SwmToken>. This command will run a shell that prints <SwmToken path="pkg/network/usm/testdata/musl/docker-compose.yml" pos="9:15:15" line-data="    entrypoint: /bin/sh -c &quot;echo started &amp;&amp; /bin/cat&quot;">`started`</SwmToken> and then waits for input, effectively keeping the container running.

```yaml
    entrypoint: /bin/sh -c "echo started && /bin/cat"
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
