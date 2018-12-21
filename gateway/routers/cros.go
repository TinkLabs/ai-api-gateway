package routers

import (
	"gateway/configs"
	"gateway/internal/public"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Cros(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	origin := string(r.Header.Get("Origin"))
	originKey := configs.Parms("REQUEST_HOST").(string)
	originName := public.OriginRequest(origin, originKey)
	public.Logger(public.Message{"prefix": "CROS", "origin": origin, "originKey": originKey, "return": originName})
	if originName != "" {
		w.Header().Set("Access-Control-Allow-Origin", originName)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	w.WriteHeader(200)
	return
}
