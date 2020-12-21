package linq

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type WhereSuite struct{
	suite.Suite
}

func TestWhereSuite(t *testing.T){
	suite.Run(t, new(WhereSuite))
}

func (f *WhereSuite) TestWhere1(){
	result := From(students).Where(func(i interface{}) bool {
		s, ok:=i.(Person)
		if !ok{
			return false
		}
		return s.Gender == "男"
	}).Result()
	result1 := From(students).Where(func(i interface{}) bool {
		s, ok:=i.(Person)
		if !ok{
			return false
		}
		return s.Gender == "女"
	}).Result()
	f.Equal(10, len(result)+len(result1))
}

func (f *WhereSuite) TestWhere2(){
	result := From(students).Where(func(i interface{}) bool {
		s, ok:=i.(Person)
		if !ok{
			return false
		}
		return s.Age<15
	}).Result()
	f.Equal(5, len(result))
	for idx := range result {
		f.Equal(fmt.Sprintf("student%d",idx+1), result[idx].(Person).Name)
	}
}