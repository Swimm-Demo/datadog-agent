---
title: Dumping Discarders Process
---
This document will cover the process of dumping discarders, which includes:

1. Initializing the runtime security client
2. Retrieving discarders
3. Processing discarders
4. Encoding discarders.

Technical document: <SwmLink doc-title="Dumping Discarders Process">[Dumping Discarders Process](/.swm/dumping-discarders-process.g8f0kba8.sw.md)</SwmLink>

# [Initializing the Runtime Security Client](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/g8f0kba8#dumping-discarders)

The process begins by initializing a runtime security client. This client is essential for managing the security operations required to retrieve and process discarders. If the client fails to initialize, the process cannot proceed, and an error is reported.

# [Retrieving Discarders](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/g8f0kba8#retrieving-discarders)

Once the runtime security client is initialized, it retrieves discarders from various eBPF maps. These maps include inode discarders, PID discarders, and discarder stats. Retrieving these discarders is crucial as they contain the data that will be processed and encoded in subsequent steps.

# [Processing Discarders](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/g8f0kba8#processing-discarders)

After retrieving the discarders, the next step is to process them. This involves gathering information about PID discarders, inode discarders, and discarder stats. The processed data is then aggregated into a structure that will be used for encoding. This step ensures that all relevant discarder information is collected and organized.

# [Encoding Discarders](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/g8f0kba8#encoding-discarders)

The final step in the process is encoding the processed discarders into a specified format, typically a YAML file. This involves creating a temporary file, setting appropriate permissions, and writing the encoded data to the file. The encoded file is then saved for further analysis and use. This step ensures that the discarders are stored in a manageable and accessible format.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
