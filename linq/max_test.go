package linq

import (
	"github.com/stretchr/testify/suite"
	"math"
	"testing"
)

type MaxSuite struct{
	suite.Suite
}

func TestMaxSuite(t *testing.T){
	suite.Run(t, new(MaxSuite))
}

func (f *MaxSuite) TestMax(){
	result := From(students).Max(func(i interface{}) float64 {
		s,ok := i.(Person)
		if !ok{
			return math.SmallestNonzeroFloat64
		}
		return float64(s.Age)
	})
	f.Equal(float64(19), result)
}
