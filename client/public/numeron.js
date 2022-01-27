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
    showNumeron(display_id).then(result => {
        if (result === null) {
            return null
        }

        let first_id = result['users'][0]['user_id'];
        let second_id = result['users'][1]['user_id'];
        const url = 'http://localhost:8082/api/numerons/' + display_id + '/start';
        const data = JSON.stringify({
            first: first_id,
            second: second_id,
        });
        const xhr = new XMLHttpRequest()
        xhr.withCredentials = true
        xhr.open('POST', url);
        xhr.setRequestHeader('content-type', 'application/json');
        xhr.send(data);

        xhr.addEventListener('readystatechange', function () {
            if (this.readyState === this.DONE) {
                if (this.status === 200) {

                }
            }
        });
    })
}

function setCodeNumeron(display_id) {
    const codeError = document.getElementById("code-error");
    let code = document.getElementById("code").value;

    const url = 'http://localhost:8082/api/numerons/' + display_id + '/code';
    const data = JSON.stringify({
        code: code,
    });
    const xhr = new XMLHttpRequest()
    xhr.withCredentials = true
    xhr.open('POST', url);
    xhr.setRequestHeader('content-type', 'application/json');
    xhr.send(data);

    xhr.addEventListener('readystatechange', function () {
        if (this.readyState === this.DONE) {
            if (this.status !== 200) {
                codeError.innerText = this.responseText;
            }
        }
    });
}