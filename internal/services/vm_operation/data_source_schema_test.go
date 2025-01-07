// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vm_operation_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/nirvana-terraform/internal/services/vm_operation"
	"github.com/stainless-sdks/nirvana-terraform/internal/test_helpers"
)

func TestVMOperationDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*vm_operation.VMOperationDataSourceModel)(nil)
	schema := vm_operation.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
