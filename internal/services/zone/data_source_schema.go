// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &ZoneDataSource{}
var _ datasource.DataSourceWithValidateConfig = &ZoneDataSource{}

func (r ZoneDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"account": schema.SingleNestedAttribute{
				Description: "The account the zone belongs to",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "The name of the account",
						Optional:    true,
					},
				},
			},
			"activated_on": schema.StringAttribute{
				Description: "The last time proof of ownership was detected and the zone was made\nactive",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the zone was created",
				Optional:    true,
			},
			"development_mode": schema.Float64Attribute{
				Description: "The interval (in seconds) from when development mode expires\n(positive integer) or last expired (negative integer) for the\ndomain. If development mode has never been enabled, this value is 0.",
				Optional:    true,
			},
			"meta": schema.SingleNestedAttribute{
				Description: "Metadata about the zone",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"cdn_only": schema.BoolAttribute{
						Description: "The zone is only configured for CDN",
						Optional:    true,
					},
					"custom_certificate_quota": schema.Int64Attribute{
						Description: "Number of Custom Certificates the zone can have",
						Optional:    true,
					},
					"dns_only": schema.BoolAttribute{
						Description: "The zone is only configured for DNS",
						Optional:    true,
					},
					"foundation_dns": schema.BoolAttribute{
						Description: "The zone is setup with Foundation DNS",
						Optional:    true,
					},
					"page_rule_quota": schema.Int64Attribute{
						Description: "Number of Page Rules a zone can have",
						Optional:    true,
					},
					"phishing_detected": schema.BoolAttribute{
						Description: "The zone has been flagged for phishing",
						Optional:    true,
					},
					"step": schema.Int64Attribute{
						Optional: true,
					},
				},
			},
			"modified_on": schema.StringAttribute{
				Description: "When the zone was last modified",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The domain name",
				Optional:    true,
			},
			"name_servers": schema.StringAttribute{
				Description: "The name servers Cloudflare assigns to a zone",
				Optional:    true,
			},
			"original_dnshost": schema.StringAttribute{
				Description: "DNS host at the time of switching to Cloudflare",
				Optional:    true,
			},
			"original_name_servers": schema.StringAttribute{
				Description: "Original name servers before moving to Cloudflare",
				Optional:    true,
			},
			"original_registrar": schema.StringAttribute{
				Description: "Registrar for the domain at the time of switching to Cloudflare",
				Optional:    true,
			},
			"owner": schema.SingleNestedAttribute{
				Description: "The owner of the zone",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Identifier",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of the owner",
						Optional:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of owner",
						Optional:    true,
					},
				},
			},
			"vanity_name_servers": schema.StringAttribute{
				Description: "An array of domains used for custom name servers. This is only available for Business and Enterprise plans.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "An account ID",
								Optional:    true,
							},
							"name": schema.StringAttribute{
								Description: "An account Name. Optional filter operators can be provided to extend refine the search:\n  * `equal` (default)\n  * `not_equal`\n  * `starts_with`\n  * `ends_with`\n  * `contains`\n  * `starts_with_case_sensitive`\n  * `ends_with_case_sensitive`\n  * `contains_case_sensitive`\n",
								Optional:    true,
							},
						},
					},
					"direction": schema.StringAttribute{
						Description: "Direction to order zones.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"match": schema.StringAttribute{
						Description: "Whether to match all search requirements or at least one (any).",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("any", "all"),
						},
					},
					"name": schema.StringAttribute{
						Description: "A domain name. Optional filter operators can be provided to extend refine the search:\n  * `equal` (default)\n  * `not_equal`\n  * `starts_with`\n  * `ends_with`\n  * `contains`\n  * `starts_with_case_sensitive`\n  * `ends_with_case_sensitive`\n  * `contains_case_sensitive`\n",
						Optional:    true,
					},
					"order": schema.StringAttribute{
						Description: "Field to order zones by.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("name", "status", "account.id", "account.name"),
						},
					},
					"page": schema.Float64Attribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
					},
					"per_page": schema.Float64Attribute{
						Description: "Number of zones per page.",
						Computed:    true,
						Optional:    true,
					},
					"status": schema.StringAttribute{
						Description: "A zone status",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("initializing", "pending", "active", "moved"),
						},
					},
				},
			},
		},
	}
}

func (r *ZoneDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *ZoneDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}