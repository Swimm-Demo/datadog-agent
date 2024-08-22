---
title: Docker
---
# Intro

This document explains how Docker is used in the <SwmPath>[pkg/network/java/testutil/JavaClientSimulator/](pkg/network/java/testutil/JavaClientSimulator/)</SwmPath> directory. It will cover the <SwmPath>[pkg/network/java/testutil/JavaClientSimulator/Dockerfile](pkg/network/java/testutil/JavaClientSimulator/Dockerfile)</SwmPath> configuration and the steps involved in building and running the Docker container.

<SwmSnippet path="/pkg/network/java/testutil/JavaClientSimulator/Dockerfile" line="1">

---

### Base Image and Maintainer

The <SwmPath>[pkg/network/java/testutil/JavaClientSimulator/Dockerfile](pkg/network/java/testutil/JavaClientSimulator/Dockerfile)</SwmPath> starts by specifying the base image and the maintainer. The base image <SwmToken path="pkg/network/java/testutil/JavaClientSimulator/Dockerfile" pos="2:2:4" line-data="FROM openjdk:11">`openjdk:11`</SwmToken> includes Java 11, which is required to run the <SwmToken path="pkg/network/java/testutil/JavaClientSimulator/Dockerfile" pos="8:4:4" line-data="COPY target/JavaClientSimulator-1.0.jar app.jar">`JavaClientSimulator`</SwmToken>. The <SwmToken path="pkg/network/java/testutil/JavaClientSimulator/Dockerfile" pos="3:0:0" line-data="MAINTAINER val">`MAINTAINER`</SwmToken> instruction specifies the maintainer of the image.

````FROM openjdk:11
MAINTAINER val```
````

```
# Base image with Java installed
FROM openjdk:11
MAINTAINER val
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
