package linq

func (q Query) Limit(limit int64) Query {
	return Query{
		Iterate: func() Iterator {
			var index int64 = 0
			next := q.Iterate()
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