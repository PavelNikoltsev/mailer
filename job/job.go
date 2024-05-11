package job

import (
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type Job struct {
	CronJob      *cron.Cron
	cronJobEntry cron.EntryID
	CronJobMu    sync.Mutex
	IsRunning    bool
}

func (j *Job) JobInit() {
	// Create a new cron job
	j.CronJob = cron.New()
	// Add a cron job to execute every second
	j.cronJobEntry, _ = j.CronJob.AddFunc("@every 1s", func() {
		fmt.Println("Cron job executed at:", time.Now())
	})
}

func (j *Job) JobStart() {
	// Start the cron scheduler
	if j.CronJob == nil {
		j.JobInit()
	}
	j.CronJob.Start()
	j.IsRunning = true
}
func (j *Job) JobStop() {
	// Stop the cron scheduler
	j.CronJob.Stop()
	j.CronJob = nil
	j.IsRunning = false
}
