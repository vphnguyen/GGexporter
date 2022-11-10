package storage

import (
	"GGexporter/entities"
	"encoding/xml"
	log "github.com/sirupsen/logrus"
	
)

/*
*
*	Lay du lieu cac groups
*	=> Moi group dua vao type se chuyen thanh cac MonitorPoint tuong ung
*	[extract, pump, mgr, pmsrvr, replicat]
*
 */
var config entities.Config

func GetGGRunningInstances(c entities.Config, groups *entities.MGroups,
	mpointsofextract *[]entities.MpointsOfExtract,
	mpointsofpump *[]entities.MpointsOfPump,
	mpointsofmgr *entities.MpointsOfMGR,
	mpointsofpmsrvr *entities.MpointsOfPMSRVR,
	mpointsofreplicat *[]entities.MpointsOfReplicat) {
	config = c
	//Lay du lieu cac groups
	data, er := fetch(config.RootURL + "/groups")
	if er != nil {
		log.Errorf("Storage. Khong the get Groups")
	}
	xml.Unmarshal(data, &groups)
	

	// Chuyen moi group thanh cac Mpoint tuong ung (Cac ham phia duoi)
	for _, aGroup := range groups.GroupRefs {
		if aGroup.Type == entities.TYPE_EXTRACT {
			appendToExtract(aGroup, mpointsofextract)
			continue
		}
		if aGroup.Type == entities.TYPE_PUMP {
			appendToPump(aGroup, mpointsofpump)
			continue
		}
		if aGroup.Type == entities.TYPE_MGR {
			appendToMGR(aGroup, mpointsofmgr)
			continue
		}
		if aGroup.Type == entities.TYPE_PMSRVR {
			appendToPMSRVR(aGroup, mpointsofpmsrvr)
			continue
		}
		if aGroup.Type == entities.TYPE_REPLICAT {
			appendToReplicat(aGroup, mpointsofreplicat)
			continue
		}
	}
}

// Chuyen doi dua theo Type
func appendToExtract(aGroup entities.GroupRep, mpointsofextract *[]entities.MpointsOfExtract) {
	var tempExtract entities.MpointsOfExtract
	data, er := fetch(config.RootURL + aGroup.URL + "/mpointsx/")
	if er != nil {
		log.Errorf("Storage. Khong the lay thong tin xml Extract")
	}
	xml.Unmarshal(data, &tempExtract)
	*mpointsofextract = append(*mpointsofextract, tempExtract)
}

func appendToPump(aGroup entities.GroupRep, mpointsofpump *[]entities.MpointsOfPump) {
	var tempPump entities.MpointsOfPump
	data, er := fetch(config.RootURL + aGroup.URL + "/mpointsx/")
	if er != nil {
		log.Errorf("Storage. Khong the lay thong tin xml Pump")
	}
	xml.Unmarshal(data, &tempPump)
	*mpointsofpump = append(*mpointsofpump, tempPump)
}

func appendToMGR(aGroup entities.GroupRep, mpointsofmgr *entities.MpointsOfMGR) {
	data, er := fetch(config.RootURL + aGroup.URL + "/mpointsx/")
	if er != nil {
		log.Errorf("Storage. Khong the lay thong tin xml MGR")
	}
	xml.Unmarshal(data, &mpointsofmgr)
}

func appendToPMSRVR(aGroup entities.GroupRep, mpointsofpmsrvr *entities.MpointsOfPMSRVR) {
	data, er := fetch(config.RootURL + aGroup.URL + "/mpointsx/")
	if er != nil {
		log.Errorf("Storage. Khong the lay thong tin xml PMSRVR")
	}
	xml.Unmarshal(data, &mpointsofpmsrvr)
}

func appendToReplicat(aGroup entities.GroupRep, mpointsofreplicat *[]entities.MpointsOfReplicat) {
	var tempReplicat entities.MpointsOfReplicat
	data, er := fetch(config.RootURL + aGroup.URL + "/mpointsx/")
	if er != nil {
		log.Errorf("Storage. Khong the lay thong tin xml REPLICAT")
	}
	xml.Unmarshal(data, &tempReplicat)
	*mpointsofreplicat = append(*mpointsofreplicat, tempReplicat)
}
