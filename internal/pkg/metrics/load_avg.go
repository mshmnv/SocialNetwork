package metrics

import (
	"github.com/shirou/gopsutil/load"
	logger "github.com/sirupsen/logrus"
)

type LoadAverage struct {
	load1  float64
	load5  float64
	load15 float64
}

func getLoadAvg() *LoadAverage {
	loadAvg, err := load.Avg()
	if err != nil {
		logger.Error("metrics: error getting metric 'load average", err)
		return &LoadAverage{}
	}
	return &LoadAverage{
		load1:  loadAvg.Load1,
		load5:  loadAvg.Load5,
		load15: loadAvg.Load15,
	}
}
