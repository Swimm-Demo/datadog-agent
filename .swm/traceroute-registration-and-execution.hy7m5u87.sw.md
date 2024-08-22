---
title: Traceroute Registration and Execution
---
This document will cover the Traceroute Registration and Execution feature, which includes:

1. Setting up the HTTP handler
2. Executing the traceroute
3. Processing the results.

Technical document: <SwmLink doc-title="Traceroute Registration and Execution">[Traceroute Registration and Execution](/.swm/traceroute-registration-and-execution.ss7tfnuf.sw.md)</SwmLink>

# [Setting up the HTTP handler](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ss7tfnuf#register)

The HTTP handler is set up to listen for incoming traceroute requests. When a request is received, it parses the parameters provided by the user. This step is crucial as it ensures that the traceroute operation is initiated correctly with the right parameters. The handler logs the request details and any errors encountered during the parsing process. This setup acts as the entry point for all traceroute operations, ensuring they are correctly processed and logged.

# [Executing the traceroute](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ss7tfnuf#runtraceroute)

Once the parameters are parsed, the traceroute operation is executed. The system resolves the destination hostname to an IP address and determines whether to use TCP or UDP protocols for the traceroute. The appropriate function (`runTCP` or `runUDP`) is then called to perform the actual network path tracing. This step is essential for gathering the network path data based on the provided configuration.

# [Processing the results](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ss7tfnuf#processudpresults)

After the traceroute operation is completed, the results are processed. For UDP traceroutes, the `processUDPResults` function constructs the network path, including details about each hop, such as the IP address and round-trip time (RTT). This processed data is then returned in a structured format, making it easy to interpret and analyze the network path. This step is crucial for converting raw traceroute data into a meaningful format that can be used for monitoring and analysis.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
