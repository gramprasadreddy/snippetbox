package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}
func snippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display from specific snippet"))
}
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// StatusMethodnotAllowed
		w.Header().Set("Allow", "POST")
		//w.WriteHeader(405)
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		//w.Write([]byte("Method not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", snippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
