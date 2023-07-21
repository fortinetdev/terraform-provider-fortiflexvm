package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func createUpdate(c *FortiSDKClient, method string, path string, rspKey string, params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	result, string_body, err := sendRequest(c, method, path, params)
	if err != nil {
		return nil, err
	}

	err = fortiAPIErrorFormat(result, string_body)
	if err == nil {
		if rspKey == "entitlements" && result[rspKey] == nil { // TMP FortiFlex API BUG fix
			rspKey = "vms"
		}
		if result[rspKey] != nil {
			if confMap, ok := result[rspKey].(map[string]interface{}); ok {
				mapTmp = confMap
				return
			} else if confList, ok := result[rspKey].([]interface{}); ok {
				if len(confList) > 1 {
					err = fmt.Errorf("Response contains multiple values: %v", len(confList))
					log.Printf("[WARN] Response contains multiple values.")
					return
				}
				if confMap, ok := confList[0].(map[string]interface{}); ok {
					mapTmp = confMap
					return
				}
			} else {
				log.Printf("[WARN] Could not parse response type: %T", result[rspKey])
			}
		}
	}
	return
}

func delete(c *FortiSDKClient, method string, path string) (err error) {
	result, string_body, err := sendRequest(c, method, path, nil)
	if err != nil {
		return err
	}
	err = fortiAPIErrorFormat(result, string_body)
	return
}

func read(c *FortiSDKClient, method string, path string, bcomplex bool, rspKey string, params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	result, string_body, err := sendRequest(c, method, path, params)
	if err != nil {
		return nil, err
	}
	if fortiAPIHttpStatus404Checking(result) == true {
		mapTmp = nil
		return
	}
	err = fortiAPIErrorFormat(result, string_body)
	if err == nil { //!!! need check whether all response is in rspKey argument
		mapTmp = map[string]interface{}{}
		if bcomplex {
			new_rspKey := rspKey
			if rspKey == "entitlements" && result[rspKey] == nil { // TMP FortiFlex API BUG fix
				new_rspKey = "vms"
			}
			mapTmp[rspKey] = result[new_rspKey]
		} else {
			mapTmp = (result[rspKey].([]interface{}))[0].(map[string]interface{})
		}
	}
	return
}

func sendRequest(c *FortiSDKClient, method string, path string, params *map[string]interface{}) (result map[string]interface{}, string_body string, err error) {
	var locJSON []byte
	if params != nil {
		locJSON, err = json.Marshal(params)
		log.Printf("FortiFlex request body: %s", string(locJSON))
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
		req := c.NewRequest(method, path, nil, bytePara)
		err = req.Send()
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
		if result["status"] != nil {
			rtStatus := fmt.Sprintf("%v", result["status"])
			if rtStatus != "0" {
				log.Printf("[Error] error response, will resend again %s, %v", path, result)
				retry++
				time.Sleep(time.Second)
				continue
			}
		}
		break
	}
	log.Printf("FortiFlex reading response %s: %s", path, string(body))
	return result, string(body), err
}
