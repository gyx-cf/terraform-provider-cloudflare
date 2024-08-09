// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_headers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &ManagedHeadersDataSource{}

func (d *ManagedHeadersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"managed_request_headers": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Human-readable identifier of the Managed Transform.",
							Computed:    true,
							Optional:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "When true, the Managed Transform is enabled.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
			"managed_response_headers": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Human-readable identifier of the Managed Transform.",
							Computed:    true,
							Optional:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "When true, the Managed Transform is enabled.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ManagedHeadersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}