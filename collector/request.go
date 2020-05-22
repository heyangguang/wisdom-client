package collector

import (
	"bytes"
	"encoding/json"
	"net/http"
	"wisdom-client/wisdom-client/logger"
)

// 请求后端
func httpRequest(url, secret string, obj interface{}) error {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(structPareJson(obj)))
	if err != nil {
		return err
	}
	// 设置头信息
	request.Header.Add("Authorization", "Bearer " + secret)
	request.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	_, err = client.Do(request)
	return err
}

// struct转json
func structPareJson(obj interface{}) []byte {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		logger.Error("structPareJson error, err:" + err.Error())
	}
	return jsonBytes
}
