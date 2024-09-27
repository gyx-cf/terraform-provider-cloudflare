// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_allow_pattern

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailSecurityAllowPatternDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account Identifier",
				Optional:    true,
			},
			"pattern_id": schema.Int64Attribute{
				Optional: true,
			},
			"comments": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"is_recipient": schema.BoolAttribute{
				Computed: true,
			},
			"is_regex": schema.BoolAttribute{
				Computed: true,
			},
			"is_sender": schema.BoolAttribute{
				Computed: true,
			},
			"is_spoof": schema.BoolAttribute{
				Computed: true,
			},
			"last_modified": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"pattern": schema.StringAttribute{
				Computed: true,
			},
			"pattern_type": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"EMAIL",
						"DOMAIN",
						"IP",
						"UNKNOWN",
					),
				},
			},
			"verify_sender": schema.BoolAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Account Identifier",
						Required:    true,
					},
					"direction": schema.StringAttribute{
						Description: "The sorting direction.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"is_recipient": schema.BoolAttribute{
						Optional: true,
					},
					"is_sender": schema.BoolAttribute{
						Optional: true,
					},
					"is_spoof": schema.BoolAttribute{
						Optional: true,
					},
					"order": schema.StringAttribute{
						Description: "The field to sort by.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("pattern", "created_at"),
						},
					},
					"pattern_type": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"EMAIL",
								"DOMAIN",
								"IP",
								"UNKNOWN",
							),
						},
					},
					"search": schema.StringAttribute{
						Description: "Allows searching in multiple properties of a record simultaneously.\nThis parameter is intended for human users, not automation. Its exact\nbehavior is intentionally left unspecified and is subject to change\nin the future.",
						Optional:    true,
					},
					"verify_sender": schema.BoolAttribute{
						Optional: true,
					},
				},
			},
		},
	}
}

func (d *EmailSecurityAllowPatternDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailSecurityAllowPatternDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("pattern_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("pattern_id")),
	}
}
