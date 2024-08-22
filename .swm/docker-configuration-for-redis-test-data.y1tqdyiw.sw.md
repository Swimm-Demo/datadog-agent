---
title: Docker Configuration for Redis Test Data
---
# Intro

This document explains how Docker is used in the <SwmPath>[pkg/network/protocols/redis/testdata/](pkg/network/protocols/redis/testdata/)</SwmPath> directory. It will go through the configuration steps in the <SwmPath>[pkg/network/protocols/redis/testdata/docker-compose.yml](pkg/network/protocols/redis/testdata/docker-compose.yml)</SwmPath> file.

<SwmSnippet path="/pkg/network/protocols/redis/testdata/docker-compose.yml" line="1">

---

## Version and Name

The <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="1:0:0" line-data="version: &#39;3&#39;">`version`</SwmToken> specifies the version of the Docker Compose file format. The <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="2:0:0" line-data="name: redis">`name`</SwmToken> sets the project name for the Docker Compose setup.

```yaml
version: '3'
name: redis
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/redis/testdata/docker-compose.yml" line="3">

---

## Redis Service Configuration

The <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="3:0:0" line-data="services:">`services`</SwmToken> section defines the services to be run. Here, the <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="4:1:1" line-data="  redis:">`redis`</SwmToken> service uses the <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="5:4:8" line-data="    image: redis:7-alpine">`redis:7-alpine`</SwmToken> image.

```yaml
services:
  redis:
    image: redis:7-alpine
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/redis/testdata/docker-compose.yml" line="6">

---

## Entrypoint and Command

The <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="6:1:1" line-data="    entrypoint: /runner.sh">`entrypoint`</SwmToken> specifies the script to be run when the container starts. The <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="7:1:1" line-data="    command: --ignore-warnings ARM64-COW-BUG ${REDIS_ARGS}">`command`</SwmToken> provides additional arguments to the entrypoint script.

```yaml
    entrypoint: /runner.sh
    command: --ignore-warnings ARM64-COW-BUG ${REDIS_ARGS}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/redis/testdata/docker-compose.yml" line="8">

---

## Ports

The <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="8:1:1" line-data="    ports:">`ports`</SwmToken> section maps the container's port 6379 to the host's port, which is specified by the <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="9:5:5" line-data="      - ${REDIS_ADDR:-127.0.0.1}:${REDIS_PORT:-6379}:6379">`REDIS_ADDR`</SwmToken> and <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="9:17:17" line-data="      - ${REDIS_ADDR:-127.0.0.1}:${REDIS_PORT:-6379}:6379">`REDIS_PORT`</SwmToken> environment variables.

```yaml
    ports:
      - ${REDIS_ADDR:-127.0.0.1}:${REDIS_PORT:-6379}:6379
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/redis/testdata/docker-compose.yml" line="10">

---

## Environment Variables

The <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="10:1:1" line-data="    environment:">`environment`</SwmToken> section sets environment variables for the container. Here, it allows an empty password for Redis.

```yaml
    environment:
      - "ALLOW_EMPTY_PASSWORD=yes"
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/redis/testdata/docker-compose.yml" line="12">

---

## Volumes

The <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="12:1:1" line-data="    volumes:">`volumes`</SwmToken> section binds host directories to container directories. It binds the <SwmToken path="pkg/network/protocols/redis/testdata/docker-compose.yml" pos="14:6:6" line-data="        source: ${CERTS_PATH}">`CERTS_PATH`</SwmToken> to <SwmPath>[pkg/network/protocols/tls/nodejs/testdata/certs/](pkg/network/protocols/tls/nodejs/testdata/certs/)</SwmPath> and the <SwmPath>[pkg/network/protocols/mysql/testdata/runner.sh](pkg/network/protocols/mysql/testdata/runner.sh)</SwmPath> script to <SwmPath>[pkg/network/protocols/mysql/testdata/runner.sh](pkg/network/protocols/mysql/testdata/runner.sh)</SwmPath> inside the container.

```yaml
    volumes:
      - type: bind
        source: ${CERTS_PATH}
        target: /certs
      - type: bind
        source: ${TESTDIR}/testdata/runner.sh
        target: /runner.sh
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
