---
title: Remote Tagger Initialization and Main Loop
---
This document will cover the Remote Tagger Initialization and Main Loop, which includes:

1. Initializing the connection to the remote tagger
2. Running the main loop for telemetry collection and stream management
3. Stream initialization and maintenance
4. Streaming entity events to clients.

Technical document: <SwmLink doc-title="Remote Tagger Initialization and Main Loop">[Remote Tagger Initialization and Main Loop](/.swm/remote-tagger-initialization-and-main-loop.ik2a04gg.sw.md)</SwmLink>

# [Initializing the Connection to the Remote Tagger](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ik2a04gg#initialization)

The initialization process begins by setting up the connection to the remote tagger. This involves configuring the telemetry ticker, which is responsible for periodically collecting telemetry data. The context and credentials for the gRPC connection are also established at this stage. The credentials ensure secure communication with the remote tagger. Once these components are set up, the tagger stream is started, and the main loop begins running in a separate thread.

# [Running the Main Loop for Telemetry Collection and Stream Management](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ik2a04gg#main-loop)

The main loop is responsible for continuously collecting telemetry data and managing the stream. It operates in a cycle where it checks for telemetry data to collect and processes it. If the stream is unavailable or an error occurs, the loop attempts to re-establish the stream. This ensures that telemetry data collection is resilient and can recover from temporary disruptions.

# [Stream Initialization and Maintenance](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ik2a04gg#stream-initialization)

Stream initialization involves establishing a connection with the remote gRPC endpoint. This process uses exponential backoff for retries, meaning that the time between retry attempts increases exponentially if the connection fails. This approach helps to manage network load and avoid overwhelming the server with frequent retry attempts. Additionally, the stream context is set up with an authorization token to ensure secure communication.

# [Streaming Entity Events to Clients](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ik2a04gg#streaming-entities)

Once the stream is established, the tagger subscribes to entity events and streams them to clients. These events include additions, removals, or changes to entities. The tagger sends these events to clients as they occur, ensuring that clients have up-to-date information. To maintain the connection, keep-alive messages are sent periodically. This helps to prevent the connection from being closed due to inactivity.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
