package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type ServerMetrics struct {
	discFree      *prometheus.Desc
	discAvailable *prometheus.Desc
	discSize      *prometheus.Desc

	loadAvg1  *prometheus.Desc
	loadAvg5  *prometheus.Desc
	loadAvg15 *prometheus.Desc
}

func newServerMetrics() *ServerMetrics {
	return &ServerMetrics{

		discFree: prometheus.NewDesc(
			"server_disc_usage_free",
			"Disc usage",
			nil, nil),
		discAvailable: prometheus.NewDesc(
			"server_disc_usage_available",
			"Disc usage",
			nil, nil),
		discSize: prometheus.NewDesc(
			"server_disc_usage_size",
			"Disc usage",
			nil, nil),

		loadAvg1: prometheus.NewDesc(
			"server_load_average_1",
			"Load Average",
			nil, nil),
		loadAvg5: prometheus.NewDesc(
			"server_load_average_5",
			"Load Average",
			nil, nil),
		loadAvg15: prometheus.NewDesc(
			"server_load_average_15",
			"Load Average",
			nil, nil),
	}
}

func (s *ServerMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- s.discFree
}

func (s *ServerMetrics) Collect(ch chan<- prometheus.Metric) {

	currentTime := time.Now()

	discUsageInfo := NewDiskUsage()

	ch <- prometheus.NewMetricWithTimestamp(currentTime, prometheus.MustNewConstMetric(s.discFree, prometheus.GaugeValue, float64(discUsageInfo.Free())))
	ch <- prometheus.NewMetricWithTimestamp(currentTime, prometheus.MustNewConstMetric(s.discAvailable, prometheus.GaugeValue, float64(discUsageInfo.Available())))
	ch <- prometheus.NewMetricWithTimestamp(currentTime, prometheus.MustNewConstMetric(s.discSize, prometheus.GaugeValue, float64(discUsageInfo.Size())))

	loadAvg := getLoadAvg()

	ch <- prometheus.NewMetricWithTimestamp(currentTime, prometheus.MustNewConstMetric(s.loadAvg1, prometheus.GaugeValue, loadAvg.load1))
	ch <- prometheus.NewMetricWithTimestamp(currentTime, prometheus.MustNewConstMetric(s.loadAvg5, prometheus.GaugeValue, loadAvg.load5))
	ch <- prometheus.NewMetricWithTimestamp(currentTime, prometheus.MustNewConstMetric(s.loadAvg15, prometheus.GaugeValue, loadAvg.load15))

}
