package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/openshift/installer/pkg/types/powervs"
)

// ValidatePlatform checks that the specified platform is valid.
func ValidatePlatform(p *powervs.Platform, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if p.Region == "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("region"), "region must be specified"))
	}

	return allErrs
}
