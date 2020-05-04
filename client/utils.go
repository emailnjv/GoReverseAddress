package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	ReponseError   = "error calling %s, got status code %d"
	RequestRetries = 10
)

type Callout struct {}

func NewCallOut() (Callout, error) {
	var result Callout
	var err error
	return result, err
}

func (c Callout) Call(endpoint string) (string, error) {
	requestCounter := 0
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	for requestCounter != RequestRetries && resp.StatusCode != http.StatusOK {
		resp, err = http.Get(endpoint)
		if err != nil {
			return "", err
		}
		if resp.StatusCode != http.StatusOK {
			requestCounter += 1
			time.Sleep(5 * time.Second)
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(ReponseError, endpoint, resp.StatusCode)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}