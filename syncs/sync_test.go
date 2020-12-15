package syncs

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Singleton struct{}

var singleton *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("init obj")
		singleton = new(Singleton)
	})
	return singleton
}

/**
 * 单例测试
 */
func TestSync(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%p\n", obj)
		}()
	}
	time.Sleep(1000)
}

var number = 0

func add(w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	number = number + 1
	defer m.Unlock()
	defer w.Done()
}

/*
 * 互斥锁
 */
func TestMutex(t *testing.T) {
	var m sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add(&wg, &m)
	}
	wg.Wait()
	fmt.Println("number:", number)
}
