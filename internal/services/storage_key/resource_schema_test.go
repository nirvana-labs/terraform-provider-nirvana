// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_key_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/storage_key"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestStorageKeyModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*storage_key.StorageKeyModel)(nil)
	schema := storage_key.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
