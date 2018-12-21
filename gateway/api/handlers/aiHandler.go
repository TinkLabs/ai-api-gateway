package handlers

import (
	"fmt"
	"gateway/configs"
	"gateway/internal/public"
	"gateway/internal/valid"
	"net/http"
)

func AiServices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ai Service")

	public.TimerStart()

	//VerifySignature
	signatureCheck := valid.VerifySignature(w, r)
	if signatureCheck == "" {
		http.Error(w, "access denied,unexpected signature or X-User-Id should not be null", 400)
		return
	}

	fmt.Println("begin to sent request")
	// Setup Request
	reqURL:= configs.Parms("NAME_HOST").(string) + r.URL.String()

	fmt.Printf("reqUrl %s\n", reqURL)

	//send request and return response
	valid.SendRequest(w, r, reqURL)

}
