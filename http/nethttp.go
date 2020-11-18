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
	ct := "application/x-www-form-urlencoded"
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
