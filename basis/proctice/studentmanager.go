package main

import "fmt"

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

// func (s *student) editInfo(key string, value interface{}) {
// 	rs := reflect.ValueOf(s).Elem()
// 	fmt.Println(rs)
// 	rs.FieldByName(key).Set(reflect.ValueOf(value))
// }

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
}
