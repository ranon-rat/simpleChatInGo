var channel = document.getElementById("channel");
var button = document.querySelector("#submit");
var redirect = function () {
    location.replace("http://" + window.location.host + "/channel/" + channel.value);
};
