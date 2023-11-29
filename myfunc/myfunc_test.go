package myfunc

import (
	"fmt"
	"testing"
)

func TestMyfunc(t *testing.T) { //用于编写测试
	sum := Myfunc(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleMyfunc() { //用于编写示例（后续可在文档中显示），其中需在最后一行通过注释方式详细写出 期望的输出，如不写这行，之后运行​​go test​​时只编译该示例函数但并不执行
	sum := Myfunc(3, 4)
	fmt.Println(sum)
	// Output: 7
}
