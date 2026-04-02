// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package instance_type_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/instance_type"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestInstanceTypesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*instance_type.InstanceTypesDataSourceModel)(nil)
	schema := instance_type.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
