package main

import (
	"fmt"
	"net/http"
)

func getAllUsers(w http.ResponseWriter , r *http.Request) string {

	return ""
}
func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter , r *http.Request) {
		w.Write([]byte("Hello world"))
		fmt.Println("request method", r.Method)
	})
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:2600", nil)
}
