package main

import "fmt"

func main() {
	//var m1 map[string]int
	//fmt.Println(m1 == nil)
	//m1 = make(map[string]int, 10)
	//m1["理"] = 18
	//m1["wsb"] = 35
	//m1["lqy"] = 19
	//fmt.Println(m1)
	//fmt.Println(m1["wsb"])
	//value, ok := m1["wsb"]
	//if !ok {
	//	fmt.Println("NO KEY")
	//} else {
	//	fmt.Println(value)
	//}
	//for k, v := range m1 {
	//	fmt.Println(k, v)
	//}
	//
	//delete(m1, "理")
	//fmt.Println(m1)
	//var scoreMap = make(map[string]int, 200)
	//for i := 0; i < 100; i++ {
	//	key := fmt.Sprintf("stu%02d", i)
	//	value := rand.Intn(100)
	//	scoreMap[key] = value
	//}
	//var keys = make([]string, 0, 200)
	//for key := range scoreMap {
	//	keys = append(keys, key)
	//}
	//sort.Strings(keys)
	//for _, key := range keys {
	//	fmt.Println(key, scoreMap[key])
	//}

	//var s1 = make([]map[int]string, 10, 10)
	//s1[0] = make(map[int]string, 1)
	//s1[0][10] = "wsb"
	//fmt.Println(s1)
	//
	//var m2 = make(map[string][]int, 10)
	//m2["aaa"] = []int{10, 20, 30}
	//fmt.Println(m2)

	var MyMap = make(map[string]int, 10)
	saveS := ""
	s := "how do you do"
	for i := 0; i < len(s); i++ {
		a := fmt.Sprintf("%c", s[i])
		if a == " " {
			MyMap[saveS] += 1
			saveS = ""
			continue
		} else {
			saveS += a
		}
	}
	MyMap[saveS] += 1
	fmt.Println(MyMap)
}
