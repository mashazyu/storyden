/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import useSwr from "swr";
import type { Arguments, Key, SWRConfiguration } from "swr";
import useSWRMutation from "swr/mutation";
import type { SWRMutationConfiguration } from "swr/mutation";

import { fetcher } from "../client";
import type {
  InternalServerErrorResponse,
  NotFoundResponse,
  PostReactAddBody,
  PostReactAddOKResponse,
  PostSearchOKResponse,
  PostSearchParams,
  PostUpdateBody,
  PostUpdateOKResponse,
  UnauthorisedResponse,
} from "../openapi-schema";

/**
 * Publish changes to a single post.
 */
export const postUpdate = (postId: string, postUpdateBody: PostUpdateBody) => {
  return fetcher<PostUpdateOKResponse>({
    url: `/posts/${postId}`,
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    data: postUpdateBody,
  });
};

export const getPostUpdateMutationFetcher = (postId: string) => {
  return (
    _: string,
    { arg }: { arg: PostUpdateBody },
  ): Promise<PostUpdateOKResponse> => {
    return postUpdate(postId, arg);
  };
};
export const getPostUpdateMutationKey = (postId: string) =>
  `/posts/${postId}` as const;

export type PostUpdateMutationResult = NonNullable<
  Awaited<ReturnType<typeof postUpdate>>
>;
export type PostUpdateMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const usePostUpdate = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  postId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof postUpdate>>,
      TError,
      string,
      PostUpdateBody,
      Awaited<ReturnType<typeof postUpdate>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getPostUpdateMutationKey(postId);
  const swrFn = getPostUpdateMutationFetcher(postId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Archive a post using soft-delete.
 */
export const postDelete = (postId: string) => {
  return fetcher<void>({ url: `/posts/${postId}`, method: "DELETE" });
};

export const getPostDeleteMutationFetcher = (postId: string) => {
  return (_: string, __: { arg: Arguments }): Promise<void> => {
    return postDelete(postId);
  };
};
export const getPostDeleteMutationKey = (postId: string) =>
  `/posts/${postId}` as const;

export type PostDeleteMutationResult = NonNullable<
  Awaited<ReturnType<typeof postDelete>>
>;
export type PostDeleteMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const usePostDelete = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  postId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof postDelete>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof postDelete>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getPostDeleteMutationKey(postId);
  const swrFn = getPostDeleteMutationFetcher(postId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Search through posts using various queries and filters.
 */
export const postSearch = (params?: PostSearchParams) => {
  return fetcher<PostSearchOKResponse>({
    url: `/posts/search`,
    method: "GET",
    params,
  });
};

export const getPostSearchKey = (params?: PostSearchParams) =>
  [`/posts/search`, ...(params ? [params] : [])] as const;

export type PostSearchQueryResult = NonNullable<
  Awaited<ReturnType<typeof postSearch>>
>;
export type PostSearchQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const usePostSearch = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  params?: PostSearchParams,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof postSearch>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getPostSearchKey(params) : null));
  const swrFn = () => postSearch(params);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  );

  return {
    swrKey,
    ...query,
  };
};
/**
 * Add a reaction to a post.
 */
export const postReactAdd = (
  postId: string,
  postReactAddBody: PostReactAddBody,
) => {
  return fetcher<PostReactAddOKResponse>({
    url: `/posts/${postId}/reacts`,
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    data: postReactAddBody,
  });
};

export const getPostReactAddMutationFetcher = (postId: string) => {
  return (
    _: string,
    { arg }: { arg: PostReactAddBody },
  ): Promise<PostReactAddOKResponse> => {
    return postReactAdd(postId, arg);
  };
};
export const getPostReactAddMutationKey = (postId: string) =>
  `/posts/${postId}/reacts` as const;

export type PostReactAddMutationResult = NonNullable<
  Awaited<ReturnType<typeof postReactAdd>>
>;
export type PostReactAddMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const usePostReactAdd = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  postId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof postReactAdd>>,
      TError,
      string,
      PostReactAddBody,
      Awaited<ReturnType<typeof postReactAdd>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getPostReactAddMutationKey(postId);
  const swrFn = getPostReactAddMutationFetcher(postId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
