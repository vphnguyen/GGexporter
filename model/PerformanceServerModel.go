// PerformanceServerModel (child)
//   - Định nghĩa struct PerformanceServerModel dùng để ánh xạ các field trong xml thành Object.
package model

import "encoding/xml"

// Định nghĩa struct PerformanceServerModel
type PerformanceServerModel struct {
	XMLName                   xml.Name `xml:"mpoints"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Name                      string   `xml:"name,attr"`
	Process                   struct {
		Text                 string `xml:",chardata"`
		Name                 string `xml:"name,attr"`
		Type                 string `xml:"type,attr"`
		Status               string `xml:"status,attr"`
		ProcessID            string `xml:"process-id"`
		StartTime            string `xml:"start-time"`
		PortNumber           string `xml:"port-number"`
		FirstMessage         string `xml:"first-message"`
		LastMessage          string `xml:"last-message"`
		ConfigurationManager struct {
			Text          string `xml:",chardata"`
			GgsVersion    string `xml:"ggs-version"`
			Hostname      string `xml:"hostname"`
			OsName        string `xml:"os-name"`
			WorkingDir    string `xml:"working-dir"`
			BuildDatabase string `xml:"build-database"`
			BuildPlatform string `xml:"build-platform"`
			BuildDate     string `xml:"build-date"`
		} `xml:"configuration-manager"`
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
		PmserverCollectorStats struct {
			Text                   string `xml:",chardata"`
			PacketsReceived        string `xml:"packets-received"`
			BytesReceived          string `xml:"bytes-received"`
			ReceiveErrors          string `xml:"receive-errors"`
			MaximumPacketSize      string `xml:"maximum-packet-size"`
			MinimumPacketSize      string `xml:"minimum-packet-size"`
			PacketSequenceErrors   string `xml:"packet-sequence-errors"`
			NumberOfMissingPackets string `xml:"number-of-missing-packets"`
			InvalidPackets         string `xml:"invalid-packets"`
			CollectorUdpPort       string `xml:"collector-udp-port"`
			CollectorBufferSize    string `xml:"collector-buffer-size"`
			DataStoreType          string `xml:"data-store-type"`
		} `xml:"pmserver-collector-stats"`
		PmserverProcessStats []struct {
			Text                     string `xml:",chardata"`
			SendingProcessName       string `xml:"sending-process-name"`
			SendingProcessID         string `xml:"sending-process-id"`
			CurrentSequenceNumber    string `xml:"current-sequence-number"`
			InboundPacketsReceived   string `xml:"inbound-packets-received"`
			InboundPacketErrors      string `xml:"inbound-packet-errors"`
			InboundBytesReceived     string `xml:"inbound-bytes-received"`
			PacketSequenceErrors     string `xml:"packet-sequence-errors"`
			NumberOfMissingPackets   string `xml:"number-of-missing-packets"`
			NumberOfDroppedPackets   string `xml:"number-of-dropped-packets"`
			NumberOfWriteOperations  string `xml:"number-of-write-operations"`
			NumberOfAppendOperations string `xml:"number-of-append-operations"`
			NumberOfWriteErrors      string `xml:"number-of-write-errors"`
			NumberOfAppendErrors     string `xml:"number-of-append-errors"`
			NumberOfOperationErrors  string `xml:"number-of-operation-errors"`
			MaximumPacketSize        string `xml:"maximum-packet-size"`
			TimesThreadStarted       string `xml:"times-thread-started"`
			WorkerThreadID           string `xml:"worker-thread-id"`
		} `xml:"pmserver-process-stats"`
	} `xml:"process"`
}
