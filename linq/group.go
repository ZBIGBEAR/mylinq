package linq

type GroupData struct {
	Datas []interface{}
	Count int64
	key interface{}
}
func (q Query) Group(group func(interface{})interface{}) []GroupData {
	set := make(map[interface{}]GroupData)
	next := q.Iterate()
	for item,ok := next(); ok; item,ok = next(){
		k := group(item)
		if _, ok := set[k]; !ok {
			set[k] = GroupData{
				Datas: []interface{}{item},
				Count: 1,
				key: k,
			}
		}else{
			tmp := set[k]
			tmp.Count+=1
			tmp.Datas = append(tmp.Datas,item)
			set[k]=tmp
		}
	}
	var datas []GroupData
	for _,val:=range set{
		datas = append(datas,val)
	}
	return datas
}