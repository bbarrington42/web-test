
package main

import (
	"html/template"
	"log"
	"net/http"
)

// Serve up the Hello message
func main() {
	http.Handle("/", NewServer())
	http.Handle("/ping", new(PingHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Server implements the server.
// It serves the user interface (it's an http.Handler)
type Server struct {
}

type PingHandler struct {}
func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

// NewServer returns an initialized server.
func NewServer() *Server {
	s := &Server{}
	return s
}

// ServeHTTP implements the HTTP user interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Name string
	}{
		"Bill",
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
	}
}

// tmpl is the HTML template that drives the user interface.
var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html><html><body><center>
	<h2>Hello, {{.Name}}!</h2>
</center></body></html>
`))
