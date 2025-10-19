// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_node_flex_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/rpc_node_flex"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestRPCNodeFlexDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*rpc_node_flex.RPCNodeFlexDataSourceModel)(nil)
	schema := rpc_node_flex.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
