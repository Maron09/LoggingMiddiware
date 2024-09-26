# Go HTTP Logging Middleware

This package provides a simple logging middleware for Go web applications. It logs every HTTP request made to your server, including the HTTP method, request path, status code, and the time taken to process each request. It is lightweight and easily integrates with the standard `net/http` package as well as routers like `gorilla/mux`.

## Features

- Logs the HTTP method (GET, POST, etc.)
- Logs the request path (e.g., `/api/v1/products`)
- Logs the status code (e.g., 200, 404)
- Logs the time taken to handle each request
- Easy to integrate with existing Go web applications
- Works with both the standard `net/http` router and other popular routers like `gorilla/mux`

## Installation

To install this package, run the following command:

```bash
go get github.com/yourusername/logmiddleware

Usage

You can use this middleware by wrapping your existing http.Handler with the LoggingMiddleware function.

Basic Example

Below is an example of how to integrate the logging middleware into a simple Go application using the standard http.ServeMux router:

package main

import (
    "net/http"
    "github.com/yourusername/logmiddleware"
)

func main() {
    // Create a new ServeMux
    mux := http.NewServeMux()

    // Define a simple handler
    mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    // Wrap the ServeMux with the logging middleware
    loggedMux := logmiddleware.LoggingMiddleware(mux)

    // Start the HTTP server
    http.ListenAndServe(":8080", loggedMux)
}

Using with gorilla/mux

You can also use this logging middleware with the gorilla/mux router:

package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/yourusername/logmiddleware"
)

func main() {
    // Create a new gorilla/mux router
    router := mux.NewRouter()

    // Define your routes
    router.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Product List"))
    }).Methods("GET")

    // Wrap the router with the logging middleware
    loggedRouter := logmiddleware.LoggingMiddleware(router)

    // Start the server
    http.ListenAndServe(":8080", loggedRouter)
}

Customizing the Logging Behavior

This middleware uses Go's built-in log package to print logs to the standard output (typically the console). If you'd like to modify this behavior (e.g., log to a file or use a structured logger), you can modify the loggingMiddleware function to suit your needs.

For instance, integrating with a structured logger like logrus or zap can be done by replacing the log.Printf line with the appropriate logger function.

Contributing

Feel free to contribute to this project by opening issues or submitting pull requests. Your contributions are highly appreciated!