package linq

func (q Query) Limit(limit int64) Query {
	return Query{
		Iterator: func() func() (interface{}, bool) {
			var index int64 = 0
			next := q.Iterator()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if index < limit {
						index++
						return item, ok
					}
					index++
				}
				return nil, false
			}
		},
	}
}