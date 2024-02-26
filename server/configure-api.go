package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"google-search/server/api"
	"google-search/service"
)

const (
	endpoint = "/top-list"
)

func configureApi(handler api.DomainHandler, contextRoot string, service service.SearchCompositionHandler, defaultQuantity int) {

	// Setup global middleware like Logging, Panic handler
	handler = handler.WithGlobalMiddleware(dontPanic)

	getGoogleSearchResultsOperation := &api.Operations{
		Method: api.GET,
		Path:   endpoint,
		Handler: func(ctx *gin.Context) {
			service.GetSearchResults(ctx.Writer, ctx.Request, defaultQuantity)
		},
		Middleware: []gin.HandlerFunc{},
	}

	handler = handler.WithOperations(contextRoot, nil, getGoogleSearchResultsOperation)

	// Setup shutdown
	handler = handler.WithShutdown(func() {
		handler.Logger(context.Background(), "🔻 Shutting down %s", "domain-handler server")
	})
}

func dontPanic(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Writer.WriteHeader(http.StatusInternalServerError)
			ctx.Writer.Write([]byte("😰PANIC!!!😰"))
		}
	}()
	ctx.Next()
}
