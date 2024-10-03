// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_infrastructure_target

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessInfrastructureTargetDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier",
				Optional:    true,
			},
			"target_id": schema.StringAttribute{
				Description: "Target identifier",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Date and time at which the target was created",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"hostname": schema.StringAttribute{
				Description: "A non-unique field that refers to a target",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Target identifier",
				Computed:    true,
			},
			"modified_at": schema.StringAttribute{
				Description: "Date and time at which the target was modified",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"ip": schema.SingleNestedAttribute{
				Description: "The IPv4/IPv6 address that identifies where to reach a target",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessInfrastructureTargetIPDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"ipv4": schema.SingleNestedAttribute{
						Description: "The target's IPv4 address",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessInfrastructureTargetIPIPV4DataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ip_addr": schema.StringAttribute{
								Description: "IP address of the target",
								Computed:    true,
							},
							"virtual_network_id": schema.StringAttribute{
								Description: "Private virtual network identifier for the target",
								Computed:    true,
							},
						},
					},
					"ipv6": schema.SingleNestedAttribute{
						Description: "The target's IPv6 address",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessInfrastructureTargetIPIPV6DataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ip_addr": schema.StringAttribute{
								Description: "IP address of the target",
								Computed:    true,
							},
							"virtual_network_id": schema.StringAttribute{
								Description: "Private virtual network identifier for the target",
								Computed:    true,
							},
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Account identifier",
						Required:    true,
					},
					"created_after": schema.StringAttribute{
						Description: "Date and time at which the target was created",
						Optional:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"hostname": schema.StringAttribute{
						Description: "Hostname of a target",
						Optional:    true,
					},
					"hostname_contains": schema.StringAttribute{
						Description: "Partial match to the hostname of a target",
						Optional:    true,
					},
					"ip_v4": schema.StringAttribute{
						Description: "IPv4 address of the target",
						Optional:    true,
					},
					"ip_v6": schema.StringAttribute{
						Description: "IPv6 address of the target",
						Optional:    true,
					},
					"modified_after": schema.StringAttribute{
						Description: "Date and time at which the target was modified",
						Optional:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"virtual_network_id": schema.StringAttribute{
						Description: "Private virtual network identifier of the target",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustAccessInfrastructureTargetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessInfrastructureTargetDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("target_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("target_id")),
	}
}