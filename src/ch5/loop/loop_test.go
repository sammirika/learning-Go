package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	// 等于for
	for n < 5 {
		t.Log(n)
		n++
	}
}
