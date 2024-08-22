---
title: Running the Datadog Cluster Agent
---
This document will cover the process of running the Datadog Cluster Agent, which includes:

1. Checking for API key and obtaining hostname
2. Initializing caches and components
3. Setting up the check collector and autoconfig
4. Setting up the server and registering endpoints

Technical document: <SwmLink doc-title="Running the Datadog Cluster Agent">[Running the Datadog Cluster Agent](/.swm/running-the-datadog-cluster-agent.4y330x8q.sw.md)</SwmLink>

# [Checking for API key and obtaining hostname](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/4y330x8q#checking-for-api-key-and-obtaining-hostname)

The first step in running the Datadog Cluster Agent is to ensure that an API key is configured. This is essential because the API key is used to authenticate the agent with the Datadog platform. If the API key is not set, the agent will not be able to send data to Datadog, and it will exit. Additionally, the hostname of the machine running the agent is obtained. This hostname is used to identify the source of the data being sent to Datadog.

# [Initializing caches and components](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/4y330x8q#initializing-caches-and-components)

Once the API key is verified and the hostname is obtained, the next step is to initialize various caches and components required for the agent's operation. These components include the tagger, demultiplexer, workload metadata, autodiscovery, secrets resolver, and collector. Initializing these components ensures that the agent has all the necessary resources to collect and process data efficiently.

# [Setting up the check collector and autoconfig](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/4y330x8q#setting-up-the-check-collector-and-autoconfig)

After initializing the necessary components, the agent sets up the check collector and autoconfig. The check collector is responsible for running various checks to collect metrics, logs, and traces from different sources. Autoconfig automatically discovers and configures integrations based on the environment in which the agent is running. This step ensures that the agent can dynamically adapt to changes in the environment and continue to collect data without manual intervention.

# [Setting up the server and registering endpoints](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/4y330x8q#setting-up-the-server-and-registering-endpoints)

The final step in running the Datadog Cluster Agent is to set up the server and register the necessary endpoints. The server is responsible for handling incoming requests and serving the API endpoints. These endpoints include routes for cluster checks, metadata, language detection, and more. Setting up the server involves creating the router, initializing authentication tokens, configuring TLS settings, and creating a gRPC server that is multiplexed with the HTTP server. Once the server is set up, it starts listening for incoming requests, ensuring that the agent is ready to handle data collection and processing tasks.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
