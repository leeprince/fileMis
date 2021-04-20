package controller

// 根据Views的配置信息返回对应的方法及描述
func ViewsFormate(view [][3]string) ([]string, []string) {
	var nextMethod []string = make([]string, len(view))
	var nextMethodDesc []string = make([]string, len(view))
	
	for k, v := range view {
		nextMethod[k] = v[0] + "::" + v[1]
		nextMethodDesc[k] = v[2]
	}
	return nextMethod, nextMethodDesc
}
