/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	providerconfig "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/providerconfig"
	action "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/action"
	domain "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/domain"
	humanuser "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/humanuser"
	machineuser "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/machineuser"
	organisation "github.com/vhdirk/crossplane-provider-zitadel/internal/controller/zitadel/organisation"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		action.Setup,
		domain.Setup,
		humanuser.Setup,
		machineuser.Setup,
		organisation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
