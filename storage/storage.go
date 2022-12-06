// Storage fetch dữ liệu từ các url và chuyển thành bytes[].
//
// Nếu một group extract có url là: http://gg-svmgr.io/groups/EXTF1/mpoints
//
// Thì rootURL: http://gg-svmgr.io/groups/ | branch: EXTF1
package storage

import (
	"GGexporter/model"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Lấy thông tin các group đã cấu hình trên GG.
func GetGroups(rootURL string) (*model.GroupsModel, error) {
	var gr model.GroupsModel
	data, err := fetch(rootURL + "/groups")
	if err != nil {
		return nil, errors.New("Storage - khong the fetch - groups - Vui long kiem tra:  " + rootURL + "/groups")
	}
	xml.Unmarshal(data, &gr)
	return &gr, nil
}

// Nếu group đó là aPump thì truyền vào địa chỉ của performanceServer và đường dẫn đến group đó.
func GetPump(rootURL string, branch string) (*model.PumpModel, error) {
	var aPump model.PumpModel
	data, err := fetch(rootURL + branch + "/mpoints")
	if err != nil {
		return nil, errors.New("Storage - khong the fetch - pump - Vui long kiem tra:  " + rootURL + branch + "/mpoints/")
	}
	xml.Unmarshal(data, &aPump)
	return &aPump, nil
}

// Nếu group đó là anExtract thì truyền vào địa chỉ của performanceServer và đường dẫn đến group đó.
func GetExtract(rootURL string, branch string) (model.ExtractModel, error) {
	var anExtract model.ExtractModel
	data, err := fetch(rootURL + branch + "/mpoints")
	if err != nil {
		item := model.ExtractModel{}
		return item, errors.New("Storage - khong the fetch - extract - Vui long kiem tra:  " + rootURL + branch + "/mpointsx/")
	}
	xml.Unmarshal(data, &anExtract)
	if anExtract.IsInit() {
		item := model.ExtractModel{}
		return item, errors.New("Storage - Extract - Founded unknown group at:  " + rootURL + branch + "/mpointsx/")
	}
	return anExtract, nil
}

// Nếu group đó là anExtract thì truyền vào địa chỉ của performanceServer và đường dẫn đến group đó.
func GetManager(rootURL string, branch string) (model.ManagerModel, error) {
	var anExtract model.ManagerModel
	data, err := fetch(rootURL + branch + "/mpoints")
	if err != nil {
		item := model.ManagerModel{}
		return item, errors.New("Storage - khong the fetch - mgr - Vui long kiem tra:  " + rootURL + branch + "/mpointsx/")
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

// Nếu group đó là aReplicat thì truyền vào địa chỉ của performanceServer và đường dẫn đến group đó.
func GetReplicat(rootURL string, branchURL string) (model.ReplicatModel, error) {
	var aReplicat model.ReplicatModel
	data, err := fetch(rootURL + branchURL + "/mpoints")
	if err != nil {
		item := model.ReplicatModel{}
		return item, errors.New("Storage - khong the fetch - Replicat - Vui long kiem tra:  " + rootURL + branchURL + "/mpoints")
	}
	xml.Unmarshal(data, &aReplicat)
	if aReplicat.IsInit() {
		item := model.ReplicatModel{}
		return item, errors.New("Storage - Replicat - Founded unknown group at:  " + rootURL + branchURL + "/mpoints")
	}
	return aReplicat, nil
}

// ===== FETCH DATA
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
