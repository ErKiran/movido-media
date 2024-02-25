package cron

import (
	"github.com/robfig/cron/v3"
)

type CronJob struct {
	cron *cron.Cron
}

func NewCronJob() *CronJob {
	return &CronJob{
		cron: cron.New(),
	}
}

func (cj *CronJob) AddFunc(spec string, cmd func()) {
	cj.cron.AddFunc(spec, cmd)
}

func (cj *CronJob) Start() {
	cj.cron.Start()
}

func (cj *CronJob) Stop() {
	cj.cron.Stop()
}
