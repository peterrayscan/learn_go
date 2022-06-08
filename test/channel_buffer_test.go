package test

import (
	"fmt"
	"learn/tools/testx"
	"learn/tools/timex"
	"testing"
	"time"
)

func Test_ChannelWithBufferInDiffGoroutine(t *testing.T) {
	testx.RunFunc(func() {
		c := make(chan struct{}, 1)

		go func(c chan struct{}) {
			time.Sleep(1 * time.Second)
			c <- struct{}{}
			// 发送值 不会被阻塞，因为允许把值放在缓冲区
			fmt.Println("done in other goroutinue")
		}(c)

		time.Sleep(3 * time.Second)
		fmt.Println("receive:", <-c)
		// 主协程会在 接收（到）值 之前被阻塞
		fmt.Println("done in main goroutinue")
	})
}

func Test_ChannelWithoutBufferInDiffGoroutine(t *testing.T) {
	testx.RunFunc(func() {
		c := make(chan struct{})

		th := timex.NewTimeHelper()
		fmt.Println("now in main goroutinue:", th.SetTime(time.Now()).GetTimeString())

		go func(c chan struct{}) {
			time.Sleep(1 * time.Second)
			c <- struct{}{}
			// 因为没有缓冲区
			// 所以这里的 协程 会被阻塞
			// 直到 channel 的值被取出，才会疏通，下面的语句才会执行
			fmt.Println("now in other goroutinue:", th.SetTime(time.Now()).GetTimeString())
		}(c)

		time.Sleep(3 * time.Second)
		fmt.Println("now in main goroutinue:", th.SetTime(time.Now()).GetTimeString())
		// 主协程会在 接收（到）值 之前被阻塞
		fmt.Println("receive:", <-c)

		time.Sleep(1 * time.Second)
	})
}

func Test_ChannelWithBufferInSameGoroutine(t *testing.T) {
	testx.RunFunc(func() {
		c := make(chan struct{}, 1)

		time.Sleep(1 * time.Second)
		// 主协程 发送值 不会被阻塞，因为有通道channel
		c <- struct{}{}

		// 主协程会在 接收（到）值 之前被阻塞
		fmt.Println("receive:", <-c)
		fmt.Println("done")
	})
}

func Test_ChannelWithoutBufferInSameGoroutine(t *testing.T) {
	testx.RunFunc(func() {
		c := make(chan struct{})

		time.Sleep(1 * time.Second)
		c <- struct{}{}
		// 因为没有 缓冲区，发送值 后
		// 通道channel 被阻塞，进一步导致 接收值 被阻塞
		// 直到test超时
		fmt.Println("receive:", <-c)
	})
}
