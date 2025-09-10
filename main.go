package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 1. Define the flags
	portPtr := flag.Int("port", 3000, "The port for the server to listen on.")
	messagePtr := flag.String("message", "Hello, World!", "A custom message to display.")

	// 2. Parse the flags
	flag.Parse()

	// 3. Use the flag values
	port := *portPtr
	message := *messagePtr

	
	// 4. Create a handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK) // Returns a 200 OK status
    })

	// 5. Start the server
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Server starting on port %d with message: \"%s\"\n", port, message)
	log.Fatal(http.ListenAndServe(addr, nil))
}