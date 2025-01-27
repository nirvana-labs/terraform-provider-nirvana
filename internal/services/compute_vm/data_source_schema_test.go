// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/compute_vm"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestComputeVMDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_vm.ComputeVMDataSourceModel)(nil)
	schema := compute_vm.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
