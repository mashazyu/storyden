/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { ProfileReference } from "./profileReference";

export interface CollectionItemMetadata {
  /** The time that the item was added to the collection. */
  added_at: string;
  owner: ProfileReference;
}