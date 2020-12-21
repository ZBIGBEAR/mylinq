package linq

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GroupSuite struct{
	suite.Suite
}

func TestGroupSuite(t *testing.T){
	suite.Run(t, new(GroupSuite))
}

func (f *GroupSuite) TestGroup1(){
	result := From(students).Group(func(i interface{}) interface{} {
		s,ok:=i.(Person)
		if !ok{
			return i
		}
		return s.Grade
	})
	fmt.Println(result)
	f.Equal(3, len(result))
	f.Equal(int64(4), result[0].Count)
	f.Equal(int64(3), result[1].Count)
	f.Equal(int64(3), result[2].Count)
}

func (f *GroupSuite) TestGroup2(){
	result := From(students).Group(func(i interface{}) interface{} {
		s,ok:=i.(Person)
		if !ok{
			return i
		}
		return s.Gender
	})
	fmt.Println(result)
	f.Equal(2, len(result))
}
