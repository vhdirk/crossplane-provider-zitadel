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

	// ProjectRef references a project by id
	ProjectGrantRef = config.Reference{
		TerraformName: "zitadel_project_grant",
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

	// HumanUserRef references a human user by id
	HumanUserRef = config.Reference{
		TerraformName: "zitadel_human_user",
		Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)`,
	}
)

// Configure configures the base provider.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("zitadel_org", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Organisation"
	})

	p.AddResourceConfigurator("zitadel_project", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Project"

		r.References["org_id"] = OrganisationRef
	})

	p.AddResourceConfigurator("zitadel_project_grant", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectGrant"

		r.References["org_id"] = OrganisationRef
		r.References["granted_org_id"] = OrganisationRef
		r.References["project_id"] = ProjectRef
	})

	p.AddResourceConfigurator("zitadel_project_grant_member", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectGrant"

		r.References["org_id"] = OrganisationRef
		r.References["granted_org_id"] = OrganisationRef
		r.References["project_id"] = ProjectRef
	})

	p.AddResourceConfigurator("zitadel_action", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Action"

		r.References["org_id"] = OrganisationRef
	})

	// zitadel_application_api
	// zitadel_application_key
	// zitadel_application_oidc
	// zitadel_application_saml
	// zitadel_default_domain_claimed_message_text
	// zitadel_default_domain_policy
	// zitadel_default_init_message_text
	// zitadel_default_label_policy
	// zitadel_default_lockout_policy
	// zitadel_default_login_policy
	// zitadel_default_login_texts
	// zitadel_default_notification_policy
	// zitadel_default_oidc_settings
	// zitadel_default_password_age_policy
	// zitadel_default_password_change_message_text
	// zitadel_default_password_complexity_policy
	// zitadel_default_password_reset_message_text
	// zitadel_default_passwordless_registration_message_text
	// zitadel_default_privacy_policy
	// zitadel_default_verify_email_message_text
	// zitadel_default_verify_email_otp_message_text
	// zitadel_default_verify_phone_message_text
	// zitadel_default_verify_sms_otp_message_text

	p.AddResourceConfigurator("zitadel_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Domain"

		r.References["org_id"] = OrganisationRef
	})

	// zitadel_domain_claimed_message_text
	// zitadel_domain_policy
	p.AddResourceConfigurator("zitadel_human_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "HumanUser"

		r.References["org_id"] = OrganisationRef
	})
	// zitadel_idp_azure_ad
	// zitadel_idp_github
	// zitadel_idp_github_es
	// zitadel_idp_gitlab
	// zitadel_idp_gitlab_self_hosted
	// zitadel_idp_google
	// zitadel_idp_ldap
	// zitadel_idp_oauth
	// zitadel_idp_oidc
	// zitadel_idp_saml
	// zitadel_init_message_text
	// zitadel_instance_member
	// zitadel_label_policy
	// zitadel_lockout_policy
	// zitadel_login_policy
	// zitadel_login_texts
	// zitadel_machine_key

	p.AddResourceConfigurator("zitadel_machine_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MachineUser"

		r.References["org_id"] = OrganisationRef
	})

	// zitadel_notification_policy

	// zitadel_org_idp_azure_ad
	// zitadel_org_idp_github
	// zitadel_org_idp_github_es
	// zitadel_org_idp_gitlab
	// zitadel_org_idp_gitlab_self_hosted
	// zitadel_org_idp_google
	// zitadel_org_idp_jwt
	// zitadel_org_idp_ldap
	// zitadel_org_idp_oauth
	// zitadel_org_idp_oidc
	// zitadel_org_idp_saml
	// zitadel_org_member
	// zitadel_org_metadata
	// zitadel_password_age_policy
	// zitadel_password_change_message_text
	// zitadel_password_complexity_policy
	// zitadel_password_reset_message_text
	// zitadel_passwordless_registration_message_text
	// zitadel_personal_access_token
	// zitadel_privacy_policy

	// zitadel_project_grant_member
	// zitadel_project_member
	// zitadel_project_role
	// zitadel_sms_provider_http
	// zitadel_sms_provider_twilio
	// zitadel_smtp_config
	// zitadel_trigger_actions
	// zitadel_user_grant
	// zitadel_user_metadata
	// zitadel_verify_email_message_text
	// zitadel_verify_email_otp_message_text
	// zitadel_verify_phone_message_text
	// zitadel_verify_sms_otp_message_text

}
