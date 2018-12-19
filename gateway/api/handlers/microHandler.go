package handlers

import (
    "gateway/internal/valid"
    "gateway/internal/public"
    "gateway/configs"
    "net/http"
    "fmt"
    "strings"
)

func StatisticsServices(w http.ResponseWriter, r *http.Request, prefix string) {
    fmt.Println("Statistics Service")

    public.TimerStart()

    //VerifySignature
    signatureCheck := valid.VerifySignature(w,r)
    if signatureCheck == ""{
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

    reqUrl      := path.(string) + pathEnd
    fmt.Printf("reqUrl %s\n",reqUrl)

    //send request and return response
    valid.SendRequest(w,r,reqUrl)

}