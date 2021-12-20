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

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// 指针不支持运算
	//aPtr = aPtr+1
	t.Log(a, aPtr)
	t.Logf("%T,%T", a, aPtr)
}

// golang中默认字符串是空字符串
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
