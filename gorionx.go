package gorionx

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// api endpoint
const (
	URL = "https://api2.orionx.io/graphql"
)

// Get secret_key, timestamp and query and return hmac512
func GetHmac512(secret_key, message string) string {
	mac := hmac.New(sha512.New, []byte(secret_key))
	mac.Write([]byte(message))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func MakeRequest(KEY, API_KEY string, requestString string) {

	queryconv := `{"query":` + strconv.QuoteToASCII(requestString) + `}`
	fmt.Printf("QUER %s", queryconv) //debug

	var str = []byte(queryconv)
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(str))

	//data to header
	now := int32(time.Now().Unix()) // timestamp
	str_timestamp := fmt.Sprintf("%d", now)

	//get length of request
	length_request := len(queryconv)
	str_length_request := fmt.Sprintf("%d", length_request)

	//calculate hmac (signature)
	message := str_timestamp + queryconv
	str_hmac512 := GetHmac512(KEY, message)

	//set header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-ORIONX-TIMESTAMP", str_timestamp)
	req.Header.Set("X-ORIONX-APIKEY", API_KEY)
	req.Header.Set("X-ORIONX-SIGNATURE", str_hmac512) // secret+me
	req.Header.Set("Content-Length", str_length_request)

	//show full request
	//fmt.Printf("\nRequest: %s\n\n", req) //debug

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("\nStatus:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}
