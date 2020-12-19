package linq

import (
	"fmt"
)

type Person struct {
	Name string
	Age int64
	Gender string
	Addr string
	Tel string
	Grade string
	Class string
	Teacher *Person
}

const (
	StudentCount = 10
)

var teachers []Person
var students []Person

func init(){
	teachers = append(teachers, Person{
		Name:    "张老师",
		Age:     21,
		Gender:"男",
		Grade:"一班",
	})
	teachers = append(teachers, Person{
		Name:    "李老师",
		Age:     22,
		Gender:"女",
		Grade:"二班",
	})
	teachers = append(teachers, Person{
		Name:    "王老师",
		Age:     23,
		Grade:"三班",
	})

	genders :=[]string{"男","女"}
	for i:=0;i<StudentCount;i++{
		teacher := teachers[i%3]
		students = append(students, Person{
			Name:    fmt.Sprintf("student%d", i+1),
			Age:     int64(i + 10),
			Gender:genders[i%2],
			Addr:    "",
			Tel:     "",
			Grade:   teacher.Grade,
			Class:   "",
			Teacher: &teacher,
		})
	}
}