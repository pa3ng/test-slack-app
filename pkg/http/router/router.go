package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pa3ng/test-slack-app/pkg/http/middleware"
	"github.com/pa3ng/test-slack-app/pkg/http/server"
)

// Route represents the route of an HTTP request
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

var health = Route{
	"Health",
	"GET",
	"/health",
	server.Health,
}

var commands = Route{
	"Commands",
	"POST",
	"/slack_events/v1/ymir_dev-v1_commands/",
	server.Commands,
}

// NewRouter creates a custom HTTP router with added middleware
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET(health.Path, health.HandlerFunc)
	router.POST(commands.Path, handler(commands))

	return router
}

func handler(route Route) httprouter.Handle {
	var handle httprouter.Handle
	handle = route.HandlerFunc
	handle = middleware.Log(handle, route.Name)
	handle = middleware.Auth(handle)
	return handle
}
