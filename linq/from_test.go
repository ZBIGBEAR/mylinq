package linq

import (
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
	printStudents("Test From", result)
}
