---
title: Data Collection and Transmission
---
This document will cover the process of collecting and sending metrics, logs, and traces using the Datadog Agent. We'll cover:

1. Collecting Data
2. Processing Data
3. Sending Data to Datadog Platform

Technical document: <SwmLink doc-title="" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" path="/.swm/.2dnbdkv4.sw.md"></SwmLink>

```mermaid
graph TD;
    A[Start] --> B[Collect Data];
    B --> C[Process Data];
    C --> D[Send Data to Datadog Platform];
    D --> E[End];
```

# [Collecting Data](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/2dnbdkv4#collecting-data)

The Datadog Agent collects metrics, logs, and traces from various sources. This involves gathering data from different systems, applications, and services that are being monitored. The goal is to ensure comprehensive coverage of all relevant data points that can provide insights into the performance and health of the infrastructure.

# [Processing Data](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/2dnbdkv4#processing-data)

Once the data is collected, it needs to be processed to ensure it is in a format that can be easily analyzed. This step involves filtering, aggregating, and transforming the raw data into meaningful metrics and logs. The processed data should be structured in a way that makes it easy to visualize and interpret on the Datadog platform.

# [Sending Data to Datadog Platform](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/2dnbdkv4#sending-data-to-datadog-platform)

After processing, the data is sent to the Datadog platform. This step involves securely transmitting the data over the network to Datadog's servers. The data is then stored and made available for real-time monitoring and analysis. Users can access dashboards, set up alerts, and generate reports based on the data sent by the agent.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
