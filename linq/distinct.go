package linq

func (q Query) Distnct(dist func(interface{})interface{})Query{
	return Query{
		Iterator: func() func() (interface{}, bool) {
			set := make(map[interface{}]struct{})
			next := q.Iterator()
			return func() (item interface{}, ok bool) {
				for item,ok=next(); ok; item, ok = next(){
					if _,ok1:=set[dist(item)]; !ok1{
						set[dist(item)] = struct{}{}
						return item,true
					}
				}
				return nil,false
			}
		},
	}
}