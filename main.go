package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"net/http/fcgi"

	"github.com/davecgh/go-spew/spew"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Content-Type", "text/html")
	url := html.EscapeString(r.URL.String())
	req := html.EscapeString(spew.Sdump(r))
	fmt.Fprintf(w, "<html><head></head><body><p>Hello from %s</p><pre>%s</pre></body></html>", url, req)
}

func main() {
	if err := fcgi.Serve(nil, &Handler{}); err != nil {
		log.Fatal(err)
	}
}
