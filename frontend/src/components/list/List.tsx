import type { Post } from "../../types"
import ListItem from "./ListItem"

interface ListProps {
    posts: Post[]
}

function List({posts}: ListProps) {
    return(
        <ul className="list bg-base-100 rounded-box shadow-md w-3/5 mx-auto mt-8">
            {
                posts.map(post => (
                    <ListItem
                        key={post.id}
                        post={post}
                    />
                ))
            }
        </ul>
    )
}

export default List
