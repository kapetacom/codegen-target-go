// Copyright 2023 Kapeta Inc.
// SPDX-License-Identifier: MIT

package providers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func hostAndFromURL(url string) (string, string) {
	hostAndPort := url[7:]
	return strings.Split(hostAndPort, ":")[0], strings.Split(hostAndPort, ":")[1]
}
func TestLocal(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if strings.HasSuffix(r.URL.Path, "/config/identity") {
			_, _ = w.Write([]byte("{\"systemId\": \"system-id\", \"instanceId\": \"instance-id\"}"))
			return
		}
		if strings.HasSuffix(r.URL.Path, "config/instance") {
			_, _ = w.Write([]byte("{\"id\": \"instance-id\"}"))
			return
		}
		if strings.HasSuffix(r.URL.Path, "provides/rest") {
			_, _ = w.Write([]byte("8080"))
			return
		} else if strings.HasSuffix(r.URL.Path, "provides/grpc") {
			_, _ = w.Write([]byte("8081"))
			return
		}
		if strings.HasSuffix(r.URL.Path, "consumes/foo/rest") {
			_, _ = w.Write([]byte("10.0.0.1:8080"))
			return
		}
		if strings.HasSuffix(r.URL.Path, "consumes/bar/grpc") {
			_, _ = w.Write([]byte("10.0.0.2:8081"))
			return
		}
		if strings.HasSuffix(r.URL.Path, "consumes/baz/rest") {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = w.Write([]byte("40004"))
	}))
	defer srv.Close()

	host, port := hostAndFromURL(srv.URL)
	os.Setenv("KAPETA_LOCAL_CLUSTER_HOST", host)
	os.Setenv("KAPETA_LOCAL_CLUSTER_PORT", port)
	defer os.Unsetenv("KAPETA_LOCAL_CLUSTER_HOST")
	defer os.Unsetenv("KAPETA_LOCAL_CLUSTER_PORT")

	// Mock environment variables for testing
	provider := CreateLocalConfigProvider("kapeta/block-type-gateway-http", "systemID", "instanceID", map[string]interface{}{})

	serverPort, err := provider.GetServerPort("http")
	assert.NoError(t, err)
	assert.Equal(t, "40004", serverPort)

	assert.Equal(t, "system-id", provider.GetSystemId())
	assert.Equal(t, "instance-id", provider.GetInstanceId())
	assert.Equal(t, "kapeta/block-type-gateway-http", provider.GetBlockReference())
	assert.Equal(t, "local", provider.GetProviderId())

	t.Run("verify that we can get port and host configurations", func(t *testing.T) {
		port, err := provider.GetServerPort("rest")
		assert.NoError(t, err)
		assert.Equal(t, "8080", port)

		port, err = provider.GetServerPort("grpc")
		assert.NoError(t, err)
		assert.Equal(t, "8081", port)

		testhost, err := provider.GetServerHost()
		assert.NoError(t, err)
		assert.Equal(t, "127.0.0.1", testhost)
	})
	t.Run("verify that we can get service address", func(t *testing.T) {

		address, err := provider.GetServiceAddress("foo", "rest")
		assert.NoError(t, err)
		assert.Equal(t, "10.0.0.1:8080", address)

		address, err = provider.GetServiceAddress("bar", "grpc")
		assert.NoError(t, err)
		assert.Equal(t, "10.0.0.2:8081", address)

		_, err = provider.GetServiceAddress("baz", "rest")
		assert.Error(t, err)
		assert.Equal(t, "failed to send GET request: request failed - Status: 500", err.Error())
	})
}

func TestGetInstanceHost1(t *testing.T) {
	// create test server that return the correct values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.HasSuffix(r.URL.Path, "/config/identity") {
			_, _ = w.Write([]byte("{\"systemId\": \"system-id\", \"instanceId\": \"instance-id\"}"))
			return
		}
		if strings.HasSuffix(r.URL.Path, "config/instance") {
			_, _ = w.Write([]byte("{\"id\": \"instance-id\"}"))
			return
		}

		if strings.Contains(r.URL.Path, "system-id/unknown-instance-id") {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if strings.Contains(r.URL.Path, "/instances/system-id/") {
			_, _ = w.Write([]byte("10.0.0.1"))
			return
		}
	}))
	defer srv.Close()

	host, port := hostAndFromURL(srv.URL)
	os.Setenv("KAPETA_LOCAL_CLUSTER_HOST", host)
	os.Setenv("KAPETA_LOCAL_CLUSTER_PORT", port)

	defer os.Unsetenv("KAPETA_LOCAL_CLUSTER_HOST")
	defer os.Unsetenv("KAPETA_LOCAL_CLUSTER_PORT")

	os.Setenv("KAPETA_INSTANCE_CONFIG", "{\"foo\": \"bar\"}")
	defer os.Unsetenv("KAPETA_INSTANCE_CONFIG")
	provider := CreateLocalConfigProvider("block-ref", "system-id", "instance-id", map[string]interface{}{
		"type": "local",
	})

	host, err := provider.GetInstanceHost("instance-id")
	assert.NoError(t, err)
	assert.Equal(t, "10.0.0.1", host)

	_, err = provider.GetInstanceHost("unknown-instance-id")
	assert.Error(t, err)
	assert.Equal(t, "failed to send GET request: request failed - Status: 500", err.Error())
}
