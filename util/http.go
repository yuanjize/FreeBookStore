package util

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func RequestGet(host string, params map[string]interface{}, headers map[string]string) *http.Request {
	logger := log.New(os.Stdout, "[RequestGet]", log.LstdFlags|log.Lshortfile)
	slice := make([]string, 0, len(params))
	for k, v := range params {
		slice = append(slice, strings.TrimSpace(k+"="+v.(string)))
	}
	param := strings.Join(slice, "&")
	url := host + "?" + param
	logger.Printf("geting accesstoken url is : %s", url)
	request, _ := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		request.Header.Set(k, v)
	}

	return request
}

func RequestPost(host string, params map[string]string, headers map[string]string) *http.Request {
	logger := log.New(os.Stdout, "[RequestGet]", log.LstdFlags|log.Lshortfile)

	data, _ := json.Marshal(params)

	logger.Printf("geting accesstoken url is : %s \n %s", host, string(data))
	request, _ := http.NewRequest("POST", host, bytes.NewBuffer(data))
	for k, v := range headers {
		request.Header.Set(k, v)
	}

	return request
}
