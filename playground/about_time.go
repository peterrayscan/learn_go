package playground

import (
	"context"
	"fmt"
	"learn/tools/timex"
	"time"
)

func AddDurationForTime() {
	fmt.Println("empty time struct: ", time.Time{})

	now := time.Now()
	fmt.Println("now: ", now)

	fmt.Println("after add")
	var hours uint32 = 2
	addDuration := time.Duration(hours) * time.Hour
	fmt.Println("addDuration is ", addDuration)
	fmt.Println("now: ", now.Add(addDuration))
}

func FormatStringTime() {
	format := "2006-01-02 15:04:05"
	timeString := "1970-01-01 00:00:00"
	time, err := time.Parse(format, timeString)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("time:", time)
}

func FormatTimeOutput() {
	format := "2006-01-02 15:04:05.000000000"
	fmt.Println("now is: ", time.Now().Format(format))
}

// time.After 实现
func TimeOutVersion1() {
	timeout := time.After(3 * time.Second)
	for {
		select {
		case <-timeout:
			fmt.Println("timeout")
			return
		default:
			fmt.Println("default, now:", timex.GetNowString())
		}
		time.Sleep(1 * time.Second)
	}
}

// time.After 实现
func TimeOutVersion2() {
	ch := make(chan struct{}, 1)
	go func() {
		var x int
		for {
			fmt.Println("x is", x)
			fmt.Println("do sth in another goroutinue")
			time.Sleep(1 * time.Second)
			if x == 3 {
				ch <- struct{}{}
				return
			}
			x++
		}
	}()

	select {
	case <-ch:
		fmt.Println("finish job")
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}
}

// context.WithTimeout 实现
func TimeOutVersion3() {
	timeOutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch := make(chan struct{}, 1)
	go func() {
		var x int
		for {
			fmt.Println("x is", x)
			fmt.Println("do sth in another goroutinue")
			time.Sleep(1 * time.Second)
			if x == 8 {
				ch <- struct{}{}
				return
			}
			x++
		}
	}()

	select {
	case <-ch:
		fmt.Println("finish job")
	case <-timeOutCtx.Done():
		fmt.Println("timeout")
		fmt.Println("timeout err:", timeOutCtx.Err())
	}
}
