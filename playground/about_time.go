package playground

import (
	"fmt"
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

func TimeOutAfter(seconds int) {
	for range time.Tick(2 * time.Second) {
		fmt.Println("xxx")
	}
	// fmt.Println(time.Now())
	// after := <-time.After(time.Duration(seconds) * time.Second)
	// fmt.Println("xxx")
	// fmt.Println(after)
}
