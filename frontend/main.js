console.log('hello');

const button = document.getElementById("get-posts");
if (!button) {
    console.warn("кнопку не нашли");
}
button.addEventListener("click", async () => {
    const res = await fetch("/api/posts", {
        method: "GET"
    })
    const data = await res.json()
    console.log(data);
})