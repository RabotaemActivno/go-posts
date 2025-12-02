import { useEffect, useState } from "react"
import { preparedFetch, ResponseMethods, StatusCode } from "./api/api"
import List from "./components/list/List"
import CreatePostModal from "./components/modal/CreatePostModal"
import Navbar from "./layouts/Navbar"
import type { Post, ResponseData } from "./types";


function App() {
  const [posts, setPosts] = useState<Post[]>([]);
  const [isCreateModalOpen, setCreateModalOpen] = useState(false);

  useEffect(() => {
    async function loadPosts() {
      const data = await preparedFetch<ResponseData>(ResponseMethods.Get);
      if (data.status === StatusCode.OK) {
        data.posts = data.posts || [];
        setPosts(data.posts);
      }
    }
    loadPosts();
  }, [])

  return (
    <div>
      <Navbar onCreateClick={() => setCreateModalOpen(true)} />
      <List posts={posts}/>
      <CreatePostModal
        open={isCreateModalOpen}
        onClose={() => setCreateModalOpen(false)}
      />
    </div>
  )
}

export default App
