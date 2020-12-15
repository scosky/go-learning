package fmt

import (
	"fmt"
	"testing"
)

func TestScanln(t *testing.T) {
	var name string
	var age int
	fmt.Print("输入姓名和年龄，使用空格分隔：")
	fmt.Print("\n")
	fmt.Scanln(&name, &age)
	fmt.Printf("name: %s\nage: %d\n", name, age)
}

func TestPrintf(t *testing.T) {
	fmt.Printf("%f", 12.56)
	fmt.Print("\n")
	fmt.Printf("%d", 102)
	fmt.Print("\n")
	fmt.Printf("%s", "中古")
	fmt.Print("\n")
	fmt.Printf("%t", true)
	fmt.Print("\n")
	fmt.Printf("%T", 123)
	fmt.Print("\n")
	fmt.Printf("%v", 123)
	fmt.Print("\n")
}

func TestPrintln(t *testing.T) {
	a := true
	b := "string"
	c := 1232
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	n, err := fmt.Println("fmt println func test")
	if err == nil {
		fmt.Println(n)
	}
}
