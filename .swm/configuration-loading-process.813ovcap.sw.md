---
title: Configuration Loading Process
---
This document explains the 'Load' process, which is responsible for loading configuration settings from a specified file path. The process involves creating a new serverless configuration, loading the configuration file, applying Datadog-specific settings, and validating the configuration.

The 'Load' process starts by creating a new serverless configuration. It then loads the configuration file from the specified path. After loading the file, it applies various Datadog-specific settings such as API keys and hostnames. Finally, it validates the configuration to ensure all necessary settings are correctly applied.

Here is a high level diagram of the flow, showing only the most important functions:

```mermaid
graph TD;
      subgraph comptraceconfigsetupgo["comp/trace/config/setup.go"]
0fa7ee0105a3a088df2c035118be72d930d1d639507257926621f137a5fd3ac4(Load):::mainFlowStyle --> 38a584d4c462a7afbc11c112f91f4cd7127b94b2595f4eabe5a92d173e69ce54(LoadConfigFile)
end

subgraph compcoreconfig["comp/core/config"]
0fa7ee0105a3a088df2c035118be72d930d1d639507257926621f137a5fd3ac4(Load):::mainFlowStyle --> 69ac91819173163d803e619eee4f740a34dc3da914054ea16529c9dd53ca9a89(NewServerlessConfig):::mainFlowStyle
end

subgraph compcoreconfig["comp/core/config"]
69ac91819173163d803e619eee4f740a34dc3da914054ea16529c9dd53ca9a89(NewServerlessConfig):::mainFlowStyle --> 3bee3fb5fc6f2458a33bfa1f10b8e201027c1f335ed35c3c567fe2578187e81a(newConfig):::mainFlowStyle
end

subgraph compcoreconfig["comp/core/config"]
3bee3fb5fc6f2458a33bfa1f10b8e201027c1f335ed35c3c567fe2578187e81a(newConfig):::mainFlowStyle --> 3c0eab777a5cd2ff0f7b9e9c47b219916ba76d590e2a8071173df8a880aac7c9(setupConfig):::mainFlowStyle
end

subgraph pkgconfig["pkg/config"]
3c0eab777a5cd2ff0f7b9e9c47b219916ba76d590e2a8071173df8a880aac7c9(setupConfig):::mainFlowStyle --> 463bbd3c45c10f62b3fe1ea96949792ee977894dd1e5c92c0c79461a7f65a8df(LoadWithSecret):::mainFlowStyle
end

subgraph pkgconfig["pkg/config"]
463bbd3c45c10f62b3fe1ea96949792ee977894dd1e5c92c0c79461a7f65a8df(LoadWithSecret):::mainFlowStyle --> 580e5d3d5ba5a3338f7b9cc1ef26f2e864292b17580cf5c51d5a2c09ce46b784(LoadDatadogCustom):::mainFlowStyle
end

subgraph pkgconfig["pkg/config"]
580e5d3d5ba5a3338f7b9cc1ef26f2e864292b17580cf5c51d5a2c09ce46b784(LoadDatadogCustom):::mainFlowStyle --> 3e92ef05633abb62bb131742ac41ac03be98a9cb1bee70b81ffe6ae0ae0f3de9(DetectFeatures):::mainFlowStyle
end

subgraph pkgconfig["pkg/config"]
3e92ef05633abb62bb131742ac41ac03be98a9cb1bee70b81ffe6ae0ae0f3de9(DetectFeatures):::mainFlowStyle --> 73864f3c169d2cbecf44f3c8cfe8e0fbde175d8d6e13f00e3d1a2a41fb9a10ad(excludeFeatures)
end


      classDef mainFlowStyle color:#000000,fill:#7CB9F4
classDef rootsStyle color:#000000,fill:#00FFF4
classDef Style1 color:#000000,fill:#00FFAA
classDef Style2 color:#000000,fill:#FFFF00
classDef Style3 color:#000000,fill:#AA7CB9
```

# Flow drill down

First, we'll zoom into this section of the flow:

```mermaid
graph TD;
      0fa7ee0105a3a088df2c035118be72d930d1d639507257926621f137a5fd3ac4(Load):::mainFlowStyle --> 38a584d4c462a7afbc11c112f91f4cd7127b94b2595f4eabe5a92d173e69ce54(LoadConfigFile)

0fa7ee0105a3a088df2c035118be72d930d1d639507257926621f137a5fd3ac4(Load):::mainFlowStyle --> 69ac91819173163d803e619eee4f740a34dc3da914054ea16529c9dd53ca9a89(NewServerlessConfig):::mainFlowStyle

69ac91819173163d803e619eee4f740a34dc3da914054ea16529c9dd53ca9a89(NewServerlessConfig):::mainFlowStyle --> wjpet(...)

38a584d4c462a7afbc11c112f91f4cd7127b94b2595f4eabe5a92d173e69ce54(LoadConfigFile) --> 5f7f71114f6b6916ed3c8e4a9f7f23db21b351019d2b446bfb4e0065156e1d58(applyDatadogConfig)

38a584d4c462a7afbc11c112f91f4cd7127b94b2595f4eabe5a92d173e69ce54(LoadConfigFile) --> 404dc35699efd326d1a928ae2728a434ebd755ab83a59d92de28f5a77ca6d013(prepareConfig)

38a584d4c462a7afbc11c112f91f4cd7127b94b2595f4eabe5a92d173e69ce54(LoadConfigFile) --> c7c1aa1bc7a6267beb3c9a67bf1b2fe41748275b1b407c56ec485c226c781751(validate)


      classDef mainFlowStyle color:#000000,fill:#7CB9F4
classDef rootsStyle color:#000000,fill:#00FFF4
classDef Style1 color:#000000,fill:#00FFAA
classDef Style2 color:#000000,fill:#FFFF00
classDef Style3 color:#000000,fill:#AA7CB9
```

<SwmSnippet path="/pkg/serverless/trace/trace.go" line="83">

---

## Load

The <SwmToken path="pkg/serverless/trace/trace.go" pos="83:2:2" line-data="// Load loads the config from a file path">`Load`</SwmToken> function is responsible for loading the configuration from a specified file path. It first creates a new serverless configuration using <SwmToken path="pkg/serverless/trace/trace.go" pos="85:10:10" line-data="	c, err := compcorecfg.NewServerlessConfig(l.Path)">`NewServerlessConfig`</SwmToken>. If successful, it then calls <SwmToken path="pkg/serverless/trace/trace.go" pos="89:5:5" line-data="	return comptracecfg.LoadConfigFile(l.Path, c)">`LoadConfigFile`</SwmToken> to load the configuration file.

```go
// Load loads the config from a file path
func (l *LoadConfig) Load() (*config.AgentConfig, error) {
	c, err := compcorecfg.NewServerlessConfig(l.Path)
	if err != nil {
		return nil, err
	}
	return comptracecfg.LoadConfigFile(l.Path, c)
}
```

---

</SwmSnippet>

<SwmSnippet path="/comp/trace/config/setup.go" line="59">

---

## <SwmToken path="comp/trace/config/setup.go" pos="59:2:2" line-data="// LoadConfigFile returns a new configuration based on the given path. The path must not necessarily exist">`LoadConfigFile`</SwmToken>

The <SwmToken path="comp/trace/config/setup.go" pos="59:2:2" line-data="// LoadConfigFile returns a new configuration based on the given path. The path must not necessarily exist">`LoadConfigFile`</SwmToken> function returns a new configuration based on the given path. It uses <SwmToken path="comp/trace/config/setup.go" pos="63:8:8" line-data="	cfg, err := prepareConfig(c)">`prepareConfig`</SwmToken> to set up the initial configuration and then applies Datadog-specific settings using <SwmToken path="comp/trace/config/setup.go" pos="73:7:7" line-data="	if err := applyDatadogConfig(cfg, c); err != nil {">`applyDatadogConfig`</SwmToken>. Finally, it validates the configuration using the <SwmToken path="comp/trace/config/setup.go" pos="77:6:6" line-data="	return cfg, validate(cfg, c)">`validate`</SwmToken> function.

```go
// LoadConfigFile returns a new configuration based on the given path. The path must not necessarily exist
// and a valid configuration can be returned based on defaults and environment variables. If a
// valid configuration can not be obtained, an error is returned.
func LoadConfigFile(path string, c corecompcfg.Component) (*config.AgentConfig, error) {
	cfg, err := prepareConfig(c)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	} else {
		cfg.ConfigPath = path
		log.Infof("Loaded configuration: %s", cfg.ConfigPath)
	}

	if err := applyDatadogConfig(cfg, c); err != nil {
		log.Error(err)
	}

	return cfg, validate(cfg, c)
```

---

</SwmSnippet>

<SwmSnippet path="/comp/trace/config/setup.go" line="145">

---

### <SwmToken path="comp/trace/config/setup.go" pos="145:2:2" line-data="func applyDatadogConfig(c *config.AgentConfig, core corecompcfg.Component) error {">`applyDatadogConfig`</SwmToken>

The <SwmToken path="comp/trace/config/setup.go" pos="145:2:2" line-data="func applyDatadogConfig(c *config.AgentConfig, core corecompcfg.Component) error {">`applyDatadogConfig`</SwmToken> function applies various Datadog-specific settings to the configuration. This includes setting API keys, hostnames, proxy settings, and other APM-related configurations. It ensures that all necessary settings are correctly applied to the configuration object.

```go
func applyDatadogConfig(c *config.AgentConfig, core corecompcfg.Component) error {
	if len(c.Endpoints) == 0 {
		c.Endpoints = []*config.Endpoint{{}}
	}
	if core.IsSet("api_key") {
		c.Endpoints[0].APIKey = utils.SanitizeAPIKey(coreconfig.Datadog().GetString("api_key"))
	}
	if core.IsSet("hostname") {
		c.Hostname = core.GetString("hostname")
	}
	if core.IsSet("dogstatsd_port") {
		c.StatsdPort = core.GetInt("dogstatsd_port")
	}

	obsPipelineEnabled, prefix := isObsPipelineEnabled(core)
	if obsPipelineEnabled {
		if host := core.GetString(fmt.Sprintf("%s.traces.url", prefix)); host == "" {
			log.Errorf("%s.traces.enabled but %s.traces.url is empty.", prefix, prefix)
		} else {
			c.Endpoints[0].Host = host
		}
```

---

</SwmSnippet>

<SwmSnippet path="/comp/trace/config/setup.go" line="80">

---

### <SwmToken path="comp/trace/config/setup.go" pos="80:2:2" line-data="func prepareConfig(c corecompcfg.Component) (*config.AgentConfig, error) {">`prepareConfig`</SwmToken>

The <SwmToken path="comp/trace/config/setup.go" pos="80:2:2" line-data="func prepareConfig(c corecompcfg.Component) (*config.AgentConfig, error) {">`prepareConfig`</SwmToken> function initializes a new configuration object with default values and settings. It also sets up logging, proxy settings, and remote configuration management if enabled.

```go
func prepareConfig(c corecompcfg.Component) (*config.AgentConfig, error) {
	cfg := config.New()
	cfg.DDAgentBin = defaultDDAgentBin
	cfg.AgentVersion = version.AgentVersion
	cfg.GitCommit = version.Commit
	cfg.ReceiverSocket = defaultReceiverSocket

	// the core config can be assumed to already be set-up as it has been
	// injected as a component dependency
	// TODO: do not interface directly with pkg/config anywhere
	coreConfigObject := c.Object()
	if coreConfigObject == nil {
		return nil, errors.New("no core config found! Bailing out")
	}

	if !coreConfigObject.GetBool("disable_file_logging") {
		cfg.LogFilePath = DefaultLogFilePath
	}

	ipcAddress, err := coreconfig.GetIPCAddress()
	if err != nil {
```

---

</SwmSnippet>

<SwmSnippet path="/comp/trace/config/setup.go" line="803">

---

### validate

The <SwmToken path="comp/trace/config/setup.go" pos="803:2:2" line-data="// validate validates if the current configuration is good for the agent to start with.">`validate`</SwmToken> function checks if the current configuration is valid for the agent to start. It ensures that essential settings like API keys and agent binary paths are set and attempts to resolve the hostname if not provided.

```go
// validate validates if the current configuration is good for the agent to start with.
func validate(c *config.AgentConfig, core corecompcfg.Component) error {
	if len(c.Endpoints) == 0 || c.Endpoints[0].APIKey == "" {
		return config.ErrMissingAPIKey
	}
	if c.DDAgentBin == "" {
		return errors.New("agent binary path not set")
	}

	if c.Hostname == "" && !core.GetBool("serverless.enabled") {
		if err := hostname(c); err != nil {
			return err
		}
	}
	return nil
}
```

---

</SwmSnippet>

Now, lets zoom into this section of the flow:

```mermaid
graph TD;
      subgraph compcoreconfig["comp/core/config"]
69ac91819173163d803e619eee4f740a34dc3da914054ea16529c9dd53ca9a89(NewServerlessConfig):::mainFlowStyle --> 3bee3fb5fc6f2458a33bfa1f10b8e201027c1f335ed35c3c567fe2578187e81a(newConfig):::mainFlowStyle
end

subgraph compcoreconfig["comp/core/config"]
3bee3fb5fc6f2458a33bfa1f10b8e201027c1f335ed35c3c567fe2578187e81a(newConfig):::mainFlowStyle --> 3c0eab777a5cd2ff0f7b9e9c47b219916ba76d590e2a8071173df8a880aac7c9(setupConfig):::mainFlowStyle
end

subgraph pkgconfig["pkg/config"]
3c0eab777a5cd2ff0f7b9e9c47b219916ba76d590e2a8071173df8a880aac7c9(setupConfig):::mainFlowStyle --> 463bbd3c45c10f62b3fe1ea96949792ee977894dd1e5c92c0c79461a7f65a8df(LoadWithSecret):::mainFlowStyle
end

subgraph pkgconfig["pkg/config"]
463bbd3c45c10f62b3fe1ea96949792ee977894dd1e5c92c0c79461a7f65a8df(LoadWithSecret):::mainFlowStyle --> 580e5d3d5ba5a3338f7b9cc1ef26f2e864292b17580cf5c51d5a2c09ce46b784(LoadDatadogCustom):::mainFlowStyle
end

subgraph pkgconfig["pkg/config"]
580e5d3d5ba5a3338f7b9cc1ef26f2e864292b17580cf5c51d5a2c09ce46b784(LoadDatadogCustom):::mainFlowStyle --> 3e92ef05633abb62bb131742ac41ac03be98a9cb1bee70b81ffe6ae0ae0f3de9(DetectFeatures):::mainFlowStyle
end


      classDef mainFlowStyle color:#000000,fill:#7CB9F4
classDef rootsStyle color:#000000,fill:#00FFF4
classDef Style1 color:#000000,fill:#00FFAA
classDef Style2 color:#000000,fill:#FFFF00
classDef Style3 color:#000000,fill:#AA7CB9
```

<SwmSnippet path="/comp/core/config/config.go" line="64">

---

## <SwmToken path="comp/core/config/config.go" pos="64:2:2" line-data="// NewServerlessConfig initializes a config component from the given config file">`NewServerlessConfig`</SwmToken> Initialization

The <SwmToken path="comp/core/config/config.go" pos="64:2:2" line-data="// NewServerlessConfig initializes a config component from the given config file">`NewServerlessConfig`</SwmToken> function initializes a serverless configuration component. It first sets up configuration options, checks if the provided path exists, and adjusts the options accordingly. Finally, it creates a new <SwmToken path="comp/core/config/config.go" pos="77:5:5" line-data="	d := dependencies{Params: NewParams(path, options...)}">`dependencies`</SwmToken> object and calls <SwmToken path="comp/core/config/config.go" pos="78:3:3" line-data="	return newConfig(d)">`newConfig`</SwmToken> to proceed with the configuration setup.

```go
// NewServerlessConfig initializes a config component from the given config file
// TODO: serverless must be eventually migrated to fx, this workaround will then become obsolete - ts should not be created directly in this fashion.
func NewServerlessConfig(path string) (Component, error) {
	options := []func(*Params){WithConfigName("serverless")}

	_, err := os.Stat(path)
	if os.IsNotExist(err) &&
		(strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml")) {
		options = append(options, WithConfigMissingOK(true))
	} else if !os.IsNotExist(err) {
		options = append(options, WithConfFilePath(path))
	}

	d := dependencies{Params: NewParams(path, options...)}
	return newConfig(d)
```

---

</SwmSnippet>

<SwmSnippet path="/comp/core/config/config.go" line="89">

---

## Configuration Setup with <SwmToken path="comp/core/config/config.go" pos="89:2:2" line-data="func newConfig(deps dependencies) (*cfg, error) {">`newConfig`</SwmToken>

The <SwmToken path="comp/core/config/config.go" pos="89:2:2" line-data="func newConfig(deps dependencies) (*cfg, error) {">`newConfig`</SwmToken> function is responsible for setting up the configuration using the provided dependencies. It initializes the Datadog configuration, calls <SwmToken path="comp/core/config/config.go" pos="91:8:8" line-data="	warnings, err := setupConfig(config, deps)">`setupConfig`</SwmToken> to load the configuration, and handles any warnings or errors that occur during the process. If security agent configuration is required, it merges the necessary settings.

```go
func newConfig(deps dependencies) (*cfg, error) {
	config := pkgconfigsetup.Datadog()
	warnings, err := setupConfig(config, deps)
	returnErrFct := func(e error) (*cfg, error) {
		if e != nil && deps.Params.ignoreErrors {
			if warnings == nil {
				warnings = &pkgconfigmodel.Warnings{}
			}
			warnings.Err = e
			e = nil
		}
		return &cfg{Config: config, warnings: warnings}, e
	}

	if err != nil {
		return returnErrFct(err)
	}

	if deps.Params.configLoadSecurityAgent {
		if err := pkgconfigsetup.Merge(deps.Params.securityAgentConfigFilePaths, config); err != nil {
			return returnErrFct(err)
```

---

</SwmSnippet>

<SwmSnippet path="/comp/core/config/setup.go" line="21">

---

### Detailed Configuration Setup

The <SwmToken path="comp/core/config/setup.go" pos="21:2:2" line-data="// setupConfig is copied from cmd/agent/common/helpers.go.">`setupConfig`</SwmToken> function performs the detailed setup of the configuration. It sets the configuration file paths, loads extra configuration paths, and handles the loading of the configuration with or without secrets. It also manages errors related to missing or inaccessible configuration files.

```go
// setupConfig is copied from cmd/agent/common/helpers.go.
func setupConfig(config pkgconfigmodel.Config, deps configDependencies) (*pkgconfigmodel.Warnings, error) {
	p := deps.getParams()

	confFilePath := p.ConfFilePath
	configName := p.configName
	failOnMissingFile := !p.configMissingOK
	defaultConfPath := p.defaultConfPath

	if configName != "" {
		config.SetConfigName(configName)
	}

	// set the paths where a config file is expected
	if len(confFilePath) != 0 {
		// if the configuration file path was supplied on the command line,
		// add that first so it's first in line
		config.AddConfigPath(confFilePath)
		// If they set a config file directly, let's try to honor that
		if strings.HasSuffix(confFilePath, ".yaml") || strings.HasSuffix(confFilePath, ".yml") {
			config.SetConfigFile(confFilePath)
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/config/setup/config.go" line="1655">

---

## Loading Configuration with Secrets

The <SwmToken path="pkg/config/setup/config.go" pos="1655:2:2" line-data="// LoadWithSecret reads config files and initializes config with decrypted secrets">`LoadWithSecret`</SwmToken> function reads the configuration files and initializes the configuration with decrypted secrets. It delegates the actual loading to <SwmToken path="pkg/config/setup/config.go" pos="1657:3:3" line-data="	return LoadDatadogCustom(config, &quot;datadog.yaml&quot;, optional.NewOption[secrets.Component](secretResolver), additionalEnvVars)">`LoadDatadogCustom`</SwmToken>.

```go
// LoadWithSecret reads config files and initializes config with decrypted secrets
func LoadWithSecret(config pkgconfigmodel.Config, secretResolver secrets.Component, additionalEnvVars []string) (*pkgconfigmodel.Warnings, error) {
	return LoadDatadogCustom(config, "datadog.yaml", optional.NewOption[secrets.Component](secretResolver), additionalEnvVars)
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/config/setup/config.go" line="1837">

---

## Custom Configuration Loading

The <SwmToken path="pkg/config/setup/config.go" pos="1837:2:2" line-data="// LoadDatadogCustom loads the datadog config in the given config">`LoadDatadogCustom`</SwmToken> function loads the Datadog configuration and performs various setup tasks, including feature detection, proxy settings, secret resolution, and conflict checks. It ensures that the configuration is properly initialized and ready for use.

```go
// LoadDatadogCustom loads the datadog config in the given config
func LoadDatadogCustom(config pkgconfigmodel.Config, origin string, secretResolver optional.Option[secrets.Component], additionalKnownEnvVars []string) (*pkgconfigmodel.Warnings, error) {
	// Feature detection running in a defer func as it always  need to run (whether config load has been successful or not)
	// Because some Agents (e.g. trace-agent) will run even if config file does not exist
	defer func() {
		// Environment feature detection needs to run before applying override funcs
		// as it may provide such overrides
		pkgconfigenv.DetectFeatures(config)
		pkgconfigmodel.ApplyOverrideFuncs(config)
	}()

	warnings := &pkgconfigmodel.Warnings{}
	err := LoadCustom(config, additionalKnownEnvVars)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			log.Warnf("Error loading config: %v (check config file permissions for dd-agent user)", err)
		} else {
			log.Warnf("Error loading config: %v", err)
		}
		return warnings, err
	}
```

---

</SwmSnippet>

<SwmSnippet path="/pkg/config/env/environment_detection.go" line="96">

---

## Feature Detection

The <SwmToken path="pkg/config/env/environment_detection.go" pos="96:2:2" line-data="// DetectFeatures runs the feature detection.">`DetectFeatures`</SwmToken> function runs the feature detection process. It ensures that the Datadog configuration is fully loaded before detecting features based on the environment. It updates the detected features and logs the results.

```go
// DetectFeatures runs the feature detection.
// We guarantee that Datadog configuration is entirely loaded (env + YAML)
// before this function is called
func DetectFeatures(cfg model.Reader) {
	featureLock.Lock()
	defer featureLock.Unlock()

	// Detection should not run in unit tests to avoid overriding features based on runner environment
	if detectionAlwaysDisabledInTests {
		return
	}

	newFeatures := make(FeatureMap)
	if IsAutoconfigEnabled(cfg) {
		detectContainerFeatures(newFeatures, cfg)
		excludedFeatures := cfg.GetStringSlice("autoconfig_exclude_features")
		excludeFeatures(newFeatures, excludedFeatures)

		includedFeatures := cfg.GetStringSlice("autoconfig_include_features")
		for _, f := range includedFeatures {
			f = strings.ToLower(f)
```

---

</SwmSnippet>

# Where is this flow used?

This flow is used multiple times in the codebase as represented in the following diagram:

```mermaid
graph TD;
      9622d44e969da5c046d7d1509f8ea860b847ee32f7b9e8ff6d8d7e779e9915f3(runAgent):::rootsStyle --> fbe9d61b43618a7cc32a758dfb8bbf1cc2a1f416600ca7a464d18d448c95325b(setupApiKey)

subgraph pkgserverless["pkg/serverless"]
fbe9d61b43618a7cc32a758dfb8bbf1cc2a1f416600ca7a464d18d448c95325b(setupApiKey) --> 544ccc9ab2236ae14f69cbfe580f5fefd000d3d8c6a779543c9fbca9d12876ed(NewWithShutdown)
end

subgraph pkgserverless["pkg/serverless"]
544ccc9ab2236ae14f69cbfe580f5fefd000d3d8c6a779543c9fbca9d12876ed(NewWithShutdown) --> 12ace7adc7424788dc53055f6b240eff573e8785d98864fca8b0d3cc7d7369c4(newAppSec)
end

subgraph pkgserverless["pkg/serverless"]
12ace7adc7424788dc53055f6b240eff573e8785d98864fca8b0d3cc7d7369c4(newAppSec) --> fd8332cdfcdd262463b6a55bbe46c2c7f910c37b4fbfe2d68467754d8e7bcb16(wafHealth)
end

subgraph pkgserverless["pkg/serverless"]
fd8332cdfcdd262463b6a55bbe46c2c7f910c37b4fbfe2d68467754d8e7bcb16(wafHealth) --> 0fa7ee0105a3a088df2c035118be72d930d1d639507257926621f137a5fd3ac4(Load):::mainFlowStyle
end

subgraph pkgserverless["pkg/serverless"]
063108a363d91ad7e53ed07c21eea179c5621cd8daff5d9feb79d8d825d9e14e(init):::rootsStyle --> 80c296e665756a9b8d2ff2d4ba70b58be50ccd36b7862d12fae142313b6b6519(New)
end

subgraph pkgserverless["pkg/serverless"]
80c296e665756a9b8d2ff2d4ba70b58be50ccd36b7862d12fae142313b6b6519(New) --> 544ccc9ab2236ae14f69cbfe580f5fefd000d3d8c6a779543c9fbca9d12876ed(NewWithShutdown)
end

e60dd21554cef837462f4efb6b6ea804e556a200dc4da10a2d0e442f279f9ffe(run):::rootsStyle --> e162c81d0025aea353a3e09f3d926d94c61095598a8f6076a5511cdbb2172015(setup)

e162c81d0025aea353a3e09f3d926d94c61095598a8f6076a5511cdbb2172015(setup) --> 027c86a32facdffa2d05edd65ecaa4ffdac97aa85cab8d1096259b5d245dcbdc(setupTraceAgent)

subgraph pkgserverless["pkg/serverless"]
027c86a32facdffa2d05edd65ecaa4ffdac97aa85cab8d1096259b5d245dcbdc(setupTraceAgent) --> 151f8a70b5bae6fffab3c0e5977e18ae241b927e358e049a4e6ac5582fd02b3c(StartServerlessTraceAgent)
end

subgraph pkgserverless["pkg/serverless"]
151f8a70b5bae6fffab3c0e5977e18ae241b927e358e049a4e6ac5582fd02b3c(StartServerlessTraceAgent) --> 0fa7ee0105a3a088df2c035118be72d930d1d639507257926621f137a5fd3ac4(Load):::mainFlowStyle
end


      classDef mainFlowStyle color:#000000,fill:#7CB9F4
classDef rootsStyle color:#000000,fill:#00FFF4
classDef Style1 color:#000000,fill:#00FFAA
classDef Style2 color:#000000,fill:#FFFF00
classDef Style3 color:#000000,fill:#AA7CB9
```

&nbsp;

*This is an auto-generated document by Swimm AI 🌊 and has not yet been verified by a human*

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBZGF0YWRvZy1hZ2VudCUzQSUzQVN3aW1tLURlbW8=" repo-name="datadog-agent"><sup>Powered by [Swimm](/)</sup></SwmMeta>
