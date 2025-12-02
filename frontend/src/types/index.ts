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