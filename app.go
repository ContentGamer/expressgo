package expressgo

import (
	"fmt"
	"net/http"
	"strings"
)

type ListenHandler func(port int)
type URLHandler func(req *Request, res *Response)

type URLDefinition struct {
	Method  Method
	Handler URLHandler
}
type Method string

const (
	GET     Method = "get"
	POST    Method = "post"
	PUT     Method = "put"
	HEAD    Method = "head"
	PATCH   Method = "patch"
	OPTIONS Method = "options"
	CONNECT Method = "connect"
	DELETE  Method = "delete"
)

type ExpressMethod func(url string, handler URLHandler)
type RouteMethod func(url string, route *ExpressApp)
type ExpressApp struct {
	methods map[string]URLDefinition

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

func call(app *ExpressApp, name string, params map[string]string, w http.ResponseWriter, r *http.Request) {
	res := Response{
		writer:  w,
		headers: make(map[string]string),
		body:    strings.Builder{},
		status:  200,
	}
	req := Request{data: r, resp: &res}

	req.setupVariables()
	req.Params = params
	app.methods[name].Handler(&req, &res)

	for n, v := range res.headers {
		res.writer.Header().Set(n, v)
	}
	res.writer.WriteHeader(res.status)
	w.Write([]byte(res.body.String()))
}

// Add a router to this app.
func (app *ExpressApp) Route(url string, router *ExpressRouter) {
	router.setupRouter(url)
	for n, v := range router.methods {
		app.baseSetup(n, v.Handler, v.Method)
	}
}

func (app *ExpressApp) baseSetup(url string, handler URLHandler, method Method) {
	app.methods[url] = URLDefinition{method, handler}
	if url == "/" {
		return
	}

	http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if strings.ToLower(r.Method) != string(method) {
			http.NotFound(w, r)
			return
		}

		res := Response{
			writer:  w,
			headers: make(map[string]string),
			body:    strings.Builder{},
			status:  200,
		}
		req := Request{data: r, resp: &res}

		req.setupVariables()
		handler(&req, &res)

		// Sorting order:
		// -> headers
		// -> status
		// -> body

		for n, v := range res.headers {
			res.writer.Header().Set(n, v)
		}
		res.writer.WriteHeader(res.status)
		w.Write([]byte(res.body.String()))
	})
}

// Listen to a port.
func (app *ExpressApp) Listen(port int, handler ListenHandler) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if _, exists := app.methods["/"]; r.URL.Path == "/" && exists {
			call(app, "/", make(map[string]string), w, r)
			return
		}

		pths := map[string](map[string]string){}
		for n := range app.methods {
			pths[n] = parsePathParams(n, r.URL.Path)
		}

		var isEmpty bool = true
		for _, v := range pths {
			if len(v) == 0 {
				continue
			}

			isEmpty = false
			break
		}

		if isEmpty {
			http.NotFound(w, r)
			return
		}

		for n, v := range pths {
			if len(v) == 0 {
				continue
			}
			if strings.ToLower(r.Method) != string(app.methods[n].Method) {
				http.NotFound(w, r)
				return
			}

			call(app, n, v, w, r)
		}
	})

	go func() {
		http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
	}()

	if handler != nil {
		handler(port)
	}

	select {}
}
