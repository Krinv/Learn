// 并发
package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Goroutines：
		Go中的并发执行单位，类似轻量级线程
		Goroutine的调度由Go运行时管理，用户无需手动分配线程
		Goroutine是非阻塞，可以高效地运行成千上万个
		go关键字启动
	Channel：
		Go中用于在Goroutine间的通信机制
		支持同步、数据共享，避免了显式的锁机制
		chan关键字创建，<-操作符发送、接收数据
	Scheduler：
		Go的调度器基于GMP模型，调度器会将Goroutine分配到系统线程中执行，通过M（Machine系统线程）、P（Processor逻辑处理器）的配合高效管理并发
*/

func main() {
	// 启动Goroutine，两者会并发执行
	go sayHello()

	// 另一个Goroutine
	for i := 0; i < 3; i++ {
		fmt.Println("Main")
		time.Sleep(100 * time.Millisecond)
	}

	// 通道例
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	// 后半部sum
	go sum(s[:len(s)/2], c)
	// 前半部sum
	go sum(s[len(s)/2:], c)
	// 由于两个Goroutine并发，顺序不定，因而打印时x、y可能互换
	x, y := <-c, <-c

	fmt.Println()
	fmt.Println(x, y, x+y)

	// 带缓冲区的chan
	//
	fmt.Println()
	// 缓冲区大小为2
	ch := make(chan int, 2)
	// 因为有缓冲区,因而可以同时发送两个数据,而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2 //不会发生阻塞

	// 获取
	// 带缓冲区的不需要发送接收并发，
	//缓冲区满了再接收/缓冲区空了即发送		仍需要并发，否则会死锁
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	/*
		对于无缓冲区的chan，要求发送和接收同步，否则会阻塞直到有接受者准备好接收数据。若不能保证同步，会导致死锁
		对于有缓冲区未满，则不会发生阻塞
	*/

	// 死锁例
	// ch := make(chan int) // 无缓冲通道
	// // 发送操作会阻塞，直到有接收者
	// ch <- 1 // 死锁：没有接收者，主goroutine永远阻塞
	// // 这行代码永远不会执行
	// fmt.Println(<-ch)
	// 上述例子的原因：发送和接收并没有并发执行

	//通道遍历、关闭例
	fmt.Println()
	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1) //cap函数是获取通道缓冲容量

	// 使用range遍历每个从管道接收的数据
	// c在接收10个后关闭，因而range接收10个关闭
	// 若c不关闭,则range函数不会结束
	for i := range c1 {
		fmt.Println(i)
	}

	// select例
	fmt.Println()
	c2 := make(chan int)
	quit := make(chan int)

	// 匿名Goroutine，并发
	go func() {
		for i := 0; i < 10; i++ {
			// 接收并打印
			fmt.Println(<-c2)
		}
		// 打印完发送结束信号
		quit <- 0
	}()

	fibonacci2(c2, quit)

	// sync.WaitGroup例
	fmt.Println()
	var wg sync.WaitGroup

	// 循环启动三个Goroutine，并且会等待某个结束再启动
	// 因而输出顺序可能不一样,由go自动运行管理
	for i := 1; i <= 3; i++ {
		wg.Add(1) //增加计数器，告诉WaitGroup等待一个额外的Goroutine完成
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")

	// 其他
	/*
		//控制Goroutine生命周期
		context.WithCancel
		contex.WithTimeout
	*/

	// 提供互斥锁保护共享资源
	var mu sync.Mutex
	mu.Lock()
	mu.Unlock()
}

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

// 通道：用于Goroutine间传递数据
func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	// 将sum传递到chan c
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	// v,ok:=<-ch在通道接收不到数据时ok为false
	// 此时用close函数关闭通道
	close(c)
}

// select:使一个Goroutine可以等待多个通信操作
// select会阻塞，直到其中某个case可以继续执行
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	// 无限循环
	for {
		// 阻塞直到某个case
		select {
		//尝试发送x到c,c阻塞,但匿名函数中有接收,因而执行
		case c <- x: //十次后发现没有接收者，因而不会执行发送
			x, y = y, x+y
		case <-quit: //十次后发现quit有接收者，因而执行此case
			fmt.Println("quit")
			return
		}
	}
}

// sync.WaitGroup用于等待多个Goroutine完成
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() //Goroutine完成时调用Done(),以及保证panic时Done的执行，不会使主Goroutine永远等待
	fmt.Printf("Worker %d started\n", id)
	fmt.Printf("Worker %d finished\n", id)
}
