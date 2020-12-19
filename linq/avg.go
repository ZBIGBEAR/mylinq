package linq

func (q Query) Avg(avg func(i interface{})float64) float64{
	convert := func(c interface{}) interface{} {
		return avg(c)
	}
	convertQ:= q.Convert(convert)
	var sum float64 = 0
	next := convertQ.Iterator()
	var count float64 = 0
	for item, ok := next(); ok; item, ok = next() {
		sum += item.(float64)
		count ++
	}
	return float64(sum*1.0/count)
}