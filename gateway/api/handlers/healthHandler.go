package handlers

import (
	"gateway/internal/public"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Write(public.Logger(public.Message{"path": r.URL.String(), "status": true, "datetime": time.Now()}))
	return
}