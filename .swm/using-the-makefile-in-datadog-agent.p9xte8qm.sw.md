---
title: Using the Makefile in Datadog Agent
---
# Intro

This document explains how to use the <SwmPath>[test/e2e/containers/dsd_sender/Makefile](test/e2e/containers/dsd_sender/Makefile)</SwmPath> in the Datadog Agent repository. It will cover the various targets and their purposes, providing snippets and explanations for each step.

<SwmSnippet path="/Makefile.trace" line="4">

---

# Default <SwmToken path="Makefile.trace" pos="4:6:6" line-data="# if the TRACE_AGENT_VERSION environment variable isn&#39;t set, default to 0.99.0">`TRACE_AGENT_VERSION`</SwmToken>

The <SwmToken path="Makefile.trace" pos="4:6:6" line-data="# if the TRACE_AGENT_VERSION environment variable isn&#39;t set, default to 0.99.0">`TRACE_AGENT_VERSION`</SwmToken> environment variable is set to <SwmToken path="Makefile.trace" pos="4:23:27" line-data="# if the TRACE_AGENT_VERSION environment variable isn&#39;t set, default to 0.99.0">`0.99.0`</SwmToken> if it is not already defined.

```trace
# if the TRACE_AGENT_VERSION environment variable isn't set, default to 0.99.0
TRACE_AGENT_VERSION := $(if $(TRACE_AGENT_VERSION),$(TRACE_AGENT_VERSION), 0.99.0)
```

---

</SwmSnippet>

<SwmSnippet path="/Makefile.trace" line="7">

---

# Version Breakdown

The version string is split into major, minor, and patch components using shell commands and <SwmToken path="Makefile.trace" pos="9:25:25" line-data="VERSION_MAJOR = $(shell echo $(word 1, $(SPLAT)) | sed &#39;s/[^0-9]*//g&#39;)">`sed`</SwmToken>.

```trace
# break up the version
SPLAT = $(subst ., ,$(TRACE_AGENT_VERSION))
VERSION_MAJOR = $(shell echo $(word 1, $(SPLAT)) | sed 's/[^0-9]*//g')
VERSION_MINOR = $(shell echo $(word 2, $(SPLAT)) | sed 's/[^0-9]*//g')
VERSION_PATCH = $(shell echo $(word 3, $(SPLAT)) | sed 's/[^0-9]*//g')
```

---

</SwmSnippet>

<SwmSnippet path="/Makefile.trace" line="13">

---

# Default Version Values

Default values are assigned to the major, minor, and patch version components if they are not already set.

```trace
# account for some defaults
VERSION_MAJOR := $(if $(VERSION_MAJOR),$(VERSION_MAJOR), 0)
VERSION_MINOR := $(if $(VERSION_MINOR),$(VERSION_MINOR), 0)
VERSION_PATCH := $(if $(VERSION_PATCH),$(VERSION_PATCH), 0)
```

---

</SwmSnippet>

<SwmSnippet path="/Makefile.trace" line="18">

---

# Install Target

The <SwmToken path="Makefile.trace" pos="18:0:0" line-data="install:">`install`</SwmToken> target generates versioning information and installs the trace agent binary using <SwmToken path="Makefile.trace" pos="20:1:3" line-data="	go generate ./pkg/trace/info">`go generate`</SwmToken> and <SwmToken path="Makefile.trace" pos="21:1:3" line-data="	go install ./cmd/trace-agent">`go install`</SwmToken>.

```trace
install:
	# generate versioning information and installing the binary.
	go generate ./pkg/trace/info
	go install ./cmd/trace-agent
```

---

</SwmSnippet>

<SwmSnippet path="/Makefile.trace" line="23">

---

# CI Target

The <SwmToken path="Makefile.trace" pos="23:0:0" line-data="ci:">`ci`</SwmToken> target is used by the CI pipeline to install the trace agent and run tests with the race detector enabled.

```trace
ci:
	# task used by CI
	go install ./cmd/trace-agent
	go test -v -race ./...
```

---

</SwmSnippet>

<SwmSnippet path="/Makefile.trace" line="28">

---

# Windows Target

The <SwmToken path="Makefile.trace" pos="28:0:0" line-data="windows:">`windows`</SwmToken> target prepares resources needed for the Windows release, including compiling message strings and resource files.

```trace
windows:
	# pre-packages resources needed for the windows release
	windmc --target pe-x86-64 -r pkg/util/winutil/messagestrings -h pkg/util/winutil/messagestrings pkg/util/winutil/messagestrings/messagestrings.mc
	windres -i pkg/util/winutil/messagestrings/messagestrings.rc --target=pe-x86-64 -O coff -o pkg/util/winutil/messagestrings/rsrc.syso
	windres --define MAJ_VER=$(VERSION_MAJOR) --define MIN_VER=$(VERSION_MINOR) --define PATCH_VER=$(VERSION_PATCH) -i cmd/trace-agent/windows/resources/trace-agent.rc --target=pe-x86-64 -O coff -o cmd/trace-agent/rsrc.syso
```

---

</SwmSnippet>

<SwmSnippet path="/test/e2e/containers/fake_datadog/Makefile" line="7">

---

# Virtual Environment Setup

The <SwmToken path="test/e2e/containers/fake_datadog/Makefile" pos="7:0:0" line-data="venv:">`venv`</SwmToken> target creates a Python virtual environment using <SwmToken path="test/e2e/containers/fake_datadog/Makefile" pos="8:1:1" line-data="	virtualenv venv -p python3">`virtualenv`</SwmToken>.

```
venv:
	virtualenv venv -p python3
```

---

</SwmSnippet>

<SwmSnippet path="/test/e2e/containers/fake_datadog/Makefile" line="10">

---

# Install Python Dependencies

The <SwmToken path="test/e2e/containers/fake_datadog/Makefile" pos="10:0:0" line-data="pip: venv">`pip`</SwmToken> target installs Python dependencies listed in <SwmPath>[test/e2e/containers/fake_datadog/app/requirements.txt](test/e2e/containers/fake_datadog/app/requirements.txt)</SwmPath> using the virtual environment.

```
pip: venv
	venv/bin/pip install -r app/requirements.txt
```

---

</SwmSnippet>

<SwmSnippet path="/test/e2e/containers/fake_datadog/Makefile" line="13">

---

# Docker Build

The <SwmToken path="test/e2e/containers/fake_datadog/Makefile" pos="13:0:0" line-data="build:">`build`</SwmToken> target builds a Docker image for the fake Datadog service, tagging it with the current date.

```
build:
	docker build --force-rm -t datadog/fake-datadog:$(TAG) .
```

---

</SwmSnippet>

<SwmSnippet path="/test/e2e/containers/fake_datadog/Makefile" line="16">

---

# Multi-Architecture Docker Build

The <SwmToken path="test/e2e/containers/fake_datadog/Makefile" pos="16:0:0" line-data="multiarch:">`multiarch`</SwmToken> target builds and pushes a multi-architecture Docker image for the fake Datadog service.

```
multiarch:
	docker buildx build --platform linux/amd64,linux/arm64 -t datadog/fake-datadog:$(TAG) . --push
```

---

</SwmSnippet>

<SwmSnippet path="/test/e2e/containers/fake_datadog/Makefile" line="19">

---

# Push Docker Image

The <SwmToken path="test/e2e/containers/fake_datadog/Makefile" pos="19:0:0" line-data="push:">`push`</SwmToken> target pushes the Docker image for the fake Datadog service to the Docker registry.

```
push:
	docker push datadog/fake-datadog:$(TAG)
```

---

</SwmSnippet>

<SwmSnippet path="/test/e2e/containers/otlp_sender/Makefile" line="5">

---

# Build OTLP Sender

The <SwmToken path="test/e2e/containers/otlp_sender/Makefile" pos="5:0:0" line-data="otlpsender:">`otlpsender`</SwmToken> target builds the OTLP sender binary for Linux using <SwmToken path="test/e2e/containers/otlp_sender/Makefile" pos="6:9:11" line-data="	GOOS=linux GOARCH=amd64 go build -o $@ ./cmd/sender">`go build`</SwmToken>.

```
otlpsender:
	GOOS=linux GOARCH=amd64 go build -o $@ ./cmd/sender
```

---

</SwmSnippet>

<SwmSnippet path="/test/e2e/containers/otlp_sender/Makefile" line="8">

---

# Docker Build for OTLP Sender

The <SwmToken path="test/e2e/containers/otlp_sender/Makefile" pos="8:0:2" line-data="docker-build: otlpsender">`docker-build`</SwmToken> target builds a Docker image for the OTLP sender service.

```
docker-build: otlpsender
	docker build --force-rm -t datadog/docker-library:e2e-otlp-sender_$(TAG) .
```

---

</SwmSnippet>

<SwmSnippet path="/test/e2e/containers/otlp_sender/Makefile" line="11">

---

# Push OTLP Sender Docker Image

The <SwmToken path="test/e2e/containers/otlp_sender/Makefile" pos="11:0:0" line-data="push: docker-build">`push`</SwmToken> target pushes the Docker image for the OTLP sender service to the Docker registry.

```
push: docker-build
	docker push datadog/docker-library:e2e-otlp-sender_$(TAG)
```

---

</SwmSnippet>

<SwmSnippet path="/test/e2e/containers/dsd_sender/Makefile" line="5">

---

# Docker Build for DSD Sender

The <SwmToken path="test/e2e/containers/dsd_sender/Makefile" pos="5:0:0" line-data="build:">`build`</SwmToken> target builds a Docker image for the DSD sender service.

```
build:
	docker build --force-rm -t datadog/docker-library:e2e-dsd-sender_$(TAG) .
```

---

</SwmSnippet>

<SwmSnippet path="/test/e2e/containers/dsd_sender/Makefile" line="8">

---

# Push DSD Sender Docker Image

The <SwmToken path="test/e2e/containers/dsd_sender/Makefile" pos="8:0:0" line-data="push:">`push`</SwmToken> target pushes the Docker image for the DSD sender service to the Docker registry.

```
push:
	docker push datadog/docker-library:e2e-dsd-sender_$(TAG)
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
