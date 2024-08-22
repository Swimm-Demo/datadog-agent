---
title: Data Collection and Transmission
---
This document will cover the process of collecting and sending metrics, logs, and traces using the Datadog Agent. We'll cover:

1. Collecting Data
2. Processing Data
3. Sending Data to Datadog Platform

Technical document: <SwmLink doc-title="" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" path="/.swm/.ppcgukys.sw.md"></SwmLink>

```mermaid
graph TD;
    A[Start] --> B[Collect Data];
    B --> C[Process Data];
    C --> D[Send Data to Datadog Platform];
    D --> E[End];
```

# [Collecting Data](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ppcgukys#collecting-data)

The Datadog Agent collects metrics, logs, and traces from various sources. This involves gathering data from different systems, applications, and services that are being monitored. The goal is to ensure comprehensive coverage of all relevant data points that can provide insights into the performance and health of the infrastructure.

# [Processing Data](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ppcgukys#processing-data)

Once the data is collected, it needs to be processed to ensure it is in a format that can be easily analyzed. This step involves filtering, aggregating, and enriching the data. For example, logs might be parsed to extract key information, metrics might be aggregated to provide summary statistics, and traces might be enriched with additional context.

# [Sending Data to Datadog Platform](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ppcgukys#sending-data-to-datadog-platform)

After processing, the data is sent to the Datadog platform. This involves transmitting the data over the network to Datadog's servers, where it can be stored, analyzed, and visualized. The goal is to provide users with real-time insights into their systems, allowing them to monitor performance, detect issues, and troubleshoot problems effectively.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
