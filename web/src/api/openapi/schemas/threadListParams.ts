/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { AccountHandle } from "./accountHandle";
import type { TagListIDs } from "./tagListIDs";
import type { CategoryNameList } from "./categoryNameList";

export type ThreadListParams = {
  /**
   * Show only results creeated by this user.
   */
  author?: AccountHandle;
  /**
   * Show only results with these tags
   */
  tags?: TagListIDs;
  /**
   * Show only results with these categories
   */
  categories?: CategoryNameList;
};
