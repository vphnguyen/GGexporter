package services


import (
	
	"time"
	"net/http"
	"io/ioutil" 
	"fmt"
)

func GetHTTPToXMLbytes(url string) ([]byte, error) {
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

