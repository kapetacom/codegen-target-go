package config

import (
	_ "embed"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClusterConfig(t *testing.T) {
	c := NewClusterConfig()
	assert.NotNil(t, c)
}

func TestGetClusterServicePort(t *testing.T) {
	os.Setenv("KAPETA_LOCAL_CLUSTER_PORT", "8080")
	c := NewClusterConfig()
	assert.Equal(t, "8080", c.getClusterServicePort())
	os.Unsetenv("KAPETA_LOCAL_CLUSTER_PORT")
}

func TestGetClusterServiceHost(t *testing.T) {
	os.Setenv("KAPETA_LOCAL_CLUSTER_HOST", "10.0.0.1")
	c := NewClusterConfig()
	assert.Equal(t, "10.0.0.1", c.getClusterServiceHost())
	os.Unsetenv("KAPETA_LOCAL_CLUSTER_HOST")
}

func TestGetDockerConfig(t *testing.T) {
	c := NewClusterConfig()
	assert.Equal(t, DockerConfig{}, c.getDockerConfig())
}

func TestGetEnvironment(t *testing.T) {
	os.Setenv("KAPETA_ENVIRONMENT", "production")
	c := NewClusterConfig()
	assert.Equal(t, "production", c.getEnvironment())
	os.Unsetenv("KAPETA_ENVIRONMENT")
}

func TestGetRemoteServices(t *testing.T) {
	c := NewClusterConfig()
	assert.Equal(t, make(RemoteServices), c.getRemoteServices())
}

func TestGetRemoteService(t *testing.T) {
	os.Setenv("KAPETA_REMOTE_SERVICE_foo", "bar")
	c := NewClusterConfig()
	assert.Equal(t, "bar", c.getRemoteService("foo", "bar"))
}

func TestGetKapetaBasedir(t *testing.T) {
	c := NewClusterConfig()
	home := os.Getenv("HOME")
	assert.Equal(t, home+"/.kapeta", c.getKapetaBasedir())
}

func TestGetAuthenticationPath(t *testing.T) {
	c := NewClusterConfig()
	home := os.Getenv("HOME")
	assert.Equal(t, home+"/.kapeta/authentication.json", c.getAuthenticationPath())
}

func TestGetRepositoryBasedir(t *testing.T) {
	c := NewClusterConfig()
	home := os.Getenv("HOME")
	assert.Equal(t, home+"/.kapeta/repository", c.getRepositoryBasedir())
}

func TestGetRepositoryAssetPath(t *testing.T) {
	c := NewClusterConfig()
	assert.Equal(t, os.Getenv("HOME")+"/.kapeta/repository/foo/bar/1.0.0", c.getRepositoryAssetPath("foo", "bar", "1.0.0"))
}

func TestGetRepositoryAssetInfoPath(t *testing.T) {
	c := NewClusterConfig()
	assetFile, versionFile := c.getRepositoryAssetInfoPath("foo", "bar", "1.0.0")
	assert.Equal(t, os.Getenv("HOME")+"/.kapeta/repository/foo/bar/1.0.0/.kapeta/version.yml", versionFile)
	assert.Equal(t, os.Getenv("HOME")+"/.kapeta/repository/foo/bar/1.0.0/kapeta.yml", assetFile)
}

func setKAPETA_HOME() {
	wd, _ := os.Getwd()
	// remove the config path from the wd
	wd = wd[:len(wd)-len("/config")]
	os.Setenv("KAPETA_HOME", wd+"/.kapeta")
}

func TestGetClusterConfigFile(t *testing.T) {
	setKAPETA_HOME()
	defer os.Unsetenv("KAPETA_HOME")
	c := NewClusterConfig()
	assert.Equal(t, os.Getenv("KAPETA_HOME")+"/cluster-service.yml", c.getClusterConfigFile())
}

func TestGetClusterConfig(t *testing.T) {
	c := NewClusterConfig()
	assert.NotNil(t, c.GetClusterConfig())
	assert.Equal(t, "production", c.GetClusterConfig().Environment)
	assert.Equal(t, "127.0.0.1", c.GetClusterConfig().Cluster["host"])
}

func TestGetClusterServiceAddress(t *testing.T) {
	c := NewClusterConfig()
	assert.Equal(t, "http://127.0.0.1:35100", c.GetClusterServiceAddress())
}
