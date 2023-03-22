package request

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/config"
)

// Request describes the request to FlexVM service
type Request struct {
	Config       config.Config
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Path         string
	Params       interface{}
	Data         *bytes.Buffer
}

// New creates reqeust object with http method, path, params and data,
// It will save the http request, path, etc. for the next operations
// such as sending data, getting response, etc.
// It returns the created request object to the gobal plugin client.
func New(c config.Config, method string, path string, params interface{}, data *bytes.Buffer) *Request {
	var h *http.Request

	if data == nil {
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", data)
	}

	r := &Request{
		Config:      c,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}
	return r
}

// Send request data to FlexVM.
// If errors are encountered, it returns the error.
func (r *Request) Send() error {
	return r.SendRequest(15)
}

// SendRequest request data to FlexVM.
// If errors are encountered, it returns the error.
func (r *Request) SendRequest(retries int) error {
	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	u := buildURL(r)

	var err error
	if r.Config.Auth.Token == "" {
		err = fmt.Errorf("Could not find a API Token!")
		return err
	}
	var bearer = "Bearer " + r.Config.Auth.Token
	r.HTTPRequest.Header.Set("Authorization", bearer)
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}

func filterapikey(v string) string {
	re, _ := regexp.Compile("access_token=.*?\"")
	res := re.ReplaceAllString(v, "access_token=***************\"")

	return res
}

func buildURL(r *Request) string {
	u := "https://support.fortinet.com"
	u += r.Path

	return u
}
