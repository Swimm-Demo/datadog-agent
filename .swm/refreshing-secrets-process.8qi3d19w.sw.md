---
title: Refreshing Secrets Process
---
This document will cover the process of refreshing secrets, which includes:

1. Initiating the refresh process
2. Fetching secrets from the backend
3. Updating the audit file
4. Executing necessary commands.

Technical document: <SwmLink doc-title="Refreshing Secrets Flow">[Refreshing Secrets Flow](/.swm/refreshing-secrets-flow.jj72nyi8.sw.md)</SwmLink>

# [Initiating the Refresh Process](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/jj72nyi8#handlerefresh)

The refresh process begins when the system initiates a request to refresh the secrets. This step is crucial as it triggers the entire flow. If an error occurs during this initiation, it is handled appropriately to ensure the system's stability. The result of this initiation is then prepared for further processing.

# [Fetching Secrets from Backend](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/jj72nyi8#refresh)

Once the refresh process is initiated, the system fetches the latest secrets from the backend. This involves identifying which secrets need to be refreshed based on a predefined allowlist. The system then retrieves these secrets, ensuring that only the necessary and allowed secrets are fetched. This step is essential to maintain the security and integrity of the secrets.

# [Updating the Audit File](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/jj72nyi8#adding-to-audit-file)

After fetching the secrets, the system updates the audit file. This file keeps a record of all the secrets that have been refreshed. The system sorts the handles of the secrets, scrubs any sensitive values, and appends the new records to the audit file. This step ensures that there is a clear and traceable record of all secret refresh activities, which is vital for auditing and compliance purposes.

# [Executing Necessary Commands](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/jj72nyi8#execcommand)

The final step in the refresh flow is executing the necessary commands to fetch the secrets. The system sets up the command context, handles input and output streams, and logs the execution details. If the command fails, the system logs the error and manages it accordingly. This step ensures that the secrets are fetched securely and efficiently.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
