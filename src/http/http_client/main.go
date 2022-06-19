package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "王色宝")
	data.Set("age", "38")
	queryStr := data.Encode() //encode之后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, _ := http.NewRequest("GET", urlObj.String(), nil)
	//req.Header.Add() //请求头增加参数
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("read resp.Body failed, err:", err)
		return
	}
	//禁用keepAlive的client
	/*
		tr := &http.Transport{
			DisableKeepAlives: true,
		}
		client := http.Client{
			Transport: tr,
		}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("read resp.Body failed, err:", err)
			return
		}
	*/
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body failed, err:", err)
		return
	}
	fmt.Println(string(b))
}
