// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*StorageKeyResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "An Object Storage Key: an S3 access key for a project's Buckets in one region.\n\n" +
			"The `secret` is returned only when the key is created. The provider keeps it in state from the create response " +
			"and never re-reads it from the API, so treat the state as sensitive. Importing an existing key cannot recover " +
			"its secret: `secret` stays null for an imported key. To get a new secret, replace the key " +
			"(for example with `terraform apply -replace`).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the Object Storage Key.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"project_id": schema.StringAttribute{
				Description:   "Project ID the Object Storage Key belongs to. The key can access the\nproject's Buckets in its region.",
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
			"name": schema.StringAttribute{
				Description: "Name of the Object Storage Key.",
				Required:    true,
			},
			"access_key": schema.StringAttribute{
				Description: "S3 access key ID.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the Object Storage Key was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"secret": schema.StringAttribute{
				Description:   "S3 secret access key. It is only returned here and cannot be retrieved\nagain: store it securely, or delete the key and create a new one if it is\nlost.",
				Computed:      true,
				Sensitive:     true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
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
				Description: "When the Object Storage Key was updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *StorageKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *StorageKeyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
