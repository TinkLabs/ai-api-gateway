package routers

import (
	"net/http"
	"net/url"
	"net/http/httptest"
	"testing"
	"fmt"
	"strings"
	"gateway/internal/public"
	"github.com/julienschmidt/httprouter"


)

func TestApi(t *testing.T) {
	fmt.Println("msg")

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

	reqURL := "localhost:80/ai/v1/user/hotel_id_verification?hotel_id=123"

	data := url.Values{}
	data.Set("hotel_id", "123")
	//data.Add("surname", "bar")

	req, err := http.NewRequest("GET", reqURL,
		strings.NewReader(data.Encode()))

	req.Header.Add("X-Signature", "pk6grd4scjouhbOJ3aZQliy2A4VvKuO1Jb8lyLxYJHM=")
	req.Header.Add("X-User-Id", "1226")

	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	//handler := http.Handler(Api)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	//handler.ServeHTTP(rr, req)

	ps ,_ := req.Context().Value("params").(httprouter.Params)
	Api(rr,req,ps)

	//rr := newRequestRecorder(req, "GET", "/ai/v1/user/hotel_id_verification", Api)

	//// Check the status code is what we expect.
	//if status := rr.Code; status != http.StatusOK {
	//	t.Errorf("handler returned wrong status code: got %v want %v",
	//		status, http.StatusOK)
	//}

	public.Logger(public.Message{"path": reqURL, "body": data, "header": req.Header})



	// Check the response body is what we expect.
	//expected := `{"alive": true}`
	//if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}

}

// Mocks a handler and returns a httptest.ResponseRecorder
func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	router.ServeHTTP(rr, req)
	return rr
}
