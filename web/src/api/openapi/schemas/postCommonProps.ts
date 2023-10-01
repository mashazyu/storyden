/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { AssetList } from "./assetList";
import type { Identifier } from "./identifier";
import type { Metadata } from "./metadata";
import type { PostContent } from "./postContent";
import type { ProfileReference } from "./profileReference";
import type { ReactList } from "./reactList";
import type { ThreadMark } from "./threadMark";
import type { Url } from "./url";

export interface PostCommonProps {
  root_id: Identifier;
  root_slug: ThreadMark;
  body: PostContent;
  author: ProfileReference;
  meta?: Metadata;
  reacts: ReactList;
  reply_to?: Identifier;
  assets: AssetList;
  url?: Url;
}
