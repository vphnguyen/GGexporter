// ExtractModel (child)
//   - Định nghĩa struct ExtractModel dùng để ánh xạ các field trong xml thành Object.
package model

import "encoding/xml"

// Định nghĩa struct ExtractModel
type ExtractModel struct {
	XMLName                   xml.Name `xml:"mpoints"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Name                      string   `xml:"name,attr"`
	Process                   struct {
		Text            string `xml:",chardata"`
		Name            string `xml:"name,attr"`
		Type            string `xml:"type,attr"`
		Mode            string `xml:"mode,attr"`
		Status          string `xml:"status,attr"`
		ProcessID       string `xml:"process-id"`
		StartTime       string `xml:"start-time"`
		PortNumber      string `xml:"port-number"`
		FirstMessage    string `xml:"first-message"`
		LastMessage     string `xml:"last-message"`
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
		PositionEr struct {
			Text            string `xml:",chardata"`
			LowWatermarkLag string `xml:"low-watermark-lag"`
			LowWatermarkTs  struct {
				Text string `xml:",chardata"`
				Nil  string `xml:"nil,attr"`
			} `xml:"low-watermark-ts"`
			HighWatermarkLag string `xml:"high-watermark-lag"`
			HighWatermarkTs  struct {
				Text string `xml:",chardata"`
				Nil  string `xml:"nil,attr"`
			} `xml:"high-watermark-ts"`
			LastOperationLag string `xml:"last-operation-lag"`
			LastOperationTs  string `xml:"last-operation-ts"`
			LastCheckpointTs string `xml:"last-checkpoint-ts"`
			InputCheckpoint  string `xml:"input-checkpoint"`
			OutputCheckpoint string `xml:"output-checkpoint"`
			EndOfFile        string `xml:"end-of-file"`
			TrailTimesAtEof  string `xml:"trail-times-at-eof"`
		} `xml:"position-er"`
		StatisticsExtract struct {
			Text                     string `xml:",chardata"`
			MappedTotalInserts       string `xml:"mapped-total-inserts"`
			MappedTotalUpdates       string `xml:"mapped-total-updates"`
			MappedTotalDeletes       string `xml:"mapped-total-deletes"`
			MappedTotalUpserts       string `xml:"mapped-total-upserts"`
			MappedTotalUnsupported   string `xml:"mapped-total-unsupported"`
			MappedTotalTruncates     string `xml:"mapped-total-truncates"`
			TotalExecutedDdls        string `xml:"total-executed-ddls"`
			TotalExecutedProcedures  string `xml:"total-executed-procedures"`
			TotalDiscards            string `xml:"total-discards"`
			TotalIgnores             string `xml:"total-ignores"`
			TotalConversionErrors    string `xml:"total-conversion-errors"`
			TotalConversionTruncates string `xml:"total-conversion-truncates"`
			TotalRowFetchAttempts    string `xml:"total-row-fetch-attempts"`
			TotalRowFetchFailures    string `xml:"total-row-fetch-failures"`
		} `xml:"statistics-extract"`
	} `xml:"process"`
}

// Là một phương thức của Extract Model dùng để kiểm tra thử xem Extract này có được dùng làm InitLoad (Sourceistable) không.
func (m *ExtractModel) IsInitLoad() bool {
	if len(m.Process.TrailOutput) == 0 {
		return true
	}
	return false
}
func (m *ExtractModel) IsANewOne() bool {
	if m.Name != "" && m.Process.Name == "" {
		return true
	}
	return false
}
