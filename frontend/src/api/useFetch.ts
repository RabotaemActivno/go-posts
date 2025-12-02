export enum ResponseMethods {
    Get = "GET",
    Post = "POST",
}

export enum StatusCode {
    Error = "Error",
    OK = "OK"
}

export type AddPostBody = {
    author: string,
    text: string
}

export async function preparedFetch<T>(method: ResponseMethods, body?: AddPostBody): Promise<T> {
    const apiRoute = "/api/posts";
    const stringifyBody = body ? JSON.stringify(body) : null;
    let data
    try {
        const res = await fetch(apiRoute, {
            method,
            body: stringifyBody
        });
        data = await res.json();
    } catch(err) {
        console.log(err);
        data = null;
    }
    return data;
}