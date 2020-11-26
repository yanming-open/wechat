package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type KV map[string]interface{}

func DoGet(url string) ([]byte, error) {
	client := http.Client{Timeout: time.Second * 15}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		} else {
			return body, nil
		}
	}
}
func DoPost(url string, params KV) ([]byte, error) {
	bytesData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	client := http.Client{Timeout: time.Second * 15}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bytesData))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		} else {
			return body, nil
		}
	}
}
