// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_connect_connection_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/networking_connect_connection"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestNetworkingConnectConnectionDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*networking_connect_connection.NetworkingConnectConnectionDataSourceModel)(nil)
	schema := networking_connect_connection.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
