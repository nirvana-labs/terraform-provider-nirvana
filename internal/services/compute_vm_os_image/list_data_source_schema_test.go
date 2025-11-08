// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_vm_os_image_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/compute_vm_os_image"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestComputeVmosImagesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_vm_os_image.ComputeVMOSImagesDataSourceModel)(nil)
	schema := compute_vm_os_image.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
