package main

import (
	"fmt"

	"github.com/robfig/cron"
)

func StartSchedule() {
	c := cron.New()
	c.AddFunc("@every 5s", func() { fmt.Println("Every 5 seconds") })
	c.Start()
}

func main() {
	fmt.Println("Hello World!")
	StartSchedule()
}
