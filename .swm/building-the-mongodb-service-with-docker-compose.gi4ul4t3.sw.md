---
title: Building the MongoDB Service with Docker Compose
---
# Intro

This document explains how Docker is used in the <SwmPath>[pkg/network/protocols/mongo/testdata/](pkg/network/protocols/mongo/testdata/)</SwmPath> directory. It will go through the configuration steps in the <SwmPath>[pkg/network/protocols/mongo/testdata/docker-compose.yml](pkg/network/protocols/mongo/testdata/docker-compose.yml)</SwmPath> file.

<SwmSnippet path="/pkg/network/protocols/mongo/testdata/docker-compose.yml" line="1">

---

## Docker-compose configuration

The <SwmPath>[pkg/network/protocols/mongo/testdata/docker-compose.yml](pkg/network/protocols/mongo/testdata/docker-compose.yml)</SwmPath> file starts by specifying the version of the Docker Compose file format being used. In this case, it is version 3.

```yaml
version: '3'
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mongo/testdata/docker-compose.yml" line="2">

---

## Service Name

The <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="2:0:0" line-data="name: mongo">`name`</SwmToken> field sets the name of the Docker Compose project. Here, it is set to <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="2:3:3" line-data="name: mongo">`mongo`</SwmToken>.

```yaml
name: mongo
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mongo/testdata/docker-compose.yml" line="3">

---

## Services Configuration

The <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="3:0:0" line-data="services:">`services`</SwmToken> section defines the services that will be run by Docker Compose. In this file, there is one service defined: <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="4:1:3" line-data="  mongodb-primary:">`mongodb-primary`</SwmToken>.

```yaml
services:
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mongo/testdata/docker-compose.yml" line="4">

---

## <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="4:1:1" line-data="  mongodb-primary:">`mongodb`</SwmToken> Primary Service

The <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="4:1:3" line-data="  mongodb-primary:">`mongodb-primary`</SwmToken> service uses the <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="5:5:11" line-data="    image: &#39;mongo:5.0.14&#39;">`mongo:5.0.14`</SwmToken> image. This specifies the Docker image to be used for this service.

```yaml
  mongodb-primary:
    image: 'mongo:5.0.14'
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mongo/testdata/docker-compose.yml" line="6">

---

## Port Mapping

The <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="6:1:1" line-data="    ports:">`ports`</SwmToken> section maps the host machine's port to the container's port. Here, it maps <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="7:3:20" line-data="      - ${MONGO_ADDR:-127.0.0.1}:${MONGO_PORT:-27017}:27017">`${MONGO_ADDR:-127.0.0.1}:${MONGO_PORT:-27017}`</SwmToken> to <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="7:19:19" line-data="      - ${MONGO_ADDR:-127.0.0.1}:${MONGO_PORT:-27017}:27017">`27017`</SwmToken> inside the container. This allows access to <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="4:1:1" line-data="  mongodb-primary:">`mongodb`</SwmToken> on the specified host and port.

```yaml
    ports:
      - ${MONGO_ADDR:-127.0.0.1}:${MONGO_PORT:-27017}:27017
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mongo/testdata/docker-compose.yml" line="8">

---

## Environment Variables

The <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="8:1:1" line-data="    environment:">`environment`</SwmToken> section sets environment variables for the container. It sets <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="9:3:3" line-data="      - MONGO_INITDB_ROOT_USERNAME=${MONGO_USER:-root}">`MONGO_INITDB_ROOT_USERNAME`</SwmToken> to `${MONGO_USER:-root}` and <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="10:3:3" line-data="      - MONGO_INITDB_ROOT_PASSWORD=${MONGO_PASSWORD:-password}">`MONGO_INITDB_ROOT_PASSWORD`</SwmToken> to `${MONGO_PASSWORD:-password}`. These variables initialize the root username and password for <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="4:1:1" line-data="  mongodb-primary:">`mongodb`</SwmToken>.

```yaml
    environment:
      - MONGO_INITDB_ROOT_USERNAME=${MONGO_USER:-root}
      - MONGO_INITDB_ROOT_PASSWORD=${MONGO_PASSWORD:-password}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mongo/testdata/docker-compose.yml" line="11">

---

## Temporary Filesystem

The <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="11:1:1" line-data="    tmpfs:">`tmpfs`</SwmToken> section mounts a temporary filesystem inside the container at <SwmPath>[pkg/fleet/internal/db/](pkg/fleet/internal/db/)</SwmPath>. This is used for storing <SwmToken path="pkg/network/protocols/mongo/testdata/docker-compose.yml" pos="4:1:1" line-data="  mongodb-primary:">`mongodb`</SwmToken> data in a non-persistent manner.

```yaml
    tmpfs:
      - /data/db
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
