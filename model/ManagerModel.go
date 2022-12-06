// ManagerModel (child)
//   - Định nghĩa struct ManagerModel dùng để ánh xạ các field trong xml thành Object.
package model

import "encoding/xml"

// Định nghĩa struct ManagerModel
type ManagerModel struct {
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
	} `xml:"process"`
}
