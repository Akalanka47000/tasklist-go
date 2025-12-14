// Package masker provides utilities for masking sensitive information in data structures based on struct tags.
package masker

import "github.com/showa-93/go-mask"

var masker = mask.NewMasker()

func init() {
	masker.SetMaskChar("*")
	masker.RegisterMaskStringFunc(mask.MaskTypeFilled, masker.MaskFilledString)
}

// Mask applies masking to the given value based on struct tags.
func Mask(v any) (any, error) {
	return masker.Mask(v)
}

// MustMask applies masking to the given value and panics if an error occurs.
func MustMask(v any) any {
	masked, err := Mask(v)
	if err != nil {
		panic(err)
	}
	return masked
}
