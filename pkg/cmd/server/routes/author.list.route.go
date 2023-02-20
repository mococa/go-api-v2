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
		Authors []models.Author `json:"authors"`
	}

	handle := func(h *internal.Handler) {
		res := h.Response
		modules := h.Modules

		authors, err := modules.Author.ListAuthors()

		if utils.GetHTTPError(err, res, 500) != nil {
			return
		}

		response, err := json.Marshal(Response{
			Authors: *authors,
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
		path:    "/authors",
		method:  "GET",
		handler: handle,
	}

	Routes.Push(route)
}
