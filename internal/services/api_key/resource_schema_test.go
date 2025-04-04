// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_key_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/api_key"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestAPIKeyModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_key.APIKeyModel)(nil)
	schema := api_key.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
