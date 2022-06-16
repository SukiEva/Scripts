addEventListener('fetch', event => {
    event.respondWith(
        handleRequest(event.request).catch((err) =>
            new Response('cfworker error:\n' + err.stack, {
                status: 502,
            })
        )
    );
});

async function handleRequest(request) {
    const url = new URL(request.url);
    switch (url.pathname) {
        case "/img":
            var type = url.searchParams.has("type") ? url.searchParams.get("type") : "pc";
            const total = getTotal(type);
            if (total == 0) return handleImage("pc", getTotal("pc"));
            return handleImage(type, total);
        case "/favicon.ico":
            return fetch("https://cdn.jsdelivr.net/gh/SukiEva/assets/blog/favicon.ico");
        default:
            return handleNotFound();
    }
}

function getTotal(type) {
    switch (type) {
        case "pc": return 175;
        case "mb": return 0;
        default: return 0;
    }
}

async function handleImage(type, total) {
    var index = Math.floor((Math.random() * total)) + 1;
    var img = "https://cdn.jsdelivr.net/gh/SukiEva/assets/webp/" + type + "/" + index + ".webp";
    res = await fetch(img);
    return new Response(res.body, {
        headers: {
            'content-type': 'image/webp',
        },
    });
}

function handleNotFound() {
    return new Response('Not Found', {
        status: 404,
    });
}