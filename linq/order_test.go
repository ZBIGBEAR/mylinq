package linq

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type OrderSuite struct{
	suite.Suite
}

func TestOrderSuite(t *testing.T){
	suite.Run(t, new(OrderSuite))
}

func (f *OrderSuite) TestOrder1(){
	result := From(students).Order(func(i interface{}) float64 {
		s,ok := i.(Person)
		if !ok {
			return 0
		}
		return float64(s.Age)
	}, true).Result()
	f.Equal(10, len(result))
	for i:=0;i<len(result);i++{
		f.Equal(fmt.Sprintf("student%d",i+1), result[i].(Person).Name)
	}
}

func (f *OrderSuite) TestOrder2(){
	result := From(students).Order(func(i interface{}) float64 {
		s,ok := i.(Person)
		if !ok {
			return 0
		}
		return float64(s.Age)
	}, false).Result()
	f.Equal(10, len(result))
	for i:=0;i<len(result);i++{
		f.Equal(fmt.Sprintf("student%d",10-i), result[i].(Person).Name)
	}
}