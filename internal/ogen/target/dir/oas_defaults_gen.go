// Code generated by ogen, DO NOT EDIT.

package ogen

// setDefaults set default value of fields.
func (s *AuthenticatorSelectionCriteria) setDefaults() {
	{
		val := UserVerificationRequirement("preferred")
		s.UserVerification.SetTo(val)
	}
}
