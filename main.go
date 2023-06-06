package main

import (
	"fmt"
	"net/http"
)

func write(w http.ResponseWriter, text string) (int, error) {
	return w.Write([]byte(text + "\n"))
}

// MyHandler Represents my handler
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "text/html; utf-8")
	w.WriteHeader(200)

	_, _ = write(w, "<html><body>")
	write(w, "<p>Possible urls:</p>")
	write(w, "<ul><li>/json</li>")
	write(w, "<li>/hello</li>")
	write(w, "<li>/foo</li>")
	write(w, "<li>/products</li>")
	write(w, "</ul>")
	write(w, "</body></html>")
	fmt.Fprint(w)
}

// HandleHello responds with hello
func HandleHello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, "{hello: true}")
}

func main() {

	host := "0.0.0.0:9090"
	http.Handle("/", new(MyHandler))
	http.HandleFunc("/hello", HandleHello)
	http.HandleFunc("/foo", HandleFoo)
	http.HandleFunc("/json", HandleJSON)
	http.HandleFunc("/products", HandleProduct)

	var err error

	defer func() {
		err = http.ListenAndServe(host, nil)
		fmt.Println("Server started a")
	}()

	if err != nil {
		fmt.Println("oops", err)
	}
}
