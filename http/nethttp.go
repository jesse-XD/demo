package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送 get 请求
func HttpGet(url string) (string, error) {
	// 设置请求超时时间，0表示没有超时限制
	client := &http.Client{Timeout: 10 * time.Second}
	// 发送请求
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	// 必须关闭响应主体
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// 发送 post 请求
func HttpPost(url string, req interface{}) (string, error) {
	var body []byte
	body, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	client := &http.Client{Timeout: 10 * time.Second}
	// 发送post请求
	ct := "application/json"
	resp, err := client.Post(url, ct, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// 自定义的请求方式，可以设置头参数
func HttpDo(method string, url string, header map[string]string, req interface{}) (string, error) {
	var body []byte
	body, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	client := &http.Client{Timeout: 10 * time.Second}
	nReq, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	// 设置请求头
	for k, v := range header{
		nReq.Header.Set(k, v)
	}
	// 发送请求
	resp, err := client.Do(nReq)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(content), nil
}
