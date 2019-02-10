package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nimrodshn/mattermost-nws-alerts/weather"
)

func main() {
	stopCh := setupSignalHandler()

	fmt.Println("Starting Alerts service!")
	alertWatcher := weather.AlertWatcher{}
	alertWatcher.Run(stopCh)
}

func setupSignalHandler() <-chan bool {
	signalCh := make(chan os.Signal, 1)
	stopCh := make(chan bool)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalCh
		fmt.Println("Stopping execution...")
		stopCh <- true
		os.Exit(0)
	}()
	return stopCh
}
