const form: HTMLFormElement = document.getElementById("form") as HTMLFormElement;
const username: HTMLInputElement = document.getElementById("username") as HTMLInputElement;
const button = document.querySelector("#submit");

if (localStorage.getItem("username")) {
    username.value = localStorage.getItem("username")
}

const redirect = (e: Event) => {
    e.preventDefault();

    const data: FormData = new FormData(form);
    
    if (data.get("channel") === "" || data.get("username") === "" ) {
        document.getElementById("alert").classList.remove("hidden")
        document.getElementById("alert").classList.add("block")
        return;
    } else {
        document.getElementById("alert").classList.remove("block")
        document.getElementById("alert").classList.add("hidden")
    };

    localStorage.setItem("username", data.get("username") as string);
    location.replace("http://"+window.location.host+"/channel/"+data.get("channel"))
}

form.addEventListener("submit", redirect)
