---
title: Overview of the Run Function
---
This document will cover the overview of the Run Function, which includes:

1. Gathering container stats
2. Processing container stats
3. Returning the results.

Technical document: <SwmLink doc-title="Overview of the Run Function">[Overview of the Run Function](/.swm/overview-of-the-run-function.5r8m72y5.sw.md)</SwmLink>

# [Gathering Container Stats](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/5r8m72y5#gathering-container-stats)

The Run function begins by gathering container-level statistics from the Cgroups and Docker APIs. This involves retrieving data about the containers running on the host, such as their CPU usage, memory consumption, and network activity. This step ensures that the function has all the necessary data to analyze the performance and health of the containers.

# [Processing Container Stats](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/5r8m72y5#processing-container-stats)

Once the container stats are gathered, the Run function processes these stats into manageable chunks. Each chunk is then converted into a message that includes details like the hostname, CPU count, total memory, and specific container stats. This step is crucial for organizing the data in a way that can be easily analyzed and reported.

# [Returning the Results](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/5r8m72y5#returning-the-results)

After processing the container stats, the Run function returns the results as a series of messages. These messages contain all the relevant container metadata and statistics, which can then be used for monitoring and analysis. This final step ensures that the data collected is made available for further use, providing insights into the performance and health of the containers.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
