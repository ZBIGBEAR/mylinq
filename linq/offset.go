package linq

func (q Query) Offset(offset int64) Query {
	return Query{
		Iterate: func() Iterator {
			next := q.Iterate()
			var index int64 = 0
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if index >= offset {
						return item, ok
					}
					index ++
				}
				return nil, false
			}
		},
	}
}