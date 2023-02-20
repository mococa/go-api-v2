package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetHTTPBody(res http.ResponseWriter, req *http.Request, model interface{}) error {
	res.Header().Set("Content-Type", "application/json")
	// res.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	err := json.NewDecoder(req.Body).Decode(&model)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"error": "Error unmarshalling"}`))

		return err
	}

	return nil
}

func GetHTTPError(err error, res http.ResponseWriter, status int) error {
	res.Header().Set("Content-Type", "application/json")

	if err != nil {
		res.WriteHeader(status)

		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))

		return err
	}

	return nil
}
