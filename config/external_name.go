/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"zitadel_action":                              config.IdentifierFromProvider,
	"zitadel_application_api":                     config.IdentifierFromProvider,
	"zitadel_application_key":                     config.IdentifierFromProvider,
	"zitadel_application_oidc":                    config.IdentifierFromProvider,
	"zitadel_application_saml":                    config.IdentifierFromProvider,
	"zitadel_default_domain_claimed_message_text": config.IdentifierFromProvider,
	"zitadel_default_domain_policy":               config.IdentifierFromProvider,
	"zitadel_default_init_message_text":           config.IdentifierFromProvider,
	"zitadel_default_label_policy":                config.IdentifierFromProvider,
	"zitadel_default_lockout_policy":              config.IdentifierFromProvider,
	"zitadel_default_login_policy":                config.IdentifierFromProvider,
	// "zitadel_default_login_texts":                            config.IdentifierFromProvider,
	"zitadel_default_notification_policy":                    config.IdentifierFromProvider,
	"zitadel_default_oidc_settings":                          config.IdentifierFromProvider,
	"zitadel_default_password_age_policy":                    config.IdentifierFromProvider,
	"zitadel_default_password_change_message_text":           config.IdentifierFromProvider,
	"zitadel_default_password_complexity_policy":             config.IdentifierFromProvider,
	"zitadel_default_password_reset_message_text":            config.IdentifierFromProvider,
	"zitadel_default_passwordless_registration_message_text": config.IdentifierFromProvider,
	"zitadel_default_privacy_policy":                         config.IdentifierFromProvider,
	"zitadel_default_verify_email_message_text":              config.IdentifierFromProvider,
	"zitadel_default_verify_email_otp_message_text":          config.IdentifierFromProvider,
	"zitadel_default_verify_phone_message_text":              config.IdentifierFromProvider,
	"zitadel_default_verify_sms_otp_message_text":            config.IdentifierFromProvider,
	"zitadel_domain":                      config.IdentifierFromProvider,
	"zitadel_domain_claimed_message_text": config.IdentifierFromProvider,
	"zitadel_domain_policy":               config.IdentifierFromProvider,
	"zitadel_human_user":                  config.IdentifierFromProvider,
	"zitadel_idp_azure_ad":                config.IdentifierFromProvider,
	"zitadel_idp_github":                  config.IdentifierFromProvider,
	"zitadel_idp_github_es":               config.IdentifierFromProvider,
	"zitadel_idp_gitlab":                  config.IdentifierFromProvider,
	"zitadel_idp_gitlab_self_hosted":      config.IdentifierFromProvider,
	"zitadel_idp_google":                  config.IdentifierFromProvider,
	"zitadel_idp_ldap":                    config.IdentifierFromProvider,
	"zitadel_idp_oauth":                   config.IdentifierFromProvider,
	"zitadel_idp_oidc":                    config.IdentifierFromProvider,
	"zitadel_idp_saml":                    config.IdentifierFromProvider,
	"zitadel_init_message_text":           config.IdentifierFromProvider,
	"zitadel_instance_member":             config.IdentifierFromProvider,
	"zitadel_label_policy":                config.IdentifierFromProvider,
	"zitadel_lockout_policy":              config.IdentifierFromProvider,
	"zitadel_login_policy":                config.IdentifierFromProvider,
	// "zitadel_login_texts":                                    config.IdentifierFromProvider,
	"zitadel_machine_key":                            config.IdentifierFromProvider,
	"zitadel_machine_user":                           config.IdentifierFromProvider,
	"zitadel_notification_policy":                    config.IdentifierFromProvider,
	"zitadel_org":                                    config.IdentifierFromProvider,
	"zitadel_org_idp_azure_ad":                       config.IdentifierFromProvider,
	"zitadel_org_idp_github":                         config.IdentifierFromProvider,
	"zitadel_org_idp_github_es":                      config.IdentifierFromProvider,
	"zitadel_org_idp_gitlab":                         config.IdentifierFromProvider,
	"zitadel_org_idp_gitlab_self_hosted":             config.IdentifierFromProvider,
	"zitadel_org_idp_google":                         config.IdentifierFromProvider,
	"zitadel_org_idp_jwt":                            config.IdentifierFromProvider,
	"zitadel_org_idp_ldap":                           config.IdentifierFromProvider,
	"zitadel_org_idp_oauth":                          config.IdentifierFromProvider,
	"zitadel_org_idp_oidc":                           config.IdentifierFromProvider,
	"zitadel_org_idp_saml":                           config.IdentifierFromProvider,
	"zitadel_org_member":                             config.IdentifierFromProvider,
	"zitadel_org_metadata":                           config.IdentifierFromProvider,
	"zitadel_password_age_policy":                    config.IdentifierFromProvider,
	"zitadel_password_change_message_text":           config.IdentifierFromProvider,
	"zitadel_password_complexity_policy":             config.IdentifierFromProvider,
	"zitadel_password_reset_message_text":            config.IdentifierFromProvider,
	"zitadel_passwordless_registration_message_text": config.IdentifierFromProvider,
	"zitadel_personal_access_token":                  config.IdentifierFromProvider,
	"zitadel_privacy_policy":                         config.IdentifierFromProvider,
	"zitadel_project":                                config.IdentifierFromProvider,
	"zitadel_project_grant":                          config.IdentifierFromProvider,
	"zitadel_project_grant_member":                   config.IdentifierFromProvider,
	"zitadel_project_member":                         config.IdentifierFromProvider,
	"zitadel_project_role":                           config.IdentifierFromProvider,
	"zitadel_sms_provider_http":                      config.IdentifierFromProvider,
	"zitadel_sms_provider_twilio":                    config.IdentifierFromProvider,
	"zitadel_smtp_config":                            config.IdentifierFromProvider,
	"zitadel_trigger_actions":                        config.IdentifierFromProvider,
	"zitadel_user_grant":                             config.IdentifierFromProvider,
	"zitadel_user_metadata":                          config.IdentifierFromProvider,
	"zitadel_verify_email_message_text":              config.IdentifierFromProvider,
	"zitadel_verify_email_otp_message_text":          config.IdentifierFromProvider,
	"zitadel_verify_phone_message_text":              config.IdentifierFromProvider,
	"zitadel_verify_sms_otp_message_text":            config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
