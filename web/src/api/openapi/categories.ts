/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import useSwr from "swr";
import type { SWRConfiguration, Key } from "swr";
import type {
  CategoriesListOKResponse,
  InternalServerErrorResponse,
} from "./schemas";
import { fetcher } from "../client";

/**
 * Get a list of all categories on the site.
 */
export const categoriesList = () => {
  return fetcher<CategoriesListOKResponse>({
    url: `/v1/categories`,
    method: "get",
  });
};

export const getCategoriesListKey = () => [`/v1/categories`];

export type CategoriesListQueryResult = NonNullable<
  Awaited<ReturnType<typeof categoriesList>>
>;
export type CategoriesListQueryError = InternalServerErrorResponse;

export const useCategoriesList = <
  TError = InternalServerErrorResponse
>(options?: {
  swr?: SWRConfiguration<Awaited<ReturnType<typeof categoriesList>>, TError> & {
    swrKey?: Key;
    enabled?: boolean;
  };
}) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getCategoriesListKey() : null));
  const swrFn = () => categoriesList();

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions
  );

  return {
    swrKey,
    ...query,
  };
};
