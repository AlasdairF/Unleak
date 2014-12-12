##Unleak

Simple function which makes a copy of a slice and returns the copy, so as to avoid potential memory leaks.

It's useful to use this with the output of `bytes.Buffer` if the buffer is to be reset. The `unleak.String()` function is also useful as there is no easy way to make a copy of a string in the current version of Go, without reslicing the original.

Example:

    b := unleak.Bytes(buf.Bytes())
    buf.Reset()

##Index

    func Bytes(b []byte) []byte
    func Runes(b []rune) []rune
    func String(b string) string
    func Ints(b []int) []int
    func Int8s(b []int8) []int8
    func Int16s(b []int16) []int16
    func Int32s(b []int32) []int32
    func Int64s(b []int64) []int64
    func Uints(b []uint) []uint
    func Uint8s(b []uint8) []uint8
    func Uint16s(b []uint16) []uint16
    func Uint32s(b []uint32) []uint32
    func Uint64s(b []uint64) []uint64
    func Float32s(b []float32) []float32
    func Float64s(b []float64) []float64
    func Bools(b []bool) []bool
