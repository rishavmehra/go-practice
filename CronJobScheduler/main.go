package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	// c := cron.New()

	// // c.AddFunc("1 * * * *", func() { Hello("rishavehra") })
	// c.AddFunc("@every 0h1m0s", func() { Hello("rishavehra") })

	// c.Start()
	// time.Sleep(70 * time.Second)

	// c.Stop()

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Panic(err)
	}

	j, err := s.NewJob(
		gocron.DurationJob(
			1*time.Second,
		), gocron.NewTask(
			func(a int, b int) {
				c := a + b
				fmt.Println(c)
			},
			1, 1,
		),
	)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(j.ID())
	s.Start()

	time.Sleep(10 * time.Second)

	err = s.Shutdown()
	if err != nil {
		log.Panic(err)
	}

}

func Hello(st string) {
	fmt.Printf("Hello, %s", st)
}
