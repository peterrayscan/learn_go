package main

import (
	"context"
	"fmt"
	"log"
	"time"

	cornV3 "github.com/robfig/cron/v3"
)

var jobs = map[string]func(ctx context.Context){}

// 耗时至少20ms
func costAtLeast20ms() {
	sli := make([]int, 0)
	for i := 0; i < 2000000; i++ {
		sli = append(sli, i)
	}
	fmt.Println("len of sli:", len(sli))
}

func main() {

	jobs["print"] = func(ctx context.Context) {

		fmt.Println()
		fmt.Println("now is:    ", time.Now().Format("2006-01-02 15:04:05.000"))

		deadline, ok := ctx.Deadline()
		if !ok {
			fmt.Println("deadline or timeout not set")
			return
		}

		costAtLeast20ms()
		// 判断是否超时
		if time.Now().After(deadline) {
			fmt.Println("timeout!")
			return
		}

		fmt.Println("now is:    ", time.Now().Format("2006-01-02 15:04:05.000"))
		fmt.Println()
	}

	corn := cornV3.New(cornV3.WithSeconds())
	err := SetCronJobs(corn)
	if err != nil {
		log.Fatalf("err:%+v", err)
	}

	fmt.Println("start at:   ", time.Now().Format("2006-01-02 15:04:05.000"))
	corn.Run()
}

func SetCronJobs(cron *cornV3.Cron) error {
	// 0，3，6，9，12...
	// 每三秒执行一次
	schedule := "*/3 * * * * *"
	for _, job := range jobs {
		_, err := cron.AddFunc(schedule, withTimeout(job))
		if err != nil {
			return err
		}
	}
	return nil
}

func withTimeout(job func(ctx context.Context)) func() {
	return func() {
		// 20 毫秒就超时
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
		defer cancel()
		job(ctx)
	}
}
