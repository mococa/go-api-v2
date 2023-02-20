package routes

import (
	"encoding/json"
	"net/http"

	"github.com/mococa/go-api-v2/pkg/db/models"
	"github.com/mococa/go-api-v2/pkg/internal"
	"github.com/mococa/go-api-v2/pkg/internal/dtos"
	"github.com/mococa/go-api-v2/pkg/utils"
)

func init() {
	type Response struct {
		Book *models.Book `json:"book"`
	}

	handle := func(h *internal.Handler) {
		req := h.Request
		res := h.Response
		modules := h.Modules

		body := dtos.CreateBookBody{}

		if utils.GetHTTPBody(res, req, &body) != nil {
			return
		}

		book, err := modules.Book.CreateBook(&body)

		if utils.GetHTTPError(err, res, 500) != nil {
			return
		}

		response, err := json.Marshal(Response{
			Book: book,
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
		method:  "POST",
		handler: handle,
	}

	Routes.Push(route)
}
