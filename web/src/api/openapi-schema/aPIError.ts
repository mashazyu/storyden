/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { APIErrorMetadata } from "./aPIErrorMetadata";

/**
 * A description of an error including a human readable message and any
related metadata from the request and associated services.

 */
export interface APIError {
  /** The internal error, not intended for end-user display. */
  error: string;
  /** A human-readable message intended for end-user display. */
  message?: string;
  /** Any additional metadata related to the error. */
  metadata?: APIErrorMetadata;
  /** A suggested action for the user. */
  suggested?: string;
}
