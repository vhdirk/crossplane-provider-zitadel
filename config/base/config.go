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
	// OrganizationRef references an organization by id
	OrganizationRef = config.Reference{
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

	// MachineUserRef references a machine user by id
	MachineUserRef = config.Reference{
		TerraformName: "zitadel_machine_user",
		Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)`,
	}

	// ApplicationApiRef references a api application by id
	ApplicationApiRef = config.Reference{
		TerraformName: "zitadel_application_api",
		Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id", true)`,
	}
)

// Configure configures the base provider.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("zitadel_org", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Organization"
	})

	p.AddResourceConfigurator("zitadel_action", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Action"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_application_api", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ApplicationApi"

		r.References["org_id"] = OrganizationRef
		r.References["project_id"] = ProjectRef
	})

	p.AddResourceConfigurator("zitadel_application_key", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ApplicationKey"

		r.References["org_id"] = OrganizationRef
		r.References["project_id"] = ProjectRef
		r.References["app_id"] = ApplicationApiRef
	})

	p.AddResourceConfigurator("zitadel_application_oidc", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ApplicationOidc"

		r.References["org_id"] = OrganizationRef
		r.References["project_id"] = ProjectRef
	})

	p.AddResourceConfigurator("zitadel_application_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ApplicationSaml"

		r.References["org_id"] = OrganizationRef
		r.References["project_id"] = ProjectRef
	})

	p.AddResourceConfigurator("zitadel_default_domain_claimed_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultDomainClaimedMessageText"
	})

	p.AddResourceConfigurator("zitadel_default_domain_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultDomainPolicy"
	})

	p.AddResourceConfigurator("zitadel_default_init_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultInitMessageText"
	})

	p.AddResourceConfigurator("zitadel_default_label_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultLabelPolicy"
	})

	p.AddResourceConfigurator("zitadel_default_lockout_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultLockoutPolicy"
	})

	p.AddResourceConfigurator("zitadel_default_login_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultLoginPolicy"
	})

	// p.AddResourceConfigurator("zitadel_default_login_texts", func(r *config.Resource) {
	// 	r.ShortGroup = shortGroup
	// 	r.Kind = "DefaultLoginTexts"
	// })

	p.AddResourceConfigurator("zitadel_default_notification_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultNotificationPolicy"
	})
	p.AddResourceConfigurator("zitadel_default_oidc_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultOidcSettings"
	})
	p.AddResourceConfigurator("zitadel_default_password_age_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultPasswordAgePolicy"
	})
	p.AddResourceConfigurator("zitadel_default_password_change_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultPasswordChangeMessageText"
	})
	p.AddResourceConfigurator("zitadel_default_password_complexity_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultPasswordComplexityPolicy"
	})
	p.AddResourceConfigurator("zitadel_default_password_reset_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultPasswordResetMessageText"
	})
	p.AddResourceConfigurator("zitadel_default_passwordless_registration_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultPasswordlessRegistrationMessageText"
	})
	p.AddResourceConfigurator("zitadel_default_privacy_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultPrivacyPolicy"
	})
	p.AddResourceConfigurator("zitadel_default_verify_email_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultVerifyEmailMessageText"
	})
	p.AddResourceConfigurator("zitadel_default_verify_email_otp_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultVerifyEmailOtpMessageText"
	})
	p.AddResourceConfigurator("zitadel_default_verify_phone_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultVerifyPhoneMessageText"
	})
	p.AddResourceConfigurator("zitadel_default_verify_sms_otp_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DefaultVerifySmsOtpMessageText"
	})

	p.AddResourceConfigurator("zitadel_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Domain"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_domain_claimed_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DomainClaimedMessageText"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_domain_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DomainPoliy"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_human_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "HumanUser"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_idp_azure_ad", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IdpAzureAd"
	})
	p.AddResourceConfigurator("zitadel_idp_github", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IdpGithub"
	})
	p.AddResourceConfigurator("zitadel_idp_github_es", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IdpGithubEs"
	})
	p.AddResourceConfigurator("zitadel_idp_gitlab", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IdpGitlab"
	})
	p.AddResourceConfigurator("zitadel_idp_gitlab_self_hosted", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IdpGitlabSelfHosted"
	})
	p.AddResourceConfigurator("zitadel_idp_google", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IdpGoogle"
	})
	p.AddResourceConfigurator("zitadel_idp_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IdpLdap"
	})
	p.AddResourceConfigurator("zitadel_idp_oauth", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IdpOauth"
	})
	p.AddResourceConfigurator("zitadel_idp_oidc", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IdpOidc"
	})
	p.AddResourceConfigurator("zitadel_idp_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IdpSaml"
	})

	p.AddResourceConfigurator("zitadel_init_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "InitMessageText"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_instance_member", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "InstanceMember"

		r.References["user_id"] = HumanUserRef
	})

	p.AddResourceConfigurator("zitadel_label_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LabelPolicy"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_lockout_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LockoutPolicy"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_login_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LoginPolicy"

		r.References["org_id"] = OrganizationRef
	})

	// p.AddResourceConfigurator("zitadel_login_texts", func(r *config.Resource) {
	// 	r.ShortGroup = shortGroup
	// 	r.Kind = "LoginTexts"

	// 	r.References["org_id"] = OrganizationRef
	// })

	p.AddResourceConfigurator("zitadel_machine_key", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MachineKey"

		r.References["org_id"] = OrganizationRef
		r.References["user_id"] = MachineUserRef
	})

	p.AddResourceConfigurator("zitadel_machine_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MachineUser"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_notification_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NotificationPolicy"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_org_idp_azure_ad", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationIdpAzureAd"

		r.References["org_ird"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_org_idp_github", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationIdpGithub"

		r.References["org_ird"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_org_idp_github_es", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationIdpGithubEs"

		r.References["org_ird"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_org_idp_gitlab", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationIdpGitlab"

		r.References["org_ird"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_org_idp_gitlab_self_hosted", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationIdpGitlabSelfHosted"

		r.References["org_ird"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_org_idp_google", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationIdpGoogle"

		r.References["org_ird"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_org_idp_jwt", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationIdpJwt"

		r.References["org_ird"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_org_idp_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationIdpLdap"

		r.References["org_ird"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_org_idp_oauth", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationIdpOauth"

		r.References["org_ird"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_org_idp_oidc", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationIdpOidc"

		r.References["org_ird"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_org_idp_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationIdpSaml"

		r.References["org_ird"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_org_member", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationMember"

		r.References["org_id"] = OrganizationRef
		r.References["user_id"] = HumanUserRef
	})

	p.AddResourceConfigurator("zitadel_org_metadata", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrganizationMetadata"

		r.References["org_ird"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_password_age_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PasswordAgePolicy"

		r.References["org_id"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_password_change_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PasswordChangeMessageText"

		r.References["org_id"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_password_complexity_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PasswordComplexityPolicy"

		r.References["org_id"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_password_reset_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PasswordResetMessageText"

		r.References["org_id"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_passwordless_registration_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PasswordlessRegistrationMessageText"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_personal_access_token", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PersonalAccessToken"

		r.References["org_id"] = OrganizationRef
		r.References["user_id"] = MachineUserRef
	})

	p.AddResourceConfigurator("zitadel_privacy_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PrivacyPolicy"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_project", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Project"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_project_grant", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectGrant"

		r.References["org_id"] = OrganizationRef
		r.References["granted_org_id"] = OrganizationRef
		r.References["project_id"] = ProjectRef
	})

	p.AddResourceConfigurator("zitadel_project_grant_member", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectGrantMember"

		r.References["org_id"] = OrganizationRef
		r.References["project_id"] = ProjectRef
		r.References["user_id"] = HumanUserRef
	})

	p.AddResourceConfigurator("zitadel_project_member", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectMember"

		r.References["org_id"] = OrganizationRef
		r.References["project_id"] = ProjectRef
		r.References["user_id"] = HumanUserRef
	})

	p.AddResourceConfigurator("zitadel_project_role", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProjectRole"

		r.References["org_id"] = OrganizationRef
		r.References["project_id"] = ProjectRef
	})

	p.AddResourceConfigurator("zitadel_sms_provider_http", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SmsProviderHttp"
	})
	p.AddResourceConfigurator("zitadel_sms_provider_twilio", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SmsProviderTwilio"
	})
	p.AddResourceConfigurator("zitadel_smtp_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SmtpConfig"
	})

	p.AddResourceConfigurator("zitadel_trigger_actions", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TriggerActions"

		r.References["org_id"] = OrganizationRef
	})

	p.AddResourceConfigurator("zitadel_user_grant", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UserGrant"

		r.References["org_id"] = OrganizationRef
		r.References["user_id"] = HumanUserRef
		r.References["project_id"] = ProjectRef
	})

	p.AddResourceConfigurator("zitadel_user_metadata", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UserMetadata"

		r.References["org_id"] = OrganizationRef
		r.References["user_id"] = HumanUserRef
	})

	p.AddResourceConfigurator("zitadel_verify_email_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VerifyEmailMessageText"

		r.References["org_id"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_verify_email_otp_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FyEmailOtpMessageText"

		r.References["org_id"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_verify_phone_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VerifyPhoneMessageText"

		r.References["org_id"] = OrganizationRef
	})
	p.AddResourceConfigurator("zitadel_verify_sms_otp_message_text", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RifySmsOtpMessageText"

		r.References["org_id"] = OrganizationRef
	})

}
