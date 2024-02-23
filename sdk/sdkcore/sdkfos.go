// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Description: SDK for FortiFlex Provider

package forticlient

import (
	"fmt"
)

// type creatUpdateOutput struct {
// 	Vdom       string  `json:"vdom"`
// 	Mkey       string  `json:"mkey"`
// 	Status     string  `json:"status"`
// 	HTTPStatus float64 `json:"http_status"`
// }

// ReadProgramsList API operation for FortiFlex gets the Programs list
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadProgramsList(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/programs/list"
	rspKey := "programs"

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}

// ReadConfigsList API operation for FortiFlex gets the Configurations list
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadConfigsList(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/configs/list"
	rspKey := "configs"

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}

// CreateConfig API operation for FortiFlex creates a Configuration.
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateConfig(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/configs/create"
	rspKey := "configs"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// UpdateConfig API operation for FortiFlex updates a Configuration.
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateConfig(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/configs/update"
	rspKey := "configs"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// UpdateConfigStatus API operation for FortiFlex updates the status of the Configuration.
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateConfigStatus(params *map[string]interface{}, op string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := fmt.Sprintf("/ES/api/fortiflex/v2/configs/%v", op)
	rspKey := "configs"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// ReadEntitlementsList API operation for FortiFlex gets the Virtual Machines list
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadEntitlementsList(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/entitlements/list"
	rspKey := "entitlements"

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}

// CreateEntitlementsVM API operation for FortiFlex creates VMs based on a Configuration.
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateEntitlementsVM(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/entitlements/vm/create"
	rspKey := "entitlements"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// CreateEntitlementsHW API operation
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateEntitlementsHW(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/entitlements/hardware/create"
	rspKey := "entitlements"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// CreateEntitlementsCloud API operation
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateEntitlementsCloud(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/entitlements/cloud/create"
	rspKey := "entitlements"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// ReadEntitlementsPoint API operation for FortiFlex gets point usage for Virtual Machines
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadEntitlementsPoint(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/entitlements/points"
	rspKey := "entitlements"

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}

// UpdateVmUpdate API operation for FortiFlex update a VM's setting
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateVmUpdate(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/entitlements/update"
	rspKey := "entitlements"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// UpdateVmUpdateStatus API operation for FortiFlex updates the status of the VM.
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateVmUpdateStatus(params *map[string]interface{}, op string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := fmt.Sprintf("/ES/api/fortiflex/v2/entitlements/%v", op)
	rspKey := "entitlements"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// UpdateVmUpdateRegenerateToken API operation for FortiFlex regenerate token for a VM
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateVmUpdateRegenerateToken(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/entitlements/vm/token"
	rspKey := "entitlements"

	mapTmp, err = createUpdate(c, HTTPMethod, path, rspKey, params)
	return
}

// ReadGroupsList API operation for FortiFlex gets the Groups list
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadGroupsList(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/flexvm/v1/groups/list"
	rspKey := "groups"
	if value, ok := (*params)["accountId"]; ok {
		if value != 0 {
			path = "/ES/api/fortiflex/v2/groups/list"
		}
	}

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}

// ReadGroupsNexttoken API operation for FortiFlex gets the next available (unused) token
// Returns the requested value when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadGroupsNexttoken(params *map[string]interface{}) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/ES/api/fortiflex/v2/groups/nexttoken"
	rspKey := "entitlements"

	mapTmp, err = read(c, HTTPMethod, path, true, rspKey, params)
	return
}
