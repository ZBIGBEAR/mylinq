package linq

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConvertSuite struct{
	suite.Suite
}

func TestConvertSuite(t *testing.T){
	suite.Run(t, new(ConvertSuite))
}

func (f *ConvertSuite) TestConvert(){
	result := From(students).Convert(func(i interface{})interface{}{
		s,ok:=i.(Person)
		if !ok {
			return nil
		}
		s.Age *=2
		return s
	}).Result()
	for i:=0;i<len(result);i++{
		f.Equal(students[i].Age*2,result[i].(Person).Age)
	}

}
