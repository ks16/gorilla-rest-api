package handlers

import (
	"github.com/ks16/gorilla-rest-api/pkg/log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	lg := log.Ctx(r.Context())
	lg.Info().Msg("heals-check request")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

	return
}
