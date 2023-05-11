/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { ThreadTitle } from "./threadTitle";
import type { PostBodyMarkdown } from "./postBodyMarkdown";
import type { TagListIDs } from "./tagListIDs";
import type { Metadata } from "./metadata";
import type { Identifier } from "./identifier";

export interface ThreadMutableProps {
  title?: ThreadTitle;
  body?: PostBodyMarkdown;
  tags?: TagListIDs;
  meta?: Metadata;
  category?: Identifier;
}
