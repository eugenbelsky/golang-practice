package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "counter_check",
		Help: "The total number of processed events",
	})
)

func main() {
	// cmd, err := exec.Command(“ls”, “-qwe”).Output()
	execAction()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
func execAction() {
	go func() {
		for {
			cmd := exec.Command("ls", "-la") // Action block
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("%s", output)
			}
			fmt.Printf("Output: %s", output)
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
