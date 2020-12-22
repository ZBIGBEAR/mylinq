package linq

import "math"

func (q Query) Max(maxBy func(interface{}) float64 ) float64 {
	m := math.SmallestNonzeroFloat64
	next := q.Iterate()
	for item,ok:=next();ok;item,ok=next(){
		if m<maxBy(item){
			m = maxBy(item)
		}
	}
	return m
}