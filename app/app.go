package app

import (
	"fmt"
	"mailer/job"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Cron   *job.Job
}

func (a *App) Init() {
	a.Router = gin.Default()
	a.Cron = &job.Job{}
	a.Cron.JobInit()
	a.InitRoutes()

	if err := a.Router.Run(":8080"); err != nil {
		fmt.Println("Error starting Gin server:", err)
	}
}

func (a *App) InitRoutes() {
	// Define a route to start the cron job
	a.Router.GET("/start", func(c *gin.Context) {
		// Lock the mutex to ensure atomicity
		a.Cron.CronJobMu.Lock()
		defer a.Cron.CronJobMu.Unlock()

		// If cron job is already running, return error
		if a.Cron.IsRunning {
			c.String(http.StatusBadRequest, "Cron job is already running")
			return
		}
		a.Cron.JobStart()
		c.String(http.StatusOK, "Cron job started successfully")
	})

	// Define a route to stop the cron job
	a.Router.GET("/stop", func(c *gin.Context) {
		// Lock the mutex to ensure atomicity
		a.Cron.CronJobMu.Lock()
		defer a.Cron.CronJobMu.Unlock()
		// If cron job is not running, return error
		if a.Cron.CronJob == nil {
			c.String(http.StatusBadRequest, "No cron job is running")
			return
		}
		a.Cron.JobStop()
		c.String(http.StatusOK, "Cron job stopped successfully")
	})
}
