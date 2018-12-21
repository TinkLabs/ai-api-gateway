package handlers

import (
	"fmt"
	"gateway/configs"
	"gateway/internal/public"
	"gateway/internal/valid"
	"net/http"
)

func Backend2(w http.ResponseWriter, r *http.Request) {
	// For re-route testing ["https://gateway-dev.handytravel.tech/apis/getlauncher?_barcode=357525081159035"]
	public.TimerStart()

	// Setup Request
	reqAuth := ""
	reqUrl := configs.Parms("DEST_HOST").(string) + r.URL.String()
	reqHost := configs.Parms("NAME_HOST").(string)
	reqHeader := make(public.JsonFormat)
	for key, val := range r.Header {
		if key == "Authorization" && val[0] != "" {
			reqAuth = val[0]
		}
		reqHeader[key] = val[0]
	}
	reqHeader["Cache-Control"] = "no-cache"
	reqHeader["X-Forwarded-For"] = r.RemoteAddr
	reqHeader["X-Forwarded-Proto"] = "https"
	reqHeader["X-Internal"] = "wrapper"

	// Setup Body
	r.ParseForm()
	body := make(public.JsonFormat)
	for key, val := range r.Form {
		body[key] = val[0]
	}

	// Valid Request
	if reqAuth != "" {
		reqAccount, err := valid.AuthDeviceUser(reqAuth)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), 500)
			return
		} else {
			for key, val := range reqAccount {
				body[key] = val
			}
		}
	}

	// Make Request
	response, header, err := public.NewRequest(reqUrl, r.Method, reqHost, reqHeader, body)

	// Request Error
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// Response Header setup
	for key, val := range header {
		w.Header().Add(key, val.(string))
	}
	originName := public.OriginRequest(string(r.Header.Get("Origin")), configs.Parms("REQUEST_HOST").(string))
	if originName != "" {
		w.Header().Set("Access-Control-Allow-Origin", originName)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// For application Log
	public.Logger(public.Message{"prefix": "backend2", "path": r.URL.String(), "body": body, "timecount": public.TimerEnd()})

	// Print the Result
	fmt.Fprintf(w, "%s", response)
	return
}
