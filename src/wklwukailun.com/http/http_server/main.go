package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write([]byte(b))
}
func f2(w http.ResponseWriter, r *http.Request) {
	//对于get请求，参数都在url上(query param),请求体中没有数据
	fmt.Println(r.URL)
	queryParam := r.URL.Query() //自动识别URL中的
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}
func main() {
	http.HandleFunc("/posts/Go/15_socket/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
