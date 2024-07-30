// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_tag

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessTagsResultListDataSourceEnvelope struct {
	Result *[]*AccessTagsResultDataSourceModel `json:"result,computed"`
}

type AccessTagsDataSourceModel struct {
	AccountID types.String                        `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                         `tfsdk:"max_items"`
	Result    *[]*AccessTagsResultDataSourceModel `tfsdk:"result"`
}

type AccessTagsResultDataSourceModel struct {
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	AppCount  types.Int64       `tfsdk:"app_count" json:"app_count"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}
