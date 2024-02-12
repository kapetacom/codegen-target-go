package config

import (
	"fmt"
	kapeta "github.com/kapetacom/sdk-go-config"
)

// Configuration for e-mails
type EmailConfigConfig struct {
	// Sender for e-mails. Note that you should be allowed to send e-mails from this domain and user
	From string `json:"from" xml:"from" yaml:"from"`
}

func GetEmailConfigConfigWithDefault(defaultValue EmailConfigConfig) EmailConfigConfig {
	anyconfig := kapeta.CONFIG.GetOrDefault("EmailConfig", defaultValue)
	result := EmailConfigConfig{}
	err := kapeta.Transcode(anyconfig, &result)
	if err != nil {
		panic(fmt.Errorf("failed to transcode config: %w", err))
	}

	return result
}

func GetEmailConfigConfig() EmailConfigConfig {
	anyconfig := kapeta.CONFIG.Get("EmailConfig")
	result := EmailConfigConfig{}
	err := kapeta.Transcode(anyconfig, &result)
	if err != nil {
		panic(fmt.Errorf("failed to transcode config: %w", err))
	}

	return result
}
