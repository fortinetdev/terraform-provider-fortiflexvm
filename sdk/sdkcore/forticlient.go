package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
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

// FortiSDKClient describes the global FortiFlex plugin client instance
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

// NewRequest creates the request to FortiFlex for the client
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
	log.Printf("FortiFlex login: %s", string(body))

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	// err = fortiAPIErrorFormat(result, string(body))

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
			result = replaceMessage(result)
		}
		result_byte, json_err := json.Marshal(result)
		if json_err != nil {
			err = fmt.Errorf("\n%v", body)
			return
		}
		err = fmt.Errorf("\n%v", string(result_byte))
		return
	}

	// Authorization Required, etc. | Attention: scalable here
	err = fmt.Errorf("\n%v", body)
	return
}

func replaceMessage(result map[string]interface{}) map[string]interface{} {
	if result["message"] != nil {
		message := fmt.Sprintf("%v", result["message"])
		re, compile_err := regexp.Compile(`Parameter id (\d+)`)
		if compile_err != nil {
			return result
		}
		newMsg := re.ReplaceAllStringFunc(message, func(match string) string {
			numberStr := re.FindStringSubmatch(match)[1]
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				return match
			}
			_, param_name, _ := paramID2Name(number)
			if param_name == "" {
				return match
			}
			return "Parameter " + param_name
		})
		result["message"] = newMsg
	}
	return result
}

func paramID2Name(p_id int) (string, string, string) {
	switch p_id {
	case 1:
		return "fgt_vm_bundle", "cpu_size", "string"
	case 2:
		return "fgt_vm_bundle", "service_pkg", "string"
	case 3: // Deprecated, keep it for backward compatible
		return "fmg_vm", "managed_dev", "int"
	case 4:
		return "fwb_vm", "cpu_size", "string"
	case 5:
		return "fwb_vm", "service_pkg", "string"
	case 6:
		return "fgt_vm_lcs", "cpu_size", "String"
	case 7:
		return "fgt_vm_lcs", "fortiguard_services", "list"
	case 8:
		return "fgt_vm_lcs", "support_service", "string"
	case 9:
		return "fmg_vm", "adom_num", "int"
	case 10:
		return "fgt_vm_bundle", "vdom_num", "int"
	case 11:
		return "fgt_vm_lcs", "vdom_num", "int"
	case 12:
		return "fgt_vm_lcs", "cloud_services", "list"
	case 13:
		return "fc_ems_op", "ztna_num", "int"
	case 14:
		return "fc_ems_op", "epp_ztna_num", "int"
	case 15:
		return "fc_ems_op", "chromebook", "int"
	case 16:
		return "fc_ems_op", "support_service", "string"
	case 21:
		return "faz_vm", "daily_storage", "int"
	case 22:
		return "faz_vm", "adom_num", "int"
	case 23:
		return "faz_vm", "support_service", "string"
	case 24:
		return "fpc_vm", "managed_dev", "int"
	case 25:
		return "fad_vm", "cpu_size", "string"
	case 26:
		return "fad_vm", "service_pkg", "string"
	case 27:
		return "fgt_hw", "device_model", "string"
	case 28:
		return "fgt_hw", "service_pkg", "string"
	case 29:
		return "fgt_hw", "addons", "list"
	case 30:
		return "fmg_vm", "managed_dev", "int"
	case 32:
		return "fwbc_private", "average_throughput", "int"
	case 33:
		return "fwbc_private", "web_applications", "int"
	case 34:
		return "fwbc_public", "average_throughput", "int"
	case 35:
		return "fwbc_public", "web_applications", "int"
	case 36:
		return "fc_ems_op", "addons", "list"
	case 37:
		return "fc_ems_cloud", "ztna_num", "int"
	case 38:
		return "fc_ems_cloud", "ztna_fgf_num", "int"
	case 39:
		return "fc_ems_cloud", "epp_ztna_num", "int"
	case 40:
		return "fc_ems_cloud", "epp_ztna_fgf_num", "int"
	case 41:
		return "fc_ems_cloud", "chromebook", "int"
	case 42:
		return "fc_ems_cloud", "addons", "list"
	case 43:
		return "fgt_vm_bundle", "fortiguard_services", "list"
	case 44:
		return "fgt_vm_bundle", "cloud_services", "list"
	case 45:
		return "fgt_vm_bundle", "support_service", "string"
	case 46:
		return "fortiedr", "service_pkg", "string"
	case 47:
		return "fortiedr", "endpoints", "int" // Read only
	case 48:
		return "fortisase", "users", "int"
	case 49:
		return "fortisase", "service_pkg", "string"
	case 50:
		return "fortisase", "bandwidth", "int"
	case 51:
		return "fortisase", "dedicated_ips", "int"
	case 52:
		return "fortiedr", "addons", "list"
	default:
		return "", "", ""
	}
}
