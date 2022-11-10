package storage

	/*
	*	Chuyen reuqest noi dung xml tu URL va chuyen thanh day cac bytes
	*/

import (

	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/xml"
	"time"
    "GGexporter/entities"
)

func GetGroups() (*entities.MGroups, error) {
	var gr entities.MGroups
	data, err := fetch(config.RootURL + "/groups")
	if err != nil {
		return nil, errors.New("Storage - khong the fetch - groups")
	}
	xml.Unmarshal(data, &gr)
	return &gr, nil
}

func GetPump(url string) (*entities.MpointsOfPump, error) {
	var mpointsofpump *entities.MpointsOfPump
	data, err := fetch(config.RootURL + "/groups")
	if err != nil {
		return nil, errors.New("Storage - khong the fetch - pump")
	}
	xml.Unmarshal(data, &mpointsofpump)
	return &mpointsofpump, nil
}




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


