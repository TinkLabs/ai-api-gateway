package valid

import (
    "gateway/internal/public"
    "gateway/configs"
    "encoding/json"
)

func AuthDeviceUser(authToken string) (public.JsonFormat, error) {

	public.TimerStart()
    reqUrl := configs.Parms("AUTH_HOST").(string) + "/oauth/verify"
    reqHeader := make(public.JsonFormat)
    reqHeader["Authorization"] = authToken
    response, _, err := public.NewRequest(reqUrl, "POST", "", reqHeader, nil)
    
    // Change the response to Map
    jsonMap := make(public.JsonFormat)
    err = json.Unmarshal([]byte(response), &jsonMap)
    if err != nil {
        return nil, err
    }

    // For application Log
    public.Logger(public.Message{"path": "/authenticate", "authToken":authToken, "timecount": public.TimerEnd()})

    // Return Formater
    jsonMap["code"] = int(jsonMap["code"].(float64))
    if jsonMap["code"] == 400 {
        jsonMap = nil
    }
    return jsonMap, err
}


