package entities


import  "encoding/xml"

// Parent
type MGroups struct {
    XMLName                   xml.Name   `xml:"groups"`
    Text                      string     `xml:",chardata"`
    Xsi                       string     `xml:"xsi,attr"`
    NoNamespaceSchemaLocation string     `xml:"noNamespaceSchemaLocation,attr"`
    GroupRefs                 []GroupRep `xml:"group-ref"`
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