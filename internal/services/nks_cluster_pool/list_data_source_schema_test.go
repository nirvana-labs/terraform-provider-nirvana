// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster_pool_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/nks_cluster_pool"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestNKSClusterPoolsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*nks_cluster_pool.NKSClusterPoolsDataSourceModel)(nil)
	schema := nks_cluster_pool.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
