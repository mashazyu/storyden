/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { PublicKeyCredentialType } from "./publicKeyCredentialType";

/**
 * https://www.w3.org/TR/webauthn-2/#dictdef-publickeycredentialparameters

 */
export interface PublicKeyCredentialParameters {
  type: PublicKeyCredentialType;
  alg: number;
}
