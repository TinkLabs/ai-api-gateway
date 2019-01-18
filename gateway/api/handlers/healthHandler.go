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
	cat.Init("aiapigateway")
}

func reportCase() {
	t := cat.NewTransaction("case_in_health_check_by_domain_2", "trans_1")
	defer t.Complete()

	e:= cat.NewEvent("case_in_health_check_by_domain_2", "event_1")
	defer e.Complete()

	cat.NewMetricHelper("magric_1_by_domain_2").Count(int(rand.Int31n(100)))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	go reportCase()
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(public.Logger(public.Message{"path": r.URL.String(), "status": true, "datetime": time.Now()}))
	return
}