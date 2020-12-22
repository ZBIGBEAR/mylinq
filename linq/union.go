package linq

func (q Query) Union(unionQuery Query) Query {
	return Query{
		Iterate: func() Iterator {
			next1 := q.Iterate()
			next2 := unionQuery.Iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next1(); ok; item, ok = next1() {
					return item, ok
				}
				for item, ok = next2(); ok; item, ok = next2() {
					return item, ok
				}
				return nil, false
			}
		},
	}
}