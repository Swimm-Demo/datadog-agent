---
title: Starting Runtime Security
---
This document will cover the process of starting the Runtime Security feature, which includes:

1. Checking Configuration
2. Creating Necessary Components
3. Setting Up Logging and Endpoints
4. Starting the Agent

Technical document: <SwmLink doc-title="Starting Runtime Security">[Starting Runtime Security](/.swm/starting-runtime-security.kz6y2nr8.sw.md)</SwmLink>

# [Checking Configuration](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/kz6y2nr8#startruntimesecurity)

The first step in starting the Runtime Security feature is to check if it is enabled in the configuration. This involves verifying a specific configuration setting that determines whether the runtime security agent should be activated. If this setting is not enabled, the process stops here, and the runtime security agent is not started. This step ensures that the feature is only activated when explicitly configured, preventing unnecessary resource usage.

# [Creating Necessary Components](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/kz6y2nr8#startruntimesecurity)

If the runtime security is enabled, the next step is to create a new instance of the runtime security agent. This involves initializing various components required for the agent to function correctly. These components include the logging context, endpoints, and the CWS reporter. Creating these components is essential for the agent to collect and report security-related data effectively.

# [Setting Up Logging and Endpoints](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/kz6y2nr8#newlogcontextruntime)

Setting up logging and endpoints is a critical step in the process. This involves configuring the logging context to ensure that all security events are properly logged and can be reviewed later. Additionally, endpoints are set up to facilitate communication between the runtime security agent and other components or services. Proper logging and endpoint configuration are crucial for monitoring and analyzing security events.

# [Starting the Agent](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/kz6y2nr8#start)

The final step is to start the runtime security agent. This involves initializing the HTTP server, setting up the necessary handlers for various endpoints, and ensuring that the agent is ready to process and report security events. Starting the agent is the culmination of the previous steps and marks the point where the runtime security feature becomes active and operational.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
