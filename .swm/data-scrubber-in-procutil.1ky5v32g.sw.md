---
title: Data Scrubber in Procutil
---
```mermaid
graph TD;
 A[Data Scrubber Initialization] --> B[Compile Default Sensitive Words];
 B --> C[Create Data Scrubber Instance];
 C --> D[Scrub Command Lines];
 D --> E[Cache Scrubbed Command Lines];
 E --> F[Reset Cache Periodically];
 D --> G[Strip All Arguments (Optional)];
```

# Overview

The Data Scrubber in Procutil is a component that allows the Datadog Agent to <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="32:12:14" line-data="// DataScrubber allows the agent to disallow-list cmdline arguments that match">`disallow-list`</SwmToken> command-line arguments that match a list of predefined and custom sensitive words. It uses regular expressions to identify and scrub sensitive information from command-line arguments, maintaining a cache of seen processes and their scrubbed command lines to optimize performance.

# Initialization

The Data Scrubber is initialized with default sensitive words and can be customized with additional sensitive words. The <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="44:2:2" line-data="// NewDefaultDataScrubber creates a DataScrubber with the default behavior: enabled">`NewDefaultDataScrubber`</SwmToken> function creates a <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="130:6:6" line-data="func (ds *DataScrubber) ScrubProcessCommand(p *Process) []string {">`DataScrubber`</SwmToken> with the default behavior: enabled and matching the default sensitive words. It initializes the <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="130:6:6" line-data="func (ds *DataScrubber) ScrubProcessCommand(p *Process) []string {">`DataScrubber`</SwmToken> with default sensitive patterns, an empty cache for seen processes and scrubbed command lines, and sets the cache cycle parameters.

# Scrubbing Command Lines

The <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="128:2:2" line-data="// ScrubProcessCommand uses a cache memory to avoid scrubbing already known">`ScrubProcessCommand`</SwmToken> method scrubs the command-line arguments of a given process, using the cache to avoid redundant scrubbing of already known processes. It first checks if all arguments should be stripped or if the scrubber is disabled. If not, it creates a unique process key and checks if the process has been seen before. If not, it scrubs the command line and stores the result in the cache.

<SwmSnippet path="/pkg/process/procutil/data_scrubber.go" line="128">

---

The <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="128:2:2" line-data="// ScrubProcessCommand uses a cache memory to avoid scrubbing already known">`ScrubProcessCommand`</SwmToken> method uses a cache to avoid scrubbing already known process command lines.

```go
// ScrubProcessCommand uses a cache memory to avoid scrubbing already known
// process' cmdlines
func (ds *DataScrubber) ScrubProcessCommand(p *Process) []string {
	if ds.StripAllArguments {
		return ds.stripArguments(p.Cmdline)
	}

	if !ds.Enabled {
		return p.Cmdline
	}

	pKey := createProcessKey(p)
	if _, ok := ds.seenProcess[pKey]; !ok {
		ds.seenProcess[pKey] = struct{}{}
		if scrubbed, changed := ds.ScrubCommand(p.Cmdline); changed {
			ds.scrubbedCmdlines[pKey] = scrubbed
		}
	}

	if scrubbed, ok := ds.scrubbedCmdlines[pKey]; ok {
		return scrubbed
```

---

</SwmSnippet>

# Scrubbing Sensitive Information

The <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="142:12:12" line-data="		if scrubbed, changed := ds.ScrubCommand(p.Cmdline); changed {">`ScrubCommand`</SwmToken> method hides the argument value for any key which matches a sensitive word pattern. It returns the updated command line and a boolean indicating whether any changes were made. It performs a fast check using a direct pattern and then uses regular expressions to replace sensitive values with asterisks.

<SwmSnippet path="/pkg/process/procutil/data_scrubber.go" line="164">

---

The <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="164:2:2" line-data="// ScrubCommand hides the argument value for any key which matches a &quot;sensitive word&quot; pattern.">`ScrubCommand`</SwmToken> method hides the argument value for any key which matches a sensitive word pattern.

```go
// ScrubCommand hides the argument value for any key which matches a "sensitive word" pattern.
// It returns the updated cmdline, as well as a boolean representing whether it was scrubbed
func (ds *DataScrubber) ScrubCommand(cmdline []string) ([]string, bool) {
	newCmdline := cmdline
	rawCmdline := strings.Join(cmdline, " ")
	lowerCaseCmdline := strings.ToLower(rawCmdline)
	changed := false
	for _, pattern := range ds.SensitivePatterns {
		// fast check with direct pattern
		if !strings.Contains(lowerCaseCmdline, pattern.FastCheck) {
			continue
		}

		if pattern.Re.MatchString(rawCmdline) {
			changed = true
			rawCmdline = pattern.Re.ReplaceAllString(rawCmdline, "${key}${delimiter}********")
		}
	}

	if changed {
```

---

</SwmSnippet>

# Compiling Sensitive Words

The <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="60:2:2" line-data="// CompileStringsToRegex compile each word in the slice into a regex pattern to match">`CompileStringsToRegex`</SwmToken> function compiles a list of sensitive words into regular expression patterns that the Data Scrubber uses to identify sensitive information in command-line arguments.

<SwmSnippet path="/pkg/process/procutil/data_scrubber.go" line="60">

---

The <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="60:2:2" line-data="// CompileStringsToRegex compile each word in the slice into a regex pattern to match">`CompileStringsToRegex`</SwmToken> function compiles a list of sensitive words into regular expression patterns.

```go
// CompileStringsToRegex compile each word in the slice into a regex pattern to match
// against the cmdline arguments
// The word must contain only word characters ([a-zA-z0-9_]) or wildcards *
func CompileStringsToRegex(words []string) []DataScrubberPattern {
	compiledRegexps := make([]DataScrubberPattern, 0, len(words))

	// forbiddenSymbolsRegex defined in `data_scrubber_<platform>.go` because it's platform dependent
	forbiddenSymbols := regexp.MustCompile(forbiddenSymbolsRegex)

	for _, word := range words {
		if forbiddenSymbols.MatchString(word) {
			log.Warnf("data scrubber: %s skipped. The sensitive word must "+
				"contain only alphanumeric characters, underscores or wildcards ('*')", word)
			continue
		}

		if word == "*" {
			log.Warn("data scrubber: ignoring wildcard-only ('*') sensitive word as it is not supported")
			continue
		}
```

---

</SwmSnippet>

# Cache Management

The Data Scrubber maintains a cache of seen processes and their scrubbed command lines to optimize performance. The cache is periodically reset to ensure that it does not grow indefinitely. The <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="153:2:2" line-data="// IncrementCacheAge increments one cycle of cache memory age. If it reaches">`IncrementCacheAge`</SwmToken> method increments the cache age and resets the cache if it reaches the maximum number of cycles.

<SwmSnippet path="/pkg/process/procutil/data_scrubber.go" line="153">

---

The <SwmToken path="pkg/process/procutil/data_scrubber.go" pos="153:2:2" line-data="// IncrementCacheAge increments one cycle of cache memory age. If it reaches">`IncrementCacheAge`</SwmToken> method increments the cache age and resets the cache if it reaches the maximum number of cycles.

```go
// IncrementCacheAge increments one cycle of cache memory age. If it reaches
// cacheMaxCycles, the cache is restarted
func (ds *DataScrubber) IncrementCacheAge() {
	ds.cacheCycles++
	if ds.cacheCycles == ds.cacheMaxCycles {
		ds.seenProcess = make(map[processCacheKey]struct{})
		ds.scrubbedCmdlines = make(map[processCacheKey][]string)
		ds.cacheCycles = 0
	}
}
```

---

</SwmSnippet>

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
