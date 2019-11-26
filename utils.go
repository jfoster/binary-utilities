package bintools

func NextPowerOfTwo(v int64) int64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	v++
	return v
}

// PrevPowerOfTwo returns the value of previous power of 2
func PrevPowerOfTwo(v int64) int64 {
	return NextPowerOfTwo(v) >> 1
}

// NearPowerOfTwo returns the value of nearest power of 2
func NearPowerOfTwo(v int64) int64 {
	next := NextPowerOfTwo(v)
	if prev := PrevPowerOfTwo(v); (next - v) > (v - prev) {
		return prev
	}
	return next
}

// Bytes creates a byte slice with all indicies initialized to a specified value
func Bytes(b byte, size int64) (bs []byte) {
	if size > 0 {
		bs = make([]byte, size)
		bs[0] = b
		for bp := int64(1); bp < size; bp *= 2 {
			copy(bs[bp:], bs[:bp])
		}
	}
	return bs
}
