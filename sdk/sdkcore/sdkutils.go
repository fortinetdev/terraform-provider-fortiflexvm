package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	request "github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/request"
)

func createUpdate(client *FortiSDKClient, method string, path string, rspKey string, params *map[string]interface{}) (map[string]interface{}, error) {
	result, string_body, err := sendRequest(client, method, path, params)
	if err != nil {
		return nil, err
	}
	err = fortiAPIErrorFormat(result, string_body)
	if err != nil {
		return nil, err
	}
	mapTmp := map[string]interface{}{}
	if rspKey == "entitlements" && result[rspKey] == nil { // TMP FortiFlex API BUG fix
		rspKey = "vms"
	}
	if result[rspKey] != nil {
		if confMap, ok := result[rspKey].(map[string]interface{}); ok {
			return confMap, nil
		} else if confList, ok := result[rspKey].([]interface{}); ok {
			if len(confList) > 1 {
				err = fmt.Errorf("Response contains multiple values: %v", len(confList))
				log.Printf("[WARN] Response contains multiple values.")
				return nil, err
			}
			if confMap, ok := confList[0].(map[string]interface{}); ok {
				return confMap, nil
			}
		} else {
			log.Printf("[WARN] Could not parse response type: %T", result[rspKey])
		}
	}
	return mapTmp, nil
}

func read(client *FortiSDKClient, method string, path string, rspKey string, params *map[string]interface{}) (map[string]interface{}, error) {
	result, string_body, err := sendRequest(client, method, path, params)
	if err != nil {
		return nil, err
	}
	err = fortiAPIErrorFormat(result, string_body)
	if err != nil {
		return nil, err
	}
	mapTmp := map[string]interface{}{}
	new_rspKey := rspKey
	if rspKey == "entitlements" && result[rspKey] == nil { // TMP FortiFlex API BUG fix
		new_rspKey = "vms"
	}
	mapTmp[rspKey] = result[new_rspKey]
	return mapTmp, nil
}

func sendRequest(client *FortiSDKClient, method string, path string, params *map[string]interface{}) (result map[string]interface{}, string_body string, err error) {
	var locJSON []byte
	if params != nil {
		locJSON, err = json.Marshal(params)
		if err != nil {
			log.Fatal(err)
			return nil, "", err
		}
	}
	retry := 0
	var body []byte
	for retry < 4 {
		var bytePara *bytes.Buffer
		if locJSON != nil {
			bytePara = bytes.NewBuffer(locJSON)
		}
		log.Printf("[INFO] Request '%s' | %s", path, string(locJSON))
		req := request.NewRequest(client.Auth, client.HTTPCon, method, path, nil, bytePara)
		err = req.Send(5) // If the connection fails, retry up to 5 times
		if err != nil || req.HTTPResponse == nil {
			err = fmt.Errorf("cannot send request: %v", err)
			return nil, "", err
		}

		body, err = ioutil.ReadAll(req.HTTPResponse.Body)
		req.HTTPResponse.Body.Close()
		if err != nil || body == nil {
			err = fmt.Errorf("cannot get response body: %v", err)
			return nil, "", err
		}

		json.Unmarshal([]byte(string(body)), &result)
		if result["status"] != nil { // Retry for FortiFlex API error
			rtStatus := fmt.Sprintf("%v", result["status"])
			if rtStatus != "0" {
				log.Printf("[ERROR] Response '%s' | retry time %v | %v", path, retry, result)
				retry++
				time.Sleep(time.Second)
				continue
			}
		}
		break
	}
	return result, string(body), err
}
