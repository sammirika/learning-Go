package operactor

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c:=[...]int{1,2,3,4,5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// 长度不同的数组不能比较，有编译错误
	//t.Log(a==c)
	t.Log(a == d)
}
