package test

import (
	"context"
	"fmt"
	"learn/tools/testx"
	"learn/tools/timex"
	"testing"
	"time"
)

// 轮询（间隔单位时间），同时一定时间后，会触发超时

var result string = "no result"
var handleResultNeeded time.Duration = 4900 // 4.9秒就提前更新值，防止出现多次结果不一致的情况
// var maxWaitTime time.Duration = 8000
var maxWaitTime time.Duration = 3500 // 3.5秒就提前达到超时时间，防止出现多次结果不一致的情况
var intervalTime time.Duration = 1

func anotherGoroutinueHandleResult() {
	time.Sleep(handleResultNeeded * time.Millisecond)
	result = "got result"
}

func pollingResult() string {
	return result
}

func Test_PollingWithTimeoutV1(t *testing.T) {
	testx.RunFunc(func() {

		go anotherGoroutinueHandleResult()

		timeout := time.After(maxWaitTime * time.Millisecond)
		for {
			select {
			case <-timeout:
				fmt.Println("timeout!")
				return
			default:
				res := pollingResult()
				fmt.Println(res)
				if res == "got result" {
					return
				}
			}
			time.Sleep(intervalTime * time.Second)
		}
	})

}

func Test_PollingWithTimeoutV2(t *testing.T) {
	testx.RunFunc(func() {
		go anotherGoroutinueHandleResult()

		notifyCh := make(chan struct{}, 1)
		resultCh := make(chan string, 1)
		pollingFunc := func() {
			for {
				res := pollingResult()
				fmt.Println(res)
				fmt.Println("now", timex.GetNowString())
				if res == "got result" {
					notifyCh <- struct{}{}
					resultCh <- res
					return
				}
				time.Sleep(intervalTime * time.Second)
			}
		}
		go pollingFunc()

		timeout := time.After(maxWaitTime * time.Millisecond)
		select {
		case <-notifyCh:
			fmt.Println(<-resultCh)
			return
		case <-timeout:
			fmt.Println("timeout!")
			fmt.Println("now", timex.GetNowString())
			return
		}
	})
}

func Test_PollingWithTimeoutV3(t *testing.T) {
	testx.RunFunc(func() {
		go anotherGoroutinueHandleResult()

		notifyCh := make(chan struct{}, 1)
		resultCh := make(chan string, 1)
		pollingFunc := func() {
			for {
				res := pollingResult()
				fmt.Println(res)
				if res == "got result" {
					notifyCh <- struct{}{}
					resultCh <- res
					return
				}
				time.Sleep(intervalTime * time.Second)
			}
		}
		go pollingFunc()

		timeOutCtx, cancel := context.WithTimeout(context.Background(), maxWaitTime*time.Millisecond)
		defer cancel()
		select {
		case <-notifyCh:
			fmt.Println(<-resultCh)
			return
		case <-timeOutCtx.Done():
			fmt.Println("timeout!")
			return
		}
	})
}
