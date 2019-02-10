package weather

import (
	fmt "fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const activeAlertsEndpoint = "https://api.weather.gov/alerts/active"
const pollPeriod = 1 * time.Minute

// AlertWatcher watches for weather alerts.
type AlertWatcher struct{}

// Run starts the worker that polls for active
// weather alerts.
func (a *AlertWatcher) Run(stopCh <-chan bool) {
	ticker := time.NewTicker(pollPeriod)
	for {
		select {
		case <-ticker.C:
			a.fetchAlerts()
		case <-stopCh:
			break
		}
	}
}

func (a *AlertWatcher) fetchAlerts() {
	resp, err := http.Get(activeAlertsEndpoint)
	if err != nil {
		fmt.Printf("Some error ocurred while trying to fetch alerts: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("api.weather.gov returned some error code: %v", resp.Status)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body of response: %v", err)
	}
	fmt.Println(string(bytes))
}
