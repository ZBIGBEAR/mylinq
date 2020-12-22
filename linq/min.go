package linq

import "math"

func (q Query) Min(MinBy func(interface{}) float64 ) float64 {
	m := math.MaxFloat64
	next := q.Iterate()
	for item,ok:=next();ok;item,ok=next(){
		if m > MinBy(item){
			m = MinBy(item)
		}
	}
	return m
}