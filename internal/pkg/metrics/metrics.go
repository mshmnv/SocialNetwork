package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/negroni"
)

var (
	requestLabels = []string{"handler", "code"}

	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_count",
			Help: "Number of requests.",
		}, requestLabels)

	responseTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration",
			Help: "Response time.",
		}, requestLabels)
)

func init() {
	prometheus.MustRegister(newServerMetrics(), totalRequests, responseTime)
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		lrw := negroni.NewResponseWriter(w)

		next.ServeHTTP(lrw, r)

		labels := []string{r.URL.Path, strconv.Itoa(lrw.Status())}

		totalRequests.WithLabelValues(labels...).Inc()
		responseTime.WithLabelValues(labels...).Observe(time.Since(start).Seconds())
	})
}
