// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_key_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/api_key"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestAPIKeysDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_key.APIKeysDataSourceModel)(nil)
	schema := api_key.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
