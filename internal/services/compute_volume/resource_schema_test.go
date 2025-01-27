// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_volume_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/compute_volume"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestComputeVolumeModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_volume.ComputeVolumeModel)(nil)
	schema := compute_volume.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
