package ini

func Filter(list interface{}, fn func(interface{}) bool) []interface{} {
	listX := list.([]interface{})
	var ret []interface{}
	for _, v := range listX {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
