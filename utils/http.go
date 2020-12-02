// 通用工具库
package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type KV map[string]interface{}

//
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

//
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

//
func DoUpload(url, filePath string, kvs KV) (body []byte, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("media", filepath.Base(filePath))
	if err != nil {
		return
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return
	}
	contentType := bodyWriter.FormDataContentType()
	for k, v := range kvs {
		out, _ := json.Marshal(v)
		bodyWriter.WriteField(k, string(out))
	}
	bodyWriter.Close()
	var resp *http.Response
	resp, err = http.Post(url, contentType, bodyBuf)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
