package main

import (
	"context"
	"fmt"
	"time"

	cornV3 "github.com/robfig/cron/v3"
)

var jobs map[string]func(ctx context.Context)

func main() {

	jobs = make(map[string]func(ctx context.Context))
	jobs["print"] = func(ctx context.Context) {
		time.Sleep(4 * time.Second) // spend time
		fmt.Println("now is:", time.Now().Format("2006-01-02 15:04:05"))
	}

	corn := cornV3.New(cornV3.WithSeconds())
	SetCronJobs(corn)

	corn.Run()
}

func SetCronJobs(cron *cornV3.Cron) {
	schedule := "*/3 * * * * *"
	cron.AddFunc(schedule, runWithTimeout(jobs["print"]))
}

func runWithTimeout(job func(ctx context.Context)) func() {
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 0*time.Second)
		defer cancel()
		job(ctx)
	}
}
