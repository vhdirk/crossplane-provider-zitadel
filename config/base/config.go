package base

import (
	"github.com/crossplane/upjet/pkg/config"

	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

const shortGroup = ""

func ExtractParamPath(sourceAttr string, isObservation bool) xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		mr.GetObjectKind()

		return ""
	}
}

// Reference definitions for resources
var (
	// OrganisationRef references an organization by id
	OrganisationRef = config.Reference{
		TerraformName: "zitadel_org",
		Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)`,
	}

	// ProjectRef references a project by id
	ProjectRef = config.Reference{
		TerraformName: "zitadel_project",
		Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)`,
	}

	// ApplicationRef references an application by id
	ApplicationRef = config.Reference{
		TerraformName: "zitadel_application",
		Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)`,
	}

	// UserRef references a user by id
	UserRef = config.Reference{
		TerraformName: "zitadel_human_user",
		Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)`,
	}

	// DomainRef references a domain by id
	DomainRef = config.Reference{
		TerraformName: "zitadel_domain",
		Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)`,
	}
)

// Configure configures the base provider.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Organisation"
	})

	p.AddResourceConfigurator("zitadel_action", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Action"

		r.References["org_id"] = OrganisationRef
	})

	p.AddResourceConfigurator("zitadel_human_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "HumanUser"

		r.References["org_id"] = OrganisationRef
	})
	p.AddResourceConfigurator("zitadel_machine_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MachineUser"

		r.References["org_id"] = OrganisationRef
	})
	p.AddResourceConfigurator("zitadel_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Domain"

		r.References["org_id"] = OrganisationRef
	})
}
