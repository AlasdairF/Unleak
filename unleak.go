package unleak

func Bytes(b []byte) []byte {
	v := make([]byte, len(b))
	copy(v, b)
	return v
}

func Runes(b []rune) []rune {
	v := make([]rune, len(b))
	copy(v, b)
	return v
}

func String(s string) string {
	b := []byte(s)
	v := make([]byte, len(b))
	copy(v, b)
	return string(v)
}

func Ints(b []int) []int {
	v := make([]int, len(b))
	copy(v, b)
	return v
}

func Int8s(b []int8) []int8 {
	v := make([]int8, len(b))
	copy(v, b)
	return v
}

func Int16s(b []int16) []int16 {
	v := make([]int16, len(b))
	copy(v, b)
	return v
}

func Int32s(b []int32) []int32 {
	v := make([]int32, len(b))
	copy(v, b)
	return v
}

func Int64s(b []int64) []int64 {
	v := make([]int64, len(b))
	copy(v, b)
	return v
}

func Uints(b []uint) []uint {
	v := make([]uint, len(b))
	copy(v, b)
	return v
}

func Uint8s(b []uint8) []uint8 {
	v := make([]uint8, len(b))
	copy(v, b)
	return v
}

func Uint16s(b []uint16) []uint16 {
	v := make([]uint16, len(b))
	copy(v, b)
	return v
}

func Uint32s(b []uint32) []uint32 {
	v := make([]uint32, len(b))
	copy(v, b)
	return v
}

func Uint64s(b []uint64) []uint64 {
	v := make([]uint64, len(b))
	copy(v, b)
	return v
}

func Float32s(b []float32) []float32 {
	v := make([]float32, len(b))
	copy(v, b)
	return v
}

func Float64s(b []float64) []float64 {
	v := make([]float64, len(b))
	copy(v, b)
	return v
}

func Bools(b []bool) []bool {
	v := make([]bool, len(b))
	copy(v, b)
	return v
}
