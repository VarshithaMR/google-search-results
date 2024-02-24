package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Method int

const (
	GET Method = iota
	POST
	PUT
	DELETE
	PATCH
)

// Operations contains the parameters of every endpoint exposed by the service
type Operations struct {
	Handler    gin.HandlerFunc
	Method     Method
	Middleware []gin.HandlerFunc
	Path       string
}

type serviceEngine struct {
	*gin.Engine
	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	serverShutdown func()

	// User defined logger function.
	logger func(context.Context, string, ...interface{})
}

// handler interface is wrapper on the http handler
type DomainHandler interface {
	http.Handler
	WithShutdown(func()) DomainHandler
	WithOperations(rootPath string, authorize gin.HandlerFunc, operations ...*Operations) DomainHandler
	WithGlobalMiddleware(globalMiddleware ...gin.HandlerFunc) DomainHandler
	Logger(context.Context, string, ...interface{})
}

func (d *serviceEngine) WithShutdown(shutdown func()) DomainHandler {
	d.serverShutdown = shutdown
	return d
}

func (d *serviceEngine) WithOperations(rootPath string, authorize gin.HandlerFunc, operations ...*Operations) DomainHandler {
	group := d.Group(rootPath)
	if authorize != nil {
		group.Use(authorize)
	}
	for _, operation := range operations {
		switch operation.Method {
		case GET:
			createGroupWithMiddleware(group, operation.Middleware...).GET(operation.Path, operation.Handler)
		}
	}
	return d
}

func (d *serviceEngine) WithGlobalMiddleware(globalMiddleware ...gin.HandlerFunc) DomainHandler {
	d.Use(globalMiddleware...)
	return d
}

func (d *serviceEngine) Logger(ctx context.Context, message string, arguments ...interface{}) {
	d.logger(ctx, message, arguments)
}

func createGroupWithMiddleware(group *gin.RouterGroup, globalMiddleware ...gin.HandlerFunc) *gin.RouterGroup {
	if len(globalMiddleware) == 0 {
		return group
	}
	return group.Group("", globalMiddleware...)
}

func NewServiceEngine() DomainHandler {
	serviceDomain := &serviceEngine{
		Engine: gin.New(),
		serverShutdown: func() {
			log.Println("🔻Server Shutdown ...")
		},
		logger: func(x context.Context, s string, arg ...interface{}) {
		},
	}
	return serviceDomain
}
