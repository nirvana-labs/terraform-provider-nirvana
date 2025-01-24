// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package volume_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/volume"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestVolumeModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*volume.VolumeModel)(nil)
	schema := volume.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
