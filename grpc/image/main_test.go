package main

import (
	"bytes"
	"fmt"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"net/http"
	"newgitlab.com/blockchain/go-ms-toolkit/encoding/xjson"
	"testing"
)

func TestURL(t *testing.T) {
	// https://gimg2.baidu.com/image_search/src=http%3A%2F%2Ffile02.16sucai.com%2Fd%2Ffile%2F2014%2F0829%2Fb871e1addf5f8e96f3b390ece2b2da0d.jpg&refer=http%3A%2F%2Ffile02.16sucai.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1640495560&t=64b384901ff1c58af87776a342b11dab
	oldURl := "http://ppiaas-starpool.oss-cn-shenzhen.aliyuncs.com/group1/default/20210929/10/39/0/a510b4f317a99deb10f03ccac34ea249"
	newURL := "http://10.8.0.16:8083/group1/upload"
	// http://pro-fastdfs.default.svc.cluster.local:8080;
	resp111, err := http.Get(oldURl)
	if err != nil {
		return
	}
	data, _ := ioutil.ReadAll(resp111.Body)
	cli := resty.New()
	resp1, err := cli.R().SetFileReader("file", "file.jpg", bytes.NewReader(data)).SetFormData(map[string]string{
		"auth_token": xjson.StringifyJson(map[string]interface{}{
			"id":    329,
			"token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzgyNTczMDQsInJpZCI6MiwidWlkIjozMjl9.W-lBGNzR_UxwS6NojahnGMcywZjXibXDMP5MmHGGIUl6Jbe7lw37_BGFoa2IBxuyJJ4IhNt_Ms1avrPndqzJb5dWqPnGg6KOV96qhfN4XhJIiOReyXTmo3TmgYtV-yHLnRDpqu7Ey7TGBXI5ffEVOCJq_MoquTrvMYLxYG9lbxI24NH-GUJSPq7a-WKryEzoC4H8Ioh6ChXfBpR9ETlua6ErN3ta0Q1K3AhgIRHAX97u_HDz7TMnFNQgXIGuYVtyXaMbhXedy-PiR6XDYwih9Y6OlXqBnUni9WRw73NQQSMsgVAokeXbFwd2ah48vIOaulTTi8H-a_KVtVxbdCjN9Q",
		}),
	}).Post(newURL)
	if err != nil {
		fmt.Println("nonono")
		return
	}
	fmt.Println(string(resp1.Body()))
}
