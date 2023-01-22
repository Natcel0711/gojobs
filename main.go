package main

import (
	"time"

	jobs "github.com/Natcel0711/gojobs/jobs/mibanco"
	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(60).Seconds().Do(func() {
		jobs.MiBanco()
	})

	s.StartBlocking()
}
