// GroupsModel (parent)
//   - Định nghĩa struct GroupsModel dùng để ánh xạ các field trong xml thành Object.
package model

import "encoding/xml"

// Định nghĩa struct GroupsModel
type GroupsModel struct {
	XMLName                   xml.Name   `xml:"groups"`
	Text                      string     `xml:",chardata"`
	Xsi                       string     `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string     `xml:"noNamespaceSchemaLocation,attr"`
	GroupRefs                 []GroupRep `xml:"group-ref"`
}

// Là một phương thức của GroupsRep dùng để kiểm tra Group này có đúng Extract hay không.
func (m *GroupRep) IsExtract() bool {
	return m.Type == TYPE_EXTRACT
}

// Là một phương thức của GroupsRep dùng để kiểm tra Group này có đúng Pump hay không.
func (m *GroupRep) IsPump() bool {
	return m.Type == TYPE_PUMP
}

// Là một phương thức của GroupsRep dùng để kiểm tra Group này có đúng Manager hay không.
func (m *GroupRep) IsManager() bool {
	return m.Type == TYPE_MGR
}

// Là một phương thức của GroupsRep dùng để kiểm tra Group này có đúng PerformanceServer hay không.
func (m *GroupRep) IsPerformanceServer() bool {
	return m.Type == TYPE_PMSRVR
}

// Là một phương thức của GroupsRep dùng để kiểm tra Group này có đúng Replicat hay không.
func (m *GroupRep) IsReplicat() bool {
	return m.Type == TYPE_REPLICAT
}

// Định nghĩa struct GroupRep nằm trong GroupsModel
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
