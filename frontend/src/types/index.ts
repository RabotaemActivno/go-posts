import type { StatusCode } from "../api/api"

export type Post = {
    id: number,
  author: string,
  text: string
}

export type ResponseData = {
  status: StatusCode,
  posts: Post[]
}

export type CreatePostResponse =
  | {
      status: StatusCode.OK,
      postID: number,
      text?: string
    }
  | {
      status: StatusCode.Error,
      text: string,
      postID?: number
    };
