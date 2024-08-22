---
title: Data Collection and Transmission
---
This document will cover the process of collecting and sending metrics, logs, and traces using the Datadog Agent. We'll cover:

1. Collecting Data
2. Processing Data
3. Sending Data to Datadog Platform

Technical document: <SwmLink doc-title="" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" path="/.swm/.ois9fuwa.sw.md"></SwmLink>

# [Collecting Data](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ois9fuwa#data-collection)

The Datadog Agent collects metrics, logs, and traces from various sources. This includes system metrics, application logs, and distributed traces. The goal is to gather comprehensive data that provides insights into the performance and health of the infrastructure and applications. The data collection process is continuous and ensures that the most up-to-date information is available for monitoring and analysis.

# [Processing Data](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ois9fuwa#data-processing)

Once the data is collected, it is processed to ensure it is in a format suitable for analysis. This involves filtering, aggregating, and enriching the data. For example, logs might be parsed to extract relevant fields, metrics might be aggregated over time intervals, and traces might be enriched with additional context. The processing step is crucial for transforming raw data into actionable insights.

# [Sending Data to Datadog Platform](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ois9fuwa#data-transmission)

After processing, the data is sent to the Datadog platform. This involves securely transmitting the data over the network to Datadog's servers. The data is then stored and made available for real-time monitoring, alerting, and analysis. Users can access the data through Datadog's dashboards, set up alerts based on specific conditions, and perform detailed analysis to identify and resolve issues.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
