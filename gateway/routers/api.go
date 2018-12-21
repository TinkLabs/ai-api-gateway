package routers

import (
	"gateway/api/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	//"fmt"
	"gateway/internal/public"
)

func Api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if ps[0].Value == "" || ps[0].Value == "/" {
		http.Error(w, "Unexpected Request Path", 502)
		return
	}
	pathArray := strings.Split(ps[0].Value, "/")
	if len(pathArray) < 3 {
		http.Error(w, "Unexpected Request Path", 502)
		return
	}
	path := pathArray[1] + "/" + pathArray[2]
	switch path {
	case "healthcheck/":
		handlers.HealthCheck(w, r)
	case "ai/v1":
		handlers.AiServices(w, r)
	case "ai/portal-stats":
		handlers.StatisticsServices(w, r, "STATISTICS")
	default:
		public.Logger(public.Message{"msg": "Unexpect Request Path"})
		//fmt.Fprintf("Unexpect Request Path")
		http.Error(w, "Unexpected Request Path", 502)
	}
	return
}
