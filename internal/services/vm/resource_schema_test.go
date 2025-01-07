// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vm_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/nirvana-terraform/internal/services/vm"
	"github.com/stainless-sdks/nirvana-terraform/internal/test_helpers"
)

func TestVMModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*vm.VMModel)(nil)
	schema := vm.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
