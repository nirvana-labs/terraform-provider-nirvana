// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_kubernetes_version_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/nks_kubernetes_version"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestNKSKubernetesVersionsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*nks_kubernetes_version.NKSKubernetesVersionsDataSourceModel)(nil)
	schema := nks_kubernetes_version.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
