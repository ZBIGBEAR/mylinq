package linq

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type IntersectSuite struct{
	suite.Suite
}

func TestIntersectSuite(t *testing.T){
	suite.Run(t, new(IntersectSuite))
}

func (f *IntersectSuite) TestIntersect1(){
	result := From(students).Intersect(From(students[3:6]), nil).Result()
	f.Equal(3, len(result))
	for i:=0;i<3;i++{
		f.Equal(fmt.Sprintf("student%d",i+4), result[i].(Person).Name)
	}
}

func (f *IntersectSuite) TestIntersect2(){
	// 求一班、二班两个老师的学生
	result := From(students).Intersect(From(teachers[:2]), func(i interface{}) interface{} {
		s, ok := i.(Person)
		if !ok {
			return i
		}
		return s.Grade
	}).Result()
	f.Equal(7, len(result))
	f.Equal("student1", result[0].(Person).Name)
	f.Equal("student2", result[1].(Person).Name)
	f.Equal("student4", result[2].(Person).Name)
	f.Equal("student5", result[3].(Person).Name)
	f.Equal("student7", result[4].(Person).Name)
	f.Equal("student8", result[5].(Person).Name)
	f.Equal("student10", result[6].(Person).Name)
}