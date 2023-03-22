// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Description: SDK for FlexVM Provider

package forticlient

import (
	"fmt"
)

type creatUpdateOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// ReadProgramsList API operation for FlexVM gets the Programs list
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadProgramsList(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/programs/list"
	rspKey := "programs"

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}

// ReadConfigsList API operation for FlexVM gets the Configurations list
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadConfigsList(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/configs/list"
	rspKey := "configs"

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}

// CreateConfig API operation for FlexVM creates a Configuration.
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateConfig(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/configs/create"
	rspKey := "configs"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// UpdateConfig API operation for FlexVM updates a Configuration.
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateConfig(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/configs/update"
	rspKey := "configs"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// UpdateConfigStatus API operation for FlexVM updates the status of the Configuration.
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateConfigStatus(params *map[string]interface{}, op string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := fmt.Sprintf("/ES/api/flexvm/v1/configs/%v", op)
	rspKey := "configs"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// ReadVmsList API operation for FlexVM gets the Virtual Machines list
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadVmsList(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/vms/list"
	rspKey := "vms"

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}

// CreateVms API operation for FlexVM creates VMs based on a Configuration.
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateVms(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/vms/create"
	rspKey := "vms"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// ReadVmsPoint API operation for FlexVM gets point usage for Virtual Machines
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadVmsPoint(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/vms/points"
	rspKey := "vms"

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}

// UpdateVmUpdate API operation for FlexVM update a VM's setting
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateVmUpdate(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/vms/update"
	rspKey := "vms"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// UpdateVmUpdateStatus API operation for FlexVM updates the status of the VM.
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateVmUpdateStatus(params *map[string]interface{}, op string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := fmt.Sprintf("/ES/api/flexvm/v1/vms/%v", op)
	rspKey := "vms"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// UpdateVmUpdateRegenerateToken API operation for FlexVM regenerate token for a VM
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateVmUpdateRegenerateToken(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/vms/token"
	rspKey := "vms"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// ReadGroupsList API operation for FlexVM gets the Groups list
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadGroupsList(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/groups/list"
	rspKey := "groups"

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}

// ReadGroupsNexttoken API operation for FlexVM gets the next available (unused) token
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadGroupsNexttoken(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/groups/nexttoken"
	rspKey := "vms"

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}
