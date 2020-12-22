package linq

func (q Query) Where(where func(interface{})bool) Query {
	return Query{
		Iterate: func() Iterator {
			next := q.Iterate()
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
