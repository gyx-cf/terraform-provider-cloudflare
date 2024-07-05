// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ByoIPPrefixesResultListDataSourceEnvelope struct {
	Result *[]*ByoIPPrefixesItemsDataSourceModel `json:"result,computed"`
}

type ByoIPPrefixesDataSourceModel struct {
	AccountID types.String                          `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                           `tfsdk:"max_items"`
	Items     *[]*ByoIPPrefixesItemsDataSourceModel `tfsdk:"items"`
}

type ByoIPPrefixesItemsDataSourceModel struct {
	ID                   types.String `tfsdk:"id" json:"id,computed"`
	AccountID            types.String `tfsdk:"account_id" json:"account_id,computed"`
	Advertised           types.Bool   `tfsdk:"advertised" json:"advertised,computed"`
	AdvertisedModifiedAt types.String `tfsdk:"advertised_modified_at" json:"advertised_modified_at,computed"`
	Approved             types.String `tfsdk:"approved" json:"approved,computed"`
	ASN                  types.Int64  `tfsdk:"asn" json:"asn,computed"`
	CIDR                 types.String `tfsdk:"cidr" json:"cidr,computed"`
	CreatedAt            types.String `tfsdk:"created_at" json:"created_at,computed"`
	Description          types.String `tfsdk:"description" json:"description,computed"`
	LOADocumentID        types.String `tfsdk:"loa_document_id" json:"loa_document_id,computed"`
	ModifiedAt           types.String `tfsdk:"modified_at" json:"modified_at,computed"`
	OnDemandEnabled      types.Bool   `tfsdk:"on_demand_enabled" json:"on_demand_enabled,computed"`
	OnDemandLocked       types.Bool   `tfsdk:"on_demand_locked" json:"on_demand_locked,computed"`
}