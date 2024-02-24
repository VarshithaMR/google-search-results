package cmd

import (
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

	startApplication()
}

func startApplication() {
	startServer()
}

func startServer() {
	providers := getProviders(properties)
	handler := service.NewSearchCompositionHandler(providers)
	servers := server.NewServer(prop)
	servers.ConfigureAPI(handler)
}

func getProviders(properties *viper.Viper) service.Providers {
	googleSearchProvider := googlesearch.NewGoogleClient(properties)

	return service.Providers{
		GoogleSearchClient: googleSearchProvider,
	}
}
