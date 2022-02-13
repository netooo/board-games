function connectSocket() {
    const url = 'ws://localhost:8082/api/ws/connect';
    let socket = new WebSocket(url);
    const chatContent = document.getElementById("chat-content");

    socket.onopen = function (event) {
    };

    socket.onclose = function(event) {
        console.log("connection closed (" + event.code + ")");
        disconnectSocket();
        window.location.href = 'http://localhost:3000/numerons.html';
    };

    socket.onmessage = function(event) {
        let msg = JSON.parse(event.data);
        const mainContent = document.getElementById("main-content");

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
        } else if (msg['Action'] === 'start') {
            mainContent.innerHTML =
                '<input type="text" id="code" name="code" maxLength="3" value=""><br>' +
                '<button onClick="setCodeNumeron(`' + display_id + '`);">設定</button><br>' +
                '<a id="code-error"></a>';
        } else if (msg['Action'] === 'set_code') {
            mainContent.innerText = "対戦相手の設定をお待ちください";
        } else if (msg['Action'] === 'completed_code') {
            showNumeron(display_id).then(result => {
                let users = result['users'];
                let my = "";
                let opponent = "";
                for (let i = 0; i < users.length; i++) {
                    if (users[i]['user_id'] === msg['Value'] ){
                        my = users[i]['name'];
                    } else {
                        opponent = users[i]['name'];
                    }
                }
                mainContent.innerHTML =
                    '<h2>部屋名: ' + result['name'] + '</h2>' +
                    '<hr>' +
                    '対戦相手: ' + opponent + '<br>' +
                    '宣言コード: ' + '<div id="opponent-call-code"></div><br>' +
                    '宣言ログ:<br>' + '<div id="opponent-log"></div>' +
                    '<hr>' +
                    '自分: ' + my + '<br>' +
                    '宣言コード: ' + '<div id="my-call-code"></div><br>' +
                    '宣言ログ:<br>' + '<div id="my-log"></div>' +
                    '<hr>' +
                    '<input type="text" id="attack-code" name="attack-code" maxLength="3" value="">' +
                    '<button onClick="attackCodeNumeron(`' + display_id + '`);">攻撃</button><br>'
            })
        } else if (msg['Action'] === 'attack') {
            const myCallCode = document.getElementById('my-call-code');
            const myLog = document.getElementById('my-log');
            const opponentCallCode = document.getElementById('opponent-call-code');
            const opponentLog = document.getElementById('opponent-log');

            if (msg['AttackUser'] === msg['UserId']) {
                myCallCode.innerHTML = msg['Code'];
                myLog.innerHTML += msg['Code'] + '(' + msg['Result'] + ')<br>';
            } else {
                opponentCallCode.innerHTML = msg['Code'];
                opponentLog.innerHTML += msg['Code'] + '(' + msg['Result'] + ')<br>';
            }
        }
    }
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

    name.innerHTML = result['name'];
    status.innerHTML = statusToString(result['status']);
    let users = result['users'];
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