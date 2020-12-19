package linq

func (q Query) Sum(sumFunc func(i interface{})float64) float64{
	convert := func(c interface{}) interface{} {
		return sumFunc(c)
	}
	convertQ:= q.Convert(convert)
	var sum float64 = 0
	next := convertQ.Iterator()
	for item, ok := next(); ok; item, ok = next() {
		sum += item.(float64)
	}
	return sum
}