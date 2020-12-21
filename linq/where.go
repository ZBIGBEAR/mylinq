package linq

func (q Query) Where(where func(interface{})bool) Query {
	return Query{
		Iterator: func() func() (interface{}, bool) {
			next := q.Iterator()
			return func() (item interface{}, ok bool) {
				for item,ok=next();ok;item,ok=next(){
					if where(item){
						return item,true
					}
				}
				return nil,false
			}
		},
	}
}
