package linq

func (q Query)Intersect(intersectQuery Query, intersectBy func(interface{})interface{}) Query {
	return Query{
		Iterator: func() func() (interface{}, bool) {
			next2 := intersectQuery.Iterator()
			set := make(map[interface{}]struct{})
			for item,ok:=next2();ok;item,ok=next2(){
				k := item
				if intersectBy != nil {
					k = intersectBy(item)
				}
				set[k] = struct{}{}
			}

			next := q.Iterator()
			return func() (i interface{}, b bool) {
				for item,ok:=next();ok;item,ok=next(){
					k := item
					if intersectBy != nil {
						k = intersectBy(item)
					}
					if _, ok := set[k];ok{
						return item,true
					}
				}
				return nil,false
			}
		},
	}
}