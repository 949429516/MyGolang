package main

import "fmt"

type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allstudent map[int64]student
}

func (self studentMgr) showStudents() {
	for _, v := range self.allstudent {
		fmt.Printf("学号:%d,姓名:%s\n", v.id, v.name)
	}
}

func (self studentMgr) addStudents() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	newStudent := student{
		id:   id,
		name: name,
	}
	self.allstudent[id] = newStudent
}

func (self studentMgr) editStudents() {
	var (
		stuId   int64
		newName string
	)
	fmt.Print("请输入需要修改学生的学号:")
	fmt.Scanln(&stuId)
	stuObj, ok := self.allstudent[stuId]
	if ok {
		fmt.Print("请输入需要修改学生的姓名:")
		fmt.Scanln(&newName)
		stuObj.name = newName
		self.allstudent[stuId] = stuObj
	} else {
		fmt.Println("查无此人!")
		return
	}
}

func (self studentMgr) deleteStudents() {
	var id int64
	fmt.Print("请输入需要删除学生的学号:")
	fmt.Scanln(&id)
	_, ok := self.allstudent[id]
	if ok {
		delete(self.allstudent, id)
	} else {
		fmt.Println("查无此人!")
	}

}
