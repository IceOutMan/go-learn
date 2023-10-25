package main

import (
	"fmt"
	http "net/http"
	"net/http/httputil"
	"strings"
)

func raw_get() {
	url := "http://www.baidu.com"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit")
	request.AddCookie(&http.Cookie{
		Name:  "test_name",
		Value: "test_value",
	})

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// use httputil -> DumpResponse -> byte[]
	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return
	}
	fmt.Print(string(response))
}

func raw_post() {

	url := "http://www.baidu.com"
	payload := "key1=value1&key2=value2" // 您的POST数据

	request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		return
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit")
	request.AddCookie(&http.Cookie{
		Name:  "test_name",
		Value: "test_value",
	})

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// use httputil -> DumpResponse -> byte[]
	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return
	}
	fmt.Print(string(response))
}

func get() {
	url := "http://www.baidu.com"
	resp, err := http.Get(url)

	defer resp.Body.Close()

	if err != nil {
		return
	}

	// use httputil -> DumpResponse -> byte[]
	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return
	}
	fmt.Print(string(response))
}

func post() {
	url := "http://www.baidu.com"
	payload := "name=LS&age=18"

	contentType := "application/x-www-form-urlencoded"
	resp, err := http.Post(url, contentType, strings.NewReader(payload))

	defer resp.Body.Close()
	if err != nil {
		return
	}

	// use httputil -> DumpResponse -> byte[]
	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return
	}
	fmt.Print(string(response))
}
