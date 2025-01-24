// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*VPCResource)(nil)

func (r *VPCResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
