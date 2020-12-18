package linq

type Query struct {
	Iterator func() func()(interface{},bool)
}
