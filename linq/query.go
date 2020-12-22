package linq

type Iterator func()(interface{}, bool)

type Query struct {
	Iterate func() Iterator
}
