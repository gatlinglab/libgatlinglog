package ilog

import "errors"

type tmpGatLingILogClient struct {
}

var g_SingleTempClient = &tmpGatLingILogClient{}
var errTmp = errors.New("this is temp client")

func GetTempLogClient() *tmpGatLingILogClient {
	return g_SingleTempClient
}

func (pInst *tmpGatLingILogClient) LogInfo(log string) error {
	return errTmp
}
func (pInst *tmpGatLingILogClient) LogError(log string) error {
	return errTmp
}
