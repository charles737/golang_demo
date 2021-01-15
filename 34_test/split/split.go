package split

import "strings"

// split.go

// 将 s 按照 sep 进行切割，返回一个字符串的切片
// Split("我爱你", "爱") => ["我", "你"]
func Split(s, sep string) (ret []string) {
	// 	ret=make 这里是一次性能优化，提前做一次内存初始化，将内存申请由原来的每次操作的3次降为了1次
	ret = make([]string, 0, strings.Count(s, sep)+1)
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx+len(sep):] // len(sep) 根据sep的长度实现切割，同时满足中文和英文的需求
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
