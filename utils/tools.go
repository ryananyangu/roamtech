package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"os"
	"strings"
	"time"
)

// MapToFormStr convert map to form string
func MapToFormStr(object map[string]string) string {
	converted := ""
	for lable, value := range object {
		converted += url.QueryEscape(lable) + "=" + url.QueryEscape(value) + "&"
	}
	return strings.TrimSuffix(converted, "&")
}

// Request post string payload to endpoint
func Request(request string, headers map[string][]string, urlPath string, method string) (string, error) {

	reqURL, _ := url.Parse(urlPath)

	reqBody := ioutil.NopCloser(strings.NewReader(request))

	req := &http.Request{
		Method: method,
		URL:    reqURL,
		Header: headers,
		Body:   reqBody,
	}

	// res, err := http.DefaultClient.Do(req)

	res, err := ExternalRequestTimer(req)
	if err != nil {
		// log http error

		ErrorLogger.Printf("SEND REQUEST | URL : %s | METHOD : %s | BODY : %s | ERROR : %v", urlPath, method, request, err)
		return "", err
	}

	data, _ := ioutil.ReadAll(res.Body)

	InfoLogger.Printf("SEND REQUEST | URL : %s | METHOD : %s | BODY : %s | STATUS : %s | HTTP_CODE : %d", urlPath, method, request, res.Status, res.StatusCode)

	if res.StatusCode > 299 || res.StatusCode <= 199 {
		ErrorLogger.Printf("SEND REQUEST | URL : %s | METHOD : %s | BODY : %s | STATUS : %s | HTTP_CODE : %d", urlPath, method, request, res.Status, res.StatusCode)
		return res.Status, fmt.Errorf("%d", res.StatusCode)
	}

	res.Body.Close()

	resbody := string(data)
	InfoLogger.Printf("SEND REQUEST | URL : %s | METHOD : %s | BODY : %s | STATUS : %s | HTTP_CODE : %d", urlPath, method, "resbody", res.Status, res.StatusCode)

	return resbody, nil
}

// ReadJSON utility function to read the json file into a map[string]interface{}
func ReadJSON(file string) map[string]string {
	jsonFile, err := os.Open(file)
	if err != nil {
		// fmt.Println(err)
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]string
	json.Unmarshal(byteValue, &result)

	return result

}

func ExternalRequestTimer(req *http.Request) (*http.Response, error) {
	// req, _ := http.NewRequest("GET", url, nil)

	var start, connect, dns, tlsHandshake time.Time

	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			InfoLogger.Printf("DNS Done: %v\n", time.Since(dns))
		},

		TLSHandshakeStart: func() { tlsHandshake = time.Now() },
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			InfoLogger.Printf("TLS Handshake: %v\n", time.Since(tlsHandshake))
		},

		ConnectStart: func(network, addr string) { connect = time.Now() },
		ConnectDone: func(network, addr string, err error) {
			InfoLogger.Printf("Connect time: %v\n", time.Since(connect))
		},

		GotFirstResponseByte: func() {
			InfoLogger.Printf("Time from start to first byte: %v\n", time.Since(start))
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return res, err
	}
	return res, nil
}
