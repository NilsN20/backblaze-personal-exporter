package Collectors

import (
	"prototype/backblaze-personal-exporter/models"
	"prototype/backblaze-personal-exporter/utils"
	"strconv"
)
import "prototype/backblaze-personal-exporter/models/xmlModels"

func CollectOverviewstatusMetrics(metrics []models.Metric, filePath string) {
	var fileContents = utils.ReadXmlFile[xmlModels.Overviewstatus](filePath)
	for i := range metrics {
		if metrics[i].Name == "transferStatus" {
			metrics[i].PromGauge.Set(utils.TransformCurrentState(fileContents.TransmitState.CurState))
			return
		}
	}
}

func CollectRemainingbackup(metrics []models.Metric, filePath string) {
	var fileContents = utils.ReadXmlFile[xmlModels.Remainingbackup](filePath)
	for i := range metrics {
		if metrics[i].Name == "remainingFilesNum" {
			// for combined volumes:
			var numFilesFloat, _ = strconv.ParseFloat(fileContents.Remaining.RemainingNumFilesForBackup, 64)
			metrics[i].PromGaugeVec.WithLabelValues("All").Set(numFilesFloat)
			// iterating through each volumeGuid:
			for _, BZVolume := range fileContents.BZVolumes {
				var numFilesFloat, _ = strconv.ParseFloat(BZVolume.PervolRemainingFilesNumFiles, 64)
				metrics[i].PromGaugeVec.WithLabelValues(BZVolume.BZVolumeGuid).Set(numFilesFloat)
			}
		}
		if metrics[i].Name == "remainingBytes" {
			var numBytesFloat, _ = strconv.ParseFloat(fileContents.Remaining.RemainingNumBytesForBackup, 64)
			metrics[i].PromGaugeVec.WithLabelValues("All").Set(numBytesFloat)
			for _, BZVolume := range fileContents.BZVolumes {
				var numBytesFloat, _ = strconv.ParseFloat(BZVolume.PervolRemainingFilesNumBytes, 64)
				metrics[i].PromGaugeVec.WithLabelValues(BZVolume.BZVolumeGuid).Set(numBytesFloat)
			}
		}

	}
}

func CollectTotalbackup(metrics []models.Metric, filePath string) {
	var fileContents = utils.ReadXmlFile[xmlModels.Totalbackup](filePath)
	for i := range metrics {
		if metrics[i].Name == "totalFilesNum" {
			// for combined volumes:
			var numFilesFloat, _ = strconv.ParseFloat(fileContents.Totals.TotalNumFiles, 64)
			metrics[i].PromGaugeVec.WithLabelValues("All").Set(numFilesFloat)
			// iterating through each volumeGuid:
			for _, BZVolume := range fileContents.BZVolumes {
				var numFilesFloat, _ = strconv.ParseFloat(BZVolume.PervolTotalNumFiles, 64)
				metrics[i].PromGaugeVec.WithLabelValues(BZVolume.BZVolumeGuid).Set(numFilesFloat)
			}
		}
		if metrics[i].Name == "totalBytes" {
			var numBytesFloat, _ = strconv.ParseFloat(fileContents.Totals.TotalNumBytes, 64)
			metrics[i].PromGaugeVec.WithLabelValues("All").Set(numBytesFloat)
			for _, BZVolume := range fileContents.BZVolumes {
				var numBytesFloat, _ = strconv.ParseFloat(BZVolume.PervolTotalNumBytes, 64)
				metrics[i].PromGaugeVec.WithLabelValues(BZVolume.BZVolumeGuid).Set(numBytesFloat)
			}
		}

	}
}

func CollectLastBackup(metrics []models.Metric, filePath string) {
	var fileContents = utils.ReadXmlFile[xmlModels.Lastbackupcompleted](filePath)
	for i := range metrics {
		if metrics[i].Name == "lastBackupComplete" {
			var gmtMillisFloat, _ = strconv.ParseFloat(fileContents.LastBackup.GmtMillis, 64)
			var localTimeFloat, _ = strconv.ParseFloat(fileContents.LastBackup.Localdatetime, 64)

			metrics[i].PromGaugeVec.WithLabelValues("GMT").Set(gmtMillisFloat)
			metrics[i].PromGaugeVec.WithLabelValues("LocalTime").Set(localTimeFloat)
			return
		}
	}
}
