package http

import (
	"testing"
)

func TestGet(t *testing.T) {
	url := "http://php10.cn/req.php?a=1"
	content, err := HttpGet(url)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(content)
}

func TestHttpPost(t *testing.T) {
	url := "http://php10.cn/req.php"
	type data struct {
		Name string `json:"name"`
	}
	content, err := HttpPost(url, data{
		Name: "jesse",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(content)
}

func TestHttpDo(t *testing.T) {
	methode := "POST"
	url := "http://php10.cn/req.php"
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Cookie"] = "token=1234"
	type data struct {
		Name string `json:"name"`
	}
	content, err := HttpDo(methode, url, header, data{
		Name: "jesse",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(content)
}
