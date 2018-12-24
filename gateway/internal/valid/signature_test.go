package valid

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

func TestVerifySignature(t *testing.T) {
	fmt.Println("TestAiServices")

	MyfuncVerifySignature("pk6grd4scjouhbOJ3aZQliy2A4VvKuO1Jb8lyLxYJHM=",t)
	fmt.Println("case1 done")

	MyfuncVerifySignature("",t)
	fmt.Println("case2 done")

}

func MyfuncVerifySignature(signature string,t *testing.T){

	// Create a request to pass to our handler.

	reqURL := "/ai/v1/user/hotel_id_verification?hotel_id=123"

	data := url.Values{}
	data.Set("hotel_id", "123")
	//data.Add("surname", "bar")

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
	VerifySignature(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	public.Logger(public.Message{"path": reqURL, "body": data, "header": req.Header})



	// Check the response body is what we expect.
	//t.Errorf("handler returned unexpected body: got %v want %v",
	//	rr.Body.String(), "")

	//expected := `{"alive": true}`
	//if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}
}
