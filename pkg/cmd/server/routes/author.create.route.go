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
		Author *models.Author `json:"author"`
	}

	handle := func(h *internal.Handler) {
		req := h.Request
		res := h.Response
		modules := h.Modules

		body := dtos.CreateAuthorBody{}

		if utils.GetHTTPBody(res, req, &body) != nil {
			return
		}

		author, err := modules.Author.CreateAuthor(&body)

		if utils.GetHTTPError(err, res, 500) != nil {
			return
		}

		response, err := json.Marshal(Response{
			Author: author,
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
		method:  "POST",
		handler: handle,
	}

	Routes.Push(route)
}
