# logmiddleware

`logmiddleware` is a simple HTTP middleware for logging HTTP requests. It logs the HTTP method, path, status code, and the duration of each request to the standard logger. This package is easy to integrate into any Go project that uses the `net/http` package.

## Features

- Logs the request method (e.g., GET, POST)
- Logs the request path (e.g., /api/v1/products)
- Logs the HTTP status code (e.g., 200, 404)
- Logs the duration of each request
- Lightweight and simple to use

## Installation

To install the package, run the following command:

```bash
go get github.com/maron09/logmiddleware
```

## Usage
You can use `logmiddleware` by wrapping your HTTP handlers or routers with the `LoggingMiddleware` function.

```go
package main

import (
    "net/http"
    "github.com/maron09/logmiddleware"
)

func main() {
    mux := http.NewServeMux()

    // Register your routes
    mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    // Wrap the mux router with the logging middleware
    loggedMux := logmiddleware.LoggingMiddleware(mux)

    // Start the server
    http.ListenAndServe(":8080", loggedMux)
}

```
## Example with `gorilla/mux`
You can also use `logmiddleware` with `gorilla/mux` or any other router:
```go
package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/maron09/logmiddleware"
)

func main() {
    router := mux.NewRouter()

    // Define your routes
    router.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Product List"))
    }).Methods("GET")

    // Apply logging middleware
    loggedRouter := logmiddleware.LoggingMiddleware(router)

    // Start the server
    http.ListenAndServe(":8080", loggedRouter)
}


```
## Customizing the Logging Behavior

This middleware uses Go's built-in `log` package to print logs to the standard output (typically the console). If you'd like to modify this behavior (e.g., log to a file or use a structured logger), you can modify the `loggingMiddleware` function to suit your needs.

For instance, integrating with a structured logger like `logrus` or `zap` can be done by replacing the `log.Printf` line with the appropriate logger function.

## Contributing

Feel free to contribute to this project by opening issues or submitting pull requests. Your contributions are highly appreciated!
## License

### Key Sections:
1. **Features**: Lists the key capabilities of the package.
2. **Installation**: Provides installation instructions using `go get`.
3. **Usage**: Demonstrates how to integrate the middleware with:
   - The basic `http.ServeMux`.
   - `gorilla/mux`.
4. **Customizing the Logging Behavior**: Provides flexibility for users to adapt the middleware with other loggers.
5. **Contributing**: Encourages contributions and involvement from other developers.
6. **License**: States that the project is licensed under the MIT License.

This **`README.md`** should provide everything users need to understand, install, and use the package effectively in their Go projects. Let me know if you want any more customizations or further examples!


[MIT](https://choosealicense.com/licenses/mit/)
