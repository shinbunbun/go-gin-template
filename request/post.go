package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Post(url string, contentType string, body []byte) ([]byte, int, error) {
	resp, err := http.Post(url, contentType, bytes.NewBuffer(body))
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
