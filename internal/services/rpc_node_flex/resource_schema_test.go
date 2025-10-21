// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_node_flex_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/rpc_node_flex"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestRPCNodeFlexModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*rpc_node_flex.RPCNodeFlexModel)(nil)
	schema := rpc_node_flex.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
