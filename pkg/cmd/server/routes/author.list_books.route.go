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
		Books []models.Book `json:"books"`
	}

	handle := func(h *internal.Handler) {
		req := h.Request
		res := h.Response
		modules := h.Modules

		params := mux.Vars(req)

		author_id := params["author_id"]

		author_books, err := modules.Author.ListAuthorBooks(author_id)

		if utils.GetHTTPError(err, res, 500) != nil {
			return
		}

		response, err := json.Marshal(Response{
			Books: *author_books,
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
		path:    "/authors/{author_id}/books",
		method:  "GET",
		handler: handle,
	}

	Routes.Push(route)
}
