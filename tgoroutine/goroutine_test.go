package tgoroutine

import (
	"fmt"
	"testing"
	"time"
)

/*
1.go 关键字后跟一个方法或者函数的调用，就可以启动一个 goroutine
*/
func TestTine(t *testing.T) {
	go fmt.Println("此事儿")
	fmt.Println("测试一哈")
	time.Sleep(time.Second)
}

/*
 channel
1.接收：获取 chan 中的值，操作符为 <- chan。               onlyReceive:=make(<-chan int)
2.发送：向 chan 发送值，把值放在 chan 中，操作符为 chan <- . onlySend := make(chan<- int)
*/
func TestChan1(t *testing.T) {
	//无缓冲channel
	//它的容量是 0，不能存储任何数据。所以无缓冲 channel 只起到传输数据的作用
	//数据并不会在 channel 中做任何停留。这也意味着，无缓冲 channel 的发送和接收操作是同时进行的
	//它也可以称为同步 channel
	ch := make(chan string)

	go func() {
		fmt.Println("接受一个字符串")
		ch <- "goroutine 完成"
	}()

	fmt.Println("我是 main goroutine")

	v := <-ch
	fmt.Println("接收到的chan中的值为：", v)
}

/*
有缓冲 channel 具备以下特点
1.有缓冲 channel 的内部有一个缓冲队列；
2.发送操作是向队列的尾部插入元素，如果队列已满，则阻塞等待，直到另一个 goroutine 执行，接收操作释放队列的空间；
3.接收操作是从队列的头部获取元素并把它从队列中删除，如果队列为空，则阻塞等待，直到另一个 goroutine 执行，发送操作插入新的元素。
*/

/*
关闭 channel  close(cacheCh)
channel 被关闭了，就不能向里面发送数据了，如果发送的话，会引起 painc 异常。
但是还可以接收 channel 里的数据，如果 channel 里没有数据的话，接收的数据是元素类型的零值
*/
func TestCacheCh(t *testing.T) {
	//有缓冲channel
	cacheCh := make(chan int, 5)
	cacheCh <- 2
	cacheCh <- 3
	close(cacheCh)
	fmt.Println("cacheCh容量为:", cap(cacheCh), ",元素个数为：", len(cacheCh))
}
