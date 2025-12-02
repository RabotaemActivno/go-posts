
type CreatePostModalProps = {
  open: boolean;
  onClose: () => void;
}

function CreatePostModal({ open, onClose }: CreatePostModalProps) {
  if (!open) return null;

  return (
    <dialog id="my_modal_1" className="modal" open>
        <div className="modal-box">
            <h3 className="font-bold text-lg">Hello!</h3>
            <p className="py-4">Press ESC key or click the button below to close</p>
            <div className="modal-action">
                <form method="dialog" onSubmit={onClose}>
                    <input type="text" placeholder="Type here" className="input"/>
                    <input type="text" placeholder="Type here" className="input"/>
                    <button className="btn" type="submit">Создать</button>
                </form>
            </div>
        </div>
    </dialog>
  )
}

export default CreatePostModal
