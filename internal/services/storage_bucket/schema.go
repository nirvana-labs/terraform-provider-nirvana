// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_bucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*StorageBucketResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the Operation.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Description:   "Name of the Bucket, used to address it over S3. It must be 3 to 63\ncharacters of lowercase letters, digits, dots and hyphens, start and end\nwith a letter or digit, not contain two adjacent dots and not be formatted\nas an IP address. \"admin\" and \"ops-api\" are reserved. The name cannot be\nchanged once the Bucket is created.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"project_id": schema.StringAttribute{
				Description:   "Project ID the Bucket belongs to.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"region": schema.StringAttribute{
				Description: "Region the resource is in.\nAvailable values: \"us-sva-2\", \"us-wdc-1\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("us-sva-2", "us-wdc-1"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"tags": schema.ListAttribute{
				Description: "Tags to attach to the Bucket.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description: "When the Bucket was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"endpoint": schema.StringAttribute{
				Description: "S3 endpoint of the Bucket's region. Point S3 clients at it, with an Object\nStorage Key of the same project and region.",
				Computed:    true,
			},
			"object_count": schema.Int64Attribute{
				Description: "Number of objects in the Bucket. Updated periodically, so it can lag behind\nrecent writes.",
				Computed:    true,
			},
			"signing_region": schema.StringAttribute{
				Description: "Region S3 clients must sign their requests for (the SigV4 region). Configure\nit as the region of any S3 client or SDK used with the Bucket.",
				Computed:    true,
			},
			"size_bytes": schema.Int64Attribute{
				Description: "Total size of the objects in the Bucket, in bytes. Updated periodically,\nso it can lag behind recent writes.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the resource.\nAvailable values: \"pending\", \"creating\", \"updating\", \"ready\", \"deleting\", \"deleted\", \"error\", \"suspended\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending",
						"creating",
						"updating",
						"ready",
						"deleting",
						"deleted",
						"error",
						"suspended",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "When the Bucket was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"url": schema.StringAttribute{
				Description: "Path-style URL of the Bucket.",
				Computed:    true,
			},
		},
	}
}

func (r *StorageBucketResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *StorageBucketResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
