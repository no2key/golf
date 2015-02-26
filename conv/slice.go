// slice
package conv

// 初始化一个长度为 size 的 []interface{} 切片
func NewInterfacePtrs(size int) []interface{} {
	buf := make([]interface{}, size)
	rs := make([]interface{}, size)

	for i := range buf {
		rs[i] = &buf[i]
	}

	return rs
}
