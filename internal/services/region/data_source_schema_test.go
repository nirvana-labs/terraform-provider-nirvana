// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package region_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/region"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestRegionDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*region.RegionDataSourceModel)(nil)
	schema := region.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
