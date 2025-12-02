import type { Post } from "../../types"

interface ListProps {
    post: Post
    handlerRemovePost: (id: number) => void
}

function ListItem({post, handlerRemovePost}: ListProps) {

    const removePost = () => {
        handlerRemovePost(post.id);
    }

    return (
        <li className="list-row bg-primary mb-2">
            <div>
                <div className="list-col-wrap text-l">{post.author}</div>
            </div>
            <p className="list-col-wrap text-s">
                {post.text}
            </p>
            <button className="btn btn-square btn-ghost" onClick={removePost}>
                <svg className="size-[1.2em]" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><g strokeLinejoin="round" strokeLinecap="round" strokeWidth="2" fill="none" stroke="currentColor"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"></path></g></svg>
            </button>
        </li>
    )
}

export default ListItem
