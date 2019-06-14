package routers

import (
	"fmt"
	"gateway/api/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	//"fmt"
	"gateway/internal/public"
)

func Api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	path := strings.Split(ps[0].Value, "/");

	profix := strings.Join(path[1:2], "/")
	fmt.Printf("%v\n", path);
	fmt.Println(profix)

	switch  profix {
	case "healthcheck":
		handlers.HealthCheck(w, r)
	case "ai/v1":
		handlers.AiServices(w, r)
	case "ai/portal-stats":
		handlers.StatisticsServices(w, r, "STATISTICS")
	default:
		public.Logger(public.Message{"msg": "Unexpect Request Path"})
		//fmt.Fprintf("Unexpect Request Path")
		http.Error(w, "Unexpected Request Path", 404)
	}
	return
}
