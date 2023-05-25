package cron

import (
	"time"

	"github.com/go-co-op/gocron"
)

func InitializeCron() *gocron.Scheduler {
	s := gocron.NewScheduler(time.UTC)

	return s
}
