package main

import (
	"context"
	"fmt"
	"time"

	cornV3 "github.com/robfig/cron/v3"
)

var jobs map[string]func(ctx context.Context)

func main() {
	corn := cornV3.New(cornV3.WithSeconds())
	// corn.Run()
	corn.Start()
}

func SetCronJobs(cron *cornV3.Cron) {

	jobs["print"] = func(ctx context.Context) {
		fmt.Println("run at:", time.Now().Format("2006-01-02 15:04:05.9"))
	}

	schedule := "*/3 * * * * *"
	cron.AddFunc(schedule, runWithTimeout(jobs["print"]))
}

func runWithTimeout(job func(ctx context.Context)) func() {
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		job(ctx)
	}
}
