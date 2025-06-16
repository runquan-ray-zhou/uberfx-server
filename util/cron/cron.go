package cron

import (
	"fmt"
	"net/http"
	"time"

	cron "github.com/robfig/cron/v3"
)

func NewCronJob() {

	c := cron.New()

	// Schedule the job to run every 15 minutes
	c.AddFunc("@every 15m", func() {
		resp, err := http.Get("https://your-backend.com/health") // or your own endpoint
		if err != nil {
			fmt.Println("Ping failed:", err)
			return
		}
		defer resp.Body.Close()

		fmt.Printf("Pinged server at %s. Status: %s\n", time.Now().Format(time.RFC3339), resp.Status)
	})

	c.Start()
}
