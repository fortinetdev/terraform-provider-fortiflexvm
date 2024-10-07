package request

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	auth "github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/auth"
)

// Request describes the request to FortiFlex service
type Request struct {
	Auth         *auth.Auth
	HTTPCon      *http.Client
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Path         string
	Params       interface{}
	Data         *bytes.Buffer
}

// NewRequest creates request object with http method, path, params and data,
// It will save the http request, path, etc. for the next operations
// such as sending data, getting response, etc.
// It returns the created request object to the gobal plugin client.
func NewRequest(author *auth.Auth, httpcon *http.Client, method string, path string, params interface{}, data *bytes.Buffer) *Request {
	var h *http.Request

	if data == nil { // This "if-else" is necessary
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", data)
	}
	h.Header.Set("Content-Type", "application/json")
	r := &Request{
		Auth:        author,
		HTTPCon:     httpcon,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}
	return r
}

// Send request data to FortiFlex.
// If errors are encountered, it returns the error.
func (r *Request) Send(retries int) error {
	u := "https://support.fortinet.com" + r.Path

	var err error
	if r.Auth.Token == "" {
		err = fmt.Errorf("Could not find a API Token!")
		return err
	}
	var bearer = "Bearer " + r.Auth.Token
	r.HTTPRequest.Header.Set("Authorization", bearer)
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	retry := 0
	for retry <= retries {
		rsp, err := r.HTTPCon.Do(r.HTTPRequest)
		if err != nil {
			log.Printf("[ERROR] Request '%s' | %v | retry time %d", u, err.Error(), retry)
		} else {
			r.HTTPResponse = rsp
			break
		}
		retry++
		time.Sleep(time.Second)
	}
	if retry > retries {
		err = fmt.Errorf("Can't connect to server, please try it later.")
	}
	return err
}
