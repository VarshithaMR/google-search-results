package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"google-search/props"
	"google-search/server/api"
	"google-search/service"
)

type Server struct {
	api             api.DomainHandler
	host            string
	port            int
	contextRoot     string
	defaultQuantity int
	doOnce          sync.Once
}

func NewServer(properties *props.Properties, viper *viper.Viper) *Server {
	if err := viper.Unmarshal(&properties, func(c *mapstructure.DecoderConfig) {
		c.DecodeHook = mapstructure.StringToTimeDurationHookFunc()
	}); err != nil {
		log.Warnf("❌ Unable to read application.yaml file : %s", err)
	}
	server := new(Server)
	server.host = properties.Server.Host
	server.port = properties.Server.Port
	server.contextRoot = properties.Server.ContextRoot
	server.defaultQuantity = properties.Server.ResultQuantity
	server.api = api.NewServiceEngine()
	return server
}

// ConfigureAPI configures the API with all the endpoints with respective handlers
func (s *Server) ConfigureAPI(myService service.SearchCompositionHandler) {
	s.doOnce.Do(func() {
		configureApi(s.api, s.contextRoot, myService, s.defaultQuantity)
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: s.api,
	}

	go func() {
		// service connections
		err := http.ListenAndServe(srv.Addr, srv.Handler)

		if err != nil {
			log.Fatalf("❌ Failure to start Go http server: %v", err)
			return
		}
	}()

	s.api.Logger(context.Background(), "✅ Started server on : %s", fmt.Sprintf("http://%s:%d", s.host, s.port))

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.api.Logger(context.Background(), "Server is shutting down ... %s", "domain-handler")

	s.Shutdown()

	s.api.Logger(context.Background(), "Until next time ... %s")
	//return nil
}

// Shutdown server and clean up resources
func (s *Server) Shutdown() {
	s.api.ServerShutdown()
}
