package model

import "encoding/xml"

// Parent
type GroupsModel struct {
	XMLName                   xml.Name   `xml:"groups"`
	Text                      string     `xml:",chardata"`
	Xsi                       string     `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string     `xml:"noNamespaceSchemaLocation,attr"`
	GroupRefs                 []GroupRep `xml:"group-ref"`
}

func (m *GroupRep) IsExtract() bool {
	return m.Type == TYPE_EXTRACT
}
func (m *GroupRep) IsPump() bool {
	return m.Type == TYPE_PUMP
}
func (m *GroupRep) IsManager() bool {
	return m.Type == TYPE_MGR
}
func (m *GroupRep) IsPerformanceServer() bool {
	return m.Type == TYPE_PMSRVR
}
func (m *GroupRep) IsReplicat() bool {
	return m.Type == TYPE_REPLICAT
}

// Child

type GroupRep struct {
	Text        string `xml:",chardata"`
	Name        string `xml:"name,attr"`
	Type        string `xml:"type,attr"` // 1 is mgr, 2 is
	Status      string `xml:"status,attr"`
	URL         string `xml:"url"`
	Description struct {
		Text string `xml:",chardata"`
		Nil  string `xml:"nil,attr"`
	} `xml:"description"`
}
