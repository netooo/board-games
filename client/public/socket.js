function connectSocket() {
    const url = 'ws://localhost:8082/api/ws/connect';
    let socket = new WebSocket(url);
    const chatContent = document.getElementById("chat-content");

    socket.onopen = function (event) {
        console.log("connected socket");
    };

    socket.onclose = function(event) {
        console.log("connection closed (" + event.code + ")");
        disconnectSocket();
        window.location.href = 'http://localhost:3000/numerons.html';
    };

    socket.onmessage = function(event) {
        let msg = JSON.parse(event.data);
        if (msg['Action'] === 'join') {
            chatContent.innerHTML += msg['Value'] + "が入室しました。";
        } else if (msg['Action'] === 'leave') {
            chatContent.innerHTML += msg['Value'] + "が退出しました。";
        }
    };
}

function disconnectSocket() {
    const url = 'http://localhost:8082/api/ws/disconnect';
    const xhr = new XMLHttpRequest()
    xhr.withCredentials = true
    xhr.open('GET', url);
    xhr.setRequestHeader('content-type', 'application/json');
    xhr.send();
}