package routes

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/gorilla/mux"
	"github.com/mococa/go-api-v2/pkg/internal"
	array "github.com/mococa/go-array"
)

type ModuleHandler func(*internal.Handler)

type Route struct {
	path    string
	method  string
	handler ModuleHandler
}

var Routes = array.NewGoArray[*Route]()

const blue_color = "\033[34m"
const reset_color = "\033[0m"

func NewRouter(s *internal.Server) {
	router := mux.NewRouter()

	fmt.Println("\n--------- Routes --------")

	Routes.Each(func(route *Route) {
		m := s.Modules
		h := route.handler

		router.HandleFunc(route.path,
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

				h(
					&internal.Handler{
						Modules:  m,
						Request:  r,
						Response: w,
					},
				)
			},
		).Methods(route.method)
	})

	s.Router = router
}

func DisplayRoutes() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	Routes.Each(func(route *Route) {
		message := fmt.Sprintf("%s\t%s[%s]%s", route.path, blue_color, route.method, reset_color)
		fmt.Fprintln(w, message)
	})

	w.Flush()
}
