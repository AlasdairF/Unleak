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

func Ints(b []int) []int {
	v := make([]int, len(b))
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

func Floats(b []float64) []float64 {
	v := make([]float64, len(b))
	copy(v, b)
	return v
}