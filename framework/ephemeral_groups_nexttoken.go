package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terraform-providers/terraform-provider-fortiflexvm/fortiflexvm"
)

var _ ephemeral.EphemeralResource = &ephemeralGroupsNexttoken{}

type ephemeralGroupsNexttoken struct {
	fortiClient *fortiflexvm.FortiClient
}

func NewEphemeralGroupsNexttoken() ephemeral.EphemeralResource {
	return &ephemeralGroupsNexttoken{}
}
func (e *ephemeralGroupsNexttoken) Metadata(ctx context.Context, request ephemeral.MetadataRequest, response *ephemeral.MetadataResponse) {
	response.TypeName = "fortiflexvm_groups_nexttoken"
}

func (e *ephemeralGroupsNexttoken) Configure(ctx context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	// Always perform a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*fortiflexvm.FortiClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Ephemeral Resource Configure Type",
			fmt.Sprintf("Expected *FortiClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	e.fortiClient = client
}

func (e *ephemeralGroupsNexttoken) Schema(ctx context.Context, request ephemeral.SchemaRequest, response *ephemeral.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.Int64Attribute{
				Optional: true,
			},
			"config_id": schema.Int64Attribute{
				Optional: true,
			},
			"folder_path": schema.StringAttribute{
				Optional: true,
			},
			"status": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"token": schema.StringAttribute{
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func (e *ephemeralGroupsNexttoken) Open(ctx context.Context, request ephemeral.OpenRequest, response *ephemeral.OpenResponse) {
	var model ephemeralGroupsNexttokenModel

	diags := request.Config.Get(ctx, &model)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	c := e.fortiClient.Client
	// Prepare data
	request_obj := make(map[string]interface{})
	if v := model.ConfigID.ValueInt64(); v != 0 {
		request_obj["configId"] = v
	}
	if v := model.FolderPath.ValueString(); v != "" {
		request_obj["folderPath"] = v
	}
	if v := model.AccountID.ValueInt64(); v != 0 {
		request_obj["accountId"] = v
	}
	if v := model.Status.Elements(); len(v) > 0 {
		status_list := make([]string, 0)
		diags.Append(model.Status.ElementsAs(ctx, &status_list, false)...)
		if diags.HasError() {
			return
		}
		request_obj["status"] = status_list
	}

	if len(request_obj) == 0 {
		response.Diagnostics.AddError(
			"Either config_id or folder_path is required",
			"",
		)
	}

	// Send request
	o, err := c.ReadGroupsNexttoken(&request_obj)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Error to get token: %v", err),
			"",
		)
		return
	}

	entitlement, err := ephemeralRefreshObjectGroupsNexttoken(o)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to get a valid token",
			fmt.Sprintf("error: %v", err),
		)
		return
	}
	model.Token = types.StringValue(entitlement)
	response.Result.Set(ctx, model)
}

func ephemeralRefreshObjectGroupsNexttoken(o map[string]interface{}) (string, error) {
	var entitlement string
	switch v := o["entitlements"].(type) {
	case map[string]interface{}:
		if v, ok := o["entitlements"].(map[string]interface{})["token"]; ok {
			entitlement = v.(string)
		}
	case []interface{}:
		if v, ok := o["entitlements"].([]interface{})[0].(map[string]interface{})["token"]; ok {
			entitlement = v.(string)
		}
	default:
		return "", fmt.Errorf("Unsupported type: %T\n", v)
	}
	return entitlement, nil
}

type ephemeralGroupsNexttokenModel struct {
	AccountID  types.Int64  `tfsdk:"account_id"`
	ConfigID   types.Int64  `tfsdk:"config_id"`
	FolderPath types.String `tfsdk:"folder_path"`
	Status     types.List   `tfsdk:"status"`
	Token      types.String `tfsdk:"token"`
}
