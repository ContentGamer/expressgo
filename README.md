# ExpressGO

**ExpressGO** is an implementation of the popular ExpressJS framework for building Websites, API's etc.., ExpressGO uses the standard 'http' system.

## Features

- Supports JSON.
- Supports Parameters (/books/:bookID/getData).
- Supports Query Parameters.
- Simple And Clean code.

## Code Examples

### Simple Homepage
```go
func main() {
  app := expressgo.Express()

  app.Get("/", func(req *expressgo.Request, res *expressgo.Response) {
    res.Json(expressgo.JSONData{
      "message": "Hello world!",
    })
  })

  app.Listen(3030, func(port int) {
    fmt.Printf("Server running at http://localhost:%d\n", port)
  })
}
```

### Simple Book System
```go
type Book struct {
	Name   string
	Author string
	Year   string

	Title string
}

var books map[int]Book = map[int]Book{
	0: {
		Name:   "1984",
		Author: "George Orwell",
		Year:   "1949",
		Title:  "Dystopian Future",
	},

	1: {
		Name:   "To Kill a Mockingbird",
		Author: "Harper Lee",
		Year:   "1960",
		Title:  "Racial Injustice",
	},

	2: {
		Name:   "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Year:   "1925",
		Title:  "American Dream",
	},

	3: {
		Name:   "Moby Dick",
		Author: "Herman Melville",
		Year:   "1851",
		Title:  "Adventure and Obsession",
	},

	4: {
		Name:   "Pride and Prejudice",
		Author: "Jane Austen",
		Year:   "1813",
		Title:  "Romance and Society",
	},
}

func structToMap(obj interface{}) map[string]interface{} {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil
	}

	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil
	}

	return result
}

func main() {
	app := expressgo.Express()

	app.Get("/books/:id/get", func(req *expressgo.Request, res *expressgo.Response) {
		id := req.Params["id"]
		num, err := strconv.Atoi(id)

		if err != nil {
			res.Status(400).Json(expressgo.JSONData{
				"error":   true,
				"message": "Expected the ID to be a number.",
			})
			return
		}

		book, exists := books[num]
		if !exists {
			res.Status(400).Json(expressgo.JSONData{
				"error":   true,
				"message": fmt.Sprintf("Book with the id of '%s' does not exist.", id),
			})
			return
		}

		res.Status(200).Json(structToMap(book))
	})

	app.Listen(3030, func(port int) {
		fmt.Printf("Server running at http://localhost:%d\n", port)
	})
}
```