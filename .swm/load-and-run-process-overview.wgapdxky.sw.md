---
title: Load and Run Process Overview
---
This document will cover the Load and Run process, which includes:

1. Initializing configurations
2. Starting configuration polling
3. Handling streaming configurations
4. Processing removed configurations

Technical document: <SwmLink doc-title="LoadAndRun Process Overview">[LoadAndRun Process Overview](/.swm/loadandrun-process-overview.fkbdn1rv.sw.md)</SwmLink>

# [Initializing Configurations](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/fkbdn1rv#loadandrun-initialization)

The LoadAndRun process begins by loading all available integration configurations and scheduling them. This ensures that providers that do not require polling are queried at least once. The process involves iterating over all configuration pollers, starting them, and logging their status. Additionally, it handles errors from file-based configuration providers to ensure smooth operation.

# [Starting Configuration Polling](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/fkbdn1rv#starting-configuration-polling)

The next step is to start the polling process for the configuration provider. This step differentiates between streaming and collecting configuration providers, initiating the appropriate method for each. For streaming providers, it sets up a continuous stream to listen for changes. For collecting providers, it retrieves configurations once and then periodically polls for updates. This ensures that the system remains up-to-date with the latest configurations.

# [Handling Streaming Configurations](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/fkbdn1rv#streaming-configurations)

For streaming configurations, the process sets up a context with cancellation, registers a health check, and listens for changes from the provider. When changes are detected, it processes removed configurations and applies new ones. This real-time processing ensures that the system's configuration is always current and responsive to changes.

# [Processing Removed Configurations](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/fkbdn1rv#handling-removed-configurations)

The final step involves handling removed configurations. This is initiated by calling the processDelConfigs function to determine the changes needed. The system then applies these changes and deletes any mappings of check <SwmToken path="tasks/gitlab_helpers.py" pos="90:10:10" line-data="def print_gitlab_object(get_object, ctx, ids, repo=&#39;DataDog/datadog-agent&#39;, jq: str | None = None, jq_colors=True):">`ids`</SwmToken> with secrets. This step ensures that the system remains clean and free of outdated configurations, maintaining optimal performance.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
