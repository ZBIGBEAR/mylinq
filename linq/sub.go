package linq

func (q Query)Sub(subQuery Query, subBy func(interface{})interface{}) Query {
	return Query{
		Iterate: func() Iterator {
			next2 := subQuery.Iterate()
			set := make(map[interface{}]struct{})
			for item,ok:=next2();ok;item,ok=next2(){
				k := item
				if subBy != nil {
					k = subBy(item)
				}
				set[k] = struct{}{}
			}

			next := q.Iterate()
			return func() (i interface{}, b bool) {
				for item,ok:=next();ok;item,ok=next(){
					k := item
					if subBy != nil {
						k = subBy(item)
					}
					if _, ok := set[k];!ok{
						return item,true
					}
				}
				return nil,false
			}
		},
	}
}