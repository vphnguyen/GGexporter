package storage

import (
    log "github.com/sirupsen/logrus"
    "encoding/xml"
    "GGexporter/entities"
)

func GetGGRunningInstances (  groups *entities.MGroups ){
	data, er := GetHTTPToXMLbytes(entities.RootURL+"/groups")
    xml.Unmarshal(data, &groups)
    if er != nil {
			log.Errorf("Storage. Khong the get Groups")
	}
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
			data, er := GetHTTPToXMLbytes( entities.RootURL + aGroup.URL + "/mpointsx/")
			if er != nil {
					log.Errorf("Storage. Khong the lay thong tin xml Extract")
			}	
			xml.Unmarshal(data, &tempExtract)
			*mpointsofextract = append (*mpointsofextract, tempExtract )
		}
	} 
}
// Pump 
func GetGGRunningPumpInstances ( groups *entities.MGroups , mpointsofpump *[]entities.MpointsOfPump ){	
	for _,aGroup := range (groups.GroupRefs) {
		if aGroup.Type == entities.TYPE_PUMP {
			var tempPump  entities.MpointsOfPump
			data, er := GetHTTPToXMLbytes( entities.RootURL + aGroup.URL + "/mpointsx/")
			if er != nil {
					log.Errorf("Storage. Khong the lay thong tin xml Pump")
			}		
			xml.Unmarshal(data, &tempPump)
			*mpointsofpump = append (*mpointsofpump, tempPump )
		}
	} 
}
// MGR 
func GetGGRunningMGRInstances ( groups *entities.MGroups , mpointsofmgr *entities.MpointsOfMGR ){
	
	for _,aGroup := range (groups.GroupRefs) {
		if aGroup.Type == entities.TYPE_MGR {
			data, er := GetHTTPToXMLbytes( entities.RootURL + aGroup.URL+ "/mpointsx/")
			if er != nil {
					log.Errorf("Storage. Khong the lay thong tin xml MGR")
			}		
			xml.Unmarshal(data, &mpointsofmgr)
		}
	} 
}


// PMSRVR
func GetGGRunningPMSRVRInstances ( groups *entities.MGroups , mpointsofpmsrvr *entities.MpointsOfPMSRVR ){
	
	for _,aGroup := range (groups.GroupRefs) {
		if aGroup.Type == entities.TYPE_PMSRVR {
			data, er := GetHTTPToXMLbytes( entities.RootURL + aGroup.URL+ "/mpointsx/")
			if er != nil {
					log.Errorf("Storage. Khong the lay thong tin xml PMSRVR")
			}		
			xml.Unmarshal(data, &mpointsofpmsrvr)

		}
	} 

}

// REPLICAT 
func GetGGRunningReplicatInstances ( groups *entities.MGroups , mpointsofreplicat *[]entities.MpointsOfReplicat ){		
	for _,aGroup := range (groups.GroupRefs) {
		if aGroup.Type == entities.TYPE_REPLICAT {
			var tempReplicat  entities.MpointsOfReplicat
			data, er := GetHTTPToXMLbytes( entities.RootURL + aGroup.URL + "/mpointsx/")
			if er != nil {
					log.Errorf("Storage. Khong the lay thong tin xml REPLICAT")
			}		
			xml.Unmarshal(data, &tempReplicat)
			*mpointsofreplicat = append (*mpointsofreplicat, tempReplicat )
		}
	} 
}
