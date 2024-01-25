package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	KAPETA_CLUSTER_SERVICE_CONFIG_FILE  = "cluster-service.yml"
	KAPETA_CLUSTER_SERVICE_DEFAULT_PORT = "35100"
	KAPETA_CLUSTER_SERVICE_DEFAULT_HOST = "127.0.0.1"
	KAPETA_CLUSTER_SERVICE_DEFAULT_ENV  = "production"
)

var (
	PROVIDER_TYPES = []string{
		"core/block-type",
		"core/block-type-operator",
		"core/resource-type-extension",
		"core/resource-type-internal",
		"core/resource-type-operator",
		"core/language-target",
		"core/deployment-target",
	}
)

type DockerConfig struct {
	SocketPath string `json:"socketPath,omitempty"`
	Protocol   string `json:"protocol,omitempty"`
	Host       string `json:"host,omitempty"`
	Port       int    `json:"port,omitempty"`
}

type RemoteServices map[string]string

type Definition struct {
	Kind     string                 `json:"kind"`
	Metadata map[string]interface{} `json:"metadata"`
	Spec     map[string]interface{} `json:"spec,omitempty"`
}

type DefinitionInfo struct {
	YmlPath    string
	Path       string
	Version    string
	Definition Definition
	HasWeb     bool
}

type ClusterConfig struct {
	Cluster        map[string]interface{} `json:"cluster,omitempty"`
	Docker         DockerConfig           `json:"docker,omitempty"`
	Environment    string                 `json:"environment,omitempty"`
	RemoteServices RemoteServices         `json:"remoteServices,omitempty"`
}

// NewClusterConfig creates a new instance of ClusterConfig
func NewClusterConfig() *ClusterConfig {
	return &ClusterConfig{}
}

func (c *ClusterConfig) getClusterServicePort() string {
	if envPort := os.Getenv("KAPETA_LOCAL_CLUSTER_PORT"); envPort != "" {
		return envPort
	}

	port := c.GetClusterConfig().Cluster["port"]
	return fmt.Sprintf("%v", port)
}

func (c *ClusterConfig) getClusterServiceHost() string {
	if envHost := os.Getenv("KAPETA_LOCAL_CLUSTER_HOST"); envHost != "" {
		return envHost
	}
	return c.GetClusterConfig().Cluster["host"].(string)
}

func (c *ClusterConfig) getDockerConfig() DockerConfig {
	return c.GetClusterConfig().Docker
}

func (c *ClusterConfig) getEnvironment() string {
	return c.GetClusterConfig().Environment
}

func (c *ClusterConfig) getRemoteServices() RemoteServices {
	remoteServices := c.GetClusterConfig().RemoteServices
	if remoteServices == nil {
		remoteServices = make(RemoteServices)
		c.GetClusterConfig().RemoteServices = remoteServices
	}
	return remoteServices
}

func (c *ClusterConfig) getRemoteService(name string, defaultValue string) string {
	if val, ok := c.getRemoteServices()[name]; ok {
		return val
	}
	return defaultValue
}

func (c *ClusterConfig) getKapetaBasedir() string {
	return getKapetaDir()
}

func (c *ClusterConfig) getAuthenticationPath() string {
	authTokenPath := os.Getenv("KAPETA_CREDENTIALS")
	if authTokenPath == "" {
		authTokenPath = filepath.Join(c.getKapetaBasedir(), "authentication.json")
	}
	return authTokenPath
}

func (c *ClusterConfig) getRepositoryBasedir() string {
	return filepath.Join(c.getKapetaBasedir(), "repository")
}

func (c *ClusterConfig) getRepositoryAssetPath(handle, name, version string) string {
	return filepath.Join(c.getRepositoryBasedir(), handle, name, version)
}

func (c *ClusterConfig) getRepositoryAssetInfoPath(handle, name, version string) (string, string) {
	assetBase := c.getRepositoryAssetPath(handle, name, version)
	assetFile, versionFile := c.getRepositoryAssetInfoRelativePath(assetBase)

	return assetFile, versionFile
}

func (c *ClusterConfig) getRepositoryAssetInfoRelativePath(assetBase string) (string, string) {
	kapetaBase := filepath.Join(assetBase, ".kapeta")

	return filepath.Join(assetBase, "kapeta.yml"), filepath.Join(kapetaBase, "version.yml")
}

func (c *ClusterConfig) GetProviderDefinitions(kindFilter ...string) []DefinitionInfo {
	resolvedFilters := []string{}
	copy(resolvedFilters, kindFilter)
	if len(kindFilter) == 0 {
		resolvedFilters = append(resolvedFilters, PROVIDER_TYPES...)
	}

	return c.getDefinitions(resolvedFilters...)
}

func (c *ClusterConfig) getDefinitions(kindFilter ...string) []DefinitionInfo {
	dir := c.getRepositoryBasedir()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil
	}

	resolvedFilters := []string{}
	copy(resolvedFilters, kindFilter)

	resolvedFilters = toLowerCase(resolvedFilters)

	ymlFiles, err := filepath.Glob(filepath.Join(c.getRepositoryBasedir(), "*/*/*/@(kapeta.yml)"))
	if err != nil {
		return nil
	}

	lists := [][]DefinitionInfo{}
	for _, folder := range ymlFiles {
		ymlPath := filepath.Join(c.getRepositoryBasedir(), folder)
		list := c.parseYmlPath(ymlPath)
		lists = append(lists, list)
	}

	definitions := []DefinitionInfo{}
	for _, list := range lists {
		definitions = append(definitions, list...)
	}

	return filterDefinitions(definitions, resolvedFilters)
}

func (c *ClusterConfig) getClusterConfigFile() string {
	return filepath.Join(c.getKapetaBasedir(), KAPETA_CLUSTER_SERVICE_CONFIG_FILE)
}

func (c *ClusterConfig) GetClusterConfig() *ClusterConfig {
	if os.Getenv("TEST_KAPETA_CLUSTER_CONFIG_FILE") != "" {
		err := yaml.Unmarshal([]byte(os.Getenv("TEST_KAPETA_CLUSTER_CONFIG_FILE")), &c)
		if err != nil {
			fmt.Printf("Error unmarshalling cluster config from test config: %s\n", err)
			return nil
		}
	} else {
		if _, err := os.Stat(c.getClusterConfigFile()); err == nil {
			rawYAML, err := os.ReadFile(c.getClusterConfigFile())
			if err != nil {
				fmt.Printf("Error reading cluster config file: %s\n", err)
				return nil
			}

			err = yaml.Unmarshal(rawYAML, &c)
			if err != nil {
				fmt.Printf("Error unmarshalling cluster config: %s\n", err)
				return nil
			}
		}
	}
	if c == nil {
		c = &ClusterConfig{}
	}

	if c.Cluster == nil {
		c.Cluster = make(map[string]interface{})
	}

	if c.Cluster["port"] == nil {
		c.Cluster["port"] = KAPETA_CLUSTER_SERVICE_DEFAULT_PORT
	}

	if c.Cluster["host"] == nil {
		c.Cluster["host"] = KAPETA_CLUSTER_SERVICE_DEFAULT_HOST
	}

	if c.Docker == (DockerConfig{}) {
		c.Docker = DockerConfig{}
	}

	if c.Environment == "" {
		c.Environment = KAPETA_CLUSTER_SERVICE_DEFAULT_ENV
	}

	fmt.Printf("Read cluster config from file: %s\n", c.getClusterConfigFile())

	return c
}

func (c *ClusterConfig) GetClusterServiceAddress() string {
	clusterPort := c.getClusterServicePort()
	host := c.getClusterServiceHost()
	return fmt.Sprintf("http://%s:%s", host, clusterPort)
}

func getKapetaDir() string {
	kapetaDir := os.Getenv("KAPETA_HOME")
	if kapetaDir == "" {
		kapetaDir = filepath.Join(os.Getenv("HOME"), ".kapeta")
	}
	return kapetaDir
}

func toLowerCase(strList []string) []string {
	for i, str := range strList {
		strList[i] = strings.ToLower(str)
	}
	return strList
}

func (c *ClusterConfig) parseYmlPath(ymlPath string) []DefinitionInfo {

	if _, err := os.Stat(ymlPath); os.IsNotExist(err) {
		return []DefinitionInfo{}
	}

	raw, err := os.ReadFile(ymlPath)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return []DefinitionInfo{}
	}

	version := "local"
	versionInfoFile := filepath.Join(filepath.Dir(ymlPath), ".kapeta", "version.yml")
	if _, err := os.Stat(versionInfoFile); err == nil {
		versionRaw, err := os.ReadFile(versionInfoFile)
		if err != nil {
			fmt.Printf("Error reading version file: %s\n", err)
			return []DefinitionInfo{}
		}

		var versionInfo map[string]interface{}
		err = yaml.Unmarshal(versionRaw, &versionInfo)
		if err != nil {
			fmt.Printf("Error unmarshalling version YAML: %s\n", err)
			return []DefinitionInfo{}
		}

		if versionInfo["version"] != nil {
			version = versionInfo["version"].(string)
		}
	}

	var result []DefinitionInfo
	docs := strings.Split(string(raw), "---")
	for _, doc := range docs {
		var data Definition
		err := yaml.Unmarshal([]byte(doc), &data)
		if err != nil {
			fmt.Printf("Error unmarshalling YAML document: %s\n", err)
			continue
		}

		result = append(result, DefinitionInfo{
			YmlPath:    ymlPath,
			Path:       filepath.Dir(ymlPath),
			Version:    version,
			Definition: data,
			HasWeb:     exists(filepath.Join(filepath.Dir(ymlPath), "web")),
		})
	}

	return result
}

func filterDefinitions(definitions []DefinitionInfo, filters []string) []DefinitionInfo {
	result := []DefinitionInfo{}
	for _, def := range definitions {
		if len(filters) > 0 {
			if def.Definition.Kind == "" || !contains(filters, strings.ToLower(def.Definition.Kind)) {
				continue
			}
		}
		result = append(result, def)
	}
	return result
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
