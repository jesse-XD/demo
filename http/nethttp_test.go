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
		Name string
	}
	content, err := HttpPost(url, data{
		Name: "jesse",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(content)
}
