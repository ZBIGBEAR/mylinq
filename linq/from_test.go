package linq

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FromSuite struct{
	suite.Suite
}

func TestFromSuite(t *testing.T){
	suite.Run(t, new(FromSuite))
}

func (f *FromSuite) TestFrom(){
	result := From(students).Result()
	f.Equal(10, len(result))
	for i:=0;i<len(result);i++{
		f.Equal(fmt.Sprintf("student%d",i+1), result[i].(Person).Name)
	}

}
