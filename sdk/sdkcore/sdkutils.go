package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func createUpdate(c *FortiSDKClient, method string, path string, rspKey string, params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest(method, path, nil, bytes)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request: %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
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

	req := c.NewRequest(method, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request: %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}
	log.Printf("FOS-flexvm delete response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

func read(c *FortiSDKClient, method string, path string, bcomplex bool, rspKey string, params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	var bytePara *bytes.Buffer
	if params != nil {
		locJSON, err := json.Marshal(params)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		bytePara = bytes.NewBuffer(locJSON)
	} else {
		bytePara = nil
	}
	req := c.NewRequest(method, path, nil, bytePara)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}
	log.Printf("Flexvm reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if fortiAPIHttpStatus404Checking(result) == true {
		mapTmp = nil
		return
	}

	err = fortiAPIErrorFormat(result, string(body))
	if err == nil { //!!! need check whether all response is in rspKey argument
		mapTmp = map[string]interface{}{}
		if bcomplex {
			mapTmp[rspKey] = result[rspKey]
		} else {
			mapTmp = (result[rspKey].([]interface{}))[0].(map[string]interface{})
		}
	}
	return
}
