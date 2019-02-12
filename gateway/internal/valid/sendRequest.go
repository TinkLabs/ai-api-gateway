package valid

import (
	//"bytes"
	"fmt"
	"gateway/configs"
	"gateway/internal/public"
	"io/ioutil"

	//"io"
	"net/http"
)

func SendRequest(w http.ResponseWriter, r *http.Request,reqUrl string)  {

	reqHeader   := make(public.JsonFormat)
	for key, val := range r.Header {
		reqHeader[key] = val[0]
	}
	reqHeader["Cache-Control"] = "no-cache"
	reqHeader["X-Forwarded-For"] = r.RemoteAddr
	reqHeader["X-Forwarded-Proto"] = "https"
	//reqHeader["X-by-pass-jwt"] = "true"
	reqHeader["Authorization"] = "73bef7wr4kw84vfu8hbrudvmfudy"

	// Setup Body
	//r.ParseForm()
	//body := make(public.JsonFormat)
	//for key, val := range r.Form {
	//	body[key] = val[0]
	//}
	//
	//fmt.Printf("body is %v\n",body)

	//var reqParams io.Reader
	//body, err := r.Body;
	//if  err == nil {
	//	reqParams = body
	//} else {
	//	reqParams = bytes.NewReader([]byte{})
	//}

	reqParam,_ := ioutil.ReadAll(r.Body)

	// Make Request
	response, header, err := public.NewRequest(reqUrl, r.Method, "", reqHeader, reqParam)

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
	public.Logger(public.Message{"path": reqUrl, "body":string(reqParam), "timecount": public.TimerEnd()})

	//fmt.Printf("\n response %s",response)
	// Print the Result
	fmt.Fprintf(w, "%s", response)
	return
}
