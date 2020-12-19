package linq

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LimitSuite struct{
	suite.Suite
}

func TestLimitSuite(t *testing.T){
	suite.Run(t, new(LimitSuite))
}

func (f *LimitSuite) TestLimit1(){
	result := From(students).Limit(0).Result()
	f.Equal(0, len(result))
}

func (f *LimitSuite) TestLimit2(){
	result := From(students).Limit(2).Result()
	f.Equal(2, len(result))
	for i:=0;i<len(result);i++{
		f.Equal(fmt.Sprintf("student%d",i+1), result[i].(Person).Name)
	}
}

func (f *LimitSuite) TestLimit3(){
	result := From(students).Limit(10).Result()
	f.Equal(10, len(result))
	for i:=0;i<len(result);i++{
		f.Equal(fmt.Sprintf("student%d",i+1), result[i].(Person).Name)
	}
}

func (f *LimitSuite) TestLimit4(){
	result := From(students).Offset(2).Limit(3).Result()
	f.Equal(3, len(result))
	for i:=0;i<len(result);i++{
		f.Equal(fmt.Sprintf("student%d",i+3), result[i].(Person).Name)
	}
}

func (f *LimitSuite) TestLimit5(){
	result := From(students).Limit(3).Offset(2).Result()
	f.Equal(1, len(result))
	f.Equal("student3", result[0].(Person).Name)
}

