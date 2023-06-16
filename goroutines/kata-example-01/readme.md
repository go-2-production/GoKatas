Let's consider a simple example of a web server built using Go, where goroutines can be beneficial. 

Imagine you have a web server that handles incoming HTTP requests and needs to perform some time-consuming tasks, such as making API calls or accessing a database. If the server were to handle these tasks sequentially, it would block and become unresponsive to other requests until each task is completed. This would lead to poor performance and a degraded user experience.

With goroutines, you can handle these time-consuming tasks concurrently, allowing the server to handle multiple requests simultaneously. Each incoming request can be assigned to a separate goroutine, and while one goroutine is waiting for an API call or database access, another goroutine can continue processing other requests. This concurrent execution ensures that the server remains responsive and can handle a higher number of concurrent users.

Here's a simplified example of how you could use goroutines in a web server:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Process the request asynchronously using a goroutine
	go processRequest(r)

	fmt.Fprint(w, "Request received, processing in the background.")
}

func processRequest(r *http.Request) {
	// Perform time-consuming tasks, such as making API calls or accessing a database
	// ...
}
```

In this example, the `handleRequest` function is invoked for each incoming HTTP request. It immediately starts a new goroutine by calling `go processRequest(r)`, which allows the server to continue accepting new requests without waiting for the time-consuming tasks to complete. The `processRequest` function performs the actual processing of the request asynchronously.

By leveraging goroutines, the server can handle multiple requests concurrently, ensuring that one slow request doesn't block others from being processed. This improves the responsiveness and scalability of the web server.

Note that in a real-world scenario, you would typically use channels to coordinate the communication and synchronization between goroutines, especially if you need to collect results or handle errors from the goroutines. However, the example above focuses on illustrating the basic concept of using goroutines in a web server.

# Kata Intructions

Using this technique let's call the endpoint that we have in the `server.go`
- ensure that we have tests for this as well using TDD best practices.
  
