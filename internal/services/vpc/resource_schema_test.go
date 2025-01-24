// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vpc_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/vpc"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestVPCModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*vpc.VPCModel)(nil)
	schema := vpc.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
