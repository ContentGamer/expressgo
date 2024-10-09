package expressgo

import (
	"fmt"
	"strings"
)

type ExpressRouter struct {
	methods map[string]URLDefinition
	baseURL string

	// Create a new GET Method Route.
	Get ExpressMethod
	// Create a new POST Method Route.
	Post ExpressMethod
	// Create a new PUT Method Route.
	Put ExpressMethod
	// Create a new HEAD Method Route.
	Head ExpressMethod
	// Create a new PATCH Method Route.
	Patch ExpressMethod
	// Create a new OPTIONS Method Route.
	Options ExpressMethod
	// Create a new CONNECT Method Route.
	Connect ExpressMethod
	// Create a new DELETE Method Route.
	Delete ExpressMethod
}

func (router *ExpressRouter) pushToStack(url string, handler URLHandler, method Method) {
	router.methods[url] = URLDefinition{method, handler}
}

func (router *ExpressRouter) setupRouter(baseURL string) {
	router.baseURL = baseURL

	for n, v := range router.methods {
		if strings.HasPrefix(n, baseURL) {
			continue
		}

		fullPath := fmt.Sprintf("%s%s", baseURL, n)
		router.methods[fullPath] = v

		delete(router.methods, n)
	}
}

// Create a new router
//
// Same as 'express.Router()'.
func Router() *ExpressRouter {
	router := ExpressRouter{
		methods: make(map[string]URLDefinition),
		baseURL: "",
	}

	router.Get = func(url string, handler URLHandler) { router.pushToStack(url, handler, GET) }
	router.Post = func(url string, handler URLHandler) { router.pushToStack(url, handler, POST) }
	router.Put = func(url string, handler URLHandler) { router.pushToStack(url, handler, PUT) }
	router.Head = func(url string, handler URLHandler) { router.pushToStack(url, handler, HEAD) }
	router.Patch = func(url string, handler URLHandler) { router.pushToStack(url, handler, PATCH) }
	router.Options = func(url string, handler URLHandler) { router.pushToStack(url, handler, OPTIONS) }
	router.Connect = func(url string, handler URLHandler) { router.pushToStack(url, handler, CONNECT) }
	router.Delete = func(url string, handler URLHandler) { router.pushToStack(url, handler, DELETE) }

	return &router
}
