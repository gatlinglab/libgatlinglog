package libgatlinglog

import ilog "github.com/gatlinglab/libgatlinglog/internal"

type GatLingILogClient interface {
	LogInfo(log string) error
	LogError(log string) error
}

var defaultLog GatLingILogClient = ilog.GetTempLogClient()

func GatlingLogLib_Initialize(url, token, appid string) (GatLingILogClient, error) {
	client1 := ilog.GetNewGatlingLogClient()
	err := client1.Initialize(url, token, appid)
	if err != nil {
		return nil, err
	}

	return client1, nil
}
func GatlingLogLib_SetDefault(logInst GatLingILogClient) {
	defaultLog = logInst
}

func GatlingLogLib_info(log string) error {
	return defaultLog.LogInfo(log)
}
func GatlingLogLib_error(log string) error {
	return defaultLog.LogError(log)
}
