package zigzag
import "C"

// 判断系统是32位还是64位
const IntBitSize  = 32 << (^uint(0) >> 63)

func Int2Zigzag(n int) uint {
	return uint((n << 1) ^ (n >> 31))
}

func Zigzag2Int(n uint) int {
	return int((n >> 1) ^ -(n & 1))
}

