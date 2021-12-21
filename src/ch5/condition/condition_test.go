package condition

import "testing"

func TestIFMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}

	// go语言方法支持多返回值用于判断是否返回正常不报错
	//if v,err := someFun(); err==nil {
	//	t.Log("1==1")
	//}
	// golang中的switch支持多变量
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it's not 0-3")
		}
	}
}
