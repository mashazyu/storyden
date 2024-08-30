/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type {
  DatagraphSearchOKResponse,
  DatagraphSearchParams,
} from "../openapi-schema";
import { fetcher } from "../server";

/**
 * Query and search content.
 */
export type datagraphSearchResponse = {
  data: DatagraphSearchOKResponse;
  status: number;
};

export const getDatagraphSearchUrl = (params?: DatagraphSearchParams) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value === null) {
      normalizedParams.append(key, "null");
    } else if (value !== undefined) {
      normalizedParams.append(key, value.toString());
    }
  });

  return `/datagraph?${normalizedParams.toString()}`;
};

export const datagraphSearch = async (
  params?: DatagraphSearchParams,
  options?: RequestInit,
): Promise<datagraphSearchResponse> => {
  return fetcher<Promise<datagraphSearchResponse>>(
    getDatagraphSearchUrl(params),
    {
      ...options,
      method: "GET",
    },
  );
};
