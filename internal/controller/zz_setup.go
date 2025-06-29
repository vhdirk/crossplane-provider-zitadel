/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	providerconfig "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/providerconfig"
	action "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/action"
	applicationapi "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/applicationapi"
	applicationkey "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/applicationkey"
	applicationoidc "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/applicationoidc"
	applicationsaml "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/applicationsaml"
	defaultdomainclaimedmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultdomainclaimedmessagetext"
	defaultdomainpolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultdomainpolicy"
	defaultinitmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultinitmessagetext"
	defaultlabelpolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultlabelpolicy"
	defaultlockoutpolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultlockoutpolicy"
	defaultloginpolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultloginpolicy"
	defaultnotificationpolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultnotificationpolicy"
	defaultoidcsettings "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultoidcsettings"
	defaultpasswordchangemessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultpasswordchangemessagetext"
	defaultpasswordcomplexitypolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultpasswordcomplexitypolicy"
	defaultpasswordlessregistrationmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultpasswordlessregistrationmessagetext"
	defaultpasswordresetmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultpasswordresetmessagetext"
	defaultprivacypolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultprivacypolicy"
	defaultverifyemailmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultverifyemailmessagetext"
	defaultverifyemailotpmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultverifyemailotpmessagetext"
	defaultverifyphonemessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultverifyphonemessagetext"
	defaultverifysmsotpmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/defaultverifysmsotpmessagetext"
	domain "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/domain"
	domainclaimedmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/domainclaimedmessagetext"
	domainpoliy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/domainpoliy"
	fyemailotpmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/fyemailotpmessagetext"
	humanuser "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/humanuser"
	idpazuread "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/idpazuread"
	idpgithub "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/idpgithub"
	idpgithubes "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/idpgithubes"
	idpgitlab "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/idpgitlab"
	idpgitlabselfhosted "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/idpgitlabselfhosted"
	idpgoogle "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/idpgoogle"
	idpldap "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/idpldap"
	idpoauth "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/idpoauth"
	idpoidc "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/idpoidc"
	idpsaml "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/idpsaml"
	initmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/initmessagetext"
	instancemember "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/instancemember"
	labelpolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/labelpolicy"
	lockoutpolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/lockoutpolicy"
	loginpolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/loginpolicy"
	machinekey "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/machinekey"
	machineuser "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/machineuser"
	notificationpolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/notificationpolicy"
	organisation "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisation"
	organisationidpazuread "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationidpazuread"
	organisationidpgithub "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationidpgithub"
	organisationidpgithubes "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationidpgithubes"
	organisationidpgitlab "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationidpgitlab"
	organisationidpgitlabselfhosted "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationidpgitlabselfhosted"
	organisationidpgoogle "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationidpgoogle"
	organisationidpjwt "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationidpjwt"
	organisationidpldap "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationidpldap"
	organisationidpoauth "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationidpoauth"
	organisationidpoidc "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationidpoidc"
	organisationidpsaml "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationidpsaml"
	organisationmember "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationmember"
	organisationmetadata "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisationmetadata"
	passwordchangemessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/passwordchangemessagetext"
	passwordcomplexitypolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/passwordcomplexitypolicy"
	passwordlessregistrationmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/passwordlessregistrationmessagetext"
	passwordresetmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/passwordresetmessagetext"
	personalaccesstoken "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/personalaccesstoken"
	privacypolicy "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/privacypolicy"
	project "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/project"
	projectgrant "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/projectgrant"
	projectgrantmember "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/projectgrantmember"
	projectmember "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/projectmember"
	projectrole "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/projectrole"
	rifysmsotpmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/rifysmsotpmessagetext"
	smsprovidertwilio "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/smsprovidertwilio"
	smtpconfig "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/smtpconfig"
	triggeractions "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/triggeractions"
	usergrant "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/usergrant"
	usermetadata "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/usermetadata"
	verifyemailmessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/verifyemailmessagetext"
	verifyphonemessagetext "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/verifyphonemessagetext"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		action.Setup,
		applicationapi.Setup,
		applicationkey.Setup,
		applicationoidc.Setup,
		applicationsaml.Setup,
		defaultdomainclaimedmessagetext.Setup,
		defaultdomainpolicy.Setup,
		defaultinitmessagetext.Setup,
		defaultlabelpolicy.Setup,
		defaultlockoutpolicy.Setup,
		defaultloginpolicy.Setup,
		defaultnotificationpolicy.Setup,
		defaultoidcsettings.Setup,
		defaultpasswordchangemessagetext.Setup,
		defaultpasswordcomplexitypolicy.Setup,
		defaultpasswordlessregistrationmessagetext.Setup,
		defaultpasswordresetmessagetext.Setup,
		defaultprivacypolicy.Setup,
		defaultverifyemailmessagetext.Setup,
		defaultverifyemailotpmessagetext.Setup,
		defaultverifyphonemessagetext.Setup,
		defaultverifysmsotpmessagetext.Setup,
		domain.Setup,
		domainclaimedmessagetext.Setup,
		domainpoliy.Setup,
		fyemailotpmessagetext.Setup,
		humanuser.Setup,
		idpazuread.Setup,
		idpgithub.Setup,
		idpgithubes.Setup,
		idpgitlab.Setup,
		idpgitlabselfhosted.Setup,
		idpgoogle.Setup,
		idpldap.Setup,
		idpoauth.Setup,
		idpoidc.Setup,
		idpsaml.Setup,
		initmessagetext.Setup,
		instancemember.Setup,
		labelpolicy.Setup,
		lockoutpolicy.Setup,
		loginpolicy.Setup,
		machinekey.Setup,
		machineuser.Setup,
		notificationpolicy.Setup,
		organisation.Setup,
		organisationidpazuread.Setup,
		organisationidpgithub.Setup,
		organisationidpgithubes.Setup,
		organisationidpgitlab.Setup,
		organisationidpgitlabselfhosted.Setup,
		organisationidpgoogle.Setup,
		organisationidpjwt.Setup,
		organisationidpldap.Setup,
		organisationidpoauth.Setup,
		organisationidpoidc.Setup,
		organisationidpsaml.Setup,
		organisationmember.Setup,
		organisationmetadata.Setup,
		passwordchangemessagetext.Setup,
		passwordcomplexitypolicy.Setup,
		passwordlessregistrationmessagetext.Setup,
		passwordresetmessagetext.Setup,
		personalaccesstoken.Setup,
		privacypolicy.Setup,
		project.Setup,
		projectgrant.Setup,
		projectgrantmember.Setup,
		projectmember.Setup,
		projectrole.Setup,
		rifysmsotpmessagetext.Setup,
		smsprovidertwilio.Setup,
		smtpconfig.Setup,
		triggeractions.Setup,
		usergrant.Setup,
		usermetadata.Setup,
		verifyemailmessagetext.Setup,
		verifyphonemessagetext.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
