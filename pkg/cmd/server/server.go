package server

import (
	"log"
	"net/http"

	"github.com/mococa/go-api-v2/pkg/cmd/server/routes"
	"github.com/mococa/go-api-v2/pkg/internal"
	"github.com/mococa/go-api-v2/pkg/modules"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func NewServer(port string, db *gorm.DB) *internal.Server {
	server := &internal.Server{
		Modules: modules.NewModules(db),
		Router:  nil,
		Port:    port,
		Start:   nil,
	}

	routes.NewRouter(server)

	server.Start = func() {
		startServer(server)
	}

	return server
}

func startServer(s *internal.Server) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		Debug:            false,
	})

	handler := c.Handler(s.Router)

	log.Println("Server listening on port", s.Port)

	routes.DisplayRoutes()

	log.Fatalln(http.ListenAndServe(s.Port, handler))
}
