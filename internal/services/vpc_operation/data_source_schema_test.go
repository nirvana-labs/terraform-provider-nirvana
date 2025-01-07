// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vpc_operation_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/nirvana-terraform/internal/services/vpc_operation"
	"github.com/stainless-sdks/nirvana-terraform/internal/test_helpers"
)

func TestVPCOperationDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*vpc_operation.VPCOperationDataSourceModel)(nil)
	schema := vpc_operation.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
