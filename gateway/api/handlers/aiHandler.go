package handlers

import (
    "gateway/internal/valid"
    "gateway/internal/public"
    "gateway/configs"
    "net/http"
    "fmt"
)

func AiServices(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Ai Service")

    public.TimerStart()

    //VerifySignature
    signatureCheck := valid.VerifySignature(w,r)
    if signatureCheck == ""{
        http.Error(w, "access denied,unexpected signature or X-User-Id should not be null", 400)
        return
    }

    fmt.Println("begin to sent request")
    // Setup Request
    reqUrl      := configs.Parms("NAME_HOST").(string) + r.URL.String()

    fmt.Println("reqUrl %s\n",reqUrl)

    //send request and return response
    valid.SendRequest(w,r,reqUrl)

}

