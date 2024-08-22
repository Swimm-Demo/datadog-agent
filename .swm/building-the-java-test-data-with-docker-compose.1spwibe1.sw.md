---
title: Building the Java Test Data with Docker Compose
---
# Intro

This document explains how Docker is used in the <SwmPath>[pkg/network/protocols/tls/java/testdata/](pkg/network/protocols/tls/java/testdata/)</SwmPath> directory. It will go through the configuration steps in the <SwmPath>[pkg/network/protocols/tls/java/testdata/docker-compose.yml](pkg/network/protocols/tls/java/testdata/docker-compose.yml)</SwmPath> file step by step.

<SwmSnippet path="/pkg/network/protocols/tls/java/testdata/docker-compose.yml" line="1">

---

# Docker Compose Configuration

The <SwmPath>[pkg/network/protocols/tls/java/testdata/docker-compose.yml](pkg/network/protocols/tls/java/testdata/docker-compose.yml)</SwmPath> file starts by specifying the version of Docker Compose being used. In this case, it is version 3.

```yaml
version: '3'
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/java/testdata/docker-compose.yml" line="2">

---

# Service Name

The <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="2:0:0" line-data="name: java">`name`</SwmToken> field sets the name of the Docker Compose project to <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="2:3:3" line-data="name: java">`java`</SwmToken>.

```yaml
name: java
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/java/testdata/docker-compose.yml" line="3">

---

# Defining Services

The <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="3:0:0" line-data="services:">`services`</SwmToken> section defines the services that will be run. In this case, there is a single service named <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="2:3:3" line-data="name: java">`java`</SwmToken>.

```yaml
services:
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/java/testdata/docker-compose.yml" line="4">

---

# Java Service Configuration

The <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="4:1:1" line-data="  java:">`java`</SwmToken> service uses an image specified by the <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="5:6:6" line-data="    image: ${IMAGE_VERSION}">`IMAGE_VERSION`</SwmToken> environment variable.

```yaml
  java:
    image: ${IMAGE_VERSION}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/java/testdata/docker-compose.yml" line="6">

---

# Entrypoint Configuration

The <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="6:1:1" line-data="    entrypoint: java -cp /v ${ENTRYCLASS}">`entrypoint`</SwmToken> field specifies the command to run when the container starts. Here, it runs a Java class specified by the <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="6:14:14" line-data="    entrypoint: java -cp /v ${ENTRYCLASS}">`ENTRYCLASS`</SwmToken> environment variable.

```yaml
    entrypoint: java -cp /v ${ENTRYCLASS}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/java/testdata/docker-compose.yml" line="7">

---

# Extra Hosts Configuration

The <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="7:1:1" line-data="    extra_hosts:">`extra_hosts`</SwmToken> field allows adding entries to the container's `/etc/hosts` file. The entries are specified by the <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="7:1:1" line-data="    extra_hosts:">`extra_hosts`</SwmToken> environment variable.

```yaml
    extra_hosts:
      - ${EXTRA_HOSTS}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/java/testdata/docker-compose.yml" line="12">

---

# Ulimits Configuration

The <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="12:1:1" line-data="    ulimits:">`ulimits`</SwmToken> section sets resource limits for the container. Here, it sets the <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="13:1:1" line-data="      nofile:">`nofile`</SwmToken> limit, which controls the number of open file descriptors. The soft limit is set to 1024, and the hard limit is set to 524288.

```yaml
    ulimits:
      nofile:
        soft: 1024
        hard: 524288
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/network/protocols/tls/java/testdata/docker-compose.yml" line="16">

---

# Volume Configuration

The <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="16:1:1" line-data="    volumes:">`volumes`</SwmToken> section mounts a directory from the host to the container. The directory is specified by the <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="17:5:5" line-data="      - ${TESTDIR}:/v:z">`TESTDIR`</SwmToken> environment variable and is mounted to <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="6:9:10" line-data="    entrypoint: java -cp /v ${ENTRYCLASS}">`/v`</SwmToken> in the container with the <SwmToken path="pkg/network/protocols/tls/java/testdata/docker-compose.yml" pos="17:9:10" line-data="      - ${TESTDIR}:/v:z">`:z`</SwmToken> option to label the content for SELinux.

```yaml
    volumes:
      - ${TESTDIR}:/v:z
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
