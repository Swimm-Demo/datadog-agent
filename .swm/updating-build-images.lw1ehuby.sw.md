---
title: Updating Build Images
---
This document will cover the process of updating build images, which includes:

1. Verifying the workspace
2. Updating configuration files
3. Triggering the build process.

Technical document: <SwmLink doc-title="Updating Build Images Flow">[Updating Build Images Flow](/.swm/updating-build-images-flow.tuaib6m2.sw.md)</SwmLink>

# [Verifying the Workspace](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/tuaib6m2#verify_workspace)

The first step in updating build images is to verify the workspace. This ensures that the workspace is clean and ready for modifications. It involves checking for any potential conflicts both locally and upstream. If no branch name is provided, a new branch name is generated based on the current user. This step is crucial to avoid any conflicts during the update process and ensures that the workspace is in a state where changes can be safely made.

# [Updating Configuration Files](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/tuaib6m2#update_buildimages)

Once the workspace is verified, the next step is to update the configuration files. This involves updating the <SwmToken path="tasks/libs/types/types.py" pos="53:1:1" line-data="    GITLAB = 6">`GITLAB`</SwmToken> and CircleCI configuration files with the new image tag. The new image tag is fetched from the `agent-buildimages`. This step ensures that the build process uses the latest image, which is essential for maintaining up-to-date and secure build environments.

# [Triggering the Build Process](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=/docs/tuaib6m2#trigger_build)

The final step is to trigger the build process. This involves committing the changes made to the configuration files and pushing them to the repository. The user is prompted to confirm if they want to trigger the pipeline. Once confirmed, the changes are committed, and the pipeline is triggered. This step ensures that the new build images are used in the build process, which helps in maintaining consistency and reliability in the build environment.

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
