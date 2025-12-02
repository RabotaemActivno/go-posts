import { type FormEvent, useEffect, useState } from "react";

type CreatePostFormValues = {
  author: string;
  text: string;
}

type CreatePostModalProps = {
  open: boolean;
  onClose: () => void;
  onSubmit: (values: CreatePostFormValues) => void | Promise<void>;
  isSubmitting?: boolean;
  error?: string | null;
}

function CreatePostModal({ open, onClose, onSubmit, isSubmitting = false, error }: CreatePostModalProps) {
  const [author, setAuthor] = useState("");
  const [text, setText] = useState("");

  useEffect(() => {
    if (!open) {
      setAuthor("");
      setText("");
    }
  }, [open]);

  if (!open) return null;

  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onSubmit({
      author: author.trim(),
      text: text.trim(),
    });
  };

  const disabled = isSubmitting || !author.trim() || !text.trim();

  return (
    <dialog id="create-post-modal" className="modal" open>
        <div className="modal-box">
            <h3 className="font-bold text-lg">Создать пост</h3>
            <p className="py-2 text-sm opacity-70">Введите автора и текст поста.</p>
            <form className="space-y-4" onSubmit={handleSubmit}>
                <label className="form-control">
                  <span className="label-text">Автор</span>
                  <input
                    type="text"
                    value={author}
                    onChange={(event) => setAuthor(event.target.value)}
                    className="input input-bordered w-full"
                    placeholder="Введите имя автора"
                  />
                </label>
                <label className="form-control">
                  <span className="label-text">Текст</span>
                  <textarea
                    value={text}
                    onChange={(event) => setText(event.target.value)}
                    className="textarea textarea-bordered w-full"
                    placeholder="Что вы хотите сказать?"
                    rows={4}
                  />
                </label>
                {error && <p className="text-sm text-error">{error}</p>}
                <div className="modal-action">
                    <button className="btn" type="button" onClick={onClose} disabled={isSubmitting}>
                      Отмена
                    </button>
                    <button className="btn btn-primary" type="submit" disabled={disabled}>
                      {isSubmitting ? "Сохранение..." : "Создать"}
                    </button>
                </div>
            </form>
        </div>
    </dialog>
  )
}

export default CreatePostModal
