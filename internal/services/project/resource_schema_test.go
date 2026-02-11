// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project_test

import (
	"context"
	"testing"

	"github.com/nirvana-labs/terraform-provider-nirvana/internal/services/project"
	"github.com/nirvana-labs/terraform-provider-nirvana/internal/test_helpers"
)

func TestProjectModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*project.ProjectModel)(nil)
	schema := project.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
