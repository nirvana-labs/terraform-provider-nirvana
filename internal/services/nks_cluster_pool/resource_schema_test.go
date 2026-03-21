// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster_pool_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/nks_cluster_pool"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestNKSClusterPoolModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*nks_cluster_pool.NKSClusterPoolModel)(nil)
	schema := nks_cluster_pool.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
