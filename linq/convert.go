package linq

func (q Query) Convert(convert func(interface{})interface{}) Query{
	return Query{
		Iterate: func() Iterator {
			next := q.Iterate()
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