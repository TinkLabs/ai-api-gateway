package handlers

import (
	"net/http"
	"net/url"
	"net/http/httptest"
	"testing"
	"fmt"
	"strings"
	"gateway/internal/public"
	//"gateway/api/handlers"


)

func TestStatisticsServices(t *testing.T) {
	fmt.Println("StatisticsServices")

	MyfuncStatisticsServices("GDX0hbr5oD2wRKvikdvlw36Jztf7DguJJr9X56qEfsk=",t,"STATISTICS")
	fmt.Println("case1 done")

	MyfuncStatisticsServices("",t,"STATISTICS")
	fmt.Println("case2 done")

	MyfuncStatisticsServices("GDX0hbr5oD2wRKvikdvlw36Jztf7DguJJr9X56qEfsk=",t,"STATISTICS1")
	fmt.Println("case3 done")

	MyfuncStatisticsServices("",t,"STATISTICS1")
	fmt.Println("case4 done")

}

func MyfuncStatisticsServices(signature string,t *testing.T,prefix string){

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

	reqURL := "/ai/portal-stats/statistics/v1/device?hotel_id=1&date_start=2018-01-01&date_end=2018-11-01"

	data := url.Values{}
	data.Set("hotel_id", "1")
	data.Add("date_start", "2018-01-01")
	data.Add("date_end", "2018-11-01")

	req, err := http.NewRequest("GET", reqURL,
		strings.NewReader(data.Encode()))

	req.Header.Add("X-Signature", signature)
	req.Header.Add("X-User-Id", "1226")

	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	//handler := http.HandlerFunc(AiServices)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	StatisticsServices(rr, req,prefix)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	public.Logger(public.Message{"path": reqURL, "body": data, "header": req.Header})



	// Check the response body is what we expect.

	//expected := `{"alive": true}`
	//if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}
}

