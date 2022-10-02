package request

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, err
	}

	return byteArray, resp.StatusCode, nil
}
