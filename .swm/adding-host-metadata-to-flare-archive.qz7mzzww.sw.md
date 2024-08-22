---
title: Adding Host Metadata to Flare Archive
---
This document will cover the process of adding host metadata to a flare archive. We'll cover:

1. Gathering Host Metadata
2. Converting Metadata to JSON
3. Adding Metadata to Flare Archive

Technical document: <SwmLink doc-title="Adding Host Metadata to Flare Archive">[Adding Host Metadata to Flare Archive](/.swm/adding-host-metadata-to-flare-archive.mw7zw1rp.sw.md)</SwmLink>

# [Gathering Host Metadata](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/mw7zw1rp#getpayload)

The process begins with gathering all the necessary host metadata. This includes common payload data, resource payload, and optional gohai payload. The common payload data consists of basic information about the host, such as hostname, system stats, and host tags. The resource payload includes data about the resources available on the host, and the gohai payload provides additional system information if enabled. This comprehensive metadata is collected to ensure that the flare archive contains all relevant information about the host.

# [Converting Metadata to JSON](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/mw7zw1rp#getpayloadasjson)

Once the metadata is gathered, it needs to be converted into a JSON format. This is done to ensure that the data is structured and can be easily read and processed. The conversion involves taking the collected metadata and formatting it into a JSON-formatted byte array. This step is crucial as it standardizes the data, making it easier to include in the flare archive and ensuring compatibility with other systems that may need to read the metadata.

# [Adding Metadata to Flare Archive](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/mw7zw1rp#fillflare)

The final step is to add the JSON-formatted metadata to the flare archive. The flare archive is a collection of diagnostic information that can be used for troubleshooting and analysis. By including the host metadata in the flare archive, we ensure that all relevant information about the host is available for review. This step involves creating a new file within the flare archive and writing the JSON-formatted metadata to this file. This makes the metadata easily accessible and ensures that it is included in any diagnostic reviews or troubleshooting processes.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
