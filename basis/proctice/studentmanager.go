package main

import (
	"fmt"
	"reflect"
)

type student struct {
	ID         int
	Name       string
	Age, Score int
}

type class struct {
	students []*student
}

func (c *class) addStudent(id int, name string, age, score int) {
	c.students = append(c.students, &student{
		id, 
		name,
		age,
		score,
	})
}

func (c *class) removeStudent(id int) {
	index := -1
	for k, v := range c.students {
		if v.ID == id {
			index = k
			break
		}
	}

	if index == -1 {
		fmt.Println("查无此人")
	} else {
		c.students = append(c.students[:index], c.students[index+1:]...)
	}
}

func (c *class) listStudent() {
	for _, v := range c.students {
		fmt.Printf("id:%d\nname:%s\nage:%d\nscore:%d\n\n", v.ID, v.Name, v.Age, v.Score)
	}
}

func (s *student) setName(val string) {
	reflect.ValueOf(s).Elem().FieldByName("Name").SetString(val)
}

func (s *student) setAge(val int) {
	reflect.ValueOf(s).Elem().FieldByName("Age").SetInt(int64(val))
}

func (s *student) setScore(val int) {
	reflect.ValueOf(s).Elem().FieldByName("Score").SetInt(int64(val))
}

func studentManagerTest() {
	c := class{students: make([]*student, 0, 200)}
	for i := 1; i < 20; i++ {
		stu := &student{
			i, 
			fmt.Sprintf("stu%02d", i),
			18,
			100,
		}
		c.students = append(c.students, stu)
	}
	fmt.Println("所有学生信息")
	c.listStudent()

	fmt.Println("增加一个学生")
	c.addStudent(20, "stu20", 18, 100)
	c.listStudent()

	fmt.Println("移除一个学生")
	c.removeStudent(20)
	c.listStudent()

	c.students[0].setName("test")
	fmt.Println(c.students[0])

	c.students[0].setAge(20)
	fmt.Println(c.students[0])

	c.students[0].setScore(150)
	fmt.Println(c.students[0])
}
