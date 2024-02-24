package main

/* Sample file to load the configurations from .env file
Usage : to read the values from
Azure Key Vaults,Azure App configuration
*/

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	EnvVarGoogleSearchURL      = "GOOGLE_SEARCH_URL"
	EnvVarGoogleAPIKey         = "GOOGLE_API_KEY"
	EnvVarCustomSearchEngineId = "GOOGLE_CUSTOM_SEARCH_ID"
)

func initialiseApplication() {
	initializeConfiguration("./env/")
}

func initializeConfiguration(path string) {
	viperConfigManager := viper.NewWithOptions(viper.KeyDelimiter("_"))
	viperConfigManager.SetConfigName("application")
	viperConfigManager.SetConfigType("yaml")
	viperConfigManager.AddConfigPath("/etc/config/")
	viperConfigManager.AddConfigPath(path)
	err := viperConfigManager.BindEnv(EnvVarGoogleSearchURL, EnvVarGoogleAPIKey, EnvVarCustomSearchEngineId)
	if err != nil {
		log.Warnf("❌ Failed to bind a configuration key to the '%v, %v, %v' environment variable with error %v",
			EnvVarGoogleSearchURL, EnvVarGoogleAPIKey, EnvVarCustomSearchEngineId, err)
	}

	viperConfigManager.AutomaticEnv()
	viperConfigManager.AllowEmptyEnv(true)
	viperConfigManager.WatchConfig()
	viperConfigManager.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("⬆️ Config file changed: %s", e.Name)
	})

	err = viperConfigManager.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("❌ unable to start google-search-results due to missing application config %v", err))
	}

	log.Infof("✅ Loading application config from %v", viperConfigManager.ConfigFileUsed())
}
