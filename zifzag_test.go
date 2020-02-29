package zigzag

import (
	"fmt"
	"testing"
)

func TestInt2Zigzag(t *testing.T) {
	a := Int2Zigzag(-37)
	fmt.Println(a)
}

func TestZigzag2Int(t *testing.T) {
	a := Zigzag2Int(73)
	fmt.Println(a)
}
