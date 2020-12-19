package linq

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SumSuite struct{
	suite.Suite
}

func TestSumSuite(t *testing.T){
	suite.Run(t, new(SumSuite))
}

func (f *SumSuite) TestSum(){
	result := From(students).Sum(func(i interface{}) float64 {
		s,ok := i.(Person)
		if !ok {
			return 0
		}
		return float64(s.Age)
	})
	f.Equal(float64(145), result)
}
