// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vm_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/nirvana-terraform/internal/services/vm"
	"github.com/stainless-sdks/nirvana-terraform/internal/test_helpers"
)

func TestVMDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*vm.VMDataSourceModel)(nil)
	schema := vm.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
