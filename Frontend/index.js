document.addEventListener("DOMContentLoaded", () => {
    fetch("/static/uploads/")
        .then(response => response.json())
        .then(videos => {
            const videoList = document.getElementById("videos");
            videos.forEach(video => {
                const link = document.createElement("a");
                link.href = `/watch.html?video=${video}`;
                link.textContent = video;
                videoList.appendChild(link);
            });
        })
        .catch(err => console.error("Error fetching videos:", err));
});
