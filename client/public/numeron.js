function getNumeron(display_id) {
    if (display_id === "") {
        return
    }
    const name = document.getElementById('name');
    const status = document.getElementById('status');
    const participants = document.getElementById('participants');
    const url = 'http://localhost:8082/api/numerons/'+display_id;
    const xhr = new XMLHttpRequest()
    xhr.withCredentials = true
    xhr.open('GET', url);
    xhr.setRequestHeader('content-type', 'application/json');
    xhr.send();

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

    xhr.addEventListener('readystatechange', function() {
        let users;

        if (this.readyState === this.DONE) {
            if (this.status === 200) {
                name.innerHTML = JSON.parse(this.responseText)['name'];
                status.innerHTML = statusToString(JSON.parse(this.responseText)['status']);
                users = JSON.parse(this.responseText)['users'];
                if (users != null) {
                    let participant_content = ''
                    for (let i = 0; i < users.length; i++) {
                        participant_content +=
                            '<ul>参加者' + (i+1) + ': ' + users[i] + '</ul>';
                    }
                    participants.innerHTML = participant_content;
                }
            }
        }
    });
}

function showNumeron(display_id) {
    return new Promise((resolve, reject) => {
        if (display_id === "") {
            return
        }
        const url = 'http://localhost:8082/api/numerons/' + display_id;
        const xhr = new XMLHttpRequest()
        xhr.withCredentials = true
        xhr.open('GET', url);
        xhr.setRequestHeader('content-type', 'application/json');
        xhr.send();

        xhr.addEventListener('readystatechange', function () {
            if (this.readyState === this.DONE) {
                if (this.status === 200) {
                    return resolve(JSON.parse(this.responseText));
                } else {
                    return resolve(null);
                }
            }
        });
    })
}

function leaveNumeron(display_id) {
    if (display_id === "") {
        return
    }

    const url = 'http://localhost:8082/api/numerons/'+display_id+'/leave';
    const xhr = new XMLHttpRequest()
    xhr.withCredentials = true
    xhr.open('POST', url);
    xhr.setRequestHeader('content-type', 'application/json');
    xhr.send();

    xhr.addEventListener('readystatechange', function() {
        if (this.readyState === this.DONE) {
            if (this.status === 200) {
                window.location.href = 'http://localhost:3000/numerons.html';
            }
        }
    });
}

function startNumeron(display_id) {
    if (display_id === "") {
        return
    }

    let numeron = showNumeron(display_id)
    if (numeron === null) {
        console.log("Error: showNumeron")
        return null
    }

    let first_id = numeron['users'][0]['user_id'];
    let second_id = numeron['users'][1]['user_id'];
    const url = 'http://localhost:8082/api/numerons/'+display_id+'/start';
    const data = JSON.stringify({
        first: first_id,
        second: second_id,
    });
    const xhr = new XMLHttpRequest()
    xhr.withCredentials = true
    xhr.open('POST', url);
    xhr.setRequestHeader('content-type', 'application/json');
    xhr.send(data);

    xhr.addEventListener('readystatechange', function() {
        if (this.readyState === this.DONE) {
            if (this.status === 200) {
                console.log("game start!!")
            }
        }
    });
}