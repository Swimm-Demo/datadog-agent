---
title: Lazy Initialization of SenderManager
---
This document will cover the Lazy Initialization of SenderManager, which includes:

1. Checking for existing SenderManager instance
2. Initializing necessary components
3. Setting up the aggregator
4. Saving and returning the SenderManager instance.

Technical document: <SwmLink doc-title="Lazy Initialization of SenderManager">[Lazy Initialization of SenderManager](/.swm/lazy-initialization-of-sendermanager.o0i7t8ta.sw.md)</SwmLink>

# [Checking for existing SenderManager instance](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/o0i7t8ta#lazygetsendermanager)

The process begins by checking if a SenderManager instance already exists. If it does, the existing instance is returned immediately. This step ensures that we do not create multiple instances of SenderManager, which could lead to redundant operations and inefficient resource usage.

# [Initializing necessary components](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/o0i7t8ta#lazygetsendermanager)

If no existing SenderManager instance is found, the next step is to initialize several critical components. These include the hostname, forwarder, and event platform forwarder. Initializing these components ensures that the SenderManager has all the necessary information and tools to function correctly. For example, the hostname is required for identifying the source of the data, while the forwarder and event platform forwarder are responsible for transmitting data to the appropriate destinations.

# [Setting up the aggregator](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/o0i7t8ta#lazygetsendermanager)

Once the necessary components are initialized, the aggregator is set up with specific options. These options include disabling the flush goroutines and starting the forwarders. Disabling the flush goroutines means that data will not be automatically sent at regular intervals, which can be useful in scenarios where manual control over data transmission is required. Starting the forwarders ensures that they are ready to transmit data as soon as it is available.

# [Saving and returning the SenderManager instance](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/o0i7t8ta#lazygetsendermanager)

After setting up the aggregator, the newly initialized SenderManager instance is saved for future use. This step ensures that subsequent requests for the SenderManager can return the existing instance, avoiding the need to reinitialize the components. Finally, the SenderManager instance is returned, completing the lazy initialization process.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
