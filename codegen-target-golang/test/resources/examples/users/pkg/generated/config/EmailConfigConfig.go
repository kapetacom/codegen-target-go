package config

import (
	"fmt"
	sdkgoconfig "github.com/kapetacom/golang-language-target/sdk-go-config"
)

// Configuration for e-mails
type EmailConfigConfig struct {
	// Sender for e-mails. Note that you should be allowed to send e-mails from this domain and user
	From string `json:"from" xml:"from" yaml:"from"`
}

func GetEmailConfigConfigWithDefault(defaultValue EmailConfigConfig) EmailConfigConfig {
	anyconfig := sdkgoconfig.CONFIG.GetOrDefault("config", defaultValue)
	result := EmailConfigConfig{}
	err := sdkgoconfig.Transcode(anyconfig, &result)
	if err != nil {
		panic(fmt.Errorf("failed to transcode config: %w", err))
	}

	return result
}

func GetEmailConfigConfig() EmailConfigConfig {
	anyconfig := sdkgoconfig.CONFIG.Get("config")
	result := EmailConfigConfig{}
	err := sdkgoconfig.Transcode(anyconfig, &result)
	if err != nil {
		panic(fmt.Errorf("failed to transcode config: %w", err))
	}

	return result
}
