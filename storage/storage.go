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
		return nil, errors.New("Storage - GetGroups - Khong the fetch - Vui long kiem tra:  " + rootURL + "/groups")
	}
	xml.Unmarshal(data, &gr)
	return &gr, nil
}

// Nếu group đó là aPump thì truyền vào địa chỉ của performanceServer và đường dẫn đến group đó.
func GetPump(rootURL string, branch string) (*model.PumpModel, error) {
	var aPump model.PumpModel
	data, err := fetch(rootURL + branch + "/mpoints")
	if err != nil {
		return nil, errors.New("Storage - Pump - Fetch failed - Check: " + rootURL + branch + "/mpoints/")
	}
	if er := xml.Unmarshal(data, &aPump); er != nil {
		return nil, errors.New("Storage - Pump - Fetched - Unmarshal error.")
	}
	if aPump.IsANewOne() {
		return nil, errors.New("Storage - Pump - Fetched - Vua duoc tao.")
	}
	return &aPump, nil
}

// Nếu group đó là anExtract thì truyền vào địa chỉ của performanceServer và đường dẫn đến group đó.
func GetExtract(rootURL string, branch string) (*model.ExtractModel, error) {
	var anExtract model.ExtractModel
	data, err := fetch(rootURL + branch + "/mpoints")
	if err != nil {
		return nil, errors.New("Storage - Extract - Fetch failed - Check: " + rootURL + branch + "/mpointsx/")
	}
	if er := xml.Unmarshal(data, &anExtract); er != nil {
		return nil, errors.New("Storage - Extract - Fetched - Unmarshal error.")
	}
	if anExtract.IsInitLoad() {
		return nil, errors.New("Storage - Extract - Could be an Initload.")
	}
	if anExtract.IsANewOne() {
		return nil, errors.New("Storage - Extract - Fetched - Just created.")
	}
	return &anExtract, nil
}

// Nếu group đó là anExtract thì truyền vào địa chỉ của performanceServer và đường dẫn đến group đó.
func GetManager(rootURL string, branch string, aManager *model.ManagerModel) error {
	data, err := fetch(rootURL + branch + "/mpoints")
	if err != nil {
		return errors.New("Storage - Manager - Fetch failed - Check: " + rootURL + branch + "/mpoints/" + err.Error())
	}
	if er := xml.Unmarshal(data, aManager); er != nil {
		return errors.New("Storage - Manager - Fetched - Unmarshal error.")
	}
	return nil
}

func GetPerformanceServer(rootURL string, branch string, anPerformanceServer *model.PerformanceServerModel) error {
	data, err := fetch(rootURL + branch + "/mpoints")
	if err != nil {
		return errors.New("Storage - PMSRVR - Fetch failed - Check: " + rootURL + branch + "/mpoints/")
	}
	if er := xml.Unmarshal(data, &anPerformanceServer); er != nil {
		return errors.New("Storage - PMSRVR - Fetched - Unmarshal error.")
	}
	return nil
}

// Nếu group đó là aReplicat thì truyền vào địa chỉ của performanceServer và đường dẫn đến group đó.
func GetReplicat(rootURL string, branch string) (*model.ReplicatModel, error) {
	var aReplicat model.ReplicatModel
	data, err := fetch(rootURL + branch + "/mpoints")
	if err != nil {
		return nil, errors.New("Storage - Replicat - Fetch failed - Check: " + rootURL + branch + "/mpoints/")
	}
	if er := xml.Unmarshal(data, &aReplicat); er != nil {
		return nil, errors.New("Storage - Replicat - Fetched - Unmarshal error.")
	}
	if aReplicat.IsInitLoad() {
		return nil, errors.New("Storage - Replicat - Could be an Initload.")
	}
	if aReplicat.IsANewOne() {
		return nil, errors.New("Storage - Replicat - Fetched - Just created.")
	}
	return &aReplicat, nil
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
