/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { AssetIDs } from "./assetIDs";
import type { AssetSourceList } from "./assetSourceList";
import type { Metadata } from "./metadata";
import type { NodeName } from "./nodeName";
import type { NullableIdentifier } from "./nullableIdentifier";
import type { PostContent } from "./postContent";
import type { Slug } from "./slug";
import type { Url } from "./url";

/**
 * Note: Properties are replace-all and are not merged with existing.

 */
export interface NodeMutableProps {
  asset_ids?: AssetIDs;
  asset_sources?: AssetSourceList;
  content?: PostContent;
  meta?: Metadata;
  name?: NodeName;
  parent?: Slug;
  primary_image_asset_id?: NullableIdentifier;
  slug?: Slug;
  url?: Url;
}
