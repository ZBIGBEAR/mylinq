package linq

func (q Query) Result() []interface{}{
	var datas []interface{}
	next := q.Iterate()
	for item,ok:=next();ok;item,ok=next(){
		datas=append(datas,item)
	}
	return datas
}
