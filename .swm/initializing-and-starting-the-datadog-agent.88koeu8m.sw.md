---
title: Initializing and Starting the Datadog Agent
---
This document will cover the process of initializing and starting the Datadog Agent, which includes:

1. Starting runtime security
2. Running the agent
3. Stopping the agent.

Technical document: <SwmLink doc-title="Run Function Overview">[Run Function Overview](/.swm/run-function-overview.hk7zu2zs.sw.md)</SwmLink>

# [Starting Runtime Security](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/hk7zu2zs#startruntimesecurity)

The process begins by checking if the runtime security feature is enabled in the configuration. If it is not enabled, a message is logged, and the process continues without starting runtime security. If enabled, the system attempts to start the runtime security agent. However, it is important to note that this feature is currently only supported on Linux and Windows platforms.

# [Running the Agent](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/hk7zu2zs#runagent)

Once runtime security is handled, the agent is set up to run. This involves initializing various components such as logging, configuration, and telemetry. The agent then starts and enters a waiting state, ready to perform its monitoring tasks. During this phase, the agent collects metrics, logs, and traces from various sources and sends them to the Datadog platform for analysis.

# [Stopping the Agent](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/hk7zu2zs#stopagent)

When a stop signal is received, the agent begins a graceful shutdown process. This involves stopping various components like the metaScheduler and statsd, and shutting down the expvar server if it is running. The agent also retrieves its health status to ensure all components are stopped correctly. This ensures that all resources are cleaned up properly, preventing potential issues during the next startup.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
