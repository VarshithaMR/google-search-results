package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"

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

func NewServer(properties *props.Properties) *Server {
	server := new(Server)
	server.host = properties.Server.Host
	server.port = properties.Server.Port
	server.contextRoot = properties.Server.ContextRoot
	server.defaultQuantity = properties.Server.NoOfResults
	server.api = api.NewServiceEngine()
	return server
}

// ConfigureAPI configures the API with all the endpoints with respective handlers
func (s *Server) ConfigureAPI(myService service.SearchCompositionHandler) {
	s.doOnce.Do(func() {
		configureApi(s.api, s.contextRoot, myService, s.port, s.defaultQuantity)
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
}
