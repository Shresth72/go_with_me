package httpservers

import (
	"html/template"
	"net/http"
)

type Handler interface {
  ServeHTTP(http.ResponseWriter, r *http.Request)
}

var t = template.Must(template.New("hello"), {{.}}")
