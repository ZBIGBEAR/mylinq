package linq

func (q Query)Order(orderBy func(interface{})float64, asc bool) Query {
	return Query{
		Iterate: func() Iterator {
			datas := q.order(orderBy, asc)
			idx :=0
			count := len(datas)
			return func() (item interface{}, ok bool) {
				if idx<count{
					item = datas[idx]
					idx++
					return item,true
				}else{
					return nil,false
				}
			}
		},
	}
}

func (q Query)order(orderBy func (interface{})float64, asc bool) []interface{}{
	datas := q.Result()
	count := len(datas)
	// 冒泡排序
	changeIndex := 0
	for i := 0;i<count;i++{
		changeIndex = count-1
		for j:=count-1;j>=i;j--{
			if asc {
				if orderBy(datas[j]) < orderBy(datas[changeIndex]) {
					changeIndex = j
				}
			}else{
				if orderBy(datas[j]) > orderBy(datas[changeIndex]) {
					changeIndex = j
				}
			}
		}
		datas[i],datas[changeIndex]=datas[changeIndex],datas[i]
	}
	return datas
}