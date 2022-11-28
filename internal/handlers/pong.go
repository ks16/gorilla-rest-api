package handlers

import (
	"github.com/ks16/gorilla-rest-api/pkg/log"
	"net/http"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	lg := log.Ctx(r.Context())
	lg.Info().Msg("GET /pond request")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("PONG"))

	return
}
