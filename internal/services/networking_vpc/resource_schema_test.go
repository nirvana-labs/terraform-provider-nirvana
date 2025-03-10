// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_vpc_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/networking_vpc"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestNetworkingVPCModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*networking_vpc.NetworkingVPCModel)(nil)
	schema := networking_vpc.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
