package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"google-search/props"
	"google-search/server"
	"google-search/service"
	"google-search/service/providers/googlesearch"
)

var (
	properties *viper.Viper
	prop       *props.Properties
)

func main() {
	initialiseApplication()
	log.Info("✅ Initialisation of application Completed")

	startApplication()
}

func startApplication() {
	startServer()
}

func startServer() {
	providers := getProviders(properties)
	log.Info("✅ All providers initialised")
	handler := service.NewSearchCompositionHandler(providers)
	servers := server.NewServer(prop, properties)
	servers.ConfigureAPI(handler)
}

func getProviders(properties *viper.Viper) service.Providers {
	googleSearchProvider := googlesearch.NewGoogleClient(properties)

	return service.Providers{
		GoogleSearchClient: googleSearchProvider,
	}
}
