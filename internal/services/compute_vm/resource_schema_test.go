// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/compute_vm"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestComputeVMModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_vm.ComputeVMModel)(nil)
	schema := compute_vm.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
