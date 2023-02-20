package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mococa/go-api-v2/pkg/db/models"
	"github.com/mococa/go-api-v2/pkg/internal"
	"github.com/mococa/go-api-v2/pkg/utils"
)

func init() {
	type Response struct {
		Book models.Book `json:"book"`
	}

	handle := func(h *internal.Handler) {
		req := h.Request
		res := h.Response
		modules := h.Modules

		params := mux.Vars(req)

		book_id := params["book_id"]

		book, err := modules.Book.FindBook(book_id)

		if utils.GetHTTPError(err, res, 500) != nil {
			return
		}

		response, err := json.Marshal(Response{
			Book: *book,
		})

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(`{"error": "Error marshalling"}`))
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write([]byte(response))
	}

	route := &Route{
		path:    "/books/{book_id}",
		method:  "GET",
		handler: handle,
	}

	Routes.Push(route)
}
