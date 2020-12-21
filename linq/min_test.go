package linq

import (
	"github.com/stretchr/testify/suite"
	"math"
	"testing"
)

type MinSuite struct{
	suite.Suite
}

func TestMinSuite(t *testing.T){
	suite.Run(t, new(MinSuite))
}

func (f *MinSuite) TestMin(){
	result := From(students).Min(func(i interface{}) float64 {
		s,ok := i.(Person)
		if !ok{
			return math.MaxFloat64
		}
		return float64(s.Age)
	})
	f.Equal(float64(10), result)
}
