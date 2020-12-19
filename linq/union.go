package linq

func (q Query) Union(unionQuery Query) Query {
	return Query{
		Iterator: func() func() (interface{}, bool) {
			next1 := q.Iterator()
			next2 := unionQuery.Iterator()
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