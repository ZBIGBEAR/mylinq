package linq

func (q Query) Convert(convert func(interface{})interface{}) Query{
	return Query{
		Iterator: func() func() (interface{}, bool) {
			next := q.Iterator()
			return func() (item interface{}, ok bool) {
				item, ok = next()
				if !ok {
					return nil, false
				}
				return convert(item),true
			}
		},
	}
}