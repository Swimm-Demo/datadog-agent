---
title: Setting Up the APM Injector
---
This document will cover the process of setting up the APM Injector, which includes:

1. Initializing the APM Injector
2. Configuring the Environment
3. Performing Instrumentation.

Technical document: <SwmLink doc-title="Setting Up the APM Injector">[Setting Up the APM Injector](/.swm/setting-up-the-apm-injector.zocju7pm.sw.md)</SwmLink>

# [Initializing the APM Injector](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/zocju7pm#setupapminjector)

The process begins by initializing the APM Injector. This involves starting a new span for tracing, which helps in monitoring and debugging the setup process. An instance of the installer is then created. This installer will be responsible for configuring the environment and performing the necessary instrumentation.

# [Configuring the Environment](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/zocju7pm#setup)

The next step is configuring the environment. This involves several sub-steps:

1. **Setting Up Default Agent Sockets**: This ensures that the agent can communicate effectively with other components.
2. **Creating Symlinks**: Symlinks are created for sysvinit to ensure that the environment variables are correctly set for the Datadog agent and trace agent. This is crucial for the proper functioning of the agents.
3. **Checking Systemd**: The system checks if systemd is running. If it is, environment overrides are added for various systemd units, and the systemd configuration is reloaded. This ensures that the systemd services are correctly configured to work with the APM Injector.

# [Performing Instrumentation](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/zocju7pm#instrument)

The final step is performing the actual instrumentation. This involves:

1. **Verifying the Shared Library**: Before any instrumentation, the shared library is verified to ensure it is working correctly. This is a critical step to prevent any issues during the instrumentation process.
2. **Host Instrumentation**: The system checks if the host should be instrumented. If so, the necessary steps are performed to instrument the host.
3. **Docker Instrumentation**: The system checks if Docker should be instrumented. If Docker instrumentation is enabled and Docker is installed, the necessary steps are performed to instrument Docker. Additionally, the Docker runtime is verified to ensure it is as expected.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
