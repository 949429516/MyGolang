package main

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
请你实现这个将字符串进行指定行数变换的函数：
string convert(string s, int numRows);

示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
*/
import "fmt"

func convert(s string, numRows int) string {
	// 如果一行以内直接返回
	if numRows <= 1 {
		return s
	}
	// 创建map存储
	m := make(map[int]string, numRows)
	for key := 0; key < numRows; key++ {
		m[key] = ""
	}
	// 定义mkey是map的key，flag是例如向三个key中存储值0 1 2 1 0 1 2 1 0 1 2，加上去在减下来，在一个区间内到达极值则向另一个方向加(减)
	mkey, flag := 0, -1
	// 遍历字符串
	for _, irune := range s {
		str := string(irune)
		// 向不同key-value中添加
		m[mkey] += str
		// 当左极限为0或者右极限为长度时，则向反方向趋近
		if mkey == 0 || mkey == numRows-1 {
			flag = -flag
		}
		mkey += flag
	}
	finalystr := ""
	for key := 0; key < numRows; key++ {
		finalystr += m[key]
	}
	return finalystr
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
