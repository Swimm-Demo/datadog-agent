---
title: HTML Rendering Flow
---
This document will cover the HTML Rendering Flow, which includes:

1. Gathering status information
2. Populating status details
3. Retrieving data from the cache
4. Generating the HTML content.

Technical document: <SwmLink doc-title="HTML Rendering Flow">[HTML Rendering Flow](/.swm/html-rendering-flow.84zgafev.sw.md)</SwmLink>

# [Gathering Status Information](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/84zgafev#retrieving-status-information)

The process begins by gathering status information. This involves initializing a map to store various status details. The map will later be populated with specific data points such as hostname statistics, metadata, and host tags. This initial step ensures that we have a structured container to hold all the necessary information for rendering the HTML content.

# [Populating Status Details](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/84zgafev#populating-status-information)

Next, we populate the status map with detailed information. This includes:

- **Hostname Statistics**: Data about the hostname is retrieved and added to the map.
- **Metadata**: Information from the cache is fetched and included.
- **Host Tags**: Tags associated with the host are collected and added.
- **Host Information**: General information about the host is gathered and stored in the map. This step ensures that all relevant data points are collected and organized, providing a comprehensive status report.

# [Retrieving Data from Cache](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/84zgafev#getting-data-from-cache)

In this step, we attempt to retrieve the payload from the cache. If the payload is found in the cache, it is used directly. If not, a new payload is generated and cached for future use. This approach optimizes performance by avoiding redundant data collection and ensuring that the most recent data is always available.

# [Generating HTML Content](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/84zgafev#html-rendering)

Finally, the HTML content is generated using the collected status information. The `HTML` function is responsible for this task. It uses a predefined template (<SwmPath>[comp/metadata/host/hostimpl/status_templates/hostHTML.tmpl](comp/metadata/host/hostimpl/status_templates/hostHTML.tmpl)</SwmPath>) and the status information to render the HTML output. This step ensures that the end user receives a well-formatted and comprehensive HTML report based on the latest status data.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
