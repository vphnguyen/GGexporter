package storage

import (
	
    "encoding/xml"
    "GGexporter/entities"
)

func GetGGRunningInstances (  groups *entities.MGroups ){
	data, _ := GetHTTPToXMLbytes(entities.RootURL+"/groups")
    xml.Unmarshal(data, &groups)
}

func GetURLOfInstances(groups *entities.MGroups) []string {
	var instanceURLs [] string
	for _,value := range (groups.GroupRefs) {
		instanceURLs= append(instanceURLs,entities.RootURL+value.URL)
	}
	return instanceURLs
}


// Extract 
func GetGGRunningExtractInstances ( groups *entities.MGroups , mpointsofextract *[]entities.MpointsOfExtract ){

	for _,aGroup := range (groups.GroupRefs) {
		if aGroup.Type == entities.TYPE_EXTRACT {
			var tempExtract  entities.MpointsOfExtract
			data, _ := GetHTTPToXMLbytes( entities.RootURL + aGroup.URL + "/mpointsx/")	
			xml.Unmarshal(data, &tempExtract)
			*mpointsofextract = append (*mpointsofextract, tempExtract )
		}
	} 

}

// MGR 
func GetGGRunningMGRInstances ( groups *entities.MGroups , mpointsofmgr *entities.MpointsOfMGR ){

	for _,aGroup := range (groups.GroupRefs) {
		if aGroup.Type == entities.TYPE_MGR {
			data, _ := GetHTTPToXMLbytes( entities.RootURL + aGroup.URL+ "/mpointsx/")	
			xml.Unmarshal(data, &mpointsofmgr)

		}
	} 

}


// PMSRVR
func GetGGRunningPMSRVRInstances ( groups *entities.MGroups , mpointsofpmsrvr *entities.MpointsOfPMSRVR ){

	for _,aGroup := range (groups.GroupRefs) {
		if aGroup.Type == entities.TYPE_PMSRVR {
			data, _ := GetHTTPToXMLbytes( entities.RootURL + aGroup.URL+ "/mpointsx/")	
			xml.Unmarshal(data, &mpointsofpmsrvr)

		}
	} 

}
