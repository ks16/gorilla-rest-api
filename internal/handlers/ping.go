package handlers

import (
	"github.com/ks16/gorilla-rest-api/pkg/httpclient"
	"github.com/ks16/gorilla-rest-api/pkg/log"
	"io"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	lg := log.Ctx(r.Context())
	lg.Info().Msg("GET /ping request")

	c := httpclient.NewClient(r.Context())
	req, err := http.NewRequest("GET", "http://service-b:8082/pong", nil)
	if err != nil {
		lg.Err(err).Msg("failed to create request to service B")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		lg.Err(err).Msg("cannot call service B")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		lg.Err(err).Msgf("cannot read response body: %+v", resp.Body)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)

	return
}
