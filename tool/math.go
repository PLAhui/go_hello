package tool

type Number interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

func SumInt[T Number](a, b T) T {
	return a + b
}

func Equal[T Number](a, b T) bool {
	return a == b
}
