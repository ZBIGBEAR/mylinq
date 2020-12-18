package linq

import "reflect"

func From(datas interface{}) Query{
	src := reflect.ValueOf(datas)
	switch src.Kind(){
	case reflect.Array,reflect.Slice:
		return Query{
			Iterator: func() func()(interface{}, bool) {
				idx :=0
				count := src.Len()
				return func()(item interface{}, ok bool) {
					if idx>=count{
						return nil,false
					}
					item = src.Index(idx).Interface()
					idx++
					return item,true
				}
			},
		}
	}
	return Query{}
}
