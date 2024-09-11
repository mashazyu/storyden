/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { AuthProvider } from "./authProvider";

/**
 * An authentication method is an active instance of an authentication
provider associated with an account. Use this to display a user's active
authentication methods so they can edit or remove it.

 */
export interface AccountAuthMethod {
  /** When this auth method was registered to the account. */
  created_at: string;
  /** The internal unique ID this method has. */
  id: string;
  /** The external identifier (third party ID or device ID) */
  identifier: string;
  /** The personal name given to the method. */
  name: string;
  provider: AuthProvider;
}