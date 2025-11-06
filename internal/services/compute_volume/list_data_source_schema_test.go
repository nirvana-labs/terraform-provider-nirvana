// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_volume_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/compute_volume"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestComputeVolumesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_volume.ComputeVolumesDataSourceModel)(nil)
	schema := compute_volume.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
