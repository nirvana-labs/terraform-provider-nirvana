// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vpc_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/nirvana-terraform/internal/services/vpc"
	"github.com/stainless-sdks/nirvana-terraform/internal/test_helpers"
)

func TestVPCDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*vpc.VPCDataSourceModel)(nil)
	schema := vpc.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
