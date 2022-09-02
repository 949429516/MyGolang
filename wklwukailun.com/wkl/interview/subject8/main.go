package main

import (
	"encoding/json"
	"fmt"
)

/*
修改代码，使status输出为200

type Result struct {
	status int
}

func main() {
	var data = []byte(`{"status":200}`)
	result := &Result{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Printf("result=%+v\n", result)
}
*/

type Result struct {
	Status int
}

func main() {
	var data = []byte(`{"status":200}`)
	result := &Result{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Printf("result=%+v\n", result.Status)
}
