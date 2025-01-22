// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package operation_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/nirvana-terraform/internal/services/operation"
	"github.com/stainless-sdks/nirvana-terraform/internal/test_helpers"
)

func TestOperationDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*operation.OperationDataSourceModel)(nil)
	schema := operation.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
