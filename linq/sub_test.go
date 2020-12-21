package linq

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SubSuite struct{
	suite.Suite
}

func TestSubSuite(t *testing.T){
	suite.Run(t, new(SubSuite))
}

func (f *SubSuite) TestSub1(){
	result := From(students).Sub(From(students[3:6]), nil).Result()
	f.Equal(7, len(result))
	for i:=0;i<3;i++{
		f.Equal(fmt.Sprintf("student%d",i+1), result[i].(Person).Name)
	}
	for i:=3;i<7;i++{
		f.Equal(fmt.Sprintf("student%d",i+4), result[i].(Person).Name)
	}
}

func (f *SubSuite) TestSub2(){
	// 减去一班、二班两个老师的学生
	result := From(students).Sub(From(teachers[:2]), func(i interface{}) interface{} {
		s, ok := i.(Person)
		if !ok {
			return i
		}
		return s.Grade
	}).Result()
	fmt.Println(result)
	f.Equal(3, len(result))
	f.Equal("student3", result[0].(Person).Name)
	f.Equal("student6", result[1].(Person).Name)
	f.Equal("student9", result[2].(Person).Name)
}