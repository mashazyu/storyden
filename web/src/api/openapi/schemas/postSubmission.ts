/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * storyden
 * OpenAPI spec version: 1
 */
import type { Post } from "./post";

/**
 * A new post within a thread of posts. A post may reply to another post in
the thread by specifying the `reply_to` property. The identifier in the
`reply_to` value must be post within the same thread.

 */
export type PostSubmission = Post;