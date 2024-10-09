// --------------------------------
// 	ExpressJS implementation in Go
// --------------------------------

package expressgo

import (
	"encoding/json"
	"fmt"
)

// Decode a byte array to JSONData (expanded to map[string]interface{}).
func DecodeJSON(data []byte) JSONData {
	var result JSONData = make(JSONData)

	err := json.Unmarshal(data, &result)
	if err != nil {
		panic(fmt.Sprintln("Error decoding JSON:", err))
	}

	return result
}

// Initiate a new express app
//
// Same as 'express()' function in expressjs.
func Express() *ExpressApp {
	app := &ExpressApp{
		methods: make(map[string]URLDefinition),
	}

	app.Get = func(url string, handler URLHandler) { app.baseSetup(url, handler, GET) }
	app.Post = func(url string, handler URLHandler) { app.baseSetup(url, handler, POST) }
	app.Put = func(url string, handler URLHandler) { app.baseSetup(url, handler, PUT) }
	app.Head = func(url string, handler URLHandler) { app.baseSetup(url, handler, HEAD) }
	app.Patch = func(url string, handler URLHandler) { app.baseSetup(url, handler, PATCH) }
	app.Options = func(url string, handler URLHandler) { app.baseSetup(url, handler, OPTIONS) }
	app.Connect = func(url string, handler URLHandler) { app.baseSetup(url, handler, CONNECT) }
	app.Delete = func(url string, handler URLHandler) { app.baseSetup(url, handler, DELETE) }

	return app
}
