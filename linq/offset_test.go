package linq

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type OffsetSuite struct{
	suite.Suite
}

func TestOffsetSuite(t *testing.T){
	suite.Run(t, new(OffsetSuite))
}

func (f *OffsetSuite) TestOffset1(){
	result := From(students).Offset(0).Result()
	f.Equal(10, len(result))
	for i:=0;i<len(result);i++{
		f.Equal(fmt.Sprintf("student%d",i+1), result[i].(Person).Name)
	}
}

func (f *OffsetSuite) TestOffset2(){
	result := From(students).Offset(8).Result()
	f.Equal(2, len(result))
	for i:=0;i<len(result);i++{
		f.Equal(fmt.Sprintf("student%d",i+9), result[i].(Person).Name)
	}
}

func (f *OffsetSuite) TestOffset3(){
	result := From(students).Offset(10).Result()
	f.Equal(0, len(result))
}
