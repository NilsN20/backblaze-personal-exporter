package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"prototype/backblaze-personal-exporter/Collectors"
	"prototype/backblaze-personal-exporter/models"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var backblazePath string
var metrics []models.Metric

func addGaugeMetric(name string, help string) {
	var metric models.Metric
	metric.Name = name
	metric.PromGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	})
	metrics = append(metrics, metric)
}
func addGaugeVecMetric(name string, help string, labels []string) {
	var metric models.Metric
	metric.Name = name
	metric.PromGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	}, labels,
	)
	metrics = append(metrics, metric)
}
func initMetrics() {

	addGaugeMetric("transferStatus", "Weather or not it is currently transferring. 1 = Transferring, 0 = Not Transferring")
	addGaugeVecMetric("remainingFilesNum", "Number of remaining files which need to be backed up", []string{"volumeGuid"})
	addGaugeVecMetric("totalFilesNum", "Number of total files marked for backup", []string{"volumeGuid"})
	addGaugeVecMetric("remainingBytes", "Number of remaining bytes which need to be backed up", []string{"volumeGuid"})
	addGaugeVecMetric("totalBytes", "Number of total bytes marked for backup", []string{"volumeGuid"})
	addGaugeVecMetric("lastBackupComplete", "Last time data was fully backed up", []string{"timeFormat"})

	for _, metric := range metrics {
		if metric.PromGauge == nil {
			prometheus.MustRegister(metric.PromGaugeVec)
		} else {
			prometheus.MustRegister(metric.PromGauge)
		}
	}
}

func update(sleepTime int) {
	for {
		Collectors.CollectRemainingbackup(metrics, filepath.Join(backblazePath, "bzdata", "bzreports", "bzstat_remainingbackup.xml"))
		Collectors.CollectTotalbackup(metrics, filepath.Join(backblazePath, "bzdata", "bzreports", "bzstat_totalbackup.xml"))
		Collectors.CollectOverviewstatusMetrics(metrics, filepath.Join(backblazePath, "bzdata", "overviewstatus.xml"))
		Collectors.CollectLastBackup(metrics, filepath.Join(backblazePath, "bzdata", "bzreports", "bzstat_lastbackupcompleted.xml"))
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}
func main() {
	port := flag.Int("port", 8090, "port to listen on")
	sleepTime := flag.Int("updateInterval", 5, "How often to fetch data from backblaze files (in seconds)")
	path := flag.String("backblazeData", "C:\\ProgramData\\Backblaze\\", "Path to backblaze directory")
	flag.Parse()
	backblazePath = *path

	initMetrics()
	go update(*sleepTime)
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening on port:", *port)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
