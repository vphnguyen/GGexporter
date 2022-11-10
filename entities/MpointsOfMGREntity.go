package entities


import  "encoding/xml"


type MpointsOfMGR struct {
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
	} `xml:"process"`
} 
