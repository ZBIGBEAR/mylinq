package linq

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UnionSuite struct{
	suite.Suite
}

func TestUnionSuite(t *testing.T){
	suite.Run(t, new(UnionSuite))
}

func (f *UnionSuite) TestUnion(){
	result := From(students).Union(From(teachers)).Result()
	f.Equal(13, len(result))
	for i:=0;i<len(students);i++{
		f.Equal(students[i].Name, result[i].(Person).Name)
	}
	for i:=0;i<len(teachers);i++{
		f.Equal(teachers[i].Name, result[i+StudentCount].(Person).Name)
	}
}
