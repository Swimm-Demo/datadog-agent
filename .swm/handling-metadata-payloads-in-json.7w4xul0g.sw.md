---
title: Handling Metadata Payloads in JSON
---
This document will cover the process of handling metadata payloads in JSON format, which includes:

1. Retrieving the metadata payload
2. Converting the payload to JSON
3. Scrubbing sensitive information from the JSON
4. Writing the scrubbed JSON as a response.

Technical document: <SwmLink doc-title="Handling Metadata Payloads in JSON">[Handling Metadata Payloads in JSON](/.swm/handling-metadata-payloads-in-json.nnlld6rf.sw.md)</SwmLink>

# [Retrieving the Metadata Payload](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/nnlld6rf#getpayload)

The process begins with retrieving the metadata payload. This payload includes various system and configuration data such as OS information, agent flavor, Python version, system stats, and more. The payload is constructed by gathering common payload data, specific payload data, and additional resources if available. If the configuration enables it, the Gohai payload is also included. This step ensures that all relevant data is collected and ready for further processing.

# [Converting the Payload to JSON](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/nnlld6rf#getpayloadasjson)

Once the metadata payload is retrieved, it is converted into a JSON format. This conversion involves marshaling the payload into a JSON byte slice with indentation. The JSON format is chosen because it is a widely used data interchange format that is easy to read and write for both humans and machines.

# [Scrubbing Sensitive Information](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/nnlld6rf#writepayloadasjson)

After converting the payload to JSON, the next step is to scrub the JSON data to remove any sensitive information. This is a critical step to ensure that no confidential or sensitive data is exposed in the JSON response. The scrubbing process involves identifying and removing specific pieces of data that are deemed sensitive, such as personal information or security-related details.

# [Writing the JSON Response](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/nnlld6rf#writepayloadasjson)

Finally, the scrubbed JSON data is written as a response. This step involves sending the cleaned JSON data back to the requester. If any errors occur during the scrubbing or writing process, appropriate error responses are set to inform the requester of the issue. This ensures that the end user receives a clean and accurate JSON response containing the metadata payload.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
