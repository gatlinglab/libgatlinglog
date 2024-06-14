package ilog

import (
	"bytes"
	"errors"
	"net/http"
)

type CGatlingLogClient struct {
	endpoint    string
	token       string
	urlInfo     string
	urlError    string
	urlAppInfo  string
	urlAppError string
}

func GetNewGatlingLogClient() *CGatlingLogClient {
	return &CGatlingLogClient{}
}

func (pInst *CGatlingLogClient) Initialize(url, token, appid string) error {
	pInst.endpoint = url
	pInst.token = token
	if url[len(url)-1] != '/' {
		pInst.endpoint = url + "/"
	}

	pInst.urlInfo = pInst.endpoint + "info/"
	pInst.urlError = pInst.endpoint + "error/" + appid
	pInst.urlAppInfo = pInst.urlInfo + appid
	pInst.urlAppError = pInst.urlError + appid

	return nil
}

func (pInst *CGatlingLogClient) LogInfo(log string) error {
	return pInst.logPut(pInst.urlAppInfo, log)
}
func (pInst *CGatlingLogClient) LogError(log string) error {
	return pInst.logPut(pInst.urlAppError, log)
}

func (pInst *CGatlingLogClient) logPut(url, log string) error {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(log)))
	if err != nil {
		return err
	}
	if pInst.token != "" {
		req.Header.Set("X-API-KEY", pInst.token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		//fmt.Println("success")
	} else {
		//fmt.Println("failed: ", resp)
		return errors.New("failed")
	}

	return nil
}
