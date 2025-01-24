// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vm_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/vm"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestVMDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*vm.VMDataSourceModel)(nil)
	schema := vm.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
