package metrics

import (
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	execTime = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "appsimulator_exectime_millisec_total",
		Help: "The time elapsed for calculation of a mandelbrot-set",
	}, []string{"function", "dockerHostname", "nodeIP"})
)

func StartMetricsEndpoint(results *map[string]float64) {

	go updateMetrics(results)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(os.Getenv("METRICS_PORT"), nil)
}

func updateMetrics(results *map[string]float64) {
	//tmpUpdateInterval, _ := strconv.Atoi(os.Getenv("UPDATE_INTERVAL_PROMETHEUS"))
	updateInterval := 1500 //1.5 sec

	for true {

		for key, val := range *results {
			execTime.With(prometheus.Labels{
				"function":       key,
				"dockerHostname": os.Getenv("DOCKERHOST"),
				"nodeIP":         os.Getenv("HOSTIP"),
			}).Set(val)
		}

		// Pause before starting next update
		time.Sleep(time.Duration(updateInterval * 1000000))

	}

}
