package linq

import "sort"

func (q Query) Sort(less func(i,j interface{})bool) Query{
	return Query{
		Iterate: func() Iterator {
			s := sortDatas{
				datas:q.Result(),
				less:less,
			}
			sort.Sort(s)
			idx := 0
			count := len(s.datas)
			return func() (item interface{}, ok bool) {
				if idx>=count{
					return nil,false
				}
				item = s.datas[idx]
				idx++
				return item,true
			}
		},
	}
}

type sortDatas struct {
	datas []interface{}
	less func(i,j interface{}) bool
}

func (s sortDatas) Len() int{
	return len(s.datas)
}

func (s sortDatas) Swap(i,j int){
	s.datas[i],s.datas[j]=s.datas[j],s.datas[i]
}

func (s sortDatas) Less(i,j int) bool {
	return s.less(s.datas[i], s.datas[j])
}