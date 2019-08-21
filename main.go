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

	write(w, "<html><body>")
	write(w, "bunda <b>mole</b>")
	write(w, "<ul><li>lelele</li></ul>")
	write(w, "</body></html>")
	fmt.Fprint(w)
}

func main() {

	host := "0.0.0.0:9090"
	http.Handle("/", new(MyHandler))

	var err error

	defer func() {
		err = http.ListenAndServe(host, nil)
		fmt.Println("Server started a")
	}()

	if err != nil {
		fmt.Println("oops", err)
	}
}
