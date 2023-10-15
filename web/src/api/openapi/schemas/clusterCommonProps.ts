/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { AssetURL } from "./assetURL";
import type { ClusterDescription } from "./clusterDescription";
import type { ClusterName } from "./clusterName";
import type { ProfileReference } from "./profileReference";
import type { Properties } from "./properties";
import type { Slug } from "./slug";

/**
 * The main properties of a cluster.
 */
export interface ClusterCommonProps {
  name: ClusterName;
  slug: Slug;
  image_url?: AssetURL;
  description: ClusterDescription;
  owner: ProfileReference;
  properties: Properties;
}
