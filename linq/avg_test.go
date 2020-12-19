package linq

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type AvgSuite struct{
	suite.Suite
}

func TestAvgSuite(t *testing.T){
	suite.Run(t, new(AvgSuite))
}

func (f *AvgSuite) TestAvg(){
	result := From(students).Avg(func(i interface{}) float64 {
		s,ok := i.(Person)
		if !ok {
			return 0
		}
		return float64(s.Age)
	})
	f.Equal(float64(14.5), result)
}
