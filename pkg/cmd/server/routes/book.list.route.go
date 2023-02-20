package routes

import (
	"encoding/json"
	"net/http"

	"github.com/mococa/go-api-v2/pkg/db/models"
	"github.com/mococa/go-api-v2/pkg/internal"
	"github.com/mococa/go-api-v2/pkg/utils"
)

func init() {
	type Response struct {
		Books []models.Book `json:"books"`
	}

	handle := func(h *internal.Handler) {
		res := h.Response
		modules := h.Modules

		books, err := modules.Book.ListBooks()

		if utils.GetHTTPError(err, res, 500) != nil {
			return
		}

		response, err := json.Marshal(Response{
			Books: *books,
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
		path:    "/books",
		method:  "GET",
		handler: handle,
	}

	Routes.Push(route)
}
