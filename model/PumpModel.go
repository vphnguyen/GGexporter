package model

import "encoding/xml"

type PumpModel struct {
	XMLName                   xml.Name `xml:"mpoints"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Name                      string   `xml:"name,attr"`
	Process                   struct {
		XMLName         xml.Name `xml:"process"`
		Text            string   `xml:",chardata"`
		Name            string   `xml:"name,attr"`
		Type            string   `xml:"type,attr"`
		Status          string   `xml:"status,attr"`
		ProcessID       string   `xml:"process-id"`
		StartTime       string   `xml:"start-time"`
		PortNumber      string   `xml:"port-number"`
		FirstMessage    string   `xml:"first-message"`
		LastMessage     string   `xml:"last-message"`
		ConfigurationEr struct {
			Text          string `xml:",chardata"`
			GgsVersion    string `xml:"ggs-version"`
			InputType     string `xml:"input-type"`
			OutputType    string `xml:"output-type"`
			BuildDatabase string `xml:"build-database"`
			BuildPlatform string `xml:"build-platform"`
			BuildDate     string `xml:"build-date"`
		} `xml:"configuration-er"`
		DatabaseInOut struct {
			Text         string `xml:",chardata"`
			DbName       string `xml:"db-name"`
			DbType       string `xml:"db-type"`
			DbVersion    string `xml:"db-version"`
			DbServerName string `xml:"db-server-name"`
			DbGlobalName string `xml:"db-global-name"`
		} `xml:"database-in-out"`
		TrailInput struct {
			Text            string `xml:",chardata"`
			TrailName       string `xml:"trail-name"`
			TrailPath       string `xml:"trail-path"`
			IoReadCount     string `xml:"io-read-count"`
			IoReadBytes     string `xml:"io-read-bytes"`
			TrailReadErrors string `xml:"trail-read-errors"`
			TrailTimesAtEof string `xml:"trail-times-at-eof"`
			TrailLobBytes   string `xml:"trail-lob-bytes"`
			TrailReadTime   string `xml:"trail-read-time"`
			TrailRba        string `xml:"trail-rba"`
			TrailSeq        string `xml:"trail-seq"`
		} `xml:"trail-input"`
		TrailOutput []struct {
			Text                   string `xml:",chardata"`
			TrailName              string `xml:"trail-name"`
			TrailPath              string `xml:"trail-path"`
			Hostname               string `xml:"hostname"`
			PortNumber             string `xml:"port-number"`
			PortType               string `xml:"port-type"`
			IoWriteCount           string `xml:"io-write-count"`
			IoWriteBytes           string `xml:"io-write-bytes"`
			TrailWriteTime         string `xml:"trail-write-time"`
			TrailLobBytes          string `xml:"trail-lob-bytes"`
			TrailBufferFlushes     string `xml:"trail-buffer-flushes"`
			TrailBufferFlushTime   string `xml:"trail-buffer-flush-time"`
			TrailBytesFlushed      string `xml:"trail-bytes-flushed"`
			TrailLastBufferFlushTs string `xml:"trail-last-buffer-flush-ts"`
			TrailRba               string `xml:"trail-rba"`
			TrailSeq               string `xml:"trail-seq"`
			TrailMaxBytes          string `xml:"trail-max-bytes"`
		} `xml:"trail-output"`
		ProcessPerformance struct {
			Text               string `xml:",chardata"`
			ProcessStartTime   string `xml:"process-start-time"`
			ProcessID          string `xml:"process-id"`
			ThreadCount        string `xml:"thread-count"`
			HandleCount        string `xml:"handle-count"`
			CpuTime            string `xml:"cpu-time"`
			KernelTime         string `xml:"kernel-time"`
			UserTime           string `xml:"user-time"`
			IoReadCount        string `xml:"io-read-count"`
			IoWriteCount       string `xml:"io-write-count"`
			IoOtherCount       string `xml:"io-other-count"`
			IoReadBytes        string `xml:"io-read-bytes"`
			IoWriteBytes       string `xml:"io-write-bytes"`
			IoOtherBytes       string `xml:"io-other-bytes"`
			PageFaults         string `xml:"page-faults"`
			WorkingSetSize     string `xml:"working-set-size"`
			PeakWorkingSetSize string `xml:"peak-working-set-size"`
			PrivateBytes       string `xml:"private-bytes"`
		} `xml:"process-performance"`
		CacheStatistics struct {
			Text                      string `xml:",chardata"`
			TotalObjectsInCache       string `xml:"total-objects-in-cache"`
			TotalObjects              string `xml:"total-objects"`
			TotalObjectsActive        string `xml:"total-objects-active"`
			TotalObjectsCommitted     string `xml:"total-objects-committed"`
			MaxActiveObjects          string `xml:"max-active-objects"`
			TimesBufferOverflowed     string `xml:"times-buffer-overflowed"`
			TimesGetNextFromFile      string `xml:"times-get-next-from-file"`
			TimesGetLastFromFile      string `xml:"times-get-last-from-file"`
			TimesSmallBuffForcedOut   string `xml:"times-small-buff-forced-out"`
			TimesRetrieved            string `xml:"times-retrieved"`
			TotalNumberOfQHits        string `xml:"total-number-of-q-hits"`
			TotalNumberOfQMisses      string `xml:"total-number-of-q-misses"`
			TotalNumberOfQPuts        string `xml:"total-number-of-q-puts"`
			TotalNumberOfQTries       string `xml:"total-number-of-q-tries"`
			TotalNumberOfQEntries     string `xml:"total-number-of-q-entries"`
			MaxNumberOfQEntries       string `xml:"max-number-of-q-entries"`
			TotalMunmap               string `xml:"total-munmap"`
			TotalCnnblAttempts        string `xml:"total-cnnbl-attempts"`
			TotalCnnblSuccess         string `xml:"total-cnnbl-success"`
			TotalCnnblMbufs           string `xml:"total-cnnbl-mbufs"`
			TotalFileCacheRequests    string `xml:"total-file-cache-requests"`
			TotalFileCacheEntries     string `xml:"total-file-cache-entries"`
			TotalFileCachePlaced      string `xml:"total-file-cache-placed"`
			MaxQlength                string `xml:"max-qlength"`
			MaxProcessed              string `xml:"max-processed"`
			TimesWaitSignaled         string `xml:"times-wait-signaled"`
			TimesFileCacheNotNeeded   string `xml:"times-file-cache-not-needed"`
			TimesRequestorNeededFc    string `xml:"times-requestor-needed-fc"`
			TotalObjectsInFileCache   string `xml:"total-objects-in-file-cache"`
			TotalFileCacheBytesToDisk string `xml:"total-file-cache-bytes-to-disk"`
			TimesCacheFlushed         string `xml:"times-cache-flushed"`
			MaxMemoryUsage            string `xml:"max-memory-usage"`
			AverageMemoryUsage        string `xml:"average-memory-usage"`
		} `xml:"cache-statistics"`
	} `xml:"process"`
}
