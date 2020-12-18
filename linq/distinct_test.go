package linq

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type DistinctSuite struct{
	suite.Suite
}

func TestDistinctSuite(t *testing.T){
	suite.Run(t, new(DistinctSuite))
}

func (d *DistinctSuite) TestDistinct1(){
	result := From(students).Distnct(func(i interface{}) interface{} {
		p,ok := i.(Person)
		if !ok{
			return i
		}
		return p.Teacher.Name
	}).Result()
	d.Equal(3, len(result))
	d.Equal("student1", result[0].(Person).Name)
	d.Equal("student2", result[1].(Person).Name)
	d.Equal("student3", result[2].(Person).Name)
}


func (d *DistinctSuite) TestDistinct2(){
	result := From(students).Distnct(func(i interface{}) interface{} {
		p,ok := i.(Person)
		if !ok{
			return i
		}
		return p.Gender
	}).Result()
	d.Equal(2, len(result))
	d.Equal("student1", result[0].(Person).Name)
	d.Equal("student2", result[1].(Person).Name)
}