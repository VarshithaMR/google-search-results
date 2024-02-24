package cmd

import (
	"github.com/spf13/viper"

	"google-search/service"
	"google-search/service/providers/googlesearch"
)

var properties *viper.Viper

func main() {
	initialiseApplication()

	startApplication()
}

func startApplication() {
	startServer()
}

func startServer() {
	providers := getProviders(properties)
}

func getProviders(properties *viper.Viper) *service.Providers {
	googleSearchProvider := googlesearch.NewGoogleClient(properties)

	return &service.Providers{
		GoogleSearchClient: googleSearchProvider,
	}
}
