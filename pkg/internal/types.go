package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mococa/go-api-v2/pkg/modules"
)

type Server struct {
	Modules modules.Modules
	Router  *mux.Router
	Port    string
	Start   func()
}

type Handler struct {
	Modules  modules.Modules
	Request  *http.Request
	Response http.ResponseWriter
}
