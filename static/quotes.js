async function setup() {
    let con = document.getElementById("quote");
    let likes = document.getElementById("heart");

    let resp = await fetch("/api/quote");
    resp = await resp.json();

    con.innerText = resp["q"];
    con.style.opacity = 1;

    resp = await fetch("/api/like/view", {
        method: "POST",
            body: JSON.stringify({
                "quote": resp["q"]
            })
        }
    )
    resp = await resp.json();
    likes.innerText = resp["likes"]
    
    likes.onclick = async function() {
        let quote = document.getElementById("quote").innerText;
        let heart = document.getElementById("heart");

        let resp = await fetch("/api/like/add", {
            method: "POST",
            body: JSON.stringify({
                "quote": quote
            })
        })

        resp = await resp.json()
        heart.innerText = resp["likes"]
    }
}

setup()
