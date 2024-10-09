# ExpressGO

**ExpressGO** is an implementation of the popular ExpressJS framework for building Websites, API's etc.., ExpressGO uses the standard 'http' system.

## Features

- Supports JSON.
- Supports Parameters (/books/:bookID/getData).
- Supports Query Parameters.
- Simple And Clean code.

## Code Examples

```go
func main() {
  app := expressgo.Express()

  app.Get("/", func(req *expressgo.Request, res *expressgo.Response) {
    res.Json(expressgo.JSONData{
      "message": "Hello world!"
    })
  })

  app.Listen(3030, func(port int) {
    fmt.Printf("Server running at http://localhost:%d\n", port)
  })
}

```
