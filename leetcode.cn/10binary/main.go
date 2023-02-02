package main

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) (s string) {
	len_a, len_b := len(a), len(b)
	if len_a >= len_b {
		b = strings.Repeat("0", len_a-len_b) + b
	} else {
		a = strings.Repeat("0", len_b-len_a) + a
	}
	var flag bool
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == b[i] && a[i] == '0' {
			if flag {
				s = "1" + s
				flag = false
			} else {
				s = "0" + s
			}
		} else if a[i] == b[i] && a[i] == '1' {
			if flag {
				s = "1" + s
			} else {
				s = "0" + s
				flag = true
			}
		} else {
			if flag {
				s = "0" + s
			} else {
				s = "1" + s
			}
		}
	}
	if flag {
		s = "1" + s
	}
	return
}

func main() {
	fmt.Println(addBinary("1010", "1011"))

}

//   1 1
//   0 1
// 1 0 0
