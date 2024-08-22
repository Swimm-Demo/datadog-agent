---
title: Handling Discarders
---
This document will cover the process of handling discarders in the Datadog Agent, which includes:

1. Initializing the runtime security client
2. Retrieving discarders
3. Encoding discarders
4. Writing discarders to a file.

Technical document: <SwmLink doc-title="Overview of the dumpDiscarders Process">[Overview of the dumpDiscarders Process](/.swm/overview-of-the-dumpdiscarders-process.4gbx046o.sw.md)</SwmLink>

# [Initializing the Runtime Security Client](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/4gbx046o#handling-discarders)

The process begins by initializing a runtime security client. This client is essential for interacting with the security monitoring system. It ensures that the system is ready to handle discarders, which are data structures used to filter out unnecessary information. This step is crucial for setting up the environment needed to manage discarders effectively.

# [Retrieving Discarders](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/4gbx046o#getdiscarders)

Once the runtime security client is initialized, the next step is to retrieve discarders from various eBPF maps. These maps contain data that helps in identifying and filtering out irrelevant security events. The discarders are collected from different sources, such as inode and PID discarders, which are specific to file and process activities respectively. This step ensures that only the necessary data is processed further.

# [Encoding Discarders](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/4gbx046o#encode-function)

After retrieving the discarders, they need to be encoded into a specific format. The encoding process involves converting the discarders into formats like JSON, Protobuf, or DOT. This step is important because it standardizes the data, making it easier to store and analyze. The choice of format depends on the requirements of the monitoring system and the type of analysis to be performed.

# [Writing Discarders to a File](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/4gbx046o#dumpdiscarders)

The final step in the process is writing the encoded discarders to a temporary file. This file serves as a storage medium for the discarders, allowing them to be accessed and analyzed later. The filename of the dumped discarders is returned, which can be used for reference in future operations. This step ensures that the discarders are safely stored and can be retrieved when needed.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
