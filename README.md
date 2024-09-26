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

##  Usage

You can use logmiddleware by wrapping your HTTP handlers or routers with the LoggingMiddleware function.
Basic Example

go

package main

import (
    "net/http"
    "github.com/yourusername/logmiddleware"
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

Example with gorilla/mux

You can also use logmiddleware with gorilla/mux or any other router:

go

package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/yourusername/logmiddleware"
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

Customizing Log Output

By default, logmiddleware uses Go's standard log package to log to the console. You can easily modify the loggingMiddleware function to log to a file, structured logger, or any other logging destination.
