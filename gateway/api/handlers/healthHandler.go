package handlers

import (
	"gateway/internal/public"
	"net/http"
	"time"
	"math/rand"
	"github.com/Meituan-Dianping/cat-go/cat"
)

func init() {
	cat.DebugOn()
	cat.Init("aiapigateway_ldev")
}

func reportCase() {
	t := cat.NewTransaction("case_in_health_check", "trans_1")
	defer t.Complete()

	e:= cat.NewEvent("case_in_health_check", "event_1")
	defer e.Complete()

	cat.NewMetricHelper("magric_1").Count(int(rand.Int31n(100)))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	go reportCase()
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(public.Logger(public.Message{"path": r.URL.String(), "status": true, "datetime": time.Now()}))
	return
}