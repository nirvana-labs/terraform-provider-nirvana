// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_key_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/storage_key"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestStorageKeysDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*storage_key.StorageKeysDataSourceModel)(nil)
	schema := storage_key.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
