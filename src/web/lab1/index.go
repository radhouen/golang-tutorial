package main
import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	port := "8083"

	ln, err := net.Listen("tcp", ":" + port)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s", port, err)
		os.Exit(1)
	}
    ln.Close()
	fmt.Printf("TCP Port %q is available", port)
	http.ListenAndServe(":8083", nil)
}
