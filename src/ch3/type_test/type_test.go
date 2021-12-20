package type_test

import "testing"

type Myint int64

func TestIm(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c Myint
	c = Myint(b)
	t.Log(b, c, a)
}
