import { useEffect, useState } from "react"
import { preparedFetch, ResponseMethods, StatusCode, type AddPostBody } from "./api/api"
import List from "./components/list/List"
import CreatePostModal from "./components/modal/CreatePostModal"
import Navbar from "./layouts/Navbar"
import type { CreatePostResponse, Post, ResponseData } from "./types";


function App() {
  const [posts, setPosts] = useState<Post[]>([]);
  const [isCreateModalOpen, setCreateModalOpen] = useState(false);
  const [isCreating, setIsCreating] = useState(false);
  const [createError, setCreateError] = useState<string | null>(null);

  useEffect(() => {
    async function loadPosts() {
      const data = await preparedFetch<ResponseData>(ResponseMethods.Get);
      if (data?.status === StatusCode.OK && Array.isArray(data.posts)) {
        setPosts(data.posts);
      }
    }
    loadPosts();
  }, [])

  const handlerUpdatePost = async (id: number) => {
    const payload = {
      author: values.author.trim(),
      text: values.text.trim(),
    };

    if (!payload.author || !payload.text) {
      setCreateError("Заполните автора и текст поста");
      return;
    }

    setIsCreating(true);
    setCreateError(null);


    const response = await preparedFetch<CreatePostResponse>(ResponseMethods.Patch, `/api/posts/${id}`, payload);
  } 

  const handleCreatePost = async (values: AddPostBody) => {
    const payload = {
      author: values.author.trim(),
      text: values.text.trim(),
    };

    if (!payload.author || !payload.text) {
      setCreateError("Заполните автора и текст поста");
      return;
    }

    setIsCreating(true);
    setCreateError(null);

    const response = await preparedFetch<CreatePostResponse>(ResponseMethods.Post, "/api/posts", payload);

    if (response?.status === StatusCode.OK && typeof response.postID === "number") {
      setPosts((prev) => [...prev, { id: response.postID, ...payload }]);
      setCreateModalOpen(false);
    } else {
      setCreateError(response?.text ?? "Не удалось создать пост");
    }

    setIsCreating(false);
  };

  const handlerRemovePost = async (id: number) => {
    const response = await preparedFetch<any>(ResponseMethods.Delete, `/api/posts/${id}`)
    if (response?.status === StatusCode.OK) {
      const newPosts = posts.filter(post => post.id !== id);
      setPosts([...newPosts])
    }
  }

  return (
    <div>
      <Navbar onCreateClick={() => {
        setCreateError(null);
        setCreateModalOpen(true);
      }} />
      <List 
        posts={posts}
        handlerRemovePost = {handlerRemovePost}
        handlerUpdatePost = {handlerUpdatePost}
      />
      <CreatePostModal
          open={isCreateModalOpen}
          onClose={() => {
            setCreateError(null);
            setCreateModalOpen(false);
          }}
          onSubmit={handleCreatePost}
          isSubmitting={isCreating}
          error={createError}
      />
    </div>
  )
}

export default App
