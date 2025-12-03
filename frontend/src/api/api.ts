export enum ResponseMethods {
    Get = "GET",
    Post = "POST",
    Delete = "DELETE",
    Patch = "PATCH",
}

export enum StatusCode {
    Error = "Error",
    OK = "OK"
}

export type AddPostBody = {
    author: string,
    text: string
}

export async function preparedFetch<T>(method: ResponseMethods, apiRoute = "/api/posts", body?: AddPostBody): Promise<T | null> {
    const stringifyBody = body ? JSON.stringify(body) : null;
    const options: RequestInit = {
        method,
    };

    if (stringifyBody) {
        options.body = stringifyBody;
        options.headers = {
            "Content-Type": "application/json",
        };
    }

    let data;
    try {
        const res = await fetch(apiRoute, options);
        if (!res.ok) {
            throw new Error(`Request failed with status ${res.status}`);
        }
        data = await res.json();
    } catch(err) {
        console.log(err);
        data = null;
    }
    return data;
}
