---
title: Handling Network Namespace Cache Dump Requests
---
This document will cover the process of handling network namespace cache dump requests. We'll cover:

1. Creating a runtime security client
2. Sending a dump request
3. Processing the request
4. Creating dump and graph files.

Technical document: <SwmLink doc-title="Handling Network Namespace Cache Dump Requests">[Handling Network Namespace Cache Dump Requests](/.swm/handling-network-namespace-cache-dump-requests.8omjib0u.sw.md)</SwmLink>

# [Creating a runtime security client](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/8omjib0u#handling-network-namespace-cache-dump-requests)

The process begins by creating a runtime security client. This client is responsible for managing the security aspects of the network namespace cache dump request. It ensures that the request is authenticated and authorized before proceeding. This step is crucial for maintaining the security and integrity of the system.

# [Sending a dump request](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/8omjib0u#handling-network-namespace-cache-dump-requests)

Once the runtime security client is created, a request is sent to dump the network namespace cache. This request includes parameters that specify the details of the dump, such as whether to include snapshot interfaces. The request is sent to the server, which will process it and generate the necessary files.

# [Processing the request](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/8omjib0u#processing-the-dump-request)

The server receives the dump request and delegates it to the NamespaceResolver. The NamespaceResolver is responsible for handling the actual dumping of the network namespaces. It checks if the platform probe is supported and then proceeds to generate the dump. This step ensures that the request is processed correctly and that the necessary files are created.

# [Creating dump and graph files](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/8omjib0u#creating-dump-and-graph-files)

The NamespaceResolver creates temporary files for the dump and graph. The dump is encoded to JSON format, and a dot graph is generated to represent the network namespaces visually. Any errors that occur during this process are handled appropriately, and the filenames of the dump and graph are printed. This step ensures that the network namespace cache is dumped correctly and that the resulting files are available for further analysis.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
