// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_cluster_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/nks_cluster"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestNKSClusterModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*nks_cluster.NKSClusterModel)(nil)
	schema := nks_cluster.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
