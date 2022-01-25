function connectSocket(display_id) {
    if (display_id === "") {
        return
    }
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
            chatContent.innerHTML += msg['Value'] + "が入室しました。<br>";
            showNumeron(display_id).then(result => {
                reflect(result);
            })
        } else if (msg['Action'] === 'leave') {
            chatContent.innerHTML += msg['Value'] + "が退出しました。<br>";
            showNumeron(display_id).then(result => {
                reflect(result);
            })
        }
    };
}

function reflect(result) {
    if (result === null) {
        return null
    }
    const name = document.getElementById('name');
    const status = document.getElementById('status');
    const participants = document.getElementById('participants');

    function statusToString(status) {
        switch (status) {
            case 0:
                return "Ready";
            case 1:
                return "Playing";
            case 2:
                return "Finished";
            default:
                return "Unknown";
        }
    }

    let users;
    name.innerHTML = result['name'];
    status.innerHTML = statusToString(result['status']);
    users = result['users'];
    if (users != null) {
        let participant_content = ''
        for (let i = 0; i < users.length; i++) {
            participant_content +=
                '<ul>参加者' + (i+1) + ': ' + users[i]['name'] + '</ul>';
        }
        participants.innerHTML = participant_content;
    }
}

function disconnectSocket() {
    const url = 'http://localhost:8082/api/ws/disconnect';
    const xhr = new XMLHttpRequest()
    xhr.withCredentials = true
    xhr.open('GET', url);
    xhr.setRequestHeader('content-type', 'application/json');
    xhr.send();
}