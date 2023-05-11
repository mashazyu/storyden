/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { AuthenticatorAttachment } from "./authenticatorAttachment";
import type { ResidentKeyRequirement } from "./residentKeyRequirement";
import type { UserVerificationRequirement } from "./userVerificationRequirement";

/**
 * https://www.w3.org/TR/webauthn-2/#dictdef-authenticatorselectioncriteria

 */
export interface AuthenticatorSelectionCriteria {
  authenticatorAttachment: AuthenticatorAttachment;
  residentKey: ResidentKeyRequirement;
  requireResidentKey?: boolean;
  userVerification?: UserVerificationRequirement;
}
