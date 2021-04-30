///ws/{name:[\w\W]}
console.log(window.location.pathname.split("/")[2]);
var ws = new WebSocket("ws://" + window.location.host + "/ws/" + window.location.pathname.split("/")[2]);
var sendMsg = function () {
    ws.send(JSON.stringify({
        author: document.getElementById("author").value,
        message: document.getElementById("message").value
    }));
};
ws.onmessage = function (event) {
    console.log(JSON.parse(event.data));
};
