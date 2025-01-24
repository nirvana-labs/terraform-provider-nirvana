// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package volume_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/volume"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestVolumeDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*volume.VolumeDataSourceModel)(nil)
	schema := volume.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
