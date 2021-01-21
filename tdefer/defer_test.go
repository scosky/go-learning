package tdefer

import (
	"fmt"
	"testing"
)

/*
1.在一个方法或者函数中，可以有多个 defer 语句
2.多个 defer 语句的执行顺序依照后进先出的原则
*/
func TestDefer(t *testing.T) {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Three defer")
	fmt.Println("函数自身代码")
}
