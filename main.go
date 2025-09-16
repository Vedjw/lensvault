package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to my great site</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:vjw55555@gmail.com\">vjw55555@gmail.com</a></p>")
}

func main() {
	http.Handle("/", http.HandlerFunc(homeHandler))
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting Server on :3000")
	http.ListenAndServe(":3000", nil)
}
