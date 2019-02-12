package valid

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	//"encoding/hex"
	"fmt"

	"net/http"
	"sort"
	"io/ioutil"
	"encoding/base64"

)

func VerifySignature(w http.ResponseWriter, r *http.Request) string {

	//get X-User-Id from header
	var xUserId string = ""

	//get signature from header
	var signatureFromRequest string = ""
	for key, val := range r.Header {
		if key == "X-Signature" && val[0] != "" {
			signatureFromRequest = val[0]
		}
		if key == "X-User-Id" && val[0] != "" {
			xUserId = val[0]
		}
	}
	fmt.Printf("xuserid: %s\n",xUserId)
	fmt.Printf("getSignature : %s\n",signatureFromRequest)

	if signatureFromRequest == "" || xUserId == ""{
		fmt.Println("signature or X-User-Id  is null ")
		//return ""
	}

	var params string = ""
	if r.Method == http.MethodGet {
		params = GenSignatureForGet(w,  r)
	} else if r.Method == http.MethodPost || r.Method == http.MethodPatch {
		params = GenSignatureForPostAndPutAndPatch(w,  r)
	}

	fmt.Printf("params : %s\n",params)

	//createSignature
	signature := GenSignature(params + "+X-User-Id=" + xUserId)

	fmt.Printf("signatureResult : %s",signature)

	if signatureFromRequest != signature{
		fmt.Println("signature not equal")
		return ""
	}

	return signature

}

func GenSignature(data string) string {

	secret := "/r9o3VKyp1/7mJYfxTMond/4vH8i2EWzbODqcl0AhzI="
	//data := "data"
	fmt.Printf("SecretData: %s\n" , data)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	//sha := hex.EncodeToString(h.Sum(nil))
	sha := base64.StdEncoding.EncodeToString(h.Sum(nil))

	//fmt.Println("Result: " + sha)
	return sha
}

func GenSignatureForGet(w http.ResponseWriter, r *http.Request) string {

	// get the value of params
	query := r.URL.Query()
	fmt.Printf("query: %s\n",query)

	/** sort the request params
		1: init a slice with map query
		2: sort the slice since map can't use sort
	 */
	paramsKey := make([]string, len(query))
	i := 0
	for k, _ := range query {
		paramsKey[i] = k
		i++
	}

	sort.Strings(paramsKey)
	fmt.Printf("paramsKey result : %s\n",paramsKey)

	var params string = ""
	j := 0

	for _, k := range paramsKey {
		fmt.Println("Key:", k, "Value:", query[k][0])

		//key
		params = fmt.Sprintf("%s%s",params,k)

		//value
		if j == len(query)-1 {
			params = fmt.Sprintf("%s=%s",params,query[k][0])
		}else{
			params = fmt.Sprintf("%s=%s&",params,query[k][0])
		}
		j++
	}

	return params

}

func GenSignatureForPostAndPutAndPatch(w http.ResponseWriter, r *http.Request) string {

	// get the value of body
	body, _ := ioutil.ReadAll(r.Body)
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	var params string = ""

	params = fmt.Sprintf("%s", body)

	return params
}