package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println(os.Args[1])
	fmt.Printf("%T\n", os.Args) //[]string
}
