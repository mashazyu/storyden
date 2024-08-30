/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type {
  ThreadCreateBody,
  ThreadCreateOKResponse,
  ThreadGetResponse,
  ThreadListOKResponse,
  ThreadListParams,
  ThreadUpdateBody,
  ThreadUpdateOKResponse,
} from "../openapi-schema";
import { fetcher } from "../server";

/**
 * Create a new thread within the specified category.
 */
export type threadCreateResponse = {
  data: ThreadCreateOKResponse;
  status: number;
};

export const getThreadCreateUrl = () => {
  return `/threads`;
};

export const threadCreate = async (
  threadCreateBody: ThreadCreateBody,
  options?: RequestInit,
): Promise<threadCreateResponse> => {
  return fetcher<Promise<threadCreateResponse>>(getThreadCreateUrl(), {
    ...options,
    method: "POST",
    body: JSON.stringify(threadCreateBody),
  });
};

/**
 * Get a list of all threads.
 */
export type threadListResponse = {
  data: ThreadListOKResponse;
  status: number;
};

export const getThreadListUrl = (params?: ThreadListParams) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value === null) {
      normalizedParams.append(key, "null");
    } else if (value !== undefined) {
      normalizedParams.append(key, value.toString());
    }
  });

  return `/threads?${normalizedParams.toString()}`;
};

export const threadList = async (
  params?: ThreadListParams,
  options?: RequestInit,
): Promise<threadListResponse> => {
  return fetcher<Promise<threadListResponse>>(getThreadListUrl(params), {
    ...options,
    method: "GET",
  });
};

/**
 * Get information about a thread such as its title, author, when it was
created as well as a list of the posts within the thread.

 * @summary Get information about a thread and the posts within the thread.
 */
export type threadGetResponse = {
  data: ThreadGetResponse;
  status: number;
};

export const getThreadGetUrl = (threadMark: string) => {
  return `/threads/${threadMark}`;
};

export const threadGet = async (
  threadMark: string,
  options?: RequestInit,
): Promise<threadGetResponse> => {
  return fetcher<Promise<threadGetResponse>>(getThreadGetUrl(threadMark), {
    ...options,
    method: "GET",
  });
};

/**
 * Publish changes to a thread.
 */
export type threadUpdateResponse = {
  data: ThreadUpdateOKResponse;
  status: number;
};

export const getThreadUpdateUrl = (threadMark: string) => {
  return `/threads/${threadMark}`;
};

export const threadUpdate = async (
  threadMark: string,
  threadUpdateBody: ThreadUpdateBody,
  options?: RequestInit,
): Promise<threadUpdateResponse> => {
  return fetcher<Promise<threadUpdateResponse>>(
    getThreadUpdateUrl(threadMark),
    {
      ...options,
      method: "PATCH",
      body: JSON.stringify(threadUpdateBody),
    },
  );
};

/**
 * Archive a thread using soft-delete.
 */
export type threadDeleteResponse = {
  data: void;
  status: number;
};

export const getThreadDeleteUrl = (threadMark: string) => {
  return `/threads/${threadMark}`;
};

export const threadDelete = async (
  threadMark: string,
  options?: RequestInit,
): Promise<threadDeleteResponse> => {
  return fetcher<Promise<threadDeleteResponse>>(
    getThreadDeleteUrl(threadMark),
    {
      ...options,
      method: "DELETE",
    },
  );
};
