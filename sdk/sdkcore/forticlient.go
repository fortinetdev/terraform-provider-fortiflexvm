package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/auth"
	"github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/config"
	"github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/request"
	// "strconv"
)

// MultValue describes the nested structure in the results
type MultValue struct {
	Name string `json:"name"`
}

// MultValues describes the nested structure in the results
type MultValues []MultValue

// FortiSDKClient describes the global FlexVM plugin client instance
type FortiSDKClient struct {
	Config  config.Config
	Retries int
}

// ExtractString extracts strings from result and put them into a string array,
// and return the string array
func ExtractString(members []MultValue) []string {
	vs := make([]string, 0, len(members))
	for _, v := range members {
		c := v.Name
		vs = append(vs, c)
	}
	return vs
}

func escapeURLString(v string) string { // doesn't support "<>()"'#"
	return strings.Replace(url.QueryEscape(v), "+", "%20", -1)
}

// NewClient initializes a new global plugin client
// It returns the created client object
func NewClient(auth *auth.Auth, client *http.Client) (*FortiSDKClient, error) {
	c := &FortiSDKClient{}

	c.Config.Auth = auth
	c.Config.HTTPCon = client

	return c, nil
}

// NewRequest creates the request to FlexVM for the client
// and return it to the client
func (c *FortiSDKClient) NewRequest(method string, path string, params interface{}, data *bytes.Buffer) *request.Request {
	return request.New(c.Config, method, path, params, data)
}

// GenerateToken() generate token from the Device
// It returns the token
func (c *FortiSDKClient) GenerateToken() error {
	var err error

	data := map[string]string{
		"username":   c.Config.Auth.Username,
		"password":   c.Config.Auth.Password,
		"client_id":  "flexvm",
		"grant_type": "password",
	}

	dataJson, err := json.Marshal(data)
	if err != nil {
		err = fmt.Errorf("Could not marshal login data: %s", err)
		return err
	}

	dataBytes := bytes.NewBuffer(dataJson)

	req := c.NewRequest("POST", "", nil, dataBytes)
	req.HTTPRequest.Header.Set("Content-Type", "application/json")
	u := "https://customerapiauth.fortinet.com/api/v1/oauth/token/"

	req.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		err = fmt.Errorf("Could not parse URL: %s", err)
		return err
	}

	rsp, err := req.Config.HTTPCon.Do(req.HTTPRequest)

	body, err := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["access_token"] != nil && result["access_token"] != "" {
			c.Config.Auth.Token = result["access_token"].(string)
		} else {
			err = fmt.Errorf("Can not get Token.")
			return err
		}
	}

	return nil
}

func fortiAPIHttpStatus404Checking(result map[string]interface{}) (b404 bool) {
	b404 = false

	if result != nil {
		if result["http_status"] != nil && result["http_status"] == 404.0 {
			b404 = true
			return
		}
	}

	return
}

func fortiAPIErrorFormat(result map[string]interface{}, body string) (err error) {
	if result != nil {
		if result["status"] != nil {
			rtStatus := fmt.Sprintf("%v", result["status"])
			if rtStatus == "success" || rtStatus == "0" {
				err = nil
				return
			}

			if result["http_status"] != nil {
				// 200	OK: Request returns successful
				if result["http_status"] == 400.0 {
					err = fmt.Errorf("Bad Request - Request cannot be processed by the API (%.0f)", result["http_status"])
				} else if result["http_status"] == 401.0 {
					err = fmt.Errorf("Not Authorized - Request without successful login session (%.0f)", result["http_status"])
				} else if result["http_status"] == 403.0 {
					err = fmt.Errorf("Forbidden - Request is missing  token or administrator is missing access profile permissions (%.0f)", result["http_status"])
				} else if result["http_status"] == 404.0 {
					err = fmt.Errorf("Resource Not Found - Unable to find the specified resource (%.0f)", result["http_status"])
				} else if result["http_status"] == 405.0 {
					err = fmt.Errorf("Method Not Allowed - Specified HTTP method is not allowed for this resource (%.0f)", result["http_status"])
				} else if result["http_status"] == 413.0 {
					err = fmt.Errorf("Request Entity Too Large - Request cannot be processed due to large entity (%.0f)", result["http_status"])
				} else if result["http_status"] == 424.0 {
					err = fmt.Errorf("Failed Dependency - Fail dependency can be duplicate resource, missing required parameter, missing required attribute, invalid attribute value (%.0f)", result["http_status"])
				} else if result["http_status"] == 429.0 {
					err = fmt.Errorf("Access temporarily blocked - Maximum failed authentications reached. The offended source is temporarily blocked for certain amount of time (%.0f)", result["http_status"])
				} else if result["http_status"] == 500.0 {
					err = fmt.Errorf("Internal Server Error - Internal error when processing the request (%.0f)", result["http_status"])
				} else {
					err = fmt.Errorf("Unknow Error (%.0f)", result["http_status"])
				}

				if result["cli_error"] != nil {
					err = fmt.Errorf(err.Error()+"\nCli response: \n%v", result["cli_error"])
				}

				return
			}

			err = fmt.Errorf("\n%v", body)
			return

		}
		err = fmt.Errorf("\n%v", body)
		return
	}

	// Authorization Required, etc. | Attention: scalable here
	err = fmt.Errorf("\n%v", body)
	return
}
