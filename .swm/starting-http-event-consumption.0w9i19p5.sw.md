---
title: Starting HTTP Event Consumption
---
This document will cover the process of starting HTTP event consumption, which includes:

1. Initializing the driver monitor
2. Reading HTTP flows
3. Handling various HTTP events.

Technical document: <SwmLink doc-title="Starting HTTP Event Consumption">[Starting HTTP Event Consumption](/.swm/starting-http-event-consumption.dugdd95n.sw.md)</SwmLink>

# [Initializing the Driver Monitor](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/dugdd95n#initialization)

The process begins by initializing the driver monitor. This involves logging the start of the driver monitor and setting up the necessary components to begin reading HTTP flows. This step ensures that the system is ready to handle incoming HTTP events by preparing the necessary infrastructure.

# [Reading HTTP Flows](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/dugdd95n#reading-http-flows)

The next step is to start reading HTTP flows. This involves initializing the ETW HTTP service subscription and starting separate processes for ETW tracing and reading HTTP transactions. The goal here is to continuously monitor and capture HTTP transactions, which are then sent to a data channel for further processing. This step ensures that all HTTP transactions are captured in real-time for analysis.

# [Handling HTTP Events](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/dugdd95n#event-handling)

The final step is handling various HTTP events. This involves processing different types of HTTP events based on their event descriptor <SwmToken path="tasks/gitlab_helpers.py" pos="90:10:10" line-data="def print_gitlab_object(get_object, ctx, ids, repo=&#39;DataDog/datadog-agent&#39;, jq: str | None = None, jq_colors=True):">`ids`</SwmToken>. The system updates the total bytes transferred and processes events such as connection traces, request traces, and cache traces. This step ensures that all relevant HTTP events are processed and logged, providing a comprehensive view of HTTP activity for monitoring and analysis.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
