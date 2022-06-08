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
			// it will not be blocked, cause this channel has buffer
			fmt.Println("done in other goroutinue")
		}(c)

		time.Sleep(3 * time.Second)
		fmt.Println("receive:", <-c)
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
			// it will be blocked, untill the main goroutinue receive the data
			// after that, the below string will print
			fmt.Println("now in other goroutinue:", th.SetTime(time.Now()).GetTimeString())
		}(c)

		time.Sleep(3 * time.Second)
		fmt.Println("now in main goroutinue:", th.SetTime(time.Now()).GetTimeString())
		fmt.Println("receive:", <-c)

		time.Sleep(1 * time.Second)
	})
}

func Test_ChannelWithBufferInSameGoroutine(t *testing.T) {
	testx.RunFunc(func() {
		c := make(chan struct{}, 1)

		time.Sleep(1 * time.Second)
		c <- struct{}{}

		fmt.Println("receive:", <-c)
	})
}

func Test_ChannelWithoutBufferInSameGoroutine(t *testing.T) {
	testx.RunFunc(func() {
		c := make(chan struct{})

		time.Sleep(1 * time.Second)
		c <- struct{}{}
		// it will be blocked until test timeout
		fmt.Println("receive:", <-c)
	})
}
