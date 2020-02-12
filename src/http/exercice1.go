package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func Home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the hand
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Radhouen"))
}

// Add a showSnippet handler function.
func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"snippet"}`))
	/*w.Write([]byte("Display a specific snippet..."))*/
}

// Add a createSnippet handler function.
func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// Use the Header().Set() method to add an 'Allow: POST' header to the
		// response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func getSnippetById(w http.ResponseWriter, r*http.Request)  {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	// Use the http.ListenAndServe() function to start a new web server. We pas
	// two parameters: the TCP network address to listen on (in this case ":4000
	// and the servemux we just created. If http.ListenAndServe() returns an er
	// we use the log.Fatal() function to log the error message and exit.
	mux.HandleFunc("/snippets", ShowSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippet)
	mux.HandleFunc("/snippet", getSnippetById)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

	/*----------------------- DefaultServeMux ---------------------------*/
	/*http.HandleFunc("/", Home)
	http.HandleFunc("/snippet", ShowSnippet)
	http.HandleFunc("/snippet/create", CreateSnippet)

	log.Println("Starting server on :4000")
	err1 := http.ListenAndServe(":4000", nil)
	log.Fatal(err1)*/
}
