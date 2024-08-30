/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { CollectionCommonProps } from "./collectionCommonProps";
import type { CollectionWithItemsAllOf } from "./collectionWithItemsAllOf";
import type { CommonProperties } from "./commonProperties";

/**
 * The full properties of a collection, for rendering a single collection
somewhere where you can afford to show all the items in the collection.

 */
export type CollectionWithItems = CommonProperties &
  CollectionCommonProps &
  CollectionWithItemsAllOf;
