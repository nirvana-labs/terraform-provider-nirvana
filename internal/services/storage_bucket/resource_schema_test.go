// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_bucket_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/storage_bucket"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestStorageBucketModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*storage_bucket.StorageBucketModel)(nil)
	schema := storage_bucket.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
