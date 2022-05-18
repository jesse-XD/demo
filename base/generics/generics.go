package generics

type Number interface {
	int64 | float64
}

// 使用泛型判断大小 [T 类型约束]
func GMin[T Number](a, b T) T {
	if a > b {
		return b
	}
	return a
}

// 求和
func GSum[k comparable, v Number](m map[k]v) v {
	var s v
	for _, val := range m {
		s += val
	}
	return s
}
