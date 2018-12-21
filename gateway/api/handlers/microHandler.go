package handlers

import (
	"fmt"
	"gateway/configs"
	"gateway/internal/public"
	"gateway/internal/valid"
	"net/http"
	"strings"
)

func StatisticsServices(w http.ResponseWriter, r *http.Request, prefix string) {
	fmt.Println("Statistics Service")

	public.TimerStart()

	//VerifySignature
	signatureCheck := valid.VerifySignature(w, r)
	if signatureCheck == "" {
		http.Error(w, "access denied,unexpected signature or X-User-Id should not be null", 400)
		return
	}

	// Setup Request
	path := configs.Parms(prefix + "_HOST")
	if path == nil {
		fmt.Println(prefix + ": undefined config")
		http.Error(w, "undefined config", 500)
		return
	}

	pathEnd := strings.Replace(r.URL.String(), "ai/portal-stats/", "", 15)

	reqURL := path.(string) + pathEnd
	fmt.Printf("reqURL %s\n", reqURL)

	//send request and return response
	valid.SendRequest(w, r, reqURL)

}
