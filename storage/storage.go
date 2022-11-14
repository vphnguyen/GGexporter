package storage

/*
*
*	Chuyen reuqest noi dung xml tu URL va chuyen thanh day cac bytes
 *
*/
import (
	"GGexporter/model"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//===== GET GROUPS
func GetGroups(url string) (*model.GroupsModel, error) {
	var gr model.GroupsModel
	data, err := fetch(url + "/groups")
	if err != nil {
		return nil, errors.New("Storage - khong the fetch - groups - Vui long kiem tra:  " + url + "/groups")
	}
	xml.Unmarshal(data, &gr)
	return &gr, nil
}

//===== GET UNITS OF GROUPS
func GetPump(root string, branch string) (*model.PumpModel, error) {
	var aPump model.PumpModel
	data, err := fetch(root + branch + "/mpointsx/")
	if err != nil {
		return nil, errors.New("Storage - khong the fetch - pump - Vui long kiem tra:  " + root + branch + "/mpointsx/")
	}
	xml.Unmarshal(data, &aPump)
	return &aPump, nil
}

func GetExtract(root string, branch string) (model.ExtractModel, error) {
	var anExtract model.ExtractModel
	data, err := fetch(root + branch + "/mpointsx/")
	if err != nil {
		item := model.ExtractModel{}
		return item, errors.New("Storage - khong the fetch - extract - Vui long kiem tra:  " + root + branch + "/mpointsx/")
	}
	xml.Unmarshal(data, &anExtract)
	return anExtract, nil
}

func GetManager(root string, branch string) (model.ManagerModel, error) {
	var anExtract model.ManagerModel
	data, err := fetch(root + branch + "/mpointsx/")
	if err != nil {
		item := model.ManagerModel{}
		return item, errors.New("Storage - khong the fetch - mgr - Vui long kiem tra:  " + root + branch + "/mpointsx/")
	}
	xml.Unmarshal(data, &anExtract)
	return anExtract, nil
}

func GetPerformanceServer(rootURL string, branchURL string) (model.PerformanceServerModel, error) {
	var anPerformanceServer model.PerformanceServerModel
	data, err := fetch(rootURL + branchURL + "/mpointsx/")
	if err != nil {
		item := model.PerformanceServerModel{}
		return item, errors.New("Storage - khong the fetch - pmsrvr - Vui long kiem tra:  " + rootURL + branchURL + "/mpointsx/")
	}
	xml.Unmarshal(data, &anPerformanceServer)
	return anPerformanceServer, nil
}

func GetReplicat(rootURL string, branchURL string) (model.ReplicatModel, error) {
	var aReplicat model.ReplicatModel
	data, err := fetch(rootURL + branchURL + "/mpointsx/")
	if err != nil {
		item := model.ReplicatModel{}
		return item, errors.New("Storage - khong the fetch - Replicat - Vui long kiem tra:  " + rootURL + branchURL + "/mpointsx/")
	}
	xml.Unmarshal(data, &aReplicat)
	return aReplicat, nil
}

//===== FETCH DATA
func fetch(url string) ([]byte, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return []byte(""), fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte(""), fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), fmt.Errorf("Read body: %v", err)
	}
	return []byte(data), nil
}
