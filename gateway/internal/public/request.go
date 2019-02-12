package public

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type JsonFormat map[string]interface{}

func NewRequest(path string, method string, host string, header JsonFormat, param []byte) ([]byte, JsonFormat, error) {

	err := errors.New("Incorrect request")
	respHeader := make(JsonFormat)
	if path == "" || method == "" {
		return nil, respHeader, err
	}

	// Request Path
	req, err := http.NewRequest(method, path, bytes.NewReader(param))
	if err != nil {
		return nil, respHeader, err
	}

	// Request Host
	if host != "" {
		req.Host = host	
	}

	// Request Header
	for key, val := range header {
		req.Header.Set(key, val.(string))
	}

	// Request have been maken
	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, respHeader, err
    }
    defer resp.Body.Close()


	// Response
    for key, val := range resp.Header {
		respHeader[key] = val[0]
    }
    response, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, respHeader, err
    }

    return response, respHeader, err
}

func OriginRequest(origin string, originKey string) string {

	keyword := strings.Index(origin, originKey)
	if keyword > -1 {
		return origin
	}
	return ""
}