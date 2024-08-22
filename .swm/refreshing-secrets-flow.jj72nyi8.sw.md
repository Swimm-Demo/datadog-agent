---
title: Refreshing Secrets Flow
---
In this document, we will explain the process of refreshing secrets. The process involves initiating the refresh, fetching secrets from the backend, updating the audit file, and executing necessary commands.

The flow starts with initiating the refresh process. This involves calling a function to refresh the secrets. If an error occurs, it is handled appropriately. Next, the secrets are fetched from the backend using a custom command. The fetched secrets are then processed and added to an audit file. Finally, the command execution details are logged, and any errors are managed.

# Flow drill down

```mermaid
graph TD;
      subgraph compcoresecretssecretsimpl["comp/core/secrets/secretsimpl"]
69dc07895e76119b2ed91838bd61c5fb2aa0b5da49ee38649662983b12611c46(handleRefresh):::mainFlowStyle --> 4f7d3af4547628df96555be719f46f6d13f9a03d817ea10c8afe5948ce298c0b(Refresh):::mainFlowStyle
end

subgraph compcoresecretssecretsimpl["comp/core/secrets/secretsimpl"]
4f7d3af4547628df96555be719f46f6d13f9a03d817ea10c8afe5948ce298c0b(Refresh):::mainFlowStyle --> c0663538e538df25cbf26ef7b5f544f7dac79acba41e36e3721865d05b9b8c67(addToAuditFile)
end

subgraph compcoresecretssecretsimpl["comp/core/secrets/secretsimpl"]
4f7d3af4547628df96555be719f46f6d13f9a03d817ea10c8afe5948ce298c0b(Refresh):::mainFlowStyle --> 685eeab3c789537c11d8ce35c0bcf328c0ad306de4e5314874df0df6467c8881(fetchSecret):::mainFlowStyle
end

subgraph compcoresecretssecretsimpl["comp/core/secrets/secretsimpl"]
685eeab3c789537c11d8ce35c0bcf328c0ad306de4e5314874df0df6467c8881(fetchSecret):::mainFlowStyle --> 68ebde6242e57749bf18e7d0608891101bb0dfbd792aee307ade9a2581ff43e3(execCommand):::mainFlowStyle
end

subgraph compcoresecretssecretsimpl["comp/core/secrets/secretsimpl"]
68ebde6242e57749bf18e7d0608891101bb0dfbd792aee307ade9a2581ff43e3(execCommand):::mainFlowStyle --> 701ffc4def2213e7fe28a30be6e9585b99d66d16c2ea36c23761350a36127d85(Add):::mainFlowStyle
end


      classDef mainFlowStyle color:#000000,fill:#7CB9F4
classDef rootsStyle color:#000000,fill:#00FFF4
classDef Style1 color:#000000,fill:#00FFAA
classDef Style2 color:#000000,fill:#FFFF00
classDef Style3 color:#000000,fill:#AA7CB9
```

<SwmSnippet path="/comp/core/secrets/secretsimpl/secrets.go" line="150">

---

## <SwmToken path="comp/core/secrets/secretsimpl/secrets.go" pos="150:9:9" line-data="func (r *secretResolver) handleRefresh(w http.ResponseWriter, _ *http.Request) {">`handleRefresh`</SwmToken>

The <SwmToken path="comp/core/secrets/secretsimpl/secrets.go" pos="150:9:9" line-data="func (r *secretResolver) handleRefresh(w http.ResponseWriter, _ *http.Request) {">`handleRefresh`</SwmToken> function initiates the refresh process by calling the <SwmToken path="comp/core/secrets/secretsimpl/secrets.go" pos="151:10:10" line-data="	result, err := r.Refresh()">`Refresh`</SwmToken> method. If an error occurs during the refresh, it sets a JSON error response; otherwise, it writes the result to the response writer.

```go
func (r *secretResolver) handleRefresh(w http.ResponseWriter, _ *http.Request) {
	result, err := r.Refresh()
	if err != nil {
		setJSONError(w, err, 500)
		return
	}
	w.Write([]byte(result))
}
```

---

</SwmSnippet>

<SwmSnippet path="/comp/core/secrets/secretsimpl/secrets.go" line="439">

---

## Refresh

The <SwmToken path="comp/core/secrets/secretsimpl/secrets.go" pos="439:2:2" line-data="// Refresh the secrets after they have been Resolved by fetching them from the backend again">`Refresh`</SwmToken> function is responsible for re-fetching secrets from the backend. It locks the resolver, filters handles based on an allowlist, and fetches the secrets using either a test hook or the <SwmToken path="comp/core/secrets/secretsimpl/fetch_secret.go" pos="105:2:2" line-data="// fetchSecret receives a list of secrets name to fetch, exec a custom">`fetchSecret`</SwmToken> function. It then processes the secret response and updates the audit file if necessary.

```go
// Refresh the secrets after they have been Resolved by fetching them from the backend again
func (r *secretResolver) Refresh() (string, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	// get handles from the cache that match the allowlist
	newHandles := maps.Keys(r.cache)
	if allowlistPaths != nil {
		filteredHandles := make([]string, 0, len(newHandles))
		for _, handle := range newHandles {
			if r.matchesAllowlist(handle) {
				filteredHandles = append(filteredHandles, handle)
			}
		}
		newHandles = filteredHandles
	}
	if len(newHandles) == 0 {
		return "", nil
	}

	log.Infof("Refreshing secrets for %d handles", len(newHandles))
```

---

</SwmSnippet>

<SwmSnippet path="/comp/core/secrets/secretsimpl/secrets.go" line="502">

---

### Adding to Audit File

The <SwmToken path="comp/core/secrets/secretsimpl/secrets.go" pos="502:2:2" line-data="// addToAuditFile adds records to the audit file based upon newly refreshed secrets">`addToAuditFile`</SwmToken> function adds records to the audit file based on newly refreshed secrets. It sorts the handles, scrubs sensitive values, and appends the new records to the audit file.

```go
// addToAuditFile adds records to the audit file based upon newly refreshed secrets
func (r *secretResolver) addToAuditFile(secretResponse map[string]string) error {
	if r.auditFilename == "" {
		return nil
	}
	if r.auditRotRecs == nil {
		r.auditRotRecs = newRotatingNDRecords(r.auditFilename, config{})
	}

	// iterate keys in deterministic order by sorting
	handles := make([]string, 0, len(secretResponse))
	for handle := range secretResponse {
		handles = append(handles, handle)
	}
	sort.Strings(handles)

	var newRows []auditRecord
	// add the newly refreshed secrets to the list of rows
	for _, handle := range handles {
		secretValue := secretResponse[handle]
		scrubbedValue := ""
```

---

</SwmSnippet>

<SwmSnippet path="/comp/core/secrets/secretsimpl/fetch_secret.go" line="105">

---

## <SwmToken path="comp/core/secrets/secretsimpl/fetch_secret.go" pos="105:2:2" line-data="// fetchSecret receives a list of secrets name to fetch, exec a custom">`fetchSecret`</SwmToken>

The <SwmToken path="comp/core/secrets/secretsimpl/fetch_secret.go" pos="105:2:2" line-data="// fetchSecret receives a list of secrets name to fetch, exec a custom">`fetchSecret`</SwmToken> function executes a custom command to fetch the actual secrets. It serializes the secret handles into a JSON payload, executes the command using <SwmToken path="comp/core/secrets/secretsimpl/fetch_secret.go" pos="116:10:10" line-data="	output, err := r.execCommand(string(jsonPayload))">`execCommand`</SwmToken>, and unmarshals the output into a map of secrets.

```go
// fetchSecret receives a list of secrets name to fetch, exec a custom
// executable to fetch the actual secrets and returns them.
func (r *secretResolver) fetchSecret(secretsHandle []string) (map[string]string, error) {
	payload := map[string]interface{}{
		"version": secrets.PayloadVersion,
		"secrets": secretsHandle,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("could not serialize secrets IDs to fetch password: %s", err)
	}
	output, err := r.execCommand(string(jsonPayload))
	if err != nil {
		return nil, err
	}

	secrets := map[string]secrets.SecretVal{}
	err = json.Unmarshal(output, &secrets)
	if err != nil {
		r.tlmSecretUnmarshalError.Inc()
		return nil, fmt.Errorf("could not unmarshal 'secret_backend_command' output: %s", err)
```

---

</SwmSnippet>

<SwmSnippet path="/comp/core/secrets/secretsimpl/fetch_secret.go" line="35">

---

### Executing Command

The <SwmToken path="comp/core/secrets/secretsimpl/fetch_secret.go" pos="35:9:9" line-data="func (r *secretResolver) execCommand(inputPayload string) ([]byte, error) {">`execCommand`</SwmToken> function runs the backend command to fetch secrets. It sets up the command context, handles input and output streams, and logs the execution details. If the command fails, it logs the error and returns it.

```go
func (r *secretResolver) execCommand(inputPayload string) ([]byte, error) {
	// hook used only for tests
	if r.commandHookFunc != nil {
		return r.commandHookFunc(inputPayload)
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(r.backendTimeout)*time.Second)
	defer cancel()

	cmd, done, err := commandContext(ctx, r.backendCommand, r.backendArguments...)
	if err != nil {
		return nil, err
	}
	defer done()

	if err := checkRights(cmd.Path, r.commandAllowGroupExec); err != nil {
		return nil, err
	}

	cmd.Stdin = strings.NewReader(inputPayload)
```

---

</SwmSnippet>

<SwmSnippet path="/comp/core/secrets/secretsimpl/rotating_ndrecords.go" line="67">

---

## Add

The <SwmToken path="comp/core/secrets/secretsimpl/rotating_ndrecords.go" pos="67:2:2" line-data="// Add adds a new record to the file with the given time and payload">`Add`</SwmToken> function adds a new record to the audit file with the given time and payload. It prunes old entries, rotates the file if it exceeds the size limit, and appends the new record to the file.

```go
// Add adds a new record to the file with the given time and payload
// old entries will be pruned, and the file will be rotated if it gets too large
func (r *rotatingNDRecords) Add(t time.Time, payload interface{}) error {
	r.ensureDefaults()

	// prune old entries
	if !r.firstEntry.IsZero() && t.Sub(r.firstEntry) > r.cfg.retention {
		if err := r.pruneOldEntries(t); err != nil {
			log.Error(err)
		}
	}
	// remove old files that were already rotated
	if !r.oldestFileMtime.IsZero() && t.Sub(*r.oldestFileMtime) > r.cfg.retention {
		if err := r.removeOldFiles(t); err != nil {
			log.Error(err)
		}
	}

	var recordData bytes.Buffer
	err := json.NewEncoder(&recordData).Encode(ndRecord{
		Time: t,
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
