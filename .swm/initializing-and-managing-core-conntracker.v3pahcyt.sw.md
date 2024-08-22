---
title: Initializing and Managing CORE Conntracker
---
This document will cover the process of initializing and managing the CORE Conntracker, which includes:

1. Checking kernel support
2. Loading the CORE asset
3. Continuously polling for updates.

Technical document: <SwmLink doc-title="Initializing and Managing CORE Conntracker">[Initializing and Managing CORE Conntracker](/.swm/initializing-and-managing-core-conntracker.ioym3vv5.sw.md)</SwmLink>

# [Checking Kernel Support](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ioym3vv5#loading-the-core-conntracker)

The process begins by verifying if the kernel supports the CO-RE eBPF conntracker. This step is essential because the CORE conntracker relies on specific kernel features to function correctly. If the kernel does not support these features, the initialization process cannot proceed. This ensures that the system only attempts to load the CORE conntracker on compatible environments, preventing potential errors and inefficiencies.

# [Loading the CORE Asset](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ioym3vv5#loading-core-asset)

Once kernel support is confirmed, the next step is to load the CORE asset. This involves several sub-steps:

1. **Finding Kernel BTF**: The system searches for the kernel's BTF (BPF Type Format) data, which is necessary for the CO-RE (Compile Once - Run Everywhere) program to adapt to different kernel versions.
2. **Reading the CO-RE Object File**: The system reads the CO-RE object file, which contains the eBPF program.
3. **Calling the Callback Function**: A callback function is invoked with the asset and BTF options pre-filled. This function is responsible for loading the CO-RE program correctly. This step ensures that the CO-RE program is loaded with the appropriate configurations and options, making it adaptable to various kernel environments.

# [Starting the Polling Loop](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ioym3vv5#starting-the-poll-loop)

After the CORE asset is loaded, the system initiates a polling loop in a separate thread. This loop continuously checks for updates to the configuration. By running in a separate thread, the polling loop does not interfere with the main operations of the system. This ensures that the system remains responsive while still keeping the configuration up-to-date.

# [Fetching and Applying Updates](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/ioym3vv5#updating-configurations)

The polling loop periodically fetches configuration updates from a remote service. This involves:

1. **Requesting Updates**: The system sends a request to the remote configuration service to fetch the latest configuration updates.
2. **Applying Updates**: Once the updates are received, the system applies them to the current configuration. This step ensures that the system always operates with the most recent configuration, improving performance and security. If the remote service is unavailable, the system handles the error and retries, ensuring robustness and reliability.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
