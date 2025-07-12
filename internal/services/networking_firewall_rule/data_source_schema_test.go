// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_firewall_rule_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/networking_firewall_rule"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestNetworkingFirewallRuleDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*networking_firewall_rule.NetworkingFirewallRuleDataSourceModel)(nil)
	schema := networking_firewall_rule.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
