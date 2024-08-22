---
title: Docker
---
# Intro

This document explains how Docker is used in the <SwmPath>[pkg/network/protocols/postgres/testdata/](pkg/network/protocols/postgres/testdata/)</SwmPath> directory, specifically focusing on the <SwmPath>[pkg/network/protocols/postgres/testdata/docker-compose.yml](pkg/network/protocols/postgres/testdata/docker-compose.yml)</SwmPath> file.

<SwmSnippet path="/pkg/network/protocols/postgres/testdata/docker-compose.yml" line="1">

---

# Docker Compose Configuration

The <SwmPath>[pkg/network/protocols/postgres/testdata/docker-compose.yml](pkg/network/protocols/postgres/testdata/docker-compose.yml)</SwmPath> file starts by specifying the version of Docker Compose being used (<SwmToken path="pkg/network/protocols/postgres/testdata/docker-compose.yml" pos="1:4:6" line-data="version: &#39;3.1&#39;">`3.1`</SwmToken>) and naming the project <SwmToken path="pkg/network/protocols/postgres/testdata/docker-compose.yml" pos="2:3:3" line-data="name: postgres">`postgres`</SwmToken>.

```yaml
version: '3.1'
name: postgres
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/postgres/testdata/docker-compose.yml" line="3">

---

# Postgres Service Configuration

The <SwmToken path="pkg/network/protocols/postgres/testdata/docker-compose.yml" pos="3:0:0" line-data="services:">`services`</SwmToken> section defines the <SwmToken path="pkg/network/protocols/postgres/testdata/docker-compose.yml" pos="4:1:1" line-data="  postgres:">`postgres`</SwmToken> service, which uses the <SwmToken path="pkg/network/protocols/postgres/testdata/docker-compose.yml" pos="5:4:8" line-data="    image: postgres:15-alpine">`postgres:15-alpine`</SwmToken> image.

```yaml
services:
  postgres:
    image: postgres:15-alpine
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/postgres/testdata/docker-compose.yml" line="6">

---

The <SwmToken path="pkg/network/protocols/postgres/testdata/docker-compose.yml" pos="6:1:1" line-data="    restart: always">`restart`</SwmToken> policy is set to <SwmToken path="pkg/network/protocols/postgres/testdata/docker-compose.yml" pos="6:4:4" line-data="    restart: always">`always`</SwmToken>, ensuring that the container restarts automatically if it stops.

```yaml
    restart: always
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
