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
go get github.com/maron09/logmiddleware

