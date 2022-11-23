package cache

type ByteView struct {
	b []byte
}

// Len 返回长度
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice 返回字节切片
func (v ByteView) ByteSlice() []byte {
	return cloneByte(v.b)
}

// String 字节转字符
func (v ByteView) String() string {
	return string(v.b)
}

// cloneByte 复制字节
func cloneByte(b []byte) []byte {
	r := make([]byte, len(b))
	copy(b, r)
	return r
}
