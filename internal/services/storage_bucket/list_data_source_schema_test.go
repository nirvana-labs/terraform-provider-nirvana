// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_bucket_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/storage_bucket"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestStorageBucketsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*storage_bucket.StorageBucketsDataSourceModel)(nil)
	schema := storage_bucket.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
