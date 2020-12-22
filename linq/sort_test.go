package linq

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SortSuite struct{
	suite.Suite
}

func TestSortSuite(t *testing.T){
	suite.Run(t, new(SortSuite))
}

func (f *SortSuite) TestSort1(){
	result := From(students).Sort(func(i,j interface{}) bool {
		s1,ok := i.(Person)
		if !ok {
			return false
		}
		s2,ok := j.(Person)
		if !ok {
			return false
		}
		return s1.Age<s2.Age
	}).Result()
	f.Equal(10, len(result))
	for i:=0;i<len(result);i++{
		f.Equal(fmt.Sprintf("student%d",i+1), result[i].(Person).Name)
	}
}

func (f *SortSuite) TestSort2(){
	result := From(students).Sort(func(i,j interface{}) bool {
		s1,ok := i.(Person)
		if !ok {
			return false
		}
		s2,ok := j.(Person)
		if !ok {
			return false
		}
		return s1.Age>s2.Age
	}).Result()
	f.Equal(10, len(result))
	for i:=0;i<len(result);i++{
		f.Equal(fmt.Sprintf("student%d",10-i), result[i].(Person).Name)
	}
}