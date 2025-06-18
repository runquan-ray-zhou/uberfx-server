package cron

import (
	"fmt"
	"net/http"
	"time"

	cron "github.com/robfig/cron/v3"
	"go.uber.org/fx"
)

var Module = fx.Invoke(
	NewCronJob,
)

func NewCronJob() {

	c := cron.New()

	// Schedule the job to ping backend server every 14 minutes to keep it running
	c.AddFunc("@every 14m", func() {
		resp, err := http.Get("http://localhost:8080/") // rest handler endpoint
		if err != nil {
			fmt.Println("Ping failed:", err)
			return
		}
		defer resp.Body.Close()

		fmt.Printf("Pinged server at %s. Status: %s\n", time.Now().Format(time.RFC3339), resp.Status)
	})

	c.Start()
}
