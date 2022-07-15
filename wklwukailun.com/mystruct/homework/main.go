package main

import (
	"fmt"
	"os"
)

var smr studentMgr

func showMenu() {
	fmt.Println("----欢迎光临学生管理系统----")
	fmt.Println(`
		1.查看所有学生
		2.增加学生
		3.修改学生
		4.删除学生
		5.退出
		`)
}

func main() {
	smr = studentMgr{
		allstudent: make(map[int64]student, 50),
	}
	for {
		showMenu()
		var choise int
		fmt.Print("请输入需要的操作序号:")
		fmt.Scanln(&choise)
		switch choise {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudents()
		case 3:
			smr.editStudents()
		case 4:
			smr.deleteStudents()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入有误,请重新输入!!!")
		}
	}
}
