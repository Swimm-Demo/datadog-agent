---
title: Building the MySQL Test Environment with Docker Compose
---
# Intro

This document explains how Docker is used in the <SwmPath>[pkg/network/protocols/mysql/testdata/](pkg/network/protocols/mysql/testdata/)</SwmPath> directory. It will go through the configuration steps in the <SwmPath>[pkg/network/protocols/mysql/testdata/docker-compose.yml](pkg/network/protocols/mysql/testdata/docker-compose.yml)</SwmPath> file step by step.

<SwmSnippet path="/pkg/network/protocols/mysql/testdata/docker-compose.yml" line="1">

---

# Docker Compose Configuration

The <SwmPath>[pkg/network/protocols/mysql/testdata/docker-compose.yml](pkg/network/protocols/mysql/testdata/docker-compose.yml)</SwmPath> file starts by specifying the version of Docker Compose being used. In this case, it is version 3.

```yaml
version: '3'
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mysql/testdata/docker-compose.yml" line="2">

---

# Service Name

The <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="2:0:0" line-data="name: mysql">`name`</SwmToken> field sets the name of the Docker Compose project to <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="2:3:3" line-data="name: mysql">`mysql`</SwmToken>.

```yaml
name: mysql
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mysql/testdata/docker-compose.yml" line="3">

---

# <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="4:1:1" line-data="  mysql:">`mysql`</SwmToken> Service Configuration

The <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="3:0:0" line-data="services:">`services`</SwmToken> section defines the services that will be run. Here, a service named <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="4:1:1" line-data="  mysql:">`mysql`</SwmToken> is defined, which uses the <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="5:4:10" line-data="    image: mysql:8.0.32">`mysql:8.0.32`</SwmToken> image.

```yaml
services:
  mysql:
    image: mysql:8.0.32
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mysql/testdata/docker-compose.yml" line="6">

---

# Entrypoint and Command

The <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="6:1:1" line-data="    entrypoint: /runner.sh">`entrypoint`</SwmToken> is set to <SwmPath>[pkg/network/protocols/mysql/testdata/runner.sh](pkg/network/protocols/mysql/testdata/runner.sh)</SwmPath>, and the <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="7:1:1" line-data="    command: --default-authentication-plugin=mysql_native_password ${MYSQL_TLS_ARGS}">`command`</SwmToken> specifies additional arguments for the <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="2:3:3" line-data="name: mysql">`mysql`</SwmToken> server, including the <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="7:4:11" line-data="    command: --default-authentication-plugin=mysql_native_password ${MYSQL_TLS_ARGS}">`--default-authentication-plugin=mysql_native_password`</SwmToken> and <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="7:13:16" line-data="    command: --default-authentication-plugin=mysql_native_password ${MYSQL_TLS_ARGS}">`${MYSQL_TLS_ARGS}`</SwmToken>.

```yaml
    entrypoint: /runner.sh
    command: --default-authentication-plugin=mysql_native_password ${MYSQL_TLS_ARGS}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mysql/testdata/docker-compose.yml" line="8">

---

# Restart Policy

The <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="8:1:1" line-data="    restart: always">`restart`</SwmToken> policy is set to <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="8:4:4" line-data="    restart: always">`always`</SwmToken>, ensuring that the <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="2:3:3" line-data="name: mysql">`mysql`</SwmToken> container will always restart if it stops.

```yaml
    restart: always
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mysql/testdata/docker-compose.yml" line="9">

---

# Environment Variables

The <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="9:1:1" line-data="    environment:">`environment`</SwmToken> section sets environment variables for the <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="2:3:3" line-data="name: mysql">`mysql`</SwmToken> container. Here, <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="10:1:1" line-data="      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASS:-root}">`MYSQL_ROOT_PASSWORD`</SwmToken> is set to <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="10:4:9" line-data="      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASS:-root}">`${MYSQL_ROOT_PASS:-root}`</SwmToken>, which defaults to <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="10:8:8" line-data="      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASS:-root}">`root`</SwmToken> if <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="10:6:6" line-data="      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASS:-root}">`MYSQL_ROOT_PASS`</SwmToken> is not provided.

```yaml
    environment:
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASS:-root}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mysql/testdata/docker-compose.yml" line="11">

---

# Port Mapping

The <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="11:1:1" line-data="    ports:">`ports`</SwmToken> section maps port <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="12:19:19" line-data="      - ${MYSQL_ADDR:-127.0.0.1}:${MYSQL_PORT:-3306}:3306">`3306`</SwmToken> of the <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="2:3:3" line-data="name: mysql">`mysql`</SwmToken> container to <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="12:3:20" line-data="      - ${MYSQL_ADDR:-127.0.0.1}:${MYSQL_PORT:-3306}:3306">`${MYSQL_ADDR:-127.0.0.1}:${MYSQL_PORT:-3306}`</SwmToken> on the host machine, defaulting to `127.0.0.1:3306` if not specified.

```yaml
    ports:
      - ${MYSQL_ADDR:-127.0.0.1}:${MYSQL_PORT:-3306}:3306
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mysql/testdata/docker-compose.yml" line="13">

---

# Temporary Filesystem

The <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="13:1:1" line-data="    tmpfs:">`tmpfs`</SwmToken> section mounts a temporary filesystem at <SwmPath>[pkg/network/ebpf/c/protocols/mysql/](pkg/network/ebpf/c/protocols/mysql/)</SwmPath> to store <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="14:8:8" line-data="      - /var/lib/mysql">`mysql`</SwmToken> data.

```yaml
    tmpfs:
      - /var/lib/mysql
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/mysql/testdata/docker-compose.yml" line="15">

---

# Volume Bindings

The <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="15:1:1" line-data="    volumes:">`volumes`</SwmToken> section binds host directories to container directories. The <SwmToken path="pkg/network/protocols/mysql/testdata/docker-compose.yml" pos="17:4:7" line-data="        source: ${CERTS_PATH}">`${CERTS_PATH}`</SwmToken> on the host is bound to <SwmPath>[pkg/network/protocols/tls/nodejs/testdata/certs/](pkg/network/protocols/tls/nodejs/testdata/certs/)</SwmPath> in the container, and <SwmPath>[pkg/network/protocols/mysql/testdata/runner.sh](pkg/network/protocols/mysql/testdata/runner.sh)</SwmPath> is bound to <SwmPath>[pkg/network/protocols/mysql/testdata/runner.sh](pkg/network/protocols/mysql/testdata/runner.sh)</SwmPath> in the container.

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
