const getAllButton = document.getElementById("get-posts");
const authorInput = document.getElementById("author");
const textInput = document.getElementById("text");
const createPostButton = document.getElementById("add-post");

if (!getAllButton || !authorInput || !textInput || !createPostButton) {
    console.warn("html element is not defined");
}
createPostButton.addEventListener("click", async () => {
    const author = authorInput.value;
    const text = textInput.value;
    if (!author || !text) {
        console.log('inputs are can not be empty');
        return;
    }

    const res = await fetch("/api/posts", {
        method: "POST",
        body: JSON.stringify({
            author,
            text
        })
    });

    const data = await res.json()
    console.log(data);
});

getAllButton.addEventListener("click", async () => {
    const res = await fetch("/api/posts", {
        method: "GET"
    })
    const data = await res.json()
    console.log(data);
});